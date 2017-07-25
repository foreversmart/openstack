package options

import (
	"net/url"
	"strconv"

	"github.com/qbox/openstack-golang-sdk/lib/enums"
	"github.com/rackspace/gophercloud/openstack/imageservice/v2/images"
)

/**
 * used to list imaegs by params
 */
type ListImagesOpts struct {
	Limit        *int                   `json:"limit"`
	Marker       *string                `json:"marker"`
	Name         *string                `json:"name"`
	Owner        *string                `json:"owner"`
	Status       *string                `json:"status"`
	Tag          *string                `json:"tag"`
	Visibility   *enums.ImageVisibility `json:"visibility"`
	MemberStatus *string                `json:"member_status"`
	SizeMax      *string                `json:"size_max"`
	SizeMin      *string                `json:"size_min"`
	CreatedAt    *string                `json:"created_at"`
	UpdatedAt    *string                `json:"updated_at"`
	SortDir      *string                `json:"sort_dir"`
	SortKey      *string                `json:"sort_key"`
	Sort         *string                `json:"sort"`
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
		options.Add("status", *opts.Status)
	}
	if opts.Tag != nil {
		options.Add("tag", *opts.Tag)
	}
	if opts.Visibility != nil {
		options.Add("visibility", string(*opts.Visibility))
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
 * used to create image
 */
type CreateImagesOpts struct {
	ID              *string                `json:"id,omitempty"`
	Name            *string                `json:"name,omitempty"`
	ContainerFormat *string                `json:"container_format,omitempty"`
	DiskFormat      *string                `json:"disk_format,omitempty"`
	MinDisk         *int                   `json:"min_disk,omitempty"`
	MinRam          *int                   `json:"min_ram,omitempty"`
	Protected       *bool                  `json:"protected,omitempty"`
	Tags            *[]string              `json:"tags,omitempty"`
	Visibility      *enums.ImageVisibility `json:"visibility,omitempty"`
}

func (opts *CreateImagesOpts) IsValid() bool {
	return opts != nil && opts.Name != nil
}

func (opts *CreateImagesOpts) ToPayload() interface{} {
	return opts
}

/**
 * used to update image
 */
type UpdateImagesOpts struct {
	Name *string   `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
}

func (opts *UpdateImagesOpts) IsValid() bool {
	return opts != nil
}

func (opts *UpdateImagesOpts) ToPatches() images.UpdateOpts {
	payload := make(images.UpdateOpts, 0, 2)
	if opts.Name != nil {
		payload = append(payload, images.ReplaceImageName{
			NewName: *opts.Name,
		})
	}

	if len(*opts.Tags) > 0 {
		payload = append(payload, images.ReplaceImageTags{
			NewTags: *opts.Tags,
		})
	}

	return payload
}
