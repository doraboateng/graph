# Contributing

>TODO

```shell
# Build dev container
docker build --tag doraboateng/graph-service:dev --target dev .

# Open IDE
code .
```

# Updating the schema

1. Update the schema file in question.
2. Run `go generate`:

```shell
./run shell
go generate ./...
```
