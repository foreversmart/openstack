package role

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

func Test_All_Role(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/roles")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /roles"))
	// mitm.Pause()

	assertion := assert.New(t)

	roles, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(3, len(roles))
	assertion.Equal(apiv3.APIString("GET /roles.roles.2.id"), roles[2].ID)
	assertion.Equal(apiv3.APIString("GET /roles.roles.2.name"), roles[2].Name)
	assertion.Empty(roles[2].DomainID)
}

func Test_Create_Role(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/roles")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /roles"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.CreateRoleOpts{
		Name: "KirkAdmin",
	}

	role, err := New(openstacker).Create(&opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /roles.role.id"), role.ID)
	assertion.Equal(opts.Name, role.Name)
	assertion.Empty(role.DomainID)
}

func Test_Show_Role(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	roleID := apiv3.APIString("POST /roles.role.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/roles/"+roleID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /roles/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	role, err := New(openstacker).Show(roleID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /roles/:id.role.id"), role.ID)
	assertion.Equal(apiv3.APIString("GET /roles/:id.role.name"), role.Name)
	assertion.Empty(role.DomainID)
}

func Test_Update_Role(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	roleID := apiv3.APIString("POST /roles.role.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/roles/"+roleID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /roles/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	name := "TestingAdmin"
	opts := options.UpdateRoleOpts{
		Name: &name,
	}

	role, err := New(openstacker).Update(roleID, &opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("PATCH /roles/:id.role.id"), role.ID)
	assertion.Equal(name, role.Name)
}

func Test_Delete_Role(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	roleID := apiv3.APIString("POST /roles.role.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/roles/"+roleID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /roles/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(roleID)
	assertion.Nil(err)
}
