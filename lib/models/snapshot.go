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

func ExtractSnapshot(r gophercloud.Result) (snapshot *SnapshotModel, err error) {
	if r.Err != nil {
		err = r.Err
		return
	}

	var response struct {
		Snapshot *SnapshotModel `mapstructure:"snapshot"`
	}
	err = mapstructure.Decode(r.Body, &response)
	snapshot = response.Snapshot
	return
}

func ExtractSnapshotsByBody(body interface{}) (snapshots []*SnapshotModel, err error) {
	var resp struct {
		Snapshots []*SnapshotModel `mapstructure:"snapshots"`
	}

	err = mapstructure.Decode(body, &resp)
	snapshots = resp.Snapshots
	return
}

func ExtractSnapshots(r gophercloud.Result) (snapshots []*SnapshotModel, err error) {
	if r.Err != nil {
		err = r.Err
		return
	}
	return ExtractSnapshotsByBody(r.Body)
}
