package images

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

var (
	testImageId string
)

const (
	imagePort = "9292"
)

func Test_Create_Images(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv2.MockResourceURLWithPort(imagePort, "/v2/images")).WithResponse(http.StatusCreated, jsonheader, apiv2.APIString("POST /images"))
	//mitm.Pause()

	assertion := assert.New(t)

	image, err := New(openstacker).Create(&options.CreateImagesOpts{
		Name: options.String("test image"),
	})

	assertion.Nil(err)
	assertion.NotNil(image)

	assertModel(assertion, "POST /images", image)

	testImageId = image.ID
}

// func Test_All_Images(t *testing.T) {
// 	mitm := mocker.StubDefaultTransport(t)

// 	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /floatingips"))
// 	// mitm.Pause()

// 	assertion := assert.New(t)

// 	ips, err := New(openstacker).All()
// 	assertion.Nil(err)
// 	assertion.NotNil(ips)
// 	assertion.EqualValues(2, len(ips))

// 	assertModel(assertion, "GET /floatingips.floatingips.0", ips[0])
// }

// func Test_Show_Images(t *testing.T) {
// 	mitm := mocker.StubDefaultTransport(t)

// 	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /floatingips/:id"))

// 	assertion := assert.New(t)

// 	ip, err := New(openstacker).Show(testFloatingipID)
// 	assertion.Nil(err)
// 	assertion.NotNil(ip)

// 	assertModel(assertion, "GET /floatingips/:id.floatingip", ip)
// }

// func Test_Update_Images(t *testing.T) {
// 	mitm := mocker.StubDefaultTransport(t)

// 	mitm.MockRequest("PUT", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusOK, nil, apiv2.APIString("PUT /floatingips/:id"))
// 	//mitm.Pause()

// 	assertion := assert.New(t)

// 	ip, err := New(openstacker).Update(testFloatingipID, &options.UpdateFloatingIPOpts{
// 		Description: options.String("update floatingip desc"),
// 	})
// 	assertion.Nil(err)
// 	assertion.NotNil(ip)

// 	assertModel(assertion, "PUT /floatingips/:id.floatingip", ip)
// }

// func Test_Delete_Images(t *testing.T) {
// 	mitm := mocker.StubDefaultTransport(t)

// 	mitm.MockRequest("DELETE", apiv2.MockResourceURLWithPort(networkPort, "/v2.0/floatingips/"+testFloatingipID)).WithResponse(http.StatusNoContent, nil, apiv2.APIString("DELETE /floatingips/:id"))
// 	//mitm.Pause()

// 	assertion := assert.New(t)

// 	err := New(openstacker).Delete(testFloatingipID)
// 	assertion.Nil(err)
// }

func assertModel(assertion *assert.Assertions, pathPrefix string, image *models.ImageModel) {
	assertion.Equal(apiv2.APIString(pathPrefix+".id"), image.ID)
	assertion.Equal(apiv2.APIString(pathPrefix+".name"), image.Name)
	assertion.Equal(apiv2.APIString(pathPrefix+".status"), image.Status)
	assertion.Equal(apiv2.APIString(pathPrefix+".visibility"), image.Visibility)
	assertion.Equal(apiv2.APIString(pathPrefix+".self"), image.Self)
	assertion.Equal(apiv2.APIString(pathPrefix+".file"), image.File)
	assertion.Equal(apiv2.APIString(pathPrefix+".schema"), image.Schema)
	assertion.Equal(apiv2.APIString(pathPrefix+".disk_format"), image.DiskFormat)
	assertion.Equal(apiv2.APIString(pathPrefix+".created_at"), image.CreatedAt)

}
