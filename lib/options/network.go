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
