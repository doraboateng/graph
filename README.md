# Local development

```shell
# Build for local development.
docker build --tag doraboateng/graph:dev --target dev .
```

<details>
    <summary>Rebuilding the schema</summary>

```shell
# TODO ...

./run shell

curl alpha:8080/alter -d '{ "drop_all": true }'
curl alpha:8080/admin/schema --data-binary "@schema/graph.gql"
curl alpha:8080/alter --data-binary "@schema/indices.dgraph"
```
</details>

# Releasing a new version

- Make sure the `stable` branch contains all the latest working changes.
- Create release on [Github](https://github.com/kwcay/boateng-graph/releases/new?target=stable) using [calendar versioning](https://calver.org) (e.g. v20.01.0).
    - See [latest releases](https://github.com/kwcay/boateng-graph/releases) for guidance.

A new release will be published to [Docker Hub](https://hub.docker.com/r/doraboateng/graph/tags?page=1).

# License

[GNU General Public License v3](https://github.com/kwcay/boateng-graph-service/blob/stable/LICENSE)

Copyright © Kwahu & Cayes
