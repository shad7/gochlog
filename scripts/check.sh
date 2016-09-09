#!/bin/bash

for pkg in $(glide nv); do
    errcheck $pkg
done
