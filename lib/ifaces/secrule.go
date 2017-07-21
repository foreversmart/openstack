package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type SecurityRuler interface {
	Create(opts *options.CreateSecruleOpts) (secrule *models.SecGroupRuleModel, err error)
	Show(id string) (*models.SecGroupRuleModel, error)
	Delete(id string) error
}
