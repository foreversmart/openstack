package openstack

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Openstack(t *testing.T) {
	assertion := assert.New(t)

	os := New()
	assertion.NotNil(os)
}
