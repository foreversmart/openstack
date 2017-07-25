package options

import (
	"net/url"
	"strconv"
)

type ListFloatingIPOpts struct {
	AllTenants        *string `mapstructure:"all_tenants" json:"all_tenants"`
	FloatingNetworkID *string `mapstructure:"floating_network_id" json:"floating_network_id"`
	PortID            *string `mapstructure:"port_id" json:"port_id"`
	FixedIP           *string `mapstructure:"fixed_ip_address" json:"fixed_ip_address"`
	FloatingIP        *string `mapstructure:"floating_ip_address" json:"floating_ip_address"`
	TenantID          *string `mapstructure:"tenant_id" json:"tenant_id"`
	Limit             *int    `mapstructure:"limit" json"limit"`
	Marker            *string `mapstructure:"marker" json:"marker"`
	SortKey           *string `mapstructure:"sort_key" json:"sort_key"`
	SortDir           *string `mapstructure:"sort_dir" json:"sort_dir"`
}

func (opts *ListFloatingIPOpts) IsValid() bool {
	return true
}

func (opts *ListFloatingIPOpts) ToQuery() url.Values {
	options := url.Values{}

	if opts != nil {
		if opts.AllTenants != nil {
			options.Add("all_tenants", *opts.AllTenants)
		}

		if opts.FloatingNetworkID != nil {
			options.Add("floating_network_id", *opts.FloatingNetworkID)
		}

		if opts.PortID != nil {
			options.Add("port_id", *opts.PortID)
		}

		if opts.FixedIP != nil {
			options.Add("fixed_ip_address", *opts.FixedIP)
		}

		if opts.FloatingIP != nil {
			options.Add("floating_ip_address", *opts.FloatingIP)
		}

		if opts.TenantID != nil {
			options.Add("tenant_id", *opts.TenantID)
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

/**
 * used to create floatingip
 */
type CreateFloatingIPOpts struct {
	ProjectID         *string `mapstructure:"project_id" json:"project_id"`
	TenantID          *string `mapstructure:"tenant_id" json:"tenant_id"`
	FloatingNetworkID *string `mapstructure:"floating_network_id" json:"floating_network_id"`

	// following are optional ref: https://developer.openstack.org/api-ref/networking/v2/index.html
	PortID      *string `mapstructure:"port_id" json:"port_id"`
	FloatingIP  *string `mapstructure:"floating_ip_address" json:"floating_ip_address"`
	FixedIP     *string `mapstructure:"fixed_ip_address" json:"fixed_ip_address"`
	SubnetID    *string `mapstructure:"subnet_id" json:"subnet_id"`
	Description *string `mapstructure:"description" json:"description"`
}

func (opts *CreateFloatingIPOpts) IsValid() bool {
	return opts != nil &&
		opts.FloatingNetworkID != nil &&
		opts.ProjectID != nil &&
		opts.TenantID != nil
}

func (opts *CreateFloatingIPOpts) ToPayload() interface{} {
	type payload struct {
		FloatingIP *CreateFloatingIPOpts `json:"floatingip"`
	}

	return payload{
		opts,
	}
}

/**
 * used to update floatingip
 */
type UpdateFloatingIPOpts struct {
	PortID      *string `mapstructure:"port_id" json:"port_id"`
	FixedIP     *string `mapstructure:"fixed_ip_address" json:"fixed_ip_address"`
	Description *string `mapstructure:"description" json:"description"`
}

func (opts *UpdateFloatingIPOpts) IsValid() bool {
	return opts != nil
}

func (opts *UpdateFloatingIPOpts) ToPayload() interface{} {
	type payload struct {
		FloatingIP *UpdateFloatingIPOpts `json:"floatingip"`
	}

	return payload{
		opts,
	}
}
