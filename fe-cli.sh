#!/bin/bash
 
DEFAULT_ARGS="--rm -ti -v $PWD:/src -v ${PWD}/src/app:/gobuild/src -p 7757:7757 -e GOPATH=/gobuild"
if [ -f "${PWD}/.netrc" ]; then
    DEFAULT_ARGS="${DEFAULT_ARGS} -v ${PWD}/.netrc:/root/.netrc"
elif [ -f "${HOME}/.netrc" ]; then
    DEFAULT_ARGS="${DEFAULT_ARGS} -v ${HOME}/.netrc:/root/.netrc"
fi
docker run ${DEFAULT_ARGS} tci-flogo-cli:latest $@