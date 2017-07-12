package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Volumer interface {
	Create(opts *options.CreateVolumeOpts) (volume *models.VolumeModel, err error)
	All() (volumes []*models.VolumeModel, err error)
	AllByParams(opts *options.ListVolumeOpts) (volumes []*models.VolumeModel, err error)
	Show(id string) (volume *models.VolumeModel, err error)
	Update(id string, opts *options.UpdateVolumeOpts) (volume *models.VolumeModel, err error)
	Resize(id string, opts *options.ResizeVolumeOpts) error
	Delete(id string) error
}
