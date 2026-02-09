package codegen

import (
    "fmt"
    "math"
    "strings"

    "github.com/arc-language/ptx-gen/builder"
)

// emitOperand formats a single operand to its PTX text representation.
func emitOperand(op builder.Operand) string {
    switch o := op.(type) {
    case *builder.Register:
        return o.Name

    case *builder.Immediate:
        return emitImmediate(o)

    case *builder.Symbol:
        return o.Name

    case *builder.Address:
        return emitAddress(o)

    case *builder.VectorOp:
        return emitVector(o)

    case *builder.SpecialRegOp:
        return o.Reg.String()

    default:
        return "???"
    }
}

// emitImmediate formats an immediate value.
//
// Integers:  42, -1, 0xFF
// Float32:   0f3F800000  (PTX hex float format)
// Float64:   0d3FF0000000000000
func emitImmediate(imm *builder.Immediate) string {
    switch v := imm.Value.(type) {
    case int:
        return fmt.Sprintf("%d", v)
    case int32:
        return fmt.Sprintf("%d", v)
    case int64:
        return fmt.Sprintf("%d", v)
    case uint32:
        return fmt.Sprintf("%d", v)
    case uint64:
        return fmt.Sprintf("%d", v)
    case float32:
        // PTX hex float: 0fXXXXXXXX
        bits := math.Float32bits(v)
        return fmt.Sprintf("0f%08X", bits)
    case float64:
        // PTX hex float: 0dXXXXXXXXXXXXXXXX
        bits := math.Float64bits(v)
        return fmt.Sprintf("0d%016X", bits)
    default:
        return fmt.Sprintf("%v", v)
    }
}

// emitAddress formats a memory address operand.
//
// [%rd0]         — register base, zero offset
// [%rd0+8]       — register base with offset
// [paramName]    — symbol (parameter) base
// [paramName+16] — symbol with offset
func emitAddress(addr *builder.Address) string {
    base := emitOperand(addr.Base)
    if addr.Offset == 0 {
        return fmt.Sprintf("[%s]", base)
    }
    if addr.Offset > 0 {
        return fmt.Sprintf("[%s+%d]", base, addr.Offset)
    }
    return fmt.Sprintf("[%s%d]", base, addr.Offset) // negative offset
}

// emitVector formats a vector operand {r0, r1, r2, r3}.
func emitVector(vec *builder.VectorOp) string {
    elems := make([]string, len(vec.Elements))
    for i, el := range vec.Elements {
        elems[i] = emitOperand(el)
    }
    return "{" + strings.Join(elems, ", ") + "}"
}