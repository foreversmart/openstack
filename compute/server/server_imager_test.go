package server

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_Create_ServerImage(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(
		http.MethodPost,
		apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).
		WithResponse(
			http.StatusAccepted,
			jsonheader,
			apiv3.APIString("POST /servers/:id/action/image"))

	mitm.MockRequest(
		http.MethodGet,
		apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID)).
		WithResponse(
			http.StatusOK,
			jsonheader,
			apiv3.APIString("GET /servers/:id"))
	// mitm.Pause()

	assertion := assert.New(t)

	imageID, err := NewServerImager(openstacker).Create(serverID, "test create server image")

	assertion.Nil(err)
	assertion.NotNil(imageID)

	assertion.Equal(apiv3.APIString("POST /servers/:id/action/image.image_id"), imageID)
}
