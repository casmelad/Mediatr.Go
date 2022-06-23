package mediatr

import (
	"context"
	"fmt"

	mediatr "github.com/casmelad/Mediatr.Go/pkg"
)

func main() {

	m := mediatr.NewMediator()
	mediatr.RegisterHandler[HandlerData](m, EventHandler{})
	mediatr.RegisterTask[TaskParameters, TaskResult](m, TaskToExecuteWithResult{})

	msg2 := HandlerData{
		Greeting: "Event 1",
	}

	taskParameter := TaskParameters{
		Data: "001 Code",
	}

	ctx := context.Background()

	mediatr.SendMsg(ctx, m, msg2)
	result, err := mediatr.ExecuteTask[TaskParameters, TaskResult](ctx, m, taskParameter)

	if err != nil {
		fmt.Println(err)
	} else {
		data := result
		fmt.Println(data.Result)
	}

}

type HandlerData struct {
	Greeting string
}

type EventHandler struct{}

func (h EventHandler) Receive(ctx context.Context, data HandlerData) error {
	fmt.Println(fmt.Sprintf("message uuid and event data %s", data.Greeting))
	return nil
}

type TaskParameters struct {
	Data string
}

type TaskResult struct {
	Result string
}

type TaskToExecuteWithResult struct {
}

func (o TaskToExecuteWithResult) Execute(ctx context.Context, params TaskParameters) (TaskResult, error) {
	return TaskResult{
		Result: "Task result info: " + params.Data,
	}, nil
}
