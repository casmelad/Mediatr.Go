package mediatr

import "context"

type Task interface {
	CanExecute(TaskParameter) (bool, error)
	Execute(context.Context, TaskParameter) (TaskResult, error)
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
