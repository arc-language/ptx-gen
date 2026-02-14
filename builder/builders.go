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

// And performs bitwise AND.
func And(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAnd, Dst: dst, Src: []Operand{src0, src1}}
}

// Or performs bitwise OR.
func Or(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpOr, Dst: dst, Src: []Operand{src0, src1}}
}

// Xor performs bitwise XOR.
func Xor(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpXor, Dst: dst, Src: []Operand{src0, src1}}
}

// Not performs bitwise negation (one's complement).
func Not(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpNot, Dst: dst, Src: []Operand{src}}
}

func Shl(dst, src, amount Operand) *Instruction {
	return &Instruction{Op: ptx.OpShl, Dst: dst, Src: []Operand{src, amount}}
}

// Shr performs shift right.
func Shr(dst, src, amt Operand) *Instruction {
	return &Instruction{Op: ptx.OpShr, Dst: dst, Src: []Operand{src, amt}}
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

// Mov performs register-to-register move, load address, or vector pack/unpack.
func Mov(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpMov, Dst: dst, Src: []Operand{src}}
}

// Ld performs a load from memory.
// Use .InSpace() for .global, .shared, etc.
// Use .WithVolatile(), .WithCache(), .WithMod() for qualifiers.
// Reference: Section 9.7.9.8
func Ld(dst, addr Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpLd,
		Dst: dst,
		Src: []Operand{addr},
	}
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

// Cvt converts values between types.
// Supports 1 source (standard), 2 sources (packed), or more (rbits, scale).
// Use Typed() for destination type and SourceTyped() for source type.
// Reference: Section 9.7.9.21
func Cvt(dst Operand, srcs ...Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpCvt,
		Dst: dst,
		Src: srcs,
	}
}

func Cvta(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvta, Dst: dst, Src: []Operand{src}}
}

// Prefetch brings a cache line into the specified level.
// Usage: Prefetch(addr).InSpace(ptx.Global).WithMod(ptx.ModL1)
func Prefetch(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpPrefetch, Src: []Operand{addr}}
}

// Prefetchu brings data into the uniform cache.
func Prefetchu(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpPrefetchu, Src: []Operand{addr}}
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

// ShflSync exchanges data within a warp with synchronization.
// Usage: ShflSync(dst, a, b, c, mask).WithMod(ptx.ModShflUp)
// Reference: Section 9.7.9.6
func ShflSync(dst, a, b, c, mask Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpShfl, // Uses OpShfl, but modifier .sync makes it shfl.sync
		Dst: dst,
		Src: []Operand{a, b, c, mask},
		Modifiers: []ptx.Modifier{ptx.ModSync},
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


// Cnot performs logical negation (C-style !).
func Cnot(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCnot, Dst: dst, Src: []Operand{src}}
}

// Lop3 performs an arbitrary logical operation on 3 inputs.
// immLut is an integer constant (0-255).
// Use WithBoolOp() and WithDst2() for the predicate output variant.
// Reference: Section 9.7.8.6
func Lop3(dst, a, b, c, immLut Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpLop3,
		Dst: dst,
		Src: []Operand{a, b, c, immLut},
	}
}

// Shf performs a funnel shift.
// Use WithMod(ptx.ModLo) for left shift (.l)
// Use WithMod(ptx.ModHi) is NOT used; .r is used for right shift.
// NOTE: We need new modifiers ModShfL and ModShfR because .l and .r are distinct
// from .lo/.hi in meaning here, though .l/.r syntax is specific to shf.
// Let's reuse ModLo/ModHi if they map to .lo/.hi, but shf uses .l/.r.
// actually shf uses .l and .r. You might need to add ModL and ModR to modifier.go
// or just handle it via opcode if you prefer shf.l and shf.r as separate ops.
// For now, let's assume specific opcodes or modifiers.
//
// Recommendation: Use ptx.OpShfL and ptx.OpShfR in opcode.go or modifiers.
// Let's implement ShfL and ShfR as builders.
func ShfL(dst, a, b, c Operand) *Instruction {
	// shf.l.mode.b32
	return &Instruction{
		Op:  ptx.OpShf,
		Dst: dst,
		Src: []Operand{a, b, c},
		// We need a way to specify .l vs .r.
		// A common way is a modifier.
		Modifiers: []ptx.Modifier{ptx.ModLeft},
	}
}

