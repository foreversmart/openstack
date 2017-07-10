package endpoint

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	EndpointUrl = "endpoints"
)

type Endpoint struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Endpoint {
	return &Endpoint{
		Client: client,
	}
}

func (ep *Endpoint) All() (endpoints []*models.EndpointModel, err error) {
	return ep.AllByParams(nil)
}

func (ep *Endpoint) AllByParams(opts *options.ListEndpointOpts) (services []*models.EndpointModel, err error) {
	client, err := ep.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(EndpointUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return models.ExtractEndpoints(result)
}

func (ep *Endpoint) Show(id string) (endpoint *models.EndpointModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := ep.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Get(client.ServiceURL(EndpointUrl, id), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractEndpoint(res)
}

func (ep *Endpoint) Create(opts options.CreateEndpointOpts) (endpoint *models.EndpointModel, err error) {
	client, err := ep.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(EndpointUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractEndpoint(res)
}

func (ep *Endpoint) Update(id string, opts options.UpdateEndpointOpts) (endpoint *models.EndpointModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := ep.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, err = client.Patch(client.ServiceURL(EndpointUrl, id), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractEndpoint(res)
}

func (ep *Endpoint) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := ep.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(EndpointUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
