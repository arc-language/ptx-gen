# ptx-gen Package Documentation

## github.com/arc-language/ptx-gen/builder



TYPES

type Address struct {
	Base   Operand // *Register or *Symbol
	Offset int64   // byte offset
}
    Address represents a memory address operand: [base], [base+offset],
    or [symbol+offset].

func Addr(base Operand, offset int64) *Address
    Addr creates a memory address operand [base+offset].

type BasicBlock struct {
	Label        string
	Instructions []*Instruction
}
    BasicBlock represents a labeled sequence of instructions within a function.
    In PTX this is a label followed by instructions until the next label or end
    of function.

func (bb *BasicBlock) Add(inst *Instruction) *BasicBlock
    Add appends an instruction to this block.

type Directive struct {
	Kind   DirectiveKind
	Values []int  // numeric values (e.g. maxnreg=128, maxntid=256,1,1)
	Text   string // for .pragma
}
    Directive represents a single performance-tuning directive on a function.

func AbiPreserve(n int) *Directive
    AbiPreserve creates a .abi_preserve N directive.

func AbiPreserveCtrl(n int) *Directive
    AbiPreserveCtrl creates a .abi_preserve_control N directive.

func Alias(alias, aliasee string) *Directive
    Alias creates a .alias directive (module scope). Note: usage typically
    requires a custom emit function in codegen if not attached to a function.

func BlocksAreClusters() *Directive
    BlocksAreClusters creates a .blocksareclusters directive.

func ExplicitCluster() *Directive
    ExplicitCluster creates a .explicitcluster directive.

func MaxClusterRank(n int) *Directive
    MaxClusterRank creates a .maxclusterrank N directive.

func MaxNCTAPerSM(n int) *Directive
    MaxNCTAPerSM creates a .maxnctapersm directive.

func MaxNReg(n int) *Directive
    MaxNReg creates a .maxnreg directive.

func MaxNTid(dims ...int) *Directive
    MaxNTid creates a .maxntid directive (1D, 2D, or 3D).

func MinNCTAPerSM(n int) *Directive
    MinNCTAPerSM creates a .minnctapersm directive.

func NoReturn() *Directive
    NoReturn creates a .noreturn directive.

func Pragma(text string) *Directive
    Pragma creates a .pragma directive.

func ReqNTid(dims ...int) *Directive
    ReqNTid creates a .reqntid directive (1D, 2D, or 3D).

type DirectiveKind int
    DirectiveKind represents PTX performance-tuning and metadata directives.

const (
	DirMaxNReg           DirectiveKind = iota // .maxnreg N
	DirMaxNTid                                // .maxntid nx, ny, nz
	DirReqNTid                                // .reqntid nx, ny, nz
	DirMinNCTAPerSM                           // .minnctapersm N
	DirMaxNCTAPerSM                           // .maxnctapersm N
	DirPragma                                 // .pragma "string"
	DirReqNCluster                            // .reqnctapercluster (sm_90+)
	DirNoReturn                               // .noreturn
	DirAbiPreserve                            // .abi_preserve N
	DirAbiPreserveCtrl                        // .abi_preserve_control N
	DirExplicitCluster                        // .explicitcluster
	DirMaxClusterRank                         // .maxclusterrank N
	DirBlocksAreClusters                      // .blocksareclusters
	DirAlias                                  // .alias
)
type Function struct {
	Name     string
	IsKernel bool        // true = .entry, false = .func
	Linkage  ptx.Linkage // .visible, .extern, .weak, etc.

	// Parameters
	Params       []*Param // input parameters
	ReturnParams []*Param // return parameters (device functions only)

	// Body
	Blocks    []*BasicBlock // ordered basic blocks
	Registers []*Register   // all declared registers (collected for .reg declarations)

	// Performance tuning directives
	Directives []*Directive

	// Function Attributes (Section 5.4.8)
	// e.g. .attribute(.unified(uuid1, uuid2))
	Attributes []VarAttribute

	// Has unexported fields.
}
    Function represents a .entry (kernel) or .func (device function) definition.

func (f *Function) AddDirective(d *Directive) *Function
    AddDirective appends a performance-tuning directive.

func (f *Function) AddParam(p *Param) *Function
    AddParam appends a kernel or function input parameter.

func (f *Function) AddReturnParam(p *Param) *Function
    AddReturnParam appends a return parameter (device functions only).

func (f *Function) NewBlock(label string) *BasicBlock
    NewBlock creates a new labeled basic block and appends it to the function.

func (f *Function) NewReg(name string, typ ptx.Type) *Register
    NewReg declares a new named register scoped to this function.

func (f *Function) Param(name string) *Symbol
    Param looks up a parameter by name and returns it as a Symbol operand
    suitable for use in ld.param / st.param instructions.

func (f *Function) TempReg(typ ptx.Type) *Register
    TempReg creates an auto-named temporary register (%t0, %t1, ...).

func (f *Function) WithAttribute(attr VarAttribute) *Function
    WithAttribute adds a function attribute (e.g., .attribute(.unified(...))).

type Global struct {
	Name        string
	Space       ptx.StateSpace // .global, .shared, .const
	Typ         ptx.Type       // element type
	Vec         ptx.VectorSize // .v2, .v4 (Scalar for non-vector)
	Count       int            // array element count (0 = scalar)
	Align       int            // alignment in bytes (0 = default)
	Linkage     ptx.Linkage    // .visible, .extern, etc.
	Initializer []interface{}  // optional initializer values (int64, float64, etc.)
	Attributes  []VarAttribute // .attribute(...)
}
    Global represents a module-scope variable declaration in .global, .shared,
    or .const space.

    Examples:

        .global .f32 gvar;
        .global .f32 garr[100];
        .shared .f32 smem[256];
        .const .b32 lookup[16] = {0, 1, 2, ...};
        .global .align 16 .b8 buffer[4096];
        .global .attribute(.managed) .s32 g;

func NewGlobal(name string, space ptx.StateSpace, typ ptx.Type) *Global
    NewGlobal creates a scalar global variable.

func NewGlobalArray(name string, space ptx.StateSpace, typ ptx.Type, count int) *Global
    NewGlobalArray creates a global array variable.

func (g *Global) WithAlign(bytes int) *Global
    WithAlign sets the alignment.

func (g *Global) WithAttribute(attr VarAttribute) *Global
    WithAttribute adds a variable attribute (e.g., .managed, .unified).

func (g *Global) WithInit(vals ...interface{}) *Global
    WithInit sets the initializer values.

func (g *Global) WithLinkage(l ptx.Linkage) *Global
    WithLinkage sets the linkage directive (.visible, .extern, etc.).

type Immediate struct {
	Value interface{} // int64, uint64, float32, float64
}
    Immediate represents a literal constant value (integer or float).

func Imm(val int64) *Immediate
    Imm creates an immediate integer operand.

func ImmF32(val float32) *Immediate
    ImmF32 creates an immediate 32-bit float operand.

func ImmF64(val float64) *Immediate
    ImmF64 creates an immediate 64-bit float operand.

func ImmU(val uint64) *Immediate
    ImmU creates an immediate unsigned integer operand.

type Instruction struct {
	Op  ptx.Opcode // opcode: add, ld, st, mov, setp, bra, etc.
	Typ ptx.Type   // instruction type: .u32, .f32, etc.

	Dst  Operand // Primary destination
	Dst2 Operand // Secondary destination (for setp p|q)

	Src   []Operand      // source operands
	Space ptx.StateSpace // state space for ld/st

	Cmp    ptx.CmpOp  // comparison operator (.eq, .lt, etc.)
	BoolOp ptx.BoolOp // boolean operator (.and, .or, .xor)

	Rounding  ptx.RoundingMode // .rn, .rz, etc.
	Cache     ptx.CacheOp      // .ca, .cg, etc.
	Scope     ptx.Scope        // .cta, .gpu, etc.
	Vec       ptx.VectorSize   // .v2, .v4
	Modifiers []ptx.Modifier   // .wide, .lo, .sat, etc.
	Guard     *Predicate       // @p

	SrcType    ptx.Type // For cvt (source type)
	CallTarget string   // For call
}
    Instruction represents a single PTX instruction node in the IR tree.

func Abs(dst, src Operand) *Instruction

func Activemask(dst Operand) *Instruction

func Add(dst, src0, src1 Operand) *Instruction

func AddCC(dst, src0, src1 Operand) *Instruction

func Addc(dst, src0, src1 Operand) *Instruction

