package ptx

// Linkage represents PTX visibility/linkage directives.
type Linkage int

const (
    LinkNone    Linkage = iota
    LinkVisible                  // .visible — externally visible
    LinkExtern                   // .extern  — declared externally
    LinkWeak                     // .weak    — weak linkage
    LinkCommon                   // .common  — common symbol (multiple definitions allowed)
)

func (l Linkage) String() string {
    switch l {
    case LinkNone:
        return ""
    case LinkVisible:
        return ".visible"
    case LinkExtern:
        return ".extern"
    case LinkWeak:
        return ".weak"
    case LinkCommon:
        return ".common"
    default:
        return ""
    }
}