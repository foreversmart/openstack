package keystone

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Keystone(t *testing.T) {
	assertion := assert.New(t)

	ks := New()
	assertion.NotNil(ks)
}
