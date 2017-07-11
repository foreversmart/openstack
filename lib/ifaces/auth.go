package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Auther interface {
	AuthScoper
	AuthCataloger

	Show(opts *options.ShowTokenOpts) (token *models.TokenModel, err error)
	HasToken(opts options.HeadTokenOpts) (exist bool, err error)
	Delete() (err error)
}

type AuthScoper interface {
	AllDomain() (domains []*models.DomainModel, err error)
	AllProject() (projects []*models.ProjectModel, err error)
}

type AuthCataloger interface {
	AllCatalog() (catalog *models.CatalogModel, err error)
}
