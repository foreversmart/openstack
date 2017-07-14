package internal

import (
	"net/http"
	"strings"

	"github.com/dolab/gogo"
	"github.com/qbox/openstack-golang-sdk/lib/auth"
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

type Openstack struct {
	options    auth.AuthOptions
	client     *gophercloud.ProviderClient
	httpclient *http.Client
	authorized bool
	logger     gogo.Logger

	region     string
	endpoint   string
	apiversion string
	projectID  string
}

func New(endpoint string) *Openstack {
	client, err := openstack.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	return &Openstack{
		client:     client,
		endpoint:   endpoint,
		apiversion: auth.V3,
	}
}

func NewWithRegion(region, endpoint string) *Openstack {
	os := New(endpoint)
	os.region = region

	return os
}

func (os *Openstack) WithHTTPClient(client *http.Client) {
	os.httpclient = client

	if os.client != nil {
		os.client.HTTPClient = *client
	}

	return
}

func (os *Openstack) ProviderClient() *gophercloud.ProviderClient {
	return os.client
}

func (os *Openstack) ServiceClient(endpoint string) *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: os.client,
		Endpoint:       gophercloud.NormalizeURL(endpoint),
	}
}

func (os *Openstack) ProjectID() string {
	return os.projectID
}

func (os *Openstack) IsAuthorized() bool {
	return os.authorized
}

func (os *Openstack) Authenticate() (err error) {
	if os.IsAuthorized() {
		return
	}

	if os.options.IsTokenValid() {
		err = os.AuthByToken(os.options)
	} else {
		// take user id as high priority
		if os.options.UserID != "" && os.options.Username != "" {
			os.options.Username = ""
		}

		err = os.AuthByPassword(os.options)
	}

	return
}

func (os *Openstack) Auth(options auth.AuthOptions) error {
	os.options = options

	return os.Authenticate()
}

func (os *Openstack) AuthByPassword(options auth.AuthOptions) (err error) {
	if (options.UserID == "" && options.Username == "") || options.Password == "" {
		return errors.ErrInvalidParams
	}

	// adjust default api version
	if !options.IsVersionValid() {
		switch options.UserID {
		case "":
			options.Version = auth.V2

		default: // The base Identity V2 API does not accept authentication by UserID
			options.Version = auth.V3

		}
	}

	err = auth.Authenticate(os.client, options)
	if err != nil {
		return err
	}

	if os.client.TokenID == "" {
		return errors.ErrTokenID
	}

	// update state after auth success
	os.options = options
	os.projectID = os.options.TenantID
	os.authorized = true

	return nil
}

func (os *Openstack) AuthByToken(options auth.AuthOptions) (err error) {
	if !options.IsTokenValid() {
		return errors.ErrInvalidParams
	}

	// adjust default api version
	if !options.IsVersionValid() {
		switch options.UserID {
		case "":
			options.Version = auth.V2

		default: // The base Identity V2 API does not accept authentication by UserID
			options.Version = auth.V3

		}
	}

	err = auth.Authenticate(os.client, options)
	if err != nil {
		return err
	}

	if os.client.TokenID == "" {
		return errors.ErrTokenID
	}

	// update state after auth success
	os.options = options
	os.projectID = os.options.TenantID
	os.authorized = true

	return nil
}

func (os *Openstack) AuthByAccessKey(id, secret string) (err error) {
	if id == "" || secret == "" {
		return errors.ErrInvalidParams
	}

	// // adjust default api version
	// if options.Version == "" {
	// 	options.Version = os.apiversion
	// }

	err = errors.ErrNotImplemented
	return
}

func (os *Openstack) AuthRequest(method, subpath string, opts gophercloud.RequestOpts) error {
	err := os.Authenticate()
	if err != nil {
		return err
	}

	if subpath[0] != '/' {
		subpath = "/" + subpath
	}

	absurl := strings.TrimSuffix(os.endpoint, "/") + subpath

	_, err = os.client.Request(method, absurl, opts)
	return err
}
