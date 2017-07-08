package domain

import (
	"github.com/kirk-enterprise/openstack/lib/errors"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
)

type Domain struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	DomainUrl = "domains"
)

func New(client ifaces.Openstacker) *Domain {
	return &Domain{
		Client: client,
	}
}

func (domain *Domain) All(opts options.ListDomainOpts) (domains []*models.DomainModel, err error) {
	client, err := domain.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(DomainUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractDomains(result)
}

func (domain *Domain) Create(opts options.CreateDomainOpts) (info *models.DomainModel, err error) {
	client, err := domain.Client.IdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(DomainUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractDomain(result)
}

func (domain *Domain) Show(id string) (info *models.DomainModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := domain.Client.IdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(DomainUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractDomain(result)
}

func (domain *Domain) Update(id string, opts options.UpdateDomainOpts) (info *models.DomainModel, err error) {
	client, err := domain.Client.IdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Patch(client.ServiceURL(DomainUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractDomain(result)
}

func (domain *Domain) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := domain.Client.IdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(DomainUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
