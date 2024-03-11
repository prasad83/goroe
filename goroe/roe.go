package goroe

import "fmt"

type ResultOrError[T any] struct {
	result T
	err    error
}

func (roe ResultOrError[T]) Unwrap() (T, error) {
	return roe.result, roe.err
}

func NewResultError[T any](res T, err error) ResultOrError[T] {
	return ResultOrError[T]{result: res, err: err}
}

func Result[T any](res T) ResultOrError[T] {
	return ResultOrError[T]{result: res}
}

func Errorf[T any](format string, args ...interface{}) ResultOrError[T] {
	return ResultOrError[T]{err: fmt.Errorf(format, args...)}
}

func Error[T any](err error) ResultOrError[T] {
	return ResultOrError[T]{err: err}
}
