package ptx

// Modifier represents miscellaneous PTX instruction modifiers.
type Modifier int

const (
    // Multiply width
    ModWide Modifier = iota // .wide — double-width result
    ModLo                   // .lo   — keep low half
    ModHi                   // .hi   — keep high half

    // Saturation
    ModSat                  // .sat  — clamp to [0.0, 1.0] or type range

    // Flush to zero
    ModFtz                  // .ftz  — flush denorms to zero

    // Approximation
    ModApprox               // .approx — fast approximation (rcp, sqrt, rsqrt, sin, cos, lg2, ex2)
    ModFull                 // .full   — full-range approximation

    // Uniformity
    ModUni                  // .uni  — all threads in warp take same path

    // Atomics
    ModAcquire              // .acquire
    ModRelease              // .release
    ModRelaxed              // .relaxed
    ModAcqRel               // .acq_rel
    ModVolatile             // .volatile
    ModMMIO                 // .mmio

    // Sync variants
    ModSync                 // .sync

    // Testp modifiers
    ModFinite               // .finite
    ModInfinite             // .infinite
    ModNumber               // .number
    ModNotANumber           // .notanumber
    ModNormal               // .normal
    ModSubnormal            // .subnormal

    // Shfl modes
    ModShflUp               // .up
    ModShflDown             // .down
    ModShflBfly             // .bfly
    ModShflIdx              // .idx

    // Atomic operations
    ModAtomAdd              // .add
    ModAtomMin              // .min
    ModAtomMax              // .max
    ModAtomInc              // .inc
    ModAtomDec              // .dec
    ModAtomCAS              // .cas
    ModAtomExch             // .exch
    ModAtomAnd              // .and
    ModAtomOr               // .or
    ModAtomXor              // .xor

    // --- New Memory Consistency Modifiers ---
    ModWeak                 // .weak (Weak operations)
    ModSC                   // .sc (Sequentially consistent fence)
    ModProxy                // .proxy (Proxy fence/membar)
    ModAlias                // .alias (Alias proxy)
    ModAsync                // .async (Used in async proxy fences)

    ModRelu // .relu (clamp to 0, used in min.relu)

    ModCC                   // .cc (Write condition code)
    ModClamp                // .clamp (Clamp mode for szext, bmsk)
    ModWrap                 // .wrap (Wrap mode for szext, bmsk)
    ModShiftAmt             // .shiftamt (Return shift amount for bfind)

    ModNaN                  // .NaN (Propagate NaN)
    ModXorsign              // .xorsign (XOR sign bits)
    ModAbs                  // .abs (Absolute value modifier for min/max)

    // --- New Modifier for Half-Precision FMA (Section 9.7.4.4) ---
    ModOOB // .oob (Out of bounds)

    ModLeft  Modifier = iota // .l (for shf.l)
    ModRight                 // .r (for shf.r)


    // --- Prmt Modes (Section 9.7.9.7) ---
	ModF4e  Modifier = iota + 100 // .f4e (Forward 4 extract)
	ModB4e                        // .b4e (Backward 4 extract)
	ModRc8                        // .rc8 (Replicate 8)
	ModEcl                        // .ecl (Edge clamp left)
	ModEcr                        // .ecr (Edge clamp right)
	ModRc16                       // .rc16 (Replicate 16)

	// --- Cache Eviction Priorities (Section 9.7.9.2) ---
	// L1 Specific
	ModL1EvictNormal    // .L1::evict_normal
	ModL1EvictUnchanged // .L1::evict_unchanged
	ModL1EvictFirst     // .L1::evict_first
	ModL1EvictLast      // .L1::evict_last
	ModL1NoAllocate     // .L1::no_allocate

	// L2 Specific
	ModL2EvictNormal // .L2::evict_normal
	ModL2EvictFirst  // .L2::evict_first
	ModL2EvictLast   // .L2::evict_last

	// --- Prefetch Sizes (Section 9.7.9.8) ---
	ModL2Prefetch64B  // .L2::64B
	ModL2Prefetch128B // .L2::128B
	ModL2Prefetch256B // .L2::256B

	// --- Cache Hints ---
	ModL2CacheHint // .L2::cache_hint
	
	// --- Misc Load Modifiers ---
	ModNC // .nc (Non-coherent, used in ld.global.nc if not implicit in opcode)

    // Completion Mechanism (st.async)
    ModMbarrierCompleteTxBytes Modifier = iota + 200 // .mbarrier::complete_tx::bytes

    // Accumulation Precision (multimem)
    ModAccF32 // .acc::f32
    ModAccF16 // .acc::f16


    // --- Cvt / Data Movement Modifiers ---
	ModTo            Modifier = iota + 300 // .to (cvta.to)
	ModTensormap                           // .tensormap (prefetch)
	ModRange                               // .range (createpolicy)
	ModFractional                          // .fractional (createpolicy)
	ModSatFinite                           // .satfinite (cvt)
	ModScaledN2UE8M0                       // .scaled::n2::ue8m0 (cvt scaling)


    // --- Async Copy Modifiers ---
    ModMulticastCluster Modifier = iota + 400 // .multicast::cluster
    ModBulkGroup                              // .bulk_group
    ModCpMask                                 // .cp_mask
    
    // State Spaces used as modifiers in instructions (like cp.async)
    // We prefix them to avoid collision with generic StateSpace enum
    ModSpaceGlobal        // .global
    ModSpaceShared        // .shared
    ModSpaceSharedCTA     // .shared::cta

    ModSpaceSharedCluster Modifier = iota + 500 // Aligning



    // --- Dimensions (Section 9.7.9.27.1.2) ---
    ModDim1D // .1d
    ModDim2D // .2d
    ModDim3D // .3d
    ModDim4D // .4d
    ModDim5D // .5d

    // --- Tensor Load Modes ---
    ModLoadTile          // .tile
    ModLoadTileGather4   // .tile::gather4
    ModLoadTileScatter4  // .tile::scatter4
    ModLoadIm2Col        // .im2col
    ModLoadIm2ColW       // .im2col::w
    ModLoadIm2ColW128    // .im2col::w::128
    ModLoadIm2ColNoOffs  // .im2col_no_offs

    // --- CTA Groups ---
    ModCtaGroup1 // .cta_group::1
    ModCtaGroup2 // .cta_group::2

    // --- Misc ---
    ModNoFtz Modifier = iota + 550 // Aligning

	// --- Texture Geometries (Section 9.7.10.3) ---
	ModGeom1D    // .1d
	ModGeom2D    // .2d
	ModGeom3D    // .3d
	ModGeomA1D   // .a1d
	ModGeomA2D   // .a2d
	ModGeomCube  // .cube
	ModGeomACube // .acube
	ModGeom2DMS  // .2dms
	ModGeomA2DMS // .a2dms

	// --- Texture Mipmap Modes ---
	ModBase  // .base
	ModLevel // .level
	ModGrad  // .grad

	// --- Texture Components (tld4) ---
	ModCompR // .r
	ModCompG // .g
	ModCompB // .b
	ModCompA // .a

	// --- Tensormap Fields (tensormap.replace) ---
	ModFieldGlobalAddr      // .global_address
	ModFieldRank            // .rank
	ModFieldBoxDim          // .box_dim
	ModFieldGlobalDim       // .global_dim
	ModFieldGlobalStride    // .global_stride
	ModFieldElementStride   // .element_stride
	ModFieldElemType        // .elemtype
	ModFieldInterleave      // .interleave_layout
	ModFieldSwizzleMode     // .swizzle_mode
	ModFieldSwizzleAtom     // .swizzle_atomicity
	ModFieldFillMode        // .fill_mode

	// --- Misc ---
	ModRead // .read (for cp.async.bulk.wait_group)

    ModLevelL2 Modifier = iota + 600 // .L2
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
    // Prmt
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

	// L1 Eviction
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

	// L2 Eviction
	case ModL2EvictNormal:
		return ".L2::evict_normal"
	case ModL2EvictFirst:
		return ".L2::evict_first"
	case ModL2EvictLast:
		return ".L2::evict_last"

	// Prefetch
	case ModL2Prefetch64B:
		return ".L2::64B"
	case ModL2Prefetch128B:
		return ".L2::128B"
	case ModL2Prefetch256B:
		return ".L2::256B"

	// Hints
	case ModL2CacheHint:
		return ".L2::cache_hint"
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

    // Dimensions
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

    // Load Modes
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

    // CTA Groups
    case ModCtaGroup1:
        return ".cta_group::1"
    case ModCtaGroup2:
        return ".cta_group::2"

    // Misc
    case ModNoFtz:
        return ".noftz"

    // Texture Geometries
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

	// Mipmap
	case ModBase:
		return ".base"
	case ModLevel:
		return ".level"
	case ModGrad:
		return ".grad"

	// Components
	case ModCompR:
		return ".r"
	case ModCompG:
		return ".g"
	case ModCompB:
		return ".b"
	case ModCompA:
		return ".a"

	// Tensormap Fields
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
    case ModLevelL2:
        return ".L2"
	case ModRead:
		return ".read"

    default:
        return ""
    }
}