package port

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

const (
	networkPort = "9696"
)

func Test_Create_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/ports")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /ports"))
	//mitm.Pause()

	testNetworkID := apiv3.APIString("GET /networks/:id.network.id")
	opts := &options.CreatePortOpts{
		NetworkID: options.String(testNetworkID),
	}

	assertion := assert.New(t)
	port, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /ports.port.device_id"), port.DeviceID)
	assertion.Equal(apiv3.APIString("POST /ports.port.device_owner"), port.DeviceOwner)
}

func Test_All_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "v2.0/ports")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /ports"))
	// mitm.Pause()

	ports, err := New(openstacker).AllByParams(&options.ListPortOpts{})
	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(1, len(ports))
	assertion.Equal(apiv3.APIString("GET /ports.ports.0.id"), ports[0].ID)
}

func Test_Show_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	portID := apiv3.APIString("GET /ports/:id.port.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/ports/"+portID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /ports/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	port, err := New(openstacker).Show(portID)
	assertion.Nil(err)
	assertion.Equal(portID, port.ID)
	assertion.True(port.AdminStateUp)

	assertion.Equal(apiv3.APIString("GET /ports/:id.port.created_at"), port.CreatedAt)
	assertion.Equal(apiv3.APIString("GET /ports/:id.port.device_id"), port.DeviceID)
	assertion.Equal(apiv3.APIString("GET /ports/:id.port.device_owner"), port.DeviceOwner)
	assertion.Equal(apiv3.APIString("GET /ports/:id.port.network_id"), port.NetworkID)
	assertion.Equal(apiv3.APIString("GET /ports/:id.port.project_id"), port.ProjectID)
	assertion.Equal(apiv3.APIString("GET /ports/:id.port.status"), port.Status)
}

func Test_Delete_Port(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	portID := apiv3.APIString("GET /ports/:id.port.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/ports/"+portID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /ports/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(portID)
	assertion.Nil(err)
}
