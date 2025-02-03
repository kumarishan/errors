package errors

const MaxStackDepth = 50

type BaseError struct {
	msg string
    type string
}

func (b *BaseError) Error() string {
	return b.msg
}

type extendedError struct {
	BaseError

	err error
}

func (e *extendedError) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return e.err.Error()
}

func (e *extendedError) Is(other error) bool {
	if e.err == other {
		return true

	} else if x, ok := other.(*extendedError); ok && e == x {
		return true
	}

	return false
}

func (e *extendedError) Unwrap() error {
	return e.err
}