// Shfl (Deprecated) exchanges data within a warp.
// Usage: Shfl(dst, a, b, c).WithMod(ptx.ModShflUp).WithDst2(p)
// Reference: Section 9.7.9.5
func Shfl(dst, a, b, c Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpShfl,
		Dst: dst,
		Src: []Operand{a, b, c},
	}
}

func ShfR(dst, a, b, c Operand) *Instruction {
	// shf.r.mode.b32
	return &Instruction{
		Op:        ptx.OpShf,
		Dst:       dst,
		Src:       []Operand{a, b, c},
		Modifiers: []ptx.Modifier{ptx.ModRight},
	}
}

// Prmt permutes bytes.
// Usage: Prmt(dst, a, b, c).WithMod(ptx.ModF4e)
// Reference: Section 9.7.9.7
func Prmt(dst, a, b, c Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpPrmt,
		Dst: dst,
		Src: []Operand{a, b, c},
	}
}

// LdGlobalNC performs a load via non-coherent cache.
// Reference: Section 9.7.9.9
func LdGlobalNC(dst, addr Operand) *Instruction {
	return &Instruction{
		Op:    ptx.OpLdNC, // Maps to ld.global.nc
		Dst:   dst,
		Src:   []Operand{addr},
		Space: ptx.Global, // Implicitly global
	}
}


// Ldu loads read-only data uniform across a warp.
// Reference: Section 9.7.9.10
func Ldu(dst, addr Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpLdu,
		Dst: dst,
		Src: []Operand{addr},
	}
}

// ---- Store (Section 9.7.9.11) ----

// St stores data to memory.
// Use .InSpace(), .WithCache(), .WithScope(), .WithMod() for qualifiers.
// Reference: Section 9.7.9.11
func St(addr, src Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpSt,
		Src: []Operand{addr, src},
	}
}

// ---- Async Store (Section 9.7.9.12) ----

// StAsync performs an asynchronous store.
// Use WithMod() for .mbarrier::complete_tx::bytes and semantics (.weak, .release).
// Reference: Section 9.7.9.12
func StAsync(addr, src Operand, mbar Operand) *Instruction {
	// If mbar is provided, it's the 3-operand variant (weak).
	// If mbar is nil, it's the 2-operand variant (release/mmio).
	if mbar != nil {
		return &Instruction{
			Op:  ptx.OpStAsync,
			Src: []Operand{addr, src, mbar},
		}
	}
	return &Instruction{
		Op:  ptx.OpStAsync,
		Src: []Operand{addr, src},
	}
}

// ---- Bulk Store (Section 9.7.9.13) ----

// StBulk initializes a memory region. initVal must be 0.
// Reference: Section 9.7.9.13
func StBulk(addr, size, initVal Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpStBulk,
		Src: []Operand{addr, size, initVal},
	}
}

// ---- Multimem Operations (Section 9.7.9.14) ----

// MultimemLdReduce loads and reduces from multiple memory locations.
// Use WithMod() for .op (.add, .min, etc.), scope, and semantics.
func MultimemLdReduce(dst, addr Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMultimemLdReduce,
		Dst: dst,
		Src: []Operand{addr},
	}
}

// MultimemSt stores to multiple memory locations.
func MultimemSt(addr, src Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMultimemSt,
		Src: []Operand{addr, src},
	}
}

// MultimemRed performs reduction on multiple memory locations.
// Use WithMod() for .op (e.g. .add).
func MultimemRed(addr, src Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMultimemRed,
		Src: []Operand{addr, src},
	}
}

// ApplyPriority applies eviction priority to an address range.
// Reference: Section 9.7.9.16
func ApplyPriority(addr, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpApplyPriority, Src: []Operand{addr, size}}
}

// Discard hints that data can be discarded without write-back.
// Reference: Section 9.7.9.17
func Discard(addr, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpDiscard, Src: []Operand{addr, size}}
}

// CreatePolicy creates a cache eviction policy.
// Supports variable args for range (4 ops) or fractional (2 ops) variants.
// Reference: Section 9.7.9.18
func CreatePolicy(dst Operand, args ...Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpCreatePolicy,
		Dst: dst,
		Src: args,
	}
}


