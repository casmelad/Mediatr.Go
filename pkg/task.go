package mediatr

import "context"

type Task[T any, U any] interface {
	callableTask
	Execute(context.Context, T) (U, error)
}

type callableTask interface{}
