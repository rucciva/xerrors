package xerrors_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/rucciva/xerrors"
)

type tError struct{}

func (t *tError) Error() string { return "" }

func TestErrorw(t *testing.T) {
	err := xerrors.Errorw(nil)
	if err != nil {
		t.Error("should return nil if no error wrapped")
	}

	terr := &tError{}
	var terr1 *tError
	err = xerrors.Errorw(xerrors.Errorw(terr))
	if !errors.Is(err, terr) {
		t.Error("should unwrap correctly")
	}
	if !errors.As(err, &terr1) {
		t.Error("should unwrap correctly")
	}

	errTarget := xerrors.Errorf("the test")
	lines := strings.Split(fmt.Sprintf("%+v", errTarget), "- ")
	if len(lines) != 1 {
		t.Error("assumption of detail printing is wrong")
	}

	err = xerrors.Errorw(errTarget)
	err = xerrors.Errorw(err)
	err = xerrors.Errorw(err)
	t.Logf("%v", err)
	t.Logf("%+v", err)
	if err.Error() != errTarget.Error() {
		t.Error("should return wrapped error's Error()")
	}
	if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", errTarget) {
		t.Error("should not add string")
	}

	err = xerrors.Errorf("the error2: %w", err)
	err = xerrors.Errorf("the error1: %w", err)
	t.Logf("%v", err)
	t.Logf("%+v", err)
	linesW := strings.Split(fmt.Sprintf("%+v", err), "- ")
	if len(linesW) != 6 {
		t.Error("should add line to error detail", fmt.Sprintf("\n%+v", err))
	}
	if err.Error() != "the error1: the error2: the test" {
		t.Errorf("should not add any string in between message")
	}
	if fmt.Sprintf("%v", err) != "the error1: the error2: the test" {
		t.Error("should not add any string in between message")
	}
}
