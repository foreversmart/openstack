package secgroup

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

type SecurityGroups struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	SecGroupUrl = "security-groups"
)

func New(client ifaces.Openstacker) *SecurityGroups {
	return &SecurityGroups{
		Client: client,
	}
}

func (s *SecurityGroups) All() (infos []*models.SecurityGroupModel, err error) {
	return s.AllByParams(nil)
}

func (s *SecurityGroups) AllByParams(opts *options.ListSecurityGroupsOpts) (securitygroups []*models.SecurityGroupModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := s.Client.NetworkClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(SecGroupUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return models.ExtractSecurityGroups(result)
}

func (s *SecurityGroups) Create(opts *options.CreateSecurityGroupOpts) (securitygroup *models.SecurityGroupModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := s.Client.NetworkClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Post(client.ServiceURL(SecGroupUrl), opts.ToPayLoad(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractSecurityGroup(result)
}
