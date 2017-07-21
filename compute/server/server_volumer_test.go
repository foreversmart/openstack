package server

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_All_Volumes(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-volume_attachments")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers/volumes"))
	//mitm.Pause()

	assertion := assert.New(t)

	volumes, err := NewServerVolumer(openstacker).All(serverID)

	assertion.Nil(err)
	assertion.NotNil(volumes)

	assertion.Equal(apiv3.APIString("GET /servers/volumes.volumeAttachments.0.volumeId"), volumes[0].VolumeID)
}

func Test_Mount_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(http.MethodPost, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-volume_attachments")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("POST /servers/volumes"))
	//mitm.Pause()

	assertion := assert.New(t)

	volume, err := NewServerVolumer(openstacker).Mount(serverID, "testVolume")

	assertion.Nil(err)
	assertion.NotNil(volume)

	assertion.Equal(apiv3.APIString("POST /servers/volumes.volumeAttachment.volumeId"), volume.VolumeID)
}

func Test_Unmount_Volume(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(http.MethodGet, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-volume_attachments")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers/volumes")).AnyTimes()
	//mitm.Pause()

	volumeID := "a26887c6-c47b-4654-abb5-dfadf7d3f803"
	mitm.MockRequest(http.MethodDelete, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-volume_attachments/"+volumeID)).WithResponse(http.StatusAccepted, jsonheader, "")
	//mitm.Pause()

	assertion := assert.New(t)

	err := NewServerVolumer(openstacker).Unmount(serverID, volumeID)

	assertion.Nil(err)
}
