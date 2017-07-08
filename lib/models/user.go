package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

// Identity V3 User
type UserModel struct {
	DomainID          string `mapstructure:"domain_id" json:"domain_id"`
	ID                string `mapstructure:"id" json:"id"`
	Name              string `mapstructure:"name" json:"name"`
	Enabled           bool   `mapstructure:"enabled" json:"enabled"`
	DefaultProjectId  string `mapstructure:"default_project_id" json:"default_project_id"`
	PasswordExpiresAt string `mapstructure:"password_expires_at" json:"password_expires_at"`
}

func ExtractUser(result gophercloud.Result) (user *UserModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		User *UserModel `mapstructure:"user" json:"user"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		user = response.User
	}

	return
}

func ExtractUsers(result gophercloud.Result) (users []*UserModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Users []*UserModel `mapstructure:"users" json:"users"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		users = response.Users
	}

	return
}
