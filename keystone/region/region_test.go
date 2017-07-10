package region

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

func Test_All_Region(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/regions")).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /regions"))
	// mitm.Pause()

	assertion := assert.New(t)

	regions, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.EqualValues(1, len(regions))
	assertion.Equal(apiv3.APIString("GET /regions.regions.0.id"), regions[0].ID)
	assertion.Equal(apiv3.APIString("GET /regions.regions.0.description"), regions[0].Description)
	assertion.Empty(regions[0].ParentRegionID)
}

func Test_Create_Region(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockAdminURL("/v3/regions")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /regions"))
	// mitm.Pause()

	assertion := assert.New(t)

	regionID := "TestingRegion"
	opts := options.CreateRegionOpts{
		Id: &regionID,
	}

	region, err := New(openstacker).Create(&opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("POST /regions.region.id"), region.ID)
	assertion.Equal(*opts.Id, region.ID)
	assertion.Empty(region.Description)
	assertion.Empty(region.ParentRegionID)
}

func Test_Show_Region(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	regionID := apiv3.APIString("POST /regions.region.id")

	mitm.MockRequest("GET", apiv3.MockAdminURL("/v3/regions/"+regionID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("GET /regions/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	region, err := New(openstacker).Show(regionID)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("GET /regions/:id.region.id"), region.ID)
	assertion.Equal(apiv3.APIString("GET /regions/:id.region.description"), region.Description)
	assertion.Empty(region.ParentRegionID)
}

func Test_Update_Region(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	regionID := apiv3.APIString("POST /regions.region.id")

	mitm.MockRequest("PATCH", apiv3.MockAdminURL("/v3/regions/"+regionID)).WithResponse(http.StatusOK, jsonheader, apiv3.APIString("PATCH /regions/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	desc := "Testing region for kirk"
	opts := options.UpdateRegionOpts{
		Description: &desc,
	}

	region, err := New(openstacker).Update(regionID, &opts)
	assertion.Nil(err)
	assertion.Equal(apiv3.APIString("PATCH /regions/:id.region.id"), region.ID)
	assertion.Equal(*opts.Description, region.Description)
	assertion.Empty(region.ParentRegionID)
}

func Test_Delete_Region(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	regionID := apiv3.APIString("POST /regions.region.id")

	mitm.MockRequest("DELETE", apiv3.MockAdminURL("/v3/regions/"+regionID)).WithResponse(http.StatusNoContent, jsonheader, apiv3.APIString("DELETE /regions/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(regionID)
	assertion.Nil(err)
}
