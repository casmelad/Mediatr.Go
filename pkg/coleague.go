package mediatr

import "context"

//Coleague
type Coleague interface {
	Receive(context.Context, RequestMessage) error
}

type SafeTypeColeague interface {
	IsColleagueFor(RequestMessage) (bool, error)
	Receive(context.Context, RequestMessage) error
}
