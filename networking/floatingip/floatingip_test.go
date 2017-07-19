package floatingip

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
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

	assertModel(assertion, "POST /floatingips.floatingip", ip)
}

func Test_All_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /floatingips"))
	// mitm.Pause()

	assertion := assert.New(t)

	ips, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.NotNil(ips)
	assertion.EqualValues(2, len(ips))

	assertModel(assertion, "GET /floatingips.floatingips.0", ips[0])
}

func Test_Show_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /floatingips/:id"))

	assertion := assert.New(t)

	ip, err := New(openstacker).Show(testFloatingipID)
	assertion.Nil(err)
	assertion.NotNil(ip)

	assertModel(assertion, "GET /floatingips/:id.floatingip", ip)
}

func Test_Update_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("PUT", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusOK, nil, apiv2.APIString("PUT /floatingips/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	ip, err := New(openstacker).Update(testFloatingipID, &options.UpdateFloatingIPOpts{
		Description: options.String("update floatingip desc"),
	})
	assertion.Nil(err)
	assertion.NotNil(ip)

	assertModel(assertion, "PUT /floatingips/:id.floatingip", ip)
}

func Test_Delete_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("DELETE", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusNoContent, nil, apiv2.APIString("DELETE /floatingips/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(testFloatingipID)
	assertion.Nil(err)
}

func assertModel(assertion *assert.Assertions, pathPrefix string, ip *models.FloatingIPModel) {
	assertion.Equal(apiv2.APIString(pathPrefix+".id"), ip.ID)
	assertion.Equal(apiv2.APIString(pathPrefix+".status"), ip.Status)
	assertion.Equal(apiv2.APIString(pathPrefix+".router_id"), ip.RouterID)
	assertion.Equal(apiv2.APIString(pathPrefix+".project_id"), ip.ProjectID)
	assertion.Equal(apiv2.APIString(pathPrefix+".tenant_id"), ip.TenantID)
	assertion.Equal(apiv2.APIString(pathPrefix+".port_id"), ip.PortID)
	assertion.Equal(apiv2.APIString(pathPrefix+".floating_ip_address"), ip.FloatingIP)
	assertion.Equal(apiv2.APIString(pathPrefix+".fixed_ip_address"), ip.FixedIP)
	assertion.Equal(apiv2.APIString(pathPrefix+".description"), ip.Description)
	assertion.Equal(apiv2.APIString(pathPrefix+".created_at"), ip.CreatedAt)
	assertion.Equal(apiv2.APIString(pathPrefix+".updated_at"), ip.UpdatedAt)
}
