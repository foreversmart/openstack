package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
)

type Networker interface {
	Create(opts *networks.CreateOpts) (network *models.NetworkModel, err error)
	All() (networks []*models.NetworkModel, err error)
	AllByParams(opts *options.ListNetworkOpt) (networks []*models.NetworkModel, err error)
	Show(id string) (network *models.NetworkModel, err error)
	Delete(id string) error
}
