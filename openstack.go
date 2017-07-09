package openstack

import (
	"github.com/kirk-enterprise/openstack/internal"
	"github.com/kirk-enterprise/openstack/keystone"
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
