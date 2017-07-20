package server

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type ServerVolumer struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerVolumer(client ifaces.Openstacker) *ServerVolumer {
	return &ServerVolumer{
		Client: client,
	}
}

// func (ser *Servers) MountVolume(id, volumeID string) error {
// 	if id == "" || volumeID == "" {
// 		return errors.ErrInvalidParams
// 	}

// 	client, err := ser.Client.ComputeClient()
// 	if err != nil {
// 		return err
// 	}

// 	opts := volumeattach.CreateOpts{
// 		VolumeID: volumeID,
// 	}

// 	return volumeattach.Create(client, id, opts).Err
// }

// func (ser *Servers) UnmountVolume(id, volumeID string) error {
// 	if id == "" || volumeID == "" {
// 		return errors.ErrInvalidParams
// 	}

// 	var volumeAttachID string

// 	client, err := ser.Client.ComputeClient()
// 	if err != nil {
// 		return err
// 	}

// 	pager := volumeattach.List(client, id)
// 	err = pager.EachPage(func(page pagination.Page) (bool, error) {
// 		vaList, err := volumeattach.ExtractVolumeAttachments(page)

// 		for _, value := range vaList {
// 			if value.VolumeID == volumeID {
// 				volumeAttachID = value.ID
// 				return false, err
// 			}
// 		}

// 		return true, err

// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if volumeAttachID == "" {
// 		return errors.ErrNotFound
// 	}

// 	return volumeattach.Delete(client, id, volumeAttachID).Err
// }
