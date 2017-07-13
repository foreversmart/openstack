package servers

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_All_Servers(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/servers/detail")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /servers"))
	//mitm.Pause()

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
