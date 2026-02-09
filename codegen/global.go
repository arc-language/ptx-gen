package codegen

import (
    "fmt"
    "strings"

    "github.com/arc-language/ptx-gen/builder"
    "github.com/arc-language/ptx-gen/ptx"
)

// emitGlobal emits a module-scope variable declaration.
//
// Examples:
//   .visible .global .f32 gvar;
//   .global .align 16 .b8 buffer[4096];
//   .shared .f32 smem[256];
//   .const .b32 lookup[16] = {0, 1, 2, 3};
func (e *Emitter) emitGlobal(g *builder.Global) {
    var parts []string

    // Linkage
    if g.Linkage != ptx.LinkNone {
        parts = append(parts, g.Linkage.String())
    }

    // State space
    parts = append(parts, g.Space.String())

    // Alignment
    if g.Align > 0 {
        parts = append(parts, fmt.Sprintf(".align %d", g.Align))
    }

    // Vector width
    if g.Vec != ptx.Scalar {
        parts = append(parts, g.Vec.String())
    }

    // Type
    parts = append(parts, g.Typ.String())

    // Name + optional array count
    if g.Count > 0 {
        parts = append(parts, fmt.Sprintf("%s[%d]", g.Name, g.Count))
    } else {
        parts = append(parts, g.Name)
    }

    // Initializer
    if len(g.Initializer) > 0 {
        vals := make([]string, len(g.Initializer))
        for i, v := range g.Initializer {
            vals[i] = fmt.Sprintf("%v", v)
        }
        e.linef("%s = {%s};", strings.Join(parts, " "), strings.Join(vals, ", "))
    } else {
        e.linef("%s;", strings.Join(parts, " "))
    }
}