package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Grouper interface {
	All(opts options.ListGroupOpts) (groups []*models.GroupModel, err error)
	Create(opts options.CreateGroupOpts) (group *models.GroupModel, err error)
	Show(id string) (group *models.GroupModel, err error)
	Update(id string, opts options.UpdateGroupOpts) (groupInfo *models.GroupModel, err error)
	Delete(id string) (err error)
}

type GroupUser interface {
	All(opts options.ListGroupUserOpts) (users []*models.UserModel, err error)
	Create(opts options.BaseGroupUserOpts) (err error) // add user to a group
	HasUser(opts options.BaseGroupUserOpts) (exist bool, err error)
	Delete(opts options.BaseGroupUserOpts) (err error)
}
