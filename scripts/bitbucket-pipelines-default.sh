#!/bin/sh

PACKAGE_PATH="${GOPATH}/src/cloud-deploy.io/terraform-provider-cloud-deploy"
mkdir -pv "${PACKAGE_PATH}"
tar -cO --exclude-vcs --exclude=bitbucket-pipelines.yml . | tar -xv -C "${PACKAGE_PATH}"
cd "${PACKAGE_PATH}"
go get github.com/kardianos/govendor
make test
make vendor-status
make vet
