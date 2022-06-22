package mediatr

import "context"

type Task[T any, U any] interface {
	CallableTask
	Execute(context.Context, T) (U, error)
}

type CallableTask interface{}
