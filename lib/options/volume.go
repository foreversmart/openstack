package options

import "net/url"

/**
 * used to list volumes by params
 */
type ListVolumeOpts struct {
	SortKey *string `json:"sort_key"`
	SortDir *string `json:"sort_dir"`
	Limit   *string `json:"limit"`
	Offset  *string `json:"offset"`
	Marker  *string `json:"marker"`
}

func (opts *ListVolumeOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.SortKey != nil {
			options.Add("sort_key", *opts.SortKey)
		}
		if opts.SortDir != nil {
			options.Add("sort_dir", *opts.SortDir)
		}
		if opts.Limit != nil {
			options.Add("limit", *opts.Limit)
		}
		if opts.Offset != nil {
			options.Add("offset", *opts.Offset)
		}
		if opts.Marker != nil {
			options.Add("marker", *opts.Marker)
		}
	}

	return
}

/**
 * used to create volume
 */
type CreateVolumeOpts struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	VolumeType  *string `json:"volume_type"`
	Size        *int    `json:"size"`
	SnapshotId  *string `json:"snapshot_id"` // create volume from a snapshot
}

func (opts *CreateVolumeOpts) IsValid() bool {
	return opts.Size != nil
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

func (opts *UpdateVolumeOpts) ToPayload() interface{} {
	type payload struct {
		Volume *UpdateVolumeOpts `json:"volume"`
	}

	return payload{
		opts,
	}
}

/**
 * used to build volume action request body
 */
func ToVolumeActionResizeMap(newSize *int) (payload map[string]interface{}) {
	payload = map[string]interface{}{
		"os-extend": map[string]interface{}{
			"new_size": newSize,
		},
	}
	return
}

func ToVolumeActionAttachMap(instanceId *string, mountPoint *string) (payload map[string]interface{}) {
	payload = map[string]interface{}{
		"os-attach": map[string]interface{}{
			"instance_uuid": instanceId,
			"mountpoint": mountPoint,
		},
	}
	return
}

func ToVolumeActionDetachMap(attachmentId *string) (payload map[string]interface{}) {
	payload = map[string]interface{}{
		"os-detach": map[string]interface{}{
			"attachment_id": attachmentId,
		},
	}
	return
}