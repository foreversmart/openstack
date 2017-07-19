package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type KeypairModel struct {
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	Deleted     *bool   `json:"deleted"`
	DeletedAt   *string `json:"deleted_at"`
	Fingerprint *string `json:"fingerprint"`
	PublicKey   *string `json:"public_key"`
	UpdatedAt   *string `json:"updated_at"`
	UserId      *string `json:"user_id"`
	Type        *string `json:"type"`
}

func ExtractKeypair(r gophercloud.Result) (keypair *KeypairModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var resp struct {
		Keypair *KeypairModel `mapstructure:"keypair"`
	}

	err = mapstructure.Decode(r.Body, &resp)
	if err == nil {
		keypair = resp.Keypair
	}
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
		keypairs = make([]*KeypairModel, 1)
		for _, kw := range resp.Keypairs {
			keypairs = append(keypairs, kw.Keypair)
		}
	}
	return
}
