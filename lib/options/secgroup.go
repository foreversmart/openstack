package options

import "net/url"

type ListSecurityGroupsOpts struct {
	Fields     *string `json:"fields"`
	AllTenants *string `json:"all_tenants"`
	TenantID   *string `json:"tenant_id"`
	Limit      *string `json:"limit"`
	Marker     *string `json:"marker"`
}

func (opts *ListSecurityGroupsOpts) IsValid() bool {
	return true
}

func (opts *ListSecurityGroupsOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts == nil {
		return
	}

	if opts.Fields != nil {
		options.Add("fields", *opts.Fields)
	}
	if opts.AllTenants != nil {
		options.Add("all_tenants", *opts.AllTenants)
	}
	if opts.TenantID != nil {
		options.Add("tenant_id", *opts.TenantID)
	}
	if opts.Limit != nil {
		options.Add("limit", *opts.Limit)
	}
	if opts.Marker != nil {
		options.Add("marker", *opts.Marker)
	}

	return
}

type CreateSecurityGroupOpts struct {
	//The ID of the project.
	TenantID *string `json:"tenant_id,omitempty"`

	// The ID of the project.
	ProjectID *string `json:"project_id,omitempty"`

	// Human-readable name of the resource.
	Name *string `json:"name,omitempty"`

	// A human-readable description for the resource. Default is an empty string
	// Optional
	Description *string `json:"description,omitempty"`
}

func (opts *CreateSecurityGroupOpts) IsValid() bool {
	return opts != nil && opts.TenantID != nil && opts.ProjectID != nil && opts.Name != nil
}

func (opts *CreateSecurityGroupOpts) ToPayload() interface{} {
	type payload struct {
		SecurityGroup *CreateSecurityGroupOpts `json:"security_group"`
	}

	return payload{
		SecurityGroup: opts,
	}
}

type ShowSecurityGroupOpts struct {
	// Show detailed information, Optional
	Verbose *bool `json:"verbose,omitempty"`

	// The fields that you want the server to return, Optional
	Fields *string `json:"string,omitempty"`
}

func (opts *ShowSecurityGroupOpts) IsValid() bool {
	return true
}

func (opts *ShowSecurityGroupOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts == nil {
		return
	}

	if opts.Verbose != nil {
		if *opts.Verbose {
			param.Add("verbose", "true")
		} else {
			param.Add("verbose", "false")
		}
	}

	if opts.Fields != nil {
		param.Add("fields", *opts.Fields)
	}

	return param
}

type UpdateSecurityGroupOpts struct {
	//Human-readable name of the resource.
	Name *string `json:"name,omitempty"`

	// A human-readable description for the resource, Optional
	Description *string `json:"description,omitempty"`
}

func (opts *UpdateSecurityGroupOpts) IsValid() bool {
	return opts != nil && opts.Name != nil
}

func (opts *UpdateSecurityGroupOpts) ToPayload() interface{} {
	type payload struct {
		SecurityGroup *UpdateSecurityGroupOpts `json:"security_group"`
	}

	return payload{
		SecurityGroup: opts,
	}
}
