package errors

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type stack []uintptr

func (st stack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			frames := runtime.CallersFrames(st)
			for {
				frame, more := frames.Next()
				if frame.Function == "runtime.main" {
					break
				}

				io.WriteString(s, "\n\t")
				io.WriteString(s, frame.Function)
				io.WriteString(s, fmt.Sprintf("(%s:%d) (0x%x)", frame.File, frame.Line, frame.PC))

				if !more {
					break
				}
			}
		}

	}
}

func StackTrace(err error) string {
	if e, ok := err.(*wrapper); ok {
		stack := e.stack
		frames := runtime.CallersFrames(stack)

		var sb strings.Builder

		for {
			frame, more := frames.Next()
			sb.WriteString(frame.Function)
			sb.WriteString("\n\t")
			sb.WriteString(frame.File)
			sb.WriteString(fmt.Sprintf(":%d (0x%x)", frame.Line, frame.PC))
			sb.WriteString("\n")
			if !more {
				break
			}

		}
		return sb.String()
	}

	return err.Error()
}
