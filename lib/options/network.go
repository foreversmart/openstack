package options

import "net/url"

type ListNetworkOpt struct {
	AllTenants *string `json:"all_tenants"`
	ProjectId  *string `json:"tenant_id"`
	Status     *string `json:"status"`
}

func (opts *ListNetworkOpt) IsValid() bool {
	return true
}

func (opts *ListNetworkOpt) ToQuery() (options url.Values) {
	options = url.Values{}
	if opts == nil {
		return
	}

	if opts.AllTenants != nil {
		options.Add("all_tenants", *opts.AllTenants)
	}

	if opts.ProjectId != nil {
		options.Add("tenant_id", *opts.ProjectId)
	}

	if opts.Status != nil {
		options.Add("status", *opts.Status)
	}

	return
}

type CreateNetworkOpts struct {
	Name         *string `json:"name,omitempty"`
	TenantID     *string `json:"tenant_id,omitempty"`
	Shared       *bool   `json:"shared,omitempty"`
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
}

func (opts *CreateNetworkOpts) ToPayload() interface{} {
	type request struct {
		Network *CreateNetworkOpts `json:"network"`
	}

	return request{
		Network: opts,
	}
}

func (opts *CreateNetworkOpts) IsValid() bool {
	return true
}
