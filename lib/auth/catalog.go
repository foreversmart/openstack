package auth

import (
	"encoding/json"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	tokens2 "github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
func CatalogV2(client *gophercloud.ProviderClient, options gophercloud.AuthOptions) (catalog *tokens2.ServiceCatalog, err error) {
	v2Client := openstack.NewIdentityV2(client)

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: v2Client.IdentityEndpoint,
		Username:         options.Username,
		Password:         options.Password,
		TenantID:         options.TenantID,
	}

	result := tokens2.Create(v2Client, tokens2.AuthOptions{AuthOptions: opts})

	catalog, err = result.ExtractServiceCatalog()
	return
}

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
func CatalogV3(client *gophercloud.ProviderClient, options gophercloud.AuthOptions) (catalog *tokens3.ServiceCatalog, err error) {
	v3Client := openstack.NewIdentityV3(client)

	// copy the auth options to a local variable that we can change. `options`
	// needs to stay as-is for reauth purposes
	v3Options := options

	var scope *tokens3.Scope
	if options.TenantID != "" {
		scope = &tokens3.Scope{
			ProjectID: options.TenantID,
		}
		v3Options.TenantID = ""
		v3Options.TenantName = ""
	} else {
		if options.TenantName != "" {
			scope = &tokens3.Scope{
				ProjectName: options.TenantName,
				DomainID:    options.DomainID,
				DomainName:  options.DomainName,
			}
			v3Options.TenantName = ""
		}
	}

	result := tokens3.Create(v3Client, tokens3.AuthOptions{AuthOptions: v3Options}, scope)

	catalog, err = result.ExtractServiceCatalog()
	return
}

func MarshalCatalog(catalog interface{}) (string, error) {
	res, err := json.Marshal(catalog)
	return string(res), err
}

func UnmarshalCatalogV2(v interface{}) (catalog *tokens2.ServiceCatalog, err error) {
	switch v.(type) {
	case *tokens2.ServiceCatalog:
		catalog, _ = v.(*tokens2.ServiceCatalog)

	case string:
		s, _ := v.(string)

		err = json.Unmarshal([]byte(s), &catalog)

	case []byte:
		data, _ := v.([]byte)

		err = json.Unmarshal(data, &catalog)

	default:
		// try to encode then decode
		var data []byte

		data, err = json.Marshal(v)
		if err == nil {
			err = json.Unmarshal(data, &catalog)
		}

	}

	return
}

func UnmarshalCatalogV3(v interface{}) (catalog *tokens3.ServiceCatalog, err error) {
	switch v.(type) {
	case *tokens3.ServiceCatalog:
		catalog, _ = v.(*tokens3.ServiceCatalog)

	case string:
		s, _ := v.(string)

		err = json.Unmarshal([]byte(s), &catalog)

	case []byte:
		data, _ := v.([]byte)

		err = json.Unmarshal(data, &catalog)

	default:
		// try to encode then decode
		var data []byte

		data, err = json.Marshal(v)
		if err == nil {
			err = json.Unmarshal(data, &catalog)
		}

	}

	return
}
