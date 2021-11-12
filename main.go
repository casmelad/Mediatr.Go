package main

import (
	"context"
	"fmt"

	mediatr "github.com/casmelad/Mediatr.Go/pkg"
)

func main() {

	mediator := mediatr.NewMediator()

	mediator.RegisterColeague(MessageHandler{})
	mediator.RegisterTask(TaskToExecuteWithResult{})

	msgWithUUID := mediatr.NewRequestWithUUID()

	msg := MediatrRequestWrapper{
		BaseRequestMessage: *msgWithUUID,
	}

	taskParameter := TaskParameter{
		DataToProccess: DataToProccess{
			Data: "001 Code",
		},
	}

	//msg.UUID=uuid.New()
	ctx := context.Background()

	mediator.Send(ctx, msg)

	result, err := mediator.ExecuteTask(ctx, taskParameter)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		data := result.(TaskResultData)
		fmt.Println(data.Result)
	}

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

type DataToProccess struct {
	Data string
}

type TaskParameter struct {
	mediatr.TaskParameter
	DataToProccess
}

type TaskResultData struct {
	mediatr.TaskResult
	Result string
}

type TaskToExecuteWithResult struct {
}

func (o TaskToExecuteWithResult) CanExecute(params mediatr.TaskParameter) (bool, error) {
	_, ok := params.(TaskParameter)

	return ok, nil
}
func (o TaskToExecuteWithResult) Execute(ctx context.Context, params mediatr.TaskParameter) (mediatr.TaskResult, error) {
	parameters := params.(TaskParameter)

	return TaskResultData{
		Result: "Task result info: " + parameters.Data,
	}, nil
}
