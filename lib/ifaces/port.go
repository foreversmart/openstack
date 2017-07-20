package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Porter interface {
	All() ([]*models.PortModel, error)
	Show(portId string) (*models.PortModel, error)
	Create(opts *options.CreatePortOpts) (*models.PortModel, error)
	Update(id string, opts *options.UpdatePortOpts) (*models.PortModel, error)
	Delete(id string) error
}
