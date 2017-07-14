package openstack

import (
	"github.com/qbox/openstack-golang-sdk/internal"
	"github.com/qbox/openstack-golang-sdk/keystone"
)

type Openstack struct {
	*internal.Openstack

	_ bool
}

func New(endpoint string) *Openstack {
	return &Openstack{
		Openstack: internal.New(endpoint),
	}
}

func (os *Openstack) NewKeystone() *keystone.Keystone {
	return keystone.New(os)
}
