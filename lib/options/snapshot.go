package options

type CreateSnapshotOpts struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Force       *bool   `json:"force"`
	VolumeID    *string `json:"volume_id"`
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
	Name        *string `json:"name"`
	Description *string `json:"description"`
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
