package subnet

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
)

type Subnet struct {
	Client ifaces.Openstacker

	_ bool
}

const (
	SubnetUrl = "subnets"
)

func New(client ifaces.Openstacker) *Subnet {
	return &Subnet{
		Client: client,
	}
}

func (n *Subnet) Create(opts *subnets.CreateOpts) (subnet *models.SubnetModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return subnet, err
	}

	result := subnets.Create(client, opts)
	return models.ExtractSubnet(result.Result)
}

func (n *Subnet) All() (subnetInfos []*models.SubnetModel, err error) {
	return n.AllByParams(nil)
}

func (n *Subnet) AllByParams(opts *subnets.ListOpts) (subnetInfos []*models.SubnetModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	page, err := subnets.List(client, opts).AllPages()
	if err != nil {
		return
	}

	return models.ExtractSubnetsByBody(page.GetBody())
}

func (n *Subnet) Show(id string) (info *models.SubnetModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	result := subnets.Get(client, id)
	return models.ExtractSubnet(result.Result)
}

func (n *Subnet) Update(id string, opts *options.UpdateSubnetOpts) (groupInfo *models.SubnetModel, err error) {
	client, err := n.Client.NetworkClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, err = client.Put(client.ServiceURL(SubnetUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractSubnet(result)
}

func (n *Subnet) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return err
	}

	return subnets.Delete(client, id).ExtractErr()
}
