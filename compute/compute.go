package compute

import (
	"github.com/qbox/openstack-golang-sdk/compute/flavors"
	"github.com/qbox/openstack-golang-sdk/compute/server"
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

func (c *Compute) NewServer() ifaces.Server {
	return server.New(c.client)
}

func (c *Compute) NewFlavors() ifaces.Flavorer {
	return flavors.New(c.client)
}
