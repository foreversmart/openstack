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
	All(groupID string, opts options.ListGroupUserOpts) (users []*models.UserModel, err error)
	Create(groupID string, opts options.CreateGroupUserOpts) (err error) // add user to a group
	HasUser(groupID, userID string) (ok bool)
	Delete(groupID, userID string) (err error)
}
