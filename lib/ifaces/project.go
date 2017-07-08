package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Projecter interface {
	All() (projs []*models.ProjectModel, err error)
	Query(projectId string) (project *models.ProjectModel, err error)
}

//Identity V3 API
type ProjecterV3 interface {
	All(opts options.ListProjectOpts) (projects []*models.ProjectModel, err error)
	Create(opts options.CreateProjectOpts) (project *models.ProjectModel, err error)
	Show(projectID string) (project *models.ProjectModel, err error)
	Update(projectID string, opts options.UpdateProjectOpts) (project *models.ProjectModel, err error)
	Delete(projectID string) (err error)
}
