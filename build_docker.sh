#!/bin/bash

docker build -t dfis-utils . &&\
docker run --rm -e UID=$UID -v `pwd`/:/go/src/dfis-utils -it dfis-utils
