package a_test

import (
	"testing"

	a "github.com/james-elicx/go-utils/assert"
)

type TestCase struct {
	a, b interface{}
}

var equalsCases = []TestCase{
	{1, 1},
	{"a", "a"},
	{true, true},
	{false, false},
	{nil, nil},
}

var notEqualsCases = []TestCase{
	{1, 2},
	{"a", "b"},
	{true, false},
	{false, true},
	{nil, 1},
}

func TestEquals(t *testing.T) {
	t.Parallel()

	for _, c := range equalsCases {
		a.Equals(t, c.a, c.b)
	}
}

func TestEqualsFail(t *testing.T) {
	t.Parallel()

	for _, c := range notEqualsCases {
		a.ShouldFailTest(t, func(mockTest *testing.T) {
			a.Equals(mockTest, c.a, c.b)
		})
	}
}

func TestNotEquals(t *testing.T) {
	t.Parallel()

	for _, c := range notEqualsCases {
		a.NotEquals(t, c.a, c.b)
	}
}

func TestNotEqualsFail(t *testing.T) {
	t.Parallel()

	for _, c := range equalsCases {
		a.ShouldFailTest(t, func(mockTest *testing.T) {
			a.NotEquals(mockTest, c.a, c.b)
		})
	}
}
