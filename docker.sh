#!/bin/bash

set -e

REPO=jnewmano/mqtt-nsq

BRANCH=$(git rev-parse --abbrev-ref HEAD)
TAG=v0.0.$(git rev-list HEAD --count)

if [ $BRANCH != "master" ]; then
  TAG="${BRANCH}-${TAG}"
fi

git tag ${TAG}
git push origin ${TAG}

GOOS=linux

CGO_ENABLED=0 \
GOOS=${GOOS} \
GOARCH=amd64 \
go build -v -i -ldflags "-s -w" -o bin/mqtt-to-nsq ./apps/mqtt-to-nsq

CGO_ENABLED=0 \
GOOS=${GOOS} \
GOARCH=amd64 \
go build -v -i -ldflags "-s -w" -o bin/nsq-to-mqtt ./apps/nsq-to-mqtt

docker build -t ${REPO}:${TAG} .

docker push ${REPO}:${TAG}
