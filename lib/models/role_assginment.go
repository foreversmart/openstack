package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type RoleDomainModel struct {
	ID   string `mapstructure:"id" json:"id,omitempty"`
	Name string `mapstructure:"name" json:"name,omitempty"`
}

type RoleComponentModel struct {
	Domain RoleDomainModel `mapstructure:"domain" json:"domain,omitempty"`
	ID     string          `mapstructure:"id" json:"id,omitempty"`
	Name   string          `mapstructure:"name" json:"name,omitempty"`
}

type RoleScopeModel struct {
	Project RoleComponentModel `mapstructure:"project" json:"project,omitempty"`
	Domain  RoleComponentModel `mapstructure:"domain" json:"domain,omitempty"`
}

type RoleAssignmentModel struct {
	Role  RoleComponentModel `mapstructure:"role" json:"role"`
	Scope RoleScopeModel     `mapstructure:"scope" json:"scope"`
	User  RoleComponentModel `mapstructure:"user" json:"user"`
}

func ExtractRoleAssignments(result gophercloud.Result) (roleAssignments []*RoleAssignmentModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		RoleAssignments []*RoleAssignmentModel `mapstructure:"role_assignments" json:"role_assignments"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		roleAssignments = response.RoleAssignments
	}

	return
}
