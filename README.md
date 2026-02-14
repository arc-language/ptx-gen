# ptx-gen

A lightweight, programmatic Go library for generating NVIDIA PTX (Parallel Thread Execution) assembly code.

**Current Coverage:** This library implements PTX features from sections 1–10 of the [NVIDIA PTX ISA documentation](https://docs.nvidia.com/cuda/parallel-thread-execution), including all fundamental types, state spaces, instructions, memory operations, control flow, synchronization, tensor core operations, and async copy primitives.

`ptx-gen` allows you to build CUDA kernels and device functions dynamically at runtime using a type-safe Go API, rather than concatenating strings. It supports modern PTX features including `sm_90` (Hopper), `sm_100` (Blackwell), tensor cores (WMMA, MMA, WGMMA, tcgen05), asynchronous copy, mbarrier, and cluster management.

## Installation

```bash
go get github.com/arc-language/ptx-gen
```

## Features

- **Type-Safe Builder**: Go structs and methods for Instructions, Operands, Registers, Parameters, and Globals.
- **Full ISA Coverage (Sections 1–10)**: All fundamental types, state spaces, instructions, and directives.
- **Modern Architecture Support**: `sm_90`, `sm_100` (Blackwell), and targets back to `sm_50`.
- **Tensor Core Support**: WMMA, MMA, WGMMA, and tcgen05 (5th-gen Tensor Core) instructions.
- **Async Primitives**: `cp.async`, `cp.async.bulk`, bulk tensor copy, and mbarrier.
- **Cluster Management**: Cluster barriers, `clusterlaunchcontrol`, and cluster-scoped memory.
- **Performance Directives**: `.maxnreg`, `.maxntid`, `.reqntid`, `.minnctapersm`, `.pragma`, `.explicitcluster`, and more.
- **Variable Attributes**: `.managed`, `.unified` for unified virtual memory.
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
	tidX  := kernel.NewReg("tid_x",  ptx.U32)
	ntidX := kernel.NewReg("ntid_x", ptx.U32)
	ctidX := kernel.NewReg("ctid_x", ptx.U32)
	idx   := kernel.NewReg("idx",    ptx.U32)
	idx64 := kernel.NewReg("idx64",  ptx.U64)
	offset := kernel.NewReg("offset", ptx.U64)
	nVal  := kernel.NewReg("n_val",  ptx.U32)
	pred  := kernel.NewReg("p",      ptx.Pred)

	addrA := kernel.NewReg("addr_a", ptx.U64)
	addrB := kernel.NewReg("addr_b", ptx.U64)
	addrC := kernel.NewReg("addr_c", ptx.U64)
	valA  := kernel.NewReg("val_a",  ptx.F32)
	valB  := kernel.NewReg("val_b",  ptx.F32)
	valC  := kernel.NewReg("val_c",  ptx.F32)

	// 5. Build Control Flow
	bbEntry   := kernel.NewBlock("entry")
	bbProcess := kernel.NewBlock("process")
	bbExit    := kernel.NewBlock("exit")

	// Calculate Global Index: idx = ctaid.x * ntid.x + tid.x
	bbEntry.Add(builder.Mov(tidX,  builder.SReg(ptx.RegTidX)).Typed(ptx.U32))
	bbEntry.Add(builder.Mov(ntidX, builder.SReg(ptx.RegNTidX)).Typed(ptx.U32))
	bbEntry.Add(builder.Mov(ctidX, builder.SReg(ptx.RegCTAIdX)).Typed(ptx.U32))
	bbEntry.Add(builder.Mad(idx, ctidX, ntidX, tidX).Typed(ptx.U32).WithMod(ptx.ModLo))

	// Bounds check: if idx >= n, exit
	bbEntry.Add(builder.LdParam(nVal, kernel.Param("n")).Typed(ptx.U32))
	bbEntry.Add(builder.Setp(ptx.CmpGe, pred, idx, nVal).Typed(ptx.U32))
	bbEntry.Add(builder.Bra("exit").Pred(pred))

	// Compute byte offset: offset = idx * 4
	bbProcess.Add(builder.Cvt(idx64, idx).Typed(ptx.U64).From(ptx.U32))
	bbProcess.Add(builder.Shl(offset, idx64, builder.Imm(2)).Typed(ptx.B64))

	// Load base pointers
	bbProcess.Add(builder.LdParam(addrA, kernel.Param("a")).Typed(ptx.U64))
	bbProcess.Add(builder.LdParam(addrB, kernel.Param("b")).Typed(ptx.U64))
	bbProcess.Add(builder.LdParam(addrC, kernel.Param("c")).Typed(ptx.U64))

	// Apply offset
	bbProcess.Add(builder.Add(addrA, addrA, offset).Typed(ptx.U64))
	bbProcess.Add(builder.Add(addrB, addrB, offset).Typed(ptx.U64))
	bbProcess.Add(builder.Add(addrC, addrC, offset).Typed(ptx.U64))

	// Load, add, store
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

You can verify syntax and register usage with NVIDIA's `ptxas` assembler (included in the CUDA Toolkit).

**Save the output:**
```bash
go run main.go > kernel.ptx
```

**Assemble with `ptxas`:**
```bash
ptxas -arch=sm_90 -v kernel.ptx -o kernel.cubin
```

**Expected output:**
```
ptxas info    : 0 bytes gmem
ptxas info    : Compiling entry function 'vec_add' for 'sm_90'
ptxas info    : Function properties for vec_add
ptxas info    : Used 12 registers, 380 bytes cmem[0]
```

If there are syntax errors, `ptxas` will report the exact line and message.

---

## API Reference

### Module

`ptxgen.NewModule(version ptx.ISAVersion, target ptx.Target) *builder.Module`

Creates a new PTX module. The module is the top-level container, mapping to a single `.ptx` file.

```go
mod := ptxgen.NewModule(ptx.ISA85, ptx.SM90)
mod.AddGlobal(builder.NewGlobalArray("smem", ptx.Shared, ptx.F32, 256))
```

**ISA Versions:** `ISA60` through `ISA91`  
**Targets:** `SM50` through `SM120` (including `SM90a`)

---

### Functions & Kernels

```go
kernel := mod.NewKernel("my_kernel")   // .entry (GPU entry point)
fn     := mod.NewFunc("my_device_fn")  // .func  (device function)
```

**Parameters:**

```go
kernel.AddParam(builder.NewParam("n", ptx.U32))                   // simple scalar
kernel.AddParam(builder.NewPtrParam("ptr", ptx.Global).WithAlign(8)) // pointer
kernel.AddParam(builder.NewByteArrayParam("mat", 64, 16))         // struct-by-value

fn.AddReturnParam(builder.NewParam("result", ptx.F32))            // return value
```

**Registers:**

```go
r   := kernel.NewReg("my_reg", ptx.F32)  // named register
tmp := kernel.TempReg(ptx.U32)            // auto-named (%t0, %t1, ...)
```

**Performance Directives:**

```go
kernel.AddDirective(builder.MaxNReg(128))
kernel.AddDirective(builder.MaxNTid(256, 1, 1))
kernel.AddDirective(builder.ReqNTid(128))
kernel.AddDirective(builder.MinNCTAPerSM(2))
kernel.AddDirective(builder.MaxNCTAPerSM(4))
kernel.AddDirective(builder.Pragma("nounroll"))
kernel.AddDirective(builder.ExplicitCluster())
kernel.AddDirective(builder.MaxClusterRank(8))
kernel.AddDirective(builder.BlocksAreClusters())
kernel.AddDirective(builder.NoReturn())
```

---

### Basic Blocks & Control Flow

```go
entry  := kernel.NewBlock("entry")
loop   := kernel.NewBlock("loop")
exit   := kernel.NewBlock("exit")

entry.Add(builder.Bra("loop"))           // unconditional branch
entry.Add(builder.Bra("exit").Pred(p))   // @p  bra exit
entry.Add(builder.Bra("exit").PredNot(p))// @!p bra exit
entry.Add(builder.BraUni("exit"))        // bra.uni
entry.Add(builder.Ret())                 // return from device function
entry.Add(builder.Exit())                // exit kernel
entry.Add(builder.Call("fn", retList, argList))
entry.Add(builder.CallIndirect(ptr, retList, argList, proto))
```

---

### Instruction Modifiers (Method Chaining)

Every `*Instruction` supports a fluent modifier API:

| Method | Purpose |
|---|---|
| `.Typed(ptx.Type)` | Set instruction type (`.u32`, `.f32`, etc.) |
| `.From(ptx.Type)` | Set source type for `cvt` |
| `.InSpace(ptx.StateSpace)` | Set state space for `ld`/`st` |
| `.WithMod(ptx.Modifier...)` | Append modifiers (`.lo`, `.hi`, `.wide`, `.sat`, `.ftz`, `.approx`, etc.) |
| `.WithRounding(ptx.RoundingMode)` | Set rounding (`.rn`, `.rz`, `.rm`, `.rp`, `.rs`) |
| `.WithCache(ptx.CacheOp)` | Set cache operator (`.ca`, `.cg`, `.cs`, `.lu`, `.cv`, `.wb`, `.wt`) |
| `.WithScope(ptx.Scope)` | Set memory scope (`.cta`, `.cluster`, `.gpu`, `.sys`) |
| `.WithVec(ptx.VectorSize)` | Set vector width (`.v2`, `.v4`) |
| `.WithBoolOp(ptx.BoolOp)` | Set boolean combiner for `setp` (`.and`, `.or`, `.xor`) |
| `.WithDst2(Operand)` | Set secondary destination for `setp p\|q` |
| `.Pred(reg)` | Guard with `@p` |
| `.PredNot(reg)` | Guard with `@!p` |

---

### Operands

```go
builder.Imm(42)           // immediate integer
builder.ImmU(0xDEAD)      // immediate unsigned
builder.ImmF32(1.0)       // immediate f32
builder.ImmF64(3.14)      // immediate f64
builder.Addr(reg, offset) // memory address [reg+offset]
builder.SReg(ptx.RegTidX) // special register (%tid.x, %ntid.x, ...)
builder.Sym("label")      // named symbol / label
builder.Vec(r0, r1, r2, r3) // vector operand {r0, r1, r2, r3}
kernel.Param("name")      // parameter symbol for ld.param/st.param
```

**Special registers** include thread/block/grid IDs, warp/lane IDs, cluster IDs (sm_90+), lane masks, clock counters, performance monitors, shared memory sizes, and environment registers (`%envreg0`–`%envreg31`).

---

### Arithmetic Instructions

```go
builder.Add(dst, a, b)               // add
builder.Sub(dst, a, b)               // sub
builder.Mul(dst, a, b)               // mul
builder.Mad(dst, a, b, c)            // mad (multiply-add)
builder.Mul24(dst, a, b)             // mul24
builder.Mad24(dst, a, b, c)          // mad24
builder.Sad(dst, a, b, c)            // sad (sum of abs difference)
builder.Div(dst, a, b)               // div
builder.Rem(dst, a, b)               // rem
builder.Abs(dst, src)                // abs
builder.Neg(dst, src)                // neg
builder.Min(dst, a, b)               // min
builder.Max(dst, a, b)               // max
builder.Fma(dst, a, b, c)            // fma
builder.Rcp(dst, src)                // rcp
builder.Sqrt(dst, src)               // sqrt
builder.Rsqrt(dst, src)              // rsqrt
builder.Sin(dst, src)                // sin.approx
builder.Cos(dst, src)                // cos.approx
builder.Lg2(dst, src)                // lg2.approx
builder.Ex2(dst, src)                // ex2.approx
builder.Tanh(dst, src)               // tanh
builder.Copysign(dst, a, b)          // copysign

// Extended precision
builder.AddCC(dst, a, b)             // add.cc
builder.Addc(dst, a, b)              // addc
builder.SubCC(dst, a, b)             // sub.cc
builder.Subc(dst, a, b)              // subc
builder.MadCC(dst, a, b, c)          // mad.cc
builder.Madc(dst, a, b, c)           // madc
```

---

### Bit Manipulation

```go
builder.And(dst, a, b)               // and
builder.Or(dst, a, b)                // or
builder.Xor(dst, a, b)               // xor
builder.Not(dst, src)                // not
builder.Cnot(dst, src)               // cnot
builder.Shl(dst, src, amt)           // shl
builder.Shr(dst, src, amt)           // shr
builder.Bfe(dst, src, start, len)    // bfe (bit field extract)
builder.Bfi(dst, base, ins, s, l)    // bfi (bit field insert)
builder.Bfind(dst, src)              // bfind
builder.Bmsk(dst, pos, width)        // bmsk
builder.Brev(dst, src)               // brev
builder.Clz(dst, src)                // clz
builder.Popc(dst, src)               // popc
builder.Lop3(dst, a, b, c, lut)      // lop3 (3-input logic)
builder.Prmt(dst, a, b, c)           // prmt (byte permute)
builder.Szext(dst, src, size)        // szext
builder.ShfL(dst, a, b, c)           // shf.l (funnel shift left)
builder.ShfR(dst, a, b, c)           // shf.r (funnel shift right)
builder.Fns(dst, mask, base, off)    // fns (find nth set bit)
builder.Dp4a(dst, a, b, c)           // dp4a (4-element dot product)
builder.Dp2a(dst, a, b, c)           // dp2a
```

---

### Comparison & Selection

```go
builder.Setp(ptx.CmpLt, pred, a, b).Typed(ptx.U32)         // setp.lt.u32
builder.Setp(ptx.CmpLt, p, a, b).WithDst2(q).Typed(ptx.S32)// setp p|q
builder.Set(ptx.CmpEq, dst, a, b).Typed(ptx.U32)            // set.eq
builder.Selp(dst, a, b, pred)                                // selp
builder.Slct(dst, a, b, c)                                   // slct
builder.Testp(dst, src)                                      // testp
builder.Min3(dst, a, b, c)                                   // min3 (sm_90+)
builder.Max3(dst, a, b, c)                                   // max3 (sm_90+)

// CmpOp values: CmpEq, CmpNe, CmpLt, CmpLe, CmpGt, CmpGe,
//               CmpLo, CmpLs, CmpHi, CmpHs (unsigned),
//               CmpEqu, CmpNeu, CmpLtu, CmpLeu, CmpGtu, CmpGeu (float unordered),
//               CmpNum, CmpNan
```

---

### Data Movement & Conversion

```go
builder.Mov(dst, src)                   // mov
builder.Ld(dst, addr).InSpace(ptx.Global)   // ld.global
builder.St(addr, src).InSpace(ptx.Shared)   // st.shared
builder.LdParam(dst, kernel.Param("n"))     // ld.param
builder.Ldu(dst, addr)                      // ldu (uniform load)
builder.LdNC(dst, addr)                     // ld.global.nc
builder.LdGlobalNC(dst, addr)               // ld.global.nc (alt)
builder.LdWeak(dst, addr)                   // ld.weak
builder.StWeak(addr, src)                   // st.weak
builder.StAsync(addr, src, mbar)            // st.async
builder.StBulk(addr, size, initVal)         // st.bulk
builder.Cvt(dst, src).Typed(ptx.F32).From(ptx.U32) // cvt
builder.CvtPack(dst, a, b)                  // cvt.pack
builder.Cvta(dst, src)                      // cvta
builder.Prefetch(addr)                      // prefetch
builder.Prefetchu(addr)                     // prefetchu
builder.IsSpacep(dst, addr)                 // isspacep

// Cache operators: CacheCA, CacheCG, CacheCS, CacheLU, CacheCV (loads)
//                  CacheWB, CacheWT (stores)
builder.Ld(dst, addr).Typed(ptx.F32).InSpace(ptx.Global).WithCache(ptx.CacheCS)
```

---

### Atomic Operations

```go
builder.Atom(ptx.ModAtomAdd, dst, addr, src).Typed(ptx.U32).InSpace(ptx.Global)
builder.AtomCAS(dst, addr, compare, val)
builder.AtomExch(ptx.U32, dst, addr, val)
builder.Red(ptx.ModAtomMax, addr, src).Typed(ptx.S32).InSpace(ptx.Global)

// Atomic ops: ModAtomAdd, ModAtomMin, ModAtomMax, ModAtomInc, ModAtomDec,
//             ModAtomAnd, ModAtomOr, ModAtomXor, ModAtomCAS, ModAtomExch
```

---

### Warp & Thread Synchronization

```go
builder.BarSync(id)                        // bar.sync
builder.BarSyncCount(id, count)            // bar.sync id, count
builder.BarWarpSync(membermask)            // bar.warp.sync
builder.BarrierCTA(id)                     // barrier.cta (sm_90+)
builder.BarrierClusterArrive()             // barrier.cluster.arrive
builder.BarrierClusterWait()               // barrier.cluster.wait
builder.Membar(ptx.ModSpaceGlobal)         // membar.gl / .sys / .cta
builder.Fence(ptx.ScopeGPU)               // fence.gpu
builder.FenceAcqRel(ptx.ScopeCTA)         // fence.acq_rel.cta
builder.FenceSC(ptx.ScopeSystem)          // fence.sc.sys
builder.FenceProxy(ptx.ModTensormap)       // fence.proxy.tensormap
builder.FenceProxyAsync(ptx.ScopeGPU)     // fence.proxy.async.gpu
builder.VoteSync(ptx.ModShflBfly, dst, mask, pred) // vote.sync
builder.ReduxSync(ptx.ModAtomAdd, dst, mask, src)  // redux.sync
builder.ShflSync(dst, a, b, c, mask)       // shfl.sync
builder.Activemask(dst)                    // activemask
```

---

### Mbarrier

```go
builder.MbarrierInit(addr, count)
builder.MbarrierInval(addr)
builder.MbarrierArrive(addr)
builder.MbarrierArriveDrop(state, addr, count)
builder.MbarrierArriveNoComplete(state, addr, count)
builder.MbarrierArriveDropNoComplete(state, addr, count)
builder.MbarrierArriveExpectTx(state, addr, txCount)
builder.MbarrierExpectTx(addr, txCount)
builder.MbarrierCompleteTx(addr, txCount)
builder.MbarrierTestWait(done, addr, state)
builder.MbarrierTestWaitParity(done, addr, parity)
builder.MbarrierTryWait(done, addr, state, hint)
builder.MbarrierPendingCount(count, state)
```

---

### Asynchronous Copy

```go
// cp.async (sm_80+)
builder.CpAsync(dst, src, size)
builder.CpAsyncCommitGroup()
builder.CpAsyncWaitGroup(n)
builder.CpAsyncWaitAll()
builder.CpAsyncMbarrierArrive(addr)
builder.CpAsyncMbarrierArriveNoInc(addr)

// cp.async.bulk (sm_90+)
builder.CpAsyncBulk(dst, src, size)
builder.CpAsyncBulkCommitGroup()
builder.CpAsyncBulkWaitGroup(n)
builder.CpAsyncBulkPrefetch(src, size, policy)
builder.CpAsyncBulkTensor(dim, dstMem, tensorMap, coords, mbar, extras)
builder.CpAsyncBulkPrefetchTensor(dim, tensorMap, coords, im2colInfo)

// cp.reduce.async.bulk
builder.CpReduceAsyncBulk(dst, src, size, mbar)
builder.CpReduceAsyncBulkTensor(dim, tensorMap, coords, srcMem)

// multimem (sm_90+)
builder.MultimemCpAsyncBulk(dst, src, size, mask)
builder.MultimemCpReduceAsyncBulk(dst, src, size)
builder.MultimemLdReduce(dst, addr)
builder.MultimemSt(addr, src)
builder.MultimemRed(addr, src)
```

---

### Tensor Map Operations

```go
builder.TensormapReplace(ptx.ModFieldGlobalAddr, addr, newVal)
builder.TensormapCpFenceproxy(dstGlobal, srcShared, size, ptx.ScopeGPU)
builder.Mapa(dst, addr, ctaRank)
builder.GetCTARank(dst, addr)
```

---

### Warp Matrix (WMMA) — sm_70+

```go
builder.WmmaLoad(ptx.ModMatrixA, ptx.ModRow, ptx.ModShapeM16N16K16,
    ptx.F16, dst, addr, stride)
builder.WmmaStore(ptx.ModMatrixD, ptx.ModRow, ptx.ModShapeM16N16K16,
    ptx.F32, addr, src, stride)
builder.WmmaMma(ptx.ModShapeM16N16K16, ptx.ModRow, ptx.ModCol,
    d, a, b, c)
builder.WmmaMmaBitOp(ptx.ModAtomXor, ptx.ModShapeM8N8K128,
    ptx.ModRow, ptx.ModCol, d, a, b, c)

// ldmatrix / stmatrix
builder.Ldmatrix(ptx.ModShapeM8N8, ptx.ModNumX4, ptx.B16, dst, addr)
builder.LdmatrixTrans(ptx.ModShapeM8N8, ptx.ModNumX2, ptx.B16, dst, addr)
builder.Stmatrix(ptx.ModShapeM8N8, ptx.ModNumX4, ptx.B16, addr, src)
builder.Movmatrix(ptx.ModShapeM8N8, ptx.B16, dst, src)
```

---

### MMA — sm_75+

```go
builder.Mma(ptx.ModShapeM16N8K16, ptx.ModRow, ptx.ModCol, d, a, b, c)
builder.MmaSparse(ptx.ModSp, ptx.ModShapeM16N8K32, ptx.ModRow, ptx.ModCol,
    d, a, b, c, metadata, selector)
builder.MmaBlockScaled(ptx.ModShapeM16N8K32, ptx.ModRow, ptx.ModCol,
    ptx.ModScaleVec2x, d, a, b, c, scaleA, selA, scaleB, selB)
builder.MmaSparseBlockScaled(ptx.ModSp, ptx.ModShapeM16N8K32,
    ptx.ModRow, ptx.ModCol, ptx.ModKindMxf8f6f4, ptx.ModScaleVec1x,
    d, a, b, c, meta, sel, scaleA, selA, scaleB, selB)
```

---

### WGMMA — sm_90+

```go
builder.WgmmaFence()
builder.WgmmaCommitGroup()
builder.WgmmaWaitGroup(n)
builder.WgmmaMmaAsync(
    ptx.ModShapeM64N128K16,
    ptx.ModTypeF32, ptx.ModTypeF16, ptx.ModTypeF16,
    d, a, b, scaleD,
    builder.Imm(1), builder.Imm(1), builder.Imm(0), builder.Imm(0),
)
builder.WgmmaMmaAsyncSparse(shape, dtype, atype, btype,
    d, a, b, spMeta, spSel, scaleD)
```

---

### tcgen05 (5th-Gen Tensor Core, sm_100+)

```go
// Allocation
builder.Tcgen05Alloc(ptx.ModCtaGroup1, dstAddr, nCols)
builder.Tcgen05Dealloc(ptx.ModCtaGroup1, tAddr, nCols)
builder.Tcgen05RelinquishAllocPermit(ptx.ModCtaGroup1)

// Load / Store (Tensor Memory ↔ Registers)
builder.Tcgen05Ld(ptx.ModShape16x128b, ptx.ModNumX4, dst, tAddr)
builder.Tcgen05St(ptx.ModShape16x128b, ptx.ModNumX4, tAddr, src)
builder.Tcgen05LdRed(shape, num, ptx.ModRedMax, ptx.ModTypeF32, dst, red, tAddr)

// Copy (Shared Memory → Tensor Memory)
builder.Tcgen05Cp(ptx.ModCtaGroup1, ptx.ModShape128x256b, tAddr, sDesc)

// MMA variants
builder.Tcgen05Mma(ptx.ModKindTf32, ptx.ModCtaGroup1,
    instrDesc, smemDescA, smemDescB, tmemD, tmemC)
builder.Tcgen05MmaWs(ptx.ModKindF16, ptx.ModCtaGroup2,
    dAddr, aDesc, bDesc, iDesc)
builder.Tcgen05MmaSp(ptx.ModKindI8, ptx.ModCtaGroup1,
    dAddr, aDesc, bDesc, spMetaAddr, iDesc)
builder.Tcgen05MmaScaled(ptx.ModKindF8f6f4, ptx.ModCtaGroup1, ptx.ModScaleVec2x,
    dAddr, aDesc, bDesc, iDesc, scaleAAddr, scaleBAddr)

// Commit / Wait / Fence
builder.Tcgen05Commit(ptx.ModCtaGroup1, mbar)
builder.Tcgen05Wait(ptx.ModWaitLd)
builder.Tcgen05Fence(ptx.ModWaitSt)
builder.Tcgen05FenceSync(ptx.ModBeforeThreadSync)
builder.Tcgen05Shift(ptx.ModCtaGroup1, tAddr)
```

---

### Cluster Launch Control — sm_100+

```go
builder.ClusterlaunchcontrolTryCancel(addr, mbar)
builder.ClusterlaunchcontrolQueryCancelIsCanceled(pred, handle)
builder.ClusterlaunchcontrolQueryCancelGetFirstCTAId(dstVec, handle)
builder.ClusterlaunchcontrolQueryCancelGetFirstCTAIdDim(dst, handle, ptx.ModDimX)
```

---

### Texture & Surface Operations

```go
builder.Tex(ptx.ModGeom2D, dst, texRef, sampler, coords)
builder.Tld4(ptx.ModCompR, ptx.ModGeom2D, dst, texRef, sampler, coords)
builder.Txq(ptx.ModQueryWidth, dst, texRef, lod)
builder.Suld(ptx.ModGeom2D, dst, surfRef, coords)
builder.Sust(ptx.ModGeom2D, surfRef, coords, val)
builder.Sured(ptx.ModGeom2D, surfRef, coords, val)
builder.Suq(ptx.ModQueryWidth, dst, surfRef)
```

---

### Video / SIMD Instructions

```go
builder.Vadd(ptx.ModTypeU32, ptx.ModTypeS32, ptx.ModTypeS32, dst, a, b)
builder.Vsub(ptx.ModTypeU32, ptx.ModTypeU32, ptx.ModTypeU32, dst, a, b)
builder.Vmax(ptx.ModTypeS32, ptx.ModTypeS32, ptx.ModTypeS32, dst, a, b)
builder.Vmin(ptx.ModTypeU32, ptx.ModTypeU32, ptx.ModTypeU32, dst, a, b)
builder.Vabsdiff(ptx.ModTypeS32, ptx.ModTypeS32, ptx.ModTypeS32, dst, a, b)
builder.Vmad(ptx.ModTypeS32, ptx.ModTypeS32, ptx.ModTypeS32, dst, a, b, c)
builder.Vshl(ptx.ModTypeU32, ptx.ModTypeU32, dst, a, b)
builder.Vshr(ptx.ModTypeU32, ptx.ModTypeU32, dst, a, b)
builder.Vset(ptx.ModTypeS32, ptx.ModTypeS32, ptx.CmpLt, dst, a, b)

// SIMD2 / SIMD4
builder.Vadd2(dtype, atype, btype, dst, a, b)
builder.Vsub4(dtype, atype, btype, dst, a, b)
builder.Vavrg2(dtype, atype, btype, dst, a, b)
builder.Vset4(atype, btype, ptx.CmpEq, dst, a, b)
```

---

### Miscellaneous

```go
builder.Alloca(dst, size)              // stack allocation
builder.StackSave(dst)                 // save stack pointer
builder.StackRestore(src)              // restore stack pointer
builder.NanoSleep(t)                   // nanosleep
builder.Brkpt()                        // breakpoint
builder.Trap()                         // abort / interrupt
builder.Discard(addr, size)            // memory discard hint
builder.CreatePolicy(dst, args...)     // create eviction policy
builder.ApplyPriority(addr, size)      // apply eviction priority
builder.Istypep(ptx.ModTypeTexRef, pred, addr) // type test
builder.Pmevent(builder.Imm(3))        // performance monitor event
builder.PmeventMask(mask)              // performance monitor mask
builder.SetMaxNReg(ptx.ModInc, builder.Imm(64)) // warp register hint
```

---

### Module-Scope Globals

```go
// Scalar global
mod.AddGlobal(builder.NewGlobal("g_val", ptx.Global, ptx.F32))

// Array with initializer
mod.AddGlobal(
    builder.NewGlobalArray("lut", ptx.Const, ptx.B32, 16).
        WithInit(int64(0), int64(1), int64(2), /* ... */),
)

// Shared memory with alignment
mod.AddGlobal(
    builder.NewGlobalArray("smem", ptx.Shared, ptx.F32, 256).WithAlign(16),
)

// Managed memory (UVM)
mod.AddGlobal(
    builder.NewGlobal("g_managed", ptx.Global, ptx.S32).
        WithAttribute(builder.Managed()).
        WithLinkage(ptx.LinkVisible),
)

// Unified memory
mod.AddGlobal(
    builder.NewGlobal("g_unified", ptx.Global, ptx.S32).
        WithAttribute(builder.Unified(uuid1, uuid2)),
)
```

---

### Types Reference

| Category | Types |
|---|---|
| Predicate | `Pred` |
| Bit-size | `B8`, `B16`, `B32`, `B64`, `B128` |
| Signed int | `S8`, `S16`, `S32`, `S64` |
| Unsigned int | `U8`, `U16`, `U32`, `U64` |
| Float | `F16`, `F32`, `F64`, `BF16`, `TF32` |
| Packed float | `F16x2`, `BF16x2`, `F32x2` |
| MX float (sm_100+) | `E2M1`, `E2M3`, `E3M2`, `E4M3`, `E5M2`, `E8M0` |
| Packed MX | `E4M3x2`, `E5M2x2`, `E4M3x4`, `E5M2x4`, `UE8M0x2`, `E2M1x4`, etc. |
| Packed int | `U16x2`, `S16x2` |
| Sub-byte tensor | `B4x16`, `B4x16_p64`, `B6x16_p32`, `B6p2x16` |
| Opaque | `TexRef`, `SamplerRef`, `SurfRef`, `TensorMap` |

### State Spaces Reference

| Constant | PTX | Description |
|---|---|---|
| `ptx.Reg` | `.reg` | Per-thread registers |
| `ptx.Global` | `.global` | Global memory |
| `ptx.Shared` | `.shared` | Per-CTA shared memory |
| `ptx.SharedCTA` | `.shared::cta` | Explicit CTA shared |
| `ptx.SharedCluster` | `.shared::cluster` | Cluster shared (sm_90+) |
| `ptx.Const` | `.const` | Read-only constant memory |
| `ptx.Local` | `.local` | Per-thread local memory |
| `ptx.Param` | `.param` | Kernel/function parameters |
| `ptx.ParamEntry` | `.param::entry` | Explicit kernel params |
| `ptx.ParamFunc` | `.param::func` | Explicit function params |

---

## Roadmap

### Complete (Sections 1–10)
- PTX module structure, versioning, and target architectures
- All fundamental data types and state spaces
- Full arithmetic, logic, shift, and bit-manipulation instructions
- All comparison, selection, and predication instructions
- Complete data movement: `ld`, `st`, `mov`, `cvt`, `cvta`, `prefetch`, etc.
- Atomic and reduction operations
- Control flow: branches, calls, returns, predicates
- Warp synchronization: `bar`, `membar`, `fence`, `vote`, `shfl`, `redux`
- Mbarrier primitives (sm_80+)
- Asynchronous copy: `cp.async`, `cp.async.bulk`, `cp.reduce.async.bulk` (sm_80+/sm_90+)
- Multimem operations (sm_90+)
- Tensor map operations and `tensormap.replace`
- WMMA (`wmma.load`, `wmma.store`, `wmma.mma`) — sm_70+
- `ldmatrix`, `stmatrix`, `movmatrix` — sm_75+
- MMA (`mma.sync`, sparse MMA, block-scaled MMA) — sm_75+
- WGMMA (`wgmma.mma_async`, sparse, fence, commit, wait) — sm_90+
- tcgen05 5th-gen Tensor Core (alloc, dealloc, ld, st, cp, mma, commit, wait, fence) — sm_100+
- Cluster management and `clusterlaunchcontrol` — sm_100+
- Texture (`tex`, `tld4`, `txq`) and surface (`suld`, `sust`, `sured`, `suq`) operations
- Video / SIMD instructions (`vadd`, `vsub`, `vmad`, `vset`, SIMD2/4 variants, etc.)
- Special registers (thread IDs, warp IDs, cluster IDs, clocks, performance monitors, etc.)
- Module-scope globals with initializers, alignment, linkage, and attributes (`.managed`, `.unified`)
- Performance-tuning directives (`.maxnreg`, `.maxntid`, `.reqntid`, `.pragma`, `.explicitcluster`, etc.)
- Stack management (`alloca`, `stacksave`, `stackrestore`)
- All miscellaneous instructions (`nanosleep`, `brkpt`, `trap`, `discard`, `pmevent`, `setmaxnreg`, etc.)

### Upcoming (Sections 11–14)
- Parallel Synchronization and Communication (advanced warp group operations)
- Additional performance instrumentation primitives
- Profiling and debug support

## License

MIT License.