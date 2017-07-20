package server

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_ChangeAdminPassword(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(http.MethodPost, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).WithResponse(http.StatusAccepted, jsonheader, "")
	//mitm.Pause()

	assertion := assert.New(t)

	err := NewServerManager(openstacker).ChangeAdminPassword(serverID, "newpass")
	assertion.Nil(err)
}

func Test_Start_Stop_Reboot_Shutdown_Server(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(http.MethodPost, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).WithResponse(http.StatusAccepted, jsonheader, "").AnyTimes()
	//mitm.Pause()

	assertion := assert.New(t)

	serverManager := NewServerManager(openstacker)

	err := serverManager.Start(serverID)
	assertion.Nil(err)

	err = serverManager.Stop(serverID)
	assertion.Nil(err)

	err = serverManager.Reboot(serverID)
	assertion.Nil(err)

	err = serverManager.Shutdown(serverID)
	assertion.Nil(err)

}

func Test_Resize(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers/:id"))
	//mitm.Pause()

	mitm.MockRequest(
		http.MethodPost,
		apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).
		WithResponse(
			http.StatusAccepted, jsonheader, "")
	//mitm.Pause()

	assertion := assert.New(t)

	serverManager := NewServerManager(openstacker)

	err := serverManager.Resize(serverID, "Test flavor")
	assertion.Nil(err)

}

func Test_Rebuild(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers/:id"))
	//mitm.Pause()

	mitm.MockRequest(http.MethodPost, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).WithResponse(http.StatusAccepted, jsonheader, apiv3.APIString("GET /servers/:id")).AnyTimes()
	//mitm.Pause()

	assertion := assert.New(t)

	serverManager := NewServerManager(openstacker)

	server, err := serverManager.Rebuild(serverID, "testImageId")

	assertion.Nil(err)
	assertion.NotNil(server)
}

func Test_Vnc(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(http.MethodPost, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("POST /servers/vnc"))
	//mitm.Pause()

	assertion := assert.New(t)

	serverManager := NewServerManager(openstacker)

	url, err := serverManager.Vnc(serverID)
	assertion.NotNil(url)
	assertion.Nil(err)
}
