package main

import (
	"fmt"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

// Performs a single 16x16x16 FP16 WMMA matrix multiply-accumulate per warp:
//   D (f32) = A (f16) * B (f16) + C (f32)
// Demonstrates wmma.load, wmma.mma, and wmma.store, plus .reqntid to lock
// block shape to exactly one warp (32 threads).
func main() {
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM80)

	kernel := mod.NewKernel("wmma_matmul")
	// Exactly one 32-thread warp per block — required for WMMA correctness
	kernel.AddDirective(builder.ReqNTid(32, 1, 1))
	kernel.AddParam(builder.NewPtrParam("d", ptx.Global).WithAlign(16)) // f32 output
	kernel.AddParam(builder.NewPtrParam("a", ptx.Global).WithAlign(16)) // f16 matrix A
	kernel.AddParam(builder.NewPtrParam("b", ptx.Global).WithAlign(16)) // f16 matrix B
	kernel.AddParam(builder.NewPtrParam("c", ptx.Global).WithAlign(16)) // f32 accumulator C
	kernel.AddParam(builder.NewParam("lda", ptx.U32))
	kernel.AddParam(builder.NewParam("ldb", ptx.U32))
	kernel.AddParam(builder.NewParam("ldc", ptx.U32))
	kernel.AddParam(builder.NewParam("ldd", ptx.U32))

	// WMMA fragment registers.
	// A: m16n16k16, row-major → 8 x f16x2 = 16 x f16 packed as 8 x b32
	aFrag := make([]*builder.Register, 8)
	for i := range aFrag {
		aFrag[i] = kernel.TempReg(ptx.B32)
	}
	// B: m16n16k16, col-major → 8 x f16x2
	bFrag := make([]*builder.Register, 8)
	for i := range bFrag {
		bFrag[i] = kernel.TempReg(ptx.B32)
	}
	// C/D: m16n16k16 accumulator → 8 x f32
	cFrag := make([]*builder.Register, 8)
	dFrag := make([]*builder.Register, 8)
	for i := range cFrag {
		cFrag[i] = kernel.TempReg(ptx.F32)
		dFrag[i] = kernel.TempReg(ptx.F32)
	}

	// Address / index registers
	ctaX    := kernel.NewReg("cta_x", ptx.U32)
	ctaY    := kernel.NewReg("cta_y", ptx.U32)
	rowA    := kernel.NewReg("row_a", ptx.U32)  // starting row in A for this CTA
	colB    := kernel.NewReg("col_b", ptx.U32)  // starting col in B for this CTA
	baseA   := kernel.NewReg("base_a", ptx.U64)
	baseB   := kernel.NewReg("base_b", ptx.U64)
	baseC   := kernel.NewReg("base_c", ptx.U64)
	baseD   := kernel.NewReg("base_d", ptx.U64)
	ldaVal  := kernel.NewReg("lda_val", ptx.U32)
	ldbVal  := kernel.NewReg("ldb_val", ptx.U32)
	ldcVal  := kernel.NewReg("ldc_val", ptx.U32)
	lddVal  := kernel.NewReg("ldd_val", ptx.U32)
	tmp64   := kernel.NewReg("tmp64", ptx.U64)
	rowOff  := kernel.NewReg("row_off", ptx.U64)
	colOff  := kernel.NewReg("col_off", ptx.U64)
	addrA   := kernel.NewReg("addr_a", ptx.U64)
	addrB   := kernel.NewReg("addr_b", ptx.U64)
	addrC   := kernel.NewReg("addr_c", ptx.U64)
	addrD   := kernel.NewReg("addr_d", ptx.U64)

	entry    := kernel.NewBlock("entry")
	loadFrag := kernel.NewBlock("load_fragments")
	mmaBlk   := kernel.NewBlock("mma")
	storeBlk := kernel.NewBlock("store")
	exitBlk  := kernel.NewBlock("exit")

	// --- entry: read params ---
	entry.Add(builder.Mov(ctaX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	entry.Add(builder.Mov(ctaY, builder.SReg(ptx.RegCTAIdY)).Typed(ptx.U32))
	entry.Add(builder.LdParam(baseA, kernel.Param("a")).Typed(ptx.U64))
	entry.Add(builder.LdParam(baseB, kernel.Param("b")).Typed(ptx.U64))
	entry.Add(builder.LdParam(baseC, kernel.Param("c")).Typed(ptx.U64))
	entry.Add(builder.LdParam(baseD, kernel.Param("d")).Typed(ptx.U64))
	entry.Add(builder.LdParam(ldaVal, kernel.Param("lda")).Typed(ptx.U32))
	entry.Add(builder.LdParam(ldbVal, kernel.Param("ldb")).Typed(ptx.U32))
	entry.Add(builder.LdParam(ldcVal, kernel.Param("ldc")).Typed(ptx.U32))
	entry.Add(builder.LdParam(lddVal, kernel.Param("ldd")).Typed(ptx.U32))

	// row_a = cta_y * 16,  col_b = cta_x * 16
	entry.Add(builder.Mul(rowA, ctaY, builder.Imm(16)).Typed(ptx.U32).WithMod(ptx.ModLo))
	entry.Add(builder.Mul(colB, ctaX, builder.Imm(16)).Typed(ptx.U32).WithMod(ptx.ModLo))

	// --- load_fragments ---
	// Compute byte offset for A: (row_a * lda) * sizeof(f16) = row_a*lda*2
	loadFrag.Add(builder.Mul(tmp64, rowA, ldaVal).Typed(ptx.U32).WithMod(ptx.ModLo)) // u32 first
	// widen & scale
	loadFrag.Add(builder.Cvt(rowOff, tmp64).Typed(ptx.U64).From(ptx.U32))
	loadFrag.Add(builder.Shl(rowOff, rowOff, builder.Imm(1)).Typed(ptx.B64)) // *2 for f16
	loadFrag.Add(builder.Add(addrA, baseA, rowOff).Typed(ptx.U64))

	// Compute byte offset for B: col_b * sizeof(f16) = col_b * 2
	loadFrag.Add(builder.Cvt(colOff, colB).Typed(ptx.U64).From(ptx.U32))
	loadFrag.Add(builder.Shl(colOff, colOff, builder.Imm(1)).Typed(ptx.B64))
	loadFrag.Add(builder.Add(addrB, baseB, colOff).Typed(ptx.U64))

	// Compute byte offset for C: (row_a * ldc) * sizeof(f32) = row_a*ldc*4
	loadFrag.Add(builder.Mul(tmp64, rowA, ldcVal).Typed(ptx.U32).WithMod(ptx.ModLo))
	loadFrag.Add(builder.Cvt(rowOff, tmp64).Typed(ptx.U64).From(ptx.U32))
	loadFrag.Add(builder.Shl(rowOff, rowOff, builder.Imm(2)).Typed(ptx.B64)) // *4 for f32
	loadFrag.Add(builder.Add(addrC, baseC, rowOff).Typed(ptx.U64))

	// wmma.load.a.sync.aligned.m16n16k16.row.f16 {aFrag...}, [addrA], lda
	aVec := builder.Vec(operandsFrom(aFrag)...)
	loadFrag.Add(builder.WmmaLoad(
		ptx.ModMatrixA, ptx.ModRow, ptx.ModShapeM16N16K16,
		ptx.F16, aVec, addrA, ldaVal,
	))

	// wmma.load.b.sync.aligned.m16n16k16.col.f16 {bFrag...}, [addrB], ldb
	bVec := builder.Vec(operandsFrom(bFrag)...)
	loadFrag.Add(builder.WmmaLoad(
		ptx.ModMatrixB, ptx.ModCol, ptx.ModShapeM16N16K16,
		ptx.F16, bVec, addrB, ldbVal,
	))

	// wmma.load.c.sync.aligned.m16n16k16.row.f32 {cFrag...}, [addrC], ldc
	cVec := builder.Vec(operandsFrom(cFrag)...)
	loadFrag.Add(builder.WmmaLoad(
		ptx.ModMatrixC, ptx.ModRow, ptx.ModShapeM16N16K16,
		ptx.F32, cVec, addrC, ldcVal,
	))

	// --- mma: wmma.mma.sync.aligned.m16n16k16.row.col.f32.f32 ---
	dVec := builder.Vec(operandsFrom(dFrag)...)
	mmaBlk.Add(builder.WmmaMma(
		ptx.ModShapeM16N16K16, ptx.ModRow, ptx.ModCol,
		dVec, aVec, bVec, cVec,
	))

	// --- store: wmma.store.d.sync.aligned.m16n16k16.row.f32 ---
	// byte offset for D: (row_a * ldd) * sizeof(f32)
	storeBlk.Add(builder.Mul(tmp64, rowA, lddVal).Typed(ptx.U32).WithMod(ptx.ModLo))
	storeBlk.Add(builder.Cvt(rowOff, tmp64).Typed(ptx.U64).From(ptx.U32))
	storeBlk.Add(builder.Shl(rowOff, rowOff, builder.Imm(2)).Typed(ptx.B64))
	storeBlk.Add(builder.Add(addrD, baseD, rowOff).Typed(ptx.U64))

	storeBlk.Add(builder.WmmaStore(
		ptx.ModMatrixD, ptx.ModRow, ptx.ModShapeM16N16K16,
		ptx.F32, addrD, dVec, lddVal,
	))
	storeBlk.Add(builder.Bra("exit"))

	exitBlk.Add(builder.Exit())

	fmt.Println(ptxgen.Build(mod))
}

// operandsFrom converts a []*builder.Register slice to []builder.Operand.
func operandsFrom(regs []*builder.Register) []builder.Operand {
	ops := make([]builder.Operand, len(regs))
	for i, r := range regs {
		ops[i] = r
	}
	return ops
}