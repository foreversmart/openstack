package options

import (
	"net/url"
)

type ListGroupOpts struct {
	DomainID *string `json:"domain_id"`
	Name     *string `json:"name"`
}

func (opts ListGroupOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts.Name != nil {
		options.Add("name", *opts.Name)
	}

	if opts.DomainID != nil {
		options.Add("domain_id", *opts.DomainID)
	}

	return
}

type CreateGroupOpts struct {
	DomainID    string `json:"domain_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (opts *CreateGroupOpts) ToPayload() interface{} {
	type payload struct {
		Group *CreateGroupOpts `json:"group"`
	}

	return payload{
		Group: opts,
	}
}

type UpdateGroupOpts struct {
	DomainID    *string `json:"domain_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (opts *UpdateGroupOpts) ToPayload() interface{} {
	type payload struct {
		Group *UpdateGroupOpts `json:"group"`
	}

	return payload{
		Group: opts,
	}
}
