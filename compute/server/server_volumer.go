package server

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type ServerVolumer struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerVolumer(client ifaces.Openstacker) *ServerVolumer {
	return &ServerVolumer{
		Client: client,
	}
}
