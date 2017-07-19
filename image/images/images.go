package images

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
)

type Images struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Images {
	return &Images{
		Client: client,
	}
}

func (i *Images) Create(opts *options.CreateImagesOpts) (image *models.ImageModel, err error) {
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

func (i *Images) All() (images []*models.ImageModel, err error) {
	return i.AllByParams(nil)
}

func (i *Images) AllByParams(opts *options.ListImagesOpts) (imageModels []*models.ImageModel, err error) {
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

func (i *Images) Show(id string) (image *models.ImageModel, err error) {
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

func (i *Images) Update(id string, opts *options.UpdateImagesOpts) (imageModel *models.ImageModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Patch(client.ServiceURL(ImagesUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractImage(result)
}

func (i *Images) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return err
	}

	return images.Delete(client, id).Err
}

type ImageReadSeeker struct {
	io.Reader
}

func (i ImageReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (i *Images) Upload(id string, data io.Reader) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return err
	}

	body := &ImageReadSeeker{
		Reader: data,
	}

	_, err = client.Request("PUT", client.ServiceURL(ImagesUrl, id, ImageUploadUrl), gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/octet-stream"},
		OkCodes:     []int{204},
		RawBody:     body,
	})

	return err
}

func (i *Images) Download(id string) (data io.Reader, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := i.Client.ImageClient()
	if err != nil {
		return
	}

	return imageservice.Download(client, id).Extract()
}
