package ifaces

import (
	"net/http"

	"github.com/rackspace/gophercloud"
	"github.com/kirk-enterprise/openstack/lib/auth"
	"github.com/kirk-enterprise/openstack/lib/models"
)

type Openstacker interface {
	OpenstackAuther
	OpenstackClienter

	ProjectID() string
	PrependSharedNetwork(networks []*models.NetworkModel) []*models.NetworkModel
	PrependSharedSubnet(subnets []*models.SubnetModel) []*models.SubnetModel
}

type OpenstackAuther interface {
	IsAuthorized() bool
	Auth(opts auth.AuthOptions) error
	AuthByPassword(opts auth.AuthOptions) (err error)
	AuthByToken(opts auth.AuthOptions) (err error)
}

type OpenstackClienter interface {
	ProviderClient() *gophercloud.ProviderClient
	ServiceClient(endpoint string) *gophercloud.ServiceClient

	WithHTTPClient(client *http.Client)

	// admin clients
	AdminIdentityClient() (client *gophercloud.ServiceClient, err error)
	AdminIdentityClientV3() (client *gophercloud.ServiceClient, err error)

	// user clients
	IdentityClient() (client *gophercloud.ServiceClient, err error)
	IdentityClientV3() (client *gophercloud.ServiceClient, err error)
}
