package options

import "net/url"

type ListCredentialOpts struct {
	Type   *string `json:"type"`
	UserID *string `json:"user_id"`
}

func (opts *ListCredentialOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.Type != nil {
			options.Add("type", *opts.Type)
		}

		if opts.UserID != nil {
			options.Add("user_id", *opts.UserID)
		}
	}

	return
}

type CreateCredentialOpts struct {
	ProjectId string `json:"project_id"`
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
	Blob      string `json:"blob"`
}

func (opts *CreateCredentialOpts) ToPayload() interface{} {
	type payload struct {
		Credential *CreateCredentialOpts `json:"credential"`
	}

	return payload{
		Credential: opts,
	}
}

type UpdateCredentialOpts struct {
	ProjectId string  `json:"project_id"`
	Type      *string `json:"type,omitempty"`
	UserID    *string `json:"user_id,omitempty"`
	Blob      *string `json:"blob,omitempty"`
}

func (opts *UpdateCredentialOpts) ToPayload() interface{} {
	type payload struct {
		Credential *UpdateCredentialOpts `json:"credential"`
	}

	return payload{
		Credential: opts,
	}
}
