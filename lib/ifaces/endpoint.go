package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Endpointer interface {
	All() (endpoints []*models.EndpointModel, err error)
	AllByParams(opts *options.ListEndpointOpts) (services []*models.EndpointModel, err error)
	Create(opts options.CreateEndpointOpts) (endpoint *models.EndpointModel, err error)
	Show(id string) (endpoint *models.EndpointModel, err error)
	Update(id string, opts options.UpdateEndpointOpts) (endpoint *models.EndpointModel, err error)
	Delete(id string) (err error)
}
