package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type DomainModel struct {
	ID          string `mapstructure:"id" json:"id"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
	Enabled     bool   `mapstructure:"enabled" json:"enabled"`
}

func ExtractDomain(result gophercloud.Result) (domain *DomainModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Domain *DomainModel `mapstructure:"domain" json:"domain"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		domain = response.Domain
	}

	return
}

func ExtractDomains(result gophercloud.Result) (domains []*DomainModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Domains []*DomainModel `mapstructure:"domains" json:"domains"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		domains = response.Domains
	}

	return
}
