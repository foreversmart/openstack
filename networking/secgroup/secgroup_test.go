package secgroup

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

func Test_Create_SecurityGroup(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)
	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort("9696", "/v2.0/security-groups")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /security-groups"))
	// mitm.Pause()

	assertion := assert.New(t)

	name := "new-webservers"
	projectID := apiv3.APIString("POST /security-groups.security_group.project_id")
	opts := options.CreateSecurityGroupOpts{
		Name:      &name,
		TenantID:  &projectID,
		ProjectID: &projectID,
	}

	securityGroup, err := New(openstacker).Create(&opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /security-groups.security_group.id"), securityGroup.ID)
	assertion.Equal(*opts.Name, securityGroup.Name)
}

func Test_All_SecurityGroup(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)
	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort("9696", "/v2.0/security-groups")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /security-groups"))
	// mitm.Pause()

	assertion := assert.New(t)

	securityGroups, err := New(openstacker).All()
	assertion.Nil(err)

	assertion.Equal(apiv3.APIString("GET /security-groups.security_groups.0.id"), securityGroups[0].ID)
	assertion.Equal(apiv3.APIString("GET /security-groups.security_groups.0.name"), securityGroups[0].Name)
}

func Test_Show_SecurityGroup(t *testing.T) {
	securityGroupID := apiv3.APIString("GET /security-groups/:id.security_group.id")

	mitm := mocker.StubDefaultTransport(t)
	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort("9696", "/v2.0/security-groups/"+securityGroupID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /security-groups/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	securityGroup, err := New(openstacker).Show(securityGroupID, nil)

	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /security-groups/:id.security_group.id"), securityGroup.ID)
	assertion.Equal(apiv3.APIString("GET /security-groups/:id.security_group.name"), securityGroup.Name)

}

func Test_Update_SecurityGroup(t *testing.T) {
	securityGroupID := apiv3.APIString("PUT /security-groups/:id.security_group.id")

	mitm := mocker.StubDefaultTransport(t)
	mitm.MockRequest("PUT", apiv3.MockResourceURLWithPort("9696", "/v2.0/security-groups/"+securityGroupID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PUT /security-groups/:id"))
	// mitm.Pause()
	assertion := assert.New(t)

	securityGroup, err := New(openstacker).Update(securityGroupID, &options.UpdateSecurityGroupOpts{
		Name: options.String("my security group"),
	})

	assertion.Nil(err)
	assertion.NotNil(securityGroup)
	assertion.Equal(apiv3.APIString("PUT /security-groups/:id.security_group.name"), securityGroup.Name)

}

func Test_Delete_SecurityGroup(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	securityGroupID := apiv3.APIString("GET /security-groups/:id.security_group.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort("9696", "/v2.0/security-groups/"+securityGroupID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /security-groups/:id"))
	// mitm.Pause()
	assertion := assert.New(t)

	err := New(openstacker).Delete(securityGroupID)
	assertion.Nil(err)

}
