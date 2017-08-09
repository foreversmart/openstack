package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Volumer interface {
	Create(opts *options.CreateVolumeOpts) (volume *models.VolumeModel, err error)
	All() (volumes []*models.VolumeModel, err error)
	AllByParams(opts *options.ListVolumeOpts) (volumes []*models.VolumeModel, err error)
	Show(opts *options.ShowVolumeOpts) (volume *models.VolumeModel, err error)
	Update(id string, opts *options.UpdateVolumeOpts) (volume *models.VolumeModel, err error)
	Resize(id string, opts *options.ResizeVolumeOpts) error
	Reset(id string, opts *options.ResetVolumeOpts) error
	Delete(id string) error
}
