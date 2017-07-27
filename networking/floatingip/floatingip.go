package floatingip

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	FloatingIPUrl = "floatingips"
)

type FloatingIP struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *FloatingIP {
	return &FloatingIP{
		Client: client,
	}
}

func (f *FloatingIP) Create(opts *options.CreateFloatingIPOpts) (ip *models.FloatingIPModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := f.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(FloatingIPUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractFloatingIP(res)
}

func (f *FloatingIP) CreateWithProvider(name, provider string, opts *options.CreateFloatingIPOpts) (ip *models.FloatingIPModel, err error) {
	return f.Create(opts)
}

func (f *FloatingIP) All() (ips []*models.FloatingIPModel, err error) {
	return f.AllByParams(nil)
}

func (f *FloatingIP) AllByParams(opts *options.ListFloatingIPOpts) (ips []*models.FloatingIPModel, err error) {
	if !opts.IsValid() {
		return nil, errors.ErrInvalidParams
	}

	client, err := f.Client.NetworkClient()
	if err != nil {
		return nil, err
	}

	var result gophercloud.Result
	_, result.Err = client.Get(client.ServiceURL(FloatingIPUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFloatingIPs(result)
}

func (f *FloatingIP) Show(id string) (ip *models.FloatingIPModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := f.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Get(client.ServiceURL(FloatingIPUrl, id), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFloatingIP(res)
}

func (f *FloatingIP) Update(id string, opts *options.UpdateFloatingIPOpts) (ip *models.FloatingIPModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := f.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Put(client.ServiceURL(FloatingIPUrl, id), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFloatingIP(res)
}

func (f *FloatingIP) Resize(id string, rateLimit int) error {
	return errors.ErrNotImplemented
}

func (f *FloatingIP) Delete(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := f.Client.NetworkClient()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(FloatingIPUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
