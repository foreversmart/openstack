package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type SubnetModel struct {
	ID              string           `mapstructure:"id" json:"id"`
	NetworkID       string           `mapstructure:"network_id" json:"network_id"`
	Name            string           `mapstructure:"name"  json:"name"`
	IpVersion       int              `mapstructure:"ip_version" json:"ip_version"`
	Cidr            string           `mapstructure:"cidr" json:"cidr"`
	GatewayIp       string           `mapstructure:"gateway_ip" json:"gateway_ip"`
	DnsNameservers  []string         `mapstructure:"dns_nameservers" json:"dns_nameservers"`
	AllocPools      []AllocationPool `mapstructure:"allocation_pools"  json:"allocation_pools"`
	HostRoutes      []HostRoute      `mapstructure:"host_routes" json:"host_routes"`
	EnableDhcp      bool             `mapstructure:"enable_dhcp" json:"enable_dhcp"`
	TenantID        string           `mapstructure:"tenant_id" json:"tenant_id"`
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

func ExtractSubnet(result gophercloud.Result) (subnet *SubnetModel, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	var response struct {
		Subnet *SubnetModel `mapstructure:"subnet" json:"subnet"`
	}

	err = mapstructure.Decode(result.Body, &response)
	return response.Subnet, err
}

func ExtractSubnetsByBody(body interface{}) (networks []*SubnetModel, err error) {
	var response struct {
		NetworkInfos []*SubnetModel `mapstructure:"subnets" json:"subnets"`
	}

	err = mapstructure.Decode(body, &response)
	return response.NetworkInfos, err
}
