package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type FlavorModel struct {
	//The ID of the flavor.
	Id string `json:"id"`

	//The display name of a flavor.
	Name string `json:"name"`

	//The amount of RAM a flavor has, MB.
	Ram int `json:"ram"`

	//The size of the root disk that will be created, GB.
	Disk int `json:"disk"`

	//The number of virtual CPUs.
	Vcpus int `json:"vcpus"`

	//The size of a dedicated swap disk, GB
	Swap string `json:"swap"`

	// The receive / transmit factor, It defaults to 1.0.
	RxtxFactor float64 `json:"rxtx_factor"`
}

func ExtractFlavors(result gophercloud.Result) (flavors []*FlavorModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Flavors []*FlavorModel `mapstructure:"flavors" json:"flavors"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		flavors = response.Flavors
	}

	return
}

func ExtractFlavor(result gophercloud.Result) (flavor *FlavorModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}
	var response struct {
		Flavor *FlavorModel `mapstructure:"flavor" json:"flavor"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		flavor = response.Flavor
	}

	return

}
