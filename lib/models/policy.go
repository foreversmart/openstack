package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type PolicyModel struct {
	ID   string      `mapstructure:"id" json:"id"`
	Type string      `mapstructure:"type" json:"type"`
	Blob interface{} `mapstructure:"blob" json:"blob"`
}

func ExtractPolicy(result gophercloud.Result) (policy *PolicyModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Policy *PolicyModel `mapstructure:"policy" json:"policy"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		policy = response.Policy
	}

	return
}

func ExtractPolicies(result gophercloud.Result) (policies []*PolicyModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		Policies []*PolicyModel `mapstructure:"policies" json:"policies"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		policies = response.Policies
	}

	return
}
