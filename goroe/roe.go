package goroe

import "fmt"

type ResultOrError struct {
	result interface{}
	err    error
}

func (roe ResultOrError) Unwrap() (interface{}, error) {
	if roe.err != nil {
		return nil, roe.err
	} else {
		return roe.result, nil
	}
}

func NewResultError(res interface{}, err error) ResultOrError {
	return ResultOrError{result: res, err: err}
}

func Result(res interface{}) ResultOrError {
	return ResultOrError{result: res}
}

func Errorf(format string, args ...interface{}) ResultOrError {
	return ResultOrError{err: fmt.Errorf(format, args...)}
}

func Error(err error) ResultOrError {
	return ResultOrError{err: err}
}
