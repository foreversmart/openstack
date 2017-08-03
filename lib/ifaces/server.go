package ifaces

import (
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
)

type Server interface {
	Create(opts *options.CreateServerOpts) (server *models.ServerModel, err error)
	All() (servers []*models.ServerModel, err error)
	AllByParams(opts *options.ListServersOpts) (servers []*models.ServerModel, err error)
	Show(id string) (server *models.ServerModel, err error)
	Update(id string, opts options.UpdateServersOpts) (server *models.ServerModel, err error)
	Delete(id string) error
}

type ServerManager interface {
	ChangeAdminPassword(serverID, newPassword string) error                                           // 重置系统管理员帐号密码
	AdminPassword(serverID string) (passwd string, err error)                                         // 查询系统管理员帐号密码
	Start(serverID string) error                                                                      // 开机
	Stop(serverID string) error                                                                       // 强制关闭
	Reboot(serverID string) error                                                                     // 重启
	Shutdown(serverID string) error                                                                   // 软关闭
	Rebuild(serverID string, opts *options.RebuildServerOpts) (server *models.ServerModel, err error) // 重置主机
	Resize(serverID, flavorID string) error                                                           // 修改主机配置
	Vnc(serverID string) (url string, err error)                                                      // vnc link

	SearchByFixedIP(ip string) (servers []*models.ServerModel, err error)
}

type ServerPorter interface {
	All(serverID string) (ports []*models.AttachPortModel, err error) // 获取虚拟机的网卡
	Bind(serverID, portID string) error                               // 绑定网卡
	Unbind(serverID, portID string) error                             // 解绑网卡
}

type ServerKeyer interface {
	Bind(serverID string, keys []string) error   // 绑定key
	Unbind(serverID string, keys []string) error // 解绑key
}

type ServerVolumer interface {
	All(serverID string) (volumes []*models.AttachVolumeModel, err error)
	Mount(serverID, volumeID string) (volume *models.AttachVolumeModel, err error) // 挂载磁盘
	Unmount(serverID, volumeID string) error                                       // 卸载磁盘
}

type ServerImager interface {
	Create(serverID, imageName string) (imageID string, err error) // 创建主机快照
}
