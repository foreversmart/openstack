package server

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_All_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-interface")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers/ports"))
	//mitm.Pause()

	assertion := assert.New(t)

	ports, err := NewServerPorter(openstacker).All(serverID)

	assertion.Nil(err)
	assertion.NotNil(ports)

	assertion.Equal(apiv3.APIString("GET /servers/ports.interfaceAttachments.0.port_id"), ports[0].PortId)
}

func Test_Bind_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(http.MethodPost, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-interface")).WithResponse(http.StatusOK, jsonheader, "{}")
	//mitm.Pause()

	assertion := assert.New(t)

	err := NewServerPorter(openstacker).Bind(serverID, "testport")

	assertion.Nil(err)
}

func Test_Unbind_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	portID := "testport"
	mitm.MockRequest(http.MethodDelete, apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/os-interface/"+portID)).WithResponse(http.StatusAccepted, jsonheader, "{}")
	//mitm.Pause()

	assertion := assert.New(t)

	err := NewServerPorter(openstacker).Unbind(serverID, portID)

	assertion.Nil(err)
}
