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

func (os *Openstack) NetworkClient() (client *gophercloud.ServiceClient, err error) {
	return openstack.NewNetworkV2(os.client, gophercloud.EndpointOpts{})
}

func (os *Openstack) VolumeClient() (client *gophercloud.ServiceClient, err error) {
	opts := gophercloud.EndpointOpts{
		Type:         "volumev2",
		Availability: gophercloud.AvailabilityPublic,
	}

	endpoint, err := os.client.EndpointLocator(opts)
	if err != nil {
		return
	}

	client = os.ServiceClient(endpoint)
	return
}
