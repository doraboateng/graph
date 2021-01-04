_Dora Boateng Graph_

# Releasing a new version

- Make sure the `stable` branch contains all the latest working changes.
- Create release on [Github](https://github.com/doraboateng/graph/releases/new?target=stable) using [calendar versioning](https://calver.org) (e.g. v20.01.0).
    - See [latest releases](https://github.com/doraboateng/graph/releases) for guidance.

# Other notes

<details>
    <summary>Loading data into Slash GraphQL</summary>

```shell
# OPTIONAL: Drop data.
curl "$(cat ./.credentials/slash-graphql-http-endpoint)/admin/slash" \
    --header "Content-Type: application/graphql" \
    --header "X-Auth-Token: $(cat ./.credentials/slash-graphql-key)" \
    --data-binary "mutation { dropData(allData: true) { response { code message } } }"

# OPTIONAL: Drop data and schema.
curl "$(cat ./.credentials/slash-graphql-http-endpoint)/admin/slash" \
    --header "Content-Type: application/graphql" \
    --header "X-Auth-Token: $(cat ./.credentials/slash-graphql-key)" \
    --data-binary "mutation { dropData(allDataAndSchema: true) { response { code message } } }"

# Load GraphQL schema.
# NOTE: this command requires Node v14 or later.
npx slash-graphql update-schema \
    --endpoint $(cat ./.credentials/slash-graphql-http-endpoint)/graphql \
    --token $(cat ./.credentials/slash-graphql-key) \
    src/schema/graph.gql

# Load Dgraph schema.
curl "$(cat ./.credentials/slash-graphql-http-endpoint)/alter" \
    --header "X-Auth-Token: $(cat ./.credentials/slash-graphql-key)" \
    --data-binary '@src/schema/indices.dgraph'

# Copy backup files into "tmp" folder.
mkdir -p tmp/data && cp <PATH TO RDF FILE> tmp/data/rdf.tar.gz
mkdir -p tmp/data/$(date +'%Y-%m-%d') \
    && mv tmp/data/rdf.tar.gz tmp/data/$(date +'%Y-%m-%d')/ \
    && cd tmp/data/$(date +'%Y-%m-%d') \
    && tar --extract --gzip --file rdf.tar.gz \
    && cp -R export/**/* . \
    && rm -rf export dgraph* \
    && cd ../../..

# Import data using the Live Loader.
docker run \
    --env-file .env \
    --interactive \
    --rm \
    --tty \
    --volume "$(pwd)/tmp/data/$(date +'%Y-%m-%d')/g01.rdf.gz/:/tmp/data.rdf.gz" \
    dgraph/dgraph:v20.11-slash \
    bash -c 'dgraph live \
        --auth_token "$SLASH_GRAPHQL_AUTH_TOKEN" \
        --files /tmp/data.rdf.gz \
        --slash_grpc_endpoint="${SLASH_GRAPHQL_GRPC_ENDPOINT}:443"'
```
</details>

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

[GNU General Public License v3](LICENSE)

Copyright © Kwahu & Cayes
