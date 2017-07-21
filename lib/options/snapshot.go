package options

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
