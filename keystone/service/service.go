package service

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	ServiceUrl = "services"
)

type Service struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Service {
	return &Service{
		Client: client,
	}
}

func (s *Service) All() (services []*models.ServiceModel, err error) {
	return s.AllByParams(nil)
}

func (s *Service) AllByParams(opts *options.ListServiceOpts) (ips []*models.ServiceModel, err error) {
	client, err := s.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServiceUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return models.ExtractServices(result)
}

func (s *Service) Create(opts options.CreateServiceOpts) (service *models.ServiceModel, err error) {
	client, err := s.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(ServiceUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractService(result)
}

func (s *Service) Show(id string) (service *models.ServiceModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := s.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServiceUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractService(result)
}

func (s *Service) Update(id string, opts options.UpdateServiceOpts) (service *models.ServiceModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := s.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Patch(client.ServiceURL(ServiceUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractService(result)
}

func (s *Service) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := s.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(ServiceUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