func Alloca(dst, size Operand) *Instruction

func And(dst, src0, src1 Operand) *Instruction

func ApplyPriority(addr, size Operand) *Instruction

func Atom(op ptx.Modifier, dst, addr, src Operand) *Instruction

func AtomCAS(dst, addr, compare, val Operand) *Instruction

func AtomExch(typ ptx.Type, dst, addr, val Operand) *Instruction

func AtomVector(op ptx.Opcode, vec ptx.VectorSize, typ ptx.Type, dst, addr, src Operand) *Instruction

func BarSync(id Operand) *Instruction

func BarSyncCount(id, count Operand) *Instruction

func BarWarpSync(membermask Operand) *Instruction

func BarrierCTA(id Operand, threadCount ...Operand) *Instruction

func BarrierClusterArrive() *Instruction

func BarrierClusterWait() *Instruction

func Bfe(dst, src, start, len Operand) *Instruction

func Bfi(dst, base, insert, start, len Operand) *Instruction

func Bfind(dst, src Operand) *Instruction

func Bmsk(dst, pos, width Operand) *Instruction

func Bra(label string) *Instruction

func BraUni(label string) *Instruction

func Brev(dst, src Operand) *Instruction

func Brkpt() *Instruction
    Brkpt suspends execution (breakpoint).

func BrxIdx(index Operand, targetList Operand) *Instruction

func Call(target string, retParams []Operand, args []Operand) *Instruction

func CallIndirect(funcPtr Operand, retParams []Operand, args []Operand, proto Operand) *Instruction

func ClusterlaunchcontrolQueryCancelGetFirstCTAId(dstVec, handle Operand) *Instruction
    ClusterlaunchcontrolQueryCancelGetFirstCTAId extracts the first CTA ID.
    Returns a vector of 4x b32.

func ClusterlaunchcontrolQueryCancelGetFirstCTAIdDim(dstReg, handle Operand, dim ptx.Modifier) *Instruction
    ClusterlaunchcontrolQueryCancelGetFirstCTAIdDim extracts a specific
    dimension (x, y, or z). dim should be ptx.ModDimX, ptx.ModDimY, or
    ptx.ModDimZ.

func ClusterlaunchcontrolQueryCancelIsCanceled(predDst, handle Operand) *Instruction
    ClusterlaunchcontrolQueryCancelIsCanceled checks if cancellation succeeded.

func ClusterlaunchcontrolTryCancel(addr, mbar Operand) *Instruction
    ClusterlaunchcontrolTryCancel requests cancellation of a cluster (sm_100+).

func Clz(dst, src Operand) *Instruction

func Cnot(dst, src Operand) *Instruction

func Copysign(dst, src0, src1 Operand) *Instruction

func Cos(dst, src Operand) *Instruction

func CpAsync(dst, src, size Operand, args ...Operand) *Instruction

func CpAsyncBulk(dst, src, size Operand, args ...Operand) *Instruction

func CpAsyncBulkCommitGroup() *Instruction

func CpAsyncBulkPrefetch(src, size Operand, policy Operand) *Instruction

func CpAsyncBulkPrefetchTensor(
	dim ptx.Modifier,
	tensorMap Operand,
	coords []Operand,
	im2colInfo []Operand,
) *Instruction

func CpAsyncBulkTensor(
	dim ptx.Modifier,
	dstMem Operand,
	tensorMap Operand,
	coords []Operand,
	mbar Operand,
	extras []Operand,
) *Instruction

func CpAsyncBulkWaitGroup(n int64) *Instruction

func CpAsyncCommitGroup() *Instruction

func CpAsyncMbarrierArrive(addr Operand) *Instruction

func CpAsyncMbarrierArriveNoInc(addr Operand) *Instruction

func CpAsyncWaitAll() *Instruction

func CpAsyncWaitGroup(n int64) *Instruction

func CpReduceAsyncBulk(dst, src, size Operand, mbar Operand) *Instruction

func CpReduceAsyncBulkTensor(
	dim ptx.Modifier,
	tensorMap Operand,
	coords []Operand,
	srcMem Operand,
) *Instruction

func CreatePolicy(dst Operand, args ...Operand) *Instruction

func Cvt(dst Operand, srcs ...Operand) *Instruction

func CvtPack(dst Operand, srcs ...Operand) *Instruction

func Cvta(dst, src Operand) *Instruction

func Discard(addr, size Operand) *Instruction

func Div(dst, src0, src1 Operand) *Instruction

func Dp2a(dst, src0, src1, src2 Operand) *Instruction

func Dp4a(dst, src0, src1, src2 Operand) *Instruction

func Ex2(dst, src Operand) *Instruction

func Exit() *Instruction

func Fence(scope ptx.Scope) *Instruction

func FenceAcqRel(scope ptx.Scope) *Instruction

func FenceProxy(kind ptx.Modifier) *Instruction

func FenceProxyAsync(scope ptx.Scope) *Instruction
    FenceProxyAsync synchronizes generic proxy with async proxy.

func FenceSC(scope ptx.Scope) *Instruction

func Fma(dst, src0, src1, src2 Operand) *Instruction

func Fns(dst, mask, base, offset Operand) *Instruction

func GetCTARank(dst, addr Operand) *Instruction

func IsSpacep(dst, addr Operand) *Instruction

func Istypep(typ ptx.Modifier, dstPred Operand, addr Operand) *Instruction

func Ld(dst, addr Operand) *Instruction

func LdGlobalNC(dst, addr Operand) *Instruction

func LdMatrix(dst, addr Operand) *Instruction

func LdNC(dst, addr Operand) *Instruction

func LdParam(dst Operand, param *Symbol) *Instruction

func LdWeak(dst, addr Operand) *Instruction

func Ldmatrix(shape, num ptx.Modifier, typ ptx.Type, dst, addr Operand) *Instruction
    Ldmatrix loads matrices from shared memory. shape: .m8n8, .m16n16, .m8n16
    num: .x1, .x2, .x4

func LdmatrixTrans(shape, num ptx.Modifier, typ ptx.Type, dst, addr Operand) *Instruction
    LdmatrixTrans loads matrices with transpose (.trans).

func Ldu(dst, addr Operand) *Instruction

func Lg2(dst, src Operand) *Instruction

func Lop3(dst, a, b, c, immLut Operand) *Instruction

func Mad(dst, src0, src1, src2 Operand) *Instruction

func Mad24(dst, src0, src1, src2 Operand) *Instruction

func MadCC(dst, src0, src1, src2 Operand) *Instruction

func Madc(dst, src0, src1, src2 Operand) *Instruction

func Mapa(dst, addr, ctaRank Operand) *Instruction

func Max(dst, src0, src1 Operand) *Instruction

func Max3(dst, src0, src1, src2 Operand) *Instruction

func MbarrierArrive(addr Operand) *Instruction

func MbarrierArriveDrop(dstState, addr, count Operand) *Instruction

func MbarrierArriveDropNoComplete(dstState, addr, count Operand) *Instruction

func MbarrierArriveExpectTx(dstState, addr, txCount Operand) *Instruction

func MbarrierArriveNoComplete(dstState, addr, count Operand) *Instruction

func MbarrierCompleteTx(addr, txCount Operand) *Instruction

func MbarrierExpectTx(addr, txCount Operand) *Instruction

func MbarrierInit(addr, count Operand) *Instruction

func MbarrierInval(addr Operand) *Instruction

func MbarrierPendingCount(countDst, state Operand) *Instruction

func MbarrierTestWait(waitComplete, addr, state Operand) *Instruction

func MbarrierTestWaitParity(waitComplete, addr, phaseParity Operand) *Instruction

func MbarrierTryWait(waitComplete, addr, state Operand, suspendHint Operand) *Instruction

func Membar(level ptx.Modifier) *Instruction

func MembarProxy() *Instruction

func Min(dst, src0, src1 Operand) *Instruction

func Min3(dst, src0, src1, src2 Operand) *Instruction

func Mma(shape, alayout, blayout ptx.Modifier, d, a, b, c Operand) *Instruction
    Mma creates an mma.sync.aligned instruction.

func MmaBlockScaled(shape, alayout, blayout, scaleVecSize ptx.Modifier, d, a, b, c, scaleA, selA, scaleB, selB Operand) *Instruction
    MmaBlockScaled creates an mma instruction with block scaling: D = (A
    * ScaleA) * (B * ScaleB) + C. scaleVecSize: e.g. ptx.ModScaleVec1x,
    ptx.ModScaleVec2x.

