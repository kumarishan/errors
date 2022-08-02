package errors

import (
	"fmt"
	"strings"
)

type Error struct {
	err   error
	cause error
	msg   string
}

// captures the callstack and returns an error
// optionaly add details or information to the error
func Return(e error, cause error, msg string) error {
	return &Error{e, cause, msg}
}

// short form for return with template and args
func Returnf(e error, cause error, args ...any) error {
	var str string
	if ee, ok := e.(*Error); ok {
		str = ee.sprintf(args...)
	}
	str = ""
	return Return(e, cause, str)
}

// create new error
func New(msg string) error {
	return &Error{nil, nil, msg}
}

// Create a new error extended from another error
// error.Is will now work on both this and the from error
func Extend(err error, msg string) error {
	return &Error{
		err,
		nil,
		msg,
	}
}

func (e *Error) Error() string {
	var sb strings.Builder

	if e.err != nil {
		sb.WriteString(e.err.Error())
		if e.msg != "" {
			sb.WriteString(": ")
		}
	}

	if e.msg != "" {
		sb.WriteString(e.msg)
	}

	return sb.String()
}

func (e *Error) sprintf(args ...any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Is(other error) bool {
	if e.err != nil {
		return e.err == other
	} else {
		if x, ok := other.(*Error); ok && e == x {
			return true
		}
	}
	return false
}

func (e *Error) Unwrap() error {
	return e.err
}
