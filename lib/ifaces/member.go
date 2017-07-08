package ifaces

import (
	"github.com/kirk-enterprise/openstack/lib/models"
	"github.com/kirk-enterprise/openstack/lib/options"
)

type Member interface {
	All(poolId string) ([]*models.MemberModel, error)
	Query(poolId, id string) (*models.MemberModel, error)
	Create(poolId string, member *options.MemberCreateOpt) (*models.MemberModel, error)
	Modify(poolId, id string, member *options.MemberUpdateOpt) error
	Delete(poolId, id string) (err error)
}

type MemberHelper interface {
	AppendVmInfos(vmInfos []*models.VmModel, memberInfos []*models.MemberModel) (err error)
	AppendVmInfo(vmInfo *models.VmModel, membersInfo *models.MemberModel) (err error)
}
