package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type CatalogModel struct {
	ID        string           `mapstructure:"id" json:"id"`
	Type      string           `mapstructure:"type" json:"type"`
	Name      string           `mapstructure:"name" json:"name"`
	Endpoints []*EndpointModel `mapstructure:"endpoints" json:"endpoints"`
}

func ExtractCatalog(result gophercloud.Result) (catalog *CatalogModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Catalog *CatalogModel `mapstructure:"catalog" json:"catalog"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		catalog = response.Catalog
	}

	return
}
