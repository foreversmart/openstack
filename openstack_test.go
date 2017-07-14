package openstack

import (
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

func Test_Openstack(t *testing.T) {
	assertion := assert.New(t)

	os := New("")
	assertion.NotNil(os)
	assertion.Implements((*ifaces.Openstacker)(nil), os)
}
