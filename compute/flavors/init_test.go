package flavors

import (
	"net/http"
	"os"
	"testing"

	"github.com/dolab/httpmitm"
	"github.com/qbox/openstack-golang-sdk/internal"
	"github.com/qbox/openstack-golang-sdk/lib/auth"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/testdata"
	"github.com/rackspace/gophercloud"
)

var (
	testProjectId string

	apiv3 *testdata.TestData

	mocker      *httpmitm.MitmTransport
	openstacker ifaces.Openstacker
	jsonheader  http.Header
)

func TestMain(m *testing.M) {
	// setup dependences
	apiv3 = testdata.NewWithFilename("../", auth.V3)

	testProjectId = apiv3.GetString("admin.project_id")

	mocker = httpmitm.NewMitmTransport().StubDefaultTransport(nil)
	defer mocker.UnstubDefaultTransport()

	jsonheader = http.Header{}
	jsonheader.Add("Content-Type", "application/json")
	jsonheader.Add("X-Subject-Token", apiv3.GetString("token.id"))

	mocker.MockRequest("POST", apiv3.MockAdminURL("/v3/auth/tokens")).WithResponse(201, jsonheader, apiv3.APIString("scoped"))
	// mocker.Pause()

	openstacker = internal.New(apiv3.GetString("admin.endpoint"))

	err := openstacker.AuthByPassword(auth.AuthOptions{
		Version: auth.V3,
		AuthOptions: &gophercloud.AuthOptions{
			DomainName: apiv3.GetString("admin.domain_name"),
			Username:   apiv3.GetString("admin.username"),
			Password:   apiv3.GetString("admin.password"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	// run testings
	code := m.Run()

	// exit
	os.Exit(code)
}
