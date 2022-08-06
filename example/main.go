package main

import (
	"fmt"

	"github.com/kumarishan/errors"
)

var AErr = errors.New("AErr")

func A() error {
	return errors.Return(AErr, nil, "some error occurred in A")
}

var BErr = errors.New("BErr")

func B() error {
	err := A()
	return errors.Return(BErr, err, "some error occured in B")
}

func C() error {
	return B()
}

func main() {
	err := C()
	fmt.Printf("Got error: %v\n", err)
	fmt.Printf("Got error with stacktrace: %+v\n", err)
}
