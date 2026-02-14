package builder

import (
	"github.com/arc-language/ptx-gen/ptx"
)

// ---- Integer & Float Arithmetic ----

func Add(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAdd, Dst: dst, Src: []Operand{src0, src1}}
}

func Sub(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSub, Dst: dst, Src: []Operand{src0, src1}}
}

func Mul(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMul, Dst: dst, Src: []Operand{src0, src1}}
}

func Mad(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMad, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Fma(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpFma, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Div(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpDiv, Dst: dst, Src: []Operand{src0, src1}}
}

func Rem(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpRem, Dst: dst, Src: []Operand{src0, src1}}
}

func Abs(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpAbs, Dst: dst, Src: []Operand{src}}
}

func Neg(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpNeg, Dst: dst, Src: []Operand{src}}
}

func Min(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMin, Dst: dst, Src: []Operand{src0, src1}}
}

func Max(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMax, Dst: dst, Src: []Operand{src0, src1}}
}

// ---- Float-only math ----

func Rcp(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpRcp, Dst: dst, Src: []Operand{src}}
}

func Sqrt(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpSqrt, Dst: dst, Src: []Operand{src}}
}

func Rsqrt(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpRsqrt, Dst: dst, Src: []Operand{src}}
}

func Sin(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpSin, Dst: dst, Src: []Operand{src}}
}

func Cos(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCos, Dst: dst, Src: []Operand{src}}
}

func Lg2(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpLg2, Dst: dst, Src: []Operand{src}}
}

func Ex2(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpEx2, Dst: dst, Src: []Operand{src}}
}

// ---- Bit manipulation ----

func Popc(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpPopc, Dst: dst, Src: []Operand{src}}
}

func Clz(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpClz, Dst: dst, Src: []Operand{src}}
}

func Brev(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpBrev, Dst: dst, Src: []Operand{src}}
}

func Bfe(dst, src, start, len Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfe, Dst: dst, Src: []Operand{src, start, len}}
}

func Bfi(dst, src, base, start, len Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfi, Dst: dst, Src: []Operand{src, base, start, len}}
}

// ---- Logic & shift ----

func And(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAnd, Dst: dst, Src: []Operand{src0, src1}}
}

func Or(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpOr, Dst: dst, Src: []Operand{src0, src1}}
}

func Xor(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpXor, Dst: dst, Src: []Operand{src0, src1}}
}

func Not(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpNot, Dst: dst, Src: []Operand{src}}
}

func Shl(dst, src, amount Operand) *Instruction {
	return &Instruction{Op: ptx.OpShl, Dst: dst, Src: []Operand{src, amount}}
}

func Shr(dst, src, amount Operand) *Instruction {
	return &Instruction{Op: ptx.OpShr, Dst: dst, Src: []Operand{src, amount}}
}

// ---- Comparison & selection ----

func Setp(cmp ptx.CmpOp, dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSetp, Cmp: cmp, Dst: dst, Src: []Operand{src0, src1}}
}

func Set(cmp ptx.CmpOp, dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSet, Cmp: cmp, Dst: dst, Src: []Operand{src0, src1}}
}

func Selp(dst, srcTrue, srcFalse, pred Operand) *Instruction {
	return &Instruction{Op: ptx.OpSelp, Dst: dst, Src: []Operand{srcTrue, srcFalse, pred}}
}

func Slct(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSlct, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// ---- Data movement ----

func Mov(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpMov, Dst: dst, Src: []Operand{src}}
}

func Ld(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLd, Dst: dst, Src: []Operand{addr}}
}

// LdNC creates a non-coherent global load (ld.global.nc).
func LdNC(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdNC, Dst: dst, Src: []Operand{addr}, Space: ptx.Global}
}

func LdParam(dst Operand, param *Symbol) *Instruction {
	return &Instruction{
		Op:    ptx.OpLd,
		Space: ptx.Param,
		Dst:   dst,
		Src:   []Operand{&Address{Base: param, Offset: 0}},
	}
}

func St(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpSt, Src: []Operand{addr, src}}
}

