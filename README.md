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

# Other notes

<details>
    <summary>Loading data into production</summary>

```shell
# TODO ...

# Copy schema files and RDF backup to graph server.
scp ./src/schema/{graph.gql,indices.dgraph} boateng@graph.doraboateng.com:/tmp/
scp /path/to/rdf.tar.gz boateng@graph.doraboateng.com:/tmp/rdf.tar.gz

ssh boateng@graph.doraboateng.com

# Move and extract RDF backup.
mkdir -p /tmp/restore/$(date +'%Y-%m-%d') \
    && mv /tmp/{graph.gql,indices.dgraph,rdf.tar.gz} /tmp/restore/$(date +'%Y-%m-%d')/ \
    && cd /tmp/restore/$(date +'%Y-%m-%d') \
    && tar --extract --gzip --file rdf.tar.gz \
    && cp export/**/* .

# OPTIONAL: Drop all records.
curl localhost:8080/alter -d '{ "drop_all": true }'

# OPTIONAL: reset database.
cd /src/graph
docker-compose down
docker-compose up --detach --force-recreate
cd /tmp/restore/$(date +'%Y-%m-%d')

# Load schema files.
curl localhost:8080/admin/schema --data-binary "@graph.gql" \
    && curl localhost:8080/alter --data-binary "@indices.dgraph"

# Load backup.
docker run \
    --interactive \
    --rm \
    --network graph_network \
    --tty \
    --volume $(pwd):/tmp \
    doraboateng/graph \
    dgraph live \
        --alpha alpha:9080 \
        --files /tmp/g01.rdf.gz \
        --zero zero:5080

exit
```
</details>

# License

[GNU General Public License v3](https://github.com/kwcay/boateng-graph-service/blob/stable/LICENSE)

Copyright © Kwahu & Cayes
