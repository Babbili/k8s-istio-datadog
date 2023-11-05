#!/bin/bash

DOCKER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# cd $DOCKER_DIR && docker build -f Dockerfile --tag go-books-api:b878944c8 ../

# docker scout cves --type image go-books-api:b878944c8

# docker run --name go-books-api-c go-books-api:b878944c8

cd $DOCKER_DIR && docker build -f Dockerfile --tag babbili/go-books-api:b878944c8 ../

# docker scout cves --type image babbili/go-books-api:b878944c8

docker push babbili/go-books-api:b878944c8
