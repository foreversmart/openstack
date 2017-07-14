package networking

import (
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
)

func Test_Networking(t *testing.T) {
	assertion := assert.New(t)

	n := New(nil)
	assertion.NotNil(n)
	assertion.Implements((*ifaces.FloatingIPer)(nil), n.NewFloatingIPer())
	assertion.Implements((*ifaces.Networker)(nil), n.NewNetworker())
	assertion.Implements((*ifaces.Subneter)(nil), n.NewSubneter())
}
