package secrule

import (
	"net/http"
	"testing"
)

const (
	networkPort = "9696"
)

func Test_Create_Secrule(t *testing.T) {
	mitm := mocker.StubDefaultTransport(t)

	mitm.MockRequest("POST", apiv3.MockResourceURLWithPort(networkPort, "/v2.0/security-group-rules")).WithResponse(http.StatusCreated, jsonheader, apiv3.APIString("POST /security-group-rules"))
	//mitm.Pause()

    
}
