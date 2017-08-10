package snapshot

import (
	"net/http"
	"testing"

	"strconv"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/blockstorage/volume"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

var (
	testSnapshotId string
	testVolumeId   string
)

func Test_Create_Snapshot(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	// create volume first
	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/volumes")).WithResponse(http.StatusAccepted, jsonheader, apiv2.APIString("POST /volumes"))
	//mitm.Pause()

	assertion := assert.New(t)

	volume, err := volume.New(openstacker).Create(&options.CreateVolumeOpts{
		Name:        options.String("test volume"),
		Description: options.String("test create volume"),
		VolumeType:  options.String("iscsi"),
		Size:        options.Int(10),
	})
	testVolumeId = volume.ID

	assertion.Nil(err)
	assertion.NotNil(volume)

	// test create snapshot
	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/snapshots")).WithResponse(http.StatusAccepted, jsonheader, apiv2.APIString("POST /snapshots"))
	// mitm.Pause()

	snapshot, err := New(openstacker).Create(&options.CreateSnapshotOpts{
		VolumeID:    options.String(testVolumeId),
		Name:        options.String("test snapshot"),
		Description: options.String("test create snapshot"),
		Force:       options.Bool(false),
	})

	testSnapshotId = snapshot.ID

	assertion.Nil(err)
	assertion.NotNil(snapshot)
	assertion.Equal(apiv2.APIString("POST /snapshots.snapshot.id"), snapshot.ID)
}

func Test_All_Snapshot(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/snapshots/detail")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /snapshots"))
	//mitm.Pause()

	assertion := assert.New(t)

	tenant_id := "5aa119a8-d25b-45a7-8d1b-88e127885635"
	all_tenants := "1"

	snapshots, err := New(openstacker).AllByParams(&options.ListSnapshotOpts{
		TenantID:   &tenant_id,
		AllTenants: &all_tenants,
	})
	assertion.Nil(err)
	assertion.NotNil(snapshots)
	assertion.EqualValues(1, len(snapshots))
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.id"), snapshots[0].ID)
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.status"), snapshots[0].Status)
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.created_at"), snapshots[0].CreatedAt)
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.description"), snapshots[0].Description)
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.name"), snapshots[0].Name)
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.size"), strconv.Itoa(snapshots[0].Size))
	assertion.Equal(apiv2.APIString("GET /snapshots.snapshots.0.volume_id"), snapshots[0].VolumeID)
}

func Test_Show_Snapshot(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/snapshots/"+testSnapshotId)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /snapshots/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	snapshot, err := New(openstacker).Show(testSnapshotId)
	assertion.Nil(err)
	assertion.NotNil(snapshot)

	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.id"), snapshot.ID)
	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.status"), snapshot.Status)
	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.created_at"), snapshot.CreatedAt)
	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.description"), snapshot.Description)
	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.name"), snapshot.Name)
	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.size"), strconv.Itoa(snapshot.Size))
	assertion.Equal(apiv2.APIString("GET /snapshots/:id.snapshot.volume_id"), snapshot.VolumeID)
}

func Test_Update_Snapshot(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	testSnapshotId := "4b502fcb-1f26-45f8-9fe5-3b9a0a52eaf2"

	mitm.MockRequest("PUT", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/snapshots/"+testSnapshotId)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("PUT /snapshots/:id"))
	//mitm.Pause()

	tenant_id := "803d363c5c4649b591fd74b96a8c30f5"
	all_tenants := "1"

	assertion := assert.New(t)

	snapshot, err := New(openstacker).Update(testSnapshotId, &options.UpdateSnapshotOpts{
		Name:        options.String("updated name"),
		Description: options.String("test update name"),
		TenantID:    &tenant_id,
		AllTenants:  &all_tenants,
	})
	assertion.Nil(err)
	assertion.NotNil(snapshot)
	assertion.Equal(apiv2.APIString("PUT /snapshots/:id.snapshot.id"), snapshot.ID)
}

func Test_Delete_Snapshot(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("DELETE", apiv2.MockResourceURLWithPort("8776", "/v2/"+testProjectId+"/snapshots/"+testSnapshotId)).WithResponse(http.StatusAccepted, jsonheader, apiv2.APIString("DELETE /snapshots/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(testSnapshotId)
	assertion.Nil(err)
}
