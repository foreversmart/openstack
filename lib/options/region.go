package options

import (
	"net/url"
)

type CreateRegionOpts struct {
	//The ID for the region, optional
	Id *string `json:"id,omitempty"`

	//The region description, optional
	Description *string `json:"description,omitempty"`

	//To make this region a child of another region,
	//set this parameter to the ID of the parent region, optional
	ParentRegionId *string `json:"parent_region_id,omitempty"`
}

type UpdateRegionOpts struct {
	//The region description, optional
	Description *string `json:"description,omitempty"`

	//To make this region a child of another region,
	//set this parameter to the ID of the parent region, optional
	ParentRegionId *string `json:"parent_region_id,omitempty"`
}

type ListRegionOpts struct {
	ParentRegionId *string `json:parent_region_id` //optional
}

func (opts *ListRegionOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts != nil {
		if opts.ParentRegionId != nil {
			param.Add("parent_region_id", *opts.ParentRegionId)
		}
	}

	return param
}

func (opts *CreateRegionOpts) ToPayLoad() interface{} {
	type payload struct {
		Region *CreateRegionOpts `json:"region"`
	}

	return payload{
		Region: opts,
	}
}

func (opts *UpdateRegionOpts) ToPayLoad() interface{} {
	type payload struct {
		Region *UpdateRegionOpts `json:"region"`
	}

	return payload{
		Region: opts,
	}
}
