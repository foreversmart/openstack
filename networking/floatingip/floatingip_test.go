package floatingip

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

var (
	testFloatingipID string
)

const (
	networkPort = "9696"
)

func Test_Create_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips")).WithResponse(http.StatusCreated, jsonheader, apiv2.APIString("POST /floatingips"))
	//mitm.Pause()

	assertion := assert.New(t)
	networkID := apiv3.APIString("GET /networks/:id.network.id")

	ip, err := New(openstacker).Create(&options.CreateFloatingIPOpts{
		FloatingNetworkID: options.String(networkID),
		TenantID:          options.String(testProjectID),
		ProjectID:         options.String(testProjectID),
		Description:       options.String("test create floatingip"),
	})
	testFloatingipID = ip.ID

	assertion.Nil(err)
	assertion.NotNil(ip)
}
