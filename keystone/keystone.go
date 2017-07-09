package keystone

import (
	"github.com/kirk-enterprise/openstack/internal"
	"github.com/kirk-enterprise/openstack/keystone/credential"
	"github.com/kirk-enterprise/openstack/keystone/domain"
	"github.com/kirk-enterprise/openstack/keystone/endpoint"
	"github.com/kirk-enterprise/openstack/keystone/group"
	"github.com/kirk-enterprise/openstack/keystone/policy"
	"github.com/kirk-enterprise/openstack/keystone/project"
	"github.com/kirk-enterprise/openstack/keystone/region"
	"github.com/kirk-enterprise/openstack/keystone/role"
	"github.com/kirk-enterprise/openstack/keystone/service"
	"github.com/kirk-enterprise/openstack/keystone/user"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
)

type Keystone struct {
	*internal.Openstack

	_ bool
}

func New(endpoint string) *Keystone {
	return &Keystone{
		Openstack: internal.New(endpoint),
	}
}

func (ks *Keystone) NewCredential() ifaces.Credentialer {
	return credential.New(ks)
}

func (ks *Keystone) NewDomain() ifaces.Domainer {
	return domain.New(ks)
}

func (ks *Keystone) NewEndpoint() ifaces.Endpointer {
	return endpoint.New(ks)
}

func (ks *Keystone) NewGroup() ifaces.Grouper {
	return group.New(ks)
}

func (ks *Keystone) NewGroupUser() ifaces.GroupUser {
	return group.NewGroupUser(ks)
}

func (ks *Keystone) NewPolicy() ifaces.Policier {
	return policy.New(ks)
}

func (ks *Keystone) NewProject() ifaces.Projecter {
	return project.New(ks)
}

func (ks *Keystone) NewRegion() ifaces.Regioner {
	return region.New(ks)
}

func (ks *Keystone) NewRole() ifaces.Roler {
	return role.New(ks)
}

func (ks *Keystone) NewDomainGroupRole() ifaces.AbstractRoler {
	return role.NewDomainGroupRole(ks)
}

func (ks *Keystone) NewDomainUserRole() ifaces.AbstractRoler {
	return role.NewDomainUserRole(ks)
}

func (ks *Keystone) NewProjectUserRole() ifaces.AbstractRoler {
	return role.NewProjectUserRole(ks)
}

func (ks *Keystone) NewProjectGroupRole() ifaces.AbstractRoler {
	return role.NewProjectGroupRole(ks)
}

func (ks *Keystone) NewService() ifaces.Servicer {
	return service.New(ks)
}

func (ks *Keystone) NewUser() ifaces.User {
	return user.New(ks)
}
