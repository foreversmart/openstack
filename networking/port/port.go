package port

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/ports"
	"qbox.us/gogo/lib/evm/platform"
)

type Port struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	PortsUrl = "ports"
)

func New(client ifaces.Openstacker) *Port {
	return &Port{
		Client: client,
	}
}

func (p *Port) Create(opts *options.CreatePortOpts) (port *models.PortModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := p.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(PortsUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractPort(res)
}

func (p *Port) All() (infos []*models.PortModel, err error) {
	return p.AllByParams(nil)
}

func (p *Port) AllByParams(opts *options.ListPortOpts) (infos []*models.PortModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}
	client, err := p.Client.NetworkClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(PortsUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return models.ExtractPorts(result)
}

func (p *Port) Show(id string) (port *models.PortModel, err error) {
	if id == "" {
		return port, platform.ErrInvalidParams
	}

	client, err := p.Client.NetworkClient()
	if err != nil {
		return port, err
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(PortsUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractPort(result)
}

func (p *Port) Delete(id string) error {
	if id == "" {
		return platform.ErrInvalidParams
	}

	client, err := p.Client.NetworkClient()
	if err != nil {
		return err
	}

	return ports.Delete(client, id).ExtractErr()
}
