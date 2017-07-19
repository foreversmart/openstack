package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type FloatingIPModel struct {
	ID          string `mapstructure:"id" json:"id"`
	Status      string `mapstructure:"status" json:"status"`
	RouterID    string `mapstructure:"router_id" json:"router_id"`
	ProjectID   string `mapstructure:"project_id" json:"project_id"`
	TenantID    string `mapstructure:"tenant_id" json:"tenant_id"`
	PortID      string `mapstructure:"port_id" json:"port_id"`
	FloatingIP  string `mapstructure:"floating_ip_address" json:"floating_ip_address"`
	FixedIP     string `mapstructure:"fixed_ip_address" json:"fixed_ip_address"`
	Description string `mapstructure:"description" json:"description"`
	CreatedAt   string `mapstructure:"created_at" json:"created_at"`
	UpdatedAt   string `mapstructure:"updated_at" json:"updated_at"`
}

func ExtractFloatingIP(result gophercloud.Result) (ip *FloatingIPModel, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	var response struct {
		FloatingIP *FloatingIPModel `mapstructure:"floatingip"`
	}

	err = mapstructure.Decode(result.Body, &response)
	ip = response.FloatingIP

	return
}

func ExtractFloatingIPs(result gophercloud.Result) (ips []*FloatingIPModel, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	return ExtractFloatingIPsByBody(result.Body)
}

func ExtractFloatingIPsByBody(body interface{}) (ips []*FloatingIPModel, err error) {
	var response struct {
		FloatingIPs []*FloatingIPModel `mapstructure:"floatingips"`
	}

	err = mapstructure.Decode(body, &response)
	ips = response.FloatingIPs

	return
}
