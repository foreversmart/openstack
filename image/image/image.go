package image

import (
	"io"

	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
	imageservice "github.com/rackspace/gophercloud/openstack/imageservice/v2/images"
)

const (
	ImagesUrl      = "images"
	ImageUploadUrl = "file"

	DefaultLimit = 50
)

type Image struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Image {
	return &Image{
		Client: client,
	}
}

func (i *Image) Create(opts *options.CreateImagesOpts) (image *models.ImageModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Post(client.ServiceURL(ImagesUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractImage(result)
}

func (i *Image) All() (images []*models.ImageModel, err error) {
	images = make([]*models.ImageModel, 0, DefaultLimit)
	for {
		opts := &options.ListImagesOpts{
			Limit: options.Int(DefaultLimit),
		}

		tempImages, hasNext, err := i.AllByParams(opts)

		if err != nil {
			return images, err
		}

		images = append(images, tempImages...)

		if !hasNext {
			break
		}

		opts.Marker = &tempImages[len(tempImages)-1].ID
	}

	return
}

// Note: image list has a default limit size
func (i *Image) AllByParams(opts *options.ListImagesOpts) (imageModels []*models.ImageModel, hasNext bool, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ImagesUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractImages(result)
}

func (i *Image) Show(id string) (image *models.ImageModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(ImagesUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractImage(result)
}

func (i *Image) Update(id string, opts *options.UpdateImagesOpts) (imageModel *models.ImageModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	res := imageservice.Update(client, id, opts.ToPatches())
	return models.ExtractImage(res.Result)
}

func (i *Image) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return err
	}

	return images.Delete(client, id).Err
}

func (i *Image) Upload(id string, data io.Reader) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return err
	}

	body := &models.ImageReadSeeker{
		Reader: data,
	}

	_, err = client.Request("PUT", client.ServiceURL(ImagesUrl, id, ImageUploadUrl), gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/octet-stream"},
		OkCodes:     []int{204},
		RawBody:     body,
	})

	return err
}

func (i *Image) Download(id string) (data io.Reader, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	return imageservice.Download(client, id).Extract()
}
