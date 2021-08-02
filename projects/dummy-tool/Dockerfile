# Build the manager binary
FROM golang:1.16 as builder

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /workspace

# install curl
RUN apt-get update \
    && apt-get install -y curl

COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY *.go /workspace/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o dummy-tool .

ENTRYPOINT ["dummy-tool"]
