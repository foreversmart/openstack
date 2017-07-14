package networking

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/networking/floatingip"
	"github.com/kirk-enterprise/openstack-golang-sdk/networking/network"
	"github.com/kirk-enterprise/openstack-golang-sdk/networking/subnet"
)

type Networking struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Networking {
	return &Networking{
		client: client,
	}
}

func (n *Networking) NewFloatingIPer() ifaces.FloatingIPer {
	return floatingip.New(n.client)
}

func (n *Networking) NewNetworker() ifaces.Networker {
	return network.New(n.client)
}

func (n *Networking) NewSubneter() ifaces.Subneter {
	return subnet.New(n.client)
}
