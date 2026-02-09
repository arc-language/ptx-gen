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

	// --- Opaque Handle Spaces (conceptually global/param, but useful for strict typing) ---
	// These usually map to specific variable declarations but aren't always used as instruction prefixes
	// unlike .global, .shared, etc.
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
	case SharedCTA:
		return ".shared::cta"
	case SharedCluster:
		return ".shared::cluster"
	case Tex:
		return ".tex"
	default:
		return ".unknown"
	}
}