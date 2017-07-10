package models

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type NetworkModel struct {
	Id           string    `mapstructure:"id" json:"id"`
	Name         string    `mapstructure:"name" json:"name"`
	Subnets      []string  `mapstructure:"subnets" json:"subnets"`
	Status       string    `mapstructure:"status" json:"status"`
	TenantId     string    `mapstructure:"tenant_id" json:"tenant_id"`
	Shared       bool      `mapstructure:"shared" json:"shared"`
	AdminStateUp bool      `mapstructure:"admin_state_up" json:"admin_state_up"`
	External     bool      `json:"router:external"`
	RateLimit    int       `json:"rate_limit"`
	CreatedAt    time.Time `json:"created_at"`
}

func ExtractNetworks(result gophercloud.Result) (networks []*NetworkModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var resp struct {
		Networks []*NetworkModel `mapstructure:"networks"`
	}

	err = mapstructure.Decode(result.Body, &resp)
	return resp.Networks, err
}

func ExtractNetworkByResult(result gophercloud.Result) (network *NetworkModel, err error) {
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

func ExtractNetworksByPage(page pagination.Page) (networks []*NetworkModel, err error) {
	var response struct {
		NetworkInfos []*NetworkModel `mapstructure:"networks"`
	}

	err = mapstructure.Decode(page.GetBody(), &response)
	return response.NetworkInfos, err
}

type NetworkIpAvailabilitiesModel struct {
	NetworkId            string                      `json:"network_id"`
	NetworkName          string                      `json:"network_name"`
	SubnetIpAvailability []SubnetIpAvailabilityModel `json:"subnet_ip_availability"` //A list of dictionaries showing subnet IP availability. It contains information for every subnet associated to the network.
	ProjectId            string                      `json:"project_id"`
	TenantId             string                      `json:"tenant_id"`
	TotalIps             float64                     `json:"total_ips"` //The total number of IP addresses in a network.
	UsedIps              float64                     `json:"used_ips"`  //The number of used IP addresses of all subnets in a network.
}

type SubnetIpAvailabilityModel struct {
	Cidr        string  `json:"cidr"`
	IpVersion   int64   `json:"ip_version"`
	SubnetId    string  `json:"subnet_id"`
	subnet_name string  `json:"subnet_name"`
	total_ips   float64 `json:"total_ips"`
	used_ips    float64 `json:"used_ips"`
}
