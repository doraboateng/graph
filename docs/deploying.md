# Deploying

>TODO

```shell
# Tag dev image.
docker build --tag doraboateng/graph-service-dev:20.04.0 --target dev .
docker image rm --force doraboateng/graph-service-dev:latest
docker tag doraboateng/graph-service-dev:20.04.0 doraboateng/graph-service-dev:latest

# Push dev image to Docker Hub.
set -a && source .env && set +a
echo $DOCKER_HUB_TOKEN | docker login --username $DOCKER_HUB_USERNAME --password-stdin
docker push doraboateng/graph-service-dev:20.04.0

# TODO: push to Docker Hub
# TODO: create .devcontainer.json file at root
# https://code.visualstudio.com/docs/remote/containers#_creating-a-devcontainerjson-file
```
