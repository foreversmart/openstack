package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Snapshoter interface {
	Create(opts *options.CreateSnapshotOpts) (snapshot *models.SnapshotModel, err error)
	All() (snapshots []*models.SnapshotModel, err error)
	Show(id string) (snapshot *models.SnapshotModel, err error)
	Update(id string, opts *options.UpdateSnapshotOpts) (snapshot *models.SnapshotModel, err error)
	AllByParams(opts *options.ListSnapshotOpts) (Snapshots []*models.SnapshotModel, err error)
	Delete(id string) error
}
