#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Excludes:
#   - when using defer there is no way to check to returned value so ignore
#   - some generated code has output parameters named as err that result in vetshadow issue so ignore
gometalinter \
    --skip=.glide \
    --exclude='^core/raw\.go:.*Subprocess launching with variable\.,HIGH,HIGH \(gas\)$' \
    --disable=aligncheck \
    --disable=gotype \
    --disable=structcheck \
    --disable=varcheck \
    --disable=interfacer \
    --disable=unconvert \
    --disable=dupl \
    --cyclo-over=15 \
    --tests \
    --deadline=60s \
    --vendor \
    ./...

# The following checks get into lower level detailed analysis of the code, but they
# all are common in that they scan the project differently then the others above that
# accept a simple base path recursion.
gometalinter \
    --skip=.glide \
    --tests \
    --vendor \
    --deadline=60s \
    --disable-all \
    --enable=unused \
    --enable=structcheck \
    --enable=varcheck \
    --enable=interfacer \
    --enable=unconvert
