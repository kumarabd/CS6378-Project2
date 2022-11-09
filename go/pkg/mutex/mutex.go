package mutex

import (
	"time"
)

type Mutex struct {
	id    string
	eTime int64
}

func New(id string, t int64) (*Mutex, error) {
	mutex := Mutex{
		id:    id,
		eTime: t,
	}
	return &mutex, nil
}

func (m *Mutex) Execute_CS() {
	//currTime := time.Now().UnixNano() / int64(time.Millisecond)
	//currClock := time.Since(currTime).Milliseconds()
	//executionTime := currClock + m.eTime
	time.Sleep(time.Duration(m.eTime) * time.Millisecond)
}
