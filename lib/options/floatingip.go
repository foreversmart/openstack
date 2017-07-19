package options

/**
 * used to create floatingip
 */
type CreateFloatingIPOpts struct {
	ProjectID         *string `mapstructure:"project_id" json:"project_id"`
	TenantID          *string `mapstructure:"tenant_id" json:"tenant_id"`
	FloatingNetworkID *string `mapstructure:"floating_network_id" json:"floating_network_id"`

	// following are optional ref: https://developer.openstack.org/api-ref/networking/v2/index.html
	PortID      *string `mapstructure:"port_id" json:"port_id"`
	FloatingIP  *string `mapstructure:"floating_ip_address" json:"floating_ip_address"`
	FixedIP     *string `mapstructure:"fixed_ip_address" json:"fixed_ip_address"`
	SubnetID    *string `mapstructure:"subnet_id" json:"subnet_id"`
	Description *string `mapstructure:"description" json:"description"`
}

func (opts *CreateFloatingIPOpts) IsValid() bool {
	return opts != nil &&
		opts.FloatingNetworkID != nil &&
		opts.ProjectID != nil &&
		opts.TenantID != nil
}

func (opts *CreateFloatingIPOpts) ToPayload() interface{} {
	type payload struct {
		FloatingIP *CreateFloatingIPOpts `json:"floatingip"`
	}

	return payload{
		opts,
	}
}

/**
 * used to update floatingip
 */
type UpdateFloatingIPOpts struct {
	PortID      *string `mapstructure:"port_id" json:"port_id"`
	FixedIP     *string `mapstructure:"fixed_ip_address" json:"fixed_ip_address"`
	Description *string `mapstructure:"description" json:"description"`
}

func (opts *UpdateFloatingIPOpts) IsValid() bool {
	return opts != nil
}

func (opts *UpdateFloatingIPOpts) ToPayload() interface{} {
	type payload struct {
		FloatingIP *UpdateFloatingIPOpts `json:"floatingip"`
	}

	return payload{
		opts,
	}
}
