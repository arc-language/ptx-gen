package builder

import (
	"github.com/arc-language/ptx-gen/ptx"
)

// Param represents a kernel or device function parameter.
//
// Kernel parameters live in .param state space and are read via ld.param.
// Device function parameters can be registers or .param byte arrays.
type Param struct {
	Name      string
	Typ       ptx.Type       // .u32, .u64, .f32, etc.
	Size      int            // byte size for .param .align N .b8 name[Size] style
	Align     int            // optional alignment in bytes (0 = default)
	IsPointer bool           // if true, this param is a pointer (.ptr attribute)
	PtrSpace  ptx.StateSpace // state space the pointer points to (.global, .shared, etc.)
}

// NewParam creates a simple typed parameter (e.g. .param .u32 N).
func NewParam(name string, typ ptx.Type) *Param {
	return &Param{
		Name: name,
		Typ:  typ,
	}
}

// NewPtrParam creates a pointer parameter with .ptr attribute.
// Pointers are passed as .u64 in 64-bit addressing mode.
func NewPtrParam(name string, pointsTo ptx.StateSpace) *Param {
	return &Param{
		Name:      name,
		Typ:       ptx.U64,
		IsPointer: true,
		PtrSpace:  pointsTo,
	}
}

// NewByteArrayParam creates a .param .align A .b8 name[size] parameter
// used for passing structures by value.
func NewByteArrayParam(name string, size int, align int) *Param {
	return &Param{
		Name:  name,
		Typ:   ptx.B8,
		Size:  size,
		Align: align,
	}
}

// WithAlign sets the alignment for the parameter.
func (p *Param) WithAlign(align int) *Param {
	p.Align = align
	return p
}