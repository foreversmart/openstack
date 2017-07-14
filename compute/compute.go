package compute

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type Compute struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Compute {
	return &Compute{
		client: client,
	}
}
