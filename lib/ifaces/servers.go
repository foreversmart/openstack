package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Serverser interface {
	Create(opts options.CreateServersOpts) (server *models.ServersModel, err error)
	All() (servers []*models.ServersModel, err error)
	AllByParams(opts *options.ListServersOpts) (servers []*models.ServersModel, err error)
	Show(id string) (server *models.ServersModel, err error)
	Update(id string, opts options.UpdateServersOpts) (server *models.ServersModel, err error)
	Delete(id string) (err error)
}
