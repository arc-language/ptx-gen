package ptx

// Modifier represents miscellaneous PTX instruction modifiers.
type Modifier int

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
	ModShapeM64N256K16

	// M64 K8 shapes
	ModShapeM64N8K8
	ModShapeM64N16K8
	ModShapeM64N32K8

	// M64 K32 shapes
	ModShapeM64N8K32
	ModShapeM64N16K32
	ModShapeM64N32K32

	// M64 K64 shapes
	ModShapeM64N8K64

	// M64 K256 shapes
	ModShapeM64N8K256


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

    // M64 N* K8 (tf32)
    ModShapeM64N24K8
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

    // M64 N* K32 (fp8, int)
    ModShapeM64N24K32
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

    // M64 N* K256 (b1)
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



	// --- Tcgen05 Data Movement Shapes (ld/st) ---
    ModShape16x64b
    ModShape16x128b
    ModShape16x256b
    ModShape16x32bx2
    ModShape32x32b

    // --- Tcgen05 Data Movement Shapes (cp) ---
    ModShape4x256b
    ModShape32x128b
    ModShape64x128b
    ModShape128x256b
    ModShape128x128b
    
    // --- Tcgen05 Shift Shape ---
    ModShape31x256b

    // --- Tcgen05 MMA Kinds ---
    ModKindF16      // .kind::f16
    ModKindTf32     // .kind::tf32
    ModKindF8f6f4   // .kind::f8f6f4
    ModKindI8       // .kind::i8

    // --- Tcgen05 Wait Operations ---
    ModWaitLd // .wait::ld
    ModWaitSt // .wait::st

    // --- Swizzle Modes (for Descriptors) ---
    ModSwizzle32B  // 32-Byte swizzle
    ModSwizzle64B  // 64-Byte swizzle
    ModSwizzle128B // 128-Byte swizzle


	// --- Tcgen05 Reduction Ops (ld.red) ---
    ModRedMin
    ModRedMax

    // --- Tcgen05 Pack/Unpack ---
    ModPack16b   // .pack::16b
    ModUnpack16b // .unpack::16b

    // --- Tcgen05 Copy Multicast ---
    ModMulticastWarpX2_02_13 // .warpx2::02_13
    ModMulticastWarpX2_01_23 // .warpx2::01_23
    ModMulticastWarpX4       // .warpx4

    // --- Tcgen05 Shift Direction ---
    ModShiftDown // .down

	// --- Tcgen05 Block Scaling Aliases ---
    ModBlock16      // .block16 (alias for .scale_vec::4X)
    ModBlock32      // .block32 (alias for .scale_vec::1X or .scale_vec::2X)


// --- Tcgen05 MMA Modifiers ---
    ModWS         // .ws

    ModAShift     // .ashift



    // --- Tcgen05 Collector Usage ---
    ModCollector // .collector
    ModBufA      // ::a
    ModBufB0     // ::b0
    ModBufB1     // ::b1
    ModBufB2     // ::b2
    ModBufB3     // ::b3
    ModOpFill    // ::fill
    ModOpUse     // ::use
    ModOpLastUse // ::lastuse
    ModOpDiscard // ::discard

    // --- Tcgen05 Fence Synchronization ---
    ModBeforeThreadSync // ::before_thread_sync
    ModAfterThreadSync  // ::after_thread_sync



)

