#!/bin/bash

gox \
    -ldflags "${LDFLAGS}" \
    -arch="amd64" \
    -arch="386" \
    -os="darwin" \
    -os="freebsd" \
    -os="linux" \
    -os="netbsd" \
    -os="openbsd" \
    -os="windows" \
    -output="build/{{.Dir}}_${VERSION}_{{.OS}}_{{.Arch}}/gochlog"
