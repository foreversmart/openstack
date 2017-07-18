package compute

import (
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

func Test_Compute(t *testing.T) {
	assertion := assert.New(t)

	cmp := New(nil)
	assertion.NotNil(cmp)
	assertion.Implements((*ifaces.Serverser)(nil), cmp.NewServerser())
	assertion.Implements((*ifaces.Flavorer)(nil), cmp.NewFlavors())
}
