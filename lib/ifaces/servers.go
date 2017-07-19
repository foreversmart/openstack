package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Serverser interface {
	Create(opts options.CreateServersOpts) (server *models.ServersModel, err error)
	All() (servers []*models.ServersModel, err error)
	AllByParams(opts *options.ListServersOpts) (servers []*models.ServersModel, err error)
	Show(id string) (server *models.ServersModel, err error)
	SearchByFixedIP(ip string) ([]*models.ServersModel, error)
	Update(id string, opts options.UpdateServersOpts) (server *models.ServersModel, err error)
	Delete(id string) error

	ServerActions
}

type ServerActions interface {
	ChangeAdminPassword(id, newPassword string) error                    // 重置系统管理员帐号密码
	AdminPassword(id string) (passwd string, err error)                  // 查询系统管理员帐号密码
	Start(id string) error                                               // 开机
	Reboot(id string) error                                              // 重启
	Shutdown(id string) error                                            // 软关闭
	Stop(id string) error                                                // 强制关闭
	CreateSnapshot(id, name string) (snapshotID string, err error)       // 创建主机快照
	BindPort(id, portID string) error                                    // 绑定网卡
	UnbindPort(id, portID string) error                                  // 解绑网卡
	Ports(id string) ([]string, error)                                   // 获取虚拟机的网卡
	BindKeys(id string, keys []string) error                             // 绑定key
	UnbindKeys(id string, keys []string) error                           // 解绑key
	MountVolume(id, volumeID string) error                               // 挂载磁盘
	UnmountVolume(id, volumeID string) error                             // 卸载磁盘
	ModifyFlavor(id, flavorID string) error                              // 修改主机配置
	Rebuild(id, imageID string) (server *models.ServersModel, err error) // 重置主机
	Vnc(id string) (string, error)                                       // vnc link
}
