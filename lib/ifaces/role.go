package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Roler interface {
	All() (roles []*models.RoleModel, err error)
	AllByParams(opts *options.ListRoleOpts) (roles []*models.RoleModel, err error)
	Create(opts *options.CreateRoleOpts) (role *models.RoleModel, err error)
	Show(id string) (role *models.RoleModel, err error)
	Update(id string, opts *options.UpdateRoleOpts) (role *models.RoleModel, err error)
	Delete(id string) (err error)
}

type AbstractRoler interface {
	All(resourceID, scopeID string) (roles []*models.RoleModel, err error)
	Create(resourceID, scopeID, roleID string) (err error)
	HasRole(resourceID, scopeID, roleID string) (exist bool)
	Delete(resourceID, scopeID, roleID string) (err error)
}

type DomainGroupRoler interface {
	AbstractRoler
}

type DomainUserRoler interface {
	AbstractRoler
}

type ProjectGroupRoler interface {
	AbstractRoler
}

type ProjectUserRoler interface {
	AbstractRoler
}
