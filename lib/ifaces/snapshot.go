package ifaces

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
)

type Snapshoter interface {
	All() (snapshotModels []*models.SnapshotModel, err error)
	Show(id string) (snapshot *models.SnapshotModel, err error)
	Create(opts options.CreateSnapshotOpts) (id string, err error)
	Update(snapshotID string, opts *options.UpdateSnapshotOpts) error
	Delete(id string) error
}
