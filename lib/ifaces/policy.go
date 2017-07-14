package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Policier interface {
	All() (policies []*models.PolicyModel, err error)
	AllByParams(opts *options.ListPolicyOpts) (policies []*models.PolicyModel, err error)
	Create(opts options.CreatePolicyOpts) (policy *models.PolicyModel, err error)
	Show(id string) (policy *models.PolicyModel, err error)
	Update(id string, opts options.UpdatePolicyOpts) (policyInfo *models.PolicyModel, err error)
	Delete(id string) (err error)
}
