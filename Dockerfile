# Base image for building and development.
FROM dgraph/dgraph:v20.03.3 AS base

WORKDIR /dgraph

# Build stage.
FROM base as build

ARG BUILD_VERSION
ARG GIT_HASH

WORKDIR /graph-src
ADD ./src/* /graph-src/

# Production stage.
FROM gcr.io/distroless/base AS prod

COPY --from=base /usr/local/bin/* /usr/local/bin/
# COPY --from=build /graph-src/bin/* /usr/local/bin/
COPY --from=build /graph-src/schema /schema

CMD ["dgraph"]

# Development stage.
FROM base AS dev

# TODO: install dev tools.