// IsSpacep queries if an address falls within a state space.
// Usage: IsSpacep(p, addr).InSpace(ptx.Shared)
// Reference: Section 9.7.9.19
func IsSpacep(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpIsSpacep, Dst: dst, Src: []Operand{addr}}
}


// ---- Packed Conversion (Section 9.7.9.22) ----

// CvtPack converts two integers and packs them.
// Usage: CvtPack(d, a, b).Typed(ptx.U8).SourceTyped(ptx.S32).WithMod(ptx.ModSat)
// For 3-input version: CvtPack(d, a, b, c)...
func CvtPack(dst Operand, srcs ...Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpCvtPack,
		Dst: dst,
		Src: srcs,
	}
}

// ---- Address Mapping (Section 9.7.9.23 - 9.7.9.24) ----

// Mapa maps a shared memory address to a target CTA rank.
func Mapa(dst, addr, ctaRank Operand) *Instruction {
	return &Instruction{Op: ptx.OpMapa, Dst: dst, Src: []Operand{addr, ctaRank}}
}

// GetCTARank gets the CTA rank of an address.
func GetCTARank(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpGetCTARank, Dst: dst, Src: []Operand{addr}}
}

// ---- Async Copy (Section 9.7.9.25) ----

// CpAsync initiates an async copy.
// Use WithMod() for .ca, .cg, .shared, .global, etc.
// Reference: Section 9.7.9.25.3.1
func CpAsync(dst, src, size Operand, args ...Operand) *Instruction {
	srcs := []Operand{dst, src, size}
	srcs = append(srcs, args...)
	return &Instruction{
		Op:  ptx.OpCpAsync,
		Src: srcs,
	}
}

// CpAsyncCommitGroup commits uncommitted cp.async instructions.
func CpAsyncCommitGroup() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncCommitGroup}
}

// CpAsyncWaitAll waits for all cp.async operations.
func CpAsyncWaitAll() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncWaitAll}
}

// CpAsyncWaitGroup waits for N groups to complete.
func CpAsyncWaitGroup(n int64) *Instruction {
	return &Instruction{
		Op:  ptx.OpCpAsyncWaitGroup,
		Src: []Operand{Imm(n)},
	}
}

// CpAsyncBulk initiates a bulk async copy.
// Reference: Section 9.7.9.25.4.1
func CpAsyncBulk(dst, src, size Operand, args ...Operand) *Instruction {
	srcs := []Operand{dst, src, size}
	srcs = append(srcs, args...)
	return &Instruction{
		Op:  ptx.OpCpAsyncBulk,
		Src: srcs,
	}
}




// For global destination:
//   CpReduceAsyncBulk(dstMem, srcMem, size, nil)
//      .WithMod(ptx.ModAtomAdd, ptx.ModBulkGroup, ptx.ModL2CacheHint)
//      .Typed(ptx.F32)
//      .InSpace(ptx.Global)
func CpReduceAsyncBulk(dst, src, size Operand, mbar Operand) *Instruction {
    srcs := []Operand{dst, src, size}
    if mbar != nil {
        srcs = append(srcs, mbar)
    }
    return &Instruction{
        Op:    ptx.OpCpReduceAsyncBulk,
        Src:   srcs,
        Space: ptx.SharedCTA, // Default .src is .shared::cta, .dst set via InSpace or modifiers
    }
}

// ---- Async Bulk Prefetch (Section 9.7.9.25.4.3) ----

// CpAsyncBulkPrefetch initiates asynchronous prefetch to L2.
// Usage: CpAsyncBulkPrefetch(src, size, policy).WithMod(ptx.ModL2CacheHint)
func CpAsyncBulkPrefetch(src, size Operand, policy Operand) *Instruction {
    srcs := []Operand{src, size}
    if policy != nil {
        srcs = append(srcs, policy)
    }
    return &Instruction{
        Op:  ptx.OpCpAsyncBulkPrefetch,
        Src: srcs,
    }
}

// ---- Multimem Async Bulk (Section 9.7.9.26) ----

