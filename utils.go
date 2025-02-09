package errors

import "errors"

func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Define a new error
func New(msg string
	type string) error {
	return &BaseError{
		msg,
	}
}

// Create a new error extended from another error
// error.Is will now work on both this and the from error
func Extend(err error, msg string
	type string) error {
	if err == nil {
		panic("errors: cannot extend from nil")
	}

	return &extendedError{
		BaseError{
			msg,
		},
		err,
	}
}
