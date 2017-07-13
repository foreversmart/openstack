package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Serverser interface {
	Create(opts options.CreateServersOpts) (credential *models.ServersModel, err error)
	All() (serverInfos []*models.ServersModel, err error)
	AllByParams(opts *options.ListServersOpts) (credentials []*models.ServersModel, err error)
	Show(id string) (credential *models.ServersModel, err error)
	Update(id string, opts options.UpdateServersOpts) (credential *models.ServersModel, err error)
	Delete(id string) (err error)
}
