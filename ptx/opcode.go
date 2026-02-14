package ptx

// Opcode represents a PTX instruction opcode.
type Opcode int

const (
	// ---- Integer arithmetic ----
	OpAdd Opcode = iota
	OpSub
	OpMul
	OpMad
	OpMul24
	OpMad24
	OpSad
	OpDiv
	OpRem
	OpAbs
	OpNeg
	OpMin
	OpMax
	OpPopc
	OpClz
	OpBfind
	OpBrev
	OpBfe
	OpBfi
	OpSzext
	OpBmsk
	OpDp4a
	OpDp2a
	OpFns

	// ---- Extended-precision integer ----
	OpAddCC // add.cc
	OpAddc
	OpSubCC // sub.cc
	OpSubc
	OpMadCC // mad.cc
	OpMadc

	// ---- Floating point ----
	OpFma
	OpRcp
	OpSqrt
	OpRsqrt
	OpSin
	OpCos
	OpLg2
	OpEx2
	OpTanh
	OpTestp
	OpCopysign

	// ---- Comparison & selection ----
	OpSet
	OpSetp
	OpSelp
	OpSlct

	// ---- Logic & shift ----
	OpAnd
	OpOr
	OpXor
	OpNot
	OpCnot
	OpLop3
	OpShf
	OpShl
	OpShr

	// ---- Data movement & conversion ----
	OpMov
	OpShfl // shfl.sync
	OpPrmt
	OpLd
	OpLdNC // ld.global.nc (non-coherent, read-only cache)
	OpSt
	OpStAsync
	OpCvt
	OpCvtPack
	OpCvta // convert address between generic and specific state space
	OpPrefetch

	// ---- Texture & surface ----
	OpTex
	OpTld4
	OpTxq
	OpSuld
	OpSust
	OpSured
	OpSuq

	// ---- Control flow ----
	OpBra
	OpBrxIdx // brx.idx or brx
	OpCall
	OpRet
	OpExit

	// ---- Parallel synchronization ----
	OpBar            // bar.sync
	OpBarWarp        // bar.warp.sync
	OpBarrierCluster // barrier.cluster
	OpMembar         // membar / fence
	OpFence
	OpAtom
	OpRed // reduction
	OpRedAsync
	OpVoteSync
	OpMatchSync
	OpActivemask
	OpReduxSync
	OpElectSync
	OpGriddepcontrol

	// ---- Async copy ----
	OpCpAsync
	OpCpAsyncCommitGroup
	OpCpAsyncWaitGroup
	OpCpAsyncBulk
	OpCpReduceAsyncBulk
	OpCpAsyncMbarrierArrive

	// ---- Warp matrix (tensor core) ----
	OpWmmaLoad
	OpWmmaStore
	OpWmmaMma
	OpMma

	// ---- Mbarrier ----
	OpMbarrierInit
	OpMbarrierInval
	OpMbarrierArrive
	OpMbarrierArriveDrop
	OpMbarrierTestWait
	OpMbarrierTryWait
	OpMbarrierExpectTx
	OpMbarrierCompleteTx
	OpMbarrierPendingCount

	// ---- New in PTX 9.1 / Recent ----

	// Vector Arithmetic / SIMD
	OpVadd
	OpVadd2
	OpVadd4
	OpVsub
	OpVsub2
	OpVsub4
	OpVmax
	OpVmax2
	OpVmax4
	OpVmin
	OpVmin2
	OpVmin4
	OpVabsdiff
	OpVabsdiff2
	OpVabsdiff4
	OpVavrg2
	OpVavrg4
	OpVset
	OpVset2
	OpVset4
	OpVshl
	OpVshr
	OpVmad

	// Matrix / Tensor / Cluster
	OpLdMatrix   // ldmatrix
	OpStMatrix   // stmatrix
	OpMovMatrix  // movmatrix
	OpWgmma      // wgmma
	OpMultimem   // multimem
	OpTensorMap  // tensormap
	OpMapA       // mapa
	OpGetCTARank // getctarank

	// Misc / Control / Debug
	OpTrap          // trap
	OpBrkpt         // brkpt
	OpDiscard       // discard
	OpNanoSleep     // nanosleep
	OpAlloca        // alloca
	OpStackRestore  // stackrestore
	OpStackSave     // stacksave
	OpCreatePolicy  // createpolicy
	OpApplyPriority // applypriority


	OpLdu Opcode = iota // ldu
    OpStBulk            // st.bulk
    OpMultimemLdReduce  // multimem.ld_reduce
    OpMultimemSt        // multimem.st
    OpMultimemRed       // multimem.red


	// ---- Data Movement (Advanced) ----
	OpPrefetchu                   // prefetchu
	OpIsSpacep                    // isspacep


	// ---- Data Movement (Async & packed) ----
    OpMapa                           // mapa
    OpCpAsyncWaitAll                 // cp.async.wait_all


    // ---- New additions for PTX 8.0 / 9.1 Bulk Async & Tensor ----

    OpCpAsyncBulkPrefetch       // cp.async.bulk.prefetch
    OpMultimemCpAsyncBulk       // multimem.cp.async.bulk
    OpMultimemCpReduceAsyncBulk // multimem.cp.reduce.async.bulk

	OpCpAsyncBulkTensor Opcode = iota + 453 // Aligning

	// ---- Tensor Reduction & Prefetch (Section 9.7.9.27.1.3 - 4) ----
	OpCpReduceAsyncBulkTensor   // cp.reduce.async.bulk.tensor
	OpCpAsyncBulkPrefetchTensor // cp.async.bulk.prefetch.tensor

	// ---- Bulk Group Management (Section 9.7.9.27.2) ----
	OpCpAsyncBulkCommitGroup // cp.async.bulk.commit_group
	OpCpAsyncBulkWaitGroup   // cp.async.bulk.wait_group

	// ---- Tensormap (Section 9.7.9.28) ----
	OpTensormapReplace // tensormap.replace

	OpIstypep                   // istypep

	// ---- Parallel Synchronization (Section 9.7.13) ----
	OpBarWarpSync      // bar.warp.sync
	OpVote             // vote (non-sync version)
)

