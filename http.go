package errors

import (
	"fmt"
	"net/http"
)

type HttpErr struct {
	BaseError
	code int
}

func NewHttpErr(msg, errType string, code int) error {
	return &HttpErr{
		BaseError{
			msg,
		},
		code,
	}
}

func (h *HttpErr) Error() string {
	return fmt.Sprintf("%s, http %s", h.BaseError.Error(), http.StatusText(h.code))
}