// MultimemCpAsyncBulk initiates async copy to multimem address range.
// Usage: MultimemCpAsyncBulk(dst, src, size, mask).WithMod(ptx.ModBulkGroup)
func MultimemCpAsyncBulk(dst, src, size Operand, byteMask Operand) *Instruction {
    srcs := []Operand{dst, src, size}
    if byteMask != nil {
        srcs = append(srcs, byteMask)
    }
    return &Instruction{
        Op:  ptx.OpMultimemCpAsyncBulk,
        Src: srcs,
    }
}

// ---- Multimem Async Reduce Bulk (Section 9.7.9.27) ----

// MultimemCpReduceAsyncBulk initiates async reduction to multimem.
// Usage:
//   MultimemCpReduceAsyncBulk(dst, src, size)
//     .WithMod(ptx.ModAtomAdd, ptx.ModBulkGroup)
//     .Typed(ptx.U32)
func MultimemCpReduceAsyncBulk(dst, src, size Operand) *Instruction {
    return &Instruction{
        Op:  ptx.OpMultimemCpReduceAsyncBulk,
        Src: []Operand{dst, src, size},
    }
}

// ---- Async Bulk Tensor (Section 9.7.9.27.1.2) ----

// CpAsyncBulkTensor initiates an async copy of tensor data.
//
// Arguments:
//   dim:        Dimensionality (.1d, .2d, ... .5d)
//   dstMem:     Destination memory address
//   tensorMap:  Tensor map object address
//   coords:     Slice of operands for coordinates (tensorCoords)
//   mbar:       Mbarrier object (optional, nil if using .bulk_group)
//   extras:     Optional args (im2colInfo vector, ctaMask, cachePolicy)
//
// Usage:
//   CpAsyncBulkTensor(ptx.ModDim2D, dst, tMap, []Operand{x, y}, mbar, nil)
//     .InSpace(ptx.SharedCluster)
//     .WithMod(ptx.ModMbarrierCompleteTxBytes, ptx.ModLoadTile)
//
// Note: .src is always .global (except for shared->global variant).
// .dst is set via InSpace().
func CpAsyncBulkTensor(
    dim ptx.Modifier,
    dstMem Operand,
    tensorMap Operand,
    coords []Operand,
    mbar Operand,
    extras []Operand,
) *Instruction {
    
    // Construct the instruction.
    // The exact operand order depends on the direction (Global->Shared vs Shared->Global).
    // The common signature is roughly [dst], [map, coords], [mbar], ...
    // or [map, coords], [src]
    
    // We assume the caller structures the "logic" of direction via InSpace/Modifiers,
    // but the operand list in the IR must match the emitting order.
    
    // Since this instruction is highly polymorphic (3 different syntax variants),
    // we collect all provided operands into Src.
    // The emitter will simply print them comma-separated.
    // It is up to the user to provide them in the correct order for the specific variant
    // if using the generic builder, OR we can try to smart-pack them.
    
    // Standard Global -> Shared packing:
    // [dstMem], [tensorMap, tensorCoords], [mbar]{, im2colInfo}{, ctaMask}{, policy}
    
    srcs := []Operand{dstMem}
    
    // Pack tensor map and coords into a vector-like structure if strictly following AST,
    // but PTX syntax treats [tensorMap, coords...] as a bracketed group.
    // Our Emitter likely handles VectorOp as {a, b}.
    // Here we need [a, b, c].
    // Let's create a special VectorOp or just rely on flat operands if the emitter handles it.
    // Given the current definitions, let's pass them as distinct operands and let the user
    // or a future specific emitter handle the brackets.
    // However, usually `[addr]` implies an Address operand.
    // `[tensorMap, tensorCoords]` is a list.
    
    // For simplicity in this IR, we will append them flatly.
    srcs = append(srcs, tensorMap)
    srcs = append(srcs, coords...)
    
    if mbar != nil {
        srcs = append(srcs, mbar)
    }
    
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }

    return &Instruction{
        Op:        ptx.OpCpAsyncBulkTensor,
        Src:       srcs,
        Modifiers: []ptx.Modifier{dim},
    }
}


// ---- Tensor Reduction Bulk (Section 9.7.9.27.1.3) ----

