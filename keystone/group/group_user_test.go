package group

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack/lib/options"
)

func Test_All_GUser(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/groups/"+groupID+"/users")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /groups/:id/users"))
	// mitm.Pause()

	assertion := assert.New(t)

	users, err := NewGroupUser(openstacker).All(groupID)
	assertion.Nil(err)
	assertion.EqualValues(1, len(users))
	assertion.Equal(apiv3.APIString("GET /groups/:id/users.users.0.id"), users[0].ID)
	assertion.Equal(apiv3.APIString("GET /groups/:id/users.users.0.name"), users[0].Name)
	assertion.Equal(apiv3.APIString("GET /groups/:id/users.users.0.email"), users[0].Email)
	assertion.Equal(apiv3.APIString("GET /groups/:id/users.users.0.domain_id"), users[0].DomainID)
	assertion.Equal(apiv3.APIString("GET /groups/:id/users.users.0.default_project_id"), users[0].DefaultProjectID)
	assertion.Empty(users[0].PasswordExpiresAt)
}

func Test_Create_GUser(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")
	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("PUT", apiv3.MockAdminURL("/v3/groups/"+groupID+"/users/"+userID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PUT /groups/:id/users/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.CreateGroupUserOpts{
		UserID: userID,
	}

	err := NewGroupUser(openstacker).Create(groupID, opts)
	assertion.Nil(err)
}

func Test_HasUser_GUser(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")
	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("HEAD", apiv3.MockAdminURL("/v3/groups/"+groupID+"/users/"+userID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("HEAD /groups/:id/users/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	ok := NewGroupUser(openstacker).HasUser(groupID, userID)
	assertion.True(ok)
}

func Test_Delete_GUser(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")
	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/groups/"+groupID+"/users/"+userID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /groups/:id/users/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := NewGroupUser(openstacker).Delete(groupID, userID)
	assertion.Nil(err)
}
