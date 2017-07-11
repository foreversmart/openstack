package options

type CreateSnapshotOpts struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Force       *bool   `json:"force"`
	VolumeID    *string `json:"volume_id"`
}

func (opts *CreateSnapshotOpts) IsValid() bool {
	return opts.VolumeID != nil && opts.Name != nil
}

type ListSnapshotOpts struct {
	ProjectId *string `json:"tenant_id"`
}

type UpdateSnapshotOpts struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (opts *UpdateSnapshotOpts) ToPayload() (payload interface{}) {
	type request struct {
		Snapshot *UpdateSnapshotOpts `json:"snapshot"`
	}

	return request{
		Snapshot: opts,
	}
}
