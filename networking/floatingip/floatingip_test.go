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

	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.id"), ip.ID)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.status"), ip.Status)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.router_id"), ip.RouterID)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.project_id"), ip.ProjectID)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.tenant_id"), ip.TenantID)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.port_id"), ip.PortID)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.floating_ip_address"), ip.FloatingIP)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.fixed_ip_address"), ip.FixedIP)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.description"), ip.Description)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.created_at"), ip.CreatedAt)
	assertion.Equal(apiv2.APIString("POST /floatingips.floatingip.updated_at"), ip.UpdatedAt)
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
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.id"), ips[0].ID)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.status"), ips[0].Status)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.router_id"), ips[0].RouterID)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.project_id"), ips[0].ProjectID)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.tenant_id"), ips[0].TenantID)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.port_id"), ips[0].PortID)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.floating_ip_address"), ips[0].FloatingIP)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.fixed_ip_address"), ips[0].FixedIP)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.description"), ips[0].Description)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.created_at"), ips[0].CreatedAt)
	assertion.Equal(apiv2.APIString("GET /floatingips.floatingips.0.updated_at"), ips[0].UpdatedAt)
}

func Test_Show_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /floatingips/:id"))

	assertion := assert.New(t)

	ip, err := New(openstacker).Show(testFloatingipID)
	assertion.Nil(err)
	assertion.NotNil(ip)

	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.id"), ip.ID)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.status"), ip.Status)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.router_id"), ip.RouterID)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.project_id"), ip.ProjectID)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.tenant_id"), ip.TenantID)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.port_id"), ip.PortID)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.floating_ip_address"), ip.FloatingIP)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.fixed_ip_address"), ip.FixedIP)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.description"), ip.Description)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.created_at"), ip.CreatedAt)
	assertion.Equal(apiv2.APIString("GET /floatingips/:id.floatingip.updated_at"), ip.UpdatedAt)
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

	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.id"), ip.ID)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.status"), ip.Status)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.router_id"), ip.RouterID)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.project_id"), ip.ProjectID)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.tenant_id"), ip.TenantID)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.port_id"), ip.PortID)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.floating_ip_address"), ip.FloatingIP)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.fixed_ip_address"), ip.FixedIP)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.description"), ip.Description)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.created_at"), ip.CreatedAt)
	assertion.Equal(apiv2.APIString("PUT /floatingips/:id.floatingip.updated_at"), ip.UpdatedAt)
}

func Test_Delete_FloatingIP(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("DELETE", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusNoContent, nil, apiv2.APIString("DELETE /floatingips/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(testFloatingipID)
	assertion.Nil(err)
}
