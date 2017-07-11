package network

import (
	"net/url"

	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
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

func (n *Network) Create(opts *networks.CreateOpts) (network *models.NetworkModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return network, err
	}

	result := networks.Create(client, opts)
	return models.ExtractNetworkByResult(result.Result)
}

func (n *Network) All() (infos []*models.NetworkModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return
	}

	opts := networks.ListOpts{
		TenantID: n.Client.ProjectID(),
		Status:   "ACTIVE",
	}

	page, err := networks.List(client, opts).AllPages()
	if err != nil {
		return
	}

	return models.ExtractNetworksByPage(page)
}

func (n *Network) AllByParams(opts *options.NetworkQueryOpt) (infos []*models.NetworkModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return
	}

	params := url.Values{}
	if opts.AllTenants != "" {
		params.Add("all_tenants", opts.AllTenants)
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL("networks")+"?"+params.Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if result.Err != nil {
		err = result.Err
		return
	}

	return models.ExtractNetworks(result)
}

func (n *Network) Show(id string) (info *models.NetworkModel, err error) {
	// find the network in the admin created networks
	networkInfos := make([]*models.NetworkModel, 0)
	for _, value := range networkInfos {
		if value.Id == id {
			info = value
			break
		}
	}

	if info == nil {
		client, err := n.Client.NetworkClient()
		if err != nil {
			return nil, err
		}

		result := networks.Get(client, id)
		return models.ExtractNetworkByResult(result.Result)
	}
	return
}

func (n *Network) Delete(networkId string) error {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return err
	}

	return networks.Delete(client, networkId).ExtractErr()
}
