package subnet

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
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

func (n *Subnet) Create(opts *options.CreateSubnetOpts) (subnet *models.SubnetModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return subnet, err
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(SubnetUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractSubnet(res)
}

func (n *Subnet) All() (subnetInfos []*models.SubnetModel, err error) {
	return n.AllByParams(nil)
}

func (n *Subnet) AllByParams(opts *options.ListSubnetOpts) (subnetInfos []*models.SubnetModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := n.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(SubnetUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractSubnets(result)
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

func (n *Subnet) Update(id string, opts *options.UpdateSubnetOpts) (info *models.SubnetModel, err error) {
	if id == "" || !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

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
