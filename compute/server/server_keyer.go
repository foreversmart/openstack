package server

import (
	"github.com/qbox/openstack-golang-sdk/lib/errors"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
	"github.com/rackspace/gophercloud"
)

type ServerKeyer struct {
	Client ifaces.Openstacker

	_ bool
}

func NewServerKeyer(client ifaces.Openstacker) *ServerKeyer {
	return &ServerKeyer{
		Client: client,
	}
}

func (sk *ServerKeyer) Bind(serverID string, keys []string) error {
	if serverID == "" || len(keys) == 0 {
		return errors.ErrInvalidParams
	}

	client, err := sk.Client.ComputeClient()
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
	_, err = client.Post(client.ServiceURL(ServersUrl, serverID, ActionUrl), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil && err.Error() == "EOF" {
		err = nil
	}

	return err
}

func (sk *ServerKeyer) Unbind(serverID string, keys []string) error {
	if serverID == "" || len(keys) == 0 {
		return errors.ErrInvalidParams
	}

	client, err := sk.Client.ComputeClient()
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
	_, err = client.Post(client.ServiceURL(ServersUrl, serverID, ActionUrl), reqBody, &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil && err.Error() == "EOF" {
		err = nil
	}

	return err
}
