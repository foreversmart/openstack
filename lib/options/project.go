package options

import (
	"net/url"
)

type ListProjectOpts struct {
	DomainID *string `json:"domain_id"`
	Name     *string `json:"name"`
	IsDomain *bool   `json:"is_domain"`
	Enabled  *bool   `json:"enabled"`
	ParentID *string `json:"parent_id"`
}

func (opts *ListProjectOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.DomainID != nil {
			options.Add("domain_id", *opts.DomainID)
		}

		if opts.Name != nil {
			options.Add("name", *opts.Name)
		}

		if opts.ParentID != nil {
			options.Add("parent_id", *opts.ParentID)
		}
	}

	return
}

type CreateProjectOpts struct {
	DomainID    *string `json:"domain_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsDomain    *bool   `json:"is_domain,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
}

func (opts *CreateProjectOpts) ToPayload() interface{} {
	type payload struct {
		Project *CreateProjectOpts `json:"project"`
	}

	return payload{
		Project: opts,
	}
}

type UpdateProjectOpts struct {
	Name        *string `json:"name,omitempty"`
	DomainID    *string `json:"domain_id,omitempty"`
	Description *string `json:"description,omitempty"`
	IsDomain    *bool   `json:"is_domain,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
}

func (opts *UpdateProjectOpts) ToPayload() interface{} {
	type payload struct {
		Project *UpdateProjectOpts `json:"project"`
	}

	return payload{
		Project: opts,
	}
}
