package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Projecter interface {
	All() (projects []*models.ProjectModel, err error)
	AllByParams(opts options.ListProjectOpts) (projects []*models.ProjectModel, err error)
	Create(opts options.CreateProjectOpts) (project *models.ProjectModel, err error)
	Show(projectID string) (project *models.ProjectModel, err error)
	Update(projectID string, opts options.UpdateProjectOpts) (project *models.ProjectModel, err error)
	Delete(projectID string) (err error)
}
