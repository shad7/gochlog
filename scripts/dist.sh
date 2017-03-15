#!/bin/bash

ROOT_DIR=$(pwd)

# Zip and copy to the release dir
echo "==> Packaging..."
for PLATFORM in $(find ./build -mindepth 1 -maxdepth 1 -type d); do
    OSARCH=$(basename ${PLATFORM})
    echo "--> ${OSARCH}"

    pushd $PLATFORM >/dev/null 2>&1
    zip ${ROOT_DIR}/release/${OSARCH}.zip ./*
    popd >/dev/null 2>&1
done

pushd ./release >/dev/null 2>&1
shasum -a256 * > ./${PROJECT}_${VERSION}_SHA256SUMS
popd >/dev/null 2>&1
