package user

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	uuid "github.com/satori/go.uuid"
)

func Test_All_User(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/users")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /users"))
	// mitm.Pause()

	assertion := assert.New(t)

	users, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(users))
	assertion.Equal(apiv3.APIString("GET /users.users.0.id"), users[0].ID)
	assertion.Equal(apiv3.APIString("GET /users.users.0.name"), users[0].Name)
	assertion.Equal(apiv3.APIString("GET /users.users.0.email"), users[0].Email)
	assertion.Empty(users[0].Options)
	assertion.True(users[0].Enabled)
	assertion.Empty(users[0].PasswordExpiresAt)
	assertion.Equal(apiv3.APIString("GET /users.users.0.domain_id"), users[0].DomainID)
	assertion.Equal(apiv3.APIString("GET /users.users.0.default_project_id"), users[0].DefaultProjectID)
}

func Test_Create_User(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/users")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /users"))
	// mitm.Pause()

	assertion := assert.New(t)

	domainid := "default"
	enable := true
	opts := options.CreateUserOpts{
		DomainID: &domainid,
		Name:     "testing kirk",
		Enabled:  &enable,
	}

	user, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /users.user.id"), user.ID)
	assertion.Equal(opts.Name, user.Name)
	assertion.Empty(user.Email)
	assertion.True(user.Enabled)
	assertion.Equal(*opts.DomainID, user.DomainID)
	assertion.Empty(user.DefaultProjectID)
	assertion.Empty(user.PasswordExpiresAt)
}

func Test_Show_User(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/users/"+userID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /users/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	user, err := New(openstacker).Show(userID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /users/:id.user.id"), user.ID)
	assertion.Equal(apiv3.APIString("GET /users/:id.user.name"), user.Name)
	// assertion.Equal(apiv3.APIString("GET /users/:id.user.email"), user.Email)
	assertion.Equal(apiv3.APIString("GET /users/:id.user.domain_id"), user.DomainID)
	// assertion.Equal(apiv3.APIString("GET /users/:id.user.default_project_id"), user.DefaultProjectID)
	assertion.Empty(user.PasswordExpiresAt)
}

func Test_Update_User(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/users/"+userID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /users/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	email := "kirk@api.testing"
	projectID := apiv3.APIString("POST /projects.project.id")
	opts := options.UpdateUserOpts{
		Email:            &email,
		DefaultProjectID: &projectID,
	}

	user, err := New(openstacker).Update(userID, opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("PATCH /users/:id.user.id"), user.ID)
	assertion.Equal(apiv3.APIString("PATCH /users/:id.user.name"), user.Name)
	assertion.Equal(apiv3.APIString("PATCH /users/:id.user.email"), user.Email)
	assertion.Equal(apiv3.APIString("PATCH /users/:id.user.extra.email"), user.Extra["email"])
	assertion.Equal(apiv3.APIString("PATCH /users/:id.user.domain_id"), user.DomainID)
	assertion.Equal(apiv3.APIString("PATCH /users/:id.user.default_project_id"), user.DefaultProjectID)
	assertion.Empty(user.PasswordExpiresAt)
}

func Test_ChangePassword_User(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/users/"+userID+"/password")).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("POST /users/:password"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.ChangeUserPasswordOpts{
		OriginalPassword: "UR6hyDfWnAbo95T$dN",
		Password:         uuid.NewV4().String(),
	}

	err := New(openstacker).ChangePasswd(userID, opts)
	assertion.Nil(err)
}

func Test_Delete_User(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/users/"+userID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /users/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(userID)
	assertion.Nil(err)
}

func Test_All_UserGroups(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/users/"+userID+"/groups")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /users/:id/groups"))
	// mitm.Pause()

	assertion := assert.New(t)

	groups, err := New(openstacker).AllGroups(userID)
	assertion.Nil(err)
	assertion.EqualValues(1, len(groups))
	assertion.Equal(apiv3.APIString("GET /users/:id/groups.groups.0.id"), groups[0].ID)
	assertion.Equal(apiv3.APIString("GET /users/:id/groups.groups.0.name"), groups[0].Name)
	assertion.Equal(apiv3.APIString("GET /users/:id/groups.groups.0.description"), groups[0].Description)
	assertion.Equal(apiv3.APIString("GET /users/:id/groups.groups.0.domain_id"), groups[0].DomainID)
}

func Test_All_UserProjects(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	userID := apiv3.APIString("POST /users.user.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/users/"+userID+"/projects")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /users/:id/projects"))
	// mitm.Pause()

	assertion := assert.New(t)

	projects, err := New(openstacker).AllProjects(userID)
	assertion.Nil(err)
	assertion.EqualValues(1, len(projects))
	assertion.Equal(apiv3.APIString("GET /users/:id/projects.projects.0.id"), projects[0].ID)
	assertion.Equal(apiv3.APIString("GET /users/:id/projects.projects.0.name"), projects[0].Name)
	assertion.Equal(apiv3.APIString("GET /users/:id/projects.projects.0.description"), projects[0].Description)
	assertion.Equal(apiv3.APIString("GET /users/:id/projects.projects.0.parent_id"), projects[0].ParentID)
	assertion.Equal(apiv3.APIString("GET /users/:id/projects.projects.0.domain_id"), projects[0].DomainID)
	assertion.False(projects[0].IsDomain)
	assertion.True(projects[0].Enabled)
}
