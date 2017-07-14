package networking

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/networking/floatingip"
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
