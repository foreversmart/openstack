package snapshot

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/blockstorage/v1/snapshots"
)

const (
	SnapshotUrl = "snapshots"
)

type Snapshot struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Snapshot {
	return &Snapshot{
		Client: client,
	}
}

func (s *Snapshot) Create(opts options.CreateSnapshotOpts) (snapshot *models.SnapshotModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := s.Client.VolumeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(SnapshotUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	snapshot, err = models.ExtractSnapshot(res)

	return
}

func (s *Snapshot) All() (snapshotModels []*models.SnapshotModel, err error) {
	client, err := s.Client.VolumeClient()
	if err != nil {
		return
	}

	page, err := snapshots.List(client, snapshots.ListOpts{}).AllPages()
	if err != nil {
		return nil, err
	}
	return models.ExtractSnapshotsByBody(page.GetBody())
}

func (s *Snapshot) Show(id string) (snapshot *models.SnapshotModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := s.Client.VolumeClient()
	if err != nil {
		return nil, err
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(SnapshotUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return models.ExtractSnapshot(result)
}

func (s *Snapshot) Update(snapshotID string, opts *options.UpdateSnapshotOpts) (snapshot *models.SnapshotModel, err error) {
	if snapshotID == "" || opts == nil {
		err = errors.ErrInvalidParams
		return
	}

	client, err := s.Client.VolumeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Put(client.ServiceURL(SnapshotUrl, snapshotID), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	snapshot, err = models.ExtractSnapshot(res)

	return
}

func (s *Snapshot) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := s.Client.VolumeClient()
	if err != nil {
		return err
	}

	return snapshots.Delete(client, id).ExtractErr()
}
