package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestHttpErrError(t *testing.T) {
	ErrHttpNotFound := NewHttpErr("not found", http.StatusNotFound)

	testcases := []struct {
		err    error
		target string
	}{
		{ErrHttpNotFound, "not found, http Not Found"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			if got := tc.err.Error(); got != tc.target {
				t.Errorf("got %v.Error() = %v; should be %v", tc.err, got, tc.target)
			}
		})
	}

}

func TestHttpErrEquals(t *testing.T) {
	ErrHttpNotFound := NewHttpErr("not found", http.StatusNotFound)

	if ErrHttpNotFound != ErrHttpNotFound {
		t.Errorf("ErrHttpNotFound != ErrHttpNotFound, should be equal")
	}
}

func TestHttpErrIs(t *testing.T) {
	ErrHttpNotFound := NewHttpErr("not found", http.StatusNotFound)
	ErrHttpNotFound2 := NewHttpErr("not found", http.StatusNotFound)

	ErrRet1 := Return(ErrHttpNotFound, nil, "entity not found")
	ErrRet2 := Return(ErrRet1, nil, "entity not found again")
	ErrRetCause := Return(ErrHttpNotFound, errors.New("some cause"), "entity not found")

	testcases := []struct {
		err    error
		target error
		match  bool
	}{
		{ErrHttpNotFound, ErrHttpNotFound, true},
		{ErrHttpNotFound, ErrHttpNotFound2, false},
		{ErrRet1, ErrHttpNotFound, true},
		{ErrRet2, ErrHttpNotFound, true},
		{ErrHttpNotFound, ErrRet1, false},
		{ErrRetCause, ErrHttpNotFound, true},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			if got := errors.Is(tc.err, tc.target); got != tc.match {
				t.Errorf("got errors.Is(%v, %v) = %v, should be %v", tc.err, tc.target, got, tc.match)
			}
		})
	}

}
