package secrule

import (
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

type Secrule struct {
	Client ifaces.Openstacker

	_ bool
}
