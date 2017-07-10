package project

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/errors"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/models"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
)

const (
	ProjectUrl = "projects"
)

type Project struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Project {
	return &Project{
		Client: client,
	}
}

func (proj *Project) All() (projs []*models.ProjectModel, err error) {
	return proj.AllByParams(nil)
}

func (proj *Project) AllByParams(opts *options.ListProjectOpts) (projects []*models.ProjectModel, err error) {
	client, err := proj.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ProjectUrl)+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractProjects(result)
}

func (proj *Project) Create(opts options.CreateProjectOpts) (project *models.ProjectModel, err error) {
	client, err := proj.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(ProjectUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractProject(result)
}

func (proj *Project) Show(projectID string) (project *models.ProjectModel, err error) {
	if projectID == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := proj.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ProjectUrl, projectID), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractProject(result)
}

func (proj *Project) Update(projectID string, opts options.UpdateProjectOpts) (project *models.ProjectModel, err error) {
	client, err := proj.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Patch(client.ServiceURL(ProjectUrl, projectID), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractProject(result)
}

func (proj *Project) Delete(projectID string) (err error) {
	if projectID == "" {
		return errors.ErrInvalidParams
	}

	client, err := proj.Client.AdminIdentityClientV3()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(ProjectUrl, projectID), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
