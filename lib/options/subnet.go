package options

import (
	"net/url"
	"strconv"

	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
)

// CreateOpts represents the attributes used when creating a new subnet.
type CreateSubnetOpts struct {
	// Required
	NetworkID *string
	CIDR      *string
	// Optional
	Name            *string
	TenantID        *string
	AllocationPools []*models.AllocationPool
	GatewayIP       *string
	NoGateway       bool
	IPVersion       int
	EnableDHCP      *bool
	DNSNameservers  []*string
	HostRoutes      []*models.HostRoute
}

func (opts *CreateSubnetOpts) IsValid() bool {
	return opts.NetworkID != nil && opts.CIDR != nil
}

func (opts *CreateSubnetOpts) ToPayload() interface{} {
	type request struct {
		Subnet *CreateSubnetOpts `json:"subnet"`
	}

	return request{
		Subnet: opts,
	}
}

type UpdateSubnetOpts struct {
	Name           *string                  `json:"name"`
	EnableDHCP     *bool                    `json:"enable_dhcp"`
	DNSNameservers []*string                `json:"dns_nameservers "`
	AllocPools     []*models.AllocationPool `json:"allocation_pools "`
	HostRoutes     []*models.HostRoute      `json:"host_routes"`
	GatewayIP      *string                  `json:"gateway_ip"`
	Description    *string                  `json:"description"`
}

func (opts *UpdateSubnetOpts) IsValid() bool {
	return true
}

func (opts *UpdateSubnetOpts) ToPayload() interface{} {
	type request struct {
		Subnet *UpdateSubnetOpts `json:"subnet"`
	}

	return request{
		Subnet: opts,
	}
}

type ListSubnetOpts struct {
	Name       *string `json:"name"`
	EnableDHCP *bool   `json:"enable_dhcp"`
	NetworkID  *string `json:"network_id"`
	TenantID   *string `json:"tenant_id"`
	IPVersion  *int    `json:"ip_version"`
	GatewayIP  *string `json:"gateway_ip"`
	CIDR       *string `json:"cidr"`
	ID         *string `json:"id"`
	Limit      *int    `json:"limit"`
	Marker     *string `json:"marker"`
	SortKey    *string `json:"sort_key"`
	SortDir    *string `json:"sort_dir"`
}

func (opts *ListSubnetOpts) IsValid() bool {
	return true
}

func (opts *ListSubnetOpts) ToQuery() url.Values {
	options := url.Values{}

	if opts != nil {
		if opts.Name != nil {
			options.Add("name", *opts.Name)
		}

		if opts.EnableDHCP != nil {
			options.Add("enable_dhcp", strconv.FormatBool(*opts.EnableDHCP))
		}

		if opts.NetworkID != nil {
			options.Add("network_id", *opts.NetworkID)
		}

		if opts.TenantID != nil {
			options.Add("tenant_id", *opts.TenantID)
		}

		if opts.IPVersion != nil {
			options.Add("ip_version", strconv.Itoa(*opts.IPVersion))
		}

		if opts.GatewayIP != nil {
			options.Add("gateway_ip", *opts.GatewayIP)
		}

		if opts.CIDR != nil {
			options.Add("cidr", *opts.CIDR)
		}

		if opts.ID != nil {
			options.Add("id", *opts.ID)
		}

		if opts.Limit != nil {
			options.Add("limit", strconv.Itoa(*opts.Limit))
		}

		if opts.Marker != nil {
			options.Add("marker", *opts.Marker)
		}

		if opts.SortKey != nil {
			options.Add("sort_key", *opts.SortKey)
		}

		if opts.SortDir != nil {
			options.Add("sort_dir", *opts.SortDir)
		}
	}

	return options
}
