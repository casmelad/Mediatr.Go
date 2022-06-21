package mediatr

import (
	"context"
	"fmt"
)

type Mediatr struct {
	types            []coleagueEntry
	safeTypes        []safeTypeColeagueEntry
	tasksProccessors []*taskProccessor
}

type coleagueEntry struct {
	coleague Coleague
	msg      RequestMessage
}

type safeTypeColeagueEntry struct {
	coleague SafeTypeColeague
	msg      RequestMessage
}

//RegisterColeagueForMessage - register a coleague to handle the message specified
func (m *Mediatr) RegisterColeagueForMessage(c Coleague, msg RequestMessage) error {

	colEntry := coleagueEntry{
		coleague: c,
		msg:      msg,
	}
	m.types = append(m.types, colEntry)
	return nil
}

//RegisterColeagueForMessage - register a coleague to handle the message specified
func (m *Mediatr) RegisterSafeTypeColeagueForMessage(c SafeTypeColeague, msg RequestMessage) error {

	colEntry := safeTypeColeagueEntry{
		coleague: c,
		msg:      msg,
	}
	m.safeTypes = append(m.safeTypes, colEntry)
	return nil
}

//RegisterTask - register a handler to execute a task and return a result
func (m *Mediatr) RegisterTask(t Task) error {

	m.tasksProccessors = append(m.tasksProccessors, newTaskProccessor(t))

	return nil
}

//SendMsg - sends the message to the correct coleague if exists in the registered coleagues collection
func (m *Mediatr) SendMsg(ctx context.Context, msg RequestMessage) error {
	for _, col := range m.types {
		if fmt.Sprintf("%T", msg) == fmt.Sprintf("%T", col.msg) {
			err := col.coleague.Receive(ctx, msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//SendMsg - sends the message to the correct coleague if exists in the registered coleagues collection
func (m *Mediatr) SendMsgWithSafeType(ctx context.Context, params RequestMessage) (TaskResult, error) {

	for _, task := range m.tasksProccessors {
		result, err := task.execute(ctx, params)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, nil
}

//ExeuteTask - Look up the handler to execute the task and returns the expected result
func (m *Mediatr) ExecuteTask(ctx context.Context, params TaskParameter) (TaskResult, error) {

	for _, task := range m.tasksProccessors {
		result, err := task.execute(ctx, params)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, nil
}

func NewMediator() *Mediatr {
	return &Mediatr{}
}
