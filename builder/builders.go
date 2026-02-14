package builder

import (
	"github.com/arc-language/ptx-gen/ptx"
)

// Integer & float arithmetic

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

func Min3(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMin, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Max3(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMax, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Mul24(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMul24, Dst: dst, Src: []Operand{src0, src1}}
}

func Mad24(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMad24, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Sad(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSad, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Dp4a(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpDp4a, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Dp2a(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpDp2a, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// Extended-precision integer arithmetic

func AddCC(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAddCC, Dst: dst, Src: []Operand{src0, src1}}
}

func Addc(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpAddc, Dst: dst, Src: []Operand{src0, src1}}
}

func SubCC(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSubCC, Dst: dst, Src: []Operand{src0, src1}}
}

func Subc(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpSubc, Dst: dst, Src: []Operand{src0, src1}}
}

func MadCC(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMadCC, Dst: dst, Src: []Operand{src0, src1, src2}}
}

func Madc(dst, src0, src1, src2 Operand) *Instruction {
	return &Instruction{Op: ptx.OpMadc, Dst: dst, Src: []Operand{src0, src1, src2}}
}

// Float-only math

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

func Tanh(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpTanh, Dst: dst, Src: []Operand{src}}
}

func Testp(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpTestp, Dst: dst, Src: []Operand{src}}
}

func Copysign(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpCopysign, Dst: dst, Src: []Operand{src0, src1}}
}

// Bit manipulation

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

func Bfi(dst, base, insert, start, len Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfi, Dst: dst, Src: []Operand{base, insert, start, len}}
}

func Bfind(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpBfind, Dst: dst, Src: []Operand{src}}
}

func Fns(dst, mask, base, offset Operand) *Instruction {
	return &Instruction{Op: ptx.OpFns, Dst: dst, Src: []Operand{mask, base, offset}}
}

func Bmsk(dst, pos, width Operand) *Instruction {
	return &Instruction{Op: ptx.OpBmsk, Dst: dst, Src: []Operand{pos, width}}
}

func Szext(dst, src, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpSzext, Dst: dst, Src: []Operand{src, size}}
}

// Logic & shift

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

func Cnot(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCnot, Dst: dst, Src: []Operand{src}}
}

func Lop3(dst, a, b, c, immLut Operand) *Instruction {
	return &Instruction{Op: ptx.OpLop3, Dst: dst, Src: []Operand{a, b, c, immLut}}
}

func Shl(dst, src, amount Operand) *Instruction {
	return &Instruction{Op: ptx.OpShl, Dst: dst, Src: []Operand{src, amount}}
}

func Shr(dst, src, amt Operand) *Instruction {
	return &Instruction{Op: ptx.OpShr, Dst: dst, Src: []Operand{src, amt}}
}

func ShfL(dst, a, b, c Operand) *Instruction {
	return &Instruction{Op: ptx.OpShf, Dst: dst, Src: []Operand{a, b, c}, Modifiers: []ptx.Modifier{ptx.ModLeft}}
}

func ShfR(dst, a, b, c Operand) *Instruction {
	return &Instruction{Op: ptx.OpShf, Dst: dst, Src: []Operand{a, b, c}, Modifiers: []ptx.Modifier{ptx.ModRight}}
}

func Prmt(dst, a, b, c Operand) *Instruction {
	return &Instruction{Op: ptx.OpPrmt, Dst: dst, Src: []Operand{a, b, c}}
}

// Comparison & selection

func Setp(cmp ptx.CmpOp, dst, a, b Operand) *Instruction {
	return &Instruction{Op: ptx.OpSetp, Cmp: cmp, Dst: dst, Src: []Operand{a, b}}
}

func Set(cmp ptx.CmpOp, dst, a, b Operand) *Instruction {
	return &Instruction{Op: ptx.OpSet, Cmp: cmp, Dst: dst, Src: []Operand{a, b}}
}

func Selp(dst, a, b, p Operand) *Instruction {
	return &Instruction{Op: ptx.OpSelp, Dst: dst, Src: []Operand{a, b, p}}
}

func Slct(dst, a, b, c Operand) *Instruction {
	return &Instruction{Op: ptx.OpSlct, Dst: dst, Src: []Operand{a, b, c}}
}

// Data movement & conversion

func Mov(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpMov, Dst: dst, Src: []Operand{src}}
}

func Ld(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLd, Dst: dst, Src: []Operand{addr}}
}

func LdNC(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdNC, Dst: dst, Src: []Operand{addr}, Space: ptx.Global}
}

func LdGlobalNC(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdNC, Dst: dst, Src: []Operand{addr}, Space: ptx.Global}
}

func LdParam(dst Operand, param *Symbol) *Instruction {
	return &Instruction{Op: ptx.OpLd, Space: ptx.Param, Dst: dst, Src: []Operand{&Address{Base: param, Offset: 0}}}
}

func Ldu(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdu, Dst: dst, Src: []Operand{addr}}
}

func LdWeak(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLd, Dst: dst, Src: []Operand{addr}, Modifiers: []ptx.Modifier{ptx.ModWeak}}
}

func St(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpSt, Src: []Operand{addr, src}}
}

func StWeak(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpSt, Src: []Operand{addr, src}, Modifiers: []ptx.Modifier{ptx.ModWeak}}
}

func StAsync(addr, src Operand, mbar Operand) *Instruction {
	if mbar != nil {
		return &Instruction{Op: ptx.OpStAsync, Src: []Operand{addr, src, mbar}}
	}
	return &Instruction{Op: ptx.OpStAsync, Src: []Operand{addr, src}}
}

func StBulk(addr, size, initVal Operand) *Instruction {
	return &Instruction{Op: ptx.OpStBulk, Src: []Operand{addr, size, initVal}}
}

func Cvt(dst Operand, srcs ...Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvt, Dst: dst, Src: srcs}
}

func CvtPack(dst Operand, srcs ...Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvtPack, Dst: dst, Src: srcs}
}

func Cvta(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpCvta, Dst: dst, Src: []Operand{src}}
}

func Prefetch(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpPrefetch, Src: []Operand{addr}}
}

func Prefetchu(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpPrefetchu, Src: []Operand{addr}}
}

func IsSpacep(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpIsSpacep, Dst: dst, Src: []Operand{addr}}
}

func Mapa(dst, addr, ctaRank Operand) *Instruction {
	return &Instruction{Op: ptx.OpMapa, Dst: dst, Src: []Operand{addr, ctaRank}}
}

