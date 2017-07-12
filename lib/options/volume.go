package options

import (
	"net/url"
	"strconv"
)

/**
 * used to list volumes by params
 */
type ListVolumeOpts struct {
	Sort   *string `json:"sort"`
	Limit  *int    `json:"limit"`
	Marker *string `json:"marker"`
}

func (opts *ListVolumeOpts) ToQuery() url.Values {
	options := url.Values{}

	if opts != nil {
		if opts.Sort != nil {
			options.Add("sort", *opts.Sort)
		}
		if opts.Limit != nil {
			options.Add("limit", strconv.Itoa(*opts.Limit))
		}
		if opts.Marker != nil {
			options.Add("marker", *opts.Marker)
		}
	}

	return options
}

/**
 * used to create volume
 */
type CreateVolumeOpts struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	VolumeType  *string `json:"volume_type"`
	Size        *int    `json:"size"`
	SnapshotID  *string `json:"snapshot_id"` // create volume from a snapshot
}

func (opts *CreateVolumeOpts) IsValid() bool {
	return opts != nil && opts.Size != nil && *opts.Size > 0
}

func (opts *CreateVolumeOpts) ToPayload() interface{} {
	type payload struct {
		Volume *CreateVolumeOpts `json:"volume"`
	}

	return payload{
		opts,
	}
}

/**
 * used to update volume name & description info
 */
type UpdateVolumeOpts struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (opts *UpdateVolumeOpts) IsValid() bool {
	return opts != nil
}

func (opts *UpdateVolumeOpts) ToPayload() interface{} {
	type payload struct {
		Volume *UpdateVolumeOpts `json:"volume"`
	}

	return payload{
		opts,
	}
}

/**
 * used to resize volume action
 */
type ResizeVolumeOpts struct {
	Size *int
}

func (opts *ResizeVolumeOpts) IsValid() bool {
	return opts != nil && opts.Size != nil && *opts.Size > 0
}

func (opts *ResizeVolumeOpts) ToPayload() interface{} {
	return map[string]interface{}{
		"os-extend": map[string]interface{}{
			"new_size": opts.Size,
		},
	}
}
