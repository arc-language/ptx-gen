package main

import (
	"fmt"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

func main() {
	// 1. Create a new PTX Module (ISA 8.5, Target Architecture sm_90 - Hopper)
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM90)

	// -------------------------------------------------------------------------
	// Define a Kernel: vec_add
	// signature: void vec_add(float* a, float* b, float* c, int n)
	// -------------------------------------------------------------------------
	kernel := mod.NewKernel("vec_add")

	// 2. Define Parameters
	// .param .u64 .ptr .global .align 8 a
	kernel.AddParam(builder.NewPtrParam("a", ptx.Global).WithAlign(8))
	kernel.AddParam(builder.NewPtrParam("b", ptx.Global).WithAlign(8))
	kernel.AddParam(builder.NewPtrParam("c", ptx.Global).WithAlign(8))
	kernel.AddParam(builder.NewParam("n", ptx.U32))

	// 3. Define Registers
	// We declare named registers. The builder will group them into .reg declarations.
	tidX := kernel.NewReg("tid_x", ptx.U32)
	ntidX := kernel.NewReg("ntid_x", ptx.U32)
	ctidX := kernel.NewReg("ctid_x", ptx.U32)
	idx := kernel.NewReg("idx", ptx.U32)
	idx64 := kernel.NewReg("idx64", ptx.U64)
	offset := kernel.NewReg("offset", ptx.U64)
	nVal := kernel.NewReg("n_val", ptx.U32)
	pred := kernel.NewReg("p", ptx.Pred)

	// Address registers
	addrA := kernel.NewReg("addr_a", ptx.U64)
	addrB := kernel.NewReg("addr_b", ptx.U64)
	addrC := kernel.NewReg("addr_c", ptx.U64)

	// Data registers
	valA := kernel.NewReg("val_a", ptx.F32)
	valB := kernel.NewReg("val_b", ptx.F32)
	valC := kernel.NewReg("val_c", ptx.F32)

	// 4. Construct the Basic Block structure
	bbEntry := kernel.NewBlock("entry")
	bbProcess := kernel.NewBlock("process")
	bbExit := kernel.NewBlock("exit")

	// -------------------------------------------------------------------------
	// Block: entry
	// Calculate Global Index: idx = tid.x + (ctaid.x * ntid.x)
	// -------------------------------------------------------------------------
	
	// mov.u32 %tid_x, %tid.x;
	bbEntry.Add(builder.Mov(tidX, builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	// mov.u32 %ntid_x, %ntid.x;
	bbEntry.Add(builder.Mov(ntidX, builder.SReg(ptx.RegNTidX)).Typed(ptx.U32))
	// mov.u32 %ctid_x, %ctaid.x;
	bbEntry.Add(builder.Mov(ctidX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))

	// mad.lo.u32 %idx, %ctid_x, %ntid_x, %tid_x;
	bbEntry.Add(builder.Mad(idx, ctidX, ntidX, tidX).Typed(ptx.U32).WithMod(ptx.ModLo))

	// ld.param.u32 %n_val, [n];
	bbEntry.Add(builder.LdParam(nVal, kernel.Param("n")).Typed(ptx.U32))

	// setp.ge.u32 %p, %idx, %n_val;
	bbEntry.Add(builder.Setp(ptx.CmpGe, pred, idx, nVal).Typed(ptx.U32))

	// @%p bra exit;
	bbEntry.Add(builder.Bra("exit").Pred(pred))

	// -------------------------------------------------------------------------
	// Block: process
	// Compute addresses, Load, Add, Store
	// -------------------------------------------------------------------------
	
	// Convert index to 64-bit for pointer arithmetic
	// cvt.u64.u32 %idx64, %idx;
	bbProcess.Add(builder.Cvt(idx64, idx).Typed(ptx.U64).From(ptx.U32))

	// Multiply index by 4 (float size) to get byte offset
	// shl.b64 %offset, %idx64, 2;
	bbProcess.Add(builder.Shl(offset, idx64, builder.Imm(2)).Typed(ptx.B64))

	// Load pointers from params
	bbProcess.Add(builder.LdParam(addrA, kernel.Param("a")).Typed(ptx.U64))
	bbProcess.Add(builder.LdParam(addrB, kernel.Param("b")).Typed(ptx.U64))
	bbProcess.Add(builder.LdParam(addrC, kernel.Param("c")).Typed(ptx.U64))

	// Add offsets to base pointers
	// add.u64 %addr_a, %addr_a, %offset;
	bbProcess.Add(builder.Add(addrA, addrA, offset).Typed(ptx.U64))
	bbProcess.Add(builder.Add(addrB, addrB, offset).Typed(ptx.U64))
	bbProcess.Add(builder.Add(addrC, addrC, offset).Typed(ptx.U64))

	// Load data from global memory
	// ld.global.f32 %val_a, [%addr_a];
	bbProcess.Add(builder.Ld(valA, builder.Addr(addrA, 0)).Typed(ptx.F32).InSpace(ptx.Global))
	bbProcess.Add(builder.Ld(valB, builder.Addr(addrB, 0)).Typed(ptx.F32).InSpace(ptx.Global))

	// Perform addition
	// add.f32 %val_c, %val_a, %val_b;
	bbProcess.Add(builder.Add(valC, valA, valB).Typed(ptx.F32))

	// Store result
	// st.global.f32 [%addr_c], %val_c;
	bbProcess.Add(builder.St(builder.Addr(addrC, 0), valC).Typed(ptx.F32).InSpace(ptx.Global))

	// Bra to exit (optional here as it falls through, but good practice)
	bbProcess.Add(builder.Bra("exit"))

	// -------------------------------------------------------------------------
	// Block: exit
	// -------------------------------------------------------------------------
	bbExit.Add(builder.Exit())

	// 5. Generate and Print the PTX String
	output := ptxgen.Build(mod)
	fmt.Println(output)
}