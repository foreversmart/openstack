package compute

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"

type Compute struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Compute {
	return &Compute{
		client: client,
	}
}
