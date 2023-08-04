package mars

import (
	"image"
	"sync"
)

type MarsGrid struct {
	grid [][]Occupier
	mtx  sync.Mutex
	X    int
	Y    int
}

func NewMarsGrid(sizeX, sizeY int) *MarsGrid {
	g := make([][]Occupier, sizeX)
	for i := range g {
		g[i] = make([]Occupier, sizeY)
	}
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			g[i][j] = NewOccupier(NONE, image.Point{i, j})
		}
	}
	return &MarsGrid{
		grid: g,
		mtx:  sync.Mutex{},
		X:    sizeX,
		Y:    sizeY,
	}
}

func (mg *MarsGrid) isPointOnGrid(p image.Point) bool {
	return (p.X < mg.X && p.Y < mg.Y) && (p.X >= 0 && p.Y >= 0)
}

func (mg *MarsGrid) Occupy(p image.Point, o Occupier) bool {
	mg.mtx.Lock()
	defer mg.mtx.Unlock()

	if !mg.isPointOnGrid(p) {
		return false
	}

	switch mg.grid[p.X][p.Y]._Kind {
	case NONE:
		mg.grid[p.X][p.Y] = o
		return true
	// TODO: Add other cases
	default:
		return false
	}

}

func (mg *MarsGrid) SetOccupier(o Occupier) {
	if mg.isPointOnGrid(o.Pos) {
		mg.grid[o.Pos.X][o.Pos.Y] = o
	}
}

func (mg *MarsGrid) GetOccupier(p image.Point) *Occupier {
	mg.mtx.Lock()
	defer mg.mtx.Unlock()

	if !mg.isPointOnGrid(p) {
		return nil
	}

	return &mg.grid[p.X][p.Y]
}
