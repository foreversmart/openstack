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

	assertion.Implements((*ifaces.Server)(nil), cmp.NewServer())

	assertion.Implements((*ifaces.ServerManager)(nil), cmp.NewServerManager())
	assertion.Implements((*ifaces.ServerImager)(nil), cmp.NewServerImager())
	assertion.Implements((*ifaces.ServerKeyer)(nil), cmp.NewServerKeyer())
	assertion.Implements((*ifaces.ServerPorter)(nil), cmp.NewServerPorter())
	assertion.Implements((*ifaces.ServerVolumer)(nil), cmp.NewServerVolumer())

	assertion.Implements((*ifaces.Flavor)(nil), cmp.NewFlavor())
	assertion.Implements((*ifaces.Keypair)(nil), cmp.NewKeypair())
}
