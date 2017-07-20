package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type SecurityGrouper interface {
	All() ([]*models.SecurityGroupModel, error)
	AllByParams(opts *options.ListSecurityGroupsOpts) (infos []*models.SecurityGroupModel, err error)
	Create(opts *options.CreateSecurityGroupOpts) (info *models.SecurityGroupModel, err error)
	Show(id string, opts *options.ShowSecurityGroupOpts) (securitygroup *models.SecurityGroupModel, err error)
	Delete(id string) error
}
