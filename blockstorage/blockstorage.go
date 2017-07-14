package blockstorage

import (
	"github.com/qbox/openstack-golang-sdk/blockstorage/snapshot"
	"github.com/qbox/openstack-golang-sdk/blockstorage/volume"
	"github.com/qbox/openstack-golang-sdk/lib/ifaces"
)

type BlockStorage struct {
	client ifaces.Openstacker

	_ bool
}

func New(client ifaces.Openstacker) *BlockStorage {
	return &BlockStorage{
		client: client,
	}
}

func (bs *BlockStorage) NewVolumer() ifaces.Volumer {
	return volume.New(bs.client)
}

func (bs *BlockStorage) NewSnapshoter() ifaces.Snapshoter {
	return snapshot.New(bs.client)
}
