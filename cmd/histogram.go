package main

import (
	"fmt"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

// Computes a 256-bin histogram of byte values using global atomic increments.
// Also shows .const memory for a lookup table and .global array initialisation.
func main() {
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM80)

	// Pre-initialised constant: bin scaling factor per bin (trivially 1 here,
	// but demonstrates .const + initialiser syntax)
	constScale := builder.NewGlobalArray("bin_scale", ptx.Const, ptx.U32, 4).
		WithInit(int64(1), int64(1), int64(1), int64(1))
	mod.AddGlobal(constScale)

	kernel := mod.NewKernel("histogram")
	kernel.AddDirective(builder.MaxNTid(256, 1, 1))
	kernel.AddParam(builder.NewPtrParam("hist", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewPtrParam("data", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewParam("n", ptx.U32))

	tidX     := kernel.NewReg("tid_x", ptx.U32)
	ctaX     := kernel.NewReg("cta_x", ptx.U32)
	ntidX    := kernel.NewReg("ntid_x", ptx.U32)
	gidx     := kernel.NewReg("gidx", ptx.U32)
	gidx64   := kernel.NewReg("gidx64", ptx.U64)
	stride   := kernel.NewReg("stride", ptx.U32)
	dataBase := kernel.NewReg("data_base", ptx.U64)
	histBase := kernel.NewReg("hist_base", ptx.U64)
	dataAddr := kernel.NewReg("data_addr", ptx.U64)
	histAddr := kernel.NewReg("hist_addr", ptx.U64)
	byteVal  := kernel.NewReg("byte_val", ptx.U8)
	binIdx   := kernel.NewReg("bin_idx", ptx.U32)
	binIdx64 := kernel.NewReg("bin_idx64", ptx.U64)
	binOff   := kernel.NewReg("bin_off", ptx.U64)
	nVal     := kernel.NewReg("n_val", ptx.U32)
	old      := kernel.NewReg("old_val", ptx.U32)
	p        := kernel.NewReg("p", ptx.Pred)
	nCTA     := kernel.NewReg("n_cta", ptx.U32)

	entry   := kernel.NewBlock("entry")
	loop    := kernel.NewBlock("loop")
	body    := kernel.NewBlock("body")
	next    := kernel.NewBlock("next")
	exitBlk := kernel.NewBlock("exit")

	// --- entry: compute gidx and grid stride ---
	entry.Add(builder.Mov(tidX, builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	entry.Add(builder.Mov(ctaX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	entry.Add(builder.Mov(ntidX, builder.SReg(ptx.RegNTidX)).Typed(ptx.U32))
	entry.Add(builder.Mov(nCTA, builder.SReg(ptx.RegNCTAIdX)).Typed(ptx.U32))
	entry.Add(builder.LdParam(nVal, kernel.Param("n")).Typed(ptx.U32))
	entry.Add(builder.LdParam(dataBase, kernel.Param("data")).Typed(ptx.U64))
	entry.Add(builder.LdParam(histBase, kernel.Param("hist")).Typed(ptx.U64))

	// gidx = cta_x * ntid_x + tid_x
	entry.Add(builder.Mad(gidx, ctaX, ntidX, tidX).Typed(ptx.U32).WithMod(ptx.ModLo))
	// stride = n_cta * ntid_x  (grid-stride loop)
	entry.Add(builder.Mul(stride, nCTA, ntidX).Typed(ptx.U32).WithMod(ptx.ModLo))

	// --- loop: bounds check ---
	loop.Add(builder.Setp(ptx.CmpGe, p, gidx, nVal).Typed(ptx.U32))
	loop.Add(builder.Bra("exit").Pred(p))

	// --- body: load byte, compute bin, atomic increment ---
	body.Add(builder.Cvt(gidx64, gidx).Typed(ptx.U64).From(ptx.U32))
	body.Add(builder.Add(dataAddr, dataBase, gidx64).Typed(ptx.U64))
	body.Add(builder.Ld(byteVal, builder.Addr(dataAddr, 0)).Typed(ptx.U8).InSpace(ptx.Global))

	// zero-extend byte to u32 for bin addressing
	body.Add(builder.Cvt(binIdx, byteVal).Typed(ptx.U32).From(ptx.U8))
	body.Add(builder.Cvt(binIdx64, binIdx).Typed(ptx.U64).From(ptx.U32))
	// bin byte offset = bin_idx * 4 (u32 elements)
	body.Add(builder.Shl(binOff, binIdx64, builder.Imm(2)).Typed(ptx.B64))
	body.Add(builder.Add(histAddr, histBase, binOff).Typed(ptx.U64))

	// atom.global.add.u32 old_val, [hist_addr], 1
	body.Add(builder.Atom(ptx.ModAtomAdd, old,
		builder.Addr(histAddr, 0), builder.Imm(1)).
		Typed(ptx.U32).InSpace(ptx.Global))

	// --- next: advance by stride and loop ---
	next.Add(builder.Add(gidx, gidx, stride).Typed(ptx.U32))
	next.Add(builder.Bra("loop"))

	exitBlk.Add(builder.Exit())

	fmt.Println(ptxgen.Build(mod))
}