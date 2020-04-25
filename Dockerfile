ARG GO_VERSION=1.14.2

# Base image for building and developmemnt.
FROM golang:${GO_VERSION}-buster AS base

RUN apt-get update \
    && apt-get upgrade --yes \
    && rm -rf /var/lib/apt/lists/*

ADD . /graph-service
WORKDIR /graph-service

# Development stage.
FROM base AS dev

RUN apt-get update \
    && apt-get upgrade --yes \
    && apt-get install --no-install-recommends --yes htop less vim \
    && apt-get remove subversion --yes \
    && apt-get autoremove --yes \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go get -v github.com/cosmtrek/air

# Build and install Go tools
# Reference: https://github.com/microsoft/vscode-dev-containers/blob/master/containers/go/.devcontainer/Dockerfile
RUN mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    && GOPATH=/tmp/gotools GO111MODULE=on go get -v golang.org/x/tools/gopls@latest 2>&1 \
    && GOPATH=/tmp/gotools GO111MODULE=on go get -v \
        honnef.co/go/tools/...@latest \
        golang.org/x/tools/cmd/gorename@latest \
        golang.org/x/tools/cmd/goimports@latest \
        golang.org/x/tools/cmd/guru@latest \
        golang.org/x/lint/golint@latest \
        github.com/mdempsky/gocode@latest \
        github.com/cweill/gotests/...@latest \
        github.com/haya14busa/goplay/cmd/goplay@latest \
        github.com/sqs/goreturns@latest \
        github.com/josharian/impl@latest \
        github.com/davidrjenni/reftools/cmd/fillstruct@latest \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest  \
        github.com/ramya-rao-a/go-outline@latest  \
        github.com/acroca/go-symbols@latest  \
        github.com/godoctor/godoctor@latest  \
        github.com/rogpeppe/godef@latest  \
        github.com/zmb3/gogetdoc@latest \
        github.com/fatih/gomodifytags@latest  \
        github.com/mgechev/revive@latest  \
        github.com/go-delve/delve/cmd/dlv@latest 2>&1 \
    && GOPATH=/tmp/gotools go get -v github.com/alecthomas/gometalinter 2>&1 \
    && GOPATH=/tmp/gotools go get -x -d github.com/stamblerre/gocode 2>&1 \
    && GOPATH=/tmp/gotools go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv /tmp/gotools/bin/* /usr/local/bin/ \
    && mv gocode-gomod /usr/local/bin/ \
    && cd \
    && rm -rf /tmp/gotools \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin 2>&1 \
    && go clean -cache -modcache

# Build stage.
FROM base as build

ARG BUILD_VERSION
ARG GIT_HASH
WORKDIR /graph-service/src
RUN CGO_ENABLED=0 GOOS=linux go build \
        -ldflags "-X main.version=${BUILD_VERSION} -X main.gitHash=${GIT_HASH}" \
        -o /tmp/graph-service
RUN chmod +x /tmp/graph-service

# Production stage.
# TODO: should we be using Alpine (alpine:3.9.6) or Distroless
# (gcr.io/distroless/static) instead?
FROM scratch AS prod

ARG BUILD_VERSION
ARG GIT_HASH

COPY --from=base /graph-service/src/schema/*.gql /opt/
COPY --from=base /graph-service/src/schema/*.dgraph /opt/
COPY --from=build /tmp/graph-service /usr/local/bin/graph-service

ENV API_ENV=production
ENV GRAPH_SCHEMA_PATH=/opt/graph.gql
ENV GRAPH_INDICES_PATH=/opt/indices.dgraph

ENTRYPOINT ["/usr/local/bin/graph-service"]
