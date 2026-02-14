package builder

// DirectiveKind represents PTX performance-tuning and metadata directives.
type DirectiveKind int

const (
	DirMaxNReg          DirectiveKind = iota // .maxnreg N
	DirMaxNTid                               // .maxntid nx, ny, nz
	DirReqNTid                               // .reqntid nx, ny, nz
	DirMinNCTAPerSM                          // .minnctapersm N
	DirMaxNCTAPerSM                          // .maxnctapersm N
	DirPragma                                // .pragma "string"
	DirReqNCluster                           // .reqnctapercluster (sm_90+)
	DirNoReturn                              // .noreturn
	DirAbiPreserve                           // .abi_preserve N
	DirAbiPreserveCtrl                       // .abi_preserve_control N
	DirExplicitCluster                       // .explicitcluster
	DirMaxClusterRank                        // .maxclusterrank N
	DirBlocksAreClusters                     // .blocksareclusters
	DirAlias                                 // .alias
)

// Directive represents a single performance-tuning directive on a function.
type Directive struct {
	Kind   DirectiveKind
	Values []int  // numeric values (e.g. maxnreg=128, maxntid=256,1,1)
	Text   string // for .pragma
}

// MaxNReg creates a .maxnreg directive.
func MaxNReg(n int) *Directive {
	return &Directive{Kind: DirMaxNReg, Values: []int{n}}
}

// MaxNTid creates a .maxntid directive (1D, 2D, or 3D).
func MaxNTid(dims ...int) *Directive {
	return &Directive{Kind: DirMaxNTid, Values: dims}
}

// ReqNTid creates a .reqntid directive (1D, 2D, or 3D).
func ReqNTid(dims ...int) *Directive {
	return &Directive{Kind: DirReqNTid, Values: dims}
}

// MinNCTAPerSM creates a .minnctapersm directive.
func MinNCTAPerSM(n int) *Directive {
	return &Directive{Kind: DirMinNCTAPerSM, Values: []int{n}}
}

// MaxNCTAPerSM creates a .maxnctapersm directive.
func MaxNCTAPerSM(n int) *Directive {
	return &Directive{Kind: DirMaxNCTAPerSM, Values: []int{n}}
}

// Pragma creates a .pragma directive.
func Pragma(text string) *Directive {
	return &Directive{Kind: DirPragma, Text: text}
}

// NoReturn creates a .noreturn directive.
func NoReturn() *Directive {
	return &Directive{Kind: DirNoReturn}
}

// AbiPreserve creates a .abi_preserve N directive.
func AbiPreserve(n int) *Directive {
	return &Directive{Kind: DirAbiPreserve, Values: []int{n}}
}

// AbiPreserveCtrl creates a .abi_preserve_control N directive.
func AbiPreserveCtrl(n int) *Directive {
	return &Directive{Kind: DirAbiPreserveCtrl, Values: []int{n}}
}

// ExplicitCluster creates a .explicitcluster directive.
func ExplicitCluster() *Directive {
	return &Directive{Kind: DirExplicitCluster}
}

// MaxClusterRank creates a .maxclusterrank N directive.
func MaxClusterRank(n int) *Directive {
	return &Directive{Kind: DirMaxClusterRank, Values: []int{n}}
}

// BlocksAreClusters creates a .blocksareclusters directive.
func BlocksAreClusters() *Directive {
	return &Directive{Kind: DirBlocksAreClusters}
}

// Alias creates a .alias directive (module scope).
// Note: usage typically requires a custom emit function in codegen if not attached to a function.
func Alias(alias, aliasee string) *Directive {
	return &Directive{Kind: DirAlias, Text: alias + ", " + aliasee}
}