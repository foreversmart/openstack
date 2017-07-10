package domain

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

func Test_All_Domain(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/domains")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /domains"))
	// mitm.Pause()

	assertion := assert.New(t)

	domains, err := New(openstacker).All(options.ListDomainOpts{})
	assertion.Nil(err)
	assertion.EqualValues(1, len(domains))
	assertion.Equal(apiv3.APIString("GET /domains.domains.0.id"), domains[0].ID)
	assertion.True(domains[0].Enabled)
}

func Test_Create_Domain(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/domains")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /domains"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.CreateDomainOpts{
		Name:        "testing domain",
		Description: "The domain for tesing",
		Enabled:     true,
	}

	domain, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /domains.domain.id"), domain.ID)
	assertion.Equal(apiv3.APIString("POST /domains.domain.name"), domain.Name)
	assertion.True(domain.Enabled)
}

func Test_Show_Domain(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	domainID := apiv3.APIString("GET /domains.domains.0.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/domains/"+domainID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /domains/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	domain, err := New(openstacker).Show(domainID)
	assertion.Nil(err)
	assertion.Equal(domainID, domain.ID)
	assertion.True(domain.Enabled)
}

func Test_Update_Domain(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	domainID := apiv3.APIString("PATCH /domains/:id.domain.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/domains/"+domainID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /domains/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	enabled := false
	opts := options.UpdateDomainOpts{
		Enabled: &enabled,
	}

	domain, err := New(openstacker).Update(domainID, opts)
	assertion.Nil(err)
	assertion.False(domain.Enabled)
}

func Test_Delete_Domain(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	domainID := apiv3.APIString("PATCH /domains/:id.domain.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/domains/"+domainID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /domains/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(domainID)
	assertion.Nil(err)
}
