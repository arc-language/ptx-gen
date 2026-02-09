package builder

// BasicBlock represents a labeled sequence of instructions within a function.
// In PTX this is a label followed by instructions until the next label or end of function.
type BasicBlock struct {
    Label        string
    Instructions []*Instruction
}

// Add appends an instruction to this block.
func (bb *BasicBlock) Add(inst *Instruction) *BasicBlock {
    bb.Instructions = append(bb.Instructions, inst)
    return bb
}