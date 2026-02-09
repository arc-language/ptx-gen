# ptx-gen

A lightweight, programmatic Go library for generating NVIDIA PTX (Parallel Thread Execution) assembly code.

`ptx-gen` allows you to build CUDA kernels and device functions dynamically at runtime using a type-safe Go API, rather than concatenating strings. It supports modern PTX features including `sm_90` (Hopper), tensor cores, asynchronous copy, and cluster management.

## Installation

```bash
go get github.com/arc-language/ptx-gen
```

## Features

- **Type-Safe Builder**: Go structs and methods for Instructions, Operands, and Registers.
- **Modern PTX Support**: Includes `sm_90`, `sm_100` (Blackwell), async copy (`cp.async`), and mbarrier.
- **Clean Output**: Generates formatted, indented, and readable PTX assembly.
- **Dependency Free**: Pure Go with no external dependencies.

## Usage Example

Here is a complete example of generating a vector addition kernel (`C = A + B`).

```go
package main

import (
	"fmt"
	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
	"github.com/arc-language/ptx-gen/ptxgen"
)

func main() {
	// 1. Initialize Module (ISA 8.5, Target sm_90)
	mod := ptxgen.NewModule(ptx.ISA85, ptx.SM90)

	// 2. Create Kernel: vec_add(float* a, float* b, float* c, int n)
	kernel := mod.NewKernel("vec_add")

	// 3. Define Parameters
	kernel.AddParam(builder.NewPtrParam("a", ptx.Global).WithAlign(8))
	kernel.AddParam(builder.NewPtrParam("b", ptx.Global).WithAlign(8))
	kernel.AddParam(builder.NewPtrParam("c", ptx.Global).WithAlign(8))
	kernel.AddParam(builder.NewParam("n", ptx.U32))

	// 4. Declare Registers
	tidX := kernel.NewReg("tid_x", ptx.U32)
	ntidX := kernel.NewReg("ntid_x", ptx.U32)
	ctidX := kernel.NewReg("ctid_x", ptx.U32)
	idx := kernel.NewReg("idx", ptx.U32)
	idx64 := kernel.NewReg("idx64", ptx.U64)
	offset := kernel.NewReg("offset", ptx.U64)
	nVal := kernel.NewReg("n_val", ptx.U32)
	pred := kernel.NewReg("p", ptx.Pred)
	
	addrA := kernel.NewReg("addr_a", ptx.U64)
	addrB := kernel.NewReg("addr_b", ptx.U64)
	addrC := kernel.NewReg("addr_c", ptx.U64)
	valA := kernel.NewReg("val_a", ptx.F32)
	valB := kernel.NewReg("val_b", ptx.F32)
	valC := kernel.NewReg("val_c", ptx.F32)

	// 5. Build Control Flow
	bbEntry := kernel.NewBlock("entry")
	bbProcess := kernel.NewBlock("process")
	bbExit := kernel.NewBlock("exit")

	// Calculate Global Index: idx = tid.x + (ctaid.x * ntid.x)
	bbEntry.Add(builder.Mov(tidX, builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	bbEntry.Add(builder.Mov(ntidX, builder.SReg(ptx.RegNTidX)).Typed(ptx.U32))
	bbEntry.Add(builder.Mov(ctidX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	bbEntry.Add(builder.Mad(idx, ctidX, ntidX, tidX).Typed(ptx.U32).WithMod(ptx.ModLo))

	// Bounds Check
	bbEntry.Add(builder.LdParam(nVal, kernel.Param("n")).Typed(ptx.U32))
	bbEntry.Add(builder.Setp(ptx.CmpGe, pred, idx, nVal).Typed(ptx.U32))
	bbEntry.Add(builder.Bra("exit").Pred(pred))

	// Load Data
	bbProcess.Add(builder.Cvt(idx64, idx).Typed(ptx.U64).From(ptx.U32))
	// Offset = idx * 4 bytes
	bbProcess.Add(builder.Shl(offset, idx64, builder.Imm(2)).Typed(ptx.B64))
    
    // Load pointers
	bbProcess.Add(builder.LdParam(addrA, kernel.Param("a")).Typed(ptx.U64))
	bbProcess.Add(builder.LdParam(addrB, kernel.Param("b")).Typed(ptx.U64))
	bbProcess.Add(builder.LdParam(addrC, kernel.Param("c")).Typed(ptx.U64))
    
    // Address math
	bbProcess.Add(builder.Add(addrA, addrA, offset).Typed(ptx.U64))
	bbProcess.Add(builder.Add(addrB, addrB, offset).Typed(ptx.U64))
	bbProcess.Add(builder.Add(addrC, addrC, offset).Typed(ptx.U64))

    // Compute
	bbProcess.Add(builder.Ld(valA, builder.Addr(addrA, 0)).Typed(ptx.F32).InSpace(ptx.Global))
	bbProcess.Add(builder.Ld(valB, builder.Addr(addrB, 0)).Typed(ptx.F32).InSpace(ptx.Global))
	bbProcess.Add(builder.Add(valC, valA, valB).Typed(ptx.F32))
	bbProcess.Add(builder.St(builder.Addr(addrC, 0), valC).Typed(ptx.F32).InSpace(ptx.Global))
	bbProcess.Add(builder.Bra("exit"))

	// Exit
	bbExit.Add(builder.Exit())

	// 6. Generate PTX
	fmt.Println(ptxgen.Build(mod))
}
```

## Verifying Generated PTX

You can verify the syntax and register usage of the generated code using NVIDIA's `ptxas` assembler (included in the CUDA Toolkit).

1. **Save the output to a file:**
   ```bash
   go run main.go > kernel.ptx
   ```

2. **Compile with `ptxas`:**
   Replace `sm_90` with your target architecture (e.g., `sm_75`, `sm_80`, `sm_86`).
   
   ```bash
   ptxas -arch=sm_90 -v kernel.ptx -o kernel.cubin
   ```

   **Expected Output:**
   ```text
   ptxas info    : 0 bytes gmem
   ptxas info    : Compiling entry function 'vec_add' for 'sm_90'
   ptxas info    : Function properties for vec_add
   ptxas info    : Used 12 registers, 380 bytes cmem[0]
   ```

   If there are syntax errors, `ptxas` will report the exact line number and error message.

## License

MIT License.