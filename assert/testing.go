package a

import (
	"testing"
)

func ShouldFailTest(t *testing.T, f func(mockTest *testing.T)) {
	mockTest := new(testing.T)
	f(mockTest)

	if !mockTest.Failed() {
		t.Error("expected a failure, but test exited successfully")
	}
}
