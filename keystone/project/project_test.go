package project

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack/lib/options"
)

func Test_All_Project(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/projects")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /projects"))
	// mitm.Pause()

	assertion := assert.New(t)

	projects, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(projects))
	assertion.Equal(apiv3.APIString("GET /projects.projects.0.id"), projects[0].ID)
	assertion.Equal(apiv3.APIString("GET /projects.projects.0.name"), projects[0].Name)
	assertion.Equal(apiv3.APIString("GET /projects.projects.0.domain_id"), projects[0].DomainID)
	assertion.Equal(apiv3.APIString("GET /projects.projects.0.parent_id"), projects[0].ParentID)
	assertion.False(projects[0].IsDomain)
	assertion.True(projects[0].Enabled)
}

func Test_Create_Project(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/projects")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /projects"))
	// mitm.Pause()

	assertion := assert.New(t)

	domainID := "default"
	isDomain := false
	isEnabled := true
	opts := options.CreateProjectOpts{
		DomainID: &domainID,
		Name:     "testing project",
		Enabled:  &isEnabled,
		IsDomain: &isDomain,
	}

	project, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(opts.Name, project.Name)
	assertion.Equal(domainID, project.DomainID)
	assertion.Equal(domainID, project.ParentID)
	assertion.False(project.IsDomain)
	assertion.True(project.Enabled)
}

func Test_Show_Project(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	projectID := apiv3.APIString("POST /projects.project.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/projects/"+projectID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /projects/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	project, err := New(openstacker).Show(projectID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /projects/:id.project.name"), project.Name)
	assertion.Equal(apiv3.APIString("GET /projects/:id.project.domain_id"), project.DomainID)
	assertion.Equal(apiv3.APIString("GET /projects/:id.project.parent_id"), project.ParentID)
	assertion.False(project.IsDomain)
	assertion.True(project.Enabled)
}

func Test_Update_Project(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	projectID := apiv3.APIString("POST /projects.project.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/projects/"+projectID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /projects/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	isEnabled := false
	opts := options.UpdateProjectOpts{
		Enabled: &isEnabled,
	}

	project, err := New(openstacker).Update(projectID, opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("PATCH /projects/:id.project.name"), project.Name)
	assertion.Equal(apiv3.APIString("PATCH /projects/:id.project.domain_id"), project.DomainID)
	assertion.Equal(apiv3.APIString("PATCH /projects/:id.project.parent_id"), project.ParentID)
	assertion.False(project.IsDomain)
	assertion.False(project.Enabled)
}

func Test_Delete_Project(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	projectID := apiv3.APIString("POST /projects.project.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/projects/"+projectID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /projects/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(projectID)
	assertion.Nil(err)
}
