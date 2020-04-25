#!/bin/sh

# Exit on error
set -e

. scripts/utils.sh

VERSION=$1
DOCKER_REPO="doraboateng/graph-service"
USAGE="Usage: ./run [VERSION]"

if [ "$VERSION" = "" ]; then
    VERSION=$(date "+%y.%m.0")
fi

if [ "$VERSION" = "help" ] || [ "$VERSION" = "-h" ] || [ "$VERSION" == "--help" ];
then
    echo "$USAGE"
    exit 0
fi

TAGGED_IMAGE="$DOCKER_REPO:$VERSION"
LATEST_IMAGE="$DOCKER_REPO:latest"

echo ""
echo "Building \"$TAGGED_IMAGE\". Continue? (yes/[no])"
read -r CONFIRMATION

if [ "$CONFIRMATION" != "yes" ]; then
    exit 0
fi

if image_exists "$TAGGED_IMAGE"; then
    docker image rm --force "$TAGGED_IMAGE"
fi

docker build \
    --build-arg BUILD_VERSION="$VERSION" \
    --force-rm \
    --tag "$TAGGED_IMAGE" \
    --target prod \
    .

echo ""
echo "Update \"latest\" tag for \"$DOCKER_REPO\" (\"$LATEST_IMAGE\")? (yes/[no])"
read -r CONFIRMATION

if [ "$CONFIRMATION" = "yes" ]; then
    if image_exists "$LATEST_IMAGE"; then
        docker image rm --force "$LATEST_IMAGE"
    fi

    docker tag "$TAGGED_IMAGE" "$LATEST_IMAGE"
fi

echo ""
echo "Publish build to Docker registry? (yes/[no])"
read -r CONFIRMATION

if [ "$CONFIRMATION" = "yes" ]; then
    get_env DOCKER_HUB_TOKEN | docker login \
        --username "$(get_env DOCKER_HUB_USERNAME)" \
        --password-stdin

    docker push "$TAGGED_IMAGE"
fi

echo ""
echo "Pruning Docker resources..."
echo ""

docker system prune