func MmaSparse(spMod, shape, alayout, blayout ptx.Modifier, d, a, b, c, metadata, selector Operand) *Instruction
    MmaSparse creates a sparse mma instruction. spMod: ptx.ModSp or
    ptx.ModSpOrderedMetadata. metadata: register containing indices of non-zero
    elements. selector: immediate or register (0 or 1) indicating metadata
    ownership.

func MmaSparseBlockScaled(
	spMod, shape, alayout, blayout ptx.Modifier,
	kind, scaleVecSize ptx.Modifier,
	d, a, b, c, metadata, selector,
	scaleA, selA, scaleB, selB Operand,
) *Instruction
    MmaSparseBlockScaled creates a sparse mma instruction with block scaling.
    selA and selB are typically VectorOp containing immediate values.

func Mov(dst, src Operand) *Instruction

func Movmatrix(shape ptx.Modifier, typ ptx.Type, dst, src Operand) *Instruction
    Movmatrix transposes a matrix in registers.

func Mul(dst, src0, src1 Operand) *Instruction

func Mul24(dst, src0, src1 Operand) *Instruction

func MultimemCpAsyncBulk(dst, src, size Operand, byteMask Operand) *Instruction

func MultimemCpReduceAsyncBulk(dst, src, size Operand) *Instruction

func MultimemLdReduce(dst, addr Operand) *Instruction

func MultimemRed(addr, src Operand) *Instruction

func MultimemSt(addr, src Operand) *Instruction

func NanoSleep(t Operand) *Instruction
    NanoSleep suspends the thread for approx 't' nanoseconds. t: register or
    immediate (u32).

func Neg(dst, src Operand) *Instruction

func Not(dst, src Operand) *Instruction

func Or(dst, src0, src1 Operand) *Instruction

func Pmevent(a Operand) *Instruction
    Pmevent triggers a single performance monitor event. a: immediate operand
    (0..15)

func PmeventMask(mask Operand) *Instruction
    PmeventMask triggers one or more performance monitor events via a mask.
    mask: 16-bit immediate mask

func Popc(dst, src Operand) *Instruction

func Prefetch(addr Operand) *Instruction

func Prefetchu(addr Operand) *Instruction

func Prmt(dst, a, b, c Operand) *Instruction

func Rcp(dst, src Operand) *Instruction

func Red(op ptx.Modifier, addr, src Operand) *Instruction

func ReduxSync(op ptx.Modifier, dst, mask, src Operand) *Instruction

func Rem(dst, src0, src1 Operand) *Instruction

func Ret() *Instruction

func Rsqrt(dst, src Operand) *Instruction

func Sad(dst, src0, src1, src2 Operand) *Instruction

func Selp(dst, a, b, p Operand) *Instruction

func Set(cmp ptx.CmpOp, dst, a, b Operand) *Instruction

func SetMaxNReg(action ptx.Modifier, count Operand) *Instruction
    SetMaxNReg hints to change the number of registers owned by the warp.
    action: ptx.ModInc or ptx.ModDec count: immediate integer (24..256, multiple
    of 8)

    Syntax: setmaxnreg.action.sync.aligned.u32 imm-reg-count;

func Setp(cmp ptx.CmpOp, dst, a, b Operand) *Instruction

func ShfL(dst, a, b, c Operand) *Instruction

func ShfR(dst, a, b, c Operand) *Instruction

func Shfl(dst, a, b, c Operand) *Instruction

func ShflSync(dst, a, b, c, mask Operand) *Instruction

func Shl(dst, src, amount Operand) *Instruction

func Shr(dst, src, amt Operand) *Instruction

func Sin(dst, src Operand) *Instruction

func Slct(dst, a, b, c Operand) *Instruction

func Sqrt(dst, src Operand) *Instruction

func St(addr, src Operand) *Instruction

func StAsync(addr, src Operand, mbar Operand) *Instruction

func StBulk(addr, size, initVal Operand) *Instruction

func StMatrix(addr, src Operand) *Instruction

func StWeak(addr, src Operand) *Instruction

func StackRestore(src Operand) *Instruction

func StackSave(dst Operand) *Instruction

func Stmatrix(shape, num ptx.Modifier, typ ptx.Type, addr, src Operand) *Instruction
    Stmatrix stores matrices to shared memory.

func Sub(dst, src0, src1 Operand) *Instruction

func SubCC(dst, src0, src1 Operand) *Instruction

func Subc(dst, src0, src1 Operand) *Instruction

func Suld(geom ptx.Modifier, dst Operand, surf Operand, coords []Operand) *Instruction

func Suq(query ptx.Modifier, dst Operand, surf Operand) *Instruction

func Sured(geom ptx.Modifier, surf Operand, coords []Operand, val Operand) *Instruction

func Sust(geom ptx.Modifier, surf Operand, coords []Operand, val Operand) *Instruction

func Szext(dst, src, size Operand) *Instruction

func Tanh(dst, src Operand) *Instruction

func Tcgen05Alloc(ctaGroup ptx.Modifier, dstAddr, nCols Operand) *Instruction
    Tcgen05Alloc allocates Tensor Memory. Syntax:
    tcgen05.alloc.cta_group.sync.aligned{.shared::cta}.b32 [dst], nCols;

func Tcgen05Commit(ctaGroup ptx.Modifier, mbar Operand, ctaMask ...Operand) *Instruction
    Tcgen05Commit tracks completion of prior async operations. Syntax:
    tcgen05.commit.cta_group.completion_mechanism{.shared::cluster}{.multicast}.b64
    [mbar] {, ctaMask};

    ctaGroup: ptx.ModCtaGroup1 or ptx.ModCtaGroup2 mbar: mbarrier
    address operand ctaMask: optional 16-bit mask for multicast (requires
    ptx.ModMulticastCluster)

func Tcgen05Cp(ctaGroup, shape ptx.Modifier, tAddr, sDesc Operand) *Instruction
    Tcgen05Cp initiates asynchronous copy. Syntax:
    tcgen05.cp.cta_group.shape{.multicast}{.dst_fmt.src_fmt} [taddr], s-desc;
    Optional format modifiers can be chained via .WithMod().

func Tcgen05Dealloc(ctaGroup ptx.Modifier, tAddr, nCols Operand) *Instruction
    Tcgen05Dealloc deallocates Tensor Memory. Syntax:
    tcgen05.dealloc.cta_group.sync.aligned.b32 taddr, nCols;

func Tcgen05Fence(waitType ptx.Modifier) *Instruction
    Tcgen05Fence fences memory accesses. waitType: ptx.ModWaitLd or
    ptx.ModWaitSt (context dependent usage, usually before ld/st).

func Tcgen05FenceSync(syncType ptx.Modifier) *Instruction
    Tcgen05FenceSync performs specialized thread synchronization. syncType:
    ptx.ModBeforeThreadSync or ptx.ModAfterThreadSync

func Tcgen05Ld(shape, num ptx.Modifier, dst, tAddr Operand, splitOff ...Operand) *Instruction
    Tcgen05Ld loads from Tensor Memory to registers. Syntax:
    tcgen05.ld.sync.aligned.shape.num{.pack}.b32 r, [taddr], {immHalfSplitoff};

func Tcgen05LdRed(shape, num, redOp, typ ptx.Modifier, dst, redVal, tAddr Operand, splitOff ...Operand) *Instruction
    Tcgen05LdRed performs load with reduction.

func Tcgen05Mma(kind, ctaGroup ptx.Modifier, instrDesc, smemDescA, smemDescB, tmemD, tmemC Operand) *Instruction
    Tcgen05Mma performs Matrix Multiply-Accumulate. kind: e.g., ptx.ModKindTf32.
    ctaGroup: ptx.ModCtaGroup1 or ptx.ModCtaGroup2. instrDesc: Instruction
    Descriptor (32-bit register or immediate). smemDescA, smemDescB: Shared
    Memory Descriptors.

func Tcgen05MmaScaled(
	kind, ctaGroup, scaleVec ptx.Modifier,
	dAddr, aDesc, bDesc, iDesc, scaleAAddr, scaleBAddr Operand,
	extras ...Operand,
) *Instruction
    Tcgen05MmaScaled performs MMA with Block Scaling. Syntax: tcgen05.mma...
    [d], a, b, idesc, [scaleA], [scaleB], ...

