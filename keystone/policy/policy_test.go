package policy

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

func Test_All_Policy(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/policies")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /policies"))
	// mitm.Pause()

	assertion := assert.New(t)

	policies, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(policies))
	assertion.Equal(apiv3.APIString("GET /policies.policies.0.id"), policies[0].ID)
	assertion.Equal(apiv3.APIString("GET /policies.policies.0.type"), policies[0].Type)
	assertion.Equal(apiv3.APIString("GET /policies.policies.0.blob"), policies[0].Blob)
}

func Test_Create_Policy(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/policies")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /policies"))
	// mitm.Pause()

	assertion := assert.New(t)

	opts := options.CreatePolicyOpts{
		Type: "application/json",
		Blob: `{}`,
	}

	policy, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /policies.policy.id"), policy.ID)
	assertion.Equal(apiv3.APIString("POST /policies.policy.type"), policy.Type)
	assertion.Equal(apiv3.APIString("POST /policies.policy.blob"), policy.Blob)
}

func Test_Show_Policy(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	policyID := apiv3.APIString("POST /policies.policy.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/policies/"+policyID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /policies/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	policy, err := New(openstacker).Show(policyID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /policies/:id.policy.id"), policy.ID)
	assertion.Equal(apiv3.APIString("GET /policies/:id.policy.type"), policy.Type)
	assertion.Equal(apiv3.APIString("GET /policies/:id.policy.blob"), policy.Blob)
}

func Test_Update_Policy(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	policyID := apiv3.APIString("POST /policies.policy.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/policies/"+policyID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /policies/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	blob := `{"user":{"id":123321},"access_key":{"id":"ak","secret":"sk"}}`
	opts := options.UpdatePolicyOpts{
		Blob: &blob,
	}

	policy, err := New(openstacker).Update(policyID, opts)
	assertion.Nil(err)
	assertion.Equal(blob, policy.Blob)
}

func Test_Delete_Policy(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	policyID := apiv3.APIString("POST /policies.policy.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/policies/"+policyID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /policies/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(policyID)
	assertion.Nil(err)
}
