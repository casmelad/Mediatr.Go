package mediatr

import "context"

//Coleague
type Colleague[T any] interface {
	callableColleague
	Receive(context.Context, T) error
}

type callableColleague interface{}
