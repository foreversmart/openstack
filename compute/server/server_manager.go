package server

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

type ServerManager struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerManager(client ifaces.Openstacker) *ServerManager {
	return &ServerManager{
		Client: client,
	}
}

func (sm *ServerManager) SearchByFixedIP(ip string) (servers []*models.ServerModel, err error) {
	if ip == "" {
		err = errors.ErrInvalidParams
		return
	}

	serverClient := New(sm.Client)

	return serverClient.AllByParams(&options.ListServersOpts{
		Ip:         options.String(ip),
		AllTenants: options.Int(1),
	})
}

func (sm *ServerManager) ChangeAdminPassword(id, newPassword string) error {
	if id == "" || newPassword == "" {
		return errors.ErrInvalidParams
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return err
	}

	return servers.ChangeAdminPassword(client, id, newPassword).ExtractErr()
}

func (sm *ServerManager) AdminPassword(id string) (password string, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return
	}

	res, err := servers.Get(client, id).Extract()
	if err != nil {
		return
	}
	if res.Metadata == nil {
		err = errors.ErrNotFound
		return
	}

	password = res.Metadata["admin_pass"].(string)
	return
}

func (sm *ServerManager) Start(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return err
	}

	return startstop.Start(client, id).ExtractErr()
}

func (sm *ServerManager) Reboot(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return err
	}

	return servers.Reboot(client, id, servers.SoftReboot).ExtractErr()
}

// TODO add own shutdown impl (openstack sdk has no shutdown method)
func (sm *ServerManager) Shutdown(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	return sm.Stop(id)
}

func (sm *ServerManager) Stop(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return err
	}

	return startstop.Stop(client, id).ExtractErr()
}

func (sm *ServerManager) Resize(id, flavorID string) error {
	if id == "" || flavorID == "" {
		return errors.ErrInvalidParams
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return err
	}

	opts := servers.ResizeOpts{
		FlavorRef: flavorID,
	}

	err = servers.Resize(client, id, opts).Err
	if err != nil {
		return err
	}

	err = gophercloud.WaitFor(30, func() (bool, error) {
		var res gophercloud.Result

		_, res.Err = client.Get(client.ServiceURL(ServersUrl, id), &res.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 201},
		})

		vmInfo, err := models.ExtractServer(res)
		if err != nil {
			return false, err
		}

		switch vmInfo.Status {
		case "VERIFY_RESIZE":
			response := servers.ConfirmResize(client, id)

			return true, response.ExtractErr()

		case "ACTIVE":
			return true, nil

		default:
			return false, nil
		}
	})

	return err
}

// TODO: after finish image package to do this then
func (sm *ServerManager) Rebuild(id, imageID string) (serverModel *models.ServerModel, err error) {
	if id == "" || imageID == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := sm.Client.ComputeClient()
	if err != nil {
		return
	}

	serverClient := New(sm.Client)

	vm, err := serverClient.Show(id)
	if err != nil {
		return
	}

	// // fetch base image id
	// baseImageId := imageID
	// imager := images.New(ser.Client)
	// image, err := imager.Show(imageID)
	// if err != nil {
	// 	return err
	// }

	// if image.BaseImageID != "" {
	// 	baseImageId = image.BaseImageId
	// }
	// 这块逻辑应该实现在controller层

	opts := options.RebuildServerOpts{
		Name:    vm.Name,
		ImageID: imageID,
		Metadata: map[string]string{
			"hypervisor_type": "qemu",
		},
	}

	res := servers.Rebuild(client, id, opts)
	if res.Err != nil {
		err = res.Err
		return
	}

	return models.ExtractServer(res.Result)

	// // update server base image id metadata
	// metaDataOpts := params.UpdateMetadataOpts{
	// 	Metadata: map[string]string{
	// 		"base_image_id": baseImageId,
	// 		"admin_pass":    res.AdminPass,
	// 	},
	// }

	// err = servers.UpdateMetadata(client, id, metaDataOpts).Err
	// return
	// 这段逻辑应该实现在controller层
}

func (sm *ServerManager) Vnc(id string) (vncURL string, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := openstack.NewComputeV2(sm.Client.ProviderClient(), gophercloud.EndpointOpts{})
	if err != nil {
		return
	}

	payload := map[string]interface{}{
		"os-getVNCConsole": map[string]interface{}{
			"type": "novnc",
		},
	}

	var result gophercloud.Result

	_, result.Err = client.Post(client.ServiceURL(ServersUrl, id, ActionUrl), payload, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	if result.Err != nil {
		err = result.Err
		return
	}

	data, err := models.ExtractOpenVNCResult(result.Body)
	if err != nil {
		return
	}

	vncURL = data.URL()
	return
}
