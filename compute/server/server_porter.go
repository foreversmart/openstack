package server

import "github.com/qbox/openstack-golang-sdk/lib/ifaces"

type ServerPorter struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerPorter(client ifaces.Openstacker) *ServerPorter {
	return &ServerPorter{
		Client: client,
	}
}

// func (sm *ServerManager) BindPort(id, portID string) error {
// 	if id == "" || portID == "" {
// 		return errors.ErrInvalidParams
// 	}

// 	client, err := ser.Client.ComputeClient()
// 	if err != nil {
// 		return err
// 	}

// 	opts := map[string]interface{}{
// 		"port_id": portID,
// 	}

// 	reqBody := map[string]interface{}{
// 		"interfaceAttachment": opts,
// 	}

// 	var res gophercloud.Result

// 	_, res.Err = client.Post(client.ServiceURL(ServersUrl, id, InterfaceUrl), reqBody, &res.Body, &gophercloud.RequestOpts{
// 		OkCodes: []int{200},
// 	})

// 	return res.Err
// }

// func (ser *Servers) Ports(id string) (portIDs []string, err error) {
// 	if id == "" {
// 		err = errors.ErrInvalidParams
// 		return
// 	}

// 	client, err := ser.Client.ComputeClient()
// 	if err != nil {
// 		return
// 	}

// 	var res gophercloud.Result

// 	_, res.Err = client.Get(client.ServiceURL(ServersUrl, id, InterfaceUrl), &res.Body, &gophercloud.RequestOpts{
// 		OkCodes: []int{200},
// 	})

// 	attachPorts, err := models.ExtractAttachPorts(res)

// 	portIDs = make([]string, len(attachPorts))
// 	for index, port := range attachPorts {
// 		portIDs[index] = port.PortId
// 	}

// 	return portIDs, err
// }

// // TODO
// func (ser *Servers) UnbindPort(id, portID string) error {
// 	if id == "" || portID == "" || ser.Client.ProjectID() == "" {
// 		return errors.ErrInvalidParams
// 	}

// 	return errors.ErrNotImplemented
// }
