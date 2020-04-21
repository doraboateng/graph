#!/bin/sh

# Exit on error
set -e

. scripts/utils.sh

set -a
. .env
set +a

IMAGE_NAME="doraboateng/graph-service:dev"
CONTAINER_NAME="boateng-graph-service"

# Check that Docker is installed.
if [ ! -x "$(command -v docker)" ]; then
    echo "Could not find Docker in path." >&2
    exit 1
fi

# Check that Docker image is built.
if ! image_exists "$IMAGE_NAME"; then
    echo "Please build Docker image first by running \"./run build-docker dev\"."
    exit 1
fi

RUNNING_CONTAINER_ID=$(docker container ls --filter name="$CONTAINER_NAME" --quiet)

if [ "$RUNNING_CONTAINER_ID" != "" ]; then
    docker exec --interactive --tty "$RUNNING_CONTAINER_ID" bash
    exit 0
fi

docker run \
    --env APP_ENV="$APP_ENV" \
    --env DGRAPH_ZERO_PORT="$DGRAPH_ZERO_PORT" \
    --env DGRAPH_ALPHA_PORT="$DGRAPH_ALPHA_PORT" \
    --env DGRAPH_RATEL_PORT="$DGRAPH_RATEL_PORT" \
    --env GRAPHIQL_PORT="$GRAPHIQL_PORT" \
    --interactive \
    --mount type="bind,source=$(pwd),target=/graph-service" \
    --name "$CONTAINER_NAME" \
    --publish "$APP_PORT:80" \
    --publish "$DGRAPH_ZERO_PORT:6080" \
    --publish "$DGRAPH_ALPHA_PORT:8080" \
    --publish "$DGRAPH_RATEL_PORT:8000" \
    --publish "$GRAPHIQL_PORT:$GRAPHIQL_PORT" \
    --rm \
    --tty \
   "$IMAGE_NAME" \
   bash
