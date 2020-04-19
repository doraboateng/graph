#!/bin/sh

# Exit on error
set -e

. scripts/utils.sh

# Check that Docker is installed.
if [ ! -x "$(command -v docker)" ]; then
    echo "Could not find Docker in path." >&2
    exit 1
fi

# Check that Docker image is built.
if ! image_exists "doraboateng/graph-service:dev"; then
    echo "Please build Docker image first by running \"./run build-docker dev\"."
    exit 1
fi

docker run \
    --interactive \
    --mount type="bind,source=$(pwd),target=/go/src/github.com/kwcay/boateng-graph-service" \
    --name boateng-graph-service \
    --publish "$(get_env APP_PORT):80" \
    --rm \
    --tty \
   doraboateng/graph-service:dev \
   ash
