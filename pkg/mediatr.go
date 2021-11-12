package mediatr

import "context"

type Mediatr struct {
	coleagues        []*callableColeague
	tasksProccessors []*taskProccessor
}

func (m *Mediatr) RegisterColeague(c Coleague) error {
	if m.coleagues == nil {
		m.coleagues = []*callableColeague{}
	}

	m.coleagues = append(m.coleagues, newCallableColeague(c))

	return nil
}

func (m *Mediatr) RegisterTask(t Task) error {
	if m.tasksProccessors == nil {
		m.tasksProccessors = []*taskProccessor{}
	}

	m.tasksProccessors = append(m.tasksProccessors, newTaskProccessor(t))

	return nil
}

func (m *Mediatr) Send(ctx context.Context, msg RequestMessage) error {

	for _, col := range m.coleagues {
		err := col.call(ctx, msg)
		if err != nil {
			return err
		}
	}

	return nil
}

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
