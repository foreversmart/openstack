all: gobuild gotest

gobuild: goclean goinstall

gorebuild: goclean goreinstall

goclean:
	go clean -i ./...

goinstall:
	go get -t -v github.com/qbox/openstack-golang-sdk
	go get github.com/dolab/gogo
	go get github.com/dolab/httpmitm
	go get github.com/dolab/httptesting
	go get github.com/mitchellh/mapstructure
	go get github.com/satori/go.uuid
	go get github.com/rackspace/gophercloud

goreinstall:
	go get -t -u -v github.com/qbox/openstack-golang-sdk
	go get -u -v github.com/dolab/gogo
	go get -u -v github.com/dolab/httpmitm
	go get -u -v github.com/dolab/httptesting
	go get -u -v github.com/mitchellh/mapstructure
	go get -u -v github.com/satori/go.uuid
	go get -u -v github.com/rackspace/gophercloud

define deep_test
    $(eval gopath = $(subst :, ,$(GOPATH)))
    $(eval firstGopath = $(word 1,$(gopath)))
    $(eval current = $(subst ...,,$(1)))
    $(eval files = $(wildcard $(firstGopath)/src/$(current)*/))
    $(eval testFiles = $(filter %_test.go, $(files)))
    $(if $(testFiles),go test $(current))
    $(eval packages = $(filter %/, $(files)))
    $(if $(packages),$(foreach package,$(packages),$(call deep_test,$(subst $(firstGopath)/src/,,$(package))...)))
endef

gotest: goinstall
	$(call deep_test,github.com/qbox/openstack-golang-sdk/...)

travis: gobuild gotest
