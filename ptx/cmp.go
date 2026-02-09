package ptx

// CmpOp represents PTX comparison operators.
type CmpOp int

const (
    // Integer & bit-size comparisons
    CmpEq  CmpOp = iota // eq  — equal
    CmpNe                // ne  — not equal
    CmpLt                // lt  — less than
    CmpLe                // le  — less than or equal
    CmpGt                // gt  — greater than
    CmpGe                // ge  — greater than or equal
    CmpLo                // lo  — lower (unsigned less than)
    CmpLs                // ls  — lower or same (unsigned <=)
    CmpHi                // hi  — higher (unsigned greater than)
    CmpHs                // hs  — higher or same (unsigned >=)

    // Float ordered comparisons (neither operand is NaN)
    CmpEqu               // equ — equal, unordered
    CmpNeu               // neu — not equal, unordered
    CmpLtu               // ltu — less than, unordered
    CmpLeu               // leu — less or equal, unordered
    CmpGtu               // gtu — greater than, unordered
    CmpGeu               // geu — greater or equal, unordered

    // Float NaN-testing comparisons
    CmpNum               // num — both operands are numbers (not NaN)
    CmpNan               // nan — either operand is NaN
)

func (c CmpOp) String() string {
    switch c {
    case CmpEq:
        return ".eq"
    case CmpNe:
        return ".ne"
    case CmpLt:
        return ".lt"
    case CmpLe:
        return ".le"
    case CmpGt:
        return ".gt"
    case CmpGe:
        return ".ge"
    case CmpLo:
        return ".lo"
    case CmpLs:
        return ".ls"
    case CmpHi:
        return ".hi"
    case CmpHs:
        return ".hs"
    case CmpEqu:
        return ".equ"
    case CmpNeu:
        return ".neu"
    case CmpLtu:
        return ".ltu"
    case CmpLeu:
        return ".leu"
    case CmpGtu:
        return ".gtu"
    case CmpGeu:
        return ".geu"
    case CmpNum:
        return ".num"
    case CmpNan:
        return ".nan"
    default:
        return ".unknown"
    }
}