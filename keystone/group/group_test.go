package group

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack/lib/options"
)

func Test_All_Group(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/groups")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /groups"))
	// mitm.Pause()

	assertion := assert.New(t)

	groups, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(groups))
	assertion.Equal(apiv3.APIString("GET /groups.groups.0.id"), groups[0].ID)
	assertion.Equal(apiv3.APIString("GET /groups.groups.0.name"), groups[0].Name)
	assertion.Equal(apiv3.APIString("GET /groups.groups.0.description"), groups[0].Description)
	assertion.Equal(apiv3.APIString("GET /groups.groups.0.domain_id"), groups[0].DomainID)
}

func Test_Create_Group(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/groups")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /groups"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.CreateGroupOpts{
		Name: "testing group",
	}

	group, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /groups.group.id"), group.ID)
	assertion.Equal(apiv3.APIString("POST /groups.group.name"), group.Name)
	assertion.Equal(apiv3.APIString("POST /groups.group.description"), group.Description)
	assertion.Equal(apiv3.APIString("POST /groups.group.domain_id"), group.DomainID)
}

func Test_Show_Group(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/groups/"+groupID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /groups/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	group, err := New(openstacker).Show(groupID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.id"), group.ID)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.name"), group.Name)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.description"), group.Description)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.domain_id"), group.DomainID)
}

func Test_Update_Group(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/groups/"+groupID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /groups/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	desc := "testing group description"
	opts := options.UpdateGroupOpts{
		Description: &desc,
	}

	group, err := New(openstacker).Update(groupID, opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.id"), group.ID)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.name"), group.Name)
	assertion.Equal(desc, group.Description)
	assertion.Equal(apiv3.APIString("GET /groups/:id.group.domain_id"), group.DomainID)
}

func Test_Delete_Group(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	groupID := apiv3.APIString("POST /groups.group.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/groups/"+groupID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /groups/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(groupID)
	assertion.Nil(err)
}
