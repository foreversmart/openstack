package options

import (
	"net/url"

	"github.com/kirk-enterprise/openstack-golang-sdk/lib/enums"
)

type ListEndpointOpts struct {
	Interface *enums.EndpointInterface `json:"interface"`  //optional
	ServiceId *string                  `json:"service_id"` //optional
}

func (opts *ListEndpointOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.Interface != nil {
			options.Add("interface", string(*opts.Interface))
		}

		if opts.ServiceId != nil {
			options.Add("service_id", *opts.ServiceId)
		}
	}

	return
}

type CreateEndpointOpts struct {
	Url       string                  `json:"url"`
	Interface enums.EndpointInterface `json:"interface"`
	Enabled   *bool                   `json:"enabled,omitempty"`   //optional
	RegionID  *string                 `json:"region_id,omitempty"` //optional
	ServiceID string                  `json:"service_id"`
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
	Url       *string                  `json:"url,omitempty"`
	Interface *enums.EndpointInterface `json:"interface,omitempty"`
	Enabled   *bool                    `json:"enabled,omitempty"`
	RegionID  *string                  `json:"region_id,omitempty"`
	ServiceID *string                  `json:"service_id,omitempty"`
}

func (opts *UpdateEndpointOpts) ToPayload() interface{} {
	type payload struct {
		Project *UpdateEndpointOpts `json:"endpoint"`
	}

	return payload{
		Project: opts,
	}
}
