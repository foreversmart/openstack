package server

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/rackspace/gophercloud/pagination"
)

type ServerVolumer struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerVolumer(client ifaces.Openstacker) *ServerVolumer {
	return &ServerVolumer{
		Client: client,
	}
}

func (sv *ServerVolumer) All(id string) (volumeModels []*models.AttachVolumeModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := sv.Client.ComputeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result

	_, res.Err = client.Get(client.ServiceURL(ServersUrl, id, VolumeUrl), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	volumeModels, err = models.ExtractAttachVolumes(res)

	return
}

func (sv *ServerVolumer) Mount(id, volumeID string) (volume *models.AttachVolumeModel, err error) {
	if id == "" || volumeID == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := sv.Client.ComputeClient()
	if err != nil {
		return
	}

	opts := volumeattach.CreateOpts{
		VolumeID: volumeID,
	}

	result := volumeattach.Create(client, id, opts)
	err = result.Err
	if err != nil {
		return
	}

	return models.ExtractAttachVolume(result.Result)
}

func (sv *ServerVolumer) Unmount(id, volumeID string) error {
	if id == "" || volumeID == "" {
		return errors.ErrInvalidParams
	}

	var volumeAttachID string

	client, err := sv.Client.ComputeClient()
	if err != nil {
		return err
	}

	pager := volumeattach.List(client, id)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		vaList, err := volumeattach.ExtractVolumeAttachments(page)

		for _, value := range vaList {
			if value.VolumeID == volumeID {
				volumeAttachID = value.ID
				return false, err
			}
		}

		return true, err

	})

	if err != nil {
		return err
	}

	if volumeAttachID == "" {
		return errors.ErrNotFound
	}

	var res gophercloud.Result

	_, res.Err = client.Delete(client.ServiceURL(ServersUrl, id, VolumeUrl, volumeAttachID), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return res.Err
}
