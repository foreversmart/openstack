package service

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack/lib/options"
)

func Test_All_Service(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/services")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /services"))
	// mitm.Pause()

	assertion := assert.New(t)

	services, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(services))
	assertion.Equal(apiv3.APIString("GET /services.services.0.id"), services[0].ID)
	assertion.Equal(apiv3.APIString("GET /services.services.0.name"), services[0].Name)
	assertion.Equal(apiv3.APIString("GET /services.services.0.type"), services[0].Type)
	assertion.Equal(apiv3.APIString("GET /services.services.0.description"), services[0].Desc)
	assertion.True(services[0].Enabled)
}

func Test_Create_Service(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/services")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /services"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.CreateServiceOpts{
		Name:    "testing service",
		Type:    "kirk",
		Desc:    "Kirk custom service",
		Enabled: false,
	}

	service, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /services.service.id"), service.ID)
	assertion.Equal(opts.Name, service.Name)
	assertion.Equal(opts.Type, service.Type)
	assertion.False(service.Enabled)
}

func Test_Show_Service(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serviceID := apiv3.APIString("POST /services.service.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/services/"+serviceID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /services/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	service, err := New(openstacker).Show(serviceID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /services/:id.service.id"), service.ID)
	assertion.Equal(apiv3.APIString("GET /services/:id.service.name"), service.Name)
	assertion.Equal(apiv3.APIString("GET /services/:id.service.type"), service.Type)
	assertion.False(service.Enabled)
}

func Test_Update_Service(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serviceID := apiv3.APIString("POST /services.service.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/services/"+serviceID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /services/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	enabled := true
	opts := options.UpdateServiceOpts{
		Name:    "testing service",
		Type:    "kirk",
		Enabled: &enabled,
	}

	service, err := New(openstacker).Update(serviceID, opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("PATCH /services/:id.service.id"), service.ID)
	assertion.Equal(apiv3.APIString("PATCH /services/:id.service.name"), service.Name)
	assertion.Equal(apiv3.APIString("PATCH /services/:id.service.type"), service.Type)
	assertion.True(service.Enabled)
}

func Test_Delete_Service(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serviceID := apiv3.APIString("POST /services.service.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/services/"+serviceID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /services/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(serviceID)
	assertion.Nil(err)
}
