package errors

import (
	"runtime"
)

const MaxStackDepth = 50

type BaseError struct {
	msg string
}

// Define a new error
func New(msg string) error {
	return &BaseError{
		msg,
	}
}

func (b *BaseError) Error() string {
	return b.msg
}

// Extended Error

type extendedError struct {
	BaseError

	err error
}

// Create a new error extended from another error
// error.Is will now work on both this and the from error
func Extend(err error, msg string) error {
	if err == nil {
		panic("errors: cannot extend from nil")
	}

	return &extendedError{
		BaseError{
			msg,
		},
		err,
	}
}

func (e *extendedError) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return e.err.Error()
}

func (e *extendedError) Is(other error) bool {
	if e.err == other {
		return true

	} else if x, ok := other.(*extendedError); ok && e == x {
		return true
	}

	return false
}

func (e *extendedError) Unwrap() error {
	return e.err
}

// errorWrapper is used to wrap errors for return
// it captures call stack for stacktrace and cause
// overrides error messages
type errorWrapper struct {
	err   error
	msg   string
	cause error

	stack []uintptr
}

func (e *errorWrapper) Error() string {
	if e.msg != "" {
		return e.msg
	}

	if e.err.Error() != "" {
		return e.err.Error()
	}

	if e.cause != nil {
		return e.cause.Error()
	}

	return ""
}

func internalreturn(err error, skip int, cause error, msg string) error {
	stack := make([]uintptr, MaxStackDepth)

	// 2 because to skip this line and the Return function call
	l := runtime.Callers(2+skip, stack)
	stack = stack[:l]

	if err == nil {
		err = &BaseError{msg}
	}

	return &errorWrapper{
		err,
		msg,
		cause,
		stack,
	}
}

// captures the callstack and returns an error
// optionaly add details or information to the error
func Return(err error, cause error, msg string) error {
	return internalreturn(err, 1, cause, msg)
}

func (e *errorWrapper) Is(other error) bool {
	if e.err != nil {
		return e.err == other
	} else {
		if x, ok := other.(*errorWrapper); ok && e == x {
			return true
		}
	}
	return false
}

func (e *errorWrapper) Unwrap() error {
	return e.err
}
