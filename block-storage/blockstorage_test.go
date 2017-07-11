package blockstorage

import (
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
)

func Test_BlockStorage(t *testing.T) {
	assertion := assert.New(t)

	bs := New(nil)
	assertion.NotNil(bs)
	assertion.Implements((*ifaces.Volumer)(nil), bs.NewVolumer())
}
