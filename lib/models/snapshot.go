package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type SnapshotModel struct {
	ID          string `mapstructure:"id" json:"id"`
	Status      string `mapstructure:"status" json:"status"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
	VolumeID    string `mapstructure:"volume_id" json:"volume_id"`
	Size        int    `mapstructure:"size" json:"size"`
	CreatedAt   string `mapstructure:"created_at" json:"created_at"`
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

func ExtractSnapshotsByBody(body interface{}) ([]*SnapshotModel, error) {
	var resp struct {
		Snapshots []*SnapshotModel `mapstructure:"snapshots"`
	}

	err := mapstructure.Decode(body, &resp)
	return resp.Snapshots, err
}

func ExtractSnapshots(r gophercloud.Result) ([]*SnapshotModel, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	return ExtractSnapshotsByBody(r.Body)
}
