package uuid

type Uuid string

func (uuid Uuid) String() string {
	return string(uuid)
}
