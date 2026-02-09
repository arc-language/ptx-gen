package ptx

// RoundingMode represents PTX rounding modifiers.
type RoundingMode int

const (
    RoundNone RoundingMode = iota

    // Floating-point rounding
    RoundNearestEven  // .rn — round to nearest even (default)
    RoundZero         // .rz — round towards zero
    RoundNegInf       // .rm — round towards negative infinity
    RoundPosInf       // .rp — round towards positive infinity

    // Integer rounding (for cvt float→int)
    RoundIntNearestEven  // .rni
    RoundIntZero         // .rzi
    RoundIntNegInf       // .rmi
    RoundIntPosInf       // .rpi
)

func (r RoundingMode) String() string {
    switch r {
    case RoundNone:
        return ""
    case RoundNearestEven:
        return ".rn"
    case RoundZero:
        return ".rz"
    case RoundNegInf:
        return ".rm"
    case RoundPosInf:
        return ".rp"
    case RoundIntNearestEven:
        return ".rni"
    case RoundIntZero:
        return ".rzi"
    case RoundIntNegInf:
        return ".rmi"
    case RoundIntPosInf:
        return ".rpi"
    default:
        return ""
    }
}