func GetCTARank(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpGetCTARank, Dst: dst, Src: []Operand{addr}}
}

// Warp shuffle

func ShflSync(dst, a, b, c, mask Operand) *Instruction {
	return &Instruction{Op: ptx.OpShfl, Dst: dst, Src: []Operand{a, b, c, mask}, Modifiers: []ptx.Modifier{ptx.ModSync}}
}

func Shfl(dst, a, b, c Operand) *Instruction {
	return &Instruction{Op: ptx.OpShfl, Dst: dst, Src: []Operand{a, b, c}}
}

// Stack manipulation

func Alloca(dst, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpAlloca, Dst: dst, Src: []Operand{size}}
}

func StackSave(dst Operand) *Instruction {
	return &Instruction{Op: ptx.OpStackSave, Dst: dst}
}

func StackRestore(src Operand) *Instruction {
	return &Instruction{Op: ptx.OpStackRestore, Src: []Operand{src}}
}

// Control flow

func Bra(label string) *Instruction {
	return &Instruction{Op: ptx.OpBra, Src: []Operand{&Symbol{Name: label}}}
}

func BraUni(label string) *Instruction {
	return &Instruction{Op: ptx.OpBra, Src: []Operand{&Symbol{Name: label}}, Modifiers: []ptx.Modifier{ptx.ModUni}}
}

func BrxIdx(index Operand, targetList Operand) *Instruction {
	return &Instruction{Op: ptx.OpBrxIdx, Src: []Operand{index, targetList}}
}

func Call(target string, retParams []Operand, args []Operand) *Instruction {
	inst := &Instruction{Op: ptx.OpCall, Src: args, CallTarget: target}
	if len(retParams) > 0 {
		inst.Dst = retParams[0]
	}
	return inst
}

func CallIndirect(funcPtr Operand, retParams []Operand, args []Operand, proto Operand) *Instruction {
	srcs := []Operand{funcPtr}
	srcs = append(srcs, args...)
	srcs = append(srcs, proto)
	inst := &Instruction{Op: ptx.OpCall, Src: srcs}
	if len(retParams) > 0 {
		inst.Dst = retParams[0]
	}
	return inst
}

func Ret() *Instruction {
	return &Instruction{Op: ptx.OpRet}
}

func Exit() *Instruction {
	return &Instruction{Op: ptx.OpExit}
}

// Synchronization

func BarSync(id Operand) *Instruction {
	return &Instruction{Op: ptx.OpBar, Src: []Operand{id}}
}

func BarSyncCount(id, count Operand) *Instruction {
	return &Instruction{Op: ptx.OpBar, Src: []Operand{id, count}}
}

func BarrierCTA(id Operand, threadCount ...Operand) *Instruction {
	srcs := []Operand{id}
	srcs = append(srcs, threadCount...)
	return &Instruction{Op: ptx.OpBar, Src: srcs, Space: ptx.SharedCTA}
}

func BarWarpSync(membermask Operand) *Instruction {
	return &Instruction{Op: ptx.OpBarWarpSync, Src: []Operand{membermask}}
}

func BarrierClusterArrive() *Instruction {
	return &Instruction{Op: ptx.OpBarrierCluster, Modifiers: []ptx.Modifier{ptx.ModArrive}}
}

func BarrierClusterWait() *Instruction {
	return &Instruction{Op: ptx.OpBarrierCluster, Modifiers: []ptx.Modifier{ptx.ModWait}}
}

func Membar(level ptx.Modifier) *Instruction {
	return &Instruction{Op: ptx.OpMembar, Modifiers: []ptx.Modifier{level}}
}

func MembarProxy() *Instruction {
	return &Instruction{Op: ptx.OpMembar, Modifiers: []ptx.Modifier{ptx.ModProxy}}
}

func Fence(scope ptx.Scope) *Instruction {
	return &Instruction{Op: ptx.OpFence, Scope: scope}
}

func FenceSC(scope ptx.Scope) *Instruction {
	return &Instruction{Op: ptx.OpFence, Scope: scope, Modifiers: []ptx.Modifier{ptx.ModSC}}
}

func FenceAcqRel(scope ptx.Scope) *Instruction {
	return &Instruction{Op: ptx.OpFence, Scope: scope, Modifiers: []ptx.Modifier{ptx.ModAcqRel}}
}

func FenceProxy(kind ptx.Modifier) *Instruction {
	return &Instruction{Op: ptx.OpFence, Modifiers: []ptx.Modifier{ptx.ModProxy, kind}}
}

// Atomics

func Atom(op ptx.Modifier, dst, addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpAtom, Dst: dst, Src: []Operand{addr, src}, Modifiers: []ptx.Modifier{op}}
}

func AtomCAS(dst, addr, compare, val Operand) *Instruction {
	return &Instruction{Op: ptx.OpAtom, Dst: dst, Src: []Operand{addr, compare, val}, Modifiers: []ptx.Modifier{ptx.ModAtomCAS}}
}

func AtomExch(typ ptx.Type, dst, addr, val Operand) *Instruction {
	return &Instruction{Op: ptx.OpAtom, Typ: typ, Dst: dst, Src: []Operand{addr, val}, Modifiers: []ptx.Modifier{ptx.ModExch}}
}

func AtomVector(op ptx.Opcode, vec ptx.VectorSize, typ ptx.Type, dst, addr, src Operand) *Instruction {
	return &Instruction{Op: op, Vec: vec, Typ: typ, Dst: dst, Src: []Operand{addr, src}, Space: ptx.Global}
}

func Red(op ptx.Modifier, addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpRed, Src: []Operand{addr, src}, Modifiers: []ptx.Modifier{op}}
}

// Warp voting

func VoteSync(mode ptx.Modifier, dst, mask, pred Operand) *Instruction {
	return &Instruction{Op: ptx.OpVoteSync, Dst: dst, Src: []Operand{mask, pred}, Modifiers: []ptx.Modifier{mode}}
}

func Activemask(dst Operand) *Instruction {
	return &Instruction{Op: ptx.OpActivemask, Dst: dst}
}

func ReduxSync(op ptx.Modifier, dst, mask, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpReduxSync, Dst: dst, Src: []Operand{mask, src}, Modifiers: []ptx.Modifier{op}}
}

// Matrix

func LdMatrix(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdMatrix, Dst: dst, Src: []Operand{addr}}
}

func StMatrix(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpStMatrix, Src: []Operand{addr, src}}
}

