package builder

import (
    "github.com/arc-language/ptx-gen/ptx"
)

// Instruction represents a single PTX instruction node in the IR tree.
type Instruction struct {
    Op        ptx.Opcode        // opcode: add, ld, st, mov, setp, bra, etc.
    Typ       ptx.Type          // instruction type: .u32, .f32, etc.
    
    Dst       Operand           // Primary destination
    Dst2      Operand           // Secondary destination (for setp p|q)
    
    Src       []Operand         // source operands
    Space     ptx.StateSpace    // state space for ld/st
    
    Cmp       ptx.CmpOp         // comparison operator (.eq, .lt, etc.)
    BoolOp    ptx.BoolOp        // boolean operator (.and, .or, .xor)
    
    Rounding  ptx.RoundingMode  // .rn, .rz, etc.
    Cache     ptx.CacheOp       // .ca, .cg, etc.
    Scope     ptx.Scope         // .cta, .gpu, etc.
    Vec       ptx.VectorSize    // .v2, .v4
    Modifiers []ptx.Modifier    // .wide, .lo, .sat, etc.
    Guard     *Predicate        // @p

    SrcType   ptx.Type          // For cvt (source type)
    CallTarget string           // For call
}

// Predicate represents a guard predicate on an instruction: @p or @!p
type Predicate struct {
    Reg    *Register
    Negate bool
}

// --- Chaining methods ---

// Typed sets the instruction type (.u32, .f32, .f64, etc.).
func (i *Instruction) Typed(t ptx.Type) *Instruction {
    i.Typ = t
    return i
}

// InSpace sets the state space (.global, .shared, .param, .local, .const).
func (i *Instruction) InSpace(s ptx.StateSpace) *Instruction {
    i.Space = s
    return i
}

// WithRounding sets the rounding modifier (.rn, .rz, .rm, .rp).
func (i *Instruction) WithRounding(r ptx.RoundingMode) *Instruction {
    i.Rounding = r
    return i
}

// WithCache sets the cache operator (.ca, .cg, .cs, .lu, .cv, .wb, .wt).
func (i *Instruction) WithCache(c ptx.CacheOp) *Instruction {
    i.Cache = c
    return i
}

// WithScope sets the memory scope (.cta, .gpu, .sys).
func (i *Instruction) WithScope(s ptx.Scope) *Instruction {
    i.Scope = s
    return i
}

// WithVec sets the vector width (.v2, .v4) for vector ld/st/mov.
func (i *Instruction) WithVec(v ptx.VectorSize) *Instruction {
    i.Vec = v
    return i
}

// WithMod appends one or more modifiers (.wide, .lo, .hi, .sat, .ftz, .approx, etc.).
func (i *Instruction) WithMod(mods ...ptx.Modifier) *Instruction {
    i.Modifiers = append(i.Modifiers, mods...)
    return i
}

// Pred sets the guard predicate (@p).
func (i *Instruction) Pred(reg *Register) *Instruction {
    i.Guard = &Predicate{Reg: reg, Negate: false}
    return i
}

// PredNot sets a negated guard predicate (@!p).
func (i *Instruction) PredNot(reg *Register) *Instruction {
    i.Guard = &Predicate{Reg: reg, Negate: true}
    return i
}

// From sets the source type for cvt instructions (cvt.dstType.srcType).
func (i *Instruction) From(t ptx.Type) *Instruction {
    i.SrcType = t
    return i
}

// WithBoolOp sets the boolean operator for set/setp (e.g., .and, .or).
func (i *Instruction) WithBoolOp(op ptx.BoolOp) *Instruction {
	i.BoolOp = op // Requires 'BoolOp ptx.BoolOp' field in Instruction struct
	return i
}

// WithDst2 sets the secondary destination (e.g., q in setp p|q, ...).
func (i *Instruction) WithDst2(dst Operand) *Instruction {
	i.Dst2 = dst // Requires 'Dst2 Operand' field in Instruction struct
	return i
}

// SourceTyped sets the source type for mixed-precision or conversion instructions.
// Example: add.f32.f16 -> Typ=F32, SrcType=F16
func (i *Instruction) SourceTyped(t ptx.Type) *Instruction {
	i.SrcType = t
	return i
}