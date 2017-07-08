package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type ServiceModel struct {
	Id      string `json:"id" mapstructure:"id"`
	Type    string `json:"type" mapstructure:"type"`
	Name    string `json:"name" mapstructure:"name"`
	Desc    string `json:"description" mapstructure:"description"`
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
}

func ExtractService(result gophercloud.Result) (service *ServiceModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Service *ServiceModel `mapstructure:"service" json:"service"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		service = response.Service
	}

	return
}

func ExtractServicesByResult(result gophercloud.Result) (services []*ServiceModel, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	var responseonse struct {
		Services []*ServiceModel `mapstructure:"services"`
	}

	err = mapstructure.Decode(result.Body, &responseonse)
	if err == nil {
		services = responseonse.Services
	}

	return
}
