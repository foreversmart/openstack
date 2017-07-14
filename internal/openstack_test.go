package internal

import (
	"testing"
	"time"

	"net/http"

	"github.com/dolab/httpmitm"
	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/auth"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/rackspace/gophercloud"
)

func Test_New(t *testing.T) {
	endpoint := apiv2.GetString("user.endpoint")

	assertion := assert.New(t)
	os := New(endpoint)

	assertion.NotNil(os.client)
	assertion.Equal(endpoint, os.endpoint)
	assertion.False(os.IsAuthorized())
}

func Test_AuthByPassword_V2(t *testing.T) {
	mitm := httpmitm.NewMitmTransport().StubDefaultTransport(t)
	defer mitm.UnstubDefaultTransport()

	endpoint := apiv2.GetString("user.endpoint")

	header := http.Header{}
	header.Add("Content-Type", "application/json")

	mitm.MockRequest("POST", apiv2.MockURL("/v2.0/tokens")).WithResponse(200, header, apiv2.APIString("scoped"))
	// mitm.Pause()

	assertion := assert.New(t)
	os := New(endpoint)

	options := auth.AuthOptions{
		Version: auth.V2,
		AuthOptions: &gophercloud.AuthOptions{
			Username: apiv2.GetString("user.username"),
			Password: apiv2.GetString("user.password"),
			TenantID: apiv2.GetString("user.project_id"),
		},
	}

	err := os.AuthByPassword(options)
	assertion.Nil(err)
	assertion.True(os.IsAuthorized())
	assertion.NotNil(os.client)
	assertion.Equal(apiv2.APIString("scoped.access.token.id"), os.client.TokenID)
}

func Test_AuthByPassword_V3(t *testing.T) {
	mitm := httpmitm.NewMitmTransport().StubDefaultTransport(t)
	defer mitm.UnstubDefaultTransport()

	endpoint := apiv3.GetString("admin.endpoint")

	header := http.Header{}
	header.Add("Content-Type", "application/json")
	header.Add("X-Subject-Token", apiv3.GetString("token.id"))

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/auth/tokens")).WithResponse(201, header, apiv3.APIString("scoped"))
	// mitm.Pause()

	assertion := assert.New(t)
	os := New(endpoint)

	options := auth.AuthOptions{
		Version: auth.V3,
		AuthOptions: &gophercloud.AuthOptions{
			DomainName: apiv3.GetString("admin.domain_name"),
			Username:   apiv3.GetString("admin.username"),
			Password:   apiv3.GetString("admin.password"),
		},
		SuccessFunc: func(tokenid string, expire time.Time, result gophercloud.Result) error {
			token, err := models.ExtractToken(result)
			assertion.Nil(err)
			assertion.Equal(token.User.ID, apiv3.GetString("v3.scoped.token.user.id"))
			return nil
		},
	}

	err := os.AuthByPassword(options)
	assertion.Nil(err)
	assertion.True(os.IsAuthorized())
	assertion.NotNil(os.client)
	assertion.Equal(header.Get("X-Subject-Token"), os.client.TokenID)
}

func Test_AuthByToken_V2(t *testing.T) {
	mitm := httpmitm.NewMitmTransport().StubDefaultTransport(t)
	defer mitm.UnstubDefaultTransport()

	endpoint := apiv2.GetString("user.endpoint")
	tokenID := apiv2.APIString("scoped.access.token.id")

	header := http.Header{}
	header.Add("Content-Type", "application/json")

	assertion := assert.New(t)
	os := New(endpoint)

	options := auth.AuthOptions{
		Version: auth.V2,
		AuthOptions: &gophercloud.AuthOptions{
			TokenID:  string(tokenID),
			TenantID: apiv2.GetString("user.project_id"),
		},
		Catalog:   apiv2.MockCatalog(),
		ExpiredAt: time.Now().Add(time.Hour * 1),
	}

	err := os.AuthByToken(options)
	assertion.Nil(err)
	assertion.True(os.IsAuthorized())
	assertion.NotNil(os.client)
	assertion.Equal(apiv2.APIString("scoped.access.token.id"), os.client.TokenID)
}

func Test_AuthByToken_V3(t *testing.T) {
	mitm := httpmitm.NewMitmTransport().StubDefaultTransport(t)
	defer mitm.UnstubDefaultTransport()

	endpoint := apiv3.GetString("admin.endpoint")
	tokenID := apiv2.GetString("token.id")

	header := http.Header{}
	header.Add("Content-Type", "application/json")

	// mitm.MockRequest("POST", apiv3.MockURL("/v2.0/tokens")).WithResponse(200, header, apiv3.API("POST /tokens", "scoped"))
	// mitm.Pause()

	assertion := assert.New(t)
	os := New(endpoint)

	options := auth.AuthOptions{
		Version: auth.V3,
		AuthOptions: &gophercloud.AuthOptions{
			TokenID:  string(tokenID),
			TenantID: apiv3.GetString("user.project_id"),
		},
		Catalog:   apiv3.MockCatalog(),
		ExpiredAt: time.Now().Add(time.Hour * 1),
	}

	err := os.AuthByToken(options)
	assertion.Nil(err)
	assertion.True(os.IsAuthorized())
	assertion.NotNil(os.client)
	assertion.Equal(tokenID, os.client.TokenID)
}
