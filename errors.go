package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const MaxStackDepth = 50

type BaseError struct {
	cause error
	msg   string
}

// create new error
func New(msg string) error {
	return &BaseError{
		nil,
		msg,
	}
}

func (b *BaseError) Error() string {
	if b.msg == "" && b.cause != nil {
		return b.cause.Error()
	} else {
		return b.msg
	}
}

// Extended Error

type extendedError struct {
	BaseError
	err error
}

// Create a new error extended from another error
// error.Is will now work on both this and the from error
func Extend(err error, msg string) error {
	return &extendedError{
		BaseError{
			nil,
			msg,
		},
		err,
	}
}

// errorWithStackTrace used to return errors
type errorWithStackTrace struct {
	BaseError
	err   error
	stack []uintptr
}

func internalreturn(err error, skip int, cause error, msg string) error {
	stack := make([]uintptr, MaxStackDepth)

	// 2 because to skip this line and the Return function call
	l := runtime.Callers(2+skip, stack)
	stack = stack[:l]

	return &errorWithStackTrace{
		BaseError{
			cause,
			msg,
		},
		err,
		stack,
	}
}

// captures the callstack and returns an error
// optionaly add details or information to the error
func Return(err error, cause error, msg string) error {
	return internalreturn(err, 1, cause, msg)
}

// short form for return with template and args
func Returnf(e error, cause error, args ...any) error {
	var str string
	if ee, ok := e.(*errorWithStackTrace); ok {
		str = ee.sprintf(args...)
	}
	str = ""
	return internalreturn(e, 1, cause, str)
}

func (e *errorWithStackTrace) Error() string {
	var sb strings.Builder

	msg := e.BaseError.Error()
	if e.err != nil {
		sb.WriteString(e.err.Error())
		if msg != "" {
			sb.WriteString(": ")
		}
	}

	if msg != "" {
		sb.WriteString(msg)
	}

	return sb.String()
}

func (e *errorWithStackTrace) sprintf(args ...any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *errorWithStackTrace) Is(other error) bool {
	if e.err != nil {
		return e.err == other
	} else {
		if x, ok := other.(*errorWithStackTrace); ok && e == x {
			return true
		}
	}
	return false
}

func (e *errorWithStackTrace) Unwrap() error {
	return e.err
}