func (m Modifier) String() string {
	switch m {
	case ModWide:
		return ".wide"
	case ModLo:
		return ".lo"
	case ModHi:
		return ".hi"
	case ModSat:
		return ".sat"
	case ModFtz:
		return ".ftz"
	case ModApprox:
		return ".approx"
	case ModFull:
		return ".full"
	case ModUni:
		return ".uni"
	case ModAcquire:
		return ".acquire"
	case ModRelease:
		return ".release"
	case ModRelaxed:
		return ".relaxed"
	case ModAcqRel:
		return ".acq_rel"
	case ModVolatile:
		return ".volatile"
	case ModMMIO:
		return ".mmio"
	case ModWeak:
		return ".weak"
	case ModSC:
		return ".sc"
	case ModProxy:
		return ".proxy"
	case ModAlias:
		return ".alias"
	case ModAsync:
		return ".async"
	case ModSync:
		return ".sync"
	case ModFinite:
		return ".finite"
	case ModInfinite:
		return ".infinite"
	case ModNumber:
		return ".number"
	case ModNotANumber:
		return ".notanumber"
	case ModNormal:
		return ".normal"
	case ModSubnormal:
		return ".subnormal"
	case ModShflUp:
		return ".up"
	case ModShflDown:
		return ".down"
	case ModShflBfly:
		return ".bfly"
	case ModShflIdx:
		return ".idx"
	case ModAtomAdd:
		return ".add"
	case ModAtomMin:
		return ".min"
	case ModAtomMax:
		return ".max"
	case ModAtomInc:
		return ".inc"
	case ModAtomDec:
		return ".dec"
	case ModAtomCAS:
		return ".cas"
	case ModAtomExch:
		return ".exch"
	case ModAtomAnd:
		return ".and"
	case ModAtomOr:
		return ".or"
	case ModAtomXor:
		return ".xor"
	case ModCas:
		return ".cas"
	case ModExch:
		return ".exch"
	case ModRelu:
		return ".relu"
	case ModCC:
		return ".cc"
	case ModClamp:
		return ".clamp"
	case ModWrap:
		return ".wrap"
	case ModShiftAmt:
		return ".shiftamt"
	case ModNaN:
		return ".NaN"
	case ModXorsign:
		return ".xorsign"
	case ModAbs:
		return ".abs"
	case ModOOB:
		return ".oob"
	case ModLeft:
		return ".l"
	case ModRight:
		return ".r"
	case ModF4e:
		return ".f4e"
	case ModB4e:
		return ".b4e"
	case ModRc8:
		return ".rc8"
	case ModEcl:
		return ".ecl"
	case ModEcr:
		return ".ecr"
	case ModRc16:
		return ".rc16"
	case ModL1EvictNormal:
		return ".L1::evict_normal"
	case ModL1EvictUnchanged:
		return ".L1::evict_unchanged"
	case ModL1EvictFirst:
		return ".L1::evict_first"
	case ModL1EvictLast:
		return ".L1::evict_last"
	case ModL1NoAllocate:
		return ".L1::no_allocate"
	case ModL2EvictNormal:
		return ".L2::evict_normal"
	case ModL2EvictFirst:
		return ".L2::evict_first"
	case ModL2EvictLast:
		return ".L2::evict_last"
	case ModL2Prefetch64B:
		return ".L2::64B"
	case ModL2Prefetch128B:
		return ".L2::128B"
	case ModL2Prefetch256B:
		return ".L2::256B"
	case ModL2CacheHint:
		return ".L2::cache_hint"
	case ModLevelL2:
		return ".L2"
	case ModNC:
		return ".nc"
	case ModMbarrierCompleteTxBytes:
		return ".mbarrier::complete_tx::bytes"
	case ModAccF32:
		return ".acc::f32"
	case ModAccF16:
		return ".acc::f16"
	case ModTo:
		return ".to"
	case ModTensormap:
		return ".tensormap"
	case ModRange:
		return ".range"
	case ModFractional:
		return ".fractional"
	case ModSatFinite:
		return ".satfinite"
	case ModScaledN2UE8M0:
		return ".scaled::n2::ue8m0"
	case ModMulticastCluster:
		return ".multicast::cluster"
	case ModBulkGroup:
		return ".bulk_group"
	case ModCpMask:
		return ".cp_mask"
	case ModSpaceGlobal:
		return ".global"
	case ModSpaceShared:
		return ".shared"
	case ModSpaceSharedCTA:
		return ".shared::cta"
	case ModSpaceSharedCluster:
		return ".shared::cluster"
	case ModDim1D:
		return ".1d"
	case ModDim2D:
		return ".2d"
	case ModDim3D:
		return ".3d"
	case ModDim4D:
		return ".4d"
	case ModDim5D:
		return ".5d"
	case ModLoadTile:
		return ".tile"
	case ModLoadTileGather4:
		return ".tile::gather4"
	case ModLoadTileScatter4:
		return ".tile::scatter4"
	case ModLoadIm2Col:
		return ".im2col"
	case ModLoadIm2ColW:
		return ".im2col::w"
	case ModLoadIm2ColW128:
		return ".im2col::w::128"
	case ModLoadIm2ColNoOffs:
		return ".im2col_no_offs"
	case ModCtaGroup1:
		return ".cta_group::1"
	case ModCtaGroup2:
		return ".cta_group::2"
	case ModNoFtz:
		return ".noftz"
	case ModGeom1D:
		return ".1d"
	case ModGeom2D:
		return ".2d"
	case ModGeom3D:
		return ".3d"
	case ModGeomA1D:
		return ".a1d"
	case ModGeomA2D:
		return ".a2d"
	case ModGeomCube:
		return ".cube"
	case ModGeomACube:
		return ".acube"
	case ModGeom2DMS:
		return ".2dms"
	case ModGeomA2DMS:
		return ".a2dms"
	case ModBase:
		return ".base"
	case ModLevel:
		return ".level"
	case ModGrad:
		return ".grad"
	case ModCompR:
		return ".r"
	case ModCompG:
		return ".g"
	case ModCompB:
		return ".b"
	case ModCompA:
		return ".a"
	case ModFieldGlobalAddr:
		return ".global_address"
	case ModFieldRank:
		return ".rank"
	case ModFieldBoxDim:
		return ".box_dim"
	case ModFieldGlobalDim:
		return ".global_dim"
	case ModFieldGlobalStride:
		return ".global_stride"
	case ModFieldElementStride:
		return ".element_stride"
	case ModFieldElemType:
		return ".elemtype"
	case ModFieldInterleave:
		return ".interleave_layout"
	case ModFieldSwizzleMode:
		return ".swizzle_mode"
	case ModFieldSwizzleAtom:
		return ".swizzle_atomicity"
	case ModFieldFillMode:
		return ".fill_mode"
	case ModRead:
		return ".read"
	case ModQueryWidth:
		return ".width"
	case ModQueryHeight:
		return ".height"
	case ModQueryDepth:
		return ".depth"
	case ModQueryChannelDataType:
		return ".channel_data_type"
	case ModQueryChannelOrder:
		return ".channel_order"
	case ModQueryNormalizedCoords:
		return ".normalized_coords"
	case ModQueryForceUnnormCoords:
		return ".force_unnormalized_coords"
	case ModQueryFilterMode:
		return ".filter_mode"
	case ModQueryAddrMode0:
		return ".addr_mode_0"
	case ModQueryAddrMode1:
		return ".addr_mode_1"
	case ModQueryAddrMode2:
		return ".addr_mode_2"
	case ModQueryArraySize:
		return ".array_size"
	case ModQueryNumMipmapLevels:
		return ".num_mipmap_levels"
	case ModQueryNumSamples:
		return ".num_samples"
	case ModQueryMemoryLayout:
		return ".memory_layout"
	case ModClampTrap:
		return ".trap"
	case ModClampClamp:
		return ".clamp"
	case ModClampZero:
		return ".zero"
	case ModTypeTexRef:
		return ".texref"
	case ModTypeSamplerRef:
		return ".samplerref"
	case ModTypeSurfRef:
		return ".surfref"
	case ModB:
		return ".b"
	case ModP:
		return ".p"
	case ModArrive:
		return ".arrive"
	case ModWait:
		return ".wait"
	case ModAligned:
		return ".aligned"
	case ModOpRestrict:
		return ".op_restrict"
	case ModSyncRestrict:
		return ".sync_restrict"
	case ModMbarrierInitRestrict:
		return ".mbarrier_init"
	case ModV8:
		return ".v8"
	case ModParity:
		return ".parity"
	case ModNoComplete:
		return ".noComplete"
	case ModNoInc:
		return ".noinc"
	case ModExpectTx:
		return ".expect_tx"
	case ModTensormapGeneric:
		return ".tensormap::generic"
	case ModMulticastClusterAll:
		return ".multicast::cluster::all"
	case ModIsCanceled:
		return ".is_canceled"
	case ModGetFirstCTAId:
		return ".get_first_ctaid"
	case ModKindMxf8f6f4:
		return ".kind::mxf8f6f4"
	case ModKindMxf4:
		return ".kind::mxf4"
	case ModKindMxf4nvf4:
		return ".kind::mxf4nvf4"
	case ModScaleVec1x:
		return ".scale_vec::1X"
	case ModScaleVec2x:
		return ".scale_vec::2X"
	case ModScaleVec4x:
		return ".scale_vec::4X"
	case ModDimX:
		return "::x"
	case ModDimY:
		return "::y"
	case ModDimZ:
		return "::z"
	case ModMatrixA:
		return ".a"
	case ModMatrixB:
		return ".b"
	case ModMatrixC:
		return ".c"
	case ModMatrixD:
		return ".d"
	case ModRow:
		return ".row"
	case ModCol:
		return ".col"
	case ModShapeM16N16K16:
		return ".m16n16k16"
	case ModShapeM8N32K16:
		return ".m8n32k16"
	case ModShapeM32N8K16:
		return ".m32n8k16"
	case ModShapeM16N16K8:
		return ".m16n16k8"
	case ModShapeM8N8K4:
		return ".m8n8k4"
	case ModShapeM8N8K32:
		return ".m8n8k32"
	case ModShapeM8N8K128:
		return ".m8n8k128"
	case ModShapeM16N8K4:
		return ".m16n8k4"
	case ModShapeM16N8K8:
		return ".m16n8k8"
	case ModShapeM16N8K16:
		return ".m16n8k16"
	case ModShapeM16N8K32:
		return ".m16n8k32"
	case ModShapeM16N8K64:
		return ".m16n8k64"
	case ModShapeM16N8K128:
		return ".m16n8k128"
	case ModPopc:
		return ".popc"
	case ModTypeF16:
		return ".f16"
	case ModTypeF32:
		return ".f32"
	case ModTypeF64:
		return ".f64"
	case ModTypeBF16:
		return ".bf16"
	case ModTypeTF32:
		return ".tf32"
	case ModTypeS32:
		return ".s32"
	case ModTypeS8:
		return ".s8"
	case ModTypeU8:
		return ".u8"
	case ModTypeS4:
		return ".s4"
	case ModTypeU4:
		return ".u4"
	case ModTypeB1:
		return ".b1"
	case ModShapeM16N8K256:
		return ".m16n8k256"
	case ModShapeM8N8:
		return ".m8n8"
	case ModShapeM16N16:
		return ".m16n16"
	case ModShapeM8N16:
		return ".m8n16"
	case ModShapeM16N8:
		return ".m16n8"
	case ModNumX1:
		return ".x1"
	case ModNumX2:
		return ".x2"
	case ModNumX4:
		return ".x4"
	case ModDstFmtB8x16:
		return ".b8x16"
	case ModSrcFmtB6x16P32:
		return ".b6x16_p32"
	case ModSrcFmtB4x16P64:
		return ".b4x16_p64"
	case ModTrans:
		return ".trans"
	case ModBlockScale:
		return ".block_scale"
	case ModSp:
		return ".sp"
	case ModSpOrderedMetadata:
		return ".sp::ordered_metadata"
	case ModShapeM64N8K16:
		return ".m64n8k16"
	case ModShapeM64N16K16:
		return ".m64n16k16"
	case ModShapeM64N32K16:
		return ".m64n32k16"
	case ModShapeM64N64K16:
		return ".m64n64k16"
	case ModShapeM64N128K16:
		return ".m64n128k16"
	case ModShapeM64N256K16:
		return ".m64n256k16"
	case ModShapeM64N8K8:
		return ".m64n8k8"
	case ModShapeM64N32K8:
		return ".m64n32k8"
	case ModShapeM64N8K32:
		return ".m64n8k32"
	case ModShapeM64N32K32:
		return ".m64n32k32"



    // ... (map the rest of the K16 constants similarly)

    case ModShapeM64N16K8: return ".m64n16k8"
    // ... (map the rest of the K8 constants)

    case ModShapeM64N16K32: return ".m64n16k32"

    case ModShapeM64N8K256: return ".m64n8k256"
    // ... (map the rest of the K256 constants)



	case ModShapeM64N8K64: return ".m64n8k64"
    case ModShapeM64N16K64: return ".m64n16k64"
    case ModShapeM64N24K64: return ".m64n24k64"
    case ModShapeM64N32K64: return ".m64n32k64"
    case ModShapeM64N48K64: return ".m64n48k64"
    case ModShapeM64N64K64: return ".m64n64k64"
    case ModShapeM64N80K64: return ".m64n80k64"
    case ModShapeM64N96K64: return ".m64n96k64"
    case ModShapeM64N112K64: return ".m64n112k64"
    case ModShapeM64N128K64: return ".m64n128k64"
    case ModShapeM64N144K64: return ".m64n144k64"
    case ModShapeM64N160K64: return ".m64n160k64"
    case ModShapeM64N176K64: return ".m64n176k64"
    case ModShapeM64N192K64: return ".m64n192k64"
    case ModShapeM64N208K64: return ".m64n208k64"
    case ModShapeM64N224K64: return ".m64n224k64"
    case ModShapeM64N240K64: return ".m64n240k64"
    case ModShapeM64N256K64: return ".m64n256k64"


	case ModShape16x64b: return ".16x64b"
    case ModShape16x128b: return ".16x128b"
    case ModShape16x256b: return ".16x256b"
    case ModShape16x32bx2: return ".16x32bx2"
    case ModShape32x32b: return ".32x32b"
    case ModShape4x256b: return ".4x256b"
    case ModShape32x128b: return ".32x128b"
    case ModShape64x128b: return ".64x128b"
    case ModShape128x256b: return ".128x256b"
    case ModShape128x128b: return ".128x128b"
    case ModShape31x256b: return ".31x256b"

    case ModKindF16: return ".kind::f16"
    case ModKindTf32: return ".kind::tf32"
    case ModKindF8f6f4: return ".kind::f8f6f4"
    case ModKindI8: return ".kind::i8"

    case ModWaitLd: return "::ld"
    case ModWaitSt: return "::st"


	case ModRedMin: return ".min"
    case ModRedMax: return ".max"
    case ModPack16b: return ".pack::16b"
    case ModUnpack16b: return ".unpack::16b"
    case ModMulticastWarpX2_02_13: return ".warpx2::02_13"
    case ModMulticastWarpX2_01_23: return ".warpx2::01_23"
    case ModMulticastWarpX4: return ".warpx4"
    case ModShiftDown: return ".down"

	case ModBlock16: return ".block16"
    case ModBlock32: return ".block32"


	case ModWS: return ".ws"
    case ModAShift: return ".ashift"


    case ModCollector: return ".collector"
    case ModBufA: return "::a"
    case ModBufB0: return "::b0"
    case ModBufB1: return "::b1"
    case ModBufB2: return "::b2"
    case ModBufB3: return "::b3"
    case ModOpFill: return "::fill"
    case ModOpUse: return "::use"
    case ModOpLastUse: return "::lastuse"
    case ModOpDiscard: return "::discard"

    case ModBeforeThreadSync: return "::before_thread_sync"
    case ModAfterThreadSync: return "::after_thread_sync"


	default:
		return ""
	}
}