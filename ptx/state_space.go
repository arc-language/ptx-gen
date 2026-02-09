package ptx

// StateSpace represents PTX memory state spaces.
type StateSpace int

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

    // Param — kernel or function parameters.
    Param

    // ParamEntry — explicitly kernel (.param::entry) parameters.
    ParamEntry

    // ParamFunc — explicitly device function (.param::func) parameters.
    ParamFunc

    // Shared — per-CTA memory visible to all threads in the CTA.
    Shared

    // Tex — deprecated texture state space (legacy).
    Tex
)

func (s StateSpace) String() string {
    switch s {
    case Reg:
        return ".reg"
    case SReg:
        return ".sreg"
    case Const:
        return ".const"
    case Global:
        return ".global"
    case Local:
        return ".local"
    case Param:
        return ".param"
    case ParamEntry:
        return ".param::entry"
    case ParamFunc:
        return ".param::func"
    case Shared:
        return ".shared"
    case Tex:
        return ".tex"
    default:
        return ".unknown"
    }
}