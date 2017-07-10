package role

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"

type DomainGroupRole struct {
	*AbstractRole

	_ bool
}

func NewDomainGroupRole(client ifaces.Openstacker) *DomainGroupRole {
	return &DomainGroupRole{
		AbstractRole: NewAbstractRole(client, DomainUrl, GroupUrl),
	}
}
