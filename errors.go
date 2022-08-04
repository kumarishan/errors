package errors

import (
	"errors"
	"fmt"
	"strings"
)

type BaseError struct {
	cause error
	msg   string
}

func (b *BaseError) Error() string {
	if b.msg == "" && b.cause != nil {
		return b.cause.Error()
	} else {
		return b.msg
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

// errorWrapper used to return errors
// supports errors.Is
type errorWrapper struct {
	BaseError
	err error
}

// captures the callstack and returns an error
// optionaly add details or information to the error
func Return(err error, cause error, msg string) error {
	return &errorWrapper{
		BaseError{
			cause,
			msg,
		},
		err,
	}
}

// short form for return with template and args
func Returnf(e error, cause error, args ...any) error {
	var str string
	if ee, ok := e.(*errorWrapper); ok {
		str = ee.sprintf(args...)
	}
	str = ""
	return Return(e, cause, str)
}

// create new error
func New(msg string) error {
	return &errorWrapper{
		BaseError{
			nil,
			msg,
		},
		nil,
	}
}

// Create a new error extended from another error
// error.Is will now work on both this and the from error
func Extend(err error, msg string) error {
	return &errorWrapper{
		BaseError{
			nil,
			msg,
		},
		err,
	}
}

func (e *errorWrapper) Error() string {
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

func (e *errorWrapper) sprintf(args ...any) string {
	return fmt.Sprintf(e.msg, args...)
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
