package auth

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	AuthUrl    = "auth"
	TokenUrl   = "tokens"
	CatalogUrl = "catalog"
	ProjectUrl = "projects"
	DomainUrl  = "domains"
)

type Auth struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Auth {
	return &Auth{
		Client: client,
	}
}

func (auth *Auth) Show(opts *options.ShowTokenOpts) (token *models.TokenModel, err error) {
	client, err := auth.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(AuthUrl, TokenUrl), &result.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"X-Subject-Token": opts.TokenID},
		OkCodes:     []int{200},
	})

	return models.ExtractToken(result)
}

func (auth *Auth) HasToken(opts *options.HeadTokenOpts) (exist bool, err error) {
	return false, errors.ErrNotImplemented
}

func (auth *Auth) Delete() (err error) {
	return errors.ErrNotImplemented
}

//AuthScoper interface
func (auth *Auth) AllDomain() (domains []*models.DomainModel, err error) {
	return nil, errors.ErrNotImplemented
}
func (auth *Auth) AllProject() (projects []*models.ProjectModel, err error) {
	return nil, errors.ErrNotImplemented
}

// AuthCataloger interface {
func (auth *Auth) AllCatalog() (catalog *models.CatalogModel, err error) {
	return nil, errors.ErrNotImplemented
}
