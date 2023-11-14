#!/bin/bash

DOCKER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# cd $DOCKER_DIR && docker build -f Dockerfile --tag text-to-speech:e3b4f2527 ../

# docker scout cves --type image text-to-speech:e3b4f2527

# docker run --name text-to-speech-e3b4f2527-speech:e3b4f2527

cd $DOCKER_DIR && docker build -f Dockerfile --tag babbili/text-to-speech:e3b4f2527 ../

# docker scout cves --type image babbili/text-to-speech:e3b4f2527

docker push babbili/text-to-speech:e3b4f2527
