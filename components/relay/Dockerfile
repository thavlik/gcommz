ARG KINDEST_REPOSITORY=""
ARG KINDEST_TAG="latest"
ARG RUNTIME_BASE_IMAGE="alpine:3.14.2"

FROM ${KINDEST_REPOSITORY}thavlik/gcommz-go-base:${KINDEST_TAG} AS builder
COPY components/relay components/relay
WORKDIR /go/src/github.com/thavlik/gcommz/components/relay/cmd/relay
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM ${RUNTIME_BASE_IMAGE}
COPY --from=builder /go/src/github.com/thavlik/gcommz/components/relay/cmd/relay/relay /usr/local/bin
CMD ["relay", "server"]

