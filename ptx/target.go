package ptx

import "fmt"

// Target represents a PTX target architecture (sm_XX).
type Target int

const (
    SM50 Target = iota   // Maxwell
    SM52
    SM53
    SM60                 // Pascal
    SM61
    SM62
    SM70                 // Volta
    SM72
    SM75                 // Turing
    SM80                 // Ampere
    SM86
    SM87
    SM89                 // Ada Lovelace
    SM90                 // Hopper
    SM90a
    SM100                // Blackwell
    SM101
    SM120
)

func (t Target) String() string {
    switch t {
    case SM50:
        return "sm_50"
    case SM52:
        return "sm_52"
    case SM53:
        return "sm_53"
    case SM60:
        return "sm_60"
    case SM61:
        return "sm_61"
    case SM62:
        return "sm_62"
    case SM70:
        return "sm_70"
    case SM72:
        return "sm_72"
    case SM75:
        return "sm_75"
    case SM80:
        return "sm_80"
    case SM86:
        return "sm_86"
    case SM87:
        return "sm_87"
    case SM89:
        return "sm_89"
    case SM90:
        return "sm_90"
    case SM90a:
        return "sm_90a"
    case SM100:
        return "sm_100"
    case SM101:
        return "sm_101"
    case SM120:
        return "sm_120"
    default:
        return "sm_50"
    }
}

// ISAVersion represents the PTX ISA version string.
type ISAVersion struct {
    Major int
    Minor int
}

// Common ISA versions.
var (
    ISA60 = ISAVersion{6, 0}
    ISA63 = ISAVersion{6, 3}
    ISA64 = ISAVersion{6, 4}
    ISA70 = ISAVersion{7, 0}
    ISA71 = ISAVersion{7, 1}
    ISA72 = ISAVersion{7, 2}
    ISA73 = ISAVersion{7, 3}
    ISA74 = ISAVersion{7, 4}
    ISA75 = ISAVersion{7, 5}
    ISA76 = ISAVersion{7, 6}
    ISA77 = ISAVersion{7, 7}
    ISA78 = ISAVersion{7, 8}
    ISA80 = ISAVersion{8, 0}
    ISA81 = ISAVersion{8, 1}
    ISA82 = ISAVersion{8, 2}
    ISA83 = ISAVersion{8, 3}
    ISA84 = ISAVersion{8, 4}
    ISA85 = ISAVersion{8, 5}
    ISA86 = ISAVersion{8, 6}
    ISA87 = ISAVersion{8, 7}
    ISA88 = ISAVersion{8, 8}
    ISA90 = ISAVersion{9, 0}
    ISA91 = ISAVersion{9, 1}
)

func (v ISAVersion) String() string {
    return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}