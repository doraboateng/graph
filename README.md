[![Build Status](https://travis-ci.com/doraboateng/graph-service.svg?branch=stable)](https://travis-ci.com/doraboateng/graph-service)
[![Maintainability](https://api.codeclimate.com/v1/badges/af6ea36778ba43f5fc1d/maintainability)](https://codeclimate.com/github/doraboateng/graph-service/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/af6ea36778ba43f5fc1d/test_coverage)](https://codeclimate.com/github/doraboateng/graph-service/test_coverage)

<details>
    <summary>Table of Contents</summary>

- [Local Setup](#local-setup)
    - [Requirements](#requirements)
    - [Running the Graph Service locally](#running-the-graph-service-locally)
    - [Published Ports](#published-ports)
    - [Viewing the log outputs from the services](#viewing-the-log-outputs-from-the-services)
- [Reporting Bugs](#reporting-bugs)
- [Reporting Security Issues](#reporting-security-issues)
- [Contributing](https://github.com/kwcay/boateng-graph-service/blob/stable/docs/contributing.md)
- [Deploying](https://github.com/kwcay/boateng-graph-service/blob/stable/docs/deploying.md)
- [License](#license)

</details>

# Local setup

## Requirements

- [Docker](https://www.docker.com) & [Docker Compose](https://docs.docker.com/compose/install)
- [Visual Studio Code](https://code.visualstudio.com)
- A POSIX-compliant terminal, such as:
    - [Visual Studio Code terminal](https://code.visualstudio.com/docs/editor/integrated-terminal)
    - [cmder](https://cmder.net)
    - [Cygwin](https://www.cygwin.com)
    - [Bash](https://www.gnu.org/software/bash)
    - [Zsh](https://www.zsh.org)

If you're on Linux or Mac, you already have a POSIX-compliant terminal.

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

## Published ports

Port numbers published to your host machine.

| Port | Service |
| --- | --- |
| 8810 | Graph API |
| 8811 | [GraphiQL (todo)](https://github.com/graphql/graphiql) |
| 8812 | Dgraph Ratel |
| 8817 | [Dgraph Alpha](https://dgraph.io/docs/deploy/#more-about-dgraph-alpha) (HTTP) |
| 8818 | [Dgraph Alpha](https://dgraph.io/docs/deploy/#more-about-dgraph-alpha) (gRPC) |
| 8819 | [Dgraph Zero](https://dgraph.io/docs/deploy/#more-about-dgraph-zero) |

## Viewing the log outputs from the services

```shell
# Displaying all logs.
docker-compose logs

# Displaying logs for the API service.
docker-compose logs api

# Displaying logs for several services, e.g. API and Dgraph Alpha
docker-compose logs api alpha

# Tailing the last 5 lines of the logs from the API service.
docker-compose logs --tail 5 api

# Following the logs for the API service as they come in (CMD/CTRL+C to exit).
docker-compose logs --follow api

# Following the logs for several services as they come in, e.g. Dgraph Alpha and Dgraph Zero.
docker-compose logs --follow api alpha zero
```

For more details, see the [docs](https://docs.docker.com/compose/reference/logs) or run the command `docker-compose logs --help`

# Reporting Bugs

>TODO

# Reporting Security Issues

>TODO

# License

[GNU General Public License v3](https://github.com/kwcay/boateng-graph-service/blob/stable/LICENSE) © Kwahu & Cayes
