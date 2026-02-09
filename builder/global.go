package builder

import (
	"github.com/arc-language/ptx-gen/ptx"
)

// VarAttribute represents attributes for variables (Section 5.4.8).
// Examples: .attribute(.managed), .attribute(.unified(uuid1, uuid2))
type VarAttribute struct {
	Name   string
	Params []interface{}
}

// Managed creates a .managed attribute.
// This attribute specifies that variable will be allocated at a location in
// unified virtual memory environment where host and other devices can reference it.
func Managed() VarAttribute {
	return VarAttribute{Name: "managed"}
}

// Unified creates a .unified attribute with UUID values.
// This attribute specifies that the variable has the same memory address on the host
// and on other devices in the system.
func Unified(uuid1, uuid2 uint64) VarAttribute {
	return VarAttribute{Name: "unified", Params: []interface{}{uuid1, uuid2}}
}

// Global represents a module-scope variable declaration in .global, .shared, or .const space.
//
// Examples:
//
//	.global .f32 gvar;
//	.global .f32 garr[100];
//	.shared .f32 smem[256];
//	.const .b32 lookup[16] = {0, 1, 2, ...};
//	.global .align 16 .b8 buffer[4096];
//	.global .attribute(.managed) .s32 g;
type Global struct {
	Name        string
	Space       ptx.StateSpace // .global, .shared, .const
	Typ         ptx.Type       // element type
	Vec         ptx.VectorSize // .v2, .v4 (Scalar for non-vector)
	Count       int            // array element count (0 = scalar)
	Align       int            // alignment in bytes (0 = default)
	Linkage     ptx.Linkage    // .visible, .extern, etc.
	Initializer []interface{}  // optional initializer values (int64, float64, etc.)
	Attributes  []VarAttribute // .attribute(...)
}

// NewGlobal creates a scalar global variable.
func NewGlobal(name string, space ptx.StateSpace, typ ptx.Type) *Global {
	return &Global{
		Name:  name,
		Space: space,
		Typ:   typ,
	}
}

// NewGlobalArray creates a global array variable.
func NewGlobalArray(name string, space ptx.StateSpace, typ ptx.Type, count int) *Global {
	return &Global{
		Name:  name,
		Space: space,
		Typ:   typ,
		Count: count,
	}
}

// WithAlign sets the alignment.
func (g *Global) WithAlign(bytes int) *Global {
	g.Align = bytes
	return g
}

// WithLinkage sets the linkage directive (.visible, .extern, etc.).
func (g *Global) WithLinkage(l ptx.Linkage) *Global {
	g.Linkage = l
	return g
}

// WithInit sets the initializer values.
func (g *Global) WithInit(vals ...interface{}) *Global {
	g.Initializer = vals
	return g
}

// WithAttribute adds a variable attribute (e.g., .managed, .unified).
func (g *Global) WithAttribute(attr VarAttribute) *Global {
	g.Attributes = append(g.Attributes, attr)
	return g
}