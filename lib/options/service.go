package options

import "net/url"

type ListServiceOpts struct {
	Type *string `json:"type"` //optional
}

func (opts *ListServiceOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts != nil {
		if opts.Type != nil {
			param.Add("type", *opts.Type)
		}
	}

	return param
}

type CreateServiceOpts struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Desc    string `json:"description,omitempty"`
	Enabled bool   `json:"enabled"`
}

func (opts *CreateServiceOpts) ToPayload() interface{} {
	type payload struct {
		Service *CreateServiceOpts `json:"service"`
	}

	return payload{
		Service: opts,
	}
}

type UpdateServiceOpts struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Desc    *string `json:"description,omitempty"` //optional
	Enabled *bool   `json:"enabled,omitempty"`     //optional
}

func (opts *UpdateServiceOpts) ToPayload() interface{} {
	type payload struct {
		Service *UpdateServiceOpts `json:"service"`
	}

	return payload{
		Service: opts,
	}
}
