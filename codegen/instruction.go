package codegen

import (
    "strings"

    "github.com/arc-language/ptx-gen/builder"
    "github.com/arc-language/ptx-gen/ptx"
)

// emitInstruction emits a single PTX instruction line.
//
// Output examples:
//   add.u32        %r0, %r1, %r2;
//   ld.global.f32  %f0, [%rd0];
//   setp.ge.u32    %p, %r0, %r1;
//   @%p bra        done;
//   cvt.u64.u32    %rd0, %r0;
//   st.global.f32  [%rd0], %f0;
//   bar.sync       0;
//   ret;
func (e *Emitter) emitInstruction(inst *builder.Instruction) {
    var line strings.Builder

    // Guard predicate: @%p or @!%p
    if inst.Guard != nil {
        if inst.Guard.Negate {
            line.WriteString("@!")
        } else {
            line.WriteString("@")
        }
        line.WriteString(inst.Guard.Reg.Name)
        line.WriteString(" ")
    }

    // Mnemonic: opcode + modifiers + space + cache + scope + rounding + vec + type
    line.WriteString(buildMnemonic(inst))

    // Operands
    ops := buildOperands(inst)
    if ops != "" {
        // Pad mnemonic to a column for alignment
        mnem := line.String()
        padded := padTo(mnem, 15)
        line.Reset()
        line.WriteString(padded)
        line.WriteString(ops)
    }

    line.WriteString(";")
    e.line(line.String())
}

// buildMnemonic constructs the full instruction mnemonic string.
//
// Examples:
//   add.u32
//   ld.global.f32
//   setp.ge.u32
//   cvt.rn.f32.u32
//   mul.wide.u32
//   atom.global.add.u32
//   st.global.f32
//   bar.sync
//   shfl.sync.up.b32
func buildMnemonic(inst *builder.Instruction) string {
    var parts []string

    // Base opcode
    parts = append(parts, inst.Op.String())

    // Modifiers that come right after opcode (.wide, .lo, .hi, etc.)
    for _, m := range inst.Modifiers {
        parts = append(parts, m.String())
    }

    // State space (.global, .shared, .param, .local, .const)
    if inst.Space != ptx.Reg && inst.Space != 0 {
        // ptx.Reg is 0/default, skip it
        parts = append(parts, inst.Space.String())
    }

    // Cache operator (.ca, .cg, .cs, .cv)
    if inst.Cache != ptx.CacheNone {
        parts = append(parts, inst.Cache.String())
    }

    // Memory scope (.cta, .gpu, .sys)
    if inst.Scope != ptx.ScopeNone {
        parts = append(parts, inst.Scope.String())
    }

    // Comparison operator (.eq, .lt, .ge, etc.)
    if inst.Op == ptx.OpSetp || inst.Op == ptx.OpSet || inst.Op == ptx.OpSlct {
        parts = append(parts, inst.Cmp.String())
    }

    // Rounding modifier (.rn, .rz, .rm, .rp)
    if inst.Rounding != ptx.RoundNone {
        parts = append(parts, inst.Rounding.String())
    }

    // Vector width (.v2, .v4)
    if inst.Vec != ptx.Scalar {
        parts = append(parts, inst.Vec.String())
    }

    // Instruction type (.u32, .f32, .f64, etc.)
    if inst.Typ != 0 {
        // For cvt: emit dstType.srcType
        if inst.Op == ptx.OpCvt && inst.SrcType != 0 {
            parts = append(parts, inst.Typ.String()+"."+stripDot(inst.SrcType.String()))
        } else {
            parts = append(parts, inst.Typ.String())
        }
    }

    return strings.Join(parts, "")
}

// buildOperands constructs the comma-separated operand string.
func buildOperands(inst *builder.Instruction) string {
    var ops []string

    // Special case: call has different formatting
    if inst.Op == ptx.OpCall {
        return buildCallOperands(inst)
    }

    // Special case: st has no Dst, operands are [addr], src
    if inst.Op == ptx.OpSt || inst.Op == ptx.OpStAsync {
        for _, s := range inst.Src {
            ops = append(ops, emitOperand(s))
        }
        return strings.Join(ops, ", ")
    }

    // Destination
    if inst.Dst != nil {
        ops = append(ops, emitOperand(inst.Dst))
    }

    // Sources
    for _, s := range inst.Src {
        ops = append(ops, emitOperand(s))
    }

    return strings.Join(ops, ", ")
}

// buildCallOperands handles the special call syntax:
//   call (retval), funcname, (arg0, arg1, ...);
//   call funcname, (arg0, arg1, ...);
func buildCallOperands(inst *builder.Instruction) string {
    var result strings.Builder

    // Return values
    if inst.Dst != nil {
        result.WriteString("(")
        result.WriteString(emitOperand(inst.Dst))
        result.WriteString("), ")
    }

    // Function name
    result.WriteString(inst.CallTarget)

    // Arguments
    if len(inst.Src) > 0 {
        result.WriteString(", (")
        for i, arg := range inst.Src {
            if i > 0 {
                result.WriteString(", ")
            }
            result.WriteString(emitOperand(arg))
        }
        result.WriteString(")")
    }

    return result.String()
}

// stripDot removes the leading dot from a type string for cvt dual-type syntax.
// ".u32" -> "u32"
func stripDot(s string) string {
    if len(s) > 0 && s[0] == '.' {
        return s[1:]
    }
    return s
}

// padTo pads a string with spaces to reach the target length.
func padTo(s string, target int) string {
    if len(s) >= target {
        return s + " "
    }
    return s + strings.Repeat(" ", target-len(s))
}