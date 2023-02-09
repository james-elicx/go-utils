package a_test

import (
	"errors"
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

func TestEqualsError(t *testing.T) {
	t.Parallel()

	a.EqualsError(t, errors.New("test"), errors.New("test"))
}

func TestNotEqualsError(t *testing.T) {
	t.Parallel()

	a.NotEqualsError(t, errors.New("test"), errors.New("not test"))
}

func TestEqualsErrorMessage(t *testing.T) {
	t.Parallel()

	a.EqualsErrorMessage(t, errors.New("test"), "test")
}

func TestNotEqualsErrorMessage(t *testing.T) {
	t.Parallel()

	a.NotEqualsErrorMessage(t, errors.New("test"), "not test")
}

func TestEqualsArray(t *testing.T) {
	t.Parallel()

	a.EqualsArray(t, []int{1, 2, 3}, []int{1, 2, 3})
	a.EqualsArray(t, []string{"a", "b", "c"}, []string{"a", "b", "c"})
	a.EqualsArray(t, []byte{1, 2, 3}, []byte{1, 2, 3})
	a.EqualsArray(t, []bool{true, false, true}, []bool{true, false, true})
	a.EqualsArray(t, []interface{}{1, "a", true}, []interface{}{1, "a", true})
	a.EqualsArray(t, []byte{}, []byte{})
}

func TestEqualsArrayFails(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.EqualsArray(mockTest, []int{1, 2, 3}, []int{1, 2, 3, 4})
	})
}

func TestNotEqualsArray(t *testing.T) {
	t.Parallel()

	a.NotEqualsArray(t, []int{1, 2, 3}, []int{1, 2, 3, 4})
	a.NotEqualsArray(t, []string{"a", "b", "c"}, []string{"a", "b", "c", "d"})
	a.NotEqualsArray(t, []byte{1, 2, 3}, []byte{1, 2, 3, 4})
	a.NotEqualsArray(t, []bool{true, false, true}, []bool{true, false, true, false})
	a.NotEqualsArray(t, []interface{}{1, "a", true}, []interface{}{1, "a", true, false})
	a.NotEqualsArray(t, []byte{}, []byte{1})
}

func TestNotEqualsArrayFails(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.NotEqualsArray(mockTest, []int{1, 2, 3}, []int{1, 2, 3})
	})
}

func TestGreaterThan(t *testing.T) {
	t.Parallel()

	a.GreaterThan(t, 2, 1)
}

func TestGreaterThanFail(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.GreaterThan(mockTest, 1, 2)
	})
	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.GreaterThan(mockTest, 1, 1)
	})
}

func TestLessThan(t *testing.T) {
	t.Parallel()

	a.LessThan(t, 1, 2)
}

func TestLessThanFail(t *testing.T) {
	t.Parallel()

	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.LessThan(mockTest, 2, 1)
	})
	a.ShouldFailTest(t, func(mockTest *testing.T) {
		a.LessThan(mockTest, 1, 1)
	})
}
