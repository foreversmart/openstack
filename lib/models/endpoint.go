package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type EndpointModel struct {
	ID        string `json:"id" mapstructure:"id"`
	Url       string `json:"url" mapstructure:"url"`
	Interface string `json:"interface" mapstructure:"interface"`
	Enabled   bool   `json:"enabled" mapstructure:"enabled"`
	RegionID  string `json:"region_id" mapstructure:"region_id"`
	Region    string `json:"region" mapstructure:"region"`
	ServiceID string `json:"service_id" mapstructure:"service_id"`
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
