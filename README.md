[![Build Status](https://travis-ci.com/doraboateng/graph-service.svg?branch=stable)](https://travis-ci.com/doraboateng/graph-service)
[![Maintainability](https://api.codeclimate.com/v1/badges/af6ea36778ba43f5fc1d/maintainability)](https://codeclimate.com/github/doraboateng/graph-service/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/af6ea36778ba43f5fc1d/test_coverage)](https://codeclimate.com/github/doraboateng/graph-service/test_coverage)

<details>
    <summary>Table of Contents</summary>

- [Local Setup](#local-setup)
    - [Requirements](#requirements)
    - [Published Ports](#published-ports)
    - [Running the Graph Service locally](#running-the-graph-service-locally)
    - [Log output from the containers](#log-output-from-the-containers)
- [Reporting Bugs](#reporting-bugs)
- [Reporting Security Issues](#reporting-security-issues)
- [Contributing](https://github.com/kwcay/boateng-graph-service/blob/stable/docs/contributing.md)
- [Deploying](https://github.com/kwcay/boateng-graph-service/blob/stable/docs/deploying.md)

</details>

# Local Setup

## Requirements

- [Docker](https://www.docker.com)
- [Visual Studio Code](https://code.visualstudio.com)
- POSIX-compliant terminal, such as:
    - [Visual Studio Code terminal](https://code.visualstudio.com/docs/editor/integrated-terminal)
    - [cmder](https://cmder.net)
    - [Cygwin](https://www.cygwin.com)
    - [Bash](https://www.gnu.org/software/bash)
    - [Zsh](https://www.zsh.org)

If you're on Linux or Mac, you already have a POSIX-compliant terminal.

## Running the Graph Service

```shell
docker-compose up

# Load schema
curl localhost:8080/admin/schema --data-binary "@src/schema/graph.gql"
curl localhost:8080/alter --data-binary "@src/schema/indices.dgraph"
```

## Published Ports

Port numbers published to your host machine.

| Port | Service |
| --- | --- |
| 8810 | API |
| 8811 | [GraphiQL (todo)](https://github.com/graphql/graphiql) |
| 8812 | Dgraph Ratel |
| 8817 | [Dgraph Alpha](https://dgraph.io/docs/deploy/#more-about-dgraph-alpha) HTTP port |
| 8818 | [Dgraph Alpha](https://dgraph.io/docs/deploy/#more-about-dgraph-alpha) gRPC port |
| 8819 | [Dgraph Zero](https://dgraph.io/docs/deploy/#more-about-dgraph-zero) |

## Running the Graph Service locally

```shell
docker-compose up --detach
```

When running for the first time, you might get this message:

```
ERROR: The image for the service you're trying to recreate has been removed. If you continue, volume data could be lost. Consider backing up your data before continuing.

Continue with the new image? [yN]
```

In which case you can go ahead and type `y` and continue.

To stop the Graph Service:

```shell
docker-compose down
```

## Log output from the containers

```shell
docker-compose logs
docker-compose logs api
docker-compose logs api alpha
docker-compose logs --tail 5 api
docker-compose logs --follow api
docker-compose logs --follow api alpha zero
```

# Reporting Bugs

>TODO

# Reporting Security Issues

>TODO

# License

Copyright © 2020 Kwahu & Cayes