// CpReduceAsyncBulkTensor initiates an async tensor reduction.
//
// Usage:
//   CpReduceAsyncBulkTensor(ptx.ModDim2D, dstMap, []Operand{x, y}, srcMem)
//     .WithMod(ptx.ModAtomAdd, ptx.ModBulkGroup, ptx.ModLoadTile)
func CpReduceAsyncBulkTensor(
	dim ptx.Modifier,
	tensorMap Operand,
	coords []Operand,
	srcMem Operand,
) *Instruction {
	srcs := []Operand{tensorMap}
	srcs = append(srcs, coords...)
	srcs = append(srcs, srcMem)

	return &Instruction{
		Op:        ptx.OpCpReduceAsyncBulkTensor,
		Src:       srcs,
		Space:     ptx.SharedCTA, // Default .src
		Modifiers: []ptx.Modifier{dim},
	}
}

// ---- Tensor Prefetch Bulk (Section 9.7.9.27.1.4) ----

// CpAsyncBulkPrefetchTensor initiates an async tensor prefetch to L2.
//
// Usage:
//   CpAsyncBulkPrefetchTensor(ptx.ModDim5D, tensorMap, []Operand{...}, im2colInfo)
//     .WithMod(ptx.ModLoadIm2Col)
func CpAsyncBulkPrefetchTensor(
    dim ptx.Modifier,
    tensorMap Operand,
    coords []Operand,
    im2colInfo []Operand,
) *Instruction {
    srcs := []Operand{tensorMap}
    srcs = append(srcs, coords...)
    if len(im2colInfo) > 0 {
        srcs = append(srcs, im2colInfo...)
    }

    return &Instruction{
        Op:        ptx.OpCpAsyncBulkPrefetchTensor,
        Src:       srcs,
        // Use ModLevelL2 here, which resolves to ".L2"
        Modifiers: []ptx.Modifier{dim, ptx.ModLevelL2}, 
    }
}

// ---- Bulk Group Management (Section 9.7.9.27.2) ----

// CpAsyncBulkCommitGroup commits uncommitted bulk async ops.
func CpAsyncBulkCommitGroup() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncBulkCommitGroup}
}

// CpAsyncBulkWaitGroup waits for bulk async groups.
// Use WithMod(ptx.ModRead) for .read variant.
func CpAsyncBulkWaitGroup(n int64) *Instruction {
	return &Instruction{
		Op:  ptx.OpCpAsyncBulkWaitGroup,
		Src: []Operand{Imm(n)},
	}
}

// ---- Tensormap Replace (Section 9.7.9.28) ----

// TensormapReplace modifies a field in a tensor map.
// Usage:
//   TensormapReplace(ptx.ModFieldGlobalAddr, addr, newVal).Typed(ptx.B64)
//   TensormapReplace(ptx.ModFieldBoxDim, addr, ord, newVal).Typed(ptx.B32)
func TensormapReplace(field ptx.Modifier, addr Operand, args ...Operand) *Instruction {
	srcs := []Operand{addr}
	srcs = append(srcs, args...) // args contains new_val, or ord, new_val
	return &Instruction{
		Op:        ptx.OpTensormapReplace,
		Src:       srcs,
		Modifiers: []ptx.Modifier{ptx.ModLoadTile, field}, // .tile is standard mode
	}
}

// ---- Texture Instructions (Section 9.7.10) ----

// Tex performs a texture lookup.
//
// Arguments:
//   geom:    Geometry (.1d, .2d, etc.)
//   dst:     Destination vector
//   tex:     Texture handle (or texture array)
//   sampler: Optional sampler (nil if using unified/implicit)
//   coords:  Coordinate vector/operands
//
// Usage:
//   Tex(ptx.ModGeom2D, dst, tex, nil, coords).Typed(ptx.S32).SourceTyped(ptx.F32)
//
// Use WithMod() for .base, .level, .grad.
// Add extra operands (lod, gradient vectors) manually via custom instruction construction
// if strict builder signature is insufficient, or extend this builder.
func Tex(
	geom ptx.Modifier,
	dst Operand,
	tex Operand,
	sampler Operand,
	coords []Operand,
) *Instruction {
	srcs := []Operand{tex}
	if sampler != nil {
		srcs = append(srcs, sampler)
	}
	// Coords are usually treated as a single vector operand or bracketed list in PTX.
	// We append them to Src. The emitter should handle formatting `{a,b}` or `[a,b]`.
	srcs = append(srcs, coords...)

	return &Instruction{
		Op:        ptx.OpTex,
		Dst:       dst,
		Src:       srcs,
		Modifiers: []ptx.Modifier{geom},
	}
}

