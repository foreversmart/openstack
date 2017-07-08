package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Domainer interface {
	All(opts options.ListDomainOpts) (domains []*models.DomainModel, err error)
	Create(opts options.CreateDomainOpts) (domain *models.DomainModel, err error)
	Show(id string) (domain *models.DomainModel, err error)
	Update(id string, opts options.UpdateDomainOpts) (domainInfo *models.DomainModel, err error)
	Delete(id string) (err error)
}
