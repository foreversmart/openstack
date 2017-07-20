package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type KeypairModel struct {
	// ID          string `mapstructure:"id" json:"id"`
	Name string `mapstructure:"name" json:"name"`
	// Deleted     bool   `mapstructure:"deleted" json:"deleted"`
	// DeletedAt   string `mapstructure:"deleted_at" json:"deleted_at"`
	Fingerprint string `mapstructure:"fingerprint" json:"fingerprint"`
	PublicKey   string `mapstructure:"public_key" json:"public_key"`
	UpdatedAt   string `mapstructure:"updated_at" json:"updated_at"`
	UserId      string `mapstructure:"user_id" json:"user_id"`
	Type        string `mapstructure:"type" json:"type"`
}

func ExtractKeypair(r gophercloud.Result) (keypair *KeypairModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var resp struct {
		Keypair *KeypairModel `mapstructure:"keypair"`
	}

	err = mapstructure.Decode(r.Body, &resp)
	keypair = resp.Keypair

	return
}

func ExtractKeypairs(r gophercloud.Result) (keypairs []*KeypairModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	return ExtractKeypairsByBody(r.Body)
}

func ExtractKeypairsByBody(body interface{}) (keypairs []*KeypairModel, err error) {

	type KeypairWrapper struct {
		Keypair *KeypairModel `mapstructure:"keypair"`
	}

	var resp struct {
		Keypairs []*KeypairWrapper `mapstructure:"keypairs"`
	}

	err = mapstructure.Decode(body, &resp)

	if err == nil {
		keypairs = make([]*KeypairModel, 0)
		for _, kw := range resp.Keypairs {
			keypairs = append(keypairs, kw.Keypair)
		}
	}
	return
}
