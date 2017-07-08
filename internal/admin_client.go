package internal

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

func (os *Openstack) AdminIdentityClient() (client *gophercloud.ServiceClient, err error) {
	return openstack.NewIdentityAdminV2(os.client, gophercloud.EndpointOpts{})
}

func (os *Openstack) AdminIdentityClientV3() (client *gophercloud.ServiceClient, err error) {
	return openstack.NewIdentityAdminV3(os.client, gophercloud.EndpointOpts{})
}
