package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type SecurityGroupModel struct {
	ID          string `json:"id" mapstructure:"id"`
	TenantID    string `json:"tenant_id" mapstructure:"tenant_id"`
	ProjectID   string `json:"project_id" mapstructure:"project_id"`
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	// Rules       []*SecurityRuleModel `json:"security_group_rules" mapstructure:"security_group_rules"`"
}

func ExtractSecurityGroups(result gophercloud.Result) (securitygroups []*SecurityGroupModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		SecurityGroups []*SecurityGroupModel `mapstructure:"security_groups" json:"security_groups"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		securitygroups = response.SecurityGroups
	}

	return
}

func ExtractSecurityGroup(result gophercloud.Result) (securitygroup *SecurityGroupModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		SecurityGroup *SecurityGroupModel `mapstructure:"security_group" json:"security_group"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		securitygroup = response.SecurityGroup
	}

	return
}
