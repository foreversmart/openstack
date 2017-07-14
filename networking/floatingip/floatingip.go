package floatingip

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	FloatingIPUrl = "floatingips"
)

type FloatingIP struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Volume {
	return &FloatingIP{
		Client: client,
	}
}

func (f *FloatingIP) Create(opts *options.CreateFloatingIPOpts) (ip *models.FloatingIPModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.Post(client.ServiceURL(FloatingIPUrl), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractFloatingIP(res)
}

func (f *FloatingIP) All() (ips []*models.FloatingIPModel, err error) {
	client, err := v.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.GET(client.ServiceURL(FloatingIPUrl), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFloatingIPs(res)
}

func (f *FloatingIP) Show(id string) (ip *models.FloatingIPModel, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.GET(client.ServiceURL(FloatingIPUrl, id), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFloatingIP(res)
}

func (f *FloatingIP) Update(id string, opts *options.UpdateFloatingIPOpts) (ip *models.FloatingIPModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.PUT(client.ServiceURL(FloatingIPUrl, id), opts.ToPayload(), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractFloatingIP(res)
}

func (f *FloatingIP) Delete(id string) error {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := v.Client.NetworkClient()
	if err != nil {
		return
	}

	var res gophercloud.Result
	_, res.Err = client.DELETE(client.ServiceURL(FloatingIPUrl, id), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return res.Err
}
