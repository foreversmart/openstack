package keystone

import (
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

func Test_Keystone(t *testing.T) {
	assertion := assert.New(t)

	ks := New(nil)
	assertion.NotNil(ks)
	// assertion.Implements((*ifaces.Openstacker)(nil), ks)
	assertion.Implements((*ifaces.Credentialer)(nil), ks.NewCredential())
	assertion.Implements((*ifaces.Domainer)(nil), ks.NewDomain())
	assertion.Implements((*ifaces.Endpointer)(nil), ks.NewEndpoint())
	assertion.Implements((*ifaces.Grouper)(nil), ks.NewGroup())
	assertion.Implements((*ifaces.GroupUser)(nil), ks.NewGroupUser())
	assertion.Implements((*ifaces.Policier)(nil), ks.NewPolicy())
	assertion.Implements((*ifaces.Projecter)(nil), ks.NewProject())
	assertion.Implements((*ifaces.Regioner)(nil), ks.NewRegion())
	assertion.Implements((*ifaces.Roler)(nil), ks.NewRole())
	assertion.Implements((*ifaces.AbstractRoler)(nil), ks.NewDomainGroupRole())
	assertion.Implements((*ifaces.AbstractRoler)(nil), ks.NewDomainUserRole())
	assertion.Implements((*ifaces.AbstractRoler)(nil), ks.NewProjectGroupRole())
	assertion.Implements((*ifaces.AbstractRoler)(nil), ks.NewProjectUserRole())
	assertion.Implements((*ifaces.Servicer)(nil), ks.NewService())
	assertion.Implements((*ifaces.User)(nil), ks.NewUser())
}
