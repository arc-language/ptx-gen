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
// Order: Opcode .Cmp .BoolOp .Modifiers .Space .Cache .Scope .Rounding .Vec .Type .SrcType
// Example: setp.lt.and.f32, ld.global.ca.v4.f32, cvt.rn.f16.f32, cp.async.ca.shared.global
func buildMnemonic(inst *builder.Instruction) string {
	var sb strings.Builder

	// 1. Opcode
	sb.WriteString(inst.Op.String())

	// 2. Comparison & Boolean Operators (set, setp)
	if inst.Op == ptx.OpSet || inst.Op == ptx.OpSetp {
		// Only append Cmp if specific logic requires it, usually strictly required for set/setp
		sb.WriteString(inst.Cmp.String())

		// Append Boolean Operator (.and, .or, .xor) if present
		if inst.BoolOp != ptx.BoolNone {
			sb.WriteString(inst.BoolOp.String())
		}
	}

	// 3. Modifiers
	// Handles .wide, .lo, .hi, .sat, .ftz, .approx, .sync, .multicast, etc.
	// Also handles explicit state space modifiers for cp.async (e.g. .shared::cluster)
	for _, mod := range inst.Modifiers {
		sb.WriteString(mod.String())
	}

	// 4. State Space
	// Standard ld/st/atom instructions use the Space field (.global, .shared, .const, etc.)
	// Note: We avoid printing .reg as it is the default implicit state.
	if inst.Space != ptx.Reg && inst.Space != ptx.StateSpace(0) {
		sb.WriteString(inst.Space.String())
	}

	// 5. Cache Operators
	// (.ca, .cg, .cs, .lu, .cv, .wb, .wt)
	sb.WriteString(inst.Cache.String())

	// 6. Scope
	// (.cta, .gpu, .sys, .cluster)
	sb.WriteString(inst.Scope.String())

	// 7. Rounding Mode
	// (.rn, .rz, .rm, .rp, .rni, .rzi, etc.)
	sb.WriteString(inst.Rounding.String())

	// 8. Vector Size
	// (.v2, .v4, .v8)
	sb.WriteString(inst.Vec.String())

	// 9. Types
	// Special handling for Conversion instructions (cvt, cvt.pack) and Mixed Precision
	if inst.Op == ptx.OpCvt || inst.Op == ptx.OpCvtPack {
		// Destination Type (convertType)
		if inst.Typ != 0 {
			sb.WriteString(inst.Typ.String())
		}
		// Source Type (abType)
		if inst.SrcType != 0 {
			sb.WriteString(inst.SrcType.String())
		}
		// cvt.pack 3rd type (cType) logic:
		// Syntax: cvt.pack.sat.convertType.abType.cType d, a, b, c
		// If 3 source operands are present (a, b, c), we append the cType.
		// In most contexts this is .b32.
		if inst.Op == ptx.OpCvtPack && len(inst.Src) > 2 {
			sb.WriteString(".b32")
		}
	} else {
		// Standard instructions (add.u32, ld.global.f32)
		if inst.Typ != 0 {
			sb.WriteString(inst.Typ.String())
		}

		// Mixed Precision / Explicit Source Type variants
		// Examples: add.f32.f16, sub.f32.bf16
		// Only append SrcType if it differs or is explicitly set and not handled above
		if inst.SrcType != 0 {
			sb.WriteString(inst.SrcType.String())
		}
	}

	return sb.String()
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