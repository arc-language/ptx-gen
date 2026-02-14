package main

import (
	"fmt"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

// Block-level sum reduction using warp shuffle intrinsics (__shfl_down_sync).
// Each warp reduces 32 elements internally; warp results are collected via shared mem.
func main() {
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM80)

	// Shared mem: one slot per warp (blockDim.x=256 → 8 warps)
	warpSums := builder.NewGlobalArray("warp_sums", ptx.Shared, ptx.F32, 8).WithAlign(4)
	mod.AddGlobal(warpSums)

	kernel := mod.NewKernel("reduce_sum")
	kernel.AddDirective(builder.MaxNTid(256, 1, 1))
	kernel.AddParam(builder.NewPtrParam("input", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewPtrParam("output", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewParam("n", ptx.U32))

	// Registers
	tidX    := kernel.NewReg("tid_x", ptx.U32)
	ctaX    := kernel.NewReg("cta_x", ptx.U32)
	gidx    := kernel.NewReg("gidx", ptx.U32)
	gidx64  := kernel.NewReg("gidx64", ptx.U64)
	off     := kernel.NewReg("off", ptx.U64)
	base    := kernel.NewReg("base", ptx.U64)
	addr    := kernel.NewReg("addr", ptx.U64)
	val     := kernel.NewReg("val", ptx.F32)
	shufVal := kernel.NewReg("shuf_val", ptx.F32)
	laneId  := kernel.NewReg("lane_id", ptx.U32)
	warpId  := kernel.NewReg("warp_id", ptx.U32)
	nVal    := kernel.NewReg("n_val", ptx.U32)
	p       := kernel.NewReg("p", ptx.Pred)
	smIdx64 := kernel.NewReg("sm_idx64", ptx.U64)
	smBase  := kernel.NewReg("sm_base", ptx.U64)
	smAddr  := kernel.NewReg("sm_addr", ptx.U64)
	fullMask := kernel.NewReg("full_mask", ptx.U32)

	entry      := kernel.NewBlock("entry")
	warpReduce := kernel.NewBlock("warp_reduce")
	storeWarp  := kernel.NewBlock("store_warp")
	finalBlk   := kernel.NewBlock("final_reduce")
	exitBlk    := kernel.NewBlock("exit")

	// --- entry ---
	entry.Add(builder.Mov(tidX, builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	entry.Add(builder.Mov(ctaX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	entry.Add(builder.LdParam(nVal, kernel.Param("n")).Typed(ptx.U32))

	// gidx = cta_x * 256 + tid_x
	entry.Add(builder.Mad(gidx, ctaX, builder.Imm(256), tidX).Typed(ptx.U32).WithMod(ptx.ModLo))

	// bounds check
	entry.Add(builder.Setp(ptx.CmpGe, p, gidx, nVal).Typed(ptx.U32))
	entry.Add(builder.Mov(val, builder.ImmF32(0.0)).Typed(ptx.F32))
	entry.Add(builder.Bra("warp_reduce").Pred(p))

	// load input[gidx]
	entry.Add(builder.Cvt(gidx64, gidx).Typed(ptx.U64).From(ptx.U32))
	entry.Add(builder.Shl(off, gidx64, builder.Imm(2)).Typed(ptx.B64))
	entry.Add(builder.LdParam(base, kernel.Param("input")).Typed(ptx.U64))
	entry.Add(builder.Add(addr, base, off).Typed(ptx.U64))
	entry.Add(builder.Ld(val, builder.Addr(addr, 0)).Typed(ptx.F32).InSpace(ptx.Global))

	// --- warp_reduce: butterfly reduction using shfl.down.sync ---
	// mask = 0xFFFFFFFF (all 32 lanes active)
	warpReduce.Add(builder.Mov(fullMask, builder.ImmU(0xFFFFFFFF)).Typed(ptx.U32))

	// shfl.sync.down.b32 shuf_val, val, 16, 31, mask  → reduce offset 16
	warpReduce.Add(builder.ShflSync(shufVal, val, builder.Imm(16), builder.Imm(31), fullMask).
		Typed(ptx.B32).WithMod(ptx.ModShflDown))
	warpReduce.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	// offset 8
	warpReduce.Add(builder.ShflSync(shufVal, val, builder.Imm(8), builder.Imm(31), fullMask).
		Typed(ptx.B32).WithMod(ptx.ModShflDown))
	warpReduce.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	// offset 4
	warpReduce.Add(builder.ShflSync(shufVal, val, builder.Imm(4), builder.Imm(31), fullMask).
		Typed(ptx.B32).WithMod(ptx.ModShflDown))
	warpReduce.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	// offset 2
	warpReduce.Add(builder.ShflSync(shufVal, val, builder.Imm(2), builder.Imm(31), fullMask).
		Typed(ptx.B32).WithMod(ptx.ModShflDown))
	warpReduce.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	// offset 1
	warpReduce.Add(builder.ShflSync(shufVal, val, builder.Imm(1), builder.Imm(31), fullMask).
		Typed(ptx.B32).WithMod(ptx.ModShflDown))
	warpReduce.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))

	// --- store_warp: lane 0 of each warp writes partial sum to shared mem ---
	storeWarp.Add(builder.Mov(laneId, builder.SReg(ptx.RegLaneId)).Typed(ptx.U32))
	storeWarp.Add(builder.Mov(warpId, builder.SReg(ptx.RegWarpId)).Typed(ptx.U32))
	storeWarp.Add(builder.Setp(ptx.CmpNe, p, laneId, builder.Imm(0)).Typed(ptx.U32))
	storeWarp.Add(builder.Bra("final_reduce").Pred(p))

	// warp_sums[warp_id] = val
	storeWarp.Add(builder.Cvt(smIdx64, warpId).Typed(ptx.U64).From(ptx.U32))
	storeWarp.Add(builder.Shl(smIdx64, smIdx64, builder.Imm(2)).Typed(ptx.B64))
	storeWarp.Add(builder.Cvta(smBase, builder.Sym("warp_sums")).Typed(ptx.U64).InSpace(ptx.Shared))
	storeWarp.Add(builder.Add(smAddr, smBase, smIdx64).Typed(ptx.U64))
	storeWarp.Add(builder.St(builder.Addr(smAddr, 0), val).Typed(ptx.F32).InSpace(ptx.Shared))

	// --- final_reduce: thread 0 sums all warp results and writes to output ---
	finalBlk.Add(builder.BarSync(builder.Imm(0)))
	finalBlk.Add(builder.Setp(ptx.CmpNe, p, tidX, builder.Imm(0)).Typed(ptx.U32))
	finalBlk.Add(builder.Bra("exit").Pred(p))

	finalBlk.Add(builder.Cvta(smBase, builder.Sym("warp_sums")).Typed(ptx.U64).InSpace(ptx.Shared))
	// manually unroll 8-warp accumulation into val
	finalBlk.Add(builder.Ld(val, builder.Addr(smBase, 0)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 4)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 8)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 12)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 16)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 20)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 24)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))
	finalBlk.Add(builder.Ld(shufVal, builder.Addr(smBase, 28)).Typed(ptx.F32).InSpace(ptx.Shared))
	finalBlk.Add(builder.Add(val, val, shufVal).Typed(ptx.F32))

	// output[cta_x] = val
	finalBlk.Add(builder.Cvt(gidx64, ctaX).Typed(ptx.U64).From(ptx.U32))
	finalBlk.Add(builder.Shl(off, gidx64, builder.Imm(2)).Typed(ptx.B64))
	finalBlk.Add(builder.LdParam(base, kernel.Param("output")).Typed(ptx.U64))
	finalBlk.Add(builder.Add(addr, base, off).Typed(ptx.U64))
	finalBlk.Add(builder.St(builder.Addr(addr, 0), val).Typed(ptx.F32).InSpace(ptx.Global))

	exitBlk.Add(builder.Exit())

	fmt.Println(ptxgen.Build(mod))
}