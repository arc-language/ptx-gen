package main

import (
	"fmt"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

// Applies ReLU element-wise: out[i] = max(0, in[i])
// Uses selp (select with predicate) instead of a branch for branchless execution.
// Also shows a device helper function called from the kernel.
func main() {
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM80)

	// -------------------------------------------------------------------------
	// Device function: relu_f32(float x) -> float
	// -------------------------------------------------------------------------
	fn := mod.NewFunc("relu_f32")
	fn.AddParam(builder.NewParam("x", ptx.F32))
	fn.AddReturnParam(builder.NewParam("result", ptx.F32))

	xIn    := fn.NewReg("x_in", ptx.F32)
	zero   := fn.NewReg("zero", ptx.F32)
	result := fn.NewReg("result", ptx.F32)
	pPos   := fn.NewReg("p_pos", ptx.Pred)

	fnBody := fn.NewBlock("body")
	fnBody.Add(builder.LdParam(xIn, fn.Param("x")).Typed(ptx.F32))
	fnBody.Add(builder.Mov(zero, builder.ImmF32(0.0)).Typed(ptx.F32))
	// setp.gt.f32 p_pos, x_in, 0.0
	fnBody.Add(builder.Setp(ptx.CmpGt, pPos, xIn, zero).Typed(ptx.F32))
	// selp.f32 result, x_in, 0.0, p_pos  → result = p_pos ? x_in : 0.0
	fnBody.Add(builder.Selp(result, xIn, zero, pPos).Typed(ptx.F32))
	fnBody.Add(builder.St(builder.Addr(fn.Param("result"), 0), result).
		Typed(ptx.F32).InSpace(ptx.Param))
	fnBody.Add(builder.Ret())

	// -------------------------------------------------------------------------
	// Kernel: relu_kernel — calls relu_f32 per element
	// -------------------------------------------------------------------------
	kernel := mod.NewKernel("relu_kernel")
	kernel.AddDirective(builder.MaxNTid(256, 1, 1))
	kernel.AddDirective(builder.MinNCTAPerSM(4))
	kernel.AddParam(builder.NewPtrParam("output", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewPtrParam("input", ptx.Global).WithAlign(16))
	kernel.AddParam(builder.NewParam("n", ptx.U32))

	tidX    := kernel.NewReg("tid_x", ptx.U32)
	ctaX    := kernel.NewReg("cta_x", ptx.U32)
	gidx    := kernel.NewReg("gidx", ptx.U32)
	gidx64  := kernel.NewReg("gidx64", ptx.U64)
	off     := kernel.NewReg("off", ptx.U64)
	inBase  := kernel.NewReg("in_base", ptx.U64)
	outBase := kernel.NewReg("out_base", ptx.U64)
	inAddr  := kernel.NewReg("in_addr", ptx.U64)
	outAddr := kernel.NewReg("out_addr", ptx.U64)
	val     := kernel.NewReg("val", ptx.F32)
	rVal    := kernel.NewReg("r_val", ptx.F32)
	nVal    := kernel.NewReg("n_val", ptx.U32)
	p       := kernel.NewReg("p", ptx.Pred)

	entry   := kernel.NewBlock("entry")
	process := kernel.NewBlock("process")
	exitBlk := kernel.NewBlock("exit")

	entry.Add(builder.Mov(tidX, builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	entry.Add(builder.Mov(ctaX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	entry.Add(builder.LdParam(nVal, kernel.Param("n")).Typed(ptx.U32))
	entry.Add(builder.Mad(gidx, ctaX, builder.Imm(256), tidX).Typed(ptx.U32).WithMod(ptx.ModLo))
	entry.Add(builder.Setp(ptx.CmpGe, p, gidx, nVal).Typed(ptx.U32))
	entry.Add(builder.Bra("exit").Pred(p))

	// compute address
	process.Add(builder.Cvt(gidx64, gidx).Typed(ptx.U64).From(ptx.U32))
	process.Add(builder.Shl(off, gidx64, builder.Imm(2)).Typed(ptx.B64))
	process.Add(builder.LdParam(inBase, kernel.Param("input")).Typed(ptx.U64))
	process.Add(builder.LdParam(outBase, kernel.Param("output")).Typed(ptx.U64))
	process.Add(builder.Add(inAddr, inBase, off).Typed(ptx.U64))
	process.Add(builder.Add(outAddr, outBase, off).Typed(ptx.U64))

	// load
	process.Add(builder.Ld(val, builder.Addr(inAddr, 0)).Typed(ptx.F32).InSpace(ptx.Global))

	// call relu_f32(val) → rVal
	process.Add(builder.Call("relu_f32", []builder.Operand{rVal}, []builder.Operand{val}))

	// store
	process.Add(builder.St(builder.Addr(outAddr, 0), rVal).Typed(ptx.F32).InSpace(ptx.Global))
	process.Add(builder.Bra("exit"))

	exitBlk.Add(builder.Exit())

	fmt.Println(ptxgen.Build(mod))
}