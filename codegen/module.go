package codegen

import (
    "github.com/arc-language/ptx-gen/builder"
)

// emitModule emits the entire PTX module.
func (e *Emitter) emitModule(mod *builder.Module) {
    // Header directives
    e.linef(".version %s", mod.Version.String())
    e.linef(".target %s", mod.Target.String())
    e.linef(".address_size %d", mod.AddressSize)
    e.blank()

    // Module-scope globals
    for _, g := range mod.Globals {
        e.emitGlobal(g)
    }
    if len(mod.Globals) > 0 {
        e.blank()
    }

    // Functions and kernels
    for i, f := range mod.Functions {
        if i > 0 {
            e.blank()
        }
        e.emitFunction(f)
    }
}