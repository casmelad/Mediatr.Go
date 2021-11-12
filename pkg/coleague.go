package mediatr

import "context"

//Coleague
type Coleague interface {
	IsColleagueFor(RequestMessage) (bool, error)
	Receive(context.Context, RequestMessage) error
}

type Task interface {
	CanExecute(TaskParameter) (bool, error)
	Execute(context.Context, TaskParameter) (TaskResult, error)
}

type Result struct {
	Error error
	RequestMessage
}

type callableColeague struct {
	coleague Coleague
}

func newCallableColeague(c Coleague) *callableColeague {
	return &callableColeague{
		coleague: c,
	}
}

//call is a template method to handle the request
func (c callableColeague) call(ctx context.Context, r RequestMessage) error {
	is, err := c.coleague.IsColleagueFor(r)
	if err != nil {
		return err
	}

	if is {
		if err := c.coleague.Receive(ctx, r); err != nil {
			return err
		}
	}

	return nil
}

type taskProccessor struct {
	task Task
}

func newTaskProccessor(t Task) *taskProccessor {
	return &taskProccessor{task: t}
}

func (t taskProccessor) execute(ctx context.Context, param TaskParameter) (TaskResult, error) {
	can, err := t.task.CanExecute(param)

	if err != nil {
		return nil, err
	}

	if can {
		return t.task.Execute(ctx, param)
	}

	return nil, nil
}
