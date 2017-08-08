package options

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strconv"
)

// ListServersOpts allows the filtering and sorting of paginated collections through
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
	ProjectID        *string `json:"project_id"`
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
	AllTenants *int `json:"all_tenants"`
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
		if opts.ProjectID != nil {
			options.Add("project_id", *opts.ProjectID)
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
		if opts.AllTenants != nil {
			options.Add("all_tenants", strconv.Itoa(*opts.AllTenants))
		}
	}

	return
}

// Network is used within CreateOpts to control a new server's network attachments.
type ServerNetworkOpts struct {
	// UUID of a nova-network to attach to the newly provisioned server.
	// Required unless Port is provided.
	UUID string `json:"uuid"`

	// Port of a neutron network to attach to the newly provisioned server.
	// Required unless UUID is provided.
	Port *string `json:"port,omitempty"`

	// FixedIP [optional] specifies a fixed IPv4 address to be used on this network.
	FixedIP *string `json:"fixed_ip,omitempty"`
}

// Personality is an array of files that are injected into the server at launch.
type ServerPersonalityOpts []*ServerFileOpts

// File is used within CreateOpts and RebuildServerOpts to inject a file into the server at launch.
// File implements the json.Marshaler interface, so when a Create or Rebuild operation is requested,
// json.Marshal will call File's MarshalJSON method.
type ServerFileOpts struct {
	// Path of the file
	Path string
	// Contents of the file. Maximum content size is 255 bytes.
	Contents []byte
}

// MarshalJSON marshals the escaped file, base64 encoding the contents.
func (f *ServerFileOpts) MarshalJSON() ([]byte, error) {
	file := struct {
		Path     string `json:"path"`
		Contents string `json:"contents"`
	}{
		Path:     f.Path,
		Contents: base64.StdEncoding.EncodeToString(f.Contents),
	}
	return json.Marshal(file)
}

// CreateOpts specifies server creation parameters.
type CreateServerOpts struct {
	// Name [required] is the name to assign to the newly launched server.
	Name *string `json:"name"`

	// ImageRef [optional; required if ImageName is not provided] is the ID or full
	// URL to the image that contains the server's OS and initial state.
	// Also optional if using the boot-from-volume extension.
	ImageRef *string `json:"imageRef,omitempty"`

	// FlavorRef [optional; required if FlavorName is not provided] is the ID or
	// full URL to the flavor that describes the server's specs.
	FlavorRef *string `json:"flavorRef,omitempty"`

	// SecurityGroups [optional] lists the names of the security groups to which this server should belong.
	SecurityGroups []string `json:"security_groups,omitempty"`

	// UserData [optional] contains configuration information or scripts to use upon launch.
	// Create will base64-encode it for you.
	UserData []byte `json:"user_data,omitempty"`

	// AvailabilityZone [optional] in which to launch the server.
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// Networks [optional] dictates how this server will be attached to available networks.
	// By default, the server will be attached to all isolated networks for the tenant.
	Networks []*ServerNetworkOpts `json:"networks,omitempty"`

	// Metadata [optional] contains key-value pairs (up to 255 bytes each) to attach to the server.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Personality [optional] includes files to inject into the server at launch.
	// Create will base64-encode file contents for you.
	Personality *ServerPersonalityOpts `json:"personality,omitempty"`

	// ConfigDrive [optional] enables metadata injection through a configuration drive.
	ConfigDrive *bool `json:"config_drive,omitempty"`

	// KeyName [optional] Key pair name.
	KeyName *bool `json:"key_name,omitempty"`

	// AdminPass [optional] sets the root user password. If not set, a randomly-generated
	// password will be created and returned in the response.
	AdminPass *string `json:"adminPass,omitempty"`

	// AccessIPv4 [optional ] specifies an IPv4 address for the instance.
	AccessIPv4 *string `json:"accessIPv4,omitempty"`

	// AccessIPv6 [optional] specifies an IPv6 address for the instance.
	AccessIPv6 *string `json:"accessIPv6,omitempty"`

	//Controls how the API partitions the disk when you create, rebuild, or resize servers.
	OSDcfDiskConfig *string `json:"OS-DCF:diskConfig,omitempty"`
}

func (opts *CreateServerOpts) IsValid() bool {
	return opts != nil && opts.Name != nil && (opts.FlavorRef != nil || opts.ImageRef != nil) &&
		(opts.OSDcfDiskConfig == nil || (*opts.OSDcfDiskConfig == "AUTO" || *opts.OSDcfDiskConfig == "MANUAL"))
}

func (opts *CreateServerOpts) ToPayload() interface{} {
	type payload struct {
		Server *CreateServerOpts `json:"server"`
	}

	return payload{
		Server: opts,
	}
}

type UpdateServersOpts struct {
	// Name [optional] changes the displayed name of the server.
	// The server host name will *not* change.
	// Server names are not constrained to be unique, even within the same tenant.
	Name *string `json:"name"`

	// AccessIPv4 [optional] provides a new IPv4 address for the instance.
	AccessIPv4 *string `json:"accessIPv4"`

	// AccessIPv6 [optional] provides a new IPv6 address for the instance.
	AccessIPv6 *string `json:"accessIPv6"`

	//Controls how the API partitions the disk when you create, rebuild, or resize servers.
	OSDcfDiskConfig *string `json:"OS-DCF:diskConfig"`
}

func (opts *UpdateServersOpts) IsValid() bool {
	return opts != nil && opts.Name != nil && opts.AccessIPv4 != nil && opts.AccessIPv6 != nil && opts.OSDcfDiskConfig != nil
}

func (opts *UpdateServersOpts) ToPayload() interface{} {
	type payload struct {
		Server *UpdateServersOpts `json:"server"`
	}

	return payload{
		Server: opts,
	}
}

// RebuildServerOpts represents the configuration options used in a server rebuild
// operation
type RebuildServerOpts struct {
	// Required. The ID of the image you want your server to be provisioned on
	ImageID string

	// Name to set the server to
	Name string

	// Metadata [optional] contains key-value pairs (up to 255 bytes each) to attach to the server.
	Metadata map[string]string
}

func (opts *RebuildServerOpts) IsValid() bool {
	if opts.ImageID == "" {
		return false
	}

	return true
}

// ToServerRebuildMap formats a RebuildServerOpts struct into a map for use in JSON
func (opts RebuildServerOpts) ToServerRebuildMap() (map[string]interface{}, error) {
	var err error
	server := make(map[string]interface{})

	if err != nil {
		return server, err
	}

	server["imageRef"] = opts.ImageID

	if opts.Name != "" {
		server["name"] = opts.Name
	}

	if opts.Metadata != nil {
		server["metadata"] = opts.Metadata
	}

	return map[string]interface{}{"rebuild": server}, nil
}
