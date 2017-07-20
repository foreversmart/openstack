package server

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type ServerKeyer struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerKeyer(client ifaces.Openstacker) *ServerKeyer {
	return &ServerKeyer{
		Client: client,
	}
}
