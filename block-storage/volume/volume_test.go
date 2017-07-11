package volume

import (
	"net/http"
	"testing"

	"strconv"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

func Test_All_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	jsonheader := http.Header{}
	jsonheader.Add("Content-Type", "application/json")

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+apiv2.GetString("admin.project_id")+"/volumes/detail")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /volumes"))
	// mitm.Pause()

	assertion := assert.New(t)

	volumes, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.NotNil(volumes)
	assertion.EqualValues(2, len(volumes))
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.id"), volumes[0].ID)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.name"), volumes[0].Name)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.size"), strconv.Itoa(volumes[0].Size))
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.status"), volumes[0].Status)
	assertion.Empty(volumes[0].Description)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.volume_type"), volumes[0].VolumeType)
	assertion.Empty(volumes[0].SnapshotID)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.bootable"), volumes[0].Bootable)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.created_at"), volumes[0].CreatedAt)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.attachments.0.id"), volumes[0].Attachments[0].ID)
}

func Test_All_Volume_By_Params(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	jsonheader := http.Header{}
	jsonheader.Add("Content-Type", "application/json")

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+apiv2.GetString("admin.project_id")+"/volumes")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /volumes"))
	// mitm.Pause()

	assertion := assert.New(t)

	volumes, err := New(openstacker).AllByParams(&options.ListVolumeOpts{})
	assertion.Nil(err)
	assertion.NotNil(volumes)
	assertion.EqualValues(2, len(volumes))
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.id"), volumes[0].ID)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.name"), volumes[0].Name)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.size"), strconv.Itoa(volumes[0].Size))
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.status"), volumes[0].Status)
	assertion.Empty(volumes[0].Description)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.volume_type"), volumes[0].VolumeType)
	assertion.Empty(volumes[0].SnapshotID)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.bootable"), volumes[0].Bootable)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.created_at"), volumes[0].CreatedAt)
	assertion.Equal(apiv2.APIString("GET /volumes.volumes.0.attachments.0.id"), volumes[0].Attachments[0].ID)
}

func Test_Create_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	jsonheader := http.Header{}
	jsonheader.Add("Content-Type", "application/json")

	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort("8776", "/v2/"+apiv2.GetString("admin.project_id")+"/volumes")).WithResponse(http.StatusAccepted, jsonheader, apiv2.APIString("POST /volumes"))
	//mitm.Pause()

	vos := New(openstacker)
	assertion := assert.New(t)

	_, err := vos.Create(&options.CreateVolumeOpts{
		Name:        options.String("test volume"),
		Description: options.String("test create volume"),
		VolumeType:  options.String("iscsi"),
		Size:        options.Int(10),
	})

	assertion.Nil(err)
}
