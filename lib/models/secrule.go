package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

// SecGroupRule represents a rule to dictate the behaviour of incoming or
// outgoing traffic for a particular security group.
type SecGroupRuleModel struct {
	// The UUID for this security group rule.
	ID string `mapstructure:"id" json:"id"`

	// The direction in which the security group rule is applied. The only values
	// allowed are "ingress" or "egress". For a compute instance, an ingress
	// security group rule is applied to incoming (ingress) traffic for that
	// instance. An egress rule is applied to traffic leaving the instance.
	Direction string `mapstructure:"direction" json:"direction"`

	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the
	// ingress or egress rules.
	EtherType string `mapstructure:"ethertype" json:"ethertype"`

	// The security group ID to associate with this security group rule.
	SecGroupID string `mapstructure:"security_group_id" json:"security_group_id"`

	// The minimum port number in the range that is matched by the security group
	// rule. If the protocol is TCP or UDP, this value must be less than or equal
	// to the value of the PortRangeMax attribute. If the protocol is ICMP, this
	// value must be an ICMP type.
	PortRangeMin int `mapstructure:"port_range_min" json:"port_range_min"`

	// The maximum port number in the range that is matched by the security group
	// rule. The PortRangeMin attribute constrains the PortRangeMax attribute. If
	// the protocol is ICMP, this value must be an ICMP type.
	PortRangeMax int `mapstructure:"port_range_max" json:"port_range_max"`

	// The protocol that is matched by the security group rule. Valid values are
	// "tcp", "udp", "icmp" or an empty string.
	Protocol string `mapstructure:"protocol" json:"protocol"`

	// The remote group ID to be associated with this security group rule. You
	// can specify either RemoteGroupID or RemoteIPPrefix.
	RemoteGroupID string `mapstructure:"remote_group_id" json:"remote_group_id"`

	RemoteGroupName string `mapstructure:"remote_group_name" json:"remote_group_name"`

	// The remote IP prefix to be associated with this security group rule. You
	// can specify either RemoteGroupID or RemoteIPPrefix . This attribute
	// matches the specified IP prefix as the source IP address of the IP packet.
	RemoteIPPrefix string `mapstructure:"remote_ip_prefix" json:"remote_ip_prefix"`

	// The owner of this security group rule.
	TenantID  string `mapstructure:"tenant_id" json:"tenant_id"`
	ProjectID string `mapstructure:"project_id" json:"project_id"`

	Description string `mapstructure:"description" json:"description"`
}

func ExtractSecRule(r gophercloud.Result) (rule *SecGroupRuleModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var resp struct {
		Secrule *SecGroupRuleModel `mapstructure:"security_group_rule" json:"security_group_rule"`
	}

	err = mapstructure.Decode(r.Body, &resp)
	if err == nil {
		rule = resp.Secrule
	}
	return
}

func ExtractSecRulesByBody(body interface{}) (rules []*SecGroupRuleModel, err error) {
	var resp struct {
		RuleInfos []*SecGroupRuleModel `mapstructure:"security_group_rules" json:"security_group_rules"`
	}

	err = mapstructure.Decode(body, &resp)
	if err == nil {
		rules = resp.RuleInfos
	}
	return
}

func ExtractSecRules(r gophercloud.Result) (rules []*SecGroupRuleModel, err error) {
	if r.Err != nil {
		return nil, r.Err
	}

	return ExtractSecRulesByBody(r.Body)
}
