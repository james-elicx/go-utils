package a

import (
	"testing"
)

func Throws(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected function to throw")
		}
	}()

	f()
}
