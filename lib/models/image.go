package models

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type ImageType string
type ImageVisibility string

const (
	// image type
	ImageTypeSnapshot ImageType = "snapshot"
	ImageTypeNormal   ImageType = "normal"
	ImageTypeAll      ImageType = "all"

	// image visibility
	ImageVisibilityPublic  ImageVisibility = "public"
	ImageVisibilityShared  ImageVisibility = "shared"
	ImageVisibilityPrivate ImageVisibility = "private"
	ImageVisibilityAll     ImageVisibility = "all"

	WindowsImageLabel = "Windows"

	OsTypeTagMark   string = "os_type:"
	OsFamilyTagMark string = "os_family:"
)

func (i ImageType) IsValid() bool {
	switch i {
	case ImageTypeSnapshot, ImageTypeNormal, ImageTypeAll:
		return true
	}

	return false
}

func (i ImageVisibility) IsValid() bool {
	switch i {
	case ImageVisibilityPublic, ImageVisibilityPrivate, ImageVisibilityShared, ImageVisibilityAll:
		return true
	}

	return false
}

type ImageModel struct {
	ID         string   `json:"id" mapstructure:"id"`
	Name       string   `json:"name" mapstructure:"name"`
	Status     string   `json:"status" mapstructure:"status"`
	Visibility string   `json:"visibility" mapstructure:"visibility"`
	Size       int      `json:"size" mapstructure:"size"`
	Tags       []string `json:"tags" mapstructure:"tags"`
	Self       string   `json:"self" mapstructure:"self"`
	File       string   `json:"file" mapstructure:"file"`
	Schema     string   `json:"schema" mapstructure:"schema"`

	// DiskFormat is the format of the disk.
	// If set, valid values are ami, ari, aki, vhd, vmdk, raw, qcow2, vdi, and iso.
	DiskFormat string `json:"disk_format" mapstructure:"disk_format"`
	// MinDiskGigabytes is the amount of disk space in GB that is required to boot the image.
	MinDisk int `json:"min_disk" mapstructure:"min_disk"`
	// MinRAMMegabytes [optional] is the amount of RAM in MB that is required to boot the image.
	MinRAM    int    `json:"min_ram" mapstructure:"min_ram"`
	UpdatedAt string `json:"updated_at" mapstructure:"updated_at"`
	CreatedAt string `json:"created_at" mapstructure:"created_at"`
}

func ExtractImages(result gophercloud.Result) (images []*ImageModel, err error) {
	if result.Err != nil {
		err = result.Err
		return
	}

	var response struct {
		Images []*ImageModel `mapstructure:"images"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		images = response.Images
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
