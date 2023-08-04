package mars

type Command string

const (
	STOP      Command = "STOP"
	START     Command = "START"
	RIGHT     Command = "RIGHT"
	LEFT      Command = "LEFT"
	DIRECTION Command = "DIRECTION"
)

func (c Command) IsValid() bool {
	switch c {
	case STOP, START, RIGHT, LEFT, DIRECTION:
		return true
	default:
		return false
	}
}
