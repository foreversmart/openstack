package compute

import (
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

type Image struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Image {
	return &Image{
		client: client,
	}
}
