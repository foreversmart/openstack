package options

import "net/url"

type ListSecurityGroupsOpts struct {
	Fields *string `json:"fields,omitempty"`
}

func (opts *ListSecurityGroupsOpts) IsValid() bool {
	return true
}

func (opts *ListSecurityGroupsOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts == nil {
		return
	}

	if opts.Fields != nil {
		param.Add("fields", *opts.Fields)
	}

	return param
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

func (opts *CreateSecurityGroupOpts) ToPayLoad() interface{} {
	type payload struct {
		SecurityGroup *CreateSecurityGroupOpts `json:"security_group"`
	}

	return payload{
		SecurityGroup: opts,
	}
}
