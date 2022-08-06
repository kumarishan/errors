package main

import (
	"fmt"
	"kumarishan/errors"
	"runtime/debug"
)

var Err = errors.New("new error")

func A() error {
	debug.PrintStack()
	return errors.Return(Err, nil, "some error")
}

func B() error {
	return A()
}

func C() error {
	return B()
}

func main() {
	err := C()
	fmt.Print("\n\nError Stack trace\n")
	fmt.Println(errors.StackTrace(err))
	fmt.Println(errors.StackTrace(Err))
}
