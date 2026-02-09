package builder

import (
    "github.com/arc-language/ptx-gen/ptx"
)

// Operand is the interface for all instruction operands.
type Operand interface {
    operandMarker()
}

// Register represents a named PTX register (%r0, %fd1, %p0).
type Register struct {
    Name string
    Typ  ptx.Type
}

// Immediate represents a literal constant value (integer or float).
type Immediate struct {
    Value interface{} // int64, uint64, float32, float64
}

// Symbol represents a named symbol: function name, label, global variable, or parameter name.
type Symbol struct {
    Name string
}

// Address represents a memory address operand: [base], [base+offset], or [symbol+offset].
type Address struct {
    Base   Operand  // *Register or *Symbol
    Offset int64    // byte offset
}

// VectorOp represents a vector operand {r0, r1, r2, r3} for .v2/.v4 ld/st.
type VectorOp struct {
    Elements []Operand
}

// SpecialRegOp represents a PTX special register operand (%tid.x, %ctaid.y, etc.).
type SpecialRegOp struct {
    Reg ptx.SpecialReg
}

// --- Interface compliance ---

func (*Register) operandMarker()     {}
func (*Immediate) operandMarker()    {}
func (*Symbol) operandMarker()       {}
func (*Address) operandMarker()      {}
func (*VectorOp) operandMarker()     {}
func (*SpecialRegOp) operandMarker() {}

// --- Convenience constructors ---

// Imm creates an immediate integer operand.
func Imm(val int64) *Immediate {
    return &Immediate{Value: val}
}

// ImmU creates an immediate unsigned integer operand.
func ImmU(val uint64) *Immediate {
    return &Immediate{Value: val}
}

// ImmF32 creates an immediate 32-bit float operand.
func ImmF32(val float32) *Immediate {
    return &Immediate{Value: val}
}

// ImmF64 creates an immediate 64-bit float operand.
func ImmF64(val float64) *Immediate {
    return &Immediate{Value: val}
}

// Addr creates a memory address operand [base+offset].
func Addr(base Operand, offset int64) *Address {
    return &Address{Base: base, Offset: offset}
}

// Sym creates a symbol operand (label, function name, global name).
func Sym(name string) *Symbol {
    return &Symbol{Name: name}
}

// SReg creates a special register operand (%tid.x, %ntid.x, etc.).
func SReg(r ptx.SpecialReg) *SpecialRegOp {
    return &SpecialRegOp{Reg: r}
}

// Vec creates a vector operand from multiple registers/operands.
func Vec(elems ...Operand) *VectorOp {
    return &VectorOp{Elements: elems}
}