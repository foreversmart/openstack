package role

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_All_RoleAssignment(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/role_assignments")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /role_assignments"))
	// mitm.Pause()

	assertion := assert.New(t)

	roleAssignments, err := NewAssignment(openstacker).All()
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /role_assignments.role_assignments.0.user.id"), roleAssignments[0].User.ID)
}
