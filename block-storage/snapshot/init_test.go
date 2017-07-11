package snapshot

import (
	"net/http"
	"os"
	"testing"

	"fmt"
	"time"

	"github.com/dolab/httpmitm"
	"github.com/kirk-enterprise/openstack-golang-sdk/internal"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/auth"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/testdata"
	"github.com/rackspace/gophercloud"
)

var (
	testVolumeId   string
	testSnapshotId string
	testProjectId  string

	apiv3 *testdata.TestData
	apiv2 *testdata.TestData

	mocker      *httpmitm.MitmTransport
	openstacker ifaces.Openstacker
	jsonheader  http.Header
)

func TestMain(m *testing.M) {
	// setup dependences
	apiv3 = testdata.NewWithFilename("../", auth.V3)
	apiv2 = testdata.NewWithFilename("../", auth.V2)
	testVolumeId = apiv3.GetString("volume.id")
	testSnapshotId = apiv3.GetString("volume.snapshot_id")
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
		SuccessFunc: func(tokenID string, catalog string, expiredAt time.Time) error {
			fmt.Printf("New token: %v \n", tokenID)
			return nil
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
