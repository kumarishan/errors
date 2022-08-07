package main

import (
	"fmt"

	"github.com/kumarishan/errors"
)

var ErrInternalError = errors.New("internal error")
var ErrInvalidInput = errors.New("invalid input")
var ErrMissingIdParameter = errors.Extend(ErrInvalidInput, "missing id parameter")

func A() error {
	return errors.Return(ErrInvalidInput, nil, "")
}

func B() error {
	err := A()

	if errors.Is(err, ErrInvalidInput) {
		return errors.Return(ErrMissingIdParameter, err, "overriden error message")
	}

	return errors.Return(ErrInternalError, nil, "")

}

func main() {
	err := B()
	fmt.Printf("Got error: %+v\n", err)
}
