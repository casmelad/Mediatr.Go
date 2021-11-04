package mediatr

import "context"

type Mediatr struct {
	coleagues []*callableColeague
}

func (m *Mediatr) RegisterColeague(c Coleague) error {
	if m.coleagues == nil {
		m.coleagues = []*callableColeague{}
	}

	m.coleagues = append(m.coleagues, newCallableColeague(c))

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

func NewMediator() *Mediatr {
	return &Mediatr{}
}
