#!/bin/bash -x
export DHUBREPO="hack4easy/kubesim_base-arm32v7"
export VERSION="0.1.0"
export DOCKER_NAMESPACE="hack4easy"
export DOCKER_USERNAME="kubedgedevops"
export DOCKER_PASSWORD=$KUBEDGEDEVOPSPWD
cd grpc/go/kubedge_server/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goserv .
cd ../../../
cd grpc/go/kubedge_client/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goclient .
cd ../../../
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build -t $DHUBREPO:$VERSION -f images/kubesim_base/Dockerfile .
docker images
docker tag $DHUBREPO:$VERSION $DHUBREPO:latest
docker tag $DHUBREPO:$VERSION $DHUBREPO:from-master-pi
docker push $DHUBREPO
