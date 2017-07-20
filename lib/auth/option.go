package auth

import (
	"time"

	"github.com/rackspace/gophercloud"
)

const (
	V2 = "v2"
	V3 = "v3"
)

type AuthSuccessFunc func(tokenID string, expiredAt time.Time, result gophercloud.Result) error

type AuthOptions struct {
	*gophercloud.AuthOptions

	Version   string
	Catalog   interface{}
	ExpiredAt time.Time

	ForceReauth bool            // set true will create new token with options even old token is valid
	SuccessFunc AuthSuccessFunc // execute when reauth success, useful for cache
}

func (options *AuthOptions) GophercloudAuthOptions() gophercloud.AuthOptions {
	return *options.AuthOptions
}

func (options *AuthOptions) IsTokenValid() bool {
	if options.TokenID == "" || options.Catalog == nil {
		return false
	}

	// is expired?
	return options.ExpiredAt.After(time.Now().Add(-time.Second * 60))
}

func (options *AuthOptions) IsVersionValid() bool {
	switch options.Version {
	case V2, V3:
		return true
	}

	return false
}
