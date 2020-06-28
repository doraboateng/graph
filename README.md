**Dora Boateng Graph Service**

# Local development

```shell
# Build for local development.
docker build --tag doraboateng/graph:dev --target dev .
```

# Rebuilding Dgraph schema

```shell
./run shell

curl alpha:8080/alter -d '{ "drop_all": true }'
curl alpha:8080/admin/schema --data-binary "@src/schema/graph.gql"
curl alpha:8080/alter --data-binary "@src/schema/indices.dgraph"
```

<details>
    <summary><h1>Releasing a new version</h1></summary>

- Make sure the `stable` branch contains all the latest working changes.
- Create release on [Github](https://github.com/kwcay/boateng-graph/releases/new?target=stable) (using [calendar versioning](https://calver.org)).
    - See [latest releases](https://github.com/kwcay/boateng-graph/releases) for guidance.
- Publish the release to Docker Hub:

```shell
# Retrieve latest tags from Github.
git checkout stable && git fetch --tags && git pull

# Build and publish the release.
./run build-docker
```
</details>

# License

[GNU General Public License v3](https://github.com/kwcay/boateng-graph-service/blob/stable/LICENSE)

Copyright © Kwahu & Cayes
