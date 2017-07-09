package group

import (
	"io"
	"net/http"

	"github.com/kirk-enterprise/openstack/lib/errors"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
)

type GroupUser struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	GroupUserUrl = "users"
)

func NewGroupUser(client ifaces.Openstacker) *GroupUser {
	return &GroupUser{
		Client: client,
	}
}

func (groupUser *GroupUser) All(groupID string) (users []*models.UserModel, err error) {
	return groupUser.AllByParams(groupID, nil)
}

func (groupUser *GroupUser) AllByParams(groupID string, opts *options.ListGroupUserOpts) (groupUsers []*models.UserModel, err error) {
	if groupID == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := groupUser.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(GroupUrl, groupID, GroupUserUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractUsers(result)
}

// TODO: support add user with roles!
func (groupUser *GroupUser) Create(groupID string, opts options.CreateGroupUserOpts) (err error) {
	if groupID == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := groupUser.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Put(client.ServiceURL(GroupUrl, groupID, GroupUserUrl, opts.UserID), nil, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{204, 200},
	})

	// NOTE: gophercloud response issue with 204
	if result.Err == io.EOF {
		result.Err = nil
	}

	return result.Err
}

func (groupUser *GroupUser) HasUser(groupID, userID string) (ok bool) {
	if groupID == "" || userID == "" {
		return false
	}

	client, err := groupUser.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	response, err := client.Request(http.MethodHead, client.ServiceURL(GroupUrl, groupID, GroupUserUrl, userID), gophercloud.RequestOpts{
		OkCodes: []int{204, 404},
	})

	if response != nil {
		switch response.StatusCode {
		case 204:
			ok = true
		}
	}

	return
}

func (groupUser *GroupUser) Delete(groupID, userID string) (err error) {
	if groupID == "" || userID == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := groupUser.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(GroupUrl, groupID, GroupUserUrl, userID), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	// NOTE: gophercloud response issue with 204
	if err == io.EOF {
		err = nil
	}

	return err
}
