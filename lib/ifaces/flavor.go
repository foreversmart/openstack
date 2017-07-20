package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Flavor interface {
	All() (flavors []*models.FlavorModel, err error)
	AllByParams(opts *options.ListFlavorsOpts) (flavors []*models.FlavorModel, err error)
	Create(opts *options.CreateFlavorOpts) (flavor *models.FlavorModel, err error)
	Show(id string) (flavor *models.FlavorModel, err error)
	Delete(id string) (err error)
}
