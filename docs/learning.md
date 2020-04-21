# Rebuilding Dgraph schema

```shell
./run shell

curl alpha:8080/admin/schema --data-binary "@src/schema/graph.gql"
curl alpha:8080/alter --data-binary "@src/schema/indices.dgraph"
```
