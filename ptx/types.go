package ptx

// Type represents a PTX fundamental data type.
type Type int

const (
    // Predicate
    Pred Type = iota

    // Bit-size (untyped)
    B8
    B16
    B32
    B64
    B128

    // Signed integer
    S8
    S16
    S32
    S64

    // Unsigned integer
    U8
    U16
    U32
    U64

    // Floating point
    F16
    F32
    F64

    // Alternate floating point (sm_80+)
    BF16
    TF32

    // Packed types
    F16x2
    BF16x2

    // Fixed-point (sm_100+)
    E2M1
    E2M3
    E3M2
    E4M3
    E5M2
    E8M0
)

// VectorSize represents PTX vector widths.
type VectorSize int

const (
    Scalar VectorSize = iota
    V2                        // .v2
    V4                        // .v4
)

// String returns the PTX type string (e.g. ".f32", ".u64", ".pred").
func (t Type) String() string {
    switch t {
    case Pred:
        return ".pred"
    case B8:
        return ".b8"
    case B16:
        return ".b16"
    case B32:
        return ".b32"
    case B64:
        return ".b64"
    case B128:
        return ".b128"
    case S8:
        return ".s8"
    case S16:
        return ".s16"
    case S32:
        return ".s32"
    case S64:
        return ".s64"
    case U8:
        return ".u8"
    case U16:
        return ".u16"
    case U32:
        return ".u32"
    case U64:
        return ".u64"
    case F16:
        return ".f16"
    case F32:
        return ".f32"
    case F64:
        return ".f64"
    case BF16:
        return ".bf16"
    case TF32:
        return ".tf32"
    case F16x2:
        return ".f16x2"
    case BF16x2:
        return ".bf16x2"
    case E2M1:
        return ".e2m1"
    case E2M3:
        return ".e2m3"
    case E3M2:
        return ".e3m2"
    case E4M3:
        return ".e4m3"
    case E5M2:
        return ".e5m2"
    case E8M0:
        return ".e8m0"
    default:
        return ".unknown"
    }
}

// BitWidth returns the width in bits for the type.
func (t Type) BitWidth() int {
    switch t {
    case Pred:
        return 1
    case B8, S8, U8, E2M1, E2M3, E3M2, E4M3, E5M2, E8M0:
        return 8
    case B16, S16, U16, F16, BF16:
        return 16
    case B32, S32, U32, F32, TF32, F16x2, BF16x2:
        return 32
    case B64, S64, U64, F64:
        return 64
    case B128:
        return 128
    default:
        return 0
    }
}

// IsFloat returns true for floating-point types.
func (t Type) IsFloat() bool {
    switch t {
    case F16, F32, F64, BF16, TF32:
        return true
    default:
        return false
    }
}

// IsSigned returns true for signed integer types.
func (t Type) IsSigned() bool {
    switch t {
    case S8, S16, S32, S64:
        return true
    default:
        return false
    }
}

func (v VectorSize) String() string {
    switch v {
    case V2:
        return ".v2"
    case V4:
        return ".v4"
    default:
        return ""
    }
}