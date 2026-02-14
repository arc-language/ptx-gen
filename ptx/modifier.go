package ptx

// Modifier represents miscellaneous PTX instruction modifiers.
type Modifier int

const (
	// ---- Multiply width ----
	ModWide Modifier = iota
	ModLo
	ModHi

	// ---- Saturation & flush ----
	ModSat
	ModFtz

	// ---- Approximation ----
	ModApprox
	ModFull

	// ---- Uniformity ----
	ModUni

	// ---- Memory consistency ----
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

	// ---- Sync ----
	ModSync

	// ---- Testp ----
	ModFinite
	ModInfinite
	ModNumber
	ModNotANumber
	ModNormal
	ModSubnormal

	// ---- Shfl modes ----
	ModShflUp
	ModShflDown
	ModShflBfly
	ModShflIdx

	// ---- Atomic operations ----
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

	// ---- Arithmetic modifiers ----
	ModRelu
	ModCC
	ModClamp
	ModWrap
	ModShiftAmt
	ModNaN
	ModXorsign
	ModAbs
	ModOOB

	// ---- Shift direction ----
	ModLeft
	ModRight

	// ---- Prmt modes ----
	ModF4e
	ModB4e
	ModRc8
	ModEcl
	ModEcr
	ModRc16

	// ---- L1 cache eviction ----
	ModL1EvictNormal
	ModL1EvictUnchanged
	ModL1EvictFirst
	ModL1EvictLast
	ModL1NoAllocate

	// ---- L2 cache eviction & prefetch ----
	ModL2EvictNormal
	ModL2EvictFirst
	ModL2EvictLast
	ModL2Prefetch64B
	ModL2Prefetch128B
	ModL2Prefetch256B
	ModL2CacheHint
	ModLevelL2
	ModNC

	// ---- Mbarrier & accumulation ----
	ModMbarrierCompleteTxBytes
	ModAccF32
	ModAccF16

	// ---- Cvt / data movement ----
	ModTo
	ModTensormap
	ModRange
	ModFractional
	ModSatFinite
	ModScaledN2UE8M0

	// ---- Async copy ----
	ModMulticastCluster
	ModBulkGroup
	ModCpMask

	// ---- State spaces (used as modifiers) ----
	ModSpaceGlobal
	ModSpaceShared
	ModSpaceSharedCTA
	ModSpaceSharedCluster

	// ---- Tensor dimensions ----
	ModDim1D
	ModDim2D
	ModDim3D
	ModDim4D
	ModDim5D

	// ---- Tensor load modes ----
	ModLoadTile
	ModLoadTileGather4
	ModLoadTileScatter4
	ModLoadIm2Col
	ModLoadIm2ColW
	ModLoadIm2ColW128
	ModLoadIm2ColNoOffs

	// ---- CTA groups ----
	ModCtaGroup1
	ModCtaGroup2

	// ---- Texture geometries ----
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

	// ---- Texture mipmap modes ----
	ModBase
	ModLevel
	ModGrad

	// ---- Texture components ----
	ModCompR
	ModCompG
	ModCompB
	ModCompA

	// ---- Tensormap fields ----
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

	// ---- Texture & surface query attributes ----
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

	// ---- Surface clamp modes ----
	ModClampTrap
	ModClampClamp
	ModClampZero

	// ---- Type checks ----
	ModTypeTexRef
	ModTypeSamplerRef
	ModTypeSurfRef

	// ---- Surface format ----
	ModB
	ModP

	// ---- Barrier ops ----
	ModArrive
	ModWait
	ModAligned

	// ---- Memory consistency restrictions ----
	ModOpRestrict
	ModSyncRestrict
	ModMbarrierInitRestrict

	// ---- Vector widths ----
	ModV8
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
	case ModLevelL2 + 1: // ModQueryWidth — guarded, see note below
		return ".L2" // unreachable; ModLevelL2 already handled
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
	default:
		return ""
	}
}