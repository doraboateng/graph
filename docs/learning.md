# Rebuilding Dgraph schema

```shell
# Start Dgraph Zero
./run shell
dgraph zero

# Start Dgraph Alpha
docker exec --interactive --tty boateng-graph-service ash
dgraph alpha --lru_mb 1024

# Load schema
docker exec --interactive --tty boateng-graph-service ash
curl localhost:8080/admin/schema --data-binary "@src/schema/graph.gql"
curl localhost:8080/alter --data-binary "@src/schema/indices.dgraph"
```