func Wgmma(dst Operand, args ...Operand) *Instruction {
	return &Instruction{Op: ptx.OpWgmma, Dst: dst, Src: args}
}

// Mbarrier

func MbarrierInit(addr, count Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierInit, Src: []Operand{addr, count}}
}

func MbarrierArrive(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierArrive, Src: []Operand{addr}}
}

func MbarrierInval(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierInval, Src: []Operand{addr}}
}

func MbarrierExpectTx(addr, txCount Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierExpectTx, Src: []Operand{addr, txCount}}
}

func MbarrierCompleteTx(addr, txCount Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierCompleteTx, Src: []Operand{addr, txCount}}
}

func MbarrierArriveNoComplete(dstState, addr, count Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpMbarrierArrive,
		Dst:       dstState,
		Src:       []Operand{addr, count},
		Modifiers: []ptx.Modifier{ptx.ModNoComplete},
	}
}

func MbarrierArriveExpectTx(dstState, addr, txCount Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpMbarrierArrive,
		Dst:       dstState,
		Src:       []Operand{addr, txCount},
		Modifiers: []ptx.Modifier{ptx.ModExpectTx},
	}
}

func MbarrierArriveDrop(dstState, addr, count Operand) *Instruction {
	srcs := []Operand{addr}
	if count != nil {
		srcs = append(srcs, count)
	}
	return &Instruction{Op: ptx.OpMbarrierArriveDrop, Dst: dstState, Src: srcs}
}

func MbarrierArriveDropNoComplete(dstState, addr, count Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpMbarrierArriveDrop,
		Dst:       dstState,
		Src:       []Operand{addr, count},
		Modifiers: []ptx.Modifier{ptx.ModNoComplete},
	}
}

func MbarrierTestWait(waitComplete, addr, state Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMbarrierTestWait,
		Dst: waitComplete,
		Src: []Operand{addr, state},
	}
}

func MbarrierTestWaitParity(waitComplete, addr, phaseParity Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpMbarrierTestWait,
		Dst:       waitComplete,
		Src:       []Operand{addr, phaseParity},
		Modifiers: []ptx.Modifier{ptx.ModParity},
	}
}

func MbarrierTryWait(waitComplete, addr, state Operand, suspendHint Operand) *Instruction {
	srcs := []Operand{addr, state}
	if suspendHint != nil {
		srcs = append(srcs, suspendHint)
	}
	return &Instruction{
		Op:  ptx.OpMbarrierTryWait,
		Dst: waitComplete,
		Src: srcs,
	}
}

func MbarrierPendingCount(countDst, state Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMbarrierPendingCount,
		Dst: countDst,
		Src: []Operand{state},
		Typ: ptx.B64,
	}
}

func CpAsyncMbarrierArrive(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncMbarrierArrive, Src: []Operand{addr}}
}

func CpAsyncMbarrierArriveNoInc(addr Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpCpAsyncMbarrierArrive,
		Src:       []Operand{addr},
		Modifiers: []ptx.Modifier{ptx.ModNoInc},
	}
}

// Multimem

func MultimemLdReduce(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpMultimemLdReduce, Dst: dst, Src: []Operand{addr}}
}

func MultimemSt(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpMultimemSt, Src: []Operand{addr, src}}
}

func MultimemRed(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpMultimemRed, Src: []Operand{addr, src}}
}

func MultimemCpAsyncBulk(dst, src, size Operand, byteMask Operand) *Instruction {
	srcs := []Operand{dst, src, size}
	if byteMask != nil {
		srcs = append(srcs, byteMask)
	}
	return &Instruction{Op: ptx.OpMultimemCpAsyncBulk, Src: srcs}
}

func MultimemCpReduceAsyncBulk(dst, src, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpMultimemCpReduceAsyncBulk, Src: []Operand{dst, src, size}}
}

// Cache & eviction policy

func ApplyPriority(addr, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpApplyPriority, Src: []Operand{addr, size}}
}

func Discard(addr, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpDiscard, Src: []Operand{addr, size}}
}

func CreatePolicy(dst Operand, args ...Operand) *Instruction {
	return &Instruction{Op: ptx.OpCreatePolicy, Dst: dst, Src: args}
}

// Async copy

func CpAsync(dst, src, size Operand, args ...Operand) *Instruction {
	srcs := []Operand{dst, src, size}
	srcs = append(srcs, args...)
	return &Instruction{Op: ptx.OpCpAsync, Src: srcs}
}

func CpAsyncCommitGroup() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncCommitGroup}
}

func CpAsyncWaitAll() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncWaitAll}
}

func CpAsyncWaitGroup(n int64) *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncWaitGroup, Src: []Operand{Imm(n)}}
}

func CpAsyncBulk(dst, src, size Operand, args ...Operand) *Instruction {
	srcs := []Operand{dst, src, size}
	srcs = append(srcs, args...)
	return &Instruction{Op: ptx.OpCpAsyncBulk, Src: srcs}
}

func CpAsyncBulkCommitGroup() *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncBulkCommitGroup}
}

func CpAsyncBulkWaitGroup(n int64) *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncBulkWaitGroup, Src: []Operand{Imm(n)}}
}

func CpAsyncBulkPrefetch(src, size Operand, policy Operand) *Instruction {
	srcs := []Operand{src, size}
	if policy != nil {
		srcs = append(srcs, policy)
	}
	return &Instruction{Op: ptx.OpCpAsyncBulkPrefetch, Src: srcs}
}

func CpReduceAsyncBulk(dst, src, size Operand, mbar Operand) *Instruction {
	srcs := []Operand{dst, src, size}
	if mbar != nil {
		srcs = append(srcs, mbar)
	}
	return &Instruction{Op: ptx.OpCpReduceAsyncBulk, Src: srcs, Space: ptx.SharedCTA}
}

// Async bulk tensor

func CpAsyncBulkTensor(
	dim ptx.Modifier,
	dstMem Operand,
	tensorMap Operand,
	coords []Operand,
	mbar Operand,
	extras []Operand,
) *Instruction {
	srcs := []Operand{dstMem, tensorMap}
	srcs = append(srcs, coords...)
	if mbar != nil {
		srcs = append(srcs, mbar)
	}
	if len(extras) > 0 {
		srcs = append(srcs, extras...)
	}
	return &Instruction{Op: ptx.OpCpAsyncBulkTensor, Src: srcs, Modifiers: []ptx.Modifier{dim}}
}

