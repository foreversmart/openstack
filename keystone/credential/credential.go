package credential

import (
	"github.com/kirk-enterprise/openstack/lib/errors"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	CredentialUrl = "credentials"
)

type Credential struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Credential {
	return &Credential{
		Client: client,
	}
}

func (cred *Credential) All() (credentials []*models.CredentialModel, err error) {
	return cred.AllByParams(nil)
}

func (cred *Credential) AllByParams(opts *options.ListCredentialOpts) (credentials []*models.CredentialModel, err error) {
	client, err := cred.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(CredentialUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractCredentials(result)
}

func (cred *Credential) Create(opts options.CreateCredentialOpts) (credential *models.CredentialModel, err error) {
	client, err := cred.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Post(client.ServiceURL(CredentialUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractCredential(result)
}

func (cred *Credential) Show(id string) (credential *models.CredentialModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := cred.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(CredentialUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractCredential(result)
}

func (cred *Credential) Update(id string, opts options.UpdateCredentialOpts) (credential *models.CredentialModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := cred.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Patch(client.ServiceURL(CredentialUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractCredential(result)
}

func (cred *Credential) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := cred.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(CredentialUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
