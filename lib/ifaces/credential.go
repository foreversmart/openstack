package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Credentialer interface {
	All() (credentials []*models.CredentialModel, err error)
	AllByParams(opts *options.ListCredentialOpts) (credentials []*models.CredentialModel, err error)
	Create(opts options.CreateCredentialOpts) (credential *models.CredentialModel, err error)
	Show(id string) (credential *models.CredentialModel, err error)
	Update(id string, opts options.UpdateCredentialOpts) (credential *models.CredentialModel, err error)
	Delete(id string) (err error)
}
