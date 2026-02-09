package builder

import (
	"fmt"

	"github.com/arc-language/ptx-gen/ptx"
)

// Function represents a .entry (kernel) or .func (device function) definition.
type Function struct {
	Name     string
	IsKernel bool        // true = .entry, false = .func
	Linkage  ptx.Linkage // .visible, .extern, .weak, etc.

	// Parameters
	Params       []*Param // input parameters
	ReturnParams []*Param // return parameters (device functions only)

	// Body
	Blocks    []*BasicBlock // ordered basic blocks
	Registers []*Register   // all declared registers (collected for .reg declarations)

	// Performance tuning directives
	Directives []*Directive

	// Function Attributes (Section 5.4.8)
	// e.g. .attribute(.unified(uuid1, uuid2))
	Attributes []VarAttribute

	// Internal counter for auto-naming registers
	regCounter map[string]int
}

// AddParam appends a kernel or function input parameter.
func (f *Function) AddParam(p *Param) *Function {
	f.Params = append(f.Params, p)
	return f
}

// AddReturnParam appends a return parameter (device functions only).
func (f *Function) AddReturnParam(p *Param) *Function {
	f.ReturnParams = append(f.ReturnParams, p)
	return f
}

// Param looks up a parameter by name and returns it as a Symbol operand
// suitable for use in ld.param / st.param instructions.
func (f *Function) Param(name string) *Symbol {
	for _, p := range f.Params {
		if p.Name == name {
			return &Symbol{Name: p.Name}
		}
	}
	for _, p := range f.ReturnParams {
		if p.Name == name {
			return &Symbol{Name: p.Name}
		}
	}
	return &Symbol{Name: name}
}

// NewReg declares a new named register scoped to this function.
func (f *Function) NewReg(name string, typ ptx.Type) *Register {
	r := &Register{
		Name: "%" + name,
		Typ:  typ,
	}
	f.Registers = append(f.Registers, r)
	return r
}

// TempReg creates an auto-named temporary register (%t0, %t1, ...).
func (f *Function) TempReg(typ ptx.Type) *Register {
	if f.regCounter == nil {
		f.regCounter = make(map[string]int)
	}
	prefix := tempPrefix(typ)
	idx := f.regCounter[prefix]
	f.regCounter[prefix] = idx + 1
	r := &Register{
		Name: fmt.Sprintf("%%%s%d", prefix, idx),
		Typ:  typ,
	}
	f.Registers = append(f.Registers, r)
	return r
}

// tempPrefix returns a conventional register prefix based on type.
func tempPrefix(t ptx.Type) string {
	switch {
	case t == ptx.Pred:
		return "p"
	case t == ptx.F32 || t == ptx.F64 || t == ptx.F16:
		return "fd"
	case t == ptx.U64 || t == ptx.S64 || t == ptx.B64:
		return "rd"
	default:
		return "r"
	}
}

// NewBlock creates a new labeled basic block and appends it to the function.
func (f *Function) NewBlock(label string) *BasicBlock {
	bb := &BasicBlock{
		Label: label,
	}
	f.Blocks = append(f.Blocks, bb)
	return bb
}

// AddDirective appends a performance-tuning directive.
func (f *Function) AddDirective(d *Directive) *Function {
	f.Directives = append(f.Directives, d)
	return f
}

// WithAttribute adds a function attribute (e.g., .attribute(.unified(...))).
func (f *Function) WithAttribute(attr VarAttribute) *Function {
	f.Attributes = append(f.Attributes, attr)
	return f
}