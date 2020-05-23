package xerrors

import (
	"fmt"
)

type NoString interface {
	Formatter

	noString()
}

type errorNoString struct {
	err   error
	frame Frame
}

func (e *errorNoString) Error() string {
	return e.err.Error()
}

// Unwrap implements xerrors.Wrapper.
func (e *errorNoString) Unwrap() (err error) {
	return e.err
}

// Format implements fmt.Formatter.
func (e *errorNoString) Format(f fmt.State, c rune) {
	FormatError(e, f, c)
}

// FormatError implements xerrors.Formatter.
func (e *errorNoString) FormatError(p Printer) (next error) {
	if p.Detail() {
		e.frame.Format(p)
	}
	return e.err
}

func (e *errorNoString) noString() {}

// Errorw wrap an error without adding additional string.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
// The returned error also implement Unwrap()
//
// Will return nil if no error to be wrapped
func Errorw(err error) error {
	if err == nil {
		return nil
	}
	return &errorNoString{
		err:   err,
		frame: Caller(1),
	}
}
