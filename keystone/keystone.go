package keystone

import (
	"github.com/qbox/openstack-golang-sdk/keystone/auth"
	"github.com/qbox/openstack-golang-sdk/keystone/credential"
	"github.com/qbox/openstack-golang-sdk/keystone/domain"
	"github.com/qbox/openstack-golang-sdk/keystone/endpoint"
	"github.com/qbox/openstack-golang-sdk/keystone/group"
	"github.com/qbox/openstack-golang-sdk/keystone/policy"
	"github.com/qbox/openstack-golang-sdk/keystone/project"
	"github.com/qbox/openstack-golang-sdk/keystone/region"
	"github.com/qbox/openstack-golang-sdk/keystone/role"
	"github.com/qbox/openstack-golang-sdk/keystone/service"
	"github.com/qbox/openstack-golang-sdk/keystone/user"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

type Keystone struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Keystone {
	return &Keystone{
		client: client,
	}
}

func (ks *Keystone) NewCredential() ifaces.Credentialer {
	return credential.New(ks.client)
}

func (ks *Keystone) NewDomain() ifaces.Domainer {
	return domain.New(ks.client)
}

func (ks *Keystone) NewEndpoint() ifaces.Endpointer {
	return endpoint.New(ks.client)
}

func (ks *Keystone) NewGroup() ifaces.Grouper {
	return group.New(ks.client)
}

func (ks *Keystone) NewGroupUser() ifaces.GroupUser {
	return group.NewGroupUser(ks.client)
}

func (ks *Keystone) NewPolicy() ifaces.Policier {
	return policy.New(ks.client)
}

func (ks *Keystone) NewProject() ifaces.Projecter {
	return project.New(ks.client)
}

func (ks *Keystone) NewRegion() ifaces.Regioner {
	return region.New(ks.client)
}

func (ks *Keystone) NewRole() ifaces.Roler {
	return role.New(ks.client)
}

func (ks *Keystone) NewDomainGroupRole() ifaces.AbstractRoler {
	return role.NewDomainGroupRole(ks.client)
}

func (ks *Keystone) NewDomainUserRole() ifaces.AbstractRoler {
	return role.NewDomainUserRole(ks.client)
}

func (ks *Keystone) NewProjectUserRole() ifaces.AbstractRoler {
	return role.NewProjectUserRole(ks.client)
}

func (ks *Keystone) NewProjectGroupRole() ifaces.AbstractRoler {
	return role.NewProjectGroupRole(ks.client)
}

func (ks *Keystone) NewService() ifaces.Servicer {
	return service.New(ks.client)
}

func (ks *Keystone) NewUser() ifaces.User {
	return user.New(ks.client)
}

func (ks *Keystone) NewAuther() ifaces.Auther {
	return auth.New(ks.client)
}
