package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type RegionModel struct {
	Id             *string `json:"id"`
	Description    *string `json:"description"`
	ParentRegionId *string `json:"parent_region_id"`
}

func ExtractRegion(result gophercloud.Result) (region *RegionModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Region *RegionModel `mapstructure:"region" json:"region"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		region = response.Region
	}

	return
}

func ExtractRegions(result gophercloud.Result) (regions []*RegionModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Regions []*RegionModel `mapstructure:"regions" json:"regions"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		regions = response.Regions
	}

	return
}
