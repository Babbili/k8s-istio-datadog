#!/bin/bash

DOCKER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# cd $DOCKER_DIR && docker build -f Dockerfile --tag alpine:c6707b6eb ../

# docker scout cves --type image alpine:c6707b6eb

# docker run --name alpine-c alpine:c6707b6eb

cd $DOCKER_DIR && docker build -f Dockerfile --tag babbili/alpine:c6707b6eb ../

# docker scout cves --type image babbili/alpine:c6707b6eb

docker push babbili/alpine:c6707b6eb
