# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22.5
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS production
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

COPY . .
 
EXPOSE 8000

ENTRYPOINT [ "/bin/server" ]
