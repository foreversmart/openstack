package role

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_All_AbstractRole(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/projects/"+abstractResourceID+"/users/"+abstractTargetID+"/roles")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /projects/:id/users/:id/roles"))
	// mitm.Pause()

	assertion := assert.New(t)

	roles, err := NewProjectUserRole(openstacker).All(abstractResourceID, abstractTargetID)
	assertion.Nil(err)
	assertion.EqualValues(1, len(roles))
	assertion.Equal(apiv3.APIString("GET /projects/:id/users/:id/roles.roles.0.id"), roles[0].ID)
	assertion.Equal(apiv3.APIString("GET /projects/:id/users/:id/roles.roles.0.name"), roles[0].Name)
	assertion.Empty(roles[0].DomainID)
}

func Test_Create_AbstractRole(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("PUT", apiv3.MockAdminURL("/v3/projects/"+abstractResourceID+"/users/"+abstractTargetID+"/roles/"+abstractRoleID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("PUT /projects/:id/users/:id/roles/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	err := NewProjectUserRole(openstacker).Create(abstractResourceID, abstractTargetID, abstractRoleID)
	assertion.Nil(err)
}

func Test_Has_AbstractRole(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("HEAD", apiv3.MockAdminURL("/v3/projects/"+abstractResourceID+"/users/"+abstractTargetID+"/roles/"+abstractRoleID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("HEAD /projects/:id/users/:id/roles/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	ok := NewProjectUserRole(openstacker).HasRole(abstractResourceID, abstractTargetID, abstractRoleID)
	assertion.True(ok)
}

func Test_Delete_AbstractRole(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/projects/"+abstractResourceID+"/users/"+abstractTargetID+"/roles/"+abstractRoleID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /projects/:id/users/:id/roles/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	err := NewProjectUserRole(openstacker).Delete(abstractResourceID, abstractTargetID, abstractRoleID)
	assertion.Nil(err)
}
