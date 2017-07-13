package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type PortModel struct {
	// UUID for the port.
	ID string `mapstructure:"id" json:"id"`
	// Network that this port is associated with.
	NetworkID string `mapstructure:"network_id" json:"network_id"`
	// Human-readable name for the port. Might not be unique.
	Name string `mapstructure:"name" json:"name"`
	// Administrative state of port. If false (down), port does not forward packets.
	AdminStateUp bool `mapstructure:"admin_state_up" json:"admin_state_up"`
	// Indicates whether network is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional values.
	Status string `mapstructure:"status" json:"status"`
	// Mac address to use on this port.
	MACAddress string `mapstructure:"mac_address" json:"mac_address"`
	// Specifies IP addresses for the port thus associating the port itself with
	// the subnets where the IP addresses are picked from
	FixedIPs []IP `mapstructure:"fixed_ips" json:"fixed_ips"`
	// Owner of network. Only admin users can specify a tenant_id other than its own.
	TenantID  string `mapstructure:"tenant_id" json:"tenant_id"`
	ProjectID string `mapstructure:"project_id" json:"project_id"`
	// Identifies the entity (e.g.: dhcp agent) using this port.
	DeviceOwner string `mapstructure:"device_owner" json:"device_owner"`
	// Specifies the IDs of any security groups associated with a port.
	SecurityGroups []string `mapstructure:"security_groups" json:"security_groups"`
	// Identifies the device (e.g., virtual server) using this port.
	DeviceID string `mapstructure:"device_id" json:"device_id"`
	// Identifies the list of IP addresses the port will recognize/accept
	AllowedAddressPairs []AddressPair `mapstructure:"allowed_address_pairs" json:"allowed_address_pairs"`

	CreatedAt string `mapstructure:"created_at" json:"created_at"`
}

type IP struct {
	SubnetID  string `mapstructure:"subnet_id" json:"subnet_id"`
	IPAddress string `mapstructure:"ip_address" json:"ip_address,omitempty"`
}

type AddressPair struct {
	IPAddress  string `mapstructure:"ip_address" json:"ip_address,omitempty"`
	MACAddress string `mapstructure:"mac_address" json:"mac_address,omitempty"`
}

func ExtractPortsByBody(body interface{}) (ports []*PortModel, err error) {
	var resp struct {
		Ports []*PortModel `mapstructure:"ports"`
	}

	err = mapstructure.Decode(body, &resp)
	if err == nil {
		ports = resp.Ports
	}
	return
}

func ExtractPort(r gophercloud.Result) (port *PortModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var resp struct {
		Port *PortModel `mapstructure:"port"`
	}

	err = mapstructure.Decode(r.Body, &resp)
	if err == nil {
		port = resp.Port
	}
	return resp.Port, err
}

func ExtractPorts(r gophercloud.Result) (port []*PortModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var resp struct {
		Ports []*PortModel `mapstructure:"ports"`
	}

	err = mapstructure.Decode(r.Body, &resp)
	if err == nil {
		port = resp.Ports
	}
	return resp.Ports, err
}
