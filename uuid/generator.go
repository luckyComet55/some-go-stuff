package uuid

import (
	"math/rand"
	"time"
)

const length = 20
const alph = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-="

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func NewUuid() Uuid {
	b := make([]byte, length)
	for i := range b {
		b[i] = alph[r.Int63()%int64(len(alph))]
	}
	return Uuid(b)
}
