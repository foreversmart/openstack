package credential

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack/lib/options"
)

func Test_All_Credential(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/credentials")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /credentials"))
	// mitm.Pause()

	assertion := assert.New(t)

	credentials, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(credentials))
	assertion.Equal(apiv3.APIString("GET /credentials.credentials.0.id"), credentials[0].ID)
	assertion.Equal(apiv3.APIString("GET /credentials.credentials.0.project_id"), credentials[0].ProjectID)
	assertion.Equal(apiv3.APIString("GET /credentials.credentials.0.user_id"), credentials[0].UserID)
	assertion.Equal(apiv3.APIString("GET /credentials.credentials.0.type"), credentials[0].Type)
	assertion.Empty(apiv3.APIString("GET /credentials.credentials.0.blob"), credentials[0].Blob)
}

func Test_Create_Credential(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/credentials")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /credentials"))
	// mitm.Pause()

	assertion := assert.New(t)

	projectID := apiv3.APIString("POST /projects.project.id")
	userID := apiv3.APIString("POST /users.user.id")
	opts := options.CreateCredentialOpts{
		ProjectID: projectID,
		UserID:    userID,
		Type:      "kirk",
	}

	credential, err := New(openstacker).Create(opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.id"), credential.ID)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.project_id"), credential.ProjectID)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.user_id"), credential.UserID)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.type"), credential.Type)
	assertion.Empty(credential.Blob)
}

func Test_Show_Credential(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	credentialID := apiv3.APIString("POST /credentials.credential.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/credentials/"+credentialID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /credentials/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	credential, err := New(openstacker).Show(credentialID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.id"), credential.ID)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.project_id"), credential.ProjectID)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.user_id"), credential.UserID)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.type"), credential.Type)
	assertion.Equal(apiv3.APIString("POST /credentials.credential.blob"), credential.Blob)
}

func Test_Update_Credential(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	credentialID := apiv3.APIString("POST /credentials.credential.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/credentials/"+credentialID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /credentials/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	blob := `{"user":{"id":123321},"access_key":{"id":"ak","secret":"sk"}}`
	opts := options.UpdateCredentialOpts{
		Blob: &blob,
	}

	credential, err := New(openstacker).Update(credentialID, opts)
	assertion.Nil(err)
	assertion.Equal(blob, credential.Blob)
}

func Test_Delete_Credential(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	credentialID := apiv3.APIString("POST /credentials.credential.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/credentials/"+credentialID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /credentials/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(credentialID)
	assertion.Nil(err)
}
