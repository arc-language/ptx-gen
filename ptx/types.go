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

	// Packed types (Legacy/Common)
	F16x2
	BF16x2

	// Fixed-point (sm_100+)
	E2M1
	E2M3
	E3M2
	E4M3
	E5M2
	E8M0

	// --- New Types added in PTX 9.1 / sm_100+ ---

	// Additional Scalar types
	UE4M3 // .ue4m3 (7-bit unsigned float, stored in .b8)
	S2F6  // .s2f6 (8-bit signed fixed-point)

	// Packed Floating Point & Fixed Point (Table 9)
	F32x2   // .f32x2
	E4M3x2  // .e4m3x2
	E5M2x2  // .e5m2x2
	E2M3x2  // .e2m3x2
	E3M2x2  // .e3m2x2
	UE8M0x2 // .ue8m0x2
	S2F6x2  // .s2f6x2
	E2M1x2  // .e2m1x2
	E4M3x4  // .e4m3x4
	E5M2x4  // .e5m2x4
	E2M3x4  // .e2m3x4
	E3M2x4  // .e3m2x4
	E2M1x4  // .e2m1x4

	// Packed Integer
	U16x2 // .u16x2
	S16x2 // .s16x2

	// Tensor Sub-byte types (Section 5.5.1.1)
	B4x16     // .b4x16
	B4x16_p64 // .b4x16_p64
	B6x16_p32 // .b6x16_p32
	B6p2x16   // .b6p2x16

	// Opaque Types (Section 5.3 & 5.5.8)
	TexRef     // .texref
	SamplerRef // .samplerref
	SurfRef    // .surfref
	TensorMap  // .tensormap
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
	case UE4M3:
		return ".ue4m3"
	case S2F6:
		return ".s2f6"
	case F32x2:
		return ".f32x2"
	case E4M3x2:
		return ".e4m3x2"
	case E5M2x2:
		return ".e5m2x2"
	case E2M3x2:
		return ".e2m3x2"
	case E3M2x2:
		return ".e3m2x2"
	case UE8M0x2:
		return ".ue8m0x2"
	case S2F6x2:
		return ".s2f6x2"
	case E2M1x2:
		return ".e2m1x2"
	case E4M3x4:
		return ".e4m3x4"
	case E5M2x4:
		return ".e5m2x4"
	case E2M3x4:
		return ".e2m3x4"
	case E3M2x4:
		return ".e3m2x4"
	case E2M1x4:
		return ".e2m1x4"
	case U16x2:
		return ".u16x2"
	case S16x2:
		return ".s16x2"
	case B4x16:
		return ".b4x16"
	case B4x16_p64:
		return ".b4x16_p64"
	case B6x16_p32:
		return ".b6x16_p32"
	case B6p2x16:
		return ".b6p2x16"
	case TexRef:
		return ".texref"
	case SamplerRef:
		return ".samplerref"
	case SurfRef:
		return ".surfref"
	case TensorMap:
		return ".tensormap"
	default:
		return ".unknown"
	}
}

// BitWidth returns the width in bits for the type.
// For packed types, this is the total width of the container.
// For opaque types, this returns the typical handle/structure size.
func (t Type) BitWidth() int {
	switch t {
	case Pred:
		return 1
	case B8, S8, U8, E2M1, E2M3, E3M2, E4M3, E5M2, E8M0, UE4M3, S2F6:
		return 8
	case E2M1x2:
		return 8 // 2 * 4 bits
	case B16, S16, U16, F16, BF16:
		return 16
	case E4M3x2, E5M2x2, E2M3x2, E3M2x2, UE8M0x2, S2F6x2, E2M1x4:
		return 16
	case B32, S32, U32, F32, TF32, F16x2, BF16x2:
		return 32
	case U16x2, S16x2, E4M3x4, E5M2x4, E2M3x4, E3M2x4:
		return 32
	case B64, S64, U64, F64, F32x2:
		return 64
	case B4x16:
		return 64 // 16 * 4 bits
	case B128:
		return 128
	case B4x16_p64, B6x16_p32, B6p2x16:
		return 128
	case TexRef, SamplerRef, SurfRef:
		return 64 // Opaque handles are usually pointers/u64
	case TensorMap:
		return 1024 // 128 bytes
	default:
		return 0
	}
}

// IsFloat returns true for floating-point types (including packed and alternate floats).
func (t Type) IsFloat() bool {
	switch t {
	case F16, F32, F64, BF16, TF32,
		F16x2, BF16x2, F32x2,
		E4M3, E5M2, E2M3, E3M2, E2M1, UE4M3,
		E4M3x2, E5M2x2, E2M3x2, E3M2x2, UE8M0x2,
		E4M3x4, E5M2x4, E2M3x4, E3M2x4:
		return true
	default:
		return false
	}
}

// IsSigned returns true for signed integer types and signed fixed-point types.
func (t Type) IsSigned() bool {
	switch t {
	case S8, S16, S32, S64, S16x2, S2F6, S2F6x2:
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