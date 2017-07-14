package servers

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

const computerPort = "8774"

func Test_All_Servers(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/detail?")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers"))
	// mitm.Pause()

	assertion := assert.New(t)

	servers, err := New(openstacker).All()

	assertion.Nil(err)

	assertion.EqualValues(1, len(servers))
	assertion.Equal(apiv3.APIString("GET /servers.servers.0.id"), servers[0].ID)
	assertion.Equal(apiv3.APIString("GET /servers.servers.0.name"), servers[0].Name)
	assertion.Equal(apiv3.APIString("GET /servers.servers.0.user_id"), servers[0].UserID)
	assertion.Equal(apiv3.APIString("GET /servers.servers.0.status"), servers[0].Status)
	assertion.Equal(apiv3.APIString("GET /servers.servers.0.hostId"), servers[0].HostID)
}

func Test_Show_Servers(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	server, err := New(openstacker).Show(serverID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /servers/:id.server.id"), server.ID)
	assertion.Equal(apiv3.APIString("GET /servers/:id.server.name"), server.Name)
	assertion.Equal(apiv3.APIString("GET /servers/:id.server.user_id"), server.UserID)
	assertion.Equal(apiv3.APIString("GET /servers/:id.server.status"), server.Status)
	assertion.Equal(apiv3.APIString("GET /servers/:id.server.hostId"), server.HostID)
}

func Test_Update_Servers(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("PUT", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PUT /servers/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	name := "TestModifyServer"
	ipv4 := "172.24.4.14"
	config := "AUTO"
	ipv6 := "80fe::"
	opts := options.UpdateServersOpts{
		Name:            &name,
		AccessIPv4:      &ipv4,
		AccessIPv6:      &ipv6,
		OSDcfDiskConfig: &config,
	}

	server, err := New(openstacker).Update(serverID, opts)
	assertion.Nil(err)
	assertion.Equal(name, server.Name)
	assertion.Equal(ipv4, server.AccessIPv4)
	assertion.Equal(ipv6, server.AccessIPv6)
	assertion.Equal(config, server.OSDcfDiskConfig)
}

func Test_Delete_Servers(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /servers/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(serverID)
	assertion.Nil(err)
}
