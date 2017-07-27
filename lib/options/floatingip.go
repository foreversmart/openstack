package options

import (
	"net/url"
	"strconv"
)

type ListFloatingIPOpts struct {
	AllTenants        *string `json:"all_tenants"`
	FloatingNetworkID *string `json:"floating_network_id"`
	PortID            *string `json:"port_id"`
	FixedIP           *string `json:"fixed_ip_address"`
	FloatingIP        *string `json:"floating_ip_address"`
	TenantID          *string `json:"tenant_id"`
	Limit             *int    `json"limit"`
	Marker            *string `json:"marker"`
	SortKey           *string `json:"sort_key"`
	SortDir           *string `json:"sort_dir"`
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
	ProjectID         *string `json:"project_id,omitempty"`
	TenantID          *string `json:"tenant_id,omitempty"`
	FloatingNetworkID *string `json:"floating_network_id,omitempty"`

	// following are optional ref: https://developer.openstack.org/api-ref/networking/v2/index.html
	PortID      *string `json:"port_id,omitempty"`
	FloatingIP  *string `json:"floating_ip_address,omitempty"`
	FixedIP     *string `json:"fixed_ip_address,omitempty"`
	SubnetID    *string `json:"subnet_id,omitempty"`
	Description *string `json:"description,omitempty"`
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
	PortID      *string `json:"port_id,omitempty"`
	FixedIP     *string `json:"fixed_ip_address,omitempty"`
	Description *string `json:"description,omitempty"`
	Unbind      *bool   `json:"-"`
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
