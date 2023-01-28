package a_test

import (
	"testing"

	a "github.com/james-elicx/go-utils/assert"
)

func TestShouldFailTest(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.Equals(mockTest, 1, 2)
	})
}

func TestShouldFailTestFail(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.ShouldFailTest(mockTest, func(subMockTest *testing.T) {
			a.Equals(subMockTest, 1, 1)
		})
	})
}
