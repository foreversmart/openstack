package internal

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

func (os *Openstack) IdentityClient() (client *gophercloud.ServiceClient, err error) {
	return openstack.NewIdentityV2(os.client), nil
}

func (os *Openstack) IdentityClientV3() (client *gophercloud.ServiceClient, err error) {
	return openstack.NewIdentityV3(os.client), nil
}
