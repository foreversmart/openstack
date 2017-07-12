package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type NetworkModel struct {
	ID           string   `mapstructure:"id" json:"id"`
	Name         string   `mapstructure:"name" json:"name"`
	Subnets      []string `mapstructure:"subnets" json:"subnets"`
	Status       string   `mapstructure:"status" json:"status"`
	ProjectID    string   `mapstructure:"project_id" json:"project_id"`
	TenantID     string   `mapstructure:"tenant_id" json:"tenant_id"`
	Shared       bool     `mapstructure:"shared" json:"shared"`
	AdminStateUp bool     `mapstructure:"admin_state_up" json:"admin_state_up"`
	CreatedAt    string   `mapstructure:"created_at" json:"created_at"`
	External     bool     `mapstructure:"router:external"  json:"router:external"`
}

func ExtractNetwork(result gophercloud.Result) (network *NetworkModel, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	var response struct {
		Network *NetworkModel `mapstructure:"network"`
	}

	err = mapstructure.Decode(result.Body, &response)
	return response.Network, err
}

func ExtractNetworks(r gophercloud.Result) (networks []*NetworkModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	return ExtractNetworksByBody(r.Body)
}

func ExtractNetworksByBody(body interface{}) (networks []*NetworkModel, err error) {
	var resp struct {
		NetworkInfos []*NetworkModel `mapstructure:"networks"`
	}

	err = mapstructure.Decode(body, &resp)
	if err == nil {
		networks = resp.NetworkInfos
	}
	return
}

type NetworkIpAvailabilitiesModel struct {
	NetworkID            string                      `json:"network_id"`
	NetworkName          string                      `json:"network_name"`
	SubnetIpAvailability []SubnetIpAvailabilityModel `json:"subnet_ip_availability"` //A list of dictionaries showing subnet IP availability. It contains information for every subnet associated to the network.
	ProjectID            string                      `json:"project_id"`
	TenantID             string                      `json:"tenant_id"`
	TotalIps             float64                     `json:"total_ips"` //The total number of IP addresses in a network.
	UsedIps              float64                     `json:"used_ips"`  //The number of used IP addresses of all subnets in a network.
}

type SubnetIpAvailabilityModel struct {
	Cidr       string  `json:"cidr"`
	IpVersion  int64   `json:"ip_version"`
	SubnetID   string  `json:"subnet_id"`
	SubnetName string  `json:"subnet_name"`
	TotalIps   float64 `json:"total_ips"`
	UsedIps    float64 `json:"used_ips"`
}
