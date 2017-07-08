package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type UserDomainModel struct {
	ID   string `mapstructure:"id" json:"id"`
	Name string `mapstructure:"name" json:"name"`
}

type TokenUserModel struct {
	Domain            UserDomainModel `mapstructure:"domain" json:"domain"`
	ID                string          `mapstructure:"id" json:"id"`
	Name              string          `mapstructure:"name" json:"name"`
	PasswordExpiresAt string          `mapstructure:"password_expires_at" json:"password_expires_at"`
}

type TokenModel struct {
	Methods   []*string       `mapstructure:"methods" json:"methods"`
	User      *TokenUserModel `mapstructure:"user" json:"user"`
	Extras    interface{}     `mapstructure:"extras" json:"extras"`
	AuditIDs  []*string       `mapstructure:"audit_ids" json:"audit_ids"`
	IssuedAt  string          `mapstructure:"issued_at" json:"issued_at"`
	ExpiresAt string          `mapstructure:"expires_at" json:"expires_at"`
}

func ExtractToken(result gophercloud.Result) (token *TokenModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Token *TokenModel `mapstructure:"token" json:"token"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		token = response.Token
	}

	return
}
