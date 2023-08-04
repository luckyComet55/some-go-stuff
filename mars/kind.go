package mars

type Kind string

const (
	ROCK   Kind = "ROCK"
	CLEFT  Kind = "CLEFT"
	DEBRIS Kind = "DEBRIS"
	LIFE   Kind = "LIFE"
	NONE   Kind = "NONE"
	ROVER  Kind = "ROVER"
)

func (k Kind) IsValid() bool {
	switch k {
	case ROCK, CLEFT, DEBRIS, LIFE, NONE, ROVER:
		return true
	default:
		return false
	}
}

func (k Kind) String() string {
	return string(k)
}
