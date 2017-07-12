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
	ActionUrl  = "action"
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

func (v *Volume) Create(param *options.CreateVolumeOpts) (volume *models.VolumeModel, err error) {
	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	if !param.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(VolumesUrl), param.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	volume, err = models.ExtractVolume(res)

	return
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

	return models.ExtractVolumesByBody(page.GetBody())
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

func (v *Volume) Show(id string) (volume *models.VolumeModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(VolumesUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractVolume(result)
}

func (v *Volume) Resize(id string, newSize int) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return err
	}

	reqBody := options.ToVolumeActionResizeMap(&newSize)

	// the response body is nil
	_, err = client.Post(client.ServiceURL(VolumesUrl, id, ActionUrl), reqBody, nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return err
}

func (v *Volume) Update(id string, param *options.UpdateVolumeOpts) (volume *models.VolumeModel, err error) {
	if id == "" || param == nil {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Put(client.ServiceURL(VolumesUrl, id), param.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	volume, err = models.ExtractVolume(res)

	return
}

func (v *Volume) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return err
	}

	return volumes.Delete(client, id).ExtractErr()
}
