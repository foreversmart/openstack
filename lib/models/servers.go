package models

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/ports"
)

type ServersModel struct {
	ID                               string                                  `mapstructure:"id" json:"id"`
	Name                             string                                  `mapstructure:"name" json:"name"`
	ConfigDrive                      string                                  `mapstructure:"config_drive" json:"config_drive"`
	OSDcfDiskConfig                  string                                  `mapstructure:"OS-DCF:diskConfig" json:"OS-DCF:diskConfig"`
	OSExtAzAvailabilityZone          string                                  `mapstructure:"OS-EXT-AZ:availability_zone" json:"OS-EXT-AZ:availability_zone"`
	OSExtSrvAttrHost                 string                                  `mapstructure:"OS-EXT-SRV-ATTR:host" json:"OS-EXT-SRV-ATTR:host"`
	OSExtSrvAttrHypervisorHostname   string                                  `mapstructure:"OS-EXT-SRV-ATTR:hypervisor_hostname" json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	OSExtSrvAttrInstanceName         string                                  `mapstructure:"OS-EXT-SRV-ATTR:instance_name" json:"OS-EXT-SRV-ATTR:instance_name"`
	OSExtStsPowerState               int64                                   `mapstructure:"OS-EXT-STS:power_state" json:"OS-EXT-STS:power_state"`
	OSExtStsTaskState                string                                  `mapstructure:"OS-EXT-STS:task_state" json:"OS-EXT-STS:task_state"`
	OSExtStsVmState                  string                                  `mapstructure:"OS-EXT-STS:vm_state" json:"OS-EXT-STS:vm_state"`
	OSExtSrvAttrKernelID             string                                  `mapstructure:"S-EXT-SRV-ATTR:kernel_id" json:"S-EXT-SRV-ATTR:kernel_id"`
	OSExtSrvAttrLaunchIndex          int64                                   `mapstructure:"OS-EXT-SRV-ATTR:launch_index" json:"OS-EXT-SRV-ATTR:launch_index"`
	OSExtSrvAttrRamdiskID            string                                  `mapstructure:"OS-EXT-SRV-ATTR:ramdisk_id" json:"OS-EXT-SRV-ATTR:ramdisk_id"`
	OSExtSrvAttrRootDeviceName       string                                  `mapstructure:"OS-EXT-SRV-ATTR:root_device_name" json:"OS-EXT-SRV-ATTR:root_device_name"`
	OSExtSrvAttrReservationID        string                                  `mapstructure:"OS-EXT-SRV-ATTR:reservation_id" json:"OS-EXT-SRV-ATTR:reservation_id"`
	OSExtSrvAttrUserData             string                                  `mapstructure:"OS-EXT-SRV-ATTR:user_data" json:"OS-EXT-SRV-ATTR:user_data"`
	OSSrvUsgLaunchedAt               string                                  `mapstructure:"OS-SRV-USG:launched_at" json:"OS-SRV-USG:launched_at"`
	OSSrvUsgTerminatedAt             string                                  `mapstructure:"OS-SRV-USG:terminated_at" json:"OS-SRV-USG:terminated_at"`
	OsExtendedVolumesVolumesAttached []OsExtendedVolumesVolumesAttachedModel `mapstructure:"os-extended-volumes:volumes_attached" json:"os-extended-volumes:volumes_attached"`
	// TenantID identifies the tenant owning this server resource.
	TenantID string `mapstructure:"tenant_id" json:"tenant_id"`

	// UserID uniquely identifies the user account owning the tenant.
	UserID string `mapstructure:"user_id" json:"user_id"`
	// Updated and Created contain ISO-8601 timestamps of when the state of the server last changed, and when it was created.
	Updated string `mapstructure:"updated" json:"updated"`
	Created string `mapstructure:"created" json:"created"`

	HostID     string `mapstructure:"hostId" json:"hostId"`
	HostStatus string `mapstructure:"host_status" json:"host_status"`

	// Status contains the current operational status of the server, such as IN_PROGRESS or ACTIVE.
	Status string `mapstructure:"status" json:"status"`

	// Progress ranges from 0..100.
	// A request made against the server completes only once Progress reaches 100.
	Progress int `mapstructure:"progress" json:"progress"`

	// AccessIPv4 and AccessIPv6 contain the IP addresses of the server, suitable for remote access for administration.
	AccessIPv4 string `mapstructure:"accessIPv4" json:"accessIPv4"`
	AccessIPv6 string `mapstructure:"accessIPv6" json:"accessIPv6"`

	// Image refers to a JSON object, which itself indicates the OS image used to deploy the server.
	Image string `mapstructure:"image" json:"image"`

	// Flavor refers to a JSON object, which itself indicates the hardware configuration of the deployed server.
	Flavor map[string]interface{} `mapstructure:"flavor" json:"flavor"`

	// Addresses includes a list of all IP addresses assigned to the server, keyed by pool.
	Addresses map[string]interface{} `mapstructure:"addresses" json:"addresses"`

	// Metadata includes a list of all user-specified key-value pairs attached to the server.
	Metadata map[string]interface{} `mapstructure:"metadata" json:"metadata"`

	// Links includes HTTP references to the itself, useful for passing along to other APIs that might want a server reference.
	Links []interface{}

	// KeyName indicates which public key was injected into the server on launch.
	KeyName string `mapstructure:"key_name" json:"key_name"`

	// AdminPass will generally be empty ("").  However, it will contain the administrative password chosen when provisioning a new server without a set AdminPass setting in the first place.
	// Note that this is the ONLY time this field will be valid.
	AdminPass string `mapstructure:"adminPass" json:"adminPass"`

	// SecurityGroups includes the security groups that this instance has applied to it
	SecurityGroups []map[string]interface{} `mapstructure:"security_groups" json:"security_groups"`

	Locked bool `mapstructure:"locked" json:"locked"`
}

type OsExtendedVolumesVolumesAttachedModel struct {
	ID                  string `mapstructure:"id" json:"id"`
	DeleteOnTermination bool   `mapstructure:"delete_on_termination" json:"delete_on_termination"`
}

func ExtractServer(result gophercloud.Result) (serverInfo *ServersModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Server *ServersModel `mapstructure:"server" json:"server"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		serverInfo = response.Server
	}

	return
}

func ExtractServers(result gophercloud.Result) (serverInfo []*ServersModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Servers []*ServersModel `mapstructure:"servers" json:"servers"`
	}

	err = mapstructure.Decode(result.Body, &response)

	if err == nil {
		serverInfo = response.Servers
	}

	return
}

type AttachPortModel struct {
	// UUID for the port.
	PortId string `mapstructure:"port_id" json:"port_id"`

	// Network that this port is associated with.
	NetworkID string `mapstructure:"net_id" json:"net_id"`

	PortState string `mapstructure:"port_state" json:"port_state"`

	// Mac address to use on this port.
	MacAddr string `mapstructure:"mac_addr" json:"mac_addr"`

	// Specifies IP addresses for the port thus associating the port itself with
	FixedIPs []ports.IP `mapstructure:"fixed_ips" json:"fixed_ips"`
}

func ExtractAttachPorts(r gophercloud.Result) ([]*AttachPortModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		AttachPorts []*AttachPortModel `mapstructure:"interfaceAttachments"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return res.AttachPorts, err
}

type OpenVNCResult struct {
	Console struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"console"`
}

func ExtractOpenVNCResult(body interface{}) (result *OpenVNCResult, err error) {
	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &result)
	return
}

func (result *OpenVNCResult) URL() string {
	return result.Console.URL + "&path=v1.0/websocket/websockify"
}
