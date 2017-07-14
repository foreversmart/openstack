package endpoint

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/enums"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

func Test_All_Endpoint(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/endpoints")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /endpoints"))
	// mitm.Pause()

	assertion := assert.New(t)

	endpoints, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(endpoints))
	assertion.Equal(apiv3.APIString("GET /endpoints.endpoints.0.id"), endpoints[0].ID)
	assertion.Equal(apiv3.APIString("GET /endpoints.endpoints.0.url"), endpoints[0].Url)
	assertion.EqualValues(apiv3.APIString("GET /endpoints.endpoints.0.interface"), endpoints[0].Interface)
	assertion.Equal(apiv3.APIString("GET /endpoints.endpoints.0.region_id"), endpoints[0].RegionID)
	assertion.Equal(apiv3.APIString("GET /endpoints.endpoints.0.region"), endpoints[0].Region)
	assertion.Equal(apiv3.APIString("GET /endpoints.endpoints.0.service_id"), endpoints[0].ServiceID)
	assertion.True(endpoints[0].Enabled)
}

func Test_Create_Endpoint(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/endpoints")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /endpoints"))
	// mitm.Pause()

	assertion := assert.New(t)

	regionID := "RegionOne"
	enabled := false
	opts := options.CreateEndpointOpts{
		Url:       "http://api.ecloud.com:8989/kirk",
		Interface: enums.EndpointInterfacePublic,
		Enabled:   &enabled,
		RegionID:  &regionID,
		ServiceID: apiv3.APIString("POST /services.service.id"),
	}

	endpoint, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /endpoints.endpoint.id"), endpoint.ID)
	assertion.Equal(opts.Url, endpoint.Url)
	assertion.Equal(opts.Interface, endpoint.Interface)
	assertion.Equal(regionID, endpoint.RegionID)
	assertion.Equal(opts.ServiceID, endpoint.ServiceID)
	assertion.False(endpoint.Enabled)
}

func Test_Show_Endpoint(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	endpointID := apiv3.APIString("GET /endpoints/:id.endpoint.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/endpoints/"+endpointID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /endpoints/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	endpoint, err := New(openstacker).Show(endpointID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /endpoints/:id.endpoint.url"), endpoint.Url)
	assertion.EqualValues(apiv3.APIString("GET /endpoints/:id.endpoint.interface"), endpoint.Interface)
	assertion.Equal(apiv3.APIString("GET /endpoints/:id.endpoint.region_id"), endpoint.RegionID)
	assertion.Equal(apiv3.APIString("GET /endpoints/:id.endpoint.region"), endpoint.Region)
	assertion.Equal(apiv3.APIString("GET /endpoints/:id.endpoint.service_id"), endpoint.ServiceID)
	assertion.False(endpoint.Enabled)
}

func Test_Update_Endpoint(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	endpointID := apiv3.APIString("GET /endpoints/:id.endpoint.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/endpoints/"+endpointID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /endpoints/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	enabled := true
	opts := options.UpdateEndpointOpts{
		Enabled: &enabled,
	}

	endpoint, err := New(openstacker).Update(endpointID, opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("PATCH /endpoints/:id.endpoint.url"), endpoint.Url)
	assertion.EqualValues(apiv3.APIString("PATCH /endpoints/:id.endpoint.interface"), endpoint.Interface)
	assertion.Equal(apiv3.APIString("PATCH /endpoints/:id.endpoint.region_id"), endpoint.RegionID)
	assertion.Equal(apiv3.APIString("PATCH /endpoints/:id.endpoint.region"), endpoint.Region)
	assertion.Equal(apiv3.APIString("PATCH /endpoints/:id.endpoint.service_id"), endpoint.ServiceID)
	assertion.True(endpoint.Enabled)
}

func Test_Delete_Endpoint(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	endpointID := apiv3.APIString("GET /endpoints/:id.endpoint.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/endpoints/"+endpointID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /endpoints/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(endpointID)
	assertion.Nil(err)
}
