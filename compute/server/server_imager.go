package server

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type ServerImager struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerImager(client ifaces.Openstacker) *ServerImager {
	return &ServerImager{
		Client: client,
	}
}
