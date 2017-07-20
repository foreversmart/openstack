package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type SecurityGrouper interface {
	All() ([]*models.SecurityGroupModel, error)
	AllByParams(opts *options.ListSecurityGroupsOpts) (infos []*models.SecurityGroupModel, err error)
	Create(opts *options.CreateSecurityGroupOpts) (info *models.SecurityGroupModel, err error)
	Query(id string) (*models.SecurityGroupModel, error)
	Delete(id string) error
}
