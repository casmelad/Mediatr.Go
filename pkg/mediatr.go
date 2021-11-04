package mediatr

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

func (m *Mediatr) Send(msg RequestMessage) error {

	for _, col := range m.coleagues {
		err := col.call(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMediator() *Mediatr {
	return &Mediatr{}
}
