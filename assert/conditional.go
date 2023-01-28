package a

import (
	"testing"
)

func Equals(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("expected %v to equal %v", a, b)
	}
}

func NotEquals(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("expected %v to not equal %v", a, b)
	}
}
