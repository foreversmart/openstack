package compute

import (
	"github.com/qbox/openstack-golang-sdk/compute/flavor"
	"github.com/qbox/openstack-golang-sdk/compute/keypair"
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

// func (c *Compute) NewServerManager() ifaces.ServerManager {
// 	return server.NewServerManager(c.client)
// }

func (c *Compute) NewServerImager() ifaces.ServerImager {
	return server.NewServerImager(c.client)
}

func (c *Compute) NewServerKeyer() ifaces.ServerKeyer {
	return server.NewServerKeyer(c.client)
}

// func (c *Compute) NewServerPorter() ifaces.ServerPorter {
// 	return server.NewServerPorter(c.client)
// }

// func (c *Compute) NewServerVolumer() ifaces.ServerVolumer {
// 	return server.NewServerVolumer(c.client)
// }

func (c *Compute) NewFlavor() ifaces.Flavor {
	return flavor.New(c.client)
}

func (c *Compute) NewKeypair() ifaces.Keypair {
	return keypair.New(c.client)
}
