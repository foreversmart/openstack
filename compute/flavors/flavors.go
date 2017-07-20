package flavors

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	FlavorUrl = "flavors"
)

type Flavors struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Flavors {
	return &Flavors{
		Client: client,
	}
}

func (f *Flavors) All() (flavors []*models.FlavorModel, err error) {
	return f.AllByParams(nil)
}
func (f *Flavors) AllByParams(opts *options.ListFlavorsOpts) (flavors []*models.FlavorModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := f.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(FlavorUrl+"/detail")+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFlavors(result)

}

func (f *Flavors) Create(opts *options.CreateFlavorOpts) (flavor *models.FlavorModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := f.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Post(client.ServiceURL(FlavorUrl), opts.ToPayLoad(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFlavor(result)
}

func (f *Flavors) Show(id string) (flavor *models.FlavorModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := f.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(FlavorUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFlavor(result)
}

func (f *Flavors) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := f.Client.ComputeClient()

	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(FlavorUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return err
}