func (o Opcode) String() string {
	switch o {
	case OpAdd:
		return "add"
	case OpSub:
		return "sub"
	case OpMul:
		return "mul"
	case OpMad:
		return "mad"
	case OpMul24:
		return "mul24"
	case OpMad24:
		return "mad24"
	case OpSad:
		return "sad"
	case OpDiv:
		return "div"
	case OpRem:
		return "rem"
	case OpAbs:
		return "abs"
	case OpNeg:
		return "neg"
	case OpMin:
		return "min"
	case OpMax:
		return "max"
	case OpPopc:
		return "popc"
	case OpClz:
		return "clz"
	case OpBfind:
		return "bfind"
	case OpBrev:
		return "brev"
	case OpBfe:
		return "bfe"
	case OpBfi:
		return "bfi"
	case OpSzext:
		return "szext"
	case OpBmsk:
		return "bmsk"
	case OpDp4a:
		return "dp4a"
	case OpDp2a:
		return "dp2a"
	case OpFns:
		return "fns"
	case OpAddCC:
		return "add.cc"
	case OpAddc:
		return "addc"
	case OpSubCC:
		return "sub.cc"
	case OpSubc:
		return "subc"
	case OpMadCC:
		return "mad.cc"
	case OpMadc:
		return "madc"
	case OpFma:
		return "fma"
	case OpRcp:
		return "rcp"
	case OpSqrt:
		return "sqrt"
	case OpRsqrt:
		return "rsqrt"
	case OpSin:
		return "sin"
	case OpCos:
		return "cos"
	case OpLg2:
		return "lg2"
	case OpEx2:
		return "ex2"
	case OpTanh:
		return "tanh"
	case OpTestp:
		return "testp"
	case OpCopysign:
		return "copysign"
	case OpSet:
		return "set"
	case OpSetp:
		return "setp"
	case OpSelp:
		return "selp"
	case OpSlct:
		return "slct"
	case OpAnd:
		return "and"
	case OpOr:
		return "or"
	case OpXor:
		return "xor"
	case OpNot:
		return "not"
	case OpCnot:
		return "cnot"
	case OpLop3:
		return "lop3"
	case OpShf:
		return "shf"
	case OpShl:
		return "shl"
	case OpShr:
		return "shr"
	case OpMov:
		return "mov"
	case OpShfl:
		return "shfl.sync"
	case OpPrmt:
		return "prmt"
	case OpLd:
		return "ld"
	case OpLdNC:
		return "ld.global.nc"
	case OpSt:
		return "st"
	case OpStAsync:
		return "st.async"
	case OpCvt:
		return "cvt"
	case OpCvtPack:
		return "cvt.pack"
	case OpCvta:
		return "cvta"
	case OpPrefetch:
		return "prefetch"
	case OpTex:
		return "tex"
	case OpTld4:
		return "tld4"
	case OpTxq:
		return "txq"
	case OpSuld:
		return "suld"
	case OpSust:
		return "sust"
	case OpSured:
		return "sured"
	case OpSuq:
		return "suq"
	case OpBra:
		return "bra"
	case OpBrxIdx:
		return "brx.idx"
	case OpCall:
		return "call"
	case OpRet:
		return "ret"
	case OpExit:
		return "exit"
	case OpBar:
		return "bar.sync"
	case OpBarWarp:
		return "bar.warp.sync"
	case OpBarrierCluster:
		return "barrier.cluster"
	case OpMembar:
		return "membar"
	case OpFence:
		return "fence"
	case OpAtom:
		return "atom"
	case OpRed:
		return "red"
	case OpRedAsync:
		return "red.async"
	case OpVoteSync:
		return "vote.sync"
	case OpMatchSync:
		return "match.sync"
	case OpActivemask:
		return "activemask"
	case OpReduxSync:
		return "redux.sync"
	case OpElectSync:
		return "elect.sync"
	case OpGriddepcontrol:
		return "griddepcontrol"
	case OpCpAsync:
		return "cp.async"
	case OpCpAsyncCommitGroup:
		return "cp.async.commit_group"
	case OpCpAsyncWaitGroup:
		return "cp.async.wait_group"
	case OpCpAsyncBulk:
		return "cp.async.bulk"
	case OpCpReduceAsyncBulk:
		return "cp.reduce.async.bulk"
	case OpCpAsyncMbarrierArrive:
        return "cp.async.mbarrier.arrive"
	case OpWmmaLoad:
		return "wmma.load"
	case OpWmmaStore:
		return "wmma.store"
	case OpWmmaMma:
		return "wmma.mma"
	case OpMma:
		return "mma"
	case OpMbarrierInit:
		return "mbarrier.init"
	case OpMbarrierInval:
		return "mbarrier.inval"
	case OpMbarrierArrive:
		return "mbarrier.arrive"
	case OpMbarrierArriveDrop:
		return "mbarrier.arrive_drop"
	case OpMbarrierTestWait:
		return "mbarrier.test_wait"
	case OpMbarrierTryWait:
		return "mbarrier.try_wait"
	case OpMbarrierExpectTx:
		return "mbarrier.expect_tx"
	case OpMbarrierCompleteTx:
		return "mbarrier.complete_tx"
	case OpMbarrierPendingCount:
		return "mbarrier.pending_count"
	case OpVadd:
		return "vadd"
	case OpVadd2:
		return "vadd2"
	case OpVadd4:
		return "vadd4"
	case OpVsub:
		return "vsub"
	case OpVsub2:
		return "vsub2"
	case OpVsub4:
		return "vsub4"
	case OpVmax:
		return "vmax"
	case OpVmax2:
		return "vmax2"
	case OpVmax4:
		return "vmax4"
	case OpVmin:
		return "vmin"
	case OpVmin2:
		return "vmin2"
	case OpVmin4:
		return "vmin4"
	case OpVabsdiff:
		return "vabsdiff"
	case OpVabsdiff2:
		return "vabsdiff2"
	case OpVabsdiff4:
		return "vabsdiff4"
	case OpVavrg2:
		return "vavrg2"
	case OpVavrg4:
		return "vavrg4"
	case OpVset:
		return "vset"
	case OpVset2:
		return "vset2"
	case OpVset4:
		return "vset4"
	case OpVshl:
		return "vshl"
	case OpVshr:
		return "vshr"
	case OpVmad:
		return "vmad"
	case OpLdMatrix:
		return "ldmatrix"
	case OpStMatrix:
		return "stmatrix"
	case OpMovMatrix:
		return "movmatrix"
	case OpWgmma:
		return "wgmma"
	case OpMultimem:
		return "multimem"
	case OpTensorMap:
		return "tensormap"
	case OpMapA:
		return "mapa"
	case OpGetCTARank:
		return "getctarank"
	case OpTrap:
		return "trap"
	case OpBrkpt:
		return "brkpt"
	case OpDiscard:
		return "discard"
	case OpNanoSleep:
		return "nanosleep"
	case OpAlloca:
		return "alloca"
	case OpStackRestore:
		return "stackrestore"
	case OpStackSave:
		return "stacksave"
	case OpCreatePolicy:
		return "createpolicy"
	case OpApplyPriority:
		return "applypriority"
	case OpLdu:
        return "ldu"
    case OpStBulk:
        return "st.bulk"
    case OpMultimemLdReduce:
        return "multimem.ld_reduce"
    case OpMultimemSt:
        return "multimem.st"
    case OpMultimemRed:
        return "multimem.red"
	case OpPrefetchu:
		return "prefetchu"
	case OpIsSpacep:
		return "isspacep"
    case OpMapa:
        return "mapa"
    case OpCpAsyncWaitAll:
        return "cp.async.wait_all"


    case OpCpAsyncBulkPrefetch:
        return "cp.async.bulk.prefetch"
    case OpMultimemCpAsyncBulk:
        return "multimem.cp.async.bulk"
    case OpMultimemCpReduceAsyncBulk:
        return "multimem.cp.reduce.async.bulk"
    case OpCpAsyncBulkTensor:
        return "cp.async.bulk.tensor"
	case OpCpReduceAsyncBulkTensor:
		return "cp.reduce.async.bulk.tensor"
	case OpCpAsyncBulkPrefetchTensor:
		return "cp.async.bulk.prefetch.tensor"
	case OpCpAsyncBulkCommitGroup:
		return "cp.async.bulk.commit_group"
	case OpCpAsyncBulkWaitGroup:
		return "cp.async.bulk.wait_group"
	case OpTensormapReplace:
		return "tensormap.replace"

	case OpBarWarpSync:
		return "bar.warp.sync"
	case OpVote:
		return "vote"

	default:
		return "unknown"
	}
}