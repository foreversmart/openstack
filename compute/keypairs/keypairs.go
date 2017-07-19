package keypairs

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/keypairs"
)

const (
	KeyUrl = "os-keypairs"
)

type Keypairs struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Keypairs {
	return &Keypairs{
		Client: client,
	}
}

func (k *Keypairs) Create(opts *options.CreateKeypairOpts) (keypair *models.KeypairModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := k.Client.ComputeClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(KeyUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return models.ExtractKeypair(res)
}

func (f *Keypairs) All() (KeypairModels []*models.KeypairModel, err error) {
	return f.AllByParams(nil)
}

func (f *Keypairs) AllByParams(opts *options.ListKeypairOpts) (keypairModels []*models.KeypairModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := f.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(KeyUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractKeypairs(result)
}

func (k *Keypairs) Show(name string) (keypairModel *models.KeypairModel, err error) {
	if name == "" {
		err = errors.ErrInvalidParams
		return
	}
	client, err := k.Client.ComputerClient()
	if err != nil {
		return
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(KeyUrl, name), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractKeypair(result)
}

func (k *Keypairs) Delete(name string) error {
	if name == "" {
		return errors.ErrInvalidParams
	}

	client, err := k.Client.ComputerClient()
	if err != nil {
		return err
	}

	err = keypairs.Delete(client, name).ExtractErr()

	return err
}
