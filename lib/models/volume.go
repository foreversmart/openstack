package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type VolumeModel struct {
	ID          string              `mapstructure:"id" json:"id"`
	Name        string              `mapstructure:"name" json:"name"`
	Size        int                 `mapstructure:"size" json:"size"`
	Status      string              `mapstructure:"status" json:"status"`
	Description string              `mapstructure:"description" json:"description"`
	VolumeType  string              `mapstructure:"volume_type" json:"volume_type"`
	SnapshotID  string              `mapstructure:"snapshot_id" json:"snapshot_id"`
	Bootable    string              `mapstructure:"bootable" json:"bootable"`
	Attachments []*VolumeAttachment `mapstructure:"attachments" json:"attachements"`
	CreatedAt   string              `mapstructure:"created_at" json:"created_at"`
}

type VolumeAttachment struct {
	ID           string `mapstructure:"id" json:"id"`
	ServerID     string `mapstructure:"server_id" json:"server_id"`
	AttachmentID string `mapstructure:"attachment_id" json:"attachment_id"`
	HostName     string `mapstructure:"host_name" json:"host_name"`
	VolumeID     string `mapstructure:"volume_id" json:"volume_id"`
	Device       string `mapstructure:"device" json:"device"`
}

func ExtractVolume(r gophercloud.Result) (*VolumeModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Volume *VolumeModel `mapstructure:"volume"`
	}
	err := mapstructure.Decode(r.Body, &response)
	return response.Volume, err
}

func ExtractVolumes(r gophercloud.Result) ([]*VolumeModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	return ExtractVolumesByBody(r.Body)
}

func ExtractVolumesByBody(body interface{}) ([]*VolumeModel, error) {
	var response struct {
		Volumes []*VolumeModel `mapstructure:"volumes"`
	}
	err := mapstructure.Decode(body, &response)
	return response.Volumes, err
}
