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
//
//	.visible .global .f32 gvar;
//	.global .align 16 .b8 buffer[4096];
//	.shared .f32 smem[256];
//	.const .b32 lookup[16] = {0, 1, 2, 3};
//	.global .attribute(.managed) .s32 g;
func (e *Emitter) emitGlobal(g *builder.Global) {
	var parts []string

	// 1. Linkage (.visible, .extern, etc.)
	if g.Linkage != ptx.LinkNone {
		parts = append(parts, g.Linkage.String())
	}

	// 2. State space (.global, .shared, .const)
	parts = append(parts, g.Space.String())

	// 3. Attributes (.attribute(.managed), etc.)
	// Section 5.4.8 syntax: .global .attribute(...) .type var
	if len(g.Attributes) > 0 {
		var attrs []string
		for _, attr := range g.Attributes {
			if len(attr.Params) > 0 {
				// Format params like .unified(0xAB, 0xCD)
				var pStrs []string
				for _, p := range attr.Params {
					pStrs = append(pStrs, fmt.Sprintf("%v", p))
				}
				attrs = append(attrs, fmt.Sprintf(".%s(%s)", attr.Name, strings.Join(pStrs, ", ")))
			} else {
				attrs = append(attrs, fmt.Sprintf(".%s", attr.Name))
			}
		}
		parts = append(parts, fmt.Sprintf(".attribute(%s)", strings.Join(attrs, ", ")))
	}

	// 4. Alignment
	if g.Align > 0 {
		parts = append(parts, fmt.Sprintf(".align %d", g.Align))
	}

	// 5. Vector width
	if g.Vec != ptx.Scalar {
		parts = append(parts, g.Vec.String())
	}

	// 6. Type
	parts = append(parts, g.Typ.String())

	// 7. Name + optional array count
	if g.Count > 0 {
		parts = append(parts, fmt.Sprintf("%s[%d]", g.Name, g.Count))
	} else {
		parts = append(parts, g.Name)
	}

	// 8. Initializer
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