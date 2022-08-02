package errors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("new error")

	if err.Error() != "new error" {
		t.Errorf("got %s wanted %s", err.Error(), "new error")
	}

}

func TestIs(t *testing.T) {

	err := New("base")
	err1 := Extend(err, "extend base 1")
	err2 := Extend(err1, "extend base 2")
	another := New("base")

	// can extend errors created using std library errors
	berr := errors.New("lib base error")
	berr1 := Extend(berr, "extend base error")
	berr2 := Extend(berr1, "second derieved")

	// Is should work on returned errors
	rerr := Return(err, nil, "return with context")
	rrerr := Return(rerr, nil, "return again")

	testcases := []struct {
		err    error
		target error
		match  bool
	}{
		{nil, nil, true},
		{err, nil, false},
		{err, err, true},
		{err1, err1, true},

		{err1, err, true},
		{err, err1, false},
		{err2, err, true},
		{err2, err1, true},
		{err1, err1, true},

		{another, err, false},
		{err, another, false},

		{berr, berr, true},
		{berr, nil, false},
		{berr1, berr, true},
		{berr, berr1, false},
		{berr2, berr, true},
		{berr, berr2, false},
		{berr2, berr1, true},
		{berr1, berr2, false},

		{rerr, err, true},
		{rrerr, err, true},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			if res := errors.Is(tc.err, tc.target); res != tc.match {
				t.Errorf("Is(%v, %v) = %v, should be %v", tc.err, tc.target, res, tc.match)
			}
		})
	}

}

func TestExtend(t *testing.T) {
	err := New("base error")

	errW := Extend(err, "derieved error")

	if errW.Error() != "derieved error" {
		t.Errorf("derieved error should only contain its own messges")
	}

	if errors.Is(errW.(*Error).Unwrap(), err) {
		t.Errorf("derieved \"%s\" error is not base error \"%s\"", errW.(*Error).Unwrap().Error(), err.Error())
	}
}

func TestReturn(t *testing.T) {

	ErrNew := New("abc")

	errR := Return(ErrNew, nil, "some error")
	errRR := Return(errR, nil, "another error")

	testcases := []struct {
		err    error
		target string
	}{
		{errR, "abc: some error"},
		{errRR, "abc: some error: another error"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			if ret := tc.err.Error(); ret != tc.target {
				t.Errorf("return: got %v; want %v", tc.err.Error(), tc.target)
			}
		})
	}
}

func TestEmbedding(t *testing.T) {

	type Embed struct {
		Error
		a int
		b string
	}

	Err := &Embed{Error{nil, nil, "Embed Err"}, 10, "something"}

	if Err.Error.Error() != "Embed Err" {
		t.Errorf("Error: got %v; want %v", Err.Error.Error(), "Embed Err")
	}

}
