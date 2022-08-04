package errors

import (
	"fmt"
	"runtime"
	"strings"
)

func StackTrace(err error) string {
	if e, ok := err.(*errorWithStackTrace); ok {
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
