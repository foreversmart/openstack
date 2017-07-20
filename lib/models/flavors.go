package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type FlavorModel struct {
	//The ID of the flavor.
	ID string `mapstructure:"id" json:"id"`

	//The display name of a flavor.
	Name string `mapstructure:"name" json:"name"`

	//The amount of RAM a flavor has, MB.
	Ram int `mapstructure:"ram" json:"ram"`

	//The size of the root disk that will be created, GB.
	Disk int `mapstructure:"disk" json:"disk"`

	//The number of virtual CPUs.
	Vcpus int `mapstructure:"vcpus" json:"vcpus"`

	//The size of a dedicated swap disk, GB
	Swap string `mapstructure:"swap" json:"swap"`

	// The receive / transmit factor, It defaults to 1.0.
	RxtxFactor float64 `mapstructure:"rxtx_factor" json:"rxtx_factor"`
}

func ExtractFlavors(result gophercloud.Result) (flavors []*FlavorModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Flavor []*FlavorModel `mapstructure:"flavors" json:"flavors"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		flavors = response.Flavor
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
