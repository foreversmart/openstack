package volume

import (
	"net/http"
	"testing"

	"strconv"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

var (
	testVolumeId string
)

func Test_Create_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes")).WithResponse(http.StatusAccepted, jsonheader, apiv2.APIString("POST /volumes"))
	//mitm.Pause()

	assertion := assert.New(t)

	volume, err := New(openstacker).Create(&options.CreateVolumeOpts{
		Name:        options.String("test volume"),
		Description: options.String("test create volume"),
		VolumeType:  options.String("iscsi"),
		Size:        options.Int(10),
	})
	testVolumeId = volume.ID

	assertion.Nil(err)
	assertion.NotNil(volume)
}

func Test_All_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/detail")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /volumes"))
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

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/detail?limit=10")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /volumes"))
	// mitm.Pause()

	assertion := assert.New(t)
	tenant_id := "803d363c5c4649b591fd74b96a8c30f5"
	all_tenants := "1"

	volumes, err := New(openstacker).AllByParams(&options.ListVolumeOpts{
		Limit:      options.Int(10),
		TenantID:   &tenant_id,
		AllTenants: &all_tenants,
	})
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

func Test_Show_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	testVolumeId = "5aa119a8-d25b-45a7-8d1b-88e127885635"

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/"+testVolumeId)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /volumes/:id"))

	assertion := assert.New(t)

	volume, err := New(openstacker).Show(testVolumeId)
	assertion.Nil(err)
	assertion.NotNil(volume)

	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.id"), volume.ID)
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.name"), volume.Name)
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.size"), strconv.Itoa(volume.Size))
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.status"), volume.Status)
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.description"), volume.Description)
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.volume_type"), volume.VolumeType)
	assertion.Empty(volume.SnapshotID)
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.bootable"), volume.Bootable)
	assertion.Equal(apiv2.APIString("GET /volumes/:id.volume.created_at"), volume.CreatedAt)
}

func Test_Resize_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/"+testVolumeId+"/action")).WithResponse(http.StatusAccepted, nil, apiv2.APIString("POST /volumes/:id/action"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Resize(testVolumeId, &options.ResizeVolumeOpts{
		Size: options.Int(12),
	})
	assertion.Nil(err)
}

func Test_Reset_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/"+testVolumeId+"/action")).WithResponse(http.StatusAccepted, nil, apiv2.APIString("POST /volumes/:id/action"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Reset(testVolumeId, &options.ResetVolumeOpts{
		Status:          options.String("available"),
		AttachStatus:    options.String("detached"),
		MigrationStatus: options.String("migrating"),
	})
	assertion.Nil(err)
}

func Test_Update_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)
	testVolumeId = "5aa119a8-d25b-45a7-8d1b-88e127885635"
	mitm.MockRequest("PUT", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/"+testVolumeId)).WithResponse(http.StatusOK, nil, apiv2.APIString("PUT /volumes/:id"))
	//mitm.Pause()

	assertion := assert.New(t)
	tenant_id := "803d363c5c4649b591fd74b96a8c30f5"
	all_tenants := "1"

	volume, err := New(openstacker).Update(testVolumeId, &options.UpdateVolumeOpts{
		Name:        options.String("update volume name"),
		Description: options.String("update volume desc"),
		TenantID:    &tenant_id,
		AllTenants:  &all_tenants,
	})
	assertion.Nil(err)

	assertion.Equal(apiv2.APIString("PUT /volumes/:id.volume.id"), volume.ID)
}

func Test_Delete_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("DELETE", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes/"+testVolumeId)).WithResponse(http.StatusAccepted, nil, apiv2.APIString("DELETE /volumes/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(testVolumeId)
	assertion.Nil(err)
}
