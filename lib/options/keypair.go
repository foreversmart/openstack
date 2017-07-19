package options

import "net/url"

type ListKeypairOpt struct {
	UserID *string `json:"user_id"`
	Limit  *string `json:"limit"`
	Marker *string `json:"marker"`
}

func (opts *ListKeypairOpt) IsValid() bool {
	return true
}

func (opts *ListKeypairOpt) ToQuery() (options url.Values) {
	options = url.Values{}
	if opts == nil {
		return
	}

	if opts.UserID != nil {
		options.Add("user_id", *opts.UserID)
	}

	if opts.Limit != nil {
		options.Add("limit", *opts.Limit)
	}

	if opts.Marker != nil {
		options.Add("marker", *opts.Marker)
	}

	return
}

type CreateKeypairOpts struct {
	Name      *string `json:"name"`
	PublicKey *string `json:"public_key"`
	Type      *string `json:"type"`
	UserId    *string `json:"user_id"`
}

func (opts *CreateKeypairOpts) ToPayload() interface{} {
	type request struct {
		Keypair *CreateKeypairOpts `json:"keypair"`
	}

	return request{
		Keypair: opts,
	}
}

func (opts *CreateNetworkOpts) IsValid() bool {
	return opts != nil && opts.Name != nil
}
