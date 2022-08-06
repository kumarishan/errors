package errors

var (
	ErrInvalidInput            = New("InvalidInputError")
	ErrInternal                = New("InternalError")
	ErrUnexpectedInternalState = New("UnexpectedInternalStateError")
	ErrNotFound                = New("NotFoundError")
)
