package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type SecurityGroupModel struct {
	ID          string `mapstructure:"id" json:"id"`
	TenantID    string `mapstructure:"tenant_id" json:"tenant_id"`
	ProjectID   string `mapstructure:"project_id" json:"project_id"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
	// Rules       []*SecurityRuleModel `json:"security_group_rules" mapstructure:"security_group_rules"`"
}

func ExtractSecurityGroups(result gophercloud.Result) (securitygroups []*SecurityGroupModel, err error) {
	if result.Err != nil {
		return nil, result.Err
	}

	var response struct {
		SecurityGroups []*SecurityGroupModel `mapstructure:"security_groups"`
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
		SecurityGroup *SecurityGroupModel `mapstructure:"security_group"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		securitygroup = response.SecurityGroup
	}

	return
}
