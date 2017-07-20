package server

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/rackspace/gophercloud"
)

type ServerPorter struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerPorter(client ifaces.Openstacker) *ServerPorter {
	return &ServerPorter{
		Client: client,
	}
}

func (sp *ServerPorter) All(id string) (portModels []*models.AttachPortModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := sp.Client.ComputeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result

	_, res.Err = client.Get(client.ServiceURL(ServersUrl, id, InterfaceUrl), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	portModels, err = models.ExtractAttachPorts(res)

	return
}

func (sp *ServerPorter) Bind(id, portID string) error {
	if id == "" || portID == "" {
		return errors.ErrInvalidParams
	}

	client, err := sp.Client.ComputeClient()
	if err != nil {
		return err
	}

	opts := map[string]interface{}{
		"port_id": portID,
	}

	reqBody := map[string]interface{}{
		"interfaceAttachment": opts,
	}

	var res gophercloud.Result

	_, res.Err = client.Post(client.ServiceURL(ServersUrl, id, InterfaceUrl), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return res.Err
}

func (sp *ServerPorter) Unbind(id, portID string) error {
	if id == "" || portID == "" {
		return errors.ErrInvalidParams
	}

	client, err := sp.Client.ComputeClient()
	if err != nil {
		return err
	}

	var res gophercloud.Result

	_, res.Err = client.Delete(client.ServiceURL(ServersUrl, id, InterfaceUrl, portID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return res.Err
}
