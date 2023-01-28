package utils_test

import (
	"testing"

	a "github.com/james-elicx/go-utils/assert"
	"github.com/james-elicx/go-utils/utils"
)

func TestTernary(t *testing.T) {
	t.Parallel()

	a.Equals(t, utils.Ternary(true, 1, 2), 1)
	a.Equals(t, utils.Ternary(false, 1, 2), 2)
}
