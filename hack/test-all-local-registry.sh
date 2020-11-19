#!/bin/bash

set -e -x -u

PORT=${1:-5000}

docker run -d -p "$PORT":5000 --restart always --name registry-"$PORT" registry:2
export IMGPKG_E2E_IMAGE="localhost:$PORT/local-tests/test-repo"
export IMGPKG_E2E_RELOCATION_REPO="localhost:$PORT/local-tests/test-relocation-repo"
./hack/test-all.sh $@

docker stop registry-"$PORT"
docker rm registry-"$PORT"