func Cvt(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvt, Dst: dst, Src: []Operand{src}}
}

// CvtPack creates a cvt.pack instruction (packing two 16-bit values into 32-bit, etc.).
func CvtPack(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvtPack, Dst: dst, Src: []Operand{src0, src1}}
}

func Cvta(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvta, Dst: dst, Src: []Operand{src}}
}

func Prefetch(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpPrefetch, Src: []Operand{addr}}
}

// ---- Stack Manipulation (Section 7.3) ----

// Alloca allocates memory on the local stack.
// size can be an immediate (constant size) or a register (variable size).
func Alloca(dst, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpAlloca, Dst: dst, Src: []Operand{size}}
}

// StackSave saves the current stack pointer to a register.
func StackSave(dst Operand) *Instruction {
	return &Instruction{Op: ptx.OpStackSave, Dst: dst}
}

// StackRestore restores the stack pointer from a register.
func StackRestore(src Operand) *Instruction {
	return &Instruction{Op: ptx.OpStackRestore, Src: []Operand{src}}
}

// ---- Warp shuffle ----

func ShflSync(mode ptx.Modifier, dst, mask, src, offset, clamp Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpShfl,
		Dst:       dst,
		Src:       []Operand{mask, src, offset, clamp},
		Modifiers: []ptx.Modifier{mode},
	}
}

// ---- Control flow ----

func Bra(label string) *Instruction {
	return &Instruction{Op: ptx.OpBra, Src: []Operand{&Symbol{Name: label}}}
}

func BraUni(label string) *Instruction {
	return &Instruction{
		Op:        ptx.OpBra,
		Src:       []Operand{&Symbol{Name: label}},
		Modifiers: []ptx.Modifier{ptx.ModUni},
	}
}

func Call(target string, retParams []Operand, args []Operand) *Instruction {
	inst := &Instruction{
		Op:         ptx.OpCall,
		Src:        args,
		CallTarget: target,
	}
	// If there are return values, first operand is the return dest
	if len(retParams) > 0 {
		inst.Dst = retParams[0]
		// If multiple returns, they are effectively handled by the emitter logic
		// or packed into the Dst as a VectorOp if strictly following IR
	}
	return inst
}

func Ret() *Instruction {
	return &Instruction{Op: ptx.OpRet}
}

func Exit() *Instruction {
	return &Instruction{Op: ptx.OpExit}
}

// ---- Synchronization ----

func BarSync(id Operand) *Instruction {
	return &Instruction{Op: ptx.OpBar, Src: []Operand{id}}
}

func BarSyncCount(id, count Operand) *Instruction {
	return &Instruction{Op: ptx.OpBar, Src: []Operand{id, count}}
}

func Membar(level ptx.Modifier) *Instruction {
	return &Instruction{Op: ptx.OpMembar, Modifiers: []ptx.Modifier{level}}
}

func Fence(scope ptx.Scope) *Instruction {
	return &Instruction{Op: ptx.OpFence, Scope: scope}
}

// ---- Atomics ----

func Atom(op ptx.Modifier, dst, addr, src Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpAtom,
		Dst:       dst,
		Src:       []Operand{addr, src},
		Modifiers: []ptx.Modifier{op},
	}
}

func AtomCAS(dst, addr, compare, val Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpAtom,
		Dst:       dst,
		Src:       []Operand{addr, compare, val},
		Modifiers: []ptx.Modifier{ptx.ModAtomCAS},
	}
}

func Red(op ptx.Modifier, addr, src Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpRed,
		Src:       []Operand{addr, src},
		Modifiers: []ptx.Modifier{op},
	}
}

// ---- Warp voting ----

func VoteSync(mode ptx.Modifier, dst, mask, pred Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpVoteSync,
		Dst:       dst,
		Src:       []Operand{mask, pred},
		Modifiers: []ptx.Modifier{mode},
	}
}

func Activemask(dst Operand) *Instruction {
	return &Instruction{Op: ptx.OpActivemask, Dst: dst}
}

