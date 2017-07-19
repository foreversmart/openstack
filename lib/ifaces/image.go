package ifaces

import (
	"io"

	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Imager interface {
	Create(opts *options.CreateImagesOpts) (image *models.ImageModel, err error)
	All() (images []*models.ImageModel, err error)
	AllByParams(opts *options.ListImagesOpts) (imageModels []*models.ImageModel, err error)
	Show(id string) (image *models.ImageModel, err error)
	Update(id string, opts *options.UpdateImagesOpts) (imageModel *models.ImageModel, err error)
	Delete(id string) error
	Upload(id string, data io.Reader) error
	Download(id string) (data io.Reader, err error)
}
