package servers

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/qbox/openstack-golang-sdk/lib/models"
	"github.com/qbox/openstack-golang-sdk/lib/options"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
)

const (
	ServersUrl   = "servers"
	InterfaceUrl = "os-interface"
	ActionUrl    = "action"
)

type Servers struct {
	Client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *Servers {
	return &Servers{
		Client: client,
	}
}

func (ser *Servers) Create(opts options.CreateServersOpts) (server *models.ServersModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Post(client.ServiceURL(ServersUrl), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return models.ExtractServer(result)
}

func (ser *Servers) All() (servers []*models.ServersModel, err error) {
	return ser.AllByParams(nil)
}

func (ser *Servers) AllByParams(opts *options.ListServersOpts) (Serverss []*models.ServersModel, err error) {
	if !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}
	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServersUrl, "detail")+"?"+opts.ToQuery().Encode(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServers(result)
}

func (ser *Servers) Show(id string) (server *models.ServersModel, err error) {
	if id == "" {
		return nil, errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, result.Err = client.Get(client.ServiceURL(ServersUrl, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServer(result)
}

func (ser *Servers) Update(id string, opts options.UpdateServersOpts) (server *models.ServersModel, err error) {
	if id == "" || !opts.IsValid() {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	var result gophercloud.Result

	_, err = client.Put(client.ServiceURL(ServersUrl, id), opts.ToPayload(), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return models.ExtractServer(result)
}

func (ser *Servers) SearchByFixedIP(ip string) (servers []*models.ServersModel, err error) {
	if ip == "" {
		err = errors.ErrInvalidParams
		return
	}

	return ser.AllByParams(&options.ListServersOpts{
		Ip:         options.String(ip),
		AllTenants: options.Int(1),
	})
}

func (ser *Servers) Delete(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	_, err = client.Delete(client.ServiceURL(ServersUrl, id), &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

// =============== servers actions ===============

func (ser *Servers) ChangeAdminPassword(id, newPassword string) error {
	if id == "" || newPassword == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	return servers.ChangeAdminPassword(client, id, newPassword).ExtractErr()
}

func (ser *Servers) AdminPassword(id string) (password string, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputerClient()
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

func (ser *Servers) Start(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	return startstop.Start(client, id).ExtractErr()
}

func (ser *Servers) Reboot(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	return servers.Reboot(client, id, servers.SoftReboot).ExtractErr()
}

// TODO add own shutdown impl (openstack sdk has no shutdown method)
func (ser *Servers) Shutdown(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	return ser.Stop(id)
}

func (ser *Servers) Stop(id string) error {
	if id == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	return startstop.Stop(client, id).ExtractErr()
}

func (ser *Servers) CreateSnapshot(id, name string) (snapshotID string, err error) {
	if id == "" || name == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return
	}

	vm, err := ser.Show(id)
	if err != nil {
		return
	}

	baseImageID := vm.Image
	if vm.Metadata != nil && vm.Metadata["base_image_id"] != "" {
		baseImageID = (vm.Metadata["base_image_id"]).(string)
	}

	opts := servers.CreateImageOpts{
		Name: name,
		Metadata: map[string]string{
			"base_image_id": baseImageID,
		},
	}

	result := servers.CreateImage(client, id, opts)

	return result.ExtractImageID()
}

func (ser *Servers) BindPort(id, portID string) error {
	if id == "" || portID == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	opts := map[string]interface{}{
		"port_id": portID,
	}

	reqBody := map[string]interface{}{
		"interfaceAttachment": opts,
	}

	var res gophercloud.Result

	_, res.Err = client.Post(client.ServiceURL(ServersUrl, id, InterfaceUrl), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return res.Err
}

func (ser *Servers) Ports(id string) (portIDs []string, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return
	}

	var res gophercloud.Result

	_, res.Err = client.Get(client.ServiceURL(ServersUrl, id, InterfaceUrl), &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	attachPorts, err := models.ExtractAttachPorts(res)

	portIDs = make([]string, len(attachPorts))
	for index, port := range attachPorts {
		portIDs[index] = port.PortId
	}

	return portIDs, err
}

// TODO
func (ser *Servers) UnbindPort(id, portID string) error {
	if id == "" || portID == "" || ser.Client.ProjectID() == "" {
		return errors.ErrInvalidParams
	}

	return errors.ErrNotImplemented
}

func (ser *Servers) BindKeys(id string, keys []string) error {
	if id == "" || len(keys) == 0 {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	opts := map[string]interface{}{
		"key_names": keys,
	}

	reqBody := map[string]interface{}{
		"attachKeypairs": opts,
	}

	var result gophercloud.Result
	_, err = client.Post(client.ServiceURL(ServersUrl, id, ActionUrl), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	if err.Error() == "EOF" {
		err = nil
	}

	return err
}

func (ser *Servers) UnbindKeys(id string, keys []string) error {
	if id == "" || len(keys) == 0 {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	opts := map[string]interface{}{
		"key_names": keys,
	}

	reqBody := map[string]interface{}{
		"detachKeypairs": opts,
	}

	var result gophercloud.Result
	_, err = client.Post(client.ServiceURL(ServersUrl, id, ActionUrl), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	if err.Error() == "EOF" {
		err = nil
	}

	return err
}

func (ser *Servers) MountVolume(id, volumeID string) error {
	if id == "" || volumeID == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	opts := volumeattach.CreateOpts{
		VolumeID: volumeID,
	}

	return volumeattach.Create(client, id, opts).Err
}

func (ser *Servers) UnmountVolume(id, volumeID string) error {
	if id == "" || volumeID == "" {
		return errors.ErrInvalidParams
	}

	var volumeAttachID string

	client, err := ser.Client.ComputerClient()
	if err != nil {
		return err
	}

	pager := volumeattach.List(client, id)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		vaList, err := volumeattach.ExtractVolumeAttachments(page)

		for _, value := range vaList {
			if value.VolumeID == volumeID {
				volumeAttachID = value.ID
				return false, err
			}
		}

		return true, err

	})

	if err != nil {
		return err
	}

	if volumeAttachID == "" {
		return errors.ErrNotFound
	}

	return volumeattach.Delete(client, id, volumeAttachID).Err
}

func (ser *Servers) ModifyFlavor(id, flavorID string) error {
	if id == "" || flavorID == "" {
		return errors.ErrInvalidParams
	}

	client, err := ser.Client.ComputerClient()
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
func (ser *Servers) Rebuild(id, imageID string) error {
	return errors.ErrNotImplemented
}

func (ser *Servers) Vnc(id string) (vncURL string, err error) {
	if id == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := openstack.NewComputeV2(ser.Client.ProviderClient(), gophercloud.EndpointOpts{})
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
