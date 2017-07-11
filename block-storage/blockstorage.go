package blockstorage

import (
	"github.com/kirk-enterprise/openstack-golang-sdk/block-storage/volume"
	"github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"
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
