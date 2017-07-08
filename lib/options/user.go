package options

import "net/url"

// Identity API V3
type ListUserOpts struct {
	DomainID          *string `json:"domain_id"`
	Name              *string `json:"name"`
	PasswordExpiresAt *string `json:"password_expires_at"`
	Enabled           *string `json:"enabled"`
	IdpID             *string `json:"idp_id"`
	UniqueID          *string `json:"unique_id"`
	ProtocolID        *string `json:"protocol_id"`
}

func (opts *ListUserOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.DomainID != nil {
			options.Add("domain_id", *opts.DomainID)
		}

		if opts.Enabled != nil {
			options.Add("enabled", *opts.Enabled)
		}

		if opts.IdpID != nil {
			options.Add("idp_id", *opts.IdpID)
		}

		if opts.Name != nil {
			options.Add("name", *opts.Name)
		}

		if opts.PasswordExpiresAt != nil {
			options.Add("password_expires_at", *opts.PasswordExpiresAt)
		}

		if opts.ProtocolID != nil {
			options.Add("protocol_id", *opts.ProtocolID)
		}

		if opts.UniqueID != nil {
			options.Add("unique_id", *opts.UniqueID)
		}
	}

	return
}

type CreateUserOpts struct {
	DomainID         *string `json:"domain_id"`
	Name             *string `json:"name"`
	Password         *string `json:"password"`
	Enabled          *string `json:"enabled"`
	DefaultProjectID *string `json:"default_project_id"`
}

func (opts *CreateUserOpts) ToPayload() interface{} {
	type payload struct {
		User *CreateUserOpts `json:"user"`
	}

	return payload{
		User: opts,
	}
}

type UpdateUserOpts struct {
	DomainID         *string `json:"domain_id"`
	Name             *string `json:"name"`
	Password         *string `json:"password"`
	Enabled          *string `json:"enabled"`
	DefaultProjectID *string `json:"default_project_id"`
}

func (opts *UpdateUserOpts) ToPayload() interface{} {
	type payload struct {
		User *UpdateUserOpts `json:"user"`
	}

	return payload{
		User: opts,
	}
}

type ChangeUserPasswordOpts struct {
	OriginalPassword *string `json:"original_password"`
	Password         *string `json:"password"`
}

func (opts *ChangeUserPasswordOpts) ToPayload() interface{} {
	type payload struct {
		User *ChangeUserPasswordOpts `json:"user"`
	}

	return payload{
		User: opts,
	}
}