func Tcgen05MmaSp(
	kind, ctaGroup ptx.Modifier,
	dAddr, aDesc, bDesc, spMetaAddr, iDesc Operand,
	extras ...Operand,
) *Instruction
    Tcgen05MmaSp performs Sparse Matrix Multiply-Accumulate. Syntax:
    tcgen05.mma.sp... [d], a, b, [sp-meta], idesc, ... Note: spMetaAddr is the
    address of the metadata in Tensor Memory.

func Tcgen05MmaWs(
	kind, ctaGroup ptx.Modifier,
	dAddr, aDesc, bDesc, iDesc Operand,
	extras ...Operand,
) *Instruction
    Tcgen05MmaWs performs Weight Stationary MMA. Syntax: tcgen05.mma.ws... [d],
    a, b, idesc, ...

func Tcgen05MmaWsSp(
	kind, ctaGroup ptx.Modifier,
	dAddr, aDesc, bDesc, spMetaAddr, iDesc Operand,
	extras ...Operand,
) *Instruction
    Tcgen05MmaWsSp performs Sparse Weight Stationary MMA. Syntax:
    tcgen05.mma.ws.sp... [d], a, b, [sp-meta], idesc, ...

func Tcgen05RelinquishAllocPermit(ctaGroup ptx.Modifier) *Instruction
    Tcgen05RelinquishAllocPermit releases allocation rights. Syntax:
    tcgen05.relinquish_alloc_permit.cta_group.sync.aligned;

func Tcgen05Shift(ctaGroup ptx.Modifier, tAddr Operand) *Instruction
    Tcgen05Shift shifts rows in Tensor Memory. Syntax:
    tcgen05.shift.cta_group.down [taddr];

func Tcgen05St(shape, num ptx.Modifier, tAddr, src Operand, splitOff ...Operand) *Instruction
    Tcgen05St stores from registers to Tensor Memory. Syntax:
    tcgen05.st.sync.aligned.shape.num{.unpack}.b32 [taddr], {immHalfSplitoff},
    r;

func Tcgen05Wait(waitType ptx.Modifier) *Instruction
    Tcgen05Wait waits for previous loads or stores. Syntax:
    tcgen05.wait::ld.sync.aligned; waitType: ptx.ModWaitLd or ptx.ModWaitSt.

func TensormapCpFenceproxy(dstGlobal, srcShared, size Operand, scope ptx.Scope) *Instruction
    TensormapCpFenceproxy performs a fused copy and fence (sm_90+). Implies
    .global.shared::cta, .tensormap::generic, .sync, .aligned.

func TensormapReplace(field ptx.Modifier, addr Operand, args ...Operand) *Instruction

func Testp(dst, src Operand) *Instruction

func Tex(geom ptx.Modifier, dst Operand, tex Operand, sampler Operand, coords []Operand) *Instruction

func Tld4(comp ptx.Modifier, geom ptx.Modifier, dst Operand, tex Operand, sampler Operand, coords []Operand) *Instruction

func Trap() *Instruction
    Trap aborts execution and generates an interrupt.

func Txq(query ptx.Modifier, dst Operand, tex Operand, lod Operand) *Instruction

