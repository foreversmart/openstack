package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Keypairer interface {
	Create(opts *options.CreateKeypairOpt) (keypair *models.KeypairModel, err error)
	All() (KeypairModels []*models.KeypairModel, err error)
	AllByParams(opts *options.ListKeypairOpt) (keypairModels []*models.KeypairModel, err error)
	Show(name string) (keypairModel *models.KeypairModel, err error)
	Delete(name string) error
}
