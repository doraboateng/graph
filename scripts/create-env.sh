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
        echo "# API"
        echo "API_ENV=local"
        echo "API_PORT=8810"

        echo ""
        echo "# Credentials"
        echo "DOCKER_HUB_USERNAME=$DOCKER_HUB_USERNAME"
        echo "DOCKER_HUB_TOKEN=$DOCKER_HUB_TOKEN"
    } >> ./.env

    echo "Environment file created."
fi
