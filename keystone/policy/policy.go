package policy

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	PolicyUrl = "policies"
)

type Policy struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Policy {
	return &Policy{
		Client: client,
	}
}

func (policy *Policy) All() (policies []*models.PolicyModel, err error) {
	return policy.AllByParams(nil)
}

func (policy *Policy) AllByParams(opts *options.ListPolicyOpts) (policies []*models.PolicyModel, err error) {
	client, err := policy.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(PolicyUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractPolicies(result)
}

func (policy *Policy) Create(opts options.CreatePolicyOpts) (policyInfo *models.PolicyModel, err error) {
	client, err := policy.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Post(client.ServiceURL(PolicyUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractPolicy(result)
}

func (policy *Policy) Show(id string) (policyInfo *models.PolicyModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := policy.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(PolicyUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractPolicy(result)
}

func (policy *Policy) Update(id string, opts options.UpdatePolicyOpts) (policyInfo *models.PolicyModel, err error) {
	client, err := policy.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Patch(client.ServiceURL(PolicyUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractPolicy(result)
}

func (policy *Policy) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := policy.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(PolicyUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
