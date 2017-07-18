package options

import "net/url"

type ListNetworkOpt struct {
	AllTenants *string `json:"all_tenants"`
	ProjectId  *string `json:"project_id"`
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
		options.Add("project_id", *opts.ProjectId)
	}

	return
}

type CreateNetworkOpts struct {
	Name         *string `json:"name"`
	TenantID     *string `json:"tenant_id"`
	Shared       *bool   `json:"shared"`
	AdminStateUp *bool   `json:"admin_state_up"`
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
