package network

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
)

type Network struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Network {
	return &Network{
		Client: client,
	}
}

const (
	NetworkUrl = "networks"
)

func (n *Network) Create(opts *options.CreateNetworkOpts) (network *models.NetworkModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return network, err
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(NetworkUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractNetwork(res)
}

func (n *Network) All() (infos []*models.NetworkModel, err error) {
	return n.AllByParams(nil)
}

func (n *Network) AllByParams(opts *options.ListNetworkOpt) (infos []*models.NetworkModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(NetworkUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return models.ExtractNetworks(result)
}

func (n *Network) Show(id string) (info *models.NetworkModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	result := networks.Get(client, id)
	return models.ExtractNetwork(result.Result)
}

func (n *Network) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return err
	}

	return networks.Delete(client, id).ExtractErr()
}
