#!/bin/bash

DOCKER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# cd $DOCKER_DIR && docker build -f Dockerfile --tag text-to-speech:py075b11e ../

# docker scout cves --type image text-to-speech:py075b11e

# docker run --name text-to-speech-py075b11e-speech:py075b11e

cd $DOCKER_DIR && docker build -f Dockerfile --tag babbili/text-to-speech:py075b11e ../

# docker scout cves --type image babbili/text-to-speech:py075b11e

docker push babbili/text-to-speech:py075b11e
