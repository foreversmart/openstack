package network

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

const (
	networkPort = "9696"
)

func Test_Create_Network(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/networks")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /networks"))
	// mitm.Pause()

	opts := &options.CreateNetworkOpts{
		Name:     options.String("TestNetwork"),
		TenantID: options.String(apiv3.GetString("user.project_id")),
	}
	assertion := assert.New(t)
	network, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /networks.network.id"), network.ID)
	assertion.True(network.AdminStateUp)
}

func Test_All_Network(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "v2.0/networks")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /networks")).AnyTimes()
	// mitm.Pause()

	networks, err := New(openstacker).All()
	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(1, len(networks))
	assertion.Equal(apiv3.APIString("GET /networks.networks.0.id"), networks[0].ID)
}

func Test_AllByParams_Network(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "v2.0/networks")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /networks"))
	// mitm.Pause()

	id := apiv3.APIString("GET /networks.networks.0.tenant_id")
	networks, err := New(openstacker).AllByParams(&options.ListNetworkOpt{
		ProjectId: &id,
	})

	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(1, len(networks))
	assertion.Equal(apiv3.APIString("GET /networks.networks.0.id"), networks[0].ID)
}

func Test_Show_Network(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	networkID := apiv3.APIString("GET /networks/:id.network.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/networks/"+networkID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /networks/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	network, err := New(openstacker).Show(networkID)
	assertion.Nil(err)
	assertion.Equal(networkID, network.ID)
	assertion.Equal(apiv3.APIString("GET /networks/:id.network.name"), network.Name)
	assertion.Equal(apiv3.APIString("GET /networks/:id.network.tenant_id"), network.TenantID)
	assertion.True(network.AdminStateUp)
	assertion.False(network.Shared)
}

func Test_Delete_Network(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	networkID := apiv3.APIString("GET /networks/:id.network.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/networks/"+networkID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /networks/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(networkID)
	assertion.Nil(err)
}