// Tld4 performs a texture fetch of the 4-texel bilerp footprint.
// Usage: Tld4(ptx.ModCompR, ptx.ModGeom2D, dst, tex, nil, coords)
func Tld4(
	comp ptx.Modifier,
	geom ptx.Modifier,
	dst Operand,
	tex Operand,
	sampler Operand,
	coords []Operand,
) *Instruction {
	srcs := []Operand{tex}
	if sampler != nil {
		srcs = append(srcs, sampler)
	}
	srcs = append(srcs, coords...)

	return &Instruction{
		Op:        ptx.OpTld4,
		Dst:       dst,
		Src:       srcs,
		Modifiers: []ptx.Modifier{comp, geom},
	}
}


//   Txq(ptx.ModQueryWidth, r1, tex, nil)
func Txq(query ptx.Modifier, dst Operand, tex Operand, lod Operand) *Instruction {
	srcs := []Operand{tex}
	mods := []ptx.Modifier{query}
	
	if lod != nil {
		srcs = append(srcs, lod)
		mods = append([]ptx.Modifier{ptx.ModLevel}, mods...) // Prepend .level
	}

	return &Instruction{
		Op:        ptx.OpTxq,
		Dst:       dst,
		Src:       srcs,
		Modifiers: mods,
	}
}

// Istypep checks opaque variable type.
// Usage: Istypep(ptx.ModTypeTexRef, p, tptr)
func Istypep(typ ptx.Modifier, dstPred Operand, addr Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpIstypep,
		Dst:       dstPred,
		Src:       []Operand{addr},
		Modifiers: []ptx.Modifier{typ},
	}
}

// ---- Surface Instructions (Section 9.7.11) ----

// Suld loads from surface memory.
//
// Arguments:
//   geom:   Geometry (e.g., ptx.ModGeom2D)
//   dst:    Destination vector
//   surf:   Surface reference
//   coords: Coordinate operands
//
// Usage:
//   Suld(ptx.ModGeom2D, dst, surf, []Operand{x, y})
//     .WithMod(ptx.ModClampTrap)
//     .WithCache(ptx.CacheCA)
func Suld(geom ptx.Modifier, dst Operand, surf Operand, coords []Operand) *Instruction {
	srcs := []Operand{surf}
	srcs = append(srcs, coords...)

	return &Instruction{
		Op:        ptx.OpSuld,
		Dst:       dst,
		Src:       srcs,
		Modifiers: []ptx.Modifier{ptx.ModB, geom}, // .b (unformatted) is default logic
	}
}

// Sust stores to surface memory.
// Usage: Sust(ptx.ModGeom2D, surf, []Operand{x, y}, val)
func Sust(geom ptx.Modifier, surf Operand, coords []Operand, val Operand) *Instruction {
	srcs := []Operand{surf}
	srcs = append(srcs, coords...)
	srcs = append(srcs, val)

	return &Instruction{
		Op:        ptx.OpSust,
		Src:       srcs,
		Modifiers: []ptx.Modifier{ptx.ModB, geom},
	}
}

// Sured reduces surface memory.
// Usage: Sured(ptx.ModGeom2D, surf, []Operand{x, y}, val).WithMod(ptx.ModAtomAdd)
func Sured(geom ptx.Modifier, surf Operand, coords []Operand, val Operand) *Instruction {
	srcs := []Operand{surf}
	srcs = append(srcs, coords...)
	srcs = append(srcs, val)

	return &Instruction{
		Op:        ptx.OpSured,
		Src:       srcs,
		Modifiers: []ptx.Modifier{ptx.ModB, geom},
	}
}

// Suq queries surface attributes.
// Usage: Suq(ptx.ModQueryWidth, dst, surf)
func Suq(query ptx.Modifier, dst Operand, surf Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpSuq,
		Dst:       dst,
		Src:       []Operand{surf},
		Modifiers: []ptx.Modifier{query},
	}
}

