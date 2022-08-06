package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("new error")

	if err.Error() != "new error" {
		t.Errorf("got %s wanted %s", err.Error(), "new error")
	}

}

func TestError(t *testing.T) {
	ErrEmpty := New("")
	Err := New("base")
	Err1 := Extend(Err, "extend base 1")
	Err2 := Extend(Err1, "extend base 2")
	Err1Empty := Extend(Err, "")
	Err2Empty := Extend(Err1, "")
	Err3Empty := Extend(ErrEmpty, "")

	rerr := Return(Err, nil, "return base")
	rerr1 := Return(Err1, nil, "return base 1")
	rerr2 := Return(Err2, nil, "return base 2")

	rrerr := Return(rerr, nil, "return return base")
	rrerr1 := Return(Err, nil, "")
	rrerr2 := Return(Err2, nil, "")
	rrerr3 := Return(ErrEmpty, nil, "")
	rrerr4 := Return(nil, nil, "some error")
	rrerr5 := Return(nil, nil, "")

	Cause := New("cause")
	cerr1 := Return(Err, Cause, "")
	cerr2 := Return(Err1, Cause, "")
	cerr3 := Return(Err, Cause, "some error")
	cerr4 := Return(ErrEmpty, Cause, "")
	cerr5 := Return(ErrEmpty, Cause, "some error")

	testcases := []struct {
		err    error
		target string
	}{
		{ErrEmpty, ""},
		{Err, "base"},
		{Err1, "extend base 1"},
		{Err2, "extend base 2"},
		{Err1Empty, "base"},
		{Err2Empty, "extend base 1"},
		{Err3Empty, ""},
		{rerr, "return base"},
		{rerr1, "return base 1"},
		{rerr2, "return base 2"},
		{rrerr, "return return base"},
		{rrerr1, "base"},
		{rrerr2, "extend base 2"},
		{rrerr3, ""},
		{rrerr4, "some error"},
		{rrerr5, ""},

		{cerr1, "base"},
		{cerr2, "extend base 1"},
		{cerr3, "some error"},
		{cerr4, "cause"},
		{cerr5, "some error"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			if res := tc.err.Error(); res != tc.target {
				t.Errorf("got [%v].Error() = %v, should be %v", tc.err, res, tc.target)
			}
		})
	}
}

func TestIs(t *testing.T) {

	ErrEmpty1 := New("")
	ErrEmpty2 := New("")
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
		{ErrEmpty1, ErrEmpty1, true},
		{ErrEmpty2, ErrEmpty2, true},
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

func TestEmbedding(t *testing.T) {

	type Embed struct {
		BaseError
		a int
		b string
	}

	Err := &Embed{BaseError{"Embed Err"}, 10, "something"}

	if Err.Error() != "Embed Err" {
		t.Errorf("Error: got %v; want %v", Err.Error(), "Embed Err")
	}

}

func TestAs(t *testing.T) {

	f := func() error {
		return &HttpErr{
			BaseError{"not found"},
			http.StatusNotFound,
		}
	}

	f2 := func() error {
		return Return(f(), nil, "extend 1")
	}

	rerr := f2()

	var httpErr *HttpErr
	if !errors.As(rerr, &httpErr) {
		t.Errorf("As not working")
	}
}
