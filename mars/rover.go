package mars

import (
	"image"

	"github.com/luckyComet55/some-go-stuff/uuid"
)

type Rover struct {
	Name string
	id   uuid.Uuid
	Occupier
}

func NewRover(name string, initPos image.Point) Rover {
	return Rover{
		Occupier: Occupier{ROVER, initPos},
		Name:     name,
		id:       uuid.NewUuid(),
	}
}
