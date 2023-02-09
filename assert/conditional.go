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

func EqualsArray[T any](t *testing.T, a, b []T) {
	if len(a) != len(b) {
		t.Errorf("expected array length %v to equal %v", a, b)
		return
	}

	for i := 0; i < len(a); i++ {
		Equals(t, a[i], b[i])
	}
}

func NotEquals(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("expected %v to not equal %v", a, b)
	}
}

func NotEqualsError(t *testing.T, a error, b interface{}) {
	NotEquals(t, a.Error(), b)
}

func anyEqualsHelper(a, b interface{}) bool {
	return a == b
}

func NotEqualsArray[T any](t *testing.T, a, b []T) {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if anyEqualsHelper(a[i], b[i]) {
				t.Errorf("expected %v to not equal %v", a[i], b[i])
				break
			}
		}
	}
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
