package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type VolumeModel struct {
	ID          string              `json:"id" mapstructure:"id"`
	Name        string              `json:"name" mapstructure:"name"`
	Size        int                 `json:"size" mapstructure:"size"`
	Status      string              `json:"status" mapstructure:"status"`
	Description string              `json:"description" mapstructure:"description"`
	VolumeType  string              `json:"volume_type" mapstructure:"volume_type"`
	SnapshotID  string              `json:"snapshot_id" mapstructure:"snapshot_id"`
	Bootable    string              `json:"bootable" mapstructure:"bootable"`
	Attachments []*VolumeAttachment `json:"-" mapstructure:"attachments"` //挂载的虚拟机id
	CreatedAt   string              `json:"created_at" mapstructure:"created_at"`
}

type VolumeAttachment struct {
	ID           string `json:"id" mapstructure:"id"`
	ServerID     string `json:"server_id" mapstructure:"server_id"`
	AttachmentID string `json:"attachment_id" mapstructure:"attachment_id"`
	HostName     string `json:"host_name" mapstructure:"host_name"`
	VolumeID     string `json:"volume_id" mapstructure:"volume_id"`
	Device       string `json:"device" mapstructure:"device"`
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

	var response struct {
		Volumes []*VolumeModel `mapstructure:"volumes" json:"volumes"`
	}
	err := mapstructure.Decode(r.Body, &response)
	return response.Volumes, err
}

type VolumeSnapshotModel struct {
	ID          string `json:"id" mapstructure:"id"`
	Status      string `json:"status" mapstructure:"status"`
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	VolumeID    string `json:"volume_id" mapstructure:"volume_id"`
	VolumeName  string `json:"volume_name" mapstructure:"volume_name"`
	Size        int    `json:"size" mapstructure:"size"`
	CreatedAt   string `json:"created_at" mapstructure:"created_at"`
}

func ExtractSnapshot(r gophercloud.Result) (*VolumeSnapshotModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Snapshots *VolumeSnapshotModel `mapstructure:"snapshot"`
	}
	err := mapstructure.Decode(r.Body, &response)
	return response.Snapshots, err
}

func ExtractSnapshots(r gophercloud.Result) ([]*VolumeSnapshotModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Snapshots []*VolumeSnapshotModel `mapstructure:"snapshots"`
	}
	err := mapstructure.Decode(r.Body, &response)
	return response.Snapshots, err
}
