package main

import (
	"fmt"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

// Transposes a matrix using shared memory tiling to avoid uncoalesced global writes.
// Each CTA loads a TILE_DIM x TILE_DIM tile into shared memory, then writes it transposed.
func main() {
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM80)

	// Shared memory tile: float[32][33] — +1 column to avoid bank conflicts
	smem := builder.NewGlobalArray("tile", ptx.Shared, ptx.F32, 32*33).
		WithAlign(4)
	mod.AddGlobal(smem)

	kernel := mod.NewKernel("matrix_transpose")
	kernel.AddDirective(builder.MaxNTid(32, 32, 1))
	kernel.AddParam(builder.NewPtrParam("dst", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewPtrParam("src", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewParam("width", ptx.U32))
	kernel.AddParam(builder.NewParam("height", ptx.U32))

	// --- Registers ---
	tidX   := kernel.NewReg("tid_x", ptx.U32)
	tidY   := kernel.NewReg("tid_y", ptx.U32)
	ctaX   := kernel.NewReg("cta_x", ptx.U32)
	ctaY   := kernel.NewReg("cta_y", ptx.U32)
	// global input coords
	gx     := kernel.NewReg("gx", ptx.U32)
	gy     := kernel.NewReg("gy", ptx.U32)
	// shared mem index (row-major into tile[33-stride])
	smIdx  := kernel.NewReg("sm_idx", ptx.U32)
	smIdx64 := kernel.NewReg("sm_idx64", ptx.U64)
	smBase := kernel.NewReg("sm_base", ptx.U64)
	smAddr := kernel.NewReg("sm_addr", ptx.U64)
	// global src/dst linear index
	srcIdx := kernel.NewReg("src_idx", ptx.U32)
	srcIdx64 := kernel.NewReg("src_idx64", ptx.U64)
	srcOff := kernel.NewReg("src_off", ptx.U64)
	srcBase := kernel.NewReg("src_base", ptx.U64)
	srcAddr := kernel.NewReg("src_addr", ptx.U64)
	dstIdx := kernel.NewReg("dst_idx", ptx.U32)
	dstIdx64 := kernel.NewReg("dst_idx64", ptx.U64)
	dstOff := kernel.NewReg("dst_off", ptx.U64)
	dstBase := kernel.NewReg("dst_base", ptx.U64)
	dstAddr := kernel.NewReg("dst_addr", ptx.U64)
	tmp    := kernel.NewReg("tmp", ptx.F32)
	w      := kernel.NewReg("w", ptx.U32)
	h      := kernel.NewReg("h", ptx.U32)
	p      := kernel.NewReg("p", ptx.Pred)

	entry   := kernel.NewBlock("entry")
	store   := kernel.NewBlock("store_shared")
	sync1   := kernel.NewBlock("sync1")
	loadSm  := kernel.NewBlock("load_shared")
	exitBlk := kernel.NewBlock("exit")

	// --- entry: read thread/block IDs, compute global coords ---
	entry.Add(builder.Mov(tidX, builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	entry.Add(builder.Mov(tidY, builder.SReg(ptx.RegTidY)).Typed(ptx.U32))
	entry.Add(builder.Mov(ctaX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	entry.Add(builder.Mov(ctaY, builder.SReg(ptx.RegCTAIdY)).Typed(ptx.U32))
	entry.Add(builder.LdParam(w, kernel.Param("width")).Typed(ptx.U32))
	entry.Add(builder.LdParam(h, kernel.Param("height")).Typed(ptx.U32))

	// gx = cta_x * 32 + tid_x
	entry.Add(builder.Mad(gx, ctaX, builder.Imm(32), tidX).Typed(ptx.U32).WithMod(ptx.ModLo))
	// gy = cta_y * 32 + tid_y
	entry.Add(builder.Mad(gy, ctaY, builder.Imm(32), tidY).Typed(ptx.U32).WithMod(ptx.ModLo))

	// bounds check: skip if gx >= width || gy >= height
	entry.Add(builder.Setp(ptx.CmpGe, p, gx, w).Typed(ptx.U32))
	entry.Add(builder.Bra("exit").Pred(p))
	entry.Add(builder.Setp(ptx.CmpGe, p, gy, h).Typed(ptx.U32))
	entry.Add(builder.Bra("exit").Pred(p))

	// --- store_shared: src[gy*width+gx] -> tile[tid_y][tid_x] ---
	// src linear index = gy * width + gx
	store.Add(builder.Mad(srcIdx, gy, w, gx).Typed(ptx.U32).WithMod(ptx.ModLo))
	store.Add(builder.Cvt(srcIdx64, srcIdx).Typed(ptx.U64).From(ptx.U32))
	store.Add(builder.Shl(srcOff, srcIdx64, builder.Imm(2)).Typed(ptx.B64))
	store.Add(builder.LdParam(srcBase, kernel.Param("src")).Typed(ptx.U64))
	store.Add(builder.Add(srcAddr, srcBase, srcOff).Typed(ptx.U64))
	store.Add(builder.Ld(tmp, builder.Addr(srcAddr, 0)).Typed(ptx.F32).InSpace(ptx.Global))

	// shared index: tid_y * 33 + tid_x  (33-stride avoids bank conflicts)
	store.Add(builder.Mad(smIdx, tidY, builder.Imm(33), tidX).Typed(ptx.U32).WithMod(ptx.ModLo))
	store.Add(builder.Cvt(smIdx64, smIdx).Typed(ptx.U64).From(ptx.U32))
	store.Add(builder.Shl(smIdx64, smIdx64, builder.Imm(2)).Typed(ptx.B64))
	store.Add(builder.Cvta(smBase, builder.Sym("tile")).Typed(ptx.U64).InSpace(ptx.Shared))
	store.Add(builder.Add(smAddr, smBase, smIdx64).Typed(ptx.U64))
	store.Add(builder.St(builder.Addr(smAddr, 0), tmp).Typed(ptx.F32).InSpace(ptx.Shared))

	// --- sync1: barrier so all threads have written shared mem ---
	sync1.Add(builder.BarSync(builder.Imm(0)))

	// --- load_shared: tile[tid_x][tid_y] -> dst[transposed index] ---
	// transposed shared index: tid_x * 33 + tid_y
	loadSm.Add(builder.Mad(smIdx, tidX, builder.Imm(33), tidY).Typed(ptx.U32).WithMod(ptx.ModLo))
	loadSm.Add(builder.Cvt(smIdx64, smIdx).Typed(ptx.U64).From(ptx.U32))
	loadSm.Add(builder.Shl(smIdx64, smIdx64, builder.Imm(2)).Typed(ptx.B64))
	loadSm.Add(builder.Cvta(smBase, builder.Sym("tile")).Typed(ptx.U64).InSpace(ptx.Shared))
	loadSm.Add(builder.Add(smAddr, smBase, smIdx64).Typed(ptx.U64))
	loadSm.Add(builder.Ld(tmp, builder.Addr(smAddr, 0)).Typed(ptx.F32).InSpace(ptx.Shared))

	// dst transposed index: gx * height + gy  (note: swapped roles)
	loadSm.Add(builder.Mad(dstIdx, gx, h, gy).Typed(ptx.U32).WithMod(ptx.ModLo))
	loadSm.Add(builder.Cvt(dstIdx64, dstIdx).Typed(ptx.U64).From(ptx.U32))
	loadSm.Add(builder.Shl(dstOff, dstIdx64, builder.Imm(2)).Typed(ptx.B64))
	loadSm.Add(builder.LdParam(dstBase, kernel.Param("dst")).Typed(ptx.U64))
	loadSm.Add(builder.Add(dstAddr, dstBase, dstOff).Typed(ptx.U64))
	loadSm.Add(builder.St(builder.Addr(dstAddr, 0), tmp).Typed(ptx.F32).InSpace(ptx.Global))

	exitBlk.Add(builder.Exit())

	fmt.Println(ptxgen.Build(mod))
}