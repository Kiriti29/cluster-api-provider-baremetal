#!/bin/sh

set -eux

IS_CONTAINER=${IS_CONTAINER:-false}
ARTIFACTS=${ARTIFACTS:-/tmp}

if [ "${IS_CONTAINER}" != "false" ]; then
  export XDG_CACHE_HOME=/tmp/.cache
  eval "$(go env)"
  cd "${GOPATH}"/src/github.com/metal3-io/cluster-api-provider-baremetal
  go test ./pkg/... ./cmd/... -coverprofile "${ARTIFACTS}"/cover.out
else
  podman run --rm \
    --env IS_CONTAINER=TRUE \
    --volume "${PWD}:/root/go/src/github.com/metal3-io/cluster-api-provider-baremetal:ro,z" \
    --entrypoint sh \
    --workdir /root/go/src/github.com/metal3-io/cluster-api-provider-baremetal \
    quay.io/metal3-io/capbm-unit:v1alpha1 \
    /root/go/src/github.com/metal3-io/cluster-api-provider-baremetal/hack/unit.sh "${@}"
fi;
