package mediatr

//Coleague
type Coleague interface {
	IsColleagueFor(RequestMessage) (bool, error)
	HandleRequest(RequestMessage) error
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
func (c callableColeague) call(r RequestMessage) error {
	is, error := c.coleague.IsColleagueFor(r)
	if error != nil {
		return error
	}

	if is {
		if error := c.coleague.HandleRequest(r); error != nil {
			return error
		}
	}

	return nil
}
