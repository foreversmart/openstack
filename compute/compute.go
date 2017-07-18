package compute

import (
	"github.com/qbox/openstack-golang-sdk/compute/flavors"
	"github.com/qbox/openstack-golang-sdk/compute/servers"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

type Compute struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Compute {
	return &Compute{
		client: client,
	}
}

func (c *Compute) NewServerser() ifaces.Serverser {
	return servers.New(c.client)
}

func (c *Compute) NewFlavors() ifaces.Flavorer {
	return flavors.New(c.client)
}
