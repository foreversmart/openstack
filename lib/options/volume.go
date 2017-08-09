package options

import (
	"net/url"
	"strconv"
)

/**
 * used to list volumes by params
 */
type ListVolumeOpts struct {
	TenantID *string `json:"tenant_id"`
	Sort     *string `json:"sort"`
	Limit    *int    `json:"limit"`
	Marker   *string `json:"marker"`
}

func (opts *ListVolumeOpts) IsValid() bool {
	return true
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
 * used to show volume
 */
type ShowVolumeOpts struct {
	//The UUID of the tenant in a multi-tenancy cloud.Optional
	TenantID *string `json:"tenant_id"`
	VolumeID *string `json:"volume_id"`
}

func (opts *ShowVolumeOpts) IsValid() bool {
	return opts != nil && opts.VolumeID != nil
}

/**
 * used to create volume
 */
type CreateVolumeOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	VolumeType  *string `json:"volume_type,omitempty"`
	Size        *int    `json:"size,omitempty"`
	SnapshotID  *string `json:"snapshot_id,omitempty"` // create volume from a snapshot
	TenantID    *string `json:"tenant_id"`             //Optional
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
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	TenantID    *string `json:"tenant_id"` //Optional
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

/**
 * used to reset volume action
 */
type ResetVolumeOpts struct {
	Status          *string
	AttachStatus    *string
	MigrationStatus *string
}

func (opts *ResetVolumeOpts) IsValid() bool {
	return opts != nil &&
		opts.Status != nil &&
		opts.AttachStatus != nil &&
		opts.MigrationStatus != nil
}

func (opts *ResetVolumeOpts) ToPayload() interface{} {
	return map[string]interface{}{
		"os-reset_status": map[string]interface{}{
			"status":           opts.Status,
			"attach_status":    opts.AttachStatus,
			"migration_status": opts.MigrationStatus,
		},
	}
}
