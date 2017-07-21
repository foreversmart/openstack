package role

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	RoleAssignmentUrl = "role_assignments"
)

type RoleAssignment struct {
	Client ifaces.Openstacker

	_ bool
}

func NewAssignment(client ifaces.Openstacker) *RoleAssignment {
	return &RoleAssignment{
		Client: client,
	}
}

func (r *RoleAssignment) All() (assignments []*models.RoleAssignmentModel, err error) {
	return r.AllByParams(nil)
}

func (r *RoleAssignment) AllByParams(opts *options.ListRoleAssignmentOpts) (assignments []*models.RoleAssignmentModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}
	client, err := r.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(RoleAssignmentUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRoleAssignments(result)
}
