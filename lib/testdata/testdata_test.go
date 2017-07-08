package testdata

import (
	"testing"

	"github.com/golib/assert"
	tokens2 "github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

func Test_TestData(t *testing.T) {
	assertion := assert.New(t)

	td := New("v2")
	assertion.Equal("v2", td.Version())
	assertion.NotEmpty(td.Data())
}

func Test_TestData_Get(t *testing.T) {
	assertion := assert.New(t)

	td := New("v2")

	// should work
	endpoint := td.GetString("endpoint")
	assertion.Equal("mitm://api.testing/v2.0", endpoint)

	// should work with embedded
	endpoint = td.GetString("user.endpoint")
	assertion.Equal("mitm://api.testing/v2.0/user", endpoint)

	// should work with array subscript
	compute := td.GetString("services.0.type")
	assertion.Equal("compute", compute)

	// should panic
	assertion.Panics(func() {
		td.GetString("unknown-json-key")
	})
}

func Test_TestData_API(t *testing.T) {
	assertion := assert.New(t)

	td := New("v2")

	// should work
	id := td.APIString("catalog.scoped.access.token.tenant.id")
	assertion.Equal("63397e0193f04dc4b2165490669ed4a1", id)

	// should panic
	assertion.Panics(func() {
		td.APIString("unknown-json-key")
	})
}

func Test_TestData_MockURL(t *testing.T) {
	assertion := assert.New(t)

	td := New("v2")

	// should work
	absurl := td.MockURLWithSSL("projects/detail")
	assertion.Equal("https://api.testing/v2.0/user/projects/detail", absurl)
}

func Test_TestData_MockAdminURL(t *testing.T) {
	assertion := assert.New(t)

	td := New("v2")

	// should work
	absurl := td.MockAdminURLWithSSL("projects/detail")
	assertion.Equal("https://api.testing/v2.0/admin/projects/detail", absurl)
}

func Test_TestData_MockResourceURL(t *testing.T) {
	assertion := assert.New(t)

	td := New("v2")

	// should work
	absurl := td.MockResourceURLWithSSL("compute/servers/detail")
	assertion.Equal("https://api.testing/v2.0/compute/servers/detail", absurl)
}

func Test_TestData_MockCatalog(t *testing.T) {
	assertion := assert.New(t)

	// should work for v2
	v2 := New("v2")

	v2catalog, ok := v2.MockCatalog().(*tokens2.ServiceCatalog)
	assertion.True(ok)
	assertion.NotEmpty(v2catalog.Entries)

	// should work for v3
	v3 := New("v3")

	v3catalog, ok := v3.MockCatalog().(*tokens3.ServiceCatalog)
	assertion.True(ok)
	assertion.NotEmpty(v3catalog.Entries)
}
