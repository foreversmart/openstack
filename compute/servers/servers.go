package servers

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	ServersUrl = "servers"
)

type Servers struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Servers {
	return &Servers{
		Client: client,
	}
}
func (server *Servers) All() (serverInfos []*models.ServersModel, err error) {
	return server.AllByParams(nil)
}

func (server *Servers) AllByParams(opts *options.ListServersOpts) (Serverss []*models.ServersModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}
	client, err := server.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServersUrl, "detail")+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServers(result)
}
