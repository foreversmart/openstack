package role

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"

type DomainUserRole struct {
	*AbstractRole

	_ bool
}

func NewDomainUserRole(client ifaces.Openstacker) *DomainUserRole {
	return &DomainUserRole{
		AbstractRole: NewAbstractRole(client, DomainUrl, UserUrl),
	}
}
