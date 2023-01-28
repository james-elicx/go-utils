package a_test

import (
	"testing"

	a "github.com/james-elicx/go-utils/assert"
)

func TestThrows(t *testing.T) {
	t.Parallel()

	a.Throws(t, func() {
		panic("test")
	})
}

func TestThrowsFail(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.Throws(mockTest, func() {})
	})
}
