package server

import (
	"net/http"
	"testing"

	"github.com/golib/assert"
)

func Test_Bind_Unbind_Key(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	serverID := apiv3.APIString("POST /servers.server.id")

	mitm.MockRequest(
		http.MethodPost,
		apiv3.MockResourceURLWithPort(computerPort, "v2.1/"+testProjectId+"/servers/"+serverID+"/action")).
		WithResponse(
			http.StatusAccepted,
			jsonheader,
			apiv3.APIString("POST /servers/:id/action/image")).AnyTimes()
	// mitm.Pause()

	assertion := assert.New(t)

	keys := []string{"test key"}

	err := NewServerKeyer(openstacker).Bind(serverID, keys)
	assertion.Nil(err)

	err = NewServerKeyer(openstacker).Unbind(serverID, keys)
	assertion.Nil(err)
}
