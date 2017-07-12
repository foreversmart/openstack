package compute

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Compute(t *testing.T) {
	assertion := assert.New(t)

	cmp := New(nil)
	assertion.NotNil(cmp)
}
