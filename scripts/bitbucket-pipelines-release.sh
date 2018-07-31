#!/bin/sh

PACKAGE_PATH="${GOPATH}/src/cloud-deploy.io/terraform-provider-cloud-deploy"
mkdir -pv "${PACKAGE_PATH}"
tar -cO --exclude-vcs --exclude=bitbucket-pipelines.yml . | tar -xv -C "${PACKAGE_PATH}"
cd "${PACKAGE_PATH}"
git ls-remote --heads --tags https://github.com/claranet/terraform-provider-cloud-deploy.git | grep -E "refs/(heads|tags)/${BITBUCKET_TAG}$"
docker run --rm goreleaser/goreleaser:v0.80
