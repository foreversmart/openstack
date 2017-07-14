package options

import (
	"net/url"
	"strconv"
)

type ListFlavorsOpts struct {
	// Sorts by a flavor attribute, Optional
	SortKey *string `json:"sort_key,omitempty"`

	// Sort direction. A valid value is asc (ascending) or desc (descending).
	// Default is asc, Optional
	SortDir *string `json:"sort_dir,omitempty"`

	// Requests a page size of items, Optional
	Limit *int `json:"limit,omitempty"`

	// The ID of the last-seen item, Optional
	Maker *string `json:"maker,omitempty"`

	// Filters the response by a minimum disk space, in GiB.
	// For example, 100. Optional
	MinDisk *string `json:"min_disk,omitempty"`

	// Filters the response by a minimum RAM, in MB.
	// For example, 512.Optional
	MinRam *string `json:"min_ram,omitempty"`

	// Filters the flavor list by only public flavors. Optional
	IsPublic *bool `json:"is_public,omitempty"`
}

func (opts *ListFlavorsOpts) IsValid() bool {
	return true
}

func (opts *ListFlavorsOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts == nil {
		return
	}

	if opts.SortKey != nil {
		param.Add("sort_key", *opts.SortKey)
	}
	if opts.SortDir != nil {
		param.Add("sort_dir", *opts.SortDir)
	}

	if opts.Limit != nil {
		param.Add("limit", strconv.Itoa(*opts.Limit))
	}

	if opts.Maker != nil {
		param.Add("maker", *opts.Maker)
	}

	if opts.MinDisk != nil {
		param.Add("min_disk", *opts.MinDisk)
	}
	if opts.MinRam != nil {
		param.Add("min_ram", *opts.MinRam)
	}

	if opts.IsPublic != nil {
		if *opts.IsPublic {
			param.Add("is_public", "true")
		} else if !*opts.IsPublic {
			param.Add("is_public", "false")
		}
	}
	return param
}

type CreateFlavorOpts struct {
	// The display name of a flavor.
	Name *string `json:"name,omitempty"`

	// The ID of the flavor, Optional.
	ID *string `json:"id,omitempty"`

	// The amount of RAM a flavor has, MB.
	Ram *int `json:"ram,omitempty"`

	// The size of the root disk that will be created, GB.
	Disk *int `json:"disk,omitempty"`

	// The size of a dedicated swap disk that will be allocated, MB.
	// Optional
	Swap *float64 `json:"swap,omitempty"`

	//The number of virtual CPUs that will be allocated to the server.
	Vcpus *int `json:"vcpus,omitempty"`
}

func (opts *CreateFlavorOpts) IsValid() bool {
	return opts != nil && opts.Name != nil && opts.Ram != nil && *opts.Ram > 0 && opts.Disk != nil && *opts.Disk > 0 && opts.Vcpus != nil && *opts.Vcpus > 0
}

func (opts *CreateFlavorOpts) ToPayLoad() interface{} {
	type payload struct {
		Flavor *CreateFlavorOpts `json:"flavor"`
	}

	return payload{
		Flavor: opts,
	}
}