func CpReduceAsyncBulkTensor(
	dim ptx.Modifier,
	tensorMap Operand,
	coords []Operand,
	srcMem Operand,
) *Instruction {
	srcs := []Operand{tensorMap}
	srcs = append(srcs, coords...)
	srcs = append(srcs, srcMem)
	return &Instruction{Op: ptx.OpCpReduceAsyncBulkTensor, Src: srcs, Space: ptx.SharedCTA, Modifiers: []ptx.Modifier{dim}}
}

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
	return &Instruction{Op: ptx.OpCpAsyncBulkPrefetchTensor, Src: srcs, Modifiers: []ptx.Modifier{dim, ptx.ModLevelL2}}
}

// Tensormap

func TensormapReplace(field ptx.Modifier, addr Operand, args ...Operand) *Instruction {
	srcs := []Operand{addr}
	srcs = append(srcs, args...)
	return &Instruction{Op: ptx.OpTensormapReplace, Src: srcs, Modifiers: []ptx.Modifier{ptx.ModLoadTile, field}}
}

// Texture

func Tex(geom ptx.Modifier, dst Operand, tex Operand, sampler Operand, coords []Operand) *Instruction {
	srcs := []Operand{tex}
	if sampler != nil {
		srcs = append(srcs, sampler)
	}
	srcs = append(srcs, coords...)
	return &Instruction{Op: ptx.OpTex, Dst: dst, Src: srcs, Modifiers: []ptx.Modifier{geom}}
}

func Tld4(comp ptx.Modifier, geom ptx.Modifier, dst Operand, tex Operand, sampler Operand, coords []Operand) *Instruction {
	srcs := []Operand{tex}
	if sampler != nil {
		srcs = append(srcs, sampler)
	}
	srcs = append(srcs, coords...)
	return &Instruction{Op: ptx.OpTld4, Dst: dst, Src: srcs, Modifiers: []ptx.Modifier{comp, geom}}
}

func Txq(query ptx.Modifier, dst Operand, tex Operand, lod Operand) *Instruction {
	srcs := []Operand{tex}
	mods := []ptx.Modifier{query}
	if lod != nil {
		srcs = append(srcs, lod)
		mods = append([]ptx.Modifier{ptx.ModLevel}, mods...)
	}
	return &Instruction{Op: ptx.OpTxq, Dst: dst, Src: srcs, Modifiers: mods}
}

func Istypep(typ ptx.Modifier, dstPred Operand, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpIstypep, Dst: dstPred, Src: []Operand{addr}, Modifiers: []ptx.Modifier{typ}}
}

// Surface

func Suld(geom ptx.Modifier, dst Operand, surf Operand, coords []Operand) *Instruction {
	srcs := []Operand{surf}
	srcs = append(srcs, coords...)
	return &Instruction{Op: ptx.OpSuld, Dst: dst, Src: srcs, Modifiers: []ptx.Modifier{ptx.ModB, geom}}
}

func Sust(geom ptx.Modifier, surf Operand, coords []Operand, val Operand) *Instruction {
	srcs := []Operand{surf}
	srcs = append(srcs, coords...)
	srcs = append(srcs, val)
	return &Instruction{Op: ptx.OpSust, Src: srcs, Modifiers: []ptx.Modifier{ptx.ModB, geom}}
}

func Sured(geom ptx.Modifier, surf Operand, coords []Operand, val Operand) *Instruction {
	srcs := []Operand{surf}
	srcs = append(srcs, coords...)
	srcs = append(srcs, val)
	return &Instruction{Op: ptx.OpSured, Src: srcs, Modifiers: []ptx.Modifier{ptx.ModB, geom}}
}

func Suq(query ptx.Modifier, dst Operand, surf Operand) *Instruction {
	return &Instruction{Op: ptx.OpSuq, Dst: dst, Src: []Operand{surf}, Modifiers: []ptx.Modifier{query}}
}

// Tensormap operations

// TensormapCpFenceproxy performs a fused copy and fence (sm_90+).
// Implies .global.shared::cta, .tensormap::generic, .sync, .aligned.
func TensormapCpFenceproxy(dstGlobal, srcShared, size Operand, scope ptx.Scope) *Instruction {
	return &Instruction{
		Op:    ptx.OpTensormapCpFenceproxy,
		Src:   []Operand{dstGlobal, srcShared, size},
		Scope: scope,
		Modifiers: []ptx.Modifier{
			ptx.ModSpaceGlobal,
			ptx.ModSpaceSharedCTA,
			ptx.ModTensormapGeneric,
			ptx.ModRelease,
			ptx.ModSync,
			ptx.ModAligned,
		},
	}
}

// Cluster launch control

// ClusterlaunchcontrolTryCancel requests cancellation of a cluster (sm_100+).
func ClusterlaunchcontrolTryCancel(addr, mbar Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpClusterlaunchcontrolTryCancel,
		Src: []Operand{addr, mbar},
		Modifiers: []ptx.Modifier{
			ptx.ModAsync,
			ptx.ModMbarrierCompleteTxBytes,
			ptx.ModMulticastClusterAll,
		},
		Typ: ptx.B128,
	}
}

// ClusterlaunchcontrolQueryCancelIsCanceled checks if cancellation succeeded.
func ClusterlaunchcontrolQueryCancelIsCanceled(predDst, handle Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpClusterlaunchcontrolQueryCancel,
		Dst:       predDst,
		Src:       []Operand{handle},
		Modifiers: []ptx.Modifier{ptx.ModIsCanceled},
		Typ:       ptx.B128,
	}
}

// ClusterlaunchcontrolQueryCancelGetFirstCTAId extracts the first CTA ID.
// Returns a vector of 4x b32.
func ClusterlaunchcontrolQueryCancelGetFirstCTAId(dstVec, handle Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpClusterlaunchcontrolQueryCancel,
		Dst:       dstVec,
		Src:       []Operand{handle},
		Modifiers: []ptx.Modifier{ptx.ModGetFirstCTAId},
		Typ:       ptx.B128,
	}
}

// ClusterlaunchcontrolQueryCancelGetFirstCTAIdDim extracts a specific dimension (x, y, or z).
// dim should be ptx.ModDimX, ptx.ModDimY, or ptx.ModDimZ.
func ClusterlaunchcontrolQueryCancelGetFirstCTAIdDim(dstReg, handle Operand, dim ptx.Modifier) *Instruction {
	return &Instruction{
		Op:        ptx.OpClusterlaunchcontrolQueryCancel,
		Dst:       dstReg,
		Src:       []Operand{handle},
		Modifiers: []ptx.Modifier{ptx.ModGetFirstCTAId, dim},
		Typ:       ptx.B128,
	}
}

