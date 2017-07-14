package internal

import (
	"os"
	"testing"

	"github.com/qbox/openstack-golang-sdk/lib/auth"
	"github.com/qbox/openstack-golang-sdk/lib/testdata"
)

var (
	apiv2 *testdata.TestData
	apiv3 *testdata.TestData
)

func TestMain(m *testing.M) {
	// setup dependences
	apiv2 = testdata.NewWithFilename("./", auth.V2)
	apiv3 = testdata.NewWithFilename("./", auth.V3)

	// run testings
	code := m.Run()

	// exit
	os.Exit(code)
}
