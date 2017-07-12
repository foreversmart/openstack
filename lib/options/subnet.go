package options

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/models"

type UpdateSubnetOpts struct {
	Name           *string                  `json:"name"`
	EnableDHCP     *bool                    `json:"enable_dhcp"`
	DNSNameservers []*string                `json:"dns_nameservers "`
	AllocPools     []*models.AllocationPool `json:"allocation_pools "`
	HostRoutes     []*models.HostRoute      `json:"host_routes"`
	GatewayIP      *string                  `json:"gateway_ip"`
	Description    *string                  `json:"description"`
}

func (opts *UpdateSubnetOpts) ToPayload() interface{} {
	type request struct {
		Subnet *UpdateSubnetOpts `json:"subnet"`
	}

	return request{
		Subnet: opts,
	}
}
