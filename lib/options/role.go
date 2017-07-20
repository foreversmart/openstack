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

func (opts *ListRoleOpts) ToQuery() (param url.Values) {
	param = url.Values{}
	if opts == nil {
		return
	}

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

type CreateRoleOpts struct {
	//The ID of the domain, optional
	DomainId *string `json:"id,omitempty"`

	//The role name
	Name string `json:"name"`
}

func (opts *CreateRoleOpts) ToPayLoad() interface{} {
	type payload struct {
		Role *CreateRoleOpts `json:"role"`
	}

	return payload{
		Role: opts,
	}
}

type UpdateRoleOpts struct {
	//The new role name, optional
	Name *string `json:"name,omitempty"`

	//The new role domain, optional
	DomainID *string `json:"domain_id,omitempty"`
}

func (opts *UpdateRoleOpts) ToPayLoad() interface{} {
	type payload struct {
		Role *UpdateRoleOpts `json:"role"`
	}

	return payload{
		Role: opts,
	}
}

type ListRoleAssignmentOpts struct {
	Effective      *bool   `json:"effective,omitempty"`
	IncludeNames   *bool   `json:"include_names,omitempty"`
	IncludeSubtree *bool   `json:"include_subtree,omitempty"`
	GroupID        *string `json:"group_id,omitempty"`
	RoleID         *string `json:"role_id,omitempty"`
	DomainID       *string `json:"scope_domain_id,omitempty"`
	ScopeProjectID *string `json:"scope_project_id ,omitempty"`
	UserID         *string `json:"user_id ,omitempty"`
}

func (opts *ListRoleAssignmentOpts) ToQuery() (param url.Values) {
	param = url.Values{}
	if opts == nil {
		return
	}

	if opts != nil {
		if opts.Effective != nil {
			param.Add("effective", "true")
		}

		if opts.IncludeNames != nil {
			param.Add("include_names", "true")
		}
		if opts.IncludeSubtree != nil {
			param.Add("include_subtree", "true")
		}
		if opts.GroupID != nil {
			param.Add("group_id", *opts.GroupID)
		}
		if opts.RoleID != nil {
			param.Add("role_id", *opts.RoleID)
		}
		if opts.DomainID != nil {
			param.Add("scope.domain.id", *opts.DomainID)
		}
		if opts.ScopeProjectID != nil {
			param.Add("scope.project.id", *opts.ScopeProjectID)
		}
		if opts.UserID != nil {
			param.Add("user_id", *opts.UserID)
		}
	}

	return param
}

func (opts *ListRoleAssignmentOpts) Valid() bool {
	return true
}
