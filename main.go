package main

import (
	"context"
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
	ctx := context.Background()

	mediator.Send(ctx, msg)
	mediator.Send(ctx, msg)

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

func (h MessageHandler) Receive(ctx context.Context, r mediatr.RequestMessage) error {

	data := r.(MediatrRequestWrapper)

	fmt.Println(fmt.Printf("message uuid %s and event data %s", data.GetUUID(), data.EventInfo))

	return nil
}
