#!/bin/sh

set -xe

LATEST_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))
git log -n1
git ls-remote --heads --tags https://github.com/claranet/terraform-provider-cloud-deploy.git | grep -E "refs/(heads|tags)/${LATEST_TAG}$"

PACKAGE_PATH="${GOPATH}/src/cloud-deploy.io/terraform-provider-cloud-deploy"
mkdir -pv "${PACKAGE_PATH}"
tar -cO . | tar -xv -C "${PACKAGE_PATH}"
cd "${PACKAGE_PATH}"

docker run --rm goreleaser/goreleaser:v0.80
