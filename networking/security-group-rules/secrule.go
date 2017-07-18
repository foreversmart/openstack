package secrule

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/extensions/security/rules"
)

type Secrule struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	SecruleUrl = "security-group-rules"
)

func New(client ifaces.Openstacker) *Secrule {
	return &Secrule{
		Client: client,
	}
}

func (s *Secrule) Create(opts *options.CreateSecruleOpts) (secrule *models.SecGroupRuleModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := s.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(SecruleUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractSecRule(res)
}

func (s *Secrule) All() (secruleInfos []*models.SecGroupRuleModel, err error) {
	return s.AllByParams(nil)
}

func (s *Secrule) AllByParams(opts *options.ListSecRuleOpts) (secruleInfos []*models.SecGroupRuleModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := s.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(SecruleUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractSecRules(result)
}

func (s *Secrule) Show(id string) (rule *models.SecGroupRuleModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := s.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(SecRuleUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return models.ExtractSecRule(result)
}

func (s *Secrule) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := s.Client.NetworkClient()
	if err != nil {
		return err
	}

	return rules.Delete(client, id).ExtractErr()
}
