package flavors

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

func Test_Create_Flavors(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)
	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort("8774", "/v2.1/fcfeddf071284e4a8c54760d4bf67c29/flavors")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("POST /flavors"))
	// mitm.Pause()

	assertion := assert.New(t)

	flavorName := "test_flavor"
	opts := options.CreateFlavorOpts{
		Name:  &flavorName,
		Ram:   1024,
		Disk:  10,
		Vcpus: 2,
	}

	flavor, err := New(openstacker).Create(&opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /flavors.flavor.id"), flavor.ID)
	assertion.Equal(*opts.Name, flavor.Name)
	assertion.Empty(flavor.Swap)

}

func Test_All_Flavors(t *testing.T) {

	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort("8774", "/v2.1/fcfeddf071284e4a8c54760d4bf67c29/flavors/detail")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /flavors"))
	// mitm.Pause()

	assertion := assert.New(t)

	flavors, err := New(openstacker).All()
	assertion.Nil(err)

	assertion.Equal(apiv3.APIString("GET /flavors.flavors.0.id"), flavors[0].ID)
	assertion.Equal(apiv3.APIString("GET /flavors.flavors.0.name"), flavors[0].Name)
	assertion.Empty(flavors[0].Swap)

}

func Test_Show_Flavors(t *testing.T) {

	mitm := mocker.StubDefaultTransport(t)
	flavorID := apiv3.APIString("GET /flavors/:id.flavor.id")

	mitm.MockRequest("GET", apiv3.MockResourceURLWithPort("8774", "/v2.1/fcfeddf071284e4a8c54760d4bf67c29/flavors/"+flavorID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /flavors/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	flavor, err := New(openstacker).Show(flavorID)
	assertion.Nil(err)

	assertion.Equal(apiv3.APIString("GET /flavors/:id.flavor.id"), flavor.ID)
	assertion.Equal(apiv3.APIString("GET /flavors/:id.flavor.name"), flavor.Name)
	assertion.Empty(flavor.Swap)

}

func Test_Delete_Flavors(t *testing.T) {

	mitm := mocker.StubDefaultTransport(t)

	flavorID := apiv3.APIString("GET /flavors/:id.flavor.id")

	mitm.MockRequest("DELETE", apiv3.MockResourceURLWithPort("8774", "/v2.1/fcfeddf071284e4a8c54760d4bf67c29/flavors/"+flavorID)).WithResponse(http.StatusAccepted, jsonheader, apiv3.APIString("DELETE /flavors/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(flavorID)
	assertion.Nil(err)
}
