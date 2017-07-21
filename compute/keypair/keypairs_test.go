package keypair

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

var (
	testKeypairName string
)

const computerPort = "8774"

func Test_Create_Keypair(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest(http.MethodPost, apiv2.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/os-keypairs")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("POST /keypairs"))
	// mitm.Pause()

	assertion := assert.New(t)

	key, err := New(openstacker).Create(&options.CreateKeypairOpts{
		Name: options.String("testkey"),
	})

	assertion.Nil(err)
	assertion.NotNil(key)

	assertModel(assertion, "POST /keypairs.keypair", key)

	testKeypairName = key.Name
}

func Test_All_Keypairs(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest(http.MethodGet, apiv2.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/os-keypairs")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /keypairs"))
	// mitm.Pause()

	assertion := assert.New(t)

	keypairs, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.NotNil(keypairs)
	assertion.EqualValues(2, len(keypairs))

	assertModel(assertion, "GET /keypairs.keypairs.0.keypair", keypairs[0])
}

func Test_Show_Keypairs(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest(http.MethodGet, apiv2.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/os-keypairs/"+testKeypairName)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /keypairs/:name"))
	// mitm.Pause()

	assertion := assert.New(t)

	key, err := New(openstacker).Show(testKeypairName)

	assertion.Nil(err)
	assertion.NotNil(key)

	assertModel(assertion, "GET /keypairs/:name.keypair", key)
}

func Test_Delete_Keypairs(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest(http.MethodDelete, apiv2.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/os-keypairs/"+testKeypairName)).WithResponse(http.StatusAccepted, jsonheader, apiv2.APIString("DELETE /keypairs/:name"))
	// mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(testKeypairName)
	assertion.Nil(err)
}

func assertModel(assertion *assert.Assertions, pathPrefix string, key *models.KeypairModel) {
	assertion.Equal(apiv2.APIString(pathPrefix+".name"), key.Name)
	assertion.Equal(apiv2.APIString(pathPrefix+".type"), key.Type)
	assertion.Equal(apiv2.APIString(pathPrefix+".public_key"), key.PublicKey)
}