func ReduxSync(op ptx.Modifier, dst, mask, src Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpReduxSync,
		Dst:       dst,
		Src:       []Operand{mask, src},
		Modifiers: []ptx.Modifier{op},
	}
}

// ---- Async copy & Matrix ----

func CpAsync(dst, src Operand, bytes int) *Instruction {
	return &Instruction{
		Op:  ptx.OpCpAsync,
		Dst: dst,
		Src: []Operand{src, Imm(int64(bytes))},
	}
}

func CpAsyncCommitGroup() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncCommitGroup}
}

func CpAsyncWaitGroup(n int) *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncWaitGroup, Src: []Operand{Imm(int64(n))}}
}

// CpAsyncBulk performs a bulk asynchronous copy (Tensor Memory Copy).
func CpAsyncBulk(dst, src, size Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpCpAsyncBulk,
		Src: []Operand{dst, src, size},
	}
}

// LdMatrix loads a matrix fragment from shared memory.
func LdMatrix(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdMatrix, Dst: dst, Src: []Operand{addr}}
}

// StMatrix stores a matrix fragment to shared memory.
func StMatrix(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpStMatrix, Src: []Operand{addr, src}}
}

// Wgmma performs a Warpgroup Matrix Multiply Accumulate.
// This is a complex instruction often taking many operands/descriptors.
func Wgmma(dst Operand, args ...Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpWgmma,
		Dst: dst,
		Src: args,
	}
}

// MbarrierInit initializes an mbarrier object.
func MbarrierInit(addr, count Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierInit, Src: []Operand{addr, count}}
}

// MbarrierArrive performs an arrive operation on an mbarrier.
func MbarrierArrive(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierArrive, Src: []Operand{addr}}
}


// ---- Memory Consistency & Fences (Section 8) ----

// FenceSC creates a sequentially consistent fence (fence.sc.scope).
// Scope determines the visibility (e.g., .cta, .gpu, .sys).
// Reference: Section 8.10.2 Fence-SC
func FenceSC(scope ptx.Scope) *Instruction {
	return &Instruction{
		Op:        ptx.OpFence,
		Scope:     scope,
		Modifiers: []ptx.Modifier{ptx.ModSC},
	}
}

// FenceAcqRel creates an acquire-release fence (fence.acq_rel.scope).
// Reference: Section 8.4 Operation types
func FenceAcqRel(scope ptx.Scope) *Instruction {
	return &Instruction{
		Op:        ptx.OpFence,
		Scope:     scope,
		Modifiers: []ptx.Modifier{ptx.ModAcqRel},
	}
}

// FenceProxy creates a proxy fence for aliasing or async operations.
// kind is typically ptx.ModAlias or ptx.ModAsync.
// Reference: Section 8.6 Proxies
func FenceProxy(kind ptx.Modifier) *Instruction {
	return &Instruction{
		Op:        ptx.OpFence,
		Modifiers: []ptx.Modifier{ptx.ModProxy, kind},
	}
}

// MembarProxy creates a proxy membar (membar.proxy).
// Reference: Section 8.4 Operation types
func MembarProxy() *Instruction {
	return &Instruction{
		Op:        ptx.OpMembar,
		Modifiers: []ptx.Modifier{ptx.ModProxy},
	}
}

// CpAsyncMbarrierArrive performs an asynchronous arrive on an mbarrier object.
// Reference: Section 8.9.1.1 Asynchronous Operations
func CpAsyncMbarrierArrive(addr Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpCpAsyncMbarrierArrive,
		Src: []Operand{addr},
	}
}

// LdWeak creates a weak load (ld.weak).
// Reference: Section 8.4 Operation types
func LdWeak(dst, addr Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpLd,
		Dst:       dst,
		Src:       []Operand{addr},
		Modifiers: []ptx.Modifier{ptx.ModWeak},
	}
}

// StWeak creates a weak store (st.weak).
// Reference: Section 8.4 Operation types
func StWeak(addr, src Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpSt,
		Src:       []Operand{addr, src},
		Modifiers: []ptx.Modifier{ptx.ModWeak},
	}
}