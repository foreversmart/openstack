package internal

import (
	"os"
	"testing"

	"github.com/kirk-enterprise/openstack/lib/auth"
	"github.com/kirk-enterprise/openstack/lib/testdata"
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
