package ptx

// Opcode represents a PTX instruction opcode.
type Opcode int

const (
	// Integer arithmetic
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

	// Extended-precision integer
	OpAddCC
	OpAddc
	OpSubCC
	OpSubc
	OpMadCC
	OpMadc

	// Floating point
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

	// Comparison & selection
	OpSet
	OpSetp
	OpSelp
	OpSlct

	// Logic & shift
	OpAnd
	OpOr
	OpXor
	OpNot
	OpCnot
	OpLop3
	OpShf
	OpShl
	OpShr

	// Data movement & conversion
	OpMov
	OpShfl
	OpPrmt
	OpLd
	OpLdNC
	OpLdu
	OpSt
	OpStAsync
	OpStBulk
	OpCvt
	OpCvtPack
	OpCvta
	OpPrefetch
	OpPrefetchu
	OpIsSpacep

	// Texture & surface
	OpTex
	OpTld4
	OpTxq
	OpSuld
	OpSust
	OpSured
	OpSuq

	// Control flow
	OpBra
	OpBrxIdx
	OpCall
	OpRet
	OpExit

	// Parallel synchronization
	OpBar
	OpBarWarp
	OpBarWarpSync
	OpBarrierCluster
	OpMembar
	OpFence
	OpAtom
	OpRed
	OpRedAsync
	OpVote
	OpVoteSync
	OpMatchSync
	OpActivemask
	OpReduxSync
	OpElectSync
	OpGriddepcontrol

	// Async copy
	OpCpAsync
	OpCpAsyncCommitGroup
	OpCpAsyncWaitGroup
	OpCpAsyncWaitAll
	OpCpAsyncBulk
	OpCpAsyncBulkCommitGroup
	OpCpAsyncBulkWaitGroup
	OpCpAsyncBulkPrefetch
	OpCpAsyncBulkTensor
	OpCpAsyncBulkPrefetchTensor
	OpCpAsyncMbarrierArrive
	OpCpReduceAsyncBulk
	OpCpReduceAsyncBulkTensor

	// Multimem
	OpMultimem
	OpMultimemLdReduce
	OpMultimemSt
	OpMultimemRed
	OpMultimemCpAsyncBulk
	OpMultimemCpReduceAsyncBulk

	// Warp matrix (tensor core)
	OpWmmaLoad
	OpWmmaStore
	OpWmmaMma
	OpMma
	OpWgmma
	OpLdMatrix
	OpStMatrix
	OpMovMatrix

	// Tensor map
	OpTensorMap
	OpTensormapReplace
	OpMapA
	OpMapa
	OpGetCTARank

	// Mbarrier
	OpMbarrierInit
	OpMbarrierInval
	OpMbarrierArrive
	OpMbarrierArriveDrop
	OpMbarrierTestWait
	OpMbarrierTryWait
	OpMbarrierExpectTx
	OpMbarrierCompleteTx
	OpMbarrierPendingCount

	// Vector / SIMD
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

	// Misc
	OpTrap
	OpBrkpt
	OpDiscard
	OpNanoSleep
	OpAlloca
	OpStackRestore
	OpStackSave
	OpCreatePolicy
	OpApplyPriority
	OpIstypep

	// Cluster launch control
	OpTensormapCpFenceproxy
	OpClusterlaunchcontrolTryCancel
	OpClusterlaunchcontrolQueryCancel

	// Warp matrix aliases
	OpLdmatrix
	OpStmatrix
	OpMovmatrix

	// WGMMA
	OpWgmmaFence
	OpWgmmaCommitGroup
	OpWgmmaWaitGroup
	OpWgmmaMmaAsync
	OpFenceProxyAsync

	// --- Tensor Core Gen 5 (sm_100+) ---
    OpTcgen05Alloc Opcode = iota + 1000 // Ensure unique ID
    OpTcgen05Dealloc
    OpTcgen05RelinquishAllocPermit
    OpTcgen05Ld
    OpTcgen05St
    OpTcgen05Cp
    OpTcgen05Shift
    OpTcgen05Mma
    OpTcgen05Commit
    OpTcgen05Wait
    OpTcgen05Fence


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
	case OpLdu:
		return "ldu"
	case OpSt:
		return "st"
	case OpStAsync:
		return "st.async"
	case OpStBulk:
		return "st.bulk"
	case OpCvt:
		return "cvt"
	case OpCvtPack:
		return "cvt.pack"
	case OpCvta:
		return "cvta"
	case OpPrefetch:
		return "prefetch"
	case OpPrefetchu:
		return "prefetchu"
	case OpIsSpacep:
		return "isspacep"
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
	case OpBarWarpSync:
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
	case OpVote:
		return "vote"
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
	case OpCpAsyncWaitAll:
		return "cp.async.wait_all"
	case OpCpAsyncBulk:
		return "cp.async.bulk"
	case OpCpAsyncBulkCommitGroup:
		return "cp.async.bulk.commit_group"
	case OpCpAsyncBulkWaitGroup:
		return "cp.async.bulk.wait_group"
	case OpCpAsyncBulkPrefetch:
		return "cp.async.bulk.prefetch"
	case OpCpAsyncBulkTensor:
		return "cp.async.bulk.tensor"
	case OpCpAsyncBulkPrefetchTensor:
		return "cp.async.bulk.prefetch.tensor"
	case OpCpAsyncMbarrierArrive:
		return "cp.async.mbarrier.arrive"
	case OpCpReduceAsyncBulk:
		return "cp.reduce.async.bulk"
	case OpCpReduceAsyncBulkTensor:
		return "cp.reduce.async.bulk.tensor"
	case OpMultimem:
		return "multimem"
	case OpMultimemLdReduce:
		return "multimem.ld_reduce"
	case OpMultimemSt:
		return "multimem.st"
	case OpMultimemRed:
		return "multimem.red"
	case OpMultimemCpAsyncBulk:
		return "multimem.cp.async.bulk"
	case OpMultimemCpReduceAsyncBulk:
		return "multimem.cp.reduce.async.bulk"
	case OpWmmaLoad:
		return "wmma.load"
	case OpWmmaStore:
		return "wmma.store"
	case OpWmmaMma:
		return "wmma.mma"
	case OpMma:
		return "mma"
	case OpWgmma:
		return "wgmma"
	case OpLdMatrix:
		return "ldmatrix"
	case OpStMatrix:
		return "stmatrix"
	case OpMovMatrix:
		return "movmatrix"
	case OpTensorMap:
		return "tensormap"
	case OpTensormapReplace:
		return "tensormap.replace"
	case OpMapA:
		return "mapa"
	case OpMapa:
		return "mapa"
	case OpGetCTARank:
		return "getctarank"
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
	case OpIstypep:
		return "istypep"
	case OpTensormapCpFenceproxy:
		return "tensormap.cp_fenceproxy"
	case OpClusterlaunchcontrolTryCancel:
		return "clusterlaunchcontrol.try_cancel"
	case OpClusterlaunchcontrolQueryCancel:
		return "clusterlaunchcontrol.query_cancel"
	case OpLdmatrix:
		return "ldmatrix"
	case OpStmatrix:
		return "stmatrix"
	case OpMovmatrix:
		return "movmatrix"
	case OpWgmmaFence:
		return "wgmma.fence"
	case OpWgmmaCommitGroup:
		return "wgmma.commit_group"
	case OpWgmmaWaitGroup:
		return "wgmma.wait_group"
	case OpWgmmaMmaAsync:
		return "wgmma.mma_async"
	case OpFenceProxyAsync:
		return "fence.proxy.async"


	case OpTcgen05Alloc: 
		return "tcgen05.alloc"
    case OpTcgen05Dealloc: 
		return "tcgen05.dealloc"
    case OpTcgen05RelinquishAllocPermit: 
		return "tcgen05.relinquish_alloc_permit"
    case OpTcgen05Ld: 
		return "tcgen05.ld"
    case OpTcgen05St: 
		return "tcgen05.st"
    case OpTcgen05Cp: 
		return "tcgen05.cp"
    case OpTcgen05Shift: 
		return "tcgen05.shift"
    case OpTcgen05Mma: 
		return "tcgen05.mma"
    case OpTcgen05Commit: 
		return "tcgen05.commit"
    case OpTcgen05Wait: 
		return "tcgen05.wait"
    case OpTcgen05Fence: 
		return "tcgen05.fence"


	default:
		return "unknown"
	}
}