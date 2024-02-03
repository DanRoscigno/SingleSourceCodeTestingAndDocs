FROM golang:1.21.6-alpine3.19 as builder

USER 0
WORKDIR /workspace


COPY ci ci

# install curl
RUN apk --no-cache add curl

# Install Ginkgo CLI to run unit tests inside the container
RUN cd ci && go install github.com/onsi/ginkgo/v2/ginkgo@v2.15.0

RUN cd ci && go mod download

