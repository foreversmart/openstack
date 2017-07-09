package role

import (
	"github.com/kirk-enterprise/openstack/lib/errors"
	"github.com/kirk-enterprise/openstack/lib/ifaces"
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	RoleUrl    = "roles"
	DomainUrl  = "domains"
	GroupUrl   = "groups"
	UserUrl    = "users"
	ProjectUrl = "projects"
)

type Role struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Role {
	return &Role{
		Client: client,
	}
}

func (r *Role) All() (roles []*models.RoleModel, err error) {
	return r.AllByParams(nil)
}

func (r *Role) AllByParams(opts *options.ListRoleOpts) (roles []*models.RoleModel, err error) {
	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(RoleUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRoles(result)

}

func (r *Role) Create(opts *options.CreateRoleOpts) (role *models.RoleModel, err error) {
	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(RoleUrl), opts.ToPayLoad(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractRole(result)
}

func (r *Role) Show(id string) (role *models.RoleModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(RoleUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRole(result)
}

func (r *Role) Update(id string, opts *options.UpdateRoleOpts) (role *models.RoleModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Patch(client.ServiceURL(RoleUrl, id), opts.ToPayLoad(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRole(result)
}

func (r *Role) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(RoleUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
