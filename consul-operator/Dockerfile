# Copyright 2021 Nokia
# Licensed under the BSD 3-Clause License.
# SPDX-License-Identifier: BSD-3-Clause

# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY ./ /workspace

# Build
RUN make all

FROM docker.io/library/alpine:latest

ENV RESREQ_DIR=/usr/src/app/resource-reqs-generated \
    DEPLOYMENT_DIR=/usr/src/app

WORKDIR /
COPY --from=builder /workspace/app-operator .

COPY deployment/helm /usr/local/bin/helm
RUN chmod +x  /usr/local/bin/helm
COPY deployment/app-deployment /usr/src/app/app-deployment
COPY deployment/config /usr/src/app/config
COPY deployment/resource-reqs /usr/src/app/resource-reqs
RUN chmod -R 770 /usr/src/app/

USER 65532:65532

ENTRYPOINT ["/app-operator"]
