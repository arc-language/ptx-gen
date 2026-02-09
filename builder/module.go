package builder

import (
    "github.com/arc-language/ptx-gen/ptx"
)

// Module is the top-level PTX program container.
// It maps to a single .ptx file with directives, globals, and functions.
type Module struct {
    Version     ptx.ISAVersion  // .version 8.5
    Target      ptx.Target      // .target sm_80
    AddressSize int             // .address_size 64 (32 or 64)
    Globals     []*Global       // module-scope variables (.global, .const, .shared)
    Functions   []*Function     // .entry and .func definitions
}

// NewModule creates a new PTX module with sensible defaults.
func NewModule(version ptx.ISAVersion, target ptx.Target) *Module {
    return &Module{
        Version:     version,
        Target:      target,
        AddressSize: 64,
    }
}

// AddGlobal appends a module-scope variable declaration.
func (m *Module) AddGlobal(g *Global) *Module {
    m.Globals = append(m.Globals, g)
    return m
}

// AddFunction appends a function or kernel definition.
func (m *Module) AddFunction(f *Function) *Module {
    m.Functions = append(m.Functions, f)
    return m
}

// NewKernel creates a new kernel (.entry) function and adds it to the module.
func (m *Module) NewKernel(name string) *Function {
    f := &Function{
        Name:     name,
        IsKernel: true,
        Linkage:  ptx.LinkVisible,
    }
    m.Functions = append(m.Functions, f)
    return f
}

// NewFunc creates a new device function (.func) and adds it to the module.
func (m *Module) NewFunc(name string) *Function {
    f := &Function{
        Name:     name,
        IsKernel: false,
    }
    m.Functions = append(m.Functions, f)
    return f
}