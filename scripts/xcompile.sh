#!/bin/bash

gox \
    -ldflags "${LDFLAGS}" \
    -arch="amd64" \
    -arch="386" \
    -os="darwin" \
    -os="linux" \
    -os="windows" \
    -output="build/{{.Dir}}_${VERSION}_{{.OS}}_{{.Arch}}/gochlog"
