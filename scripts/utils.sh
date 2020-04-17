#!/bin/sh

get_env() {
    ENV_NAME=$1
    [ "$ENV_NAME" = "" ] && return 1

    ./scripts/create-env.sh

    set -a
    . .env
    set +a

    eval "ENV_VALUE=\$$ENV_NAME"
    echo "$ENV_VALUE"

    return 0
}

image_exists() {
    IMAGE_ID=$(docker images --quiet "$1")
    if [ "$IMAGE_ID" = "" ]; then
        return 1
    fi

    return 0
}
