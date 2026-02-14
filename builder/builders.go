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

// Max finds the maximum of two values.
// Supports .relu via WithMod(ptx.ModRelu).
// Reference: Section 9.7.1.13
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

// Ex2 computes the base-2 exponential (2^a).
// Reference: Section 9.7.4.10
func Ex2(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpEx2, Dst: dst, Src: []Operand{src}}
}

// ---- Bit manipulation ----
// Popc counts the number of one bits (population count).
// Reference: Section 9.7.1.14
func Popc(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpPopc, Dst: dst, Src: []Operand{src}}
}

// Clz counts the leading zeros.
// Reference: Section 9.7.1.15
func Clz(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpClz, Dst: dst, Src: []Operand{src}}
}

// Brev performs bitwise reversal.
// Reference: Section 9.7.1.18
func Brev(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpBrev, Dst: dst, Src: []Operand{src}}
}

// Bfe extracts a bit field (Bit Field Extract).
// Reference: Section 9.7.1.19
func Bfe(dst, src, start, len Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfe, Dst: dst, Src: []Operand{src, start, len}}
}

// Bfi inserts a bit field (Bit Field Insert).
// Reference: Section 9.7.1.20
func Bfi(dst, base, insert, start, len Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfi, Dst: dst, Src: []Operand{base, insert, start, len}}
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

// Usage: Setp(ptx.CmpLt, p, a, b) or Setp(..., p, a, b).WithDst2(q)
func Setp(cmp ptx.CmpOp, dst, a, b Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpSetp,
		Cmp: cmp,
		Dst: dst,
		Src: []Operand{a, b},
	}
}

// Set compares two values and optionally applies a boolean op.
// Usage: Set(ptx.CmpEq, dst, a, b).WithBoolOp(ptx.BoolAnd)
func Set(cmp ptx.CmpOp, dst, a, b Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpSet,
		Cmp: cmp,
		Dst: dst,
		Src: []Operand{a, b},
	}
}

// Selp selects between a and b based on predicate p.
// d = (p == 1) ? a : b
func Selp(dst, a, b, p Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpSelp,
		Dst: dst,
		Src: []Operand{a, b, p},
	}
}

// Slct selects based on the sign of c.
// d = (c >= 0) ? a : b
func Slct(dst, a, b, c Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpSlct,
		Dst: dst,
		Src: []Operand{a, b, c},
	}
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

// Mul24 multiplies two 24-bit integer values.
// Reference: Section 9.7.1.5
func Mul24(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMul24, Dst: dst, Src: []Operand{src0, src1}}
}

// Mad24 multiplies two 24-bit integer values and adds a third value.
// Reference: Section 9.7.1.6
func Mad24(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMad24, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// Fns finds the n-th set bit.
// Reference: Section 9.7.1
func Fns(dst, mask, base, offset Operand) *Instruction {
	return &Instruction{Op: ptx.OpFns, Dst: dst, Src: []Operand{mask, base, offset}}
}

// Bmsk creates a bit mask.
// Reference: Section 9.7.1
func Bmsk(dst, pos, width Operand) *Instruction {
	return &Instruction{Op: ptx.OpBmsk, Dst: dst, Src: []Operand{pos, width}}
}

// Szext performs sign or zero extension.
// Reference: Section 9.7.1
func Szext(dst, src, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpSzext, Dst: dst, Src: []Operand{src, size}}
}

// Dp4a performs a 4-way dot product with accumulation.
// Reference: Section 9.7.1
func Dp4a(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpDp4a, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// Dp2a performs a 2-way dot product with accumulation.
// Reference: Section 9.7.1
func Dp2a(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpDp2a, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// Bfind finds the most significant non-sign bit.
// Use WithMod(ptx.ModShiftAmt) for bfind.shiftamt.
// Reference: Section 9.7.1.16
func Bfind(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfind, Dst: dst, Src: []Operand{src}}
}

// ---- Extended-Precision Integer Arithmetic (Section 9.7.2) ----

// AddCC performs addition and writes the carry-out to CC.CF (add.cc).
func AddCC(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAddCC, Dst: dst, Src: []Operand{src0, src1}}
}

// Addc performs addition with carry-in (addc).
// Use WithMod(ptx.ModCC) for addc.cc (write carry-out).
func Addc(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAddc, Dst: dst, Src: []Operand{src0, src1}}
}

// SubCC performs subtraction and writes the borrow-out to CC.CF (sub.cc).
func SubCC(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSubCC, Dst: dst, Src: []Operand{src0, src1}}
}

// Subc performs subtraction with borrow-in (subc).
// Use WithMod(ptx.ModCC) for subc.cc (write borrow-out).
func Subc(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSubc, Dst: dst, Src: []Operand{src0, src1}}
}

// MadCC performs multiply-add and writes carry-out to CC.CF (mad.cc).
func MadCC(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMadCC, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// Madc performs multiply-add with carry-in (madc).
// Use WithMod(ptx.ModCC) for madc.cc.
func Madc(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMadc, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// ---- Floating-Point Instructions (Section 9.7.3) ----

// Testp tests a floating-point property.
// Use WithMod(...) to specify the property (e.g., ptx.ModFinite, ptx.ModNotANumber).
// Reference: Section 9.7.3.1
func Testp(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpTestp, Dst: dst, Src: []Operand{src}}
}

// Copysign copies the sign of src0 to src1.
// Reference: Section 9.7.3.2
func Copysign(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpCopysign, Dst: dst, Src: []Operand{src0, src1}}
}


// Min3 finds the minimum of three values.
// Reference: Section 9.7.3.11 (Introduced in PTX 8.8)
func Min3(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMin,
		Dst: dst,
		Src: []Operand{src0, src1, src2},
	}
}

// Max3 finds the maximum of three values.
// Reference: Section 9.7.3.12 (Introduced in PTX 8.8)
func Max3(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMax,
		Dst: dst,
		Src: []Operand{src0, src1, src2},
	}
}

// Tanh calculates the hyperbolic tangent.
// Reference: Section 9.7.3.22
func Tanh(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpTanh, Dst: dst, Src: []Operand{src}}
}