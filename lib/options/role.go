package options

import (
	"net/url"
)

type ListRoleOpts struct {
	// Filters the response by a domain ID, optional
	DomainId *string `json:"id,omitempty"`

	// Filters the response by a role name, optional
	Name *string `json:"name,omitempty"`
}

type CreateRoleOpts struct {
	//The ID of the domain, optional
	DomainId *string `json:"id,omitempty"`

	//The role name, optional
	Name *string `json:"name,omitempty"`
}

type UpdateRoleOpts struct {
	//The new role name, optional
	Name *string `json:"name,omitempty"`
}

func (opts *ListRoleOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts != nil {
		if opts.Name != nil {
			param.Add("name", *opts.Name)
		}

		if opts.DomainId != nil {
			param.Add("domain_id", *opts.DomainId)
		}
	}

	return param
}

func (opts *CreateRoleOpts) ToPayLoad() interface{} {
	type payload struct {
		Role *CreateRoleOpts `json:"role"`
	}

	return payload{
		Role: opts,
	}
}

func (opts *UpdateRoleOpts) ToPayLoad() interface{} {
	type payload struct {
		Role *UpdateRoleOpts `json:"role"`
	}

	return payload{
		Role: opts,
	}
}
