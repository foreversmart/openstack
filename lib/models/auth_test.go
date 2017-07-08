package models

import (
	"encoding/json"
	"testing"

	"github.com/golib/assert"
)

func Test_AuthModel(t *testing.T) {
	assertion := assert.New(t)

	auth := &AuthModel{}
	assertion.Empty(auth.ProjectID)
	assertion.Empty(auth.TokenID)
	assertion.True(auth.ExpiredAt.IsZero())

	data, _ := json.Marshal(&auth)

	sdata := string(data)
	assertion.NotContains(sdata, "password")
	assertion.NotContains(sdata, "token_id")
	assertion.NotContains(sdata, "expired_at")
}
