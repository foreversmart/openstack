package subnet

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

const (
	networkPort = "9696"
)

func Test_Create_Subnet(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/subnets")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /subnets"))
	//mitm.Pause()

	testTenantID := apiv3.APIString("POST /subnets.subnet.project_id")
	testNetworkID := apiv3.APIString("POST /subnets.subnet.network_id")

	opts := &options.CreateSubnetOpts{
		Name:      options.String("testSubnet"),
		IPVersion: 4,
		CIDR:      options.String("172.30.248.0/22"),
		TenantID:  options.String(testTenantID),
		NetworkID: options.String(testNetworkID),
	}

	assertion := assert.New(t)
	subnet, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(testNetworkID, subnet.NetworkID)
	assertion.Equal(testTenantID, subnet.TenantID)
}

func Test_All_Subnet(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "v2.0/subnets")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /subnets"))
	// mitm.Pause()

	subnets, err := New(openstacker).AllByParams(&options.ListSubnetOpts{})
	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(1, len(subnets))
	assertion.Equal(apiv3.APIString("GET /subnets.subnets.0.id"), subnets[0].ID)
}

func Test_Show_Subnet(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	subnetID := apiv3.APIString("GET /subnets/:id.subnet.network_id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/subnets/"+subnetID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /subnets/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	subnet, err := New(openstacker).Show(subnetID)
	assertion.Nil(err)
	assertion.Equal(subnetID, subnet.NetworkID)
	assertion.Equal(apiv3.APIString("GET /subnets/:id.subnet.name"), subnet.Name)
	assertion.Equal(apiv3.APIString("GET /subnets/:id.subnet.tenant_id"), subnet.TenantID)
	assertion.Equal(apiv3.APIString("GET /subnets/:id.subnet.created_at"), subnet.CreatedAt)
	assertion.Equal(apiv3.APIString("GET /subnets/:id.subnet.cidr"), subnet.CIDR)
	assertion.Equal(apiv3.APIString("GET /subnets/:id.subnet.gateway_ip"), subnet.GatewayIP)
}

func Test_Update_Subnet(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	subnetID := apiv3.APIString("PUT /subnets/:id.subnet.network_id")

	mitm.MockRequest("PUT", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/subnets/"+subnetID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PUT /subnets/:id"))

	subnet, err := New(openstacker).Update(subnetID, &options.UpdateSubnetOpts{
		Name: options.String("updated name"),
	})

	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.NotNil(subnet)
	assertion.Equal(apiv3.APIString("PUT /subnets/:id.subnet.id"), subnet.ID)
}

func Test_Delete_Subnet(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	subnetID := apiv3.APIString("GET /subnets/:id.subnet.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/subnets/"+subnetID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /subnets/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(subnetID)
	assertion.Nil(err)
}
