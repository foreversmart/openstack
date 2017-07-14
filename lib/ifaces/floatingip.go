package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type FloatingIPer interface {
	Create(opts *options.CreateFloatingIPOpts) (ip *models.FloatingIPModel, err error)
	All() (ips []*models.FloatingIPModel, err error)
	Show(id string) (ip *models.FloatingIPModel, err error)
	Update(id string, opts *options.UpdateFloatingIPOpts) (ip *models.FloatingIPModel, err error)
	Delete(id string) error
}
