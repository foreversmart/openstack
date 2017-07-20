package secrule

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

const (
	networkPort = "9696"
)

func Test_Create_Secrule(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/security-group-rules")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /secrules"))
	//mitm.Pause()

	opts := &options.CreateSecruleOpts{
		SecGroupID:     options.String(apiv3.APIString("GET /secrules/:id.security_group_rule.security_group_id")),
		Direction:      options.String(apiv3.APIString("GET /secrules/:id.security_group_rule.direction")),
		RemoteIPPrefix: options.String(apiv3.APIString("GET /secrules/:id.security_group_rule.remote_ip_prefix")),
	}

	secrule, err := New(openstacker).Create(opts)
	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /secrules.security_group_rule.direction"), secrule.Direction)
	assertion.Equal(apiv3.APIString("POST /secrules.security_group_rule.security_group_id"), secrule.SecGroupID)
}

func Test_All_SecRules(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "v2.0/security-group-rules")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /secrules"))
	// mitm.Pause()

	secrules, err := New(openstacker).AllByParams(&options.ListSecRuleOpts{})
	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(1, len(secrules))
	assertion.Equal(apiv3.APIString("GET /secrules.security_group_rules.0.id"), secrules[0].ID)
}

func Test_Show_SecRule(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	secruleID := apiv3.APIString("GET /secrules/:id.security_group_rule.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/security-group-rules/"+secruleID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /secrules/:id"))
	// mitm.Pause()

	secrule, err := New(openstacker).Show(secruleID)
	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /secrules/:id.security_group_rule.id"), secrule.ID)
	assertion.Equal(apiv3.APIString("GET /secrules/:id.security_group_rule.direction"), secrule.Direction)
	assertion.Equal(apiv3.APIString("GET /secrules/:id.security_group_rule.ethertype"), secrule.EtherType)
	assertion.Equal(apiv3.APIString("GET /secrules/:id.security_group_rule.security_group_id"), secrule.SecGroupID)
	assertion.Equal(apiv3.APIString("GET /secrules/:id.security_group_rule.project_id"), secrule.ProjectID)
}

func Test_Delete_SecRule(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	secruleID := apiv3.APIString("GET /secrules/:id.security_group_rule.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/security-group-rules/"+secruleID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /secrules/:id"))
	//mitm.Pause()

	err := New(openstacker).Delete(secruleID)
	assertion := assert.New(t)
	assertion.Nil(err)
}
