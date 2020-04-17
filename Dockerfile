ARG DGRAPH_VERSION=20.03.0
ARG GO_VERSION=1.14.2

# Dev stage.
FROM golang:${GO_VERSION}-alpine AS dev

ARG DGRAPH_VERSION
RUN apk add --no-cache gcc git htop make vim && \
    git clone https://github.com/dgraph-io/dgraph.git /go/src/dgraph && \
        cd /go/src/dgraph && \
        git checkout v${DGRAPH_VERSION} --quiet && \
        CGO_ENABLED=0 make install --jobs=4 && \
        cd /go && \
        rm -rf /go/src/dgraph && \
    go get -u -v github.com/cosmtrek/air && \
    go clean -cache -modcache

ADD . /graph-service
WORKDIR /graph-service/src

# Build stage.
FROM dev as build

ARG BUILD_VERSION
ARG GIT_HASH
ARG BUILD_NAME
RUN CGO_ENABLED=0 GOOS=linux go build \
        -ldflags "-X main.version=${BUILD_VERSION} -X main.gitHash=${GIT_HASH}" \
        -o /tmp/${BUILD_NAME}
RUN chmod +x /tmp/${BUILD_NAME}

# Production stage.
FROM scratch AS dist

ARG BUILD_NAME
COPY --from=build /go/bin/dgraph /usr/local/bin/dgraph
COPY --from=build /tmp/${BUILD_NAME} /usr/local/bin/graph-service

EXPOSE ${APP_PORT}
ENTRYPOINT ["/usr/local/bin/graph-service"]