// WMMA

// WmmaLoad creates a wmma.load instruction.
// role: ModMatrixA, ModMatrixB, or ModMatrixC.
// layout: ModRow or ModCol.
// stride: optional (pass nil for default).
func WmmaLoad(role, layout, shape ptx.Modifier, typ ptx.Type, dst, addr, stride Operand) *Instruction {
	srcs := []Operand{addr}
	if stride != nil {
		srcs = append(srcs, stride)
	}
	return &Instruction{
		Op:  ptx.OpWmmaLoad,
		Typ: typ,
		Dst: dst,
		Src: srcs,
		Modifiers: []ptx.Modifier{
			role,
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			layout,
		},
	}
}

// WmmaStore creates a wmma.store instruction.
// role: ModMatrixD.
// layout: ModRow or ModCol.
// stride: optional (pass nil for default).
func WmmaStore(role, layout, shape ptx.Modifier, typ ptx.Type, addr, src, stride Operand) *Instruction {
	srcs := []Operand{addr, src}
	if stride != nil {
		srcs = append(srcs, stride)
	}
	return &Instruction{
		Op:  ptx.OpWmmaStore,
		Typ: typ,
		Src: srcs,
		Modifiers: []ptx.Modifier{
			role,
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			layout,
		},
	}
}

// WmmaMma creates a wmma.mma instruction performing D = A * B + C.
func WmmaMma(shape, alayout, blayout ptx.Modifier, d, a, b, c Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpWmmaMma,
		Dst: d,
		Src: []Operand{a, b, c},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			alayout,
			blayout,
			shape,
		},
	}
}

// WmmaMmaBitOp creates a wmma.mma.op.popc instruction for .b1 types.
// op: ptx.ModAtomXor or ptx.ModAtomAnd.
func WmmaMmaBitOp(op, shape, alayout, blayout ptx.Modifier, d, a, b, c Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpWmmaMma,
		Dst: d,
		Src: []Operand{a, b, c},
		Modifiers: []ptx.Modifier{
			op,
			ptx.ModPopc,
			ptx.ModSync,
			ptx.ModAligned,
			alayout,
			blayout,
			shape,
		},
	}
}

// MMA

// Mma creates an mma.sync.aligned instruction.
func Mma(shape, alayout, blayout ptx.Modifier, d, a, b, c Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMma,
		Dst: d,
		Src: []Operand{a, b, c},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			alayout,
			blayout,
			shape,
		},
	}
}

// Ldmatrix loads matrices from shared memory.
// shape: .m8n8, .m16n16, .m8n16
// num: .x1, .x2, .x4
func Ldmatrix(shape, num ptx.Modifier, typ ptx.Type, dst, addr Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpLdMatrix,
		Typ: typ,
		Dst: dst,
		Src: []Operand{addr},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			num,
		},
	}
}

// LdmatrixTrans loads matrices with transpose (.trans).
func LdmatrixTrans(shape, num ptx.Modifier, typ ptx.Type, dst, addr Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpLdMatrix,
		Typ: typ,
		Dst: dst,
		Src: []Operand{addr},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			num,
			ptx.ModTrans,
		},
	}
}

// Stmatrix stores matrices to shared memory.
func Stmatrix(shape, num ptx.Modifier, typ ptx.Type, addr, src Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpStMatrix,
		Typ: typ,
		Src: []Operand{addr, src},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			num,
		},
	}
}

// Movmatrix transposes a matrix in registers.
func Movmatrix(shape ptx.Modifier, typ ptx.Type, dst, src Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMovMatrix,
		Typ: typ,
		Dst: dst,
		Src: []Operand{src},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			ptx.ModTrans,
		},
	}
}

// MmaBlockScaled creates an mma instruction with block scaling: D = (A * ScaleA) * (B * ScaleB) + C.
// scaleVecSize: e.g. ptx.ModScaleVec1x, ptx.ModScaleVec2x.
func MmaBlockScaled(shape, alayout, blayout, scaleVecSize ptx.Modifier, d, a, b, c, scaleA, selA, scaleB, selB Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMma,
		Dst: d,
		Src: []Operand{a, b, c, scaleA, selA, scaleB, selB},
		Modifiers: []ptx.Modifier{
			ptx.ModSync,
			ptx.ModAligned,
			alayout,
			blayout,
			shape,
			ptx.ModBlockScale,
			scaleVecSize,
		},
	}
}

// MmaSparse creates a sparse mma instruction.
// spMod: ptx.ModSp or ptx.ModSpOrderedMetadata.
// metadata: register containing indices of non-zero elements.
// selector: immediate or register (0 or 1) indicating metadata ownership.
func MmaSparse(spMod, shape, alayout, blayout ptx.Modifier, d, a, b, c, metadata, selector Operand) *Instruction {
	return &Instruction{
		Op:  ptx.OpMma,
		Dst: d,
		Src: []Operand{a, b, c, metadata, selector},
		Modifiers: []ptx.Modifier{
			spMod,
			ptx.ModSync,
			ptx.ModAligned,
			alayout,
			blayout,
			shape,
		},
	}
}

// MmaSparseBlockScaled creates a sparse mma instruction with block scaling.
// selA and selB are typically VectorOp containing immediate values.
func MmaSparseBlockScaled(
	spMod, shape, alayout, blayout ptx.Modifier,
	kind, scaleVecSize ptx.Modifier,
	d, a, b, c, metadata, selector,
	scaleA, selA, scaleB, selB Operand,
) *Instruction {
	return &Instruction{
		Op:  ptx.OpMma,
		Dst: d,
		Src: []Operand{a, b, c, metadata, selector, scaleA, selA, scaleB, selB},
		Modifiers: []ptx.Modifier{
			spMod,
			ptx.ModSync,
			ptx.ModAligned,
			shape,
			alayout,
			blayout,
			kind,
			ptx.ModBlockScale,
			scaleVecSize,
		},
	}
}

// WGMMA synchronization

// WgmmaFence enforces memory consistency for wgmma operations.
func WgmmaFence() *Instruction {
	return &Instruction{
		Op:        ptx.OpWgmmaFence,
		Modifiers: []ptx.Modifier{ptx.ModSync, ptx.ModAligned},
	}
}

// WgmmaCommitGroup commits prior wgmma.mma_async operations to a group.
func WgmmaCommitGroup() *Instruction {
	return &Instruction{
		Op:        ptx.OpWgmmaCommitGroup,
		Modifiers: []ptx.Modifier{ptx.ModSync, ptx.ModAligned},
	}
}

// WgmmaWaitGroup waits for completion of a specific number of groups.
func WgmmaWaitGroup(n Operand) *Instruction {
	return &Instruction{
		Op:        ptx.OpWgmmaWaitGroup,
		Src:       []Operand{n},
		Modifiers: []ptx.Modifier{ptx.ModSync, ptx.ModAligned},
	}
}

// FenceProxyAsync synchronizes generic proxy with async proxy.
func FenceProxyAsync(scope ptx.Scope) *Instruction {
	return &Instruction{
		Op:    ptx.OpFenceProxyAsync,
		Scope: scope,
	}
}


// WgmmaMmaAsync creates a wgmma.mma_async instruction.
//
// Variants:
// 1. FP16/BF16: d, a, b, scaleD, immScaleA, immScaleB, immTransA, immTransB
// 2. TF32/FP8:  d, a, b, scaleD, immScaleA, immScaleB
// 3. Int/Bit:   d, a, b, scaleD
//
// Modifiers required: Shape (.m64...), Dtype (.f32, .s32), Atype (.f16...), Btype.
// Usage example:
//   WgmmaMmaAsync(
//     ModShapeM64N128K16, ModTypeF32, ModTypeF16, ModTypeF16,
//     d, a, b, scaleD, Imm(1), Imm(1), Imm(0), Imm(0)
//   )
func WgmmaMmaAsync(
    shape ptx.Modifier,
    dtype, atype, btype ptx.Modifier, // e.g. ModTypeF32, ModTypeF16, ModTypeF16
    d, a, b, scaleD Operand,
    extras ...Operand, // Immediates for scaling/transpose
) *Instruction {
    srcs := []Operand{a, b, scaleD}
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }
    
    return &Instruction{
        Op:  ptx.OpWgmmaMmaAsync,
        Dst: d,
        Src: srcs,
        Modifiers: []ptx.Modifier{
            ptx.ModSync,
            ptx.ModAligned,
            shape,
            dtype,
            atype,
            btype,
        },
    }
}



// WgmmaMmaAsyncSparse creates a sparse wgmma.mma_async.sp instruction.
//
// Syntax: wgmma.mma_async.sp.sync.aligned.shape.dtype.atype.btype d, a, b, sp-meta, sp-sel, scale-d, [imms...]
//
// spMeta: sparsity metadata register or immediate.
// spSel: sparsity selector immediate (0 or 1 for f16/tf32, 0 for int/fp8).
//
// To add optional modifiers like .satfinite (for integer types), chain .WithMod(ptx.ModSatFinite).
func WgmmaMmaAsyncSparse(
    shape ptx.Modifier,
    dtype, atype, btype ptx.Modifier,
    d, a, b, spMeta, spSel, scaleD Operand,
    extras ...Operand, // Immediates for scaling/transpose
) *Instruction {
    srcs := []Operand{a, b, spMeta, spSel, scaleD}
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }

    return &Instruction{
        Op:  ptx.OpWgmmaMmaAsync,
        Dst: d,
        Src: srcs,
        Modifiers: []ptx.Modifier{
            ptx.ModSp, // .sp
            ptx.ModSync,
            ptx.ModAligned,
            shape,
            dtype,
            atype,
            btype,
        },
    }
}



// --- 5th Gen Tensor Core (tcgen05) ---

// Tcgen05Mma performs Matrix Multiply-Accumulate.
// kind: e.g., ptx.ModKindTf32.
// ctaGroup: ptx.ModCtaGroup1 or ptx.ModCtaGroup2.
// instrDesc: Instruction Descriptor (32-bit register or immediate).
// smemDescA, smemDescB: Shared Memory Descriptors.
func Tcgen05Mma(kind, ctaGroup ptx.Modifier, instrDesc, smemDescA, smemDescB, tmemD, tmemC Operand) *Instruction {
    // Note: tmemC is optional in some contexts but typical for MMA (D = A*B + C).
    // The instruction signature is: tcgen05.mma... desc, descA, descB, tmemD, tmemC
    srcs := []Operand{instrDesc, smemDescA, smemDescB, tmemD}
    if tmemC != nil {
        srcs = append(srcs, tmemC)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Mma,
        Src:       srcs,
        Modifiers: []ptx.Modifier{ctaGroup, ptx.ModSync, ptx.ModAligned, kind},
    }
}

// Tcgen05Fence fences memory accesses.
// waitType: ptx.ModWaitLd or ptx.ModWaitSt (context dependent usage, usually before ld/st).
func Tcgen05Fence(waitType ptx.Modifier) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Fence,
        Modifiers: []ptx.Modifier{waitType, ptx.ModSync, ptx.ModAligned},
    }
}





// Tcgen05Alloc allocates Tensor Memory.
// Syntax: tcgen05.alloc.cta_group.sync.aligned{.shared::cta}.b32 [dst], nCols;
func Tcgen05Alloc(ctaGroup ptx.Modifier, dstAddr, nCols Operand) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Alloc,
        Src:       []Operand{dstAddr, nCols},
        Modifiers: []ptx.Modifier{ctaGroup, ptx.ModSync, ptx.ModAligned, ptx.ModSpaceSharedCTA, ptx.ModTypeB32},
    }
}

// Tcgen05Dealloc deallocates Tensor Memory.
// Syntax: tcgen05.dealloc.cta_group.sync.aligned.b32 taddr, nCols;
func Tcgen05Dealloc(ctaGroup ptx.Modifier, tAddr, nCols Operand) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Dealloc,
        Src:       []Operand{tAddr, nCols},
        Modifiers: []ptx.Modifier{ctaGroup, ptx.ModSync, ptx.ModAligned, ptx.ModTypeB32},
    }
}

// Tcgen05RelinquishAllocPermit releases allocation rights.
// Syntax: tcgen05.relinquish_alloc_permit.cta_group.sync.aligned;
func Tcgen05RelinquishAllocPermit(ctaGroup ptx.Modifier) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05RelinquishAllocPermit,
        Modifiers: []ptx.Modifier{ctaGroup, ptx.ModSync, ptx.ModAligned},
    }
}

