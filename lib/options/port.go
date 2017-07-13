package options

import (
	"net/url"
	"strconv"

	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
)

type ListPortOpts struct {
	Status       *string `json:"status"`
	Name         *string `json:"name"`
	AdminStateUp *bool   `json:"admin_state_up"`
	NetworkID    *string `json:"network_id"`
	TenantID     *string `json:"tenant_id"`
	DeviceOwner  *string `json:"device_owner"`
	MACAddress   *string `json:"mac_address"`
	ID           *string `json:"id"`
	DeviceID     *string `json:"device_id"`
	Limit        *int    `json:"limit"`
	Marker       *string `json:"marker"`
	SortKey      *string `json:"sort_key"`
	SortDir      *string `json:"sort_dir"`
}

func (opts *ListPortOpts) IsValid() bool {
	return true
}

func (opts *ListPortOpts) ToQuery() (options url.Values) {
	options = url.Values{}
	if opts == nil {
		return
	}

	if opts.Status != nil {
		options.Add("status", *opts.Status)
	}
	if opts.Name != nil {
		options.Add("name", *opts.Name)
	}
	if opts.AdminStateUp != nil {
		options.Add("admin_state_up", strconv.FormatBool(*opts.AdminStateUp))
	}
	if opts.NetworkID != nil {
		options.Add("network_id", *opts.NetworkID)
	}
	if opts.TenantID != nil {
		options.Add("tenant_id", *opts.TenantID)
	}
	if opts.DeviceOwner != nil {
		options.Add("device_owner", *opts.DeviceOwner)
	}
	if opts.MACAddress != nil {
		options.Add("mac_address", *opts.MACAddress)
	}
	if opts.ID != nil {
		options.Add("id", *opts.ID)
	}
	if opts.DeviceID != nil {
		options.Add("device_id", *opts.DeviceID)
	}
	if opts.Limit != nil {
		options.Add("limit", strconv.Itoa(*opts.Limit))
	}
	if opts.Marker != nil {
		options.Add("marker", *opts.Marker)
	}
	if opts.SortKey != nil {
		options.Add("sort_key", *opts.SortKey)
	}
	if opts.SortDir != nil {
		options.Add("sort_dir", *opts.SortDir)
	}

	return
}

type CreatePortOpts struct {
	NetworkID           *string              `json:"network_id"`
	Name                *string              `json:"name"`
	AdminStateUp        *bool                `json:"admin_state_up"`
	MACAddress          *string              `json:"mac_address"`
	FixedIPs            interface{}          `json:"fixed_ips"`
	DeviceID            *string              `json:"device_id"`
	DeviceOwner         *string              `json:"device_owner"`
	TenantID            *string              `json:"tenant_id"`
	SecurityGroups      []*string            `json:"security_groups"`
	AllowedAddressPairs []models.AddressPair `json:"allowed_address_pairs "`
}

func (opts *CreatePortOpts) IsValid() bool {
	return opts != nil && opts.NetworkID != nil
}

func (opts *CreatePortOpts) ToPayload() interface{} {
	type request struct {
		Port *CreatePortOpts `json:"port"`
	}

	return request{
		Port: opts,
	}
}

type UpdatePortOpts struct {
	Name                *string               `json:"name"`
	AdminStateUp        *bool                 `json:"admin_state_up"`
	FixedIPs            interface{}           `json:"fixed_ips"`
	DeviceID            *string               `json:"deivce_id"`
	DeviceOwner         *string               `json:"device_owner"`
	SecurityGroups      []*string             `json:"security_groups"`
	AllowedAddressPairs []*models.AddressPair `json:"allowed_address_pairs "`
}

func (opts *UpdatePortOpts) IsValid() bool {
	return true
}

func (opts *UpdatePortOpts) ToPayload() interface{} {
	type request struct {
		Port *UpdatePortOpts `json:"port"`
	}

	return request{
		Port: opts,
	}
}
