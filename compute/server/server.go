package server

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	ServersUrl   = "servers"
	InterfaceUrl = "os-interface"
	VolumeUrl    = "os-volume_attachments"
	ActionUrl    = "action"
)

type Server struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Server {
	return &Server{
		Client: client,
	}
}

func (ser *Server) Create(opts options.CreateServerOpts) (server *models.ServerModel, err error) {
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
		OkCodes: []int{202},
	})

	return models.ExtractServer(result)
}

func (ser *Server) All() (servers []*models.ServerModel, err error) {
	return ser.AllByParams(nil)
}

func (ser *Server) AllByParams(opts *options.ListServersOpts) (Servers []*models.ServerModel, err error) {
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

func (ser *Server) Show(id string) (server *models.ServerModel, err error) {
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

func (ser *Server) Update(id string, opts options.UpdateServersOpts) (server *models.ServerModel, err error) {
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

func (ser *Server) Delete(id string) (err error) {
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
