package options

import (
	"net/url"
	"strconv"
)

type CreateSecruleOpts struct {
	RemoteGroupID  *string `json:"remote_group_id"`
	Direction      *string `json:"direction"`
	Protocol       *string `json:"protocol"`
	EtherType      *string `json:"ethertype"`
	PortRangeMin   *int    `json:"port_range_min"`
	PortRangeMax   *int    `json:"port_range_max"`
	SecGroupID     *string `json:"security_group_id"`
	RemoteIPPrefix *string `json:"remote_ip_prefix"`
	Description    *string `json:"description"`
}

func (opts *CreateSecruleOpts) IsValid() bool {
	return opts != nil && opts.Direction != nil && opts.RemoteIPPrefix != nil
}

func (opts *CreateSecruleOpts) ToPayload() interface{} {
	type request struct {
		Secrule *CreateSecruleOpts `json:"secrule"`
	}

	return request{
		Secrule: opts,
	}
}

type ListSecRuleOpts struct {
	ID             *string `json:"id"`
	RemoteGroupID  *string `json:"remote_group_id"`
	Direction      *string `json:"direction"`
	Protocol       *string `json:"protocol"`
	EtherType      *string `json:"ethertype"`
	PortRangeMin   *int    `json:"port_range_min"`
	PortRangeMax   *int    `json:"port_range_max"`
	SecGroupID     *string `json:"security_group_id"`
	TenantID       *string `json:"tenant_id"`
	ProjectID      *string `json:"project_id"`
	RemoteIPPrefix *string `json:"remote_ip_prefix"`
	Description    *string `json:"description"`
}

func (opts *ListSecRuleOpts) IsValid() bool {
	return true
}

func (opts *ListSecRuleOpts) ToQuery() url.Values {
	options := url.Values{}

	if opts != nil {
		if opts.ID != nil {
			options.Add("id", *opts.ID)
		}

		if opts.RemoteGroupID != nil {
			options.Add("remote_group_id", *opts.RemoteGroupID)
		}

		if opts.Direction != nil {
			options.Add("direction", *opts.Direction)
		}

		if opts.Protocol != nil {
			options.Add("protocol", *opts.Protocol)
		}

		if opts.EtherType != nil {
			options.Add("ethertype", *opts.EtherType)
		}

		if opts.PortRangeMin != nil {
			options.Add("port_range_min", strconv.Itoa(*opts.PortRangeMin))
		}

		if opts.PortRangeMax != nil {
			options.Add("port_range_max", strconv.Itoa(*opts.PortRangeMax))
		}

		if opts.SecGroupID != nil {
			options.Add("security_group_id", *opts.SecGroupID)
		}

		if opts.TenantID != nil {
			options.Add("tenant_id", *opts.TenantID)
		}

		if opts.ProjectID != nil {
			options.Add("project_id", *opts.ProjectID)
		}

		if opts.RemoteIPPrefix != nil {
			options.Add("remote_ip_prefix", *opts.RemoteIPPrefix)
		}

		if opts.Description != nil {
			options.Add("description", *opts.Description)
		}
	}

	return options
}
