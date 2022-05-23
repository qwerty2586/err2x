package tryx

import (
	"errors"
	"github.com/lainio/err2"
	"testing"
)

var (
	errNormal   = errors.New("normal")
	errOverRide = errors.New("override")
)

func TestOverride(t *testing.T) {
	defer err2.Catch(func(err error) {
		if !errors.Is(err, errOverRide) {
			t.Errorf("expected error to be %v, got %v", errOverRide, err)
		}
	})
	val := To1(func() (string, error) {
		return fReturningError()
	}, errOverRide)

	t.Error("should not be reached, but if it is, the value is:", val)

}

func fReturningError() (string, error) {
	return "", errNormal
}
