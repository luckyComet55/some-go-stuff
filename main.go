package main

import (
	"image"

	"github.com/luckyComet55/some-go-stuff/mars"
)

func main() {
	mainDataChannel := make(chan string)
	st := mars.NewSatellite(mainDataChannel)
	r1 := mars.NewRover("PIVO", image.Point{0, 0})
	r2 := mars.NewRover("VODKA", image.Point{50, 50})
	r3 := mars.NewRover("SAMOGON", image.Point{70, 70})
	r4 := mars.NewRover("Le Tequilla", image.Point{90, 90})
	rd1 := mars.NewRoverDriver(r1)
	rd2 := mars.NewRoverDriver(r2)
	rd3 := mars.NewRoverDriver(r3)
	rd4 := mars.NewRoverDriver(r4)

	mg := mars.NewMarsGrid(100, 100)
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{17, 24}))
	mg.SetOccupier(mars.NewOccupier(mars.CLEFT, image.Point{11, 0}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{90, 35}))
	mg.SetOccupier(mars.NewOccupier(mars.DEBRIS, image.Point{70, 51}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{12, 24}))
	mg.SetOccupier(mars.NewOccupier(mars.ROCK, image.Point{0, 57}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{63, 60}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{90, 91}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{12, 33}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{1, 43}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{2, 6}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{1, 8}))
	mg.SetOccupier(mars.NewOccupier(mars.LIFE, image.Point{98, 70}))

	go rd1.Drive(mg, mainDataChannel)
	go rd2.Drive(mg, mainDataChannel)
	go rd3.Drive(mg, mainDataChannel)
	go rd4.Drive(mg, mainDataChannel)
	go st.Operate()
	for {

	}
}
