package errors

import (
	"fmt"
	"io"
	"runtime"
)

// wrapper is used to wrap errors for return
// it captures call stack for stacktrace and cause
// overrides error messages
type wrapper struct {
	err   error
	msg   string
	cause error

	stack stack
}

func internalreturn(err error, skip int, cause error, msg string
	type string) error {
	stack := make(stack, MaxStackDepth)

	// 2 because to skip this line and the Return function call
	l := runtime.Callers(2+skip, stack)
	stack = stack[:l]

	if err == nil {
		err = &BaseError{msg}
	}

	if cause != nil {
		cause = pruned(cause, stack)
	}

	return &wrapper{
		err,
		msg,
		cause,
		stack,
	}
}

func pruned(cause error, stack stack) error {
	if w, ok := cause.(*wrapper); ok {
		ws := w.stack

		var li int = len(ws)
		for i := len(ws) - 1; i >= 0; i-- {
			var found bool
			for j := len(stack) - 1; j >= 0; j-- {
				if ws[i] == stack[j] {
					found = true
					break
				}
			}

			if found {
				li = i
			} else {
				break
			}
		}
		ws = ws[:li]

		return &wrapper{
			w.err,
			w.msg,
			w.cause,
			ws,
		}
	}

	return cause
}

func Return(err error, cause error, msg string
	type string) error {
	return internalreturn(err, 1, cause, msg)
}

func Returnf(err error, cause error, format string, args ...any) error {
	return internalreturn(err, 1, cause, fmt.Sprintf(format, args...))
}

func (e *wrapper) Error() string {
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

func (e *wrapper) Is(other error) bool {
	if e.err != nil {
		return e.err == other
	} else {
		if x, ok := other.(*wrapper); ok && e == x {
			return true
		}
	}
	return false
}

func (e *wrapper) Unwrap() error {
	return e.err
}

func (e *wrapper) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Error())
			e.stack.Format(s, verb)

			if e.cause != nil {
				io.WriteString(s, "\nCaused by: ")
				if w, ok := e.cause.(*wrapper); ok {
					w.Format(s, verb)
				} else {

					io.WriteString(s, e.cause.Error())
				}
			}

			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
