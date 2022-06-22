package mediatr

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrHandlerNotFound error = errors.New("handler not found")
)

type Mediatr struct {
	types            map[string]CallableColleague
	tasksProccessors []*CallableTask
}

//RegisterColeagueForMessage - register a coleague to handle the message specified
func RegisterColeagueForMessage[T any](m *Mediatr, c CallableColleague, msg T) error {
	m.types[fmt.Sprintf("%T", msg)] = c
	return nil
}

//RegisterTask - register a handler to execute a task and return a result
func RegisterTask(m *Mediatr, t CallableTask) error {

	var task CallableTask = t

	m.tasksProccessors = append(m.tasksProccessors, &task)

	return nil
}

//SendMsg - sends the message to the correct coleague if exists in the registered coleagues collection
func SendMsg[T any](ctx context.Context, m *Mediatr, msg T) error {

	col, is := m.types[fmt.Sprintf("%T", msg)]
	if !is {
		return ErrHandlerNotFound
	}

	colleague, is := col.(Colleague[T])
	if !is {
		return ErrHandlerNotFound
	}

	err := colleague.Receive(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

//ExeuteTask - Look up the handler to execute the task and returns the expected result
func ExecuteTask[T any, U any](ctx context.Context, m *Mediatr, params T) (U, error) {
	var result U

	for _, task := range m.tasksProccessors {
		tp := *task
		if t, is := tp.(Task[T, U]); is {

			result, err := t.Execute(ctx, params)

			if err == nil {
				return result, err
			}
		}
		return result, nil
	}
	return result, ErrHandlerNotFound
}

func NewMediator() *Mediatr {
	return &Mediatr{
		types: map[string]CallableColleague{},
	}
}
