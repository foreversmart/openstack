package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type SubnetModel struct {
	ID              string           `mapstructure:"id" json:"id"`
	NetworkID       string           `mapstructure:"network_id" json:"network_id"`
	Name            string           `mapstructure:"name"  json:"name"`
	IPVersion       int              `mapstructure:"ip_version" json:"ip_version"`
	CIDR            string           `mapstructure:"cidr" json:"cidr"`
	GatewayIP       string           `mapstructure:"gateway_ip" json:"gateway_ip"`
	DNSNameservers  []string         `mapstructure:"dns_nameservers" json:"dns_nameservers"`
	AllocPools      []AllocationPool `mapstructure:"allocation_pools"  json:"allocation_pools"`
	HostRoutes      []HostRoute      `mapstructure:"host_routes" json:"host_routes"`
	EnableDHCP      bool             `mapstructure:"enable_dhcp" json:"enable_dhcp"`
	TenantID        string           `mapstructure:"tenant_id" json:"tenant_id"`
	ProjectID       string           `mapstructure:"project_id" json:"project_id"`
	Ipv6RaMode      bool             `mapstructure:"ipv6_ra_mode" json:"ipv6_ra_mode"`
	Ipv6AddressMode bool             `mapstructure:"ipv6_address_mode" json:"ipv6_address_mode"`
	CreatedAt       string           `mapstructure:"created_at" json:"created_at"`
}

type AllocationPool struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// HostRoute represents a route that should be used by devices with IPs from
// a subnet (not including local subnet route).
type HostRoute struct {
	DestinationCIDR string `mapstructure:"destination" json:"destination"`
	NextHop         string `mapstructure:"nexthop" json:"nexthop"`
}

func ExtractSubnet(r gophercloud.Result) (subnet *SubnetModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var resp struct {
		Subnet *SubnetModel `mapstructure:"subnet" json:"subnet"`
	}

	err = mapstructure.Decode(r.Body, &resp)
	if err == nil {
		subnet = resp.Subnet
	}
	return
}

func ExtractSubnetsByBody(body interface{}) (networks []*SubnetModel, err error) {
	var resp struct {
		NetworkInfos []*SubnetModel `mapstructure:"subnets" json:"subnets"`
	}

	err = mapstructure.Decode(body, &resp)
	if err == nil {
		networks = resp.NetworkInfos
	}
	return
}
