ARG GO_VERSION=1.14.2

# Base image for building and developmemnt.
FROM golang:${GO_VERSION}-alpine AS base

RUN apk update && apk add --no-cache curl gcc git htop make vim

ENV CGO_ENABLED=0

ADD . /graph-service
WORKDIR /graph-service

# Development stage.
FROM base AS dev

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
ARG BUILD_NAME
WORKDIR /graph-service/src
RUN GOOS=linux go build \
        -ldflags "-X main.version=${BUILD_VERSION} -X main.gitHash=${GIT_HASH}" \
        -o /tmp/${BUILD_NAME}
RUN chmod +x /tmp/${BUILD_NAME}

# Production stage.
FROM scratch AS prod

ARG BUILD_NAME
COPY --from=build /tmp/${BUILD_NAME} /usr/local/bin/graph-service

ENTRYPOINT ["/usr/local/bin/graph-service"]
