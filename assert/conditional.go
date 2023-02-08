package a

import (
	"testing"
)

func Equals(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("expected %v to equal %v", a, b)
	}
}

func EqualsError(t *testing.T, a error, b interface{}) {
	Equals(t, a.Error(), b)
}

func NotEquals(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("expected %v to not equal %v", a, b)
	}
}

func NotEqualsError(t *testing.T, a error, b interface{}) {
	NotEquals(t, a.Error(), b)
}

func GreaterThan(t *testing.T, a, b int) {
	if a <= b {
		t.Errorf("expected %v to be greater than %v", a, b)
	}
}

func LessThan(t *testing.T, a, b int) {
	if a >= b {
		t.Errorf("expected %v to be less than %v", a, b)
	}
}
