package role

import (
	"net/http"

	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/rackspace/gophercloud"
)

type AbstractRole struct {
	Client ifaces.Openstacker

	resource string
	scope    string

	_ bool
}

func (role *AbstractRole) ensureResourceAndScope() {
	switch role.resource {
	case DomainUrl, ProjectUrl:
	// ignore

	default:
		panic("invalid resource")
	}

	switch role.scope {
	case GroupUrl, UserUrl:
	// ignore

	default:
		panic("invalid scope")
	}
}

func NewAbstractRole(client ifaces.Openstacker, resource, scope string) *AbstractRole {
	ar := &AbstractRole{
		Client:   client,
		resource: resource,
		scope:    scope,
	}

	ar.ensureResourceAndScope()

	return ar
}

func (role *AbstractRole) All(resourceID, scopeID string) (roles []*models.RoleModel, err error) {
	if resourceID == "" || scopeID == "" {
		return nil, errors.ErrInvalidParams
	}

	role.ensureResourceAndScope()

	client, err := role.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(role.resource, resourceID, role.scope, scopeID, RoleUrl), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractRoles(result)
}

func (role *AbstractRole) Create(resourceID, scopeID, roleID string) (err error) {
	if resourceID == "" || scopeID == "" || roleID == "" {
		return errors.ErrInvalidParams
	}

	role.ensureResourceAndScope()

	client, err := role.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Put(client.ServiceURL(role.resource, resourceID, role.scope, scopeID, RoleUrl, roleID), nil, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

func (role *AbstractRole) HasRole(resourceID, scopeID, roleID string) (ok bool) {
	if resourceID == "" || scopeID == "" || roleID == "" {
		return false
	}

	role.ensureResourceAndScope()

	client, err := role.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	response, _ := client.Request(http.MethodHead, client.ServiceURL(role.resource, resourceID, role.scope, scopeID, RoleUrl, roleID), gophercloud.RequestOpts{
		OkCodes: []int{204, 404},
	})

	if response != nil {
		switch response.StatusCode {
		case 204:
			ok = true
		}
	}

	return
}

func (role *AbstractRole) Delete(resourceID, scopeID, roleID string) (err error) {
	if resourceID == "" || scopeID == "" || roleID == "" {
		return errors.ErrInvalidParams
	}

	role.ensureResourceAndScope()

	client, err := role.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(role.resource, resourceID, role.scope, scopeID, RoleUrl, roleID), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
