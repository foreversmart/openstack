package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Networker interface {
	Create(opts *options.CreateNetworkOpts) (network *models.NetworkModel, err error)
	All() (networks []*models.NetworkModel, err error)
	AllByParams(opts *options.ListNetworkOpt) (networks []*models.NetworkModel, err error)
	Show(id string) (network *models.NetworkModel, err error)
	Delete(id string) error
}
