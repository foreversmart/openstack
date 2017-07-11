package auth

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

func Test_Show_Token(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/auth/tokens")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("scoped"))
	// mitm.Pause()

	assertion := assert.New(t)

	token, err := New(openstacker).Show(&options.ShowTokenOpts{})

	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /auth/tokens.token.user.id"), token.User.ID)
}
