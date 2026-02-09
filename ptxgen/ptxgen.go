// ptxgen.go
package ptxgen

import (
    "github.com/arc-language/ptx-gen/builder"
    "github.com/arc-language/ptx-gen/codegen"
    "github.com/arc-language/ptx-gen/ptx"
)

func NewModule(version ptx.ISAVersion, target ptx.Target) *builder.Module {
    return builder.NewModule(version, target)
}

func Build(mod *builder.Module) string {
    return codegen.Emit(mod)
}