// Tcgen05Ld loads from Tensor Memory to registers.
// Syntax: tcgen05.ld.sync.aligned.shape.num{.pack}.b32 r, [taddr], {immHalfSplitoff};
func Tcgen05Ld(shape, num ptx.Modifier, dst, tAddr Operand, splitOff ...Operand) *Instruction {
    srcs := []Operand{tAddr}
    if len(splitOff) > 0 {
        srcs = append(srcs, splitOff...)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Ld,
        Dst:       dst,
        Src:       srcs,
        Modifiers: []ptx.Modifier{ptx.ModSync, ptx.ModAligned, shape, num, ptx.ModTypeB32},
    }
}

// Tcgen05LdRed performs load with reduction.
func Tcgen05LdRed(shape, num, redOp, typ ptx.Modifier, dst, redVal, tAddr Operand, splitOff ...Operand) *Instruction {
    srcs := []Operand{tAddr}
    if len(splitOff) > 0 {
        srcs = append(srcs, splitOff...)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Ld,
        Dst:       dst, 
        Src:       srcs,
        // CHANGED: Use ptx.ModRed instead of ptx.OpRed
        Modifiers: []ptx.Modifier{ptx.ModRed, ptx.ModSync, ptx.ModAligned, shape, num, redOp, typ},
    }
}

// Tcgen05St stores from registers to Tensor Memory.
// Syntax: tcgen05.st.sync.aligned.shape.num{.unpack}.b32 [taddr], {immHalfSplitoff}, r;
func Tcgen05St(shape, num ptx.Modifier, tAddr, src Operand, splitOff ...Operand) *Instruction {
    srcs := []Operand{tAddr}
    if len(splitOff) > 0 {
        srcs = append(srcs, splitOff...)
    }
    srcs = append(srcs, src)
    return &Instruction{
        Op:        ptx.OpTcgen05St,
        Src:       srcs,
        Modifiers: []ptx.Modifier{ptx.ModSync, ptx.ModAligned, shape, num, ptx.ModTypeB32},
    }
}

// Tcgen05Wait waits for previous loads or stores.
// Syntax: tcgen05.wait::ld.sync.aligned;
// waitType: ptx.ModWaitLd or ptx.ModWaitSt.
func Tcgen05Wait(waitType ptx.Modifier) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Wait,
        Modifiers: []ptx.Modifier{waitType, ptx.ModSync, ptx.ModAligned},
    }
}

// Tcgen05Cp initiates asynchronous copy.
// Syntax: tcgen05.cp.cta_group.shape{.multicast}{.dst_fmt.src_fmt} [taddr], s-desc;
// Optional format modifiers can be chained via .WithMod().
func Tcgen05Cp(ctaGroup, shape ptx.Modifier, tAddr, sDesc Operand) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Cp,
        Src:       []Operand{tAddr, sDesc},
        Modifiers: []ptx.Modifier{ctaGroup, shape},
    }
}

// Tcgen05Shift shifts rows in Tensor Memory.
// Syntax: tcgen05.shift.cta_group.down [taddr];
func Tcgen05Shift(ctaGroup ptx.Modifier, tAddr Operand) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Shift,
        Src:       []Operand{tAddr},
        Modifiers: []ptx.Modifier{ctaGroup, ptx.ModShiftDown},
    }
}


// --- 5th Gen Tensor Core (tcgen05) Advanced MMA ---

// Tcgen05MmaSp performs Sparse Matrix Multiply-Accumulate.
// Syntax: tcgen05.mma.sp... [d], a, b, [sp-meta], idesc, ...
// Note: spMetaAddr is the address of the metadata in Tensor Memory.
func Tcgen05MmaSp(
    kind, ctaGroup ptx.Modifier,
    dAddr, aDesc, bDesc, spMetaAddr, iDesc Operand,
    extras ...Operand, // disable-output-lane, enable-input-d, etc.
) *Instruction {
    srcs := []Operand{aDesc, bDesc, spMetaAddr, iDesc}
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Mma,
        Dst:       dAddr, // [d-tmem] is conceptually a destination here
        Src:       srcs,
        Modifiers: []ptx.Modifier{ptx.ModSp, ctaGroup, kind},
    }
}

// Tcgen05MmaWs performs Weight Stationary MMA.
// Syntax: tcgen05.mma.ws... [d], a, b, idesc, ...
func Tcgen05MmaWs(
    kind, ctaGroup ptx.Modifier,
    dAddr, aDesc, bDesc, iDesc Operand,
    extras ...Operand, // enable-input-d, zero-column-mask-desc
) *Instruction {
    srcs := []Operand{aDesc, bDesc, iDesc}
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Mma,
        Dst:       dAddr,
        Src:       srcs,
        Modifiers: []ptx.Modifier{ptx.ModWS, ctaGroup, kind},
    }
}

// Tcgen05MmaWsSp performs Sparse Weight Stationary MMA.
// Syntax: tcgen05.mma.ws.sp... [d], a, b, [sp-meta], idesc, ...
func Tcgen05MmaWsSp(
    kind, ctaGroup ptx.Modifier,
    dAddr, aDesc, bDesc, spMetaAddr, iDesc Operand,
    extras ...Operand,
) *Instruction {
    srcs := []Operand{aDesc, bDesc, spMetaAddr, iDesc}
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Mma,
        Dst:       dAddr,
        Src:       srcs,
        Modifiers: []ptx.Modifier{ptx.ModWS, ptx.ModSp, ctaGroup, kind},
    }
}

// Tcgen05MmaScaled performs MMA with Block Scaling.
// Syntax: tcgen05.mma... [d], a, b, idesc, [scaleA], [scaleB], ...
func Tcgen05MmaScaled(
    kind, ctaGroup, scaleVec ptx.Modifier,
    dAddr, aDesc, bDesc, iDesc, scaleAAddr, scaleBAddr Operand,
    extras ...Operand,
) *Instruction {
    srcs := []Operand{aDesc, bDesc, iDesc, scaleAAddr, scaleBAddr}
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }
    return &Instruction{
        Op:        ptx.OpTcgen05Mma,
        Dst:       dAddr,
        Src:       srcs,
        Modifiers: []ptx.Modifier{ctaGroup, kind, ptx.ModBlockScale, scaleVec},
    }
}

// Tcgen05FenceSync performs specialized thread synchronization.
// syncType: ptx.ModBeforeThreadSync or ptx.ModAfterThreadSync
func Tcgen05FenceSync(syncType ptx.Modifier) *Instruction {
    return &Instruction{
        Op:        ptx.OpTcgen05Fence,
        Modifiers: []ptx.Modifier{syncType},
    }
}


