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
	default:
		return ""
	}
}