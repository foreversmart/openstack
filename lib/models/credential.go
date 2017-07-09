package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type CredentialModel struct {
	ID        string `mapstructure:"id" json:"id"`
	ProjectID string `mapstructure:"project_id" json:"project_id"`
	UserID    string `mapstructure:"user_id" json:"user_id"`
	Type      string `mapstructure:"type" json:"type"`
	Blob      string `mapstructure:"blob" json:"blob"`
}

func ExtractCredential(result gophercloud.Result) (credential *CredentialModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Credential *CredentialModel `mapstructure:"credential" json:"credential"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		credential = response.Credential
	}

	return
}

func ExtractCredentials(result gophercloud.Result) (credentials []*CredentialModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Credentials []*CredentialModel `mapstructure:"credentials" json:"credentials"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		credentials = response.Credentials
	}

	return
}
