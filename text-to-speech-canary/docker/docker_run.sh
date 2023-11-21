#!/bin/bash

DOCKER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# cd $DOCKER_DIR && docker build -f Dockerfile --tag text-to-speech-py:79931c98a ../

# docker scout cves --type image text-to-speech-py:79931c98a

# docker run --name text-to-speech-py-79931c98a-speech-py:79931c98a

cd $DOCKER_DIR && docker build -f Dockerfile --tag babbili/text-to-speech-py:79931c98a ../

# docker scout cves --type image babbili/text-to-speech-py:79931c98a

docker push babbili/text-to-speech-py:79931c98a
