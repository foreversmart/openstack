package user

import (
	"io"

	"github.com/kirk-enterprise/openstack/lib/errors"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	UsersUrl    = "users"
	GroupUrl    = "groups"
	ProjectUrl  = "projects"
	PasswordUrl = "password"
)

type User struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *User {
	return &User{
		Client: client,
	}
}

func (user *User) All() (users []*models.UserModel, err error) {
	return user.AllByParams(nil)
}

// Identity API V3
func (user *User) AllByParams(opts *options.ListUserOpts) (users []*models.UserModel, err error) {
	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(UsersUrl)+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractUsers(result)
}

func (user *User) Create(opts options.CreateUserOpts) (info *models.UserModel, err error) {
	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(UsersUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractUser(result)
}

func (user *User) Show(userID string) (info *models.UserModel, err error) {
	if userID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(UsersUrl, userID), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractUser(result)
}

func (user *User) Update(userID string, opts options.UpdateUserOpts) (userInfo *models.UserModel, err error) {
	if userID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Patch(client.ServiceURL(UsersUrl, userID), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractUser(result)
}

func (user *User) ChangePasswd(userID string, opts options.ChangeUserPasswordOpts) (err error) {
	if userID == "" {
		return errors.ErrInvalidParams
	}

	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Post(client.ServiceURL(UsersUrl, userID, PasswordUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	// NOTE: gophercloud request issue for 204 response.
	if err == io.EOF {
		err = nil
	}

	return
}

func (user *User) Delete(userID string) (err error) {
	if userID == "" {
		return errors.ErrInvalidParams
	}

	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(UsersUrl, userID), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

//List groups to which a user belongs
func (user *User) AllGroups(userID string) (groups []*models.GroupModel, err error) {
	if userID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(UsersUrl, userID, GroupUrl), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractUserGroups(result)
}

func (user *User) AllProjects(userID string) (projects []*models.ProjectModel, err error) {
	if userID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := user.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(UsersUrl, userID, ProjectUrl), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractUserProjects(result)
}