// Brkpt suspends execution (breakpoint).
func Brkpt() *Instruction {
    return &Instruction{Op: ptx.OpBrkpt}
}

// NanoSleep suspends the thread for approx 't' nanoseconds.
// t: register or immediate (u32).
func NanoSleep(t Operand) *Instruction {
    return &Instruction{
        Op:  ptx.OpNanoSleep,
        Src: []Operand{t},
        Typ: ptx.U32,
    }
}


// Tcgen05Commit tracks completion of prior async operations.
// Syntax: tcgen05.commit.cta_group.completion_mechanism{.shared::cluster}{.multicast}.b64 [mbar] {, ctaMask};
//
// ctaGroup: ptx.ModCtaGroup1 or ptx.ModCtaGroup2
// mbar: mbarrier address operand
// ctaMask: optional 16-bit mask for multicast (requires ptx.ModMulticastCluster)
func Tcgen05Commit(ctaGroup ptx.Modifier, mbar Operand, ctaMask ...Operand) *Instruction {
    srcs := []Operand{mbar}
    if len(ctaMask) > 0 {
        srcs = append(srcs, ctaMask...)
    }
    
    // Base modifiers required by spec
    mods := []ptx.Modifier{
        ctaGroup,
        ptx.ModMbarrierArriveOne, // .mbarrier::arrive::one
        ptx.ModTypeB64,           // .b64
    }
    
    return &Instruction{
        Op:        ptx.OpTcgen05Commit,
        Src:       srcs,
        Modifiers: mods,
    }
}



// --- Video Instructions (Scalar) ---

// Vadd performs scalar video addition: d = a + b (+ c).
// dtype, atype, btype: .u32 or .s32
// extras: optional 'c' operand (for merge/accumulate) or modifiers (.sat, .add, .min, .max, selectors .b0, .h1, etc.)
func Vadd(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVadd, dtype, atype, btype, dst, a, b, extras...)
}

// Vsub performs scalar video subtraction: d = a - b (+ c).
func Vsub(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVsub, dtype, atype, btype, dst, a, b, extras...)
}

// Vabsdiff performs scalar video absolute difference: d = |a - b| (+ c).
func Vabsdiff(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVabsdiff, dtype, atype, btype, dst, a, b, extras...)
}

// Vmin performs scalar video minimum.
func Vmin(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVmin, dtype, atype, btype, dst, a, b, extras...)
}

// Vmax performs scalar video maximum.
func Vmax(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVmax, dtype, atype, btype, dst, a, b, extras...)
}

// Vshl performs scalar video shift left.
// Note: middle types are .atype.u32 (btype is always u32 for shift amount).
func Vshl(dtype, atype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVshl, dtype, atype, ptx.ModTypeU32, dst, a, b, extras...)
}

// Vshr performs scalar video shift right.
func Vshr(dtype, atype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVshr, dtype, atype, ptx.ModTypeU32, dst, a, b, extras...)
}

// Vmad performs scalar video multiply-accumulate: d = a * b + c.
// Supports .po (plus one) and scaling via modifiers in extras.
func Vmad(dtype, atype, btype ptx.Modifier, dst, a, b, c Operand, extras ...Operand) *Instruction {
    // Vmad always takes 3 sources
    args := []Operand{c} 
    args = append(args, extras...)
    return buildVideoInst(ptx.OpVmad, dtype, atype, btype, dst, a, b, args...)
}

// Vset performs scalar video comparison.
// cmp: ptx.CmpEq, ptx.CmpLt, etc.
func Vset(atype, btype ptx.Modifier, cmp ptx.CmpOp, dst, a, b Operand, extras ...Operand) *Instruction {
    inst := buildVideoInst(ptx.OpVset, atype, btype, 0, dst, a, b, extras...)
    inst.Cmp = cmp
    return inst
}

// --- Video Instructions (SIMD: .v2 and .v4) ---

// Vadd2 performs dual half-word addition.
func Vadd2(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVadd2, dtype, atype, btype, dst, a, b, extras...)
}

// Vsub2 performs dual half-word subtraction.
func Vsub2(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVsub2, dtype, atype, btype, dst, a, b, extras...)
}

// Vavrg2 performs dual half-word average.
func Vavrg2(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVavrg2, dtype, atype, btype, dst, a, b, extras...)
}

// Vset2 performs dual half-word comparison.
func Vset2(atype, btype ptx.Modifier, cmp ptx.CmpOp, dst, a, b Operand, extras ...Operand) *Instruction {
    inst := buildVideoInst(ptx.OpVset2, atype, btype, 0, dst, a, b, extras...)
    inst.Cmp = cmp
    return inst
}

// Vadd4 performs quad byte addition.
func Vadd4(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    return buildVideoInst(ptx.OpVadd4, dtype, atype, btype, dst, a, b, extras...)
}

// Vset4 performs quad byte comparison.
func Vset4(atype, btype ptx.Modifier, cmp ptx.CmpOp, dst, a, b Operand, extras ...Operand) *Instruction {
    inst := buildVideoInst(ptx.OpVset4, atype, btype, 0, dst, a, b, extras...)
    inst.Cmp = cmp
    return inst
}

// --- Helper for Video Instructions ---

func buildVideoInst(op ptx.Opcode, m1, m2, m3 ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction {
    srcs := []Operand{a, b}
    
    // Filter operands vs modifiers from extras (e.g., 'c' operand vs .sat)
    // In PTX video instructions, 'c' is the 3rd source if present.
    // We assume any Operand in extras is 'c', any Modifier is a modifier.
    // Note: The caller might pass 'c' as the first element of extras.
    
    // Base modifiers: dtype, atype, btype
    mods := []ptx.Modifier{m1, m2}
    if m3 != 0 {
        mods = append(mods, m3)
    }

    // Since our Instruction struct separates Src operands from Modifiers,
    // we don't need to manually filter if the user constructs 'extras' carefully,
    // but typically builder helpers don't mix them in one varargs. 
    // However, to support the signature `extras ...Operand` we assume extras contains 'c'.
    // If the user wants to add .sat, they should use .WithMod() on the result.
    // EXCEPT: if the user passes 'c', it must go into Src.
    
    if len(extras) > 0 {
        srcs = append(srcs, extras...)
    }

    return &Instruction{
        Op:        op,
        Dst:       dst,
        Src:       srcs,
        Modifiers: mods,
    }
}