func Vabsdiff(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vabsdiff performs scalar video absolute difference: d = |a - b| (+ c).

func Vadd(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vadd performs scalar video addition: d = a + b (+ c). dtype, atype,
    btype: .u32 or .s32 extras: optional 'c' operand (for merge/accumulate) or
    modifiers (.sat, .add, .min, .max, selectors .b0, .h1, etc.)

func Vadd2(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vadd2 performs dual half-word addition.

func Vadd4(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vadd4 performs quad byte addition.

func Vavrg2(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vavrg2 performs dual half-word average.

func Vmad(dtype, atype, btype ptx.Modifier, dst, a, b, c Operand, extras ...Operand) *Instruction
    Vmad performs scalar video multiply-accumulate: d = a * b + c. Supports .po
    (plus one) and scaling via modifiers in extras.

func Vmax(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vmax performs scalar video maximum.

func Vmin(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vmin performs scalar video minimum.

func VoteSync(mode ptx.Modifier, dst, mask, pred Operand) *Instruction

func Vset(atype, btype ptx.Modifier, cmp ptx.CmpOp, dst, a, b Operand, extras ...Operand) *Instruction
    Vset performs scalar video comparison. cmp: ptx.CmpEq, ptx.CmpLt, etc.

func Vset2(atype, btype ptx.Modifier, cmp ptx.CmpOp, dst, a, b Operand, extras ...Operand) *Instruction
    Vset2 performs dual half-word comparison.

func Vset4(atype, btype ptx.Modifier, cmp ptx.CmpOp, dst, a, b Operand, extras ...Operand) *Instruction
    Vset4 performs quad byte comparison.

func Vshl(dtype, atype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vshl performs scalar video shift left. Note: middle types are .atype.u32
    (btype is always u32 for shift amount).

func Vshr(dtype, atype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vshr performs scalar video shift right.

func Vsub(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vsub performs scalar video subtraction: d = a - b (+ c).

func Vsub2(dtype, atype, btype ptx.Modifier, dst, a, b Operand, extras ...Operand) *Instruction
    Vsub2 performs dual half-word subtraction.

func Wgmma(dst Operand, args ...Operand) *Instruction

func WgmmaCommitGroup() *Instruction
    WgmmaCommitGroup commits prior wgmma.mma_async operations to a group.

func WgmmaFence() *Instruction
    WgmmaFence enforces memory consistency for wgmma operations.

func WgmmaMmaAsync(
	shape ptx.Modifier,
	dtype, atype, btype ptx.Modifier,
	d, a, b, scaleD Operand,
	extras ...Operand,
) *Instruction
    WgmmaMmaAsync creates a wgmma.mma_async instruction.

    Variants: 1. FP16/BF16: d, a, b, scaleD, immScaleA, immScaleB, immTransA,
    immTransB 2. TF32/FP8: d, a, b, scaleD, immScaleA, immScaleB 3. Int/Bit: d,
    a, b, scaleD

    Modifiers required: Shape (.m64...), Dtype (.f32, .s32), Atype (.f16...),
    Btype. Usage example:

        WgmmaMmaAsync(
          ModShapeM64N128K16, ModTypeF32, ModTypeF16, ModTypeF16,
          d, a, b, scaleD, Imm(1), Imm(1), Imm(0), Imm(0)
        )

func WgmmaMmaAsyncSparse(
	shape ptx.Modifier,
	dtype, atype, btype ptx.Modifier,
	d, a, b, spMeta, spSel, scaleD Operand,
	extras ...Operand,
) *Instruction
    WgmmaMmaAsyncSparse creates a sparse wgmma.mma_async.sp instruction.

    Syntax: wgmma.mma_async.sp.sync.aligned.shape.dtype.atype.btype d, a, b,
    sp-meta, sp-sel, scale-d, [imms...]

    spMeta: sparsity metadata register or immediate. spSel: sparsity selector
    immediate (0 or 1 for f16/tf32, 0 for int/fp8).

    To add optional modifiers like .satfinite (for integer types), chain
    .WithMod(ptx.ModSatFinite).

func WgmmaWaitGroup(n Operand) *Instruction
    WgmmaWaitGroup waits for completion of a specific number of groups.

func WmmaLoad(role, layout, shape ptx.Modifier, typ ptx.Type, dst, addr, stride Operand) *Instruction
    WmmaLoad creates a wmma.load instruction. role: ModMatrixA, ModMatrixB,
    or ModMatrixC. layout: ModRow or ModCol. stride: optional (pass nil for
    default).

func WmmaMma(shape, alayout, blayout ptx.Modifier, d, a, b, c Operand) *Instruction
    WmmaMma creates a wmma.mma instruction performing D = A * B + C.

func WmmaMmaBitOp(op, shape, alayout, blayout ptx.Modifier, d, a, b, c Operand) *Instruction
    WmmaMmaBitOp creates a wmma.mma.op.popc instruction for .b1 types. op:
    ptx.ModAtomXor or ptx.ModAtomAnd.

func WmmaStore(role, layout, shape ptx.Modifier, typ ptx.Type, addr, src, stride Operand) *Instruction
    WmmaStore creates a wmma.store instruction. role: ModMatrixD. layout:
    ModRow or ModCol. stride: optional (pass nil for default).

func Xor(dst, src0, src1 Operand) *Instruction

func (i *Instruction) From(t ptx.Type) *Instruction
    From sets the source type for cvt instructions (cvt.dstType.srcType).

func (i *Instruction) InSpace(s ptx.StateSpace) *Instruction
    InSpace sets the state space (.global, .shared, .param, .local, .const).

func (i *Instruction) Pred(reg *Register) *Instruction
    Pred sets the guard predicate (@p).

func (i *Instruction) PredNot(reg *Register) *Instruction
    PredNot sets a negated guard predicate (@!p).

func (i *Instruction) SourceTyped(t ptx.Type) *Instruction
    SourceTyped sets the source type for mixed-precision or conversion
    instructions. Example: add.f32.f16 -> Typ=F32, SrcType=F16

func (i *Instruction) Typed(t ptx.Type) *Instruction
    Typed sets the instruction type (.u32, .f32, .f64, etc.).

func (i *Instruction) WithBoolOp(op ptx.BoolOp) *Instruction
    WithBoolOp sets the boolean operator for set/setp (e.g., .and, .or).

func (i *Instruction) WithCache(c ptx.CacheOp) *Instruction
    WithCache sets the cache operator (.ca, .cg, .cs, .lu, .cv, .wb, .wt).

func (i *Instruction) WithDst2(dst Operand) *Instruction
    WithDst2 sets the secondary destination (e.g., q in setp p|q, ...).

func (i *Instruction) WithMod(mods ...ptx.Modifier) *Instruction
    WithMod appends one or more modifiers (.wide, .lo, .hi, .sat, .ftz, .approx,
    etc.).

func (i *Instruction) WithRounding(r ptx.RoundingMode) *Instruction
    WithRounding sets the rounding modifier (.rn, .rz, .rm, .rp).

func (i *Instruction) WithScope(s ptx.Scope) *Instruction
    WithScope sets the memory scope (.cta, .gpu, .sys).

func (i *Instruction) WithVec(v ptx.VectorSize) *Instruction
    WithVec sets the vector width (.v2, .v4) for vector ld/st/mov.

type Module struct {
	Version     ptx.ISAVersion // .version 8.5
	Target      ptx.Target     // .target sm_80
	AddressSize int            // .address_size 64 (32 or 64)
	Globals     []*Global      // module-scope variables (.global, .const, .shared)
	Functions   []*Function    // .entry and .func definitions
}
    Module is the top-level PTX program container. It maps to a single .ptx file
    with directives, globals, and functions.

func NewModule(version ptx.ISAVersion, target ptx.Target) *Module
    NewModule creates a new PTX module with sensible defaults.

func (m *Module) AddFunction(f *Function) *Module
    AddFunction appends a function or kernel definition.

func (m *Module) AddGlobal(g *Global) *Module
    AddGlobal appends a module-scope variable declaration.

func (m *Module) NewFunc(name string) *Function
    NewFunc creates a new device function (.func) and adds it to the module.

func (m *Module) NewKernel(name string) *Function
    NewKernel creates a new kernel (.entry) function and adds it to the module.

type Operand interface {
	// Has unexported methods.
}
    Operand is the interface for all instruction operands.

type Param struct {
	Name      string
	Typ       ptx.Type       // .u32, .u64, .f32, etc.
	Size      int            // byte size for .param .align N .b8 name[Size] style
	Align     int            // optional alignment in bytes (0 = default)
	IsPointer bool           // if true, this param is a pointer (.ptr attribute)
	PtrSpace  ptx.StateSpace // state space the pointer points to (.global, .shared, etc.)
}
    Param represents a kernel or device function parameter.

    Kernel parameters live in .param state space and are read via ld.param.
    Device function parameters can be registers or .param byte arrays.

func NewByteArrayParam(name string, size int, align int) *Param
    NewByteArrayParam creates a .param .align A .b8 name[size] parameter used
    for passing structures by value.

func NewParam(name string, typ ptx.Type) *Param
    NewParam creates a simple typed parameter (e.g. .param .u32 N).

func NewPtrParam(name string, pointsTo ptx.StateSpace) *Param
    NewPtrParam creates a pointer parameter with .ptr attribute. Pointers are
    passed as .u64 in 64-bit addressing mode.

func (p *Param) WithAlign(align int) *Param
    WithAlign sets the alignment for the parameter.

type Predicate struct {
	Reg    *Register
	Negate bool
}
    Predicate represents a guard predicate on an instruction: @p or @!p

type Register struct {
	Name string
	Typ  ptx.Type
}
    Register represents a named PTX register (%r0, %fd1, %p0).

type SpecialRegOp struct {
	Reg ptx.SpecialReg
}
    SpecialRegOp represents a PTX special register operand (%tid.x, %ctaid.y,
    etc.).

func SReg(r ptx.SpecialReg) *SpecialRegOp
    SReg creates a special register operand (%tid.x, %ntid.x, etc.).

type Symbol struct {
	Name string
}
    Symbol represents a named symbol: function name, label, global variable,
    or parameter name.

func Sym(name string) *Symbol
    Sym creates a symbol operand (label, function name, global name).

type VarAttribute struct {
	Name   string
	Params []interface{}
}
    VarAttribute represents attributes for variables (Section 5.4.8). Examples:
    .attribute(.managed), .attribute(.unified(uuid1, uuid2))

func Managed() VarAttribute
    Managed creates a .managed attribute. This attribute specifies that variable
    will be allocated at a location in unified virtual memory environment where
    host and other devices can reference it.

func Unified(uuid1, uuid2 uint64) VarAttribute
    Unified creates a .unified attribute with UUID values. This attribute
    specifies that the variable has the same memory address on the host and on
    other devices in the system.

type VectorOp struct {
	Elements []Operand
}
    VectorOp represents a vector operand {r0, r1, r2, r3} for .v2/.v4 ld/st.

func Vec(elems ...Operand) *VectorOp
    Vec creates a vector operand from multiple registers/operands.


## github.com/arc-language/ptx-gen/codegen



FUNCTIONS

func Emit(mod *builder.Module) string
    Emit takes a complete builder.Module and returns the PTX source string.


TYPES

type Emitter struct {
	// Has unexported fields.
}
    Emitter holds state during PTX text generation.


## github.com/arc-language/ptx-gen/ptx



VARIABLES

var (
	ISA60 = ISAVersion{6, 0}
	ISA63 = ISAVersion{6, 3}
	ISA64 = ISAVersion{6, 4}
	ISA70 = ISAVersion{7, 0}
	ISA71 = ISAVersion{7, 1}
	ISA72 = ISAVersion{7, 2}
	ISA73 = ISAVersion{7, 3}
	ISA74 = ISAVersion{7, 4}
	ISA75 = ISAVersion{7, 5}
	ISA76 = ISAVersion{7, 6}
	ISA77 = ISAVersion{7, 7}
	ISA78 = ISAVersion{7, 8}
	ISA80 = ISAVersion{8, 0}
	ISA81 = ISAVersion{8, 1}
	ISA82 = ISAVersion{8, 2}
	ISA83 = ISAVersion{8, 3}
	ISA84 = ISAVersion{8, 4}
	ISA85 = ISAVersion{8, 5}
	ISA86 = ISAVersion{8, 6}
	ISA87 = ISAVersion{8, 7}
	ISA88 = ISAVersion{8, 8}
	ISA90 = ISAVersion{9, 0}
	ISA91 = ISAVersion{9, 1}
)
    Common ISA versions.


TYPES

type BoolOp int
    BoolOp represents PTX boolean operators for set/setp instructions.

const (
	BoolNone BoolOp = iota
	BoolAnd         // .and
	BoolOr          // .or
	BoolXor         // .xor
)
func (b BoolOp) String() string

type CacheOp int
    CacheOp represents PTX cache operators for ld/st instructions.

const (
	CacheNone CacheOp = iota

	// Load cache operators
	CacheCA // .ca — cache at all levels (default for ld)
	CacheCG // .cg — cache at global level (L2 only)
	CacheCS // .cs — cache streaming (likely to be accessed once)
	CacheLU // .lu — last use (evict first)
	CacheCV // .cv — don't cache, volatile (bypass L1)

	// Store cache operators
	CacheWB // .wb — write back at all levels (default for st)
	CacheWT // .wt — write through to system memory
)
func (c CacheOp) String() string

type CmpOp int
    CmpOp represents PTX comparison operators.

const (
	// Integer & bit-size comparisons
	CmpEq CmpOp = iota // eq  — equal
	CmpNe              // ne  — not equal
	CmpLt              // lt  — less than
	CmpLe              // le  — less than or equal
	CmpGt              // gt  — greater than
	CmpGe              // ge  — greater than or equal
	CmpLo              // lo  — lower (unsigned less than)
	CmpLs              // ls  — lower or same (unsigned <=)
	CmpHi              // hi  — higher (unsigned greater than)
	CmpHs              // hs  — higher or same (unsigned >=)

	// Float ordered comparisons (neither operand is NaN)
	CmpEqu // equ — equal, unordered
	CmpNeu // neu — not equal, unordered
	CmpLtu // ltu — less than, unordered
	CmpLeu // leu — less or equal, unordered
	CmpGtu // gtu — greater than, unordered
	CmpGeu // geu — greater or equal, unordered

	// Float NaN-testing comparisons
	CmpNum // num — both operands are numbers (not NaN)
	CmpNan // nan — either operand is NaN
)
func (c CmpOp) String() string

type ISAVersion struct {
	Major int
	Minor int
}
    ISAVersion represents the PTX ISA version string.

func (v ISAVersion) String() string

type Linkage int
    Linkage represents PTX visibility/linkage directives.

const (
	LinkNone    Linkage = iota
	LinkVisible         // .visible — externally visible
	LinkExtern          // .extern  — declared externally
	LinkWeak            // .weak    — weak linkage
	LinkCommon          // .common  — common symbol (multiple definitions allowed)
)
func (l Linkage) String() string

type MembarLevel int
    MembarLevel represents membar fence levels.

const (
	MembarCTA MembarLevel = iota // membar.cta
	MembarGL                     // membar.gl  (global)
	MembarSys                    // membar.sys (system)
)
func (m MembarLevel) String() string

type Modifier int
    Modifier represents miscellaneous PTX instruction modifiers.

const (
	// Multiply width
	ModWide Modifier = iota
	ModLo
	ModHi

	// Saturation & flush
	ModSat
	ModFtz

	// Approximation
	ModApprox
	ModFull

	// Uniformity
	ModUni

	// Memory consistency
	ModAcquire
	ModRelease
	ModRelaxed
	ModAcqRel
	ModVolatile
	ModMMIO
	ModWeak
	ModSC
	ModProxy
	ModAlias
	ModAsync

	// Sync
	ModSync

	// Testp
	ModFinite
	ModInfinite
	ModNumber
	ModNotANumber
	ModNormal
	ModSubnormal

	// Shfl modes
	ModShflUp
	ModShflDown
	ModShflBfly
	ModShflIdx

	// Atomic operations
	ModAtomAdd
	ModAtomMin
	ModAtomMax
	ModAtomInc
	ModAtomDec
	ModAtomCAS
	ModAtomExch
	ModAtomAnd
	ModAtomOr
	ModAtomXor
	ModCas
	ModExch

	// Arithmetic modifiers
	ModRelu
	ModCC
	ModClamp
	ModWrap
	ModShiftAmt
	ModNaN
	ModXorsign
	ModAbs
	ModOOB

	// Shift direction
	ModLeft
	ModRight

	// Prmt modes
	ModF4e
	ModB4e
	ModRc8
	ModEcl
	ModEcr
	ModRc16

	// L1 cache eviction
	ModL1EvictNormal
	ModL1EvictUnchanged
	ModL1EvictFirst
	ModL1EvictLast
	ModL1NoAllocate

	// L2 cache eviction & prefetch
	ModL2EvictNormal
	ModL2EvictFirst
	ModL2EvictLast
	ModL2Prefetch64B
	ModL2Prefetch128B
	ModL2Prefetch256B
	ModL2CacheHint
	ModLevelL2
	ModNC

	// Mbarrier & accumulation
	ModMbarrierCompleteTxBytes
	ModAccF32
	ModAccF16

	// Cvt / data movement
	ModTo
	ModTensormap
	ModRange
	ModFractional
	ModSatFinite
	ModScaledN2UE8M0

	// Async copy
	ModMulticastCluster
	ModBulkGroup
	ModCpMask

	// State spaces
	ModSpaceGlobal
	ModSpaceShared
	ModSpaceSharedCTA
	ModSpaceSharedCluster

	// Tensor dimensions
	ModDim1D
	ModDim2D
	ModDim3D
	ModDim4D
	ModDim5D

	// Tensor load modes
	ModLoadTile
	ModLoadTileGather4
	ModLoadTileScatter4
	ModLoadIm2Col
	ModLoadIm2ColW
	ModLoadIm2ColW128
	ModLoadIm2ColNoOffs

	// CTA groups
	ModCtaGroup1
	ModCtaGroup2

	// Texture geometries
	ModNoFtz
	ModGeom1D
	ModGeom2D
	ModGeom3D
	ModGeomA1D
	ModGeomA2D
	ModGeomCube
	ModGeomACube
	ModGeom2DMS
	ModGeomA2DMS

	// Texture mipmap modes
	ModBase
	ModLevel
	ModGrad

	// Texture components
	ModCompR
	ModCompG
	ModCompB
	ModCompA

	// Tensormap fields
	ModFieldGlobalAddr
	ModFieldRank
	ModFieldBoxDim
	ModFieldGlobalDim
	ModFieldGlobalStride
	ModFieldElementStride
	ModFieldElemType
	ModFieldInterleave
	ModFieldSwizzleMode
	ModFieldSwizzleAtom
	ModFieldFillMode
	ModRead

	// Texture & surface query attributes
	ModQueryWidth
	ModQueryHeight
	ModQueryDepth
	ModQueryChannelDataType
	ModQueryChannelOrder
	ModQueryNormalizedCoords
	ModQueryForceUnnormCoords
	ModQueryFilterMode
	ModQueryAddrMode0
	ModQueryAddrMode1
	ModQueryAddrMode2
	ModQueryArraySize
	ModQueryNumMipmapLevels
	ModQueryNumSamples
	ModQueryMemoryLayout

	// Surface clamp modes
	ModClampTrap
	ModClampClamp
	ModClampZero

	// Type checks
	ModTypeTexRef
	ModTypeSamplerRef
	ModTypeSurfRef

	// Surface format
	ModB
	ModP

	// Barrier ops
	ModArrive
	ModWait
	ModAligned

	// Memory consistency restrictions
	ModOpRestrict
	ModSyncRestrict
	ModMbarrierInitRestrict

	// Vector widths
	ModV8

	// Barrier / completion
	ModParity
	ModNoComplete
	ModNoInc
	ModExpectTx

	// Tensormap & cluster
	ModTensormapGeneric
	ModMulticastClusterAll
	ModIsCanceled
	ModGetFirstCTAId

	// MMA block scaling (.kind)
	ModKindMxf8f6f4
	ModKindMxf4
	ModKindMxf4nvf4

	// MMA scale vector (.scale_vec)
	ModScaleVec1x
	ModScaleVec2x
	ModScaleVec4x

	// Dimension query
	ModDimX
	ModDimY
	ModDimZ

	// Matrix roles
	ModMatrixA
	ModMatrixB
	ModMatrixC
	ModMatrixD

	// Matrix layouts
	ModRow
	ModCol

	// WMMA/MMA shapes
	ModShapeM16N16K16
	ModShapeM8N32K16
	ModShapeM32N8K16
	ModShapeM16N16K8
	ModShapeM8N8K4
	ModShapeM8N8K32
	ModShapeM8N8K128
	ModShapeM16N8K4
	ModShapeM16N8K8
	ModShapeM16N8K16
	ModShapeM16N8K32
	ModShapeM16N8K64
	ModShapeM16N8K128
	ModShapeM16N8K256
	ModShapeM8N8
	ModShapeM16N16
	ModShapeM8N16
	ModShapeM16N8

	// Matrix operations
	ModPopc

	// Type modifiers
	ModTypeF16
	ModTypeF32
	ModTypeF64
	ModTypeBF16
	ModTypeTF32
	ModTypeS32
	ModTypeS8
	ModTypeU8
	ModTypeS4
	ModTypeU4
	ModTypeB1
	ModTypeB32
	ModTypeB64
	ModTypeU32
	ModTypeS64
	ModTypeU64

	// Matrix counts (.num)
	ModNumX1
	ModNumX2
	ModNumX4

	// Matrix data formats
	ModDstFmtB8x16
	ModSrcFmtB6x16P32
	ModSrcFmtB4x16P64

	// Transpose & block scale
	ModTrans
	ModBlockScale

	// Sparse MMA
	ModSp
	ModSpOrderedMetadata

	// M64 K8 shapes
	ModShapeM64N8K8
	ModShapeM64N16K8
	ModShapeM64N24K8
	ModShapeM64N32K8
	ModShapeM64N40K8
	ModShapeM64N48K8
	ModShapeM64N56K8
	ModShapeM64N64K8
	ModShapeM64N72K8
	ModShapeM64N80K8
	ModShapeM64N88K8
	ModShapeM64N96K8
	ModShapeM64N104K8
	ModShapeM64N112K8
	ModShapeM64N120K8
	ModShapeM64N128K8
	ModShapeM64N136K8
	ModShapeM64N144K8
	ModShapeM64N152K8
	ModShapeM64N160K8
	ModShapeM64N168K8
	ModShapeM64N176K8
	ModShapeM64N184K8
	ModShapeM64N192K8
	ModShapeM64N200K8
	ModShapeM64N208K8
	ModShapeM64N216K8
	ModShapeM64N224K8
	ModShapeM64N232K8
	ModShapeM64N240K8
	ModShapeM64N248K8
	ModShapeM64N256K8

	// M64 K16 shapes
	ModShapeM64N8K16
	ModShapeM64N16K16
	ModShapeM64N24K16
	ModShapeM64N32K16
	ModShapeM64N40K16
	ModShapeM64N48K16
	ModShapeM64N56K16
	ModShapeM64N64K16
	ModShapeM64N72K16
	ModShapeM64N80K16
	ModShapeM64N88K16
	ModShapeM64N96K16
	ModShapeM64N104K16
	ModShapeM64N112K16
	ModShapeM64N120K16
	ModShapeM64N128K16
	ModShapeM64N136K16
	ModShapeM64N144K16
	ModShapeM64N152K16
	ModShapeM64N160K16
	ModShapeM64N168K16
	ModShapeM64N176K16
	ModShapeM64N184K16
	ModShapeM64N192K16
	ModShapeM64N200K16
	ModShapeM64N208K16
	ModShapeM64N216K16
	ModShapeM64N224K16
	ModShapeM64N232K16
	ModShapeM64N240K16
	ModShapeM64N248K16
	ModShapeM64N256K16

	// M64 K32 shapes
	ModShapeM64N8K32
	ModShapeM64N16K32
	ModShapeM64N24K32
	ModShapeM64N32K32
	ModShapeM64N40K32
	ModShapeM64N48K32
	ModShapeM64N56K32
	ModShapeM64N64K32
	ModShapeM64N72K32
	ModShapeM64N80K32
	ModShapeM64N88K32
	ModShapeM64N96K32
	ModShapeM64N104K32
	ModShapeM64N112K32
	ModShapeM64N120K32
	ModShapeM64N128K32
	ModShapeM64N136K32
	ModShapeM64N144K32
	ModShapeM64N152K32
	ModShapeM64N160K32
	ModShapeM64N168K32
	ModShapeM64N176K32
	ModShapeM64N184K32
	ModShapeM64N192K32
	ModShapeM64N208K32
	ModShapeM64N224K32
	ModShapeM64N240K32
	ModShapeM64N256K32

	// M64 K64 shapes
	ModShapeM64N8K64
	ModShapeM64N16K64
	ModShapeM64N24K64
	ModShapeM64N32K64
	ModShapeM64N40K64
	ModShapeM64N48K64
	ModShapeM64N56K64
	ModShapeM64N64K64
	ModShapeM64N72K64
	ModShapeM64N80K64
	ModShapeM64N88K64
	ModShapeM64N96K64
	ModShapeM64N104K64
	ModShapeM64N112K64
	ModShapeM64N120K64
	ModShapeM64N128K64
	ModShapeM64N136K64
	ModShapeM64N144K64
	ModShapeM64N152K64
	ModShapeM64N160K64
	ModShapeM64N168K64
	ModShapeM64N176K64
	ModShapeM64N184K64
	ModShapeM64N192K64
	ModShapeM64N200K64
	ModShapeM64N208K64
	ModShapeM64N216K64
	ModShapeM64N224K64
	ModShapeM64N232K64
	ModShapeM64N240K64
	ModShapeM64N248K64
	ModShapeM64N256K64

	// M64 K256 shapes
	ModShapeM64N8K256
	ModShapeM64N16K256
	ModShapeM64N24K256
	ModShapeM64N32K256
	ModShapeM64N48K256
	ModShapeM64N64K256
	ModShapeM64N80K256
	ModShapeM64N96K256
	ModShapeM64N112K256
	ModShapeM64N128K256
	ModShapeM64N144K256
	ModShapeM64N160K256
	ModShapeM64N176K256
	ModShapeM64N192K256
	ModShapeM64N208K256
	ModShapeM64N224K256
	ModShapeM64N240K256
	ModShapeM64N256K256

	// Tcgen05 data movement shapes (ld/st)
	ModShape16x64b
	ModShape16x128b
	ModShape16x256b
	ModShape16x32bx2
	ModShape32x32b

	// Tcgen05 data movement shapes (cp)
	ModShape4x256b
	ModShape32x128b
	ModShape64x128b
	ModShape128x256b
	ModShape128x128b

	// Tcgen05 shift shape
	ModShape31x256b

	// Tcgen05 MMA kinds
	ModKindF16
	ModKindTf32
	ModKindF8f6f4
	ModKindI8

	// Tcgen05 wait operations
	ModWaitLd
	ModWaitSt

	// Swizzle modes
	ModSwizzle32B
	ModSwizzle64B
	ModSwizzle128B

	// Tcgen05 reduction ops (ld.red)
	ModRedMin
	ModRedMax

	// Tcgen05 pack/unpack
	ModPack16b
	ModUnpack16b

	// Tcgen05 copy multicast
	ModMulticastWarpX2_02_13
	ModMulticastWarpX2_01_23
	ModMulticastWarpX4

	// Tcgen05 shift direction
	ModShiftDown

	// Tcgen05 block scaling aliases
	ModBlock16
	ModBlock32

	// Tcgen05 MMA modifiers
	ModWS
	ModAShift

	// Tcgen05 collector usage
	ModCollector
	ModBufA
	ModBufB0
	ModBufB1
	ModBufB2
	ModBufB3
	ModOpFill
	ModOpUse
	ModOpLastUse
	ModOpDiscard

	// Tcgen05 fence synchronization
	ModBeforeThreadSync
	ModAfterThreadSync

	// Video/SIMD selectors & masks
	ModB0
	ModB1
	ModB2
	ModB3
	ModH0
	ModH1
	ModH10
	ModB00
	ModB10
	ModB3210
	ModB7654

	// Video scaling & modes
	ModShr7
	ModShr15
	ModPo

	// Mbarrier completion
	ModMbarrierArriveOne

	// SetMaxNReg actions
	ModInc
	ModDec

	// Pmevent
	ModMask

	// Miscellaneous
	ModRed
)
func (m Modifier) String() string

type Opcode int
    Opcode represents a PTX instruction opcode.

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
	OpPmevent
	OpSetMaxNReg

	// Cluster launch control
	OpTensormapCpFenceproxy
	OpClusterlaunchcontrolTryCancel
	OpClusterlaunchcontrolQueryCancel

	// WGMMA
	OpWgmmaFence
	OpWgmmaCommitGroup
	OpWgmmaWaitGroup
	OpWgmmaMmaAsync
	OpFenceProxyAsync

	// Tensor Core Gen 5 (sm_100+)
	OpTcgen05Alloc
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
func (o Opcode) String() string

type RoundingMode int
    RoundingMode represents PTX rounding modifiers.

const (
	RoundNone RoundingMode = iota

	// Floating-point rounding
	RoundNearestEven // .rn — round to nearest even (default)
	RoundNearestAway // .rna — round to nearest, ties away from zero
	RoundZero        // .rz — round towards zero
	RoundNegInf      // .rm — round towards negative infinity
	RoundPosInf      // .rp — round towards positive infinity
	RoundStochastic  // .rs — stochastic rounding (sm_70+)

	// Integer rounding (for cvt float→int)
	RoundIntNearestEven // .rni
	RoundIntZero        // .rzi
	RoundIntNegInf      // .rmi
	RoundIntPosInf      // .rpi
)
func (r RoundingMode) String() string

type Scope int
    Scope represents the memory ordering scope for atomics, fences, and ld/st.

const (
	ScopeNone    Scope = iota
	ScopeCTA           // .cta  — within the CTA
	ScopeCluster       // .cluster — within the cluster (sm_90+)
	ScopeGPU           // .gpu  — within the GPU device
	ScopeSystem        // .sys  — across host + all devices
)
func (s Scope) String() string

type SpecialReg int
    SpecialReg represents PTX predefined special registers.

const (
	// ---- Thread identification ----
	RegTidX SpecialReg = iota // %tid.x
	RegTidY                   // %tid.y
	RegTidZ                   // %tid.z

	// ---- CTA (block) dimensions ----
	RegNTidX // %ntid.x
	RegNTidY // %ntid.y
	RegNTidZ // %ntid.z

	// ---- Warp identification ----
	RegLaneId  // %laneid
	RegWarpId  // %warpid
	RegNWarpId // %nwarpid

	// ---- CTA (block) identification ----
	RegCTAIdX // %ctaid.x
	RegCTAIdY // %ctaid.y
	RegCTAIdZ // %ctaid.z

	// ---- Grid dimensions ----
	RegNCTAIdX // %nctaid.x
	RegNCTAIdY // %nctaid.y
	RegNCTAIdZ // %nctaid.z

	// ---- SM identification ----
	RegSMId  // %smid
	RegNSMId // %nsmid

	// ---- Grid identification ----
	RegGridId // %gridid

	// ---- Cluster identification (sm_90+) ----
	RegIsExplicitCluster // %is_explicit_cluster
	RegClusterIdX        // %clusterid.x
	RegClusterIdY        // %clusterid.y
	RegClusterIdZ        // %clusterid.z
	RegNClusterIdX       // %nclusterid.x
	RegNClusterIdY       // %nclusterid.y
	RegNClusterIdZ       // %nclusterid.z
	RegClusterCTAIdX     // %cluster_ctaid.x
	RegClusterCTAIdY     // %cluster_ctaid.y
	RegClusterCTAIdZ     // %cluster_ctaid.z
	RegClusterNCTAIdX    // %cluster_nctaid.x
	RegClusterNCTAIdY    // %cluster_nctaid.y
	RegClusterNCTAIdZ    // %cluster_nctaid.z
	RegClusterCTARank    // %cluster_ctarank
	RegClusterNCTARank   // %cluster_nctarank

	// ---- Lane masks ----
	RegLanemaskEq // %lanemask_eq
	RegLanemaskLe // %lanemask_le
	RegLanemaskLt // %lanemask_lt
	RegLanemaskGe // %lanemask_ge
	RegLanemaskGt // %lanemask_gt

	// ---- Clock ----
	RegClock   // %clock
	RegClockHi // %clock_hi
	RegClock64 // %clock64

	// ---- Performance monitoring (32-bit) ----
	RegPM0 // %pm0
	RegPM1 // %pm1
	RegPM2 // %pm2
	RegPM3 // %pm3
	RegPM4 // %pm4
	RegPM5 // %pm5
	RegPM6 // %pm6
	RegPM7 // %pm7

	// ---- Performance monitoring (64-bit) ----
	RegPM0_64 // %pm0_64
	RegPM1_64 // %pm1_64
	RegPM2_64 // %pm2_64
	RegPM3_64 // %pm3_64
	RegPM4_64 // %pm4_64
	RegPM5_64 // %pm5_64
	RegPM6_64 // %pm6_64
	RegPM7_64 // %pm7_64

	// ---- Shared memory ----
	RegDynamicSmemSize         // %dynamic_smem_size
	RegTotalSmemSize           // %total_smem_size
	RegAggrSmemSize            // %aggr_smem_size
	RegReservedSmemOffsetBegin // %reserved_smem_offset_begin
	RegReservedSmemOffsetEnd   // %reserved_smem_offset_end
	RegReservedSmemOffsetCap   // %reserved_smem_offset_cap
	RegReservedSmemOffset2     // %reserved_smem_offset_2

	// ---- Global timer ----
	RegGlobalTimer   // %globaltimer
	RegGlobalTimerLo // %globaltimer_lo
	RegGlobalTimerHi // %globaltimer_hi

	// ---- Execution graph ----
	RegCurrentGraphExec // %current_graph_exec

	// ---- Environment registers (%envreg0–%envreg31) ----
	// RegEnvReg0 is the base; specific indices are computed as RegEnvReg0 + n.
	RegEnvReg0
	RegEnvReg31 = RegEnvReg0 + 31
)
func (r SpecialReg) String() string

func (r SpecialReg) Type() Type
    Type returns the PTX type of this special register.

type StateSpace int
    StateSpace represents PTX memory state spaces.

const (
	// Reg — fast per-thread registers.
	Reg StateSpace = iota

	// SReg — predefined read-only special registers.
	SReg

	// Const — read-only memory initialized by host.
	Const

	// Global — visible to all threads across all CTAs.
	Global

	// Local — per-thread private memory.
	Local

	// Param — kernel or function parameters (generic .param).
	Param

	// ParamEntry — explicitly kernel (.param::entry) parameters.
	ParamEntry

	// ParamFunc — explicitly device function (.param::func) parameters.
	ParamFunc

	// Shared — per-CTA memory visible to all threads in the CTA (generic .shared).
	Shared

	// SharedCTA — explicitly (.shared::cta) memory.
	SharedCTA

	// SharedCluster — (.shared::cluster) memory visible to all threads in the cluster (sm_90+).
	SharedCluster

	// Tex — deprecated texture state space (legacy).
	Tex
)
func (s StateSpace) String() string

type Target int
    Target represents a PTX target architecture (sm_XX).

const (
	SM50 Target = iota // Maxwell
	SM52
	SM53
	SM60 // Pascal
	SM61
	SM62
	SM70 // Volta
	SM72
	SM75 // Turing
	SM80 // Ampere
	SM86
	SM87
	SM89 // Ada Lovelace
	SM90 // Hopper
	SM90a
	SM100 // Blackwell
	SM101
	SM120
)
func (t Target) String() string

type Type int
    Type represents a PTX fundamental data type.

const (
	// Predicate
	Pred Type = iota

	// Bit-size (untyped)
	B8
	B16
	B32
	B64
	B128

	// Signed integer
	S8
	S16
	S32
	S64

	// Unsigned integer
	U8
	U16
	U32
	U64

	// Floating point
	F16
	F32
	F64

	// Alternate floating point (sm_80+)
	BF16
	TF32

	// Packed types (Legacy/Common)
	F16x2
	BF16x2

	// Fixed-point (sm_100+)
	E2M1
	E2M3
	E3M2
	E4M3
	E5M2
	E8M0

	// Additional Scalar types
	UE4M3 // .ue4m3 (7-bit unsigned float, stored in .b8)
	S2F6  // .s2f6 (8-bit signed fixed-point)

	// Packed Floating Point & Fixed Point (Table 9)
	F32x2   // .f32x2
	E4M3x2  // .e4m3x2
	E5M2x2  // .e5m2x2
	E2M3x2  // .e2m3x2
	E3M2x2  // .e3m2x2
	UE8M0x2 // .ue8m0x2
	S2F6x2  // .s2f6x2
	E2M1x2  // .e2m1x2
	E4M3x4  // .e4m3x4
	E5M2x4  // .e5m2x4
	E2M3x4  // .e2m3x4
	E3M2x4  // .e3m2x4
	E2M1x4  // .e2m1x4

	// Packed Integer
	U16x2 // .u16x2
	S16x2 // .s16x2

	// Tensor Sub-byte types (Section 5.5.1.1)
	B4x16     // .b4x16
	B4x16_p64 // .b4x16_p64
	B6x16_p32 // .b6x16_p32
	B6p2x16   // .b6p2x16

	// Opaque Types (Section 5.3 & 5.5.8)
	TexRef     // .texref
	SamplerRef // .samplerref
	SurfRef    // .surfref
	TensorMap  // .tensormap
)
func (t Type) BitWidth() int
    BitWidth returns the width in bits for the type. For packed types, this is
    the total width of the container. For opaque types, this returns the typical
    handle/structure size.

func (t Type) IsFloat() bool
    IsFloat returns true for floating-point types (including packed and
    alternate floats).

func (t Type) IsSigned() bool
    IsSigned returns true for signed integer types and signed fixed-point types.

func (t Type) String() string
    String returns the PTX type string (e.g. ".f32", ".u64", ".pred").

type VectorSize int
    VectorSize represents PTX vector widths.

const (
	Scalar VectorSize = iota
	V2                // .v2
	V4                // .v4
)
func (v VectorSize) String() string


## github.com/arc-language/ptx-gen/ptxgen

ptxgen.go

FUNCTIONS

func Build(mod *builder.Module) string
func NewModule(version ptx.ISAVersion, target ptx.Target) *builder.Module

