#!/bin/bash

result=0

for pkg in $(glide nv); do
    if ! golint -set_exit_status $pkg; then
        result=1
    fi
done

exit $result
