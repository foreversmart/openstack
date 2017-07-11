package volume

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/blockstorage/v2/volumes"
)

const (
	VolumesUrl = "volumes"
)

type Volume struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Volume {
	return &Volume{
		Client: client,
	}
}

func (v *Volume) All() (volumeInfos []*models.VolumeModel, err error) {
	client, err := v.Client.VolumeClient()
	if err != nil {
		return nil, err
	}

	page, err := volumes.List(client, volumes.ListOpts{}).AllPages()
	if err != nil {
		return nil, err
	}

	return models.ExtractVolumesFromBody(page.GetBody())
}

func (v *Volume) AllByParams(opts *options.ListVolumeOpts) (volumes []*models.VolumeModel, err error) {
	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(VolumesUrl)+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	volumes, err = models.ExtractVolumes(result)
	return
}

func (v *Volume) Create(param *options.CreateVolumeOpts) (id string, err error) {
	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	if !param.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	// snapshot'volume type should be same as volume type
	opts := volumes.CreateOpts{}
	if param.Name != nil {
		opts.Name = *param.Name
	}
	if param.Description != nil {
		opts.Description = *param.Description
	}
	if param.Size != nil {
		opts.Size = *param.Size
	}
	if param.SnapshotID != nil {
		opts.SnapshotID = *param.SnapshotID
	}
	if param.VolumeType != nil {
		opts.VolumeType = *param.VolumeType
	}

	volumeInfo, err := volumes.Create(client, opts).Extract()
	if err != nil {
		return "", err
	}

	return volumeInfo.ID, err
}
