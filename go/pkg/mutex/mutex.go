package mutex

import (
	"math"
	"time"
)

type Mutex struct {
	id    string
	eTime float64
}

func New(id string, t float64) (*Mutex, error) {
	mutex := Mutex{
		id:    id,
		eTime: t,
	}
	return &mutex, nil
}

func (m *Mutex) Execute_CS() {
	currClock := math.Floor(time.Since(time.Now()).Seconds()*1000000) / 1000000
	executionTime := currClock + m.eTime
	time.Sleep(time.Duration(executionTime) * time.Second)
}
