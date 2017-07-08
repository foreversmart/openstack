package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type ProjectModel struct {
	// ID is a unique identifier for this tenant.
	ID string `mapstructure:"id" json:"id"`

	// Name is a friendlier user-facing name for this tenant.
	Name string `mapstructure:"name" json:"name"`

	// Description is a human-readable explanation of this Tenant's purpose.
	Description string `mapstructure:"description" json:"description"`

	ParentID string `mapstructure:"parent_id" json:"parent_id"`

	// Indicates whether the project also acts as a domain.
	IsDomain bool `mapstructure:"is_domain" json:"is_domain"`

	DomainID string `mapstructure:"domain_id" json:"domain_id"`

	// Enabled indicates whether or not a tenant is active.
	Enabled bool `mapstructure:"enabled" json:"enabled"`
}

func ExtractProject(result gophercloud.Result) (project *ProjectModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Project *ProjectModel `mapstructure:"project" json:"project"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		project = response.Project
	}

	return
}

func ExtractProjects(result gophercloud.Result) (projects []*ProjectModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Projects []*ProjectModel `mapstructure:"projects" json:"projects"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		projects = response.Projects
	}

	return
}

func ExtractTenant(result gophercloud.Result) (project *ProjectModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Project *ProjectModel `mapstructure:"tenant" json:"tenant"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		project = response.Project
	}

	return
}

func ExtractTenants(result gophercloud.Result) (projects []*ProjectModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Projects []*ProjectModel `mapstructure:"tenants" json:"tenants"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		projects = response.Projects
	}

	return
}
