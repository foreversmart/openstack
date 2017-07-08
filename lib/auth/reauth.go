package auth

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

const (
	MaxReauthAttempts int32 = 3
)

// ReauthRetryer satisfies the http.RoundTripper interface and is used to
// customize the reauth times limit
type ReauthRetryer struct {
	rt                 http.RoundTripper
	numReauthAttempted int32
}

func NewReauthRetryer(rt http.RoundTripper) *ReauthRetryer {
	return &ReauthRetryer{
		rt:                 rt,
		numReauthAttempted: 0,
	}
}

func (retryer *ReauthRetryer) RoundTrip(request *http.Request) (response *http.Response, err error) {
	if atomic.LoadInt32(&retryer.numReauthAttempted) >= MaxReauthAttempts {
		return response, fmt.Errorf("Tried to re-authenticate %d times without success.", MaxReauthAttempts)
	}

	response, err = retryer.rt.RoundTrip(request)
	if response == nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		atomic.AddInt32(&retryer.numReauthAttempted, 1)
	}

	return response, err
}
