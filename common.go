package errors

var (
	ErrInvalidInput            = New("invalid input error")
	ErrInternal                = New("internal error")
	ErrUnexpectedInternalState = New("unexpected internal state error")
	ErrNotFound                = New("not found erro")
)
