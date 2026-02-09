package ptx

// CacheOp represents PTX cache operators for ld/st instructions.
type CacheOp int

const (
    CacheNone CacheOp = iota

    // Load cache operators
    CacheCA           // .ca — cache at all levels (default for ld)
    CacheCG           // .cg — cache at global level (L2 only)
    CacheCS           // .cs — cache streaming (likely to be accessed once)
    CacheLU           // .lu — last use (evict first)
    CacheCV           // .cv — don't cache, volatile (bypass L1)

    // Store cache operators
    CacheWB           // .wb — write back at all levels (default for st)
    CacheWT           // .wt — write through to system memory
)

func (c CacheOp) String() string {
    switch c {
    case CacheNone:
        return ""
    case CacheCA:
        return ".ca"
    case CacheCG:
        return ".cg"
    case CacheCS:
        return ".cs"
    case CacheLU:
        return ".lu"
    case CacheCV:
        return ".cv"
    case CacheWB:
        return ".wb"
    case CacheWT:
        return ".wt"
    default:
        return ""
    }
}