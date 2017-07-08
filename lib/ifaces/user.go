package ifaces

import "github.com/kirk-enterprise/openstack/lib/models"
import "github.com/kirk-enterprise/openstack/lib/options"

type User interface {
	UserGrouper
	UserProjecter

	// old apis
	Register(username, password, email string, enabled bool) (user *models.UserResult, err error)
	ChangePassword(userId, newPassword string) error
	AddAccessKey(user, accessKey, secretKey string) error
	All() (userInfos []*models.UserResult, err error)
	Query(id string) (userInfo *models.UserResult, err error)

	// new apis
	AllV3(opts options.ListUserOpts) (user []*models.UserModel, err error)
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
