package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type GroupModel struct {
	DomainID    string `mapstructure:"domain_id" json:"domain_id"`
	ID          string `mapstructure:"id" json:"id"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
}

func ExtractGroup(result gophercloud.Result) (group *GroupModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Group *GroupModel `mapstructure:"group" json:"group"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		group = response.Group
	}

	return
}

func ExtractGroups(result gophercloud.Result) (groups []*GroupModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Groups []*GroupModel `mapstructure:"groups" json:"groups"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		groups = response.Groups
	}

	return
}

func ExtractUserGroups(result gophercloud.Result) (groups []*GroupModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Groups []*GroupModel `mapstructure:"groups"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		groups = response.Groups
	}

	return
}
