# Contributing

>TODO

```shell
# Build dev container
docker build --tag doraboateng/graph-service:dev --target dev .

# Open IDE
code .
```

# Updating the schema

>TODO

```shell
CGO_ENABLED=0 go run github.com/99designs/gqlgen generate
```
