package options

import (
	"net/url"
)

type ListPolicyOpts struct {
	Type *string `json:"type"`
}

func (opts *ListPolicyOpts) ToQuery() (options url.Values) {
	options = url.Values{}
	if opts == nil {
		return
	}

	if opts.Type != nil {
		options.Add("type", *opts.Type)
	}

	return
}

type CreatePolicyOpts struct {
	Type string `json:"type,omitempty"`
	Blob string `json:"blob,omitempty"`
}

func (opts *CreatePolicyOpts) ToPayload() interface{} {
	type payload struct {
		Policy *CreatePolicyOpts `json:"policy"`
	}

	return payload{
		Policy: opts,
	}
}

type UpdatePolicyOpts struct {
	Type string      `json:"type"`
	Blob interface{} `json:"blob"`
}

func (opts *UpdatePolicyOpts) ToPayload() interface{} {
	type payload struct {
		Policy *UpdatePolicyOpts `json:"policy"`
	}

	return payload{
		Policy: opts,
	}
}
