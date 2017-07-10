package auth

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	tokens2 "github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

// Authenticate or re-authenticate against the most recent identity service supported at the provided endpoint.
func Authenticate(client *gophercloud.ProviderClient, options AuthOptions) error {
	switch options.Version {
	case V2:
		return v2auth(client, "", options)

	case V3:
		return v3auth(client, "", options)
	}

	// The switch statement must be out of date from the versions list.
	return fmt.Errorf("Unrecognized identity version: %s", options.Version)
}

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
func AuthenticateV2(client *gophercloud.ProviderClient, options AuthOptions) (err error) {
	return v2auth(client, "", options)
}

func v2auth(client *gophercloud.ProviderClient, endpoint string, options AuthOptions) (err error) {
	if options.IsTokenValid() {
		catalog, err := UnmarshalCatalogV2(options.Catalog)
		if err != nil {
			return err
		}

		client.TokenID = options.TokenID
		client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
			return openstack.V2EndpointURL(catalog, opts)
		}
		if options.AllowReauth {
			client.ReauthFunc = func() error {
				// reset token id
				client.TokenID = ""

				// try to reauth by password
				return v2auth(client, endpoint, options)
			}
		}

		return nil
	}

	v2client := openstack.NewIdentityV2(client)
	if endpoint != "" {
		v2client.Endpoint = endpoint
	}

	result := tokens2.Create(v2client, tokens2.AuthOptions{AuthOptions: options.GophercloudAuthOptions()})

	token, err := result.ExtractToken()
	if err != nil {
		return
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return
	}

	// callback if exists
	// NOTE: is it safety to ignore callback error?
	if options.SuccessFunc != nil {
		err = options.SuccessFunc(token.ID, token.ExpiresAt, result.Result)
		if err != nil {
			return
		}
	}

	// setup client
	client.TokenID = token.ID
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return openstack.V2EndpointURL(catalog, opts)
	}
	if options.AllowReauth {
		client.ReauthFunc = func() error {
			// reset token id
			client.TokenID = ""

			// try to reauth by password
			return v2auth(client, endpoint, options)
		}
	}

	return
}

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(client *gophercloud.ProviderClient, options AuthOptions) (err error) {
	return v3auth(client, "", options)
}

func v3auth(client *gophercloud.ProviderClient, endpoint string, options AuthOptions) (err error) {
	if options.IsTokenValid() {
		catalog, err := UnmarshalCatalogV3(options.Catalog)
		if err != nil {
			return err
		}

		client.TokenID = options.TokenID
		client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
			return openstack.V3EndpointURL(catalog, opts)
		}
		if options.AllowReauth {
			client.ReauthFunc = func() error {
				// reset token id
				client.TokenID = ""

				// try to reauth by password
				return v3auth(client, endpoint, options)
			}
		}

		return nil
	}

	// Override the generated service endpoint with the one returned by the version endpoint.
	v3client := openstack.NewIdentityV3(client)
	if endpoint != "" {
		v3client.Endpoint = endpoint
	}

	// copy the auth options to a local variable that we can change. `options`
	// needs to stay as-is for reauth purposes
	v3options := options.GophercloudAuthOptions()

	var scope *tokens3.Scope
	if options.TenantID != "" {
		scope = &tokens3.Scope{
			ProjectID: options.TenantID,
		}
		v3options.TenantID = ""
		v3options.TenantName = ""
	} else {
		if options.TenantName != "" {
			scope = &tokens3.Scope{
				ProjectName: options.TenantName,
				DomainID:    options.DomainID,
				DomainName:  options.DomainName,
			}
			v3options.TenantName = ""
		}
	}

	result := tokens3.Create(v3client, tokens3.AuthOptions{AuthOptions: v3options}, scope)

	token, err := result.ExtractToken()
	if err != nil {
		return
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return
	}

	// callback if exists
	// NOTE: is it safety to ignore callback error?
	if options.SuccessFunc != nil {
		err = options.SuccessFunc(token.ID, token.ExpiresAt, result.Result)
		if err != nil {
			return
		}
	}

	// setup client
	client.TokenID = token.ID
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return openstack.V3EndpointURL(catalog, opts)
	}
	if options.AllowReauth {
		client.ReauthFunc = func() error {
			// reset token id
			client.TokenID = ""

			// auth by password
			return v3auth(client, endpoint, options)
		}
	}

	return
}
