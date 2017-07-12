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
	VolumesUrl       = "volumes"
	VolumesDetailUrl = "volumes/detail"
	ActionUrl        = "action"
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

func (v *Volume) Create(opts *options.CreateVolumeOpts) (volume *models.VolumeModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(VolumesUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return models.ExtractVolume(res)
}

func (v *Volume) All() (volumeInfos []*models.VolumeModel, err error) {
	return v.AllByParams(nil)
}

func (v *Volume) AllByParams(opts *options.ListVolumeOpts) (volumes []*models.VolumeModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(VolumesDetailUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractVolumes(result)
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

func (v *Volume) Resize(id string, opts *options.ResizeVolumeOpts) error {
	if id == "" || !opts.IsValid() {
		return errors.ErrInvalidParams
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return err
	}

	// the response body is nil
	_, err = client.Post(client.ServiceURL(VolumesUrl, id, ActionUrl), opts.ToPayload(), nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return err
}

func (v *Volume) Reset(id string, opts *options.ResetVolumeOpts) error {
	if id == "" || !opts.IsValid() {
		return errors.ErrInvalidParams
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return err
	}

	// the response body is nil
	_, err = client.Post(client.ServiceURL(VolumesUrl, id, ActionUrl), opts.ToPayload(), nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return err
}

func (v *Volume) Update(id string, opts *options.UpdateVolumeOpts) (volume *models.VolumeModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.VolumeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Put(client.ServiceURL(VolumesUrl, id), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractVolume(res)
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
