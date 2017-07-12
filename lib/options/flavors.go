package options

import (
	"net/url"
	"strconv"
)

type ListFlavorsOpts struct {
	// Sorts by a flavor attribute, Optional
	SortKey *string `sort_key`

	// Sort direction. A valid value is asc (ascending) or desc (descending).
	// Default is asc, Optional
	SortDir *string `sort_dir`

	// Requests a page size of items, Optional
	Limit int `limit`

	// The ID of the last-seen item, Optional
	Maker *string `maker`

	// Filters the response by a minimum disk space, in GiB.
	// For example, 100. Optional
	MinDisk *string `min_disk`

	// Filters the response by a minimum RAM, in MB.
	// For example, 512.Optional
	MinRam *string `min_ram`

	// Filters the flavor list by only public flavors. Optional
	IsPublic bool `is_public`
}

func (opts *ListFlavorsOpts) ToQuery() (param url.Values) {
	param = url.Values{}

	if opts != nil {
		if opts.SortKey != nil {
			param.Add("sort_key", *opts.SortKey)
		}
		if opts.SortDir != nil {
			param.Add("sort_dir", *opts.SortDir)
		}

		if opts.Limit > 0 {
			param.Add("limit", strconv.Itoa(opts.Limit))
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
		if opts.IsPublic {
			param.Add("is_public", "true")
		} else if !opts.IsPublic {
			param.Add("is_public", "false")
		}

	}

	return param
}

type CreateFlavorOpts struct {
	// The display name of a flavor.
	Name *string `name`

	// The ID of the flavor, Optional.
	ID *string `id`

	// The amount of RAM a flavor has, MB.
	Ram *int `ram`

	// The size of the root disk that will be created, GB.
	Disk *int `disk`

	// The size of a dedicated swap disk that will be allocated, MB.
	// Optional
	Swap float64 `swap`

	// The receive / transmit factor, It defaults to 1.0.
	RxtxFactor float64 `json:"rxtx_factor"`
}

func (opts *CreateFlavorOpts) ToPayLoad() interface{} {
	type payload struct {
		Flavor *CreateFlavorOpts `json:"flavor"`
	}

	return payload{
		Flavor: opts,
	}
}
