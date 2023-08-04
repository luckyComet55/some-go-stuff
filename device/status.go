package device

type Status string

const (
	ONLINE     Status = "ONLINE"
	LAG        Status = "LAG"
	OFFLINE    Status = "OFFLINE"
	CONNECTING Status = "CONNECTING"
)

func (s Status) IsValid() bool {
	switch s {
	case ONLINE, LAG, OFFLINE, CONNECTING:
		return true
	default:
		return false
	}
}

func (s Status) String() string {
	return string(s)
}
