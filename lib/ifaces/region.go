package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Regioner interface {
	All() (regions []*models.RegionModel, err error)
	AllByParams(opts *options.ListRegionOpts) (regions []*models.RegionModel, err error)
	Create(opts *options.CreateRegionOpts) (region *models.RegionModel, err error)
	Show(id string) (region *models.RegionModel, err error)
	Update(id string, opts *options.UpdateRegionOpts) (region *models.RegionModel, err error)
	Delete(id string) (err error)
}
