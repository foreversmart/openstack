package region

import (
	"github.com/kirk-enterprise/openstack/lib/errors"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	RegionUrl = "regions"
)

type Region struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Region {
	return &Region{
		Client: client,
	}
}

func (r *Region) All() (regions []*models.RegionModel, err error) {
	return r.AllByParams(nil)
}

func (r *Region) AllByParams(opts *options.ListRegionOpts) (regions []*models.RegionModel, err error) {
	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(RegionUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRegions(result)
}

func (r *Region) Create(opts *options.CreateRegionOpts) (region *models.RegionModel, err error) {
	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(RegionUrl), opts.ToPayLoad(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractRegion(result)
}

func (r *Region) Update(regionID string, opts *options.UpdateRegionOpts) (region *models.RegionModel, err error) {
	if regionID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Patch(client.ServiceURL(RegionUrl, regionID), opts.ToPayLoad(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRegion(result)
}

func (r *Region) Show(regionID string) (region *models.RegionModel, err error) {
	if regionID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(RegionUrl, regionID), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRegion(result)
}

func (r *Region) Delete(regionID string) (err error) {
	if regionID == "" {
		return errors.ErrInvalidParams
	}

	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(RegionUrl, regionID), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
