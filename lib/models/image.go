package models

import (
	"io"

	"github.com/mitchellh/mapstructure"
	"github.com/qbox/openstack-golang-sdk/lib/enums"
	"github.com/rackspace/gophercloud"
)

type ImageModel struct {
	ID         string                `mapstructure:"id" json:"id"`
	Name       string                `mapstructure:"name" json:"name"`
	Status     string                `mapstructure:"status" json:"status"`
	Visibility enums.ImageVisibility `mapstructure:"visibility" json:"visibility"`
	Size       int                   `mapstructure:"size" json:"size"`
	Tags       []string              `mapstructure:"tags" json:"tags"`
	Self       string                `mapstructure:"self" json:"self"`
	File       string                `mapstructure:"file" json:"file"`
	Schema     string                `mapstructure:"schema" json:"schema"`

	ImageType string `mapstructure:"image_type" json:"image_type"`

	// DiskFormat is the format of the disk.
	// If set, valid values are ami, ari, aki, vhd, vmdk, raw, qcow2, vdi, and iso.
	DiskFormat string `mapstructure:"disk_format" json:"disk_format"`
	// MinDiskGigabytes is the amount of disk space in GB that is required to boot the image.
	MinDisk int `mapstructure:"min_disk" json:"min_disk"`
	// MinRAMMegabytes [optional] is the amount of RAM in MB that is required to boot the image.
	MinRAM    int    `mapstructure:"min_ram" json:"min_ram"`
	UpdatedAt string `mapstructure:"updated_at" json:"updated_at"`
	CreatedAt string `mapstructure:"created_at" json:"created_at"`
}

func ExtractImages(result gophercloud.Result) (images []*ImageModel, hasNext bool, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	var response struct {
		Images []*ImageModel `mapstructure:"images"`
		Next string `mapstructure:"next"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		images = response.Images
	}

	if response.Next == "" {
		hasNext = true
	}

	return
}

func ExtractImage(res gophercloud.Result) (image *ImageModel, err error) {
	if res.Err != nil {
		err = res.Err
		return
	}

	err = mapstructure.Decode(res.Body, &image)

	return
}

type ImageReadSeeker struct {
	io.Reader
}

func (i ImageReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}
