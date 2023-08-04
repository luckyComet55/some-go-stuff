package mars

import (
	"fmt"
	"image"
)

type Occupier struct {
	_Kind Kind
	Pos   image.Point
}

func NewOccupier(k Kind, initPos image.Point) Occupier {
	return Occupier{k, initPos}
}

func (o *Occupier) Move(p image.Point, mg *MarsGrid) bool {
	res := mg.Occupy(p, *o)

	if !res {
		return false
	}

	o.Pos = p
	return true
}

func (o Occupier) String() string {
	return fmt.Sprintf("%s AT %s", o._Kind, o.Pos)
}
