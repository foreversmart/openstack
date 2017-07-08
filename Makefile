all: gobuild gotest

gobuild: goclean goinstall

gorebuild: goclean goreinstall

goclean:
	go clean -i ./...

goinstall:
	go get -t -v github.com/kirk-enterprise/openstack
	go get github.com/dolab/gogo
	go get github.com/dolab/httpmitm
	go get github.com/dolab/httptesting
	go get github.com/mitchellh/mapstructure
	go get github.com/satori/go.uuid
	go get github.com/rackspace/gophercloud

goreinstall:
	go get -t -u -v github.com/kirk-enterprise/openstack
	go get -u -v github.com/dolab/gogo
	go get -u -v github.com/dolab/httpmitm
	go get -u -v github.com/dolab/httptesting
	go get -u -v github.com/mitchellh/mapstructure
	go get -u -v github.com/satori/go.uuid
	go get -u -v github.com/rackspace/gophercloud

gotest: goinstall
	go test github.com/kirk-enterprise/openstack
	go test github.com/kirk-enterprise/openstack/lib/...
	go test github.com/kirk-enterprise/openstack/internal/...
	go test github.com/kirk-enterprise/openstack/keystone/...

travis: gobuild gotest
