package options

import "net/url"

type ListEndpointOpts struct {
	Interface *string `json:"interface"`  //optional
	ServiceId *string `json:"service_id"` //optional
}

func (opts *ListEndpointOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.Interface != nil {
			options.Add("interface", *opts.Interface)
		}

		if opts.ServiceId != nil {
			options.Add("service_id", *opts.ServiceId)
		}
	}

	return
}

type CreateEndpointOpts struct {
	RegionId  *string `json:"region_id,omitempty"` //optional
	Interface string  `json:"interface"`
	ServiceId string  `json:"service_id"`
	Url       string  `json:"url"`
	Enabled   *bool   `json:"enabled,omitempty"` //optional
}

func (opts *CreateEndpointOpts) ToPayload() interface{} {
	type payload struct {
		Project *CreateEndpointOpts `json:"endpoint"`
	}

	return payload{
		Project: opts,
	}
}

type UpdateEndpointOpts struct {
	RegionId  string `json:"region_id"`
	Interface string `json:"interface"`
	ServiceId string `json:"service_id"`
}

func (opts *UpdateEndpointOpts) ToPayload() interface{} {
	type payload struct {
		Project *UpdateEndpointOpts `json:"endpoint"`
	}

	return payload{
		Project: opts,
	}
}
