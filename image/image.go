package image

import (
	"github.com/qbox/openstack-golang-sdk/image/image"
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

func (i *Image) NewImager() ifaces.Imager {
	return image.New(i.client)
}
