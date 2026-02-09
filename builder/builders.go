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

func Cvta(dst, src Operand) *Instruction {
    return &Instruction{Op: ptx.OpCvta, Dst: dst, Src: []Operand{src}}
}

func Prefetch(addr Operand) *Instruction {
    return &Instruction{Op: ptx.OpPrefetch, Src: []Operand{addr}}
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
        // If multiple returns, store extras in Src prefix (codegen handles this)
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

// ---- Async copy ----

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