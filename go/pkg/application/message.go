package application

import "encoding/json"

type MessageType string

const (
	REQUEST MessageType = "request"
	REPLY               = "reply"
	RELEASE             = "release"
	EXIT                = "exit"
)

type Message struct {
	Source string      `json:"source"`
	Type   MessageType `json:"type"`
	Time   float64     `json:"time"`
}

func (m *Message) Encode() []byte {
	mBytes, err := json.Marshal(m)
	if err != nil {
		mBytes = []byte("Error")
	}
	return mBytes
}

func (m *Message) Decode(data []byte) error {
	mObj := Message{}
	err := json.Unmarshal(data, &mObj)
	if err != nil {
		return err
	}
	m.Source = mObj.Source
	m.Type = mObj.Type
	m.Time = mObj.Time
	return nil
}
