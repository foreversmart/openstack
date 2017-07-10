package role

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"

type ProjectUserRole struct {
	*AbstractRole

	_ bool
}

func NewProjectUserRole(client ifaces.Openstacker) *ProjectUserRole {
	return &ProjectUserRole{
		AbstractRole: NewAbstractRole(client, ProjectUrl, UserUrl),
	}
}
