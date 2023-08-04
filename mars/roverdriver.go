package mars

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"
)

type RoverDriver struct {
	rover       Rover
	Direction   image.Point
	commandChan chan Command
	r           *rand.Rand
}

func NewRoverDriver(rover Rover) RoverDriver {
	return RoverDriver{
		rover:       rover,
		Direction:   image.Point{1, 0},
		commandChan: make(chan Command),
		r:           rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rd *RoverDriver) logInfo(msg string) {
	log.Printf("Rover %s ==> %s\n", rd.rover.id, msg)
}

func (rd *RoverDriver) Left() {
	rd.commandChan <- LEFT
}

func (rd *RoverDriver) Right() {
	rd.commandChan <- RIGHT
}

func (rd *RoverDriver) Stop() {
	rd.commandChan <- STOP
}

func (rd *RoverDriver) Start() {
	rd.commandChan <- START
}

func (rd *RoverDriver) ChangeDir() {
	rd.commandChan <- DIRECTION
}

func (rd *RoverDriver) getRandomDirection() image.Point {
	randR := rd.r.Int63n(4)
	randL := rd.r.Int63n(4)

	nextDir := rd.Direction

	for i := 0; i < int(randR); i++ {
		nextDir = image.Point{
			X: -nextDir.Y,
			Y: nextDir.X,
		}
	}
	for i := 0; i < int(randL); i++ {
		nextDir = image.Point{
			X: nextDir.Y,
			Y: -nextDir.X,
		}
	}

	return nextDir
}

func (rd *RoverDriver) sendData(data string, dataChannel chan string) {
	dataChannel <- data
}

func (rd *RoverDriver) Drive(mg *MarsGrid, dataChannel chan string) {

	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-rd.commandChan:
			switch c {
			case RIGHT:
				rd.Direction = image.Point{
					X: -rd.Direction.X,
					Y: rd.Direction.X,
				}
			case LEFT:
				rd.Direction = image.Point{
					X: rd.Direction.Y,
					Y: -rd.Direction.X,
				}
			}
		case <-nextMove:
			nextPoint := rd.rover.Pos.Add(rd.Direction)
			res := rd.rover.Move(nextPoint, mg)

			if !res {
				obstacle := mg.GetOccupier(nextPoint)
				if obstacle == nil {
					rd.logInfo("unable to move forward, end of grid")
				} else {
					if obstacle._Kind == LIFE {
						rd.sendData(obstacle.String(), dataChannel)
					}
					rd.logInfo(fmt.Sprintf("unable to move forward, obstacle %s", obstacle._Kind))
				}
				rd.Direction = rd.getRandomDirection()
			} else {
				rd.logInfo(fmt.Sprintf("moved to [%d, %d]", rd.rover.Pos.X, rd.rover.Pos.Y))
			}

			nextMove = time.After(updateInterval)
		}
	}

}
