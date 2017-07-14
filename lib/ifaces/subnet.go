package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Subneter interface {
	Create(opts *options.CreateSubnetOpts) (subnet *models.SubnetModel, err error)
	All() (subnets []*models.SubnetModel, err error)
	AllByParams(opts *options.ListSubnetOpts) (subnets []*models.SubnetModel, err error)
	Show(id string) (subnet *models.SubnetModel, err error)
	Update(id string, opts *options.UpdateSubnetOpts) (subnet *models.SubnetModel, err error)
	Delete(id string) error
}
