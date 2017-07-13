package options

import (
	"net/url"
	"strconv"
)

type ListOpts struct {
	Type   *string `json:"type"`
	UserID *string `json:"user_id"`
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListServersOpts struct {
	AccessIpV4       *string `json:"access_ip_v4"`
	AccessIpV6       *string `json:"access_ip_v6"`
	AutoDiskConfig   *string `json:"auto_disk_config "`
	AvailabilityZone *string `json:"availability_zone "`
	ConfigDrive      *bool   `json:"config_drive"`
	CreatedAt        *string `json:"created_at"`
	Deleted          *bool   `json:"deleted"`
	Description      *string `json:"description"`
	Ip               *string `json:"ip"`
	Ip6              *string `json:"ip6"`
	KernelID         *string `json:"kernel_id"`
	KeyName          *string `json:"key_name"`
	LaunchIndex      *string `json:"launch_index"`
	LaunchedAt       *string `json:"launched_at"`
	LockedBy         *string `json:"locked_by"`
	Node             *string `json:"node"`
	NotTags          *string `json:"not-tags"`
	NotTagsAny       *string `json:"not-tags-any"`
	PowerState       *string `json:"power_state"`
	Progress         *string `json:"progress"`
	ProjectId        *string `json:"project_id"`
	RamdiskID        *string `json:"ramdisk_id"`
	ReservationID    *string `json:"reservation_id "`
	RootDeviceName   *string `json:"root_device_name"`
	SortDir          *string `json:"sort_dir"`
	SortKey          *string `json:"sort_key"`
	TagsAny          *string `json:"tags-any"`
	TaskState        *string `json:"task_state"`
	TerminatedAt     *string `json:"terminated_at "`
	UserID           *string `json:"user_id"`
	Uuid             *string `json:"uuid"`
	VmState          *string `json:"vm_state"`
	Tags             *string `json:"tags"`

	// A time/date stamp for when the server last changed status.
	ChangesSince *string `json:"changes-since"`

	// Name of the image in URL format.
	Image *string `json:"image"`

	// Name of the flavor in URL format.
	Flavor *string `json:"flavor"`

	// Name of the server as a string; can be jsonueried with regular expressions.
	// Realize that ?name=bob returns both bob and bobb. If you need to match bob
	// only, you can use a regular expression matching the syntax of the
	// underlying database server implemented for Compute.
	Name *string `json:"name"`

	// Value of the status of the server so that you can filter on "ACTIVE" for example.
	Status *string `json:"status"`

	// Name of the host as a string.
	Host     *string `json:"host"`
	HostName *string `json:"hostname"`

	// UUID of the server at which you want to set a marker.
	Marker *string `json:"marker"`

	// Integer value for the limit of values to return.
	Limit *int `json:"limit"`

	// Bool to show all tenants
	AllTenants *bool `json:"all_tenants"`
}

func (opts *ListServersOpts) IsValid() bool {
	return true
}

func (opts *ListServersOpts) ToQuery() (options url.Values) {
	options = url.Values{}

	if opts != nil {
		if opts.UserID != nil {
			options.Add("user_id", *opts.UserID)
		}
		if opts.AccessIpV4 != nil {
			options.Add("access_ip_v4", *opts.AccessIpV4)
		}
		if opts.AccessIpV6 != nil {
			options.Add("access_ip_v6", *opts.AccessIpV6)
		}
		if opts.AutoDiskConfig != nil {
			options.Add("auto_disk_config", *opts.AutoDiskConfig)
		}
		if opts.AvailabilityZone != nil {
			options.Add("availability_zone", *opts.AvailabilityZone)
		}
		if opts.ConfigDrive != nil && *opts.ConfigDrive == true {
			options.Add("config_drive", "true")
		}
		if opts.CreatedAt != nil {
			options.Add("created_at", *opts.CreatedAt)
		}
		if opts.Deleted != nil && *opts.Deleted == true {
			options.Add("deleted", "true")
		}
		if opts.Description != nil {
			options.Add("description", *opts.Description)
		}
		if opts.Ip != nil {
			options.Add("ip", *opts.Ip)
		}
		if opts.Ip6 != nil {
			options.Add("ip6", *opts.Ip6)
		}
		if opts.KernelID != nil {
			options.Add("kernel_id", *opts.KernelID)
		}
		if opts.KeyName != nil {
			options.Add("key_name", *opts.KeyName)
		}
		if opts.LaunchIndex != nil {
			options.Add("launch_index", *opts.LaunchIndex)
		}
		if opts.LaunchedAt != nil {
			options.Add("launched_at", *opts.LaunchedAt)
		}
		if opts.LockedBy != nil {
			options.Add("locked_by", *opts.LockedBy)
		}
		if opts.Node != nil {
			options.Add("node", *opts.Node)
		}
		if opts.NotTags != nil {
			options.Add("not-tags", *opts.NotTags)
		}
		if opts.NotTagsAny != nil {
			options.Add("not-tags-any", *opts.NotTagsAny)
		}
		if opts.PowerState != nil {
			options.Add("power_state", *opts.PowerState)
		}
		if opts.Progress != nil {
			options.Add("progress", *opts.Progress)
		}
		if opts.ProjectId != nil {
			options.Add("project_id", *opts.ProjectId)
		}
		if opts.RamdiskID != nil {
			options.Add("ramdisk_id", *opts.RamdiskID)
		}
		if opts.ReservationID != nil {
			options.Add("reservation_id", *opts.ReservationID)
		}
		if opts.RootDeviceName != nil {
			options.Add("root_device_name", *opts.RootDeviceName)
		}
		if opts.SortDir != nil {
			options.Add("sort_dir", *opts.SortDir)
		}
		if opts.SortKey != nil {
			options.Add("sort_key", *opts.SortKey)
		}
		if opts.TagsAny != nil {
			options.Add("tags-any", *opts.TagsAny)
		}
		if opts.TaskState != nil {
			options.Add("task_state", *opts.TaskState)
		}
		if opts.TerminatedAt != nil {
			options.Add("terminated_at", *opts.TerminatedAt)
		}
		if opts.UserID != nil {
			options.Add("user_id", *opts.UserID)
		}
		if opts.Uuid != nil {
			options.Add("uuid", *opts.Uuid)
		}
		if opts.VmState != nil {
			options.Add("vm_state", *opts.VmState)
		}
		if opts.Tags != nil {
			options.Add("tags", *opts.Tags)
		}
		if opts.ChangesSince != nil {
			options.Add("changes-since", *opts.ChangesSince)
		}
		if opts.Image != nil {
			options.Add("image", *opts.Image)
		}
		if opts.Flavor != nil {
			options.Add("flavor", *opts.Flavor)
		}
		if opts.Name != nil {
			options.Add("name", *opts.Name)
		}
		if opts.Status != nil {
			options.Add("status", *opts.Status)
		}
		if opts.Host != nil {
			options.Add("host", *opts.Host)
		}
		if opts.HostName != nil {
			options.Add("hostname", *opts.HostName)
		}
		if opts.Marker != nil {
			options.Add("marker", *opts.Marker)
		}
		if opts.Limit != nil {
			options.Add("limit", strconv.Itoa(*opts.Limit))
		}
		if opts.AllTenants != nil && *opts.AllTenants == true {
			options.Add("all_tenants", "true")
		}
	}

	return
}

type CreateServersOpts struct {
}

func (opts *CreateServersOpts) ToPayload() interface{} {
	type payload struct {
		Servers *CreateServersOpts `json:"servers"`
	}

	return payload{
		Servers: opts,
	}
}

type UpdateServersOpts struct {
}

func (opts *UpdateServersOpts) ToPayload() interface{} {
	type payload struct {
		Servers *UpdateServersOpts `json:"servers"`
	}

	return payload{
		Servers: opts,
	}
}
