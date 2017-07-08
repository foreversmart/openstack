package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type RoleModel struct {
	Id       *string `mapstructure:"id" json:"id"`
	Name     *string `mapstructure:"name" json:"name"`
	DomainId *string `mapstructure:"domain_id" json:"domain_id"`
}

func ExtractRole(result gophercloud.Result) (role *RoleModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Role *RoleModel `mapstructure:"role" json:"role"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		role = response.Role
	}

	return
}

func ExtractRoles(result gophercloud.Result) (roles []*RoleModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Roles []*RoleModel `mapstructure:"roles" json:"roles"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		roles = response.Roles
	}

	return
}
