package group

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

type Group struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	GroupUrl = "groups"
)

func New(client ifaces.Openstacker) *Group {
	return &Group{
		Client: client,
	}
}

func (group *Group) All() (groups []*models.GroupModel, err error) {
	return group.AllByParams(nil)
}

func (group *Group) AllByParams(opts *options.ListGroupOpts) (groups []*models.GroupModel, err error) {
	client, err := group.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(GroupUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractGroups(result)
}

func (group *Group) Create(opts options.CreateGroupOpts) (groupInfo *models.GroupModel, err error) {
	client, err := group.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Post(client.ServiceURL(GroupUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractGroup(result)
}

func (group *Group) Show(id string) (groupInfo *models.GroupModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := group.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(GroupUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractGroup(result)
}

func (group *Group) Update(id string, opts options.UpdateGroupOpts) (groupInfo *models.GroupModel, err error) {
	client, err := group.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, err = client.Patch(client.ServiceURL(GroupUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractGroup(result)
}

func (group *Group) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := group.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(GroupUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
