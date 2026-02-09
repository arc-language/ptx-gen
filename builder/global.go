package builder

import (
    "github.com/arc-language/ptx-gen/ptx"
)

// Global represents a module-scope variable declaration in .global, .shared, or .const space.
//
// Examples:
//   .global .f32 gvar;
//   .global .f32 garr[100];
//   .shared .f32 smem[256];
//   .const .b32 lookup[16] = {0, 1, 2, ...};
//   .global .align 16 .b8 buffer[4096];
type Global struct {
    Name        string
    Space       ptx.StateSpace  // .global, .shared, .const
    Typ         ptx.Type        // element type
    Vec         ptx.VectorSize  // .v2, .v4 (Scalar for non-vector)
    Count       int             // array element count (0 = scalar)
    Align       int             // alignment in bytes (0 = default)
    Linkage     ptx.Linkage     // .visible, .extern, etc.
    Initializer []interface{}   // optional initializer values (int64, float64, etc.)
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