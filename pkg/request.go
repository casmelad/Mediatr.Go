package mediatr

import (
	"github.com/google/uuid"
)

func prepareRequest(m *BaseRequestMessage) {
	m.UUID = uuid.New()
}

func NewRequestWithUUID() *BaseRequestMessage {
	return &BaseRequestMessage{
		UUID: uuid.New(),
	}
}

type BaseRequestMessage struct {
	UUID uuid.UUID
}

func (r BaseRequestMessage) GetUUID() uuid.UUID {
	return r.UUID
}

//Message
type RequestMessage interface {
	GetUUID() uuid.UUID
}
