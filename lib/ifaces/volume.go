package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Volumer interface {
	All() (volumes []*models.VolumeModel, err error)
	AllByParams(opts *options.ListVolumeOpts) (volumes []*models.VolumeModel, err error)
	Create(param *options.CreateVolumeOpts) (id string, err error)
	Show(id string) (*models.VolumeModel, error)
	Delete(id string) error
	Resize(id string, newSize int) error
	Update(id string, param *options.UpdateVolumeOpts) error
}
