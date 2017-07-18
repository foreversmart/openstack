package options

import (
	"net/url"
	"strconv"
)

/**
 * used to list imaegs by params
 */
type ListImagesOpts struct {
	Limit        *int    `json:"limit"`
	Marker       *string `json:"marker"`
	Name         *string `json:"name"`
	Owner        *string `json:"owner"`
	Status       *int    `json:"status"`
	Tag          *string `json:"tag"`
	Visibility   *string `json:"visibility"`
	MemberStatus *string `json:"member_status"`
	SizeMax      *string `json:"size_max"`
	SizeMin      *string `json:"size_min"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
	SortDir      *string `json:"sort_dir"`
	SortKey      *string `json:"sort_key"`
	Sort         *string `json:"sort"`
}

func (opts *ListImagesOpts) IsValid() bool {
	return true
}

func (opts *ListImagesOpts) ToQuery() (options url.Values) {
	if opts == nil {
		return
	}

	options = url.Values{}

	if opts.Limit != nil {
		options.Add("limit", strconv.Itoa(*opts.Limit))
	}
	if opts.Marker != nil {
		options.Add("marker", *opts.Marker)
	}
	if opts.Name != nil {
		options.Add("name", *opts.Name)
	}
	if opts.Owner != nil {
		options.Add("owner", *opts.Owner)
	}
	if opts.Status != nil {
		options.Add("status", strconv.Itoa(*opts.Status))
	}
	if opts.Tag != nil {
		options.Add("tag", *opts.Tag)
	}
	if opts.Visibility != nil {
		options.Add("visibility", *opts.Visibility)
	}
	if opts.MemberStatus != nil {
		options.Add("member_status", *opts.MemberStatus)
	}
	if opts.SizeMax != nil {
		options.Add("size_max", *opts.SizeMax)
	}
	if opts.SizeMin != nil {
		options.Add("size_min", *opts.SizeMin)
	}
	if opts.CreatedAt != nil {
		options.Add("created_at", *opts.CreatedAt)
	}
	if opts.UpdatedAt != nil {
		options.Add("updated_at", *opts.UpdatedAt)
	}
	if opts.SortDir != nil {
		options.Add("sort_dir", *opts.SortDir)
	}
	if opts.SortKey != nil {
		options.Add("sort_key", *opts.SortKey)
	}
	if opts.Sort != nil {
		options.Add("sort", *opts.Sort)
	}

	return
}

/**
 * used to create volume
 */
type CreateImagesOpts struct {
	ID              *string   `json:"id"`
	Name            *string   `json:"name"`
	ContainerFormat *string   `json:"container_format"`
	DiskFormat      *string   `json:"disk_format"`
	MinDisk         *int      `json:"min_disk"`
	MinRam          *int      `json:"min_ram"`
	Protected       *bool     `json:"protected"`
	Tags            *[]string `json:"tags"`
	Visibility      *string   `json:"visibility"`
}

func (opts *CreateImagesOpts) IsValid() bool {
	return opts != nil && opts.Name != nil
}

func (opts *CreateImagesOpts) ToPayload() interface{} {
	return opts
}
