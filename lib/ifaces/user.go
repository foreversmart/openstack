package ifaces

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
import "github.com/kirk-enterprise/openstack-golang-sdk/lib/options"

type User interface {
	UserGrouper
	UserProjecter

	All() (user []*models.UserModel, err error)
	AllByParams(opts *options.ListUserOpts) (user []*models.UserModel, err error)
	Create(opts options.CreateUserOpts) (user *models.UserModel, err error)
	Show(UserID string) (user *models.UserModel, err error)
	Update(UserID string, opts options.UpdateUserOpts) (userInfo *models.UserModel, err error)
	ChangePasswd(UserID string, opts options.ChangeUserPasswordOpts) (err error)
	Delete(UserID string) (err error)
}

type UserGrouper interface {
	AllGroups(userID string) (groups []*models.GroupModel, err error)
}

type UserProjecter interface {
	AllProjects(userID string) (projects []*models.ProjectModel, err error)
}
