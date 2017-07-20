package server

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type ServerPorter struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerPorter(client ifaces.Openstacker) *ServerPorter {
	return &ServerPorter{
		Client: client,
	}
}
