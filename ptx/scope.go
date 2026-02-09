package ptx

// Scope represents the memory ordering scope for atomics, fences, and ld/st.
type Scope int

const (
    ScopeNone Scope = iota
    ScopeCTA                // .cta  — within the CTA
    ScopeCluster            // .cluster — within the cluster (sm_90+)
    ScopeGPU                // .gpu  — within the GPU device
    ScopeSystem             // .sys  — across host + all devices
)

func (s Scope) String() string {
    switch s {
    case ScopeNone:
        return ""
    case ScopeCTA:
        return ".cta"
    case ScopeCluster:
        return ".cluster"
    case ScopeGPU:
        return ".gpu"
    case ScopeSystem:
        return ".sys"
    default:
        return ""
    }
}

// MembarLevel represents membar fence levels.
type MembarLevel int

const (
    MembarCTA    MembarLevel = iota  // membar.cta
    MembarGL                          // membar.gl  (global)
    MembarSys                         // membar.sys (system)
)

func (m MembarLevel) String() string {
    switch m {
    case MembarCTA:
        return ".cta"
    case MembarGL:
        return ".gl"
    case MembarSys:
        return ".sys"
    default:
        return ""
    }
}