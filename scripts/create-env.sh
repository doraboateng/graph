#!/bin/sh

RESET_ENV=$1
if [ "$RESET_ENV" = "reset" ]; then
    rm ./.env
fi

if [ ! -f ./.env ]; then
    echo "Creating environment file..."
    touch ./.env

    echo "Docker hub username:"
    read -r DOCKER_HUB_USERNAME

    echo "Docker hub token:"
    read -r DOCKER_HUB_TOKEN

    {
        echo "# App"
        echo "APP_ENV=local"
        echo "APP_PORT=8810"

        echo ""
        echo "# Credentials"
        echo "DOCKER_HUB_USERNAME=$DOCKER_HUB_USERNAME"
        echo "DOCKER_HUB_TOKEN=$DOCKER_HUB_TOKEN"

        echo ""
        echo "# Dgraph"
        echo "DGRAPH_ZERO_HTTP_PORT=8819"
        echo "DGRAPH_ALPHA_PORT=8818"
        echo "DGRAPH_RATEL_PORT=8812"
        echo "GRAPHIQL_PORT=8811"
    } >> ./.env

    echo "Environment file created."
fi
