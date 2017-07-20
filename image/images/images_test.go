package images

import (
	"io/ioutil"
	"net/http"
	"strings"
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

func Test_All_Images(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(imagePort, "/v2/images")).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /images"))
	// mitm.Pause()

	assertion := assert.New(t)

	images, err := New(openstacker).All()
	assertion.Nil(err)
	assertion.NotNil(images)
	assertion.EqualValues(2, len(images))

	assertModel(assertion, "GET /images.images.0", images[0])
}

func Test_Show_Images(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("GET", apiv2.MockResourceURLWithPort(imagePort, "/v2/images/"+testImageId)).WithResponse(http.StatusOK, jsonheader, apiv2.APIString("GET /images/:id"))

	assertion := assert.New(t)

	image, err := New(openstacker).Show(testImageId)
	assertion.Nil(err)
	assertion.NotNil(image)

	assertModel(assertion, "GET /images/:id", image)
}

func Test_Update_Images(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("PATCH", apiv2.MockResourceURLWithPort(imagePort, "/v2/images/"+testImageId)).WithResponse(http.StatusOK, nil, apiv2.APIString("PATCH /images/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	tags := []string{"tag1", "tag2"}

	image, err := New(openstacker).Update(testImageId, &options.UpdateImagesOpts{
		Name: options.String("updatename"),
		Tags: &tags,
	})
	assertion.Nil(err)
	assertion.NotNil(image)

	assertModel(assertion, "PATCH /images/:id", image)
}

func Test_Delete_Images(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("DELETE", apiv2.MockResourceURLWithPort(imagePort, "/v2/images/"+testImageId)).WithResponse(http.StatusNoContent, nil, apiv2.APIString("DELETE /images/:id"))
	//mitm.Pause()

	assertion := assert.New(t)

	err := New(openstacker).Delete(testImageId)
	assertion.Nil(err)
}

func Test_Upload_Image(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	jsonheader := http.Header{}
	jsonheader.Add("Content-Type", "application/json")

	mitm.MockRequest(http.MethodPut, apiv2.MockResourceURLWithPort(imagePort, "/v2/images/"+testImageId+"/file")).WithResponse(http.StatusNoContent, jsonheader, "")
	//mitm.Pause()

	i := New(openstacker)
	assertion := assert.New(t)

	data := strings.NewReader("test")
	err := i.Upload(testImageId, data)

	assertion.Nil(err)
}

func Test_Download_Image(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	jsonheader := http.Header{}
	jsonheader.Add("Content-Type", "application/json")
	jsonheader.Add("Content-Md5", "sdlfjladmxllckvk")
	jsonheader.Add("Content-Length", "28")

	imageBody := "test-image"
	mitm.MockRequest(http.MethodGet, apiv2.MockResourceURLWithPort(imagePort, "/v2/images/"+testImageId+"/file")).WithResponse(http.StatusOK, jsonheader, imageBody)
	//mitm.Pause()

	i := New(openstacker)
	assertion := assert.New(t)

	data, err := i.Download(testImageId)
	assertion.Nil(err)

	body, err := ioutil.ReadAll(data)

	assertion.Nil(err)
	assertion.Equal(imageBody, string(body))
}

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
