package server

import (
	"github.com/mitchellh/mapstructure"
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

type ServerImager struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerImager(client ifaces.Openstacker) *ServerImager {
	return &ServerImager{
		Client: client,
	}
}

func (ser *ServerImager) Create(serverID, imageName string) (imageID string, err error) {
	if serverID == "" || imageName == "" {
		err = errors.ErrInvalidParams
		return
	}

	client, err := ser.Client.ComputeClient()
	if err != nil {
		return
	}

	serverClient := New(ser.Client)

	server, err := serverClient.Show(serverID)
	if err != nil {
		return
	}

	reqBody, err := servers.CreateImageOpts{
		Name: imageName,
		Metadata: map[string]string{
			"base_image_id": server.BaseImageID(),
		},
	}.ToServerCreateImageMap()
	if err != nil {

		return
	}

	var result gophercloud.Result

	_, err = client.Post(client.ServiceURL(ServersUrl, serverID, "action"), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	var response struct {
		ImageID string `mapstructure:"image_id"`
	}

	err = mapstructure.Decode(result.Body, &response)
	if err == nil {
		imageID = response.ImageID
	}

	return
}
