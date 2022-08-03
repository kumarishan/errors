package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestHttpErr(t *testing.T) {
	ErrHttpNotFound := NewHttpErr(nil, "not found", http.StatusNotFound)

	if ErrHttpNotFound.Error() != "not found, http Not Found" {
		t.Errorf("Error got %v, should be %v", ErrHttpNotFound.Error(), "not found, http Not Found")
	}

}

func TestHttpErrEquals(t *testing.T) {
	ErrHttpNotFound := NewHttpErr(nil, "not found", http.StatusNotFound)

	if ErrHttpNotFound != ErrHttpNotFound {
		t.Errorf("ErrHttpNotFound != ErrHttpNotFound, should be equal")
	}
}

func TestHttpErrIs(t *testing.T) {
	ErrHttpNotFound := NewHttpErr(nil, "not found", http.StatusNotFound)

	ErrRet := Return(ErrHttpNotFound, nil, "entity not found")

	testcases := []struct {
		err    error
		target error
		match  bool
	}{
		{ErrHttpNotFound, ErrHttpNotFound, true},
		{ErrRet, ErrHttpNotFound, true},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			if got := errors.Is(tc.err, tc.target); got != tc.match {
				t.Errorf("got errors.Is(%v, %v) = %v, should be %v", tc.err, tc.target, got, tc.match)
			}
		})
	}

}
