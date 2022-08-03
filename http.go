package errors

import (
	"fmt"
	"net/http"
)

type HttpErr struct {
	BaseError
	code int
}

func NewHttpErr(cause error, msg string, code int) error {
	return &HttpErr{
		BaseError{
			cause,
			msg,
		},
		code,
	}
}

func (h *HttpErr) Error() string {
	return fmt.Sprintf("%s, http %s", h.BaseError.Error(), http.StatusText(h.code))
}