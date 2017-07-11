package volume

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
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
