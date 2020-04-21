[![Build Status](https://travis-ci.com/doraboateng/graph-service.svg?branch=stable)](https://travis-ci.com/doraboateng/graph-service)
[![Maintainability](https://api.codeclimate.com/v1/badges/af6ea36778ba43f5fc1d/maintainability)](https://codeclimate.com/github/doraboateng/graph-service/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/af6ea36778ba43f5fc1d/test_coverage)](https://codeclimate.com/github/doraboateng/graph-service/test_coverage)

<details>
    <summary>Table of Contents</summary>

- [Local Setup](#local-setup)
    - [Requirements](#requirements)
    - [Initial Setup](#initial-setup)
    - [Published Ports](#published-ports)
- [Reporting Bugs](#reporting-bugs)
- [Reporting Security Issues](#reporting-security-issues)
- [Contributing](https://github.com/kwcay/boateng-graph-service/blob/stable/docs/contributing.md)
- [Deploying](https://github.com/kwcay/boateng-graph-service/blob/stable/docs/deploying.md)

</details>

# Local Setup

## Requirements

- Docker
- Visual Studio Code

## Initial Setup

```shell
./run build-docker dev
```

## Running the Graph Service

```shell
./run shell

# Run Dgraph services
dgraph zero --cwd /dgraph &
dgraph alpha --cwd /dgraph --lru_mb 1024 &
dgraph-ratel &

# TODO: redirect terminal output to log?

# Load schema
curl localhost:8080/admin/schema --data-binary "@src/schema/graph.gql"
curl localhost:8080/alter --data-binary "@src/schema/indices.dgraph"

# Stop Dgraph services
pkill -f dgraph
```

## Published Ports

Port numbers published to your host machine.

| Port | Service |
| --- | --- |
| 8810 | Graph Service |
| 8811 | [GraphiQL (todo)](https://github.com/graphql/graphiql) |
| 8812 | Dgraph Ratel |
| 8818 | [Dgraph Alpha](https://dgraph.io/docs/deploy/#more-about-dgraph-alpha) |
| 8819 | [Dgraph Zero](https://dgraph.io/docs/deploy/#more-about-dgraph-zero) |

# Reporting Bugs

>TODO

# Reporting Security Issues

>TODO

# License

Copyright © 2020 Kwahu & Cayes
