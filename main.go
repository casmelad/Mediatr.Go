package mediatr

/* package main

import (
	"context"
	"fmt"

	mediatr "github.com/casmelad/Mediatr.Go/pkg"
)

func main() {

	mediator := mediatr.NewMediator()
	mediator.RegisterColeagueForMessage(EventHandler{}, EventData{})
	mediator.RegisterTask(TaskToExecuteWithResult{})

	msgWithUUID := mediatr.NewRequestWithUUID()

	msg2 := EventData{
		BaseRequestMessage: *msgWithUUID,
		Greeting:           "Event 1",
	}

	msg3 := EventData{
		BaseRequestMessage: *msgWithUUID,
		Greeting:           "Event 2",
	}

	taskParameter := TaskParameter{
		DataToProccess: DataToProccess{
			Data: "001 Code",
		},
	}

	//msg.UUID=uuid.New()
	ctx := context.Background()
	mediator.SendMsg(ctx, msg2)
	mediator.SendMsg(ctx, msg3)

	result, err := mediator.ExecuteTask(ctx, taskParameter)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		data := result.(TaskResultData)
		fmt.Println(data.Result)
	}

}

type EventData struct {
	mediatr.BaseRequestMessage
	Greeting string
}

type EventHandler struct{}

func (h EventHandler) Receive(ctx context.Context, r mediatr.RequestMessage) error {
	data := r.(EventData)
	fmt.Println(fmt.Sprintf("message uuid and event data %s", data.Greeting))
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
*/
