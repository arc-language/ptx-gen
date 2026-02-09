package codegen

import (
	"fmt"
	"strings"

	"github.com/arc-language/ptx-gen/builder"
	"github.com/arc-language/ptx-gen/ptx"
)

// emitFunction emits a .entry or .func definition with params, registers, directives, and body.
func (e *Emitter) emitFunction(f *builder.Function) {
	// --- Signature line ---
	e.emitFunctionSignature(f)

	// --- Open body ---
	e.line("{")
	e.push()

	// --- Register declarations ---
	e.emitRegisterDecls(f)

	// --- Performance tuning directives ---
	for _, d := range f.Directives {
		e.emitDirective(d)
	}

	if len(f.Registers) > 0 || len(f.Directives) > 0 {
		e.blank()
	}

	// --- Basic blocks ---
	for i, bb := range f.Blocks {
		if i > 0 {
			e.blank()
		}
		e.emitBlock(bb)
	}

	// --- Close body ---
	e.pop()
	e.line("}")
}

// emitFunctionSignature emits the .entry/.func line with parameters and attributes.
//
// Output examples:
//
//	.visible .entry vec_add(
//	    .param .u64 A,
//	    .param .u64 B,
//	    .param .u32 N
//	)
//
//	.func .attribute(.unified(0xAB, 0xCD)) bar()
func (e *Emitter) emitFunctionSignature(f *builder.Function) {
	var prefix []string

	// 1. Linkage
	if f.Linkage != ptx.LinkNone {
		prefix = append(prefix, f.Linkage.String())
	}

	// 2. .entry or .func
	if f.IsKernel {
		prefix = append(prefix, ".entry")
	} else {
		prefix = append(prefix, ".func")
	}

	// 3. Attributes (Section 5.4.8)
	// Format: .func .attribute(...) name
	if len(f.Attributes) > 0 {
		var attrs []string
		for _, attr := range f.Attributes {
			if len(attr.Params) > 0 {
				var pStrs []string
				for _, p := range attr.Params {
					pStrs = append(pStrs, fmt.Sprintf("%v", p))
				}
				attrs = append(attrs, fmt.Sprintf(".%s(%s)", attr.Name, strings.Join(pStrs, ", ")))
			} else {
				attrs = append(attrs, fmt.Sprintf(".%s", attr.Name))
			}
		}
		prefix = append(prefix, fmt.Sprintf(".attribute(%s)", strings.Join(attrs, ", ")))
	}

	// For .func with return params, emit return param list before name
	// Syntax: .func (.reg .u32 r) name ...
	head := strings.Join(prefix, " ")

	if !f.IsKernel && len(f.ReturnParams) > 0 {
		retParts := make([]string, len(f.ReturnParams))
		for i, rp := range f.ReturnParams {
			retParts[i] = emitParamDecl(rp, f.IsKernel)
		}
		e.writef("%s (%s) %s", head, strings.Join(retParts, ", "), f.Name)
	} else {
		// Just name
		if len(prefix) > 0 {
			e.writeIndent() // Ensure indentation if we started a new line logic, though usually top level
			e.writef("%s %s", head, f.Name)
		} else {
			e.writeIndent()
			e.writef("%s", f.Name)
		}
	}

	// Input parameters
	if len(f.Params) == 0 {
		e.write("()\n")
		return
	}

	e.write("(\n")
	e.push()
	for i, p := range f.Params {
		e.writeIndent()
		e.write(emitParamDecl(p, f.IsKernel))
		if i < len(f.Params)-1 {
			e.write(",")
		}
		e.write("\n")
	}
	e.pop()
	e.line(")")
}

// emitParamDecl formats a single parameter declaration string.
//
// Kernel params:  .param .u64 .ptr .global .align 8 A
// Kernel params:  .param .u32 N
// Kernel params:  .param .align 8 .b8 buffer[64]
// Func params:    .reg .u32 a
func emitParamDecl(p *builder.Param, isKernel bool) string {
	var parts []string

	if isKernel {
		parts = append(parts, ".param")

		// Alignment (for byte array params)
		if p.Align > 0 {
			parts = append(parts, fmt.Sprintf(".align %d", p.Align))
		}

		// Type
		parts = append(parts, p.Typ.String())

		// Pointer attribute
		if p.IsPointer {
			parts = append(parts, ".ptr")
			parts = append(parts, p.PtrSpace.String())
			// Default alignment for pointers
			if p.Align > 0 {
				parts = append(parts, fmt.Sprintf(".align %d", p.Align))
			}
		}

		// Name + optional byte array size
		if p.Size > 0 {
			parts = append(parts, fmt.Sprintf("%s[%d]", p.Name, p.Size))
		} else {
			parts = append(parts, p.Name)
		}
	} else {
		// Device function params: can be .reg or .param
		if p.Size > 0 {
			// Pass-by-value struct: .param .align A .b8 name[size]
			parts = append(parts, ".param")
			if p.Align > 0 {
				parts = append(parts, fmt.Sprintf(".align %d", p.Align))
			}
			parts = append(parts, p.Typ.String())
			parts = append(parts, fmt.Sprintf("%s[%d]", p.Name, p.Size))
		} else {
			parts = append(parts, ".reg")
			parts = append(parts, p.Typ.String())
			parts = append(parts, p.Name)
		}
	}

	return strings.Join(parts, " ")
}

// emitRegisterDecls groups registers by type and emits .reg declarations.
//
// Output example:
//
//	.reg .u32  %tid, %n;
//	.reg .u64  %a_ptr, %b_ptr, %offset;
//	.reg .f32  %a_val, %b_val, %result;
//	.reg .pred %p;
func (e *Emitter) emitRegisterDecls(f *builder.Function) {
	// Group registers by type
	groups := make(map[ptx.Type][]*builder.Register)
	var order []ptx.Type

	for _, r := range f.Registers {
		if _, exists := groups[r.Typ]; !exists {
			order = append(order, r.Typ)
		}
		groups[r.Typ] = append(groups[r.Typ], r)
	}

	for _, typ := range order {
		regs := groups[typ]
		names := make([]string, len(regs))
		for i, r := range regs {
			names[i] = r.Name
		}
		e.linef(".reg %-6s %s;", typ.String(), strings.Join(names, ", "))
	}
}

// emitDirective emits a performance-tuning directive.
func (e *Emitter) emitDirective(d *builder.Directive) {
	switch d.Kind {
	case builder.DirMaxNReg:
		e.linef(".maxnreg %d", d.Values[0])
	case builder.DirMaxNTid:
		e.linef(".maxntid %s", joinInts(d.Values))
	case builder.DirReqNTid:
		e.linef(".reqntid %s", joinInts(d.Values))
	case builder.DirMinNCTAPerSM:
		e.linef(".minnctapersm %d", d.Values[0])
	case builder.DirMaxNCTAPerSM:
		e.linef(".maxnctapersm %d", d.Values[0])
	case builder.DirPragma:
		e.linef(".pragma \"%s\";", d.Text)
	case builder.DirReqNCluster:
		e.linef(".reqnctapercluster %s", joinInts(d.Values))
	}
}

// emitBlock emits a labeled basic block.
func (e *Emitter) emitBlock(bb *builder.BasicBlock) {
	// Label (outdented one level from instructions)
	if bb.Label != "" {
		e.pop()
		e.linef("%s:", bb.Label)
		e.push()
	}

	// Instructions
	for _, inst := range bb.Instructions {
		e.emitInstruction(inst)
	}
}

// joinInts formats a slice of ints as comma-separated string.
func joinInts(vals []int) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, ", ")
}