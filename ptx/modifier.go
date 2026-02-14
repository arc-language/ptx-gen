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

    default:
        return ""
    }
}