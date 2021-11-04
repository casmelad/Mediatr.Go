package main

import (
	"fmt"

	mediatr "github.com/casmelad/Mediatr.Go/pkg"
)

func main() {

	mediator := mediatr.NewMediator()

	mediator.RegisterColeague(MessageHandler{})

	msgWithUUID := mediatr.NewRequestWithUUID()

	msg := MediatrRequestWrapper{
		BaseRequestMessage: *msgWithUUID,
	}

	//msg.UUID=uuid.New()

	mediator.Proccess(msg)
	mediator.Proccess(msg)

}

type MessageToHandle struct {
	EventInfo string
}

type MediatrRequestWrapper struct {
	mediatr.BaseRequestMessage
	MessageToHandle
}

type MessageHandler struct{}

func (h MessageHandler) IsColleagueFor(r mediatr.RequestMessage) (bool, error) {

	_, ok := r.(MediatrRequestWrapper)

	return ok, nil
}

func (h MessageHandler) HandleRequest(r mediatr.RequestMessage) error {

	data := r.(MediatrRequestWrapper)

	fmt.Println(fmt.Printf("message uuid %s and event data %s", data.GetUUID(), data.EventInfo))

	return nil
}
