#!/bin/sh

# Exit on error
set -e

. scripts/utils.sh

set -a
. .env
set +a

IMAGE_NAME="doraboateng/graph-service:dev"
CONTAINER_NAME="boateng-graph-service-api"

# Check that Docker is installed.
if [ ! -x "$(command -v docker)" ]; then
    echo "Could not find Docker in path." >&2
    exit 1
fi

# Check that Docker image is built.
if ! image_exists "$IMAGE_NAME"; then
    echo "Please build Docker image first by running \"docker-compose build api\"."
    exit 1
fi

RUNNING_CONTAINER_ID=$(docker container ls --filter name="$CONTAINER_NAME" --quiet)

if [ "$RUNNING_CONTAINER_ID" = "" ]; then
    echo "Please make sure the Graph Service is running (\"docker-compose up --detach\")."
    exit 1
fi

docker exec --interactive --tty "$RUNNING_CONTAINER_ID" ash
