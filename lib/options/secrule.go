package options

type CreateSecruleOpts struct {
	RemoteGroupID  *string `json:"remote_group_id"`
	Direction      *string `json:"direction"`
	Protocol       *string `json:"protocol"`
	EtherType      *string `json:"ethertype"`
	PortRangeMin   *int    `json:"port_range_min"`
	PortRangeMax   *int    `json:"port_range_max"`
	SecGroupID     *string `json:"security_group_id"`
	RemoteIPPrefix *string `json:"remote_ip_prefix"`
	Description    *string `json:"description"`
}

func (opts *CreateSecruleOpts) IsValid() bool {
	return opts != nil && opts.Direction != nil && opts.RemoteIPPrefix != nil
}

func (opts *CreateSecruleOpts) ToPayload() interface{} {
	type request struct {
		Secrule *CreateSecruleOpts `json:"secrule"`
	}

	return request{
		Secrule: opts,
	}
}

type ListSecRuleOpts struct {
	ID             *string `json:"id"`
	RemoteGroupID  *string `json:"remote_group_id"`
	Direction      *string `json:"direction"`
	Protocol       *string `json:"protocol"`
	EtherType      *string `json:"ethertype"`
	PortRangeMin   *int    `json:"port_range_min"`
	PortRangeMax   *int    `json:"port_range_max"`
	SecGroupID     *string `json:"security_group_id"`
	TenantID       *string `json:"tenant_id"`
	ProjectID      *string `json:"project_id"`
	RemoteIPPrefix *string `json:"remote_ip_prefix"`
	Description    *string `json:"description"`
}

func (opts *ListSecRuleOpts) IsValid() bool {
	return true
}
