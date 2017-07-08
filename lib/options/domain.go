package options

import (
	"net/url"
)

/*
	name: Filters the response by a domain name.

	enabled: If set to true, then only domains that are enabled will be returned,
if set to false only that are disabled will be returned.
Any value other than 0, including no value, will be interpreted as true.
*/
type ListDomainOpts struct {
	Name    *string `json:"name"`
	Enabled *string `json:"enabled"`
}

func (opts ListDomainOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts.Name != nil {
		options.Add("name", *opts.Name)
	}

	if opts.Enabled != nil {
		options.Add("enabled", *opts.Enabled)
	}

	return
}

/*
	enabled: If set to true, domain is created enabled.
If set to false, domain is created disabled.
The default is true.

	Users can only authorize against an enabled domain (and any of its projects).
In addition, users can only authenticate if the domain that owns them is also enabled.
Disabling a domain prevents both of these things.
*/
type CreateDomainOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

func (opts *CreateDomainOpts) ToPayload() interface{} {
	type payload struct {
		Domain *CreateDomainOpts `json:"domain"`
	}

	return payload{
		Domain: opts,
	}
}

type UpdateDomainOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
}

func (opts *UpdateDomainOpts) ToPayload() interface{} {
	type payload struct {
		Domain *UpdateDomainOpts `json:"domain"`
	}

	return payload{
		Domain: opts,
	}
}
