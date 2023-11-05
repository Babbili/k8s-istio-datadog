#!/bin/bash

DOCKER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# cd $DOCKER_DIR && docker build -f Dockerfile --tag alpine ../

# docker scout cves --type image alpine

# docker run --name alpine-c alpine

cd $DOCKER_DIR && docker build -f Dockerfile --tag babbili/alpine:c6702847 ../

# docker scout cves --type image babbili/alpine:c6702847

docker push babbili/alpine:c6702847
