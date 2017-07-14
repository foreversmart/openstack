package servers

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	ServersUrl = "servers"
)

type Servers struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Servers {
	return &Servers{
		Client: client,
	}
}

func (ser *Servers) Create(opts options.CreateServersOpts) (server *models.ServersModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Post(client.ServiceURL(ServersUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractServer(result)
}

func (ser *Servers) All() (servers []*models.ServersModel, err error) {
	return ser.AllByParams(nil)
}

func (ser *Servers) AllByParams(opts *options.ListServersOpts) (Serverss []*models.ServersModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}
	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServersUrl, "detail")+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServers(result)
}

func (ser *Servers) Show(id string) (server *models.ServersModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServersUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServer(result)
}

func (ser *Servers) Update(id string, opts options.UpdateServersOpts) (server *models.ServersModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Put(client.ServiceURL(ServersUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServer(result)
}

func (ser *Servers) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(ServersUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
