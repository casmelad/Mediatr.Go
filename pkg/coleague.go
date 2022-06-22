package mediatr

import "context"

//Coleague
type Colleague[T any] interface {
	CallableColleague
	Receive(context.Context, T) error
}

type CallableColleague interface{}
