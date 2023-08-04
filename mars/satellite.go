package mars

import (
	"log"
	"math/rand"
	"time"

	"github.com/luckyComet55/some-go-stuff/circlebuffer"
	"github.com/luckyComet55/some-go-stuff/uuid"
)

type Satellite struct {
	id          uuid.Uuid
	dataBuffer  *circlebuffer.CircleBuffer[string]
	dataChannel chan string
	r           *rand.Rand
}

func NewSatellite(dataChannel chan string) *Satellite {
	return &Satellite{
		id:          uuid.NewUuid(),
		dataBuffer:  circlebuffer.NewCircleBuffer[string](128),
		dataChannel: dataChannel,
		r:           rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (st *Satellite) logData(data string) {
	log.Printf("%s ST %s received ==> %s", time.Now(), st.id, data)
}

func (st *Satellite) Operate() {
	for {

		for data, err := st.dataBuffer.PopFront(); err == nil; {
			st.logData(data)
		}

		data := <-st.dataChannel
		st.logData(data)

		newTimeout := st.r.Int63() % 2000
		time.Sleep(time.Duration(newTimeout) * time.Millisecond)
	}
}
