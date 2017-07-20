package image

import (
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

func Test_Image(t *testing.T) {
	assertion := assert.New(t)

	i := New(nil)
	assertion.NotNil(i)
	assertion.Implements((*ifaces.Imager)(nil), i.NewImager())
}
