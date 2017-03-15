#!/bin/bash

echo "==> Generating docs..."

for pkg in $(glide nv);
do
    for subpkg in $(go list ${pkg});
    do
        SUBPKG_DIR=${subpkg#$IMPORT_PATH}
        echo "--> ${subpkg}"
        mkdir -p ./docs/${SUBPKG_DIR}
        godoc2md ${subpkg} > ./docs$SUBPKG_DIR/README.md
    done
done

echo "docs ready"
