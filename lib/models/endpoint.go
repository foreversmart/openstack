package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type EndpointModel struct {
	Id        string `json:"id" mapstructure:"id"`
	RegionId  string `json:"region_id" mapstructure:"region_id"`
	Region    string `json:"region" mapstructure:"region"`
	Interface string `json:"interface" mapstructure:"interface"`
	ServiceId string `json:"service_id" mapstructure:"service_id"`
	Enabled   bool   `json:"enabled" mapstructure:"enabled"`
}

func ExtractEndpoint(result gophercloud.Result) (endpoint *EndpointModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Endpoint *EndpointModel `mapstructure:"endpoint" json:"endpoint"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		endpoint = response.Endpoint
	}

	return
}

func ExtractEndpoints(result gophercloud.Result) (endpoints []*EndpointModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Endpoints []*EndpointModel `mapstructure:"endpoints" json:"endpoints"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		endpoints = response.Endpoints
	}

	return
}
