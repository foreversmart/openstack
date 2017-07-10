package network

import (
	"net/url"

	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
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

func (n *Network) CreateNetwork(name, tenantId string, shared, adminStateUp bool) (network *models.NetworkModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return network, err
	}

	opts := networks.CreateOpts{
		Name:         name,
		AdminStateUp: &adminStateUp,
		Shared:       &shared,
		TenantID:     tenantId,
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

func (n *Network) AllByParams(param *options.NetworkQueryOpt) (infos []*models.NetworkModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return
	}

	opts := url.Values{}
	if param.AllTenants != "" {
		opts.Add("all_tenants", param.AllTenants)
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL("networks")+"?"+opts.Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if result.Err != nil {
		err = result.Err
		return
	}

	return models.ExtractNetworks(result)
}

func (n *Network) CreateSubnet(name, networkId, cidr, tenantId string, ipVersion int) (subnet *models.SubnetModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return subnet, err
	}

	opts := subnets.CreateOpts{
		Name:      name,
		NetworkID: networkId,
		CIDR:      cidr,
		TenantID:  tenantId,
		IPVersion: ipVersion,
	}

	result := subnets.Create(client, opts)
	return models.ExtractSubnetByResult(result.Result)
}

func (n *Network) AllSubnets() (subnetInfos []*models.SubnetModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return subnetInfos, err
	}

	subnetInfos = make([]*models.SubnetModel, 0, 5)

	opts := subnets.ListOpts{
		TenantID: n.Client.ProjectID(),
	}

	page, err := subnets.List(client, opts).AllPages()
	if err != nil {
		return
	}
	subnetInfos, err = models.ExtractSubnetsByPage(page)
	if err != nil {
		return
	}

	return subnetInfos, err
}

func (n *Network) Query(id string) (info *models.NetworkModel, err error) {
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

func (n *Network) AllSubnetsByID(networkId string) (subnetIds []string, err error) {
	networkInfo, err := n.Query(networkId)
	if err != nil || networkInfo == nil {
		return
	}

	return networkInfo.Subnets, err
}

func (n *Network) DeleteNetwork(networkId string) error {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return err
	}

	return networks.Delete(client, networkId).ExtractErr()
}
