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

// ---- Extended-precision integer arithmetic ----

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

func Tanh(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpTanh, Dst: dst, Src: []Operand{src}}
}

func Testp(dst, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpTestp, Dst: dst, Src: []Operand{src}}
}

func Copysign(dst, src0, src1 Operand) *Instruction {
	return &Instruction{Op: ptx.OpCopysign, Dst: dst, Src: []Operand{src0, src1}}
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

// ---- Comparison & selection ----

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

// ---- Data movement & conversion ----

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

// ---- Warp shuffle ----

func ShflSync(dst, a, b, c, mask Operand) *Instruction {
	return &Instruction{Op: ptx.OpShfl, Dst: dst, Src: []Operand{a, b, c, mask}, Modifiers: []ptx.Modifier{ptx.ModSync}}
}

func Shfl(dst, a, b, c Operand) *Instruction {
	return &Instruction{Op: ptx.OpShfl, Dst: dst, Src: []Operand{a, b, c}}
}

// ---- Stack manipulation ----

func Alloca(dst, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpAlloca, Dst: dst, Src: []Operand{size}}
}

func StackSave(dst Operand) *Instruction {
	return &Instruction{Op: ptx.OpStackSave, Dst: dst}
}

func StackRestore(src Operand) *Instruction {
	return &Instruction{Op: ptx.OpStackRestore, Src: []Operand{src}}
}

// ---- Control flow ----

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

// ---- Synchronization ----

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

func FenceProxyAsync(space ptx.StateSpace) *Instruction {
	return &Instruction{Op: ptx.OpFence, Space: space, Modifiers: []ptx.Modifier{ptx.ModProxy, ptx.ModAsync}}
}

// ---- Atomics ----

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

// ---- Warp voting ----

func VoteSync(mode ptx.Modifier, dst, mask, pred Operand) *Instruction {
	return &Instruction{Op: ptx.OpVoteSync, Dst: dst, Src: []Operand{mask, pred}, Modifiers: []ptx.Modifier{mode}}
}

func Activemask(dst Operand) *Instruction {
	return &Instruction{Op: ptx.OpActivemask, Dst: dst}
}

func ReduxSync(op ptx.Modifier, dst, mask, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpReduxSync, Dst: dst, Src: []Operand{mask, src}, Modifiers: []ptx.Modifier{op}}
}

// ---- Matrix ----

func LdMatrix(dst, addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpLdMatrix, Dst: dst, Src: []Operand{addr}}
}

func StMatrix(addr, src Operand) *Instruction {
	return &Instruction{Op: ptx.OpStMatrix, Src: []Operand{addr, src}}
}

func Wgmma(dst Operand, args ...Operand) *Instruction {
	return &Instruction{Op: ptx.OpWgmma, Dst: dst, Src: args}
}

// ---- Mbarrier ----

func MbarrierInit(addr, count Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierInit, Src: []Operand{addr, count}}
}

func MbarrierArrive(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpMbarrierArrive, Src: []Operand{addr}}
}

func CpAsyncMbarrierArrive(addr Operand) *Instruction {
	return &Instruction{Op: ptx.OpCpAsyncMbarrierArrive, Src: []Operand{addr}}
}

// ---- Multimem ----

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

// ---- Cache & eviction policy ----

func ApplyPriority(addr, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpApplyPriority, Src: []Operand{addr, size}}
}

func Discard(addr, size Operand) *Instruction {
	return &Instruction{Op: ptx.OpDiscard, Src: []Operand{addr, size}}
}

func CreatePolicy(dst Operand, args ...Operand) *Instruction {
	return &Instruction{Op: ptx.OpCreatePolicy, Dst: dst, Src: args}
}

// ---- Async copy ----

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

// ---- Async bulk tensor ----

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

// ---- Tensormap ----

func TensormapReplace(field ptx.Modifier, addr Operand, args ...Operand) *Instruction {
	srcs := []Operand{addr}
	srcs = append(srcs, args...)
	return &Instruction{Op: ptx.OpTensormapReplace, Src: srcs, Modifiers: []ptx.Modifier{ptx.ModLoadTile, field}}
}

// ---- Texture ----

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

// ---- Surface ----

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