package options

import "net/url"

type CreateSnapshotOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Force       *bool   `json:"force,omitempty"`
	VolumeID    *string `json:"volume_id,omitempty"`
}

func (opts *CreateSnapshotOpts) IsValid() bool {
	return opts != nil && opts.VolumeID != nil && opts.Name != nil
}

func (opts *CreateSnapshotOpts) ToPayload() interface{} {
	type request struct {
		Snapshot *CreateSnapshotOpts `json:"snapshot"`
	}

	return request{
		Snapshot: opts,
	}
}

type UpdateSnapshotOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	TenantID    *string `json:"tenant_id"` //Optional
	AllTenants  *string `json:"all_tenants"`
}

func (opts *UpdateSnapshotOpts) IsValid() bool {
	return opts != nil
}

func (opts *UpdateSnapshotOpts) ToPayload() interface{} {
	type request struct {
		Snapshot *UpdateSnapshotOpts `json:"snapshot"`
	}

	return request{
		Snapshot: opts,
	}
}

type ListSnapshotOpts struct {
	TenantID   *string `json:"project_id,omitempty"`
	AllTenants *string `json:"all_tenants,omitempty"`
}

func (opts *ListSnapshotOpts) IsValid() bool {
	return true
}

func (opts *ListSnapshotOpts) ToQuery() url.Values {
	options := url.Values{}

	if opts != nil {
		if opts.AllTenants != nil {
			options.Add("all_tenants", *opts.AllTenants)
		}

		if opts.TenantID != nil {
			options.Add("project_id", *opts.TenantID)
		}

	}

	return options
}
