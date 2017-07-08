package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Servicer interface {
	All() (services []*models.ServiceModel, err error)
	AllByParams(opts *options.ListServiceOpts) (services []*models.ServiceModel, err error)
	Create(opts options.CreateServiceOpts) (service *models.ServiceModel, err error)
	Show(id string) (service *models.ServiceModel, err error)
	Update(id string, opts options.UpdateServiceOpts) (service *models.ServiceModel, err error)
	Delete(id string) (err error)
}