// ---- Advanced Control Flow (Section 9.7.12) ----

// BrxIdx branches to a label indexed from a list.
// targetList must be a label pointing to a .branchtargets directive.
// Usage: BrxIdx(idx, Sym("L_targets")).WithMod(ptx.ModUni)
func BrxIdx(index Operand, targetList Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpBrxIdx,
		Src: []Operand{index, targetList},
	}
}

// CallIndirect performs an indirect function call.
//
// Arguments:
//   funcPtr:   Register holding function address
//   retParams: Optional return operands
//   args:      Call arguments
//   proto:     Prototype label (.callprototype) OR Target list label (.calltargets)
//
// Usage:
//   CallIndirect(fptr, ret, args, Sym("Fproto"))
func CallIndirect(funcPtr Operand, retParams []Operand, args []Operand, proto Operand) *Instruction {
	srcs := []Operand{funcPtr}
	srcs = append(srcs, args...)
	srcs = append(srcs, proto)

	inst := &Instruction{
		Op:  ptx.OpCall,
		Src: srcs,
	}
	
	if len(retParams) > 0 {
		inst.Dst = retParams[0]
	}
	return inst
}


// ---- Thread Termination ----

func Exit() *Instruction {
	return &Instruction{Op: ptx.OpExit}
}

// ---- Parallel Synchronization (Section 9.7.13) ----

// BarrierCTA performs CTA-level barrier synchronization.
// Usage: BarrierCTA(Imm(0)).WithMod(ptx.ModSync, ptx.ModAligned)
func BarrierCTA(id Operand, threadCount ...Operand) *Instruction {
	srcs := []Operand{id}
	srcs = append(srcs, threadCount...)
	return &Instruction{
		Op:    ptx.OpBar, // bar and barrier are often synonymous in IR
		Src:   srcs,
		Space: ptx.SharedCTA,
	}
}

// BarWarpSync synchronizes threads within a warp.
// membermask is a 32-bit bitmask of lanes.
func BarWarpSync(membermask Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpBarWarpSync,
		Src: []Operand{membermask},
	}
}

// BarrierClusterArrive marks arrival at a cluster barrier.
// Use WithMod(ptx.ModRelaxed, ptx.ModAligned) for variants.
func BarrierClusterArrive() *Instruction {
	return &Instruction{
		Op:        ptx.OpBarrierCluster,
		Modifiers: []ptx.Modifier{ptx.ModArrive},
	}
}

// BarrierClusterWait waits for cluster barrier completion.
func BarrierClusterWait() *Instruction {
	return &Instruction{
		Op:        ptx.OpBarrierCluster,
		Modifiers: []ptx.Modifier{ptx.ModWait},
	}
}

// ---- Memory Ordering (Section 9.7.13.4) ----


// FenceProxyAsync creates an asynchronous proxy fence.
func FenceProxyAsync(space ptx.StateSpace) *Instruction {
	return &Instruction{
		Op:        ptx.OpFence,
		Space:     space,
		Modifiers: []ptx.Modifier{ptx.ModProxy, ptx.ModAsync},
	}
}

// ---- Atomic Operations (Section 9.7.13.5) ----

// AtomCAS performs a Compare-And-Swap.

// AtomExch performs an Exchange.
func AtomExch(typ ptx.Type, dst, addr, val Operand) *Instruction {
	return &Instruction{
		Op:    ptx.OpAtom,
		Typ:   typ,
		Dst:   dst,
		Src:   []Operand{addr, val},
		Modifiers: []ptx.Modifier{ptx.ModExch},
	}
}

// AtomVector performs a vector atomic operation (sm_90+).
// Usage: AtomVector(ptx.OpAdd, ptx.V4, ptx.F32, dstVec, addr, srcVec)
func AtomVector(op ptx.Opcode, vec ptx.VectorSize, typ ptx.Type, dst, addr, src Operand) *Instruction {
	return &Instruction{
		Op:    op,
		Vec:   vec,
		Typ:   typ,
		Dst:   dst,
		Src:   []Operand{addr, src},
		Space: ptx.Global,
	}
}