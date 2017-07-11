package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type SnapshotModel struct {
	ID          string `json:"id" mapstructure:"id"`
	Status      string `json:"status" mapstructure:"status"`
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	VolumeID    string `json:"volume_id" mapstructure:"volume_id"`
	Size        int    `json:"size" mapstructure:"size"`
	CreatedAt   string `json:"created_at" mapstructure:"created_at"`
}

func ExtractSnapshot(r gophercloud.Result) (*SnapshotModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Snapshots *SnapshotModel `mapstructure:"snapshot"`
	}
	err := mapstructure.Decode(r.Body, &response)
	return response.Snapshots, err
}

func ExtractSnapshots(r gophercloud.Result) ([]*SnapshotModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Snapshots []*SnapshotModel `mapstructure:"snapshots"`
	}
	err := mapstructure.Decode(r.Body, &response)
	return response.Snapshots, err
}
