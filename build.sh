#!/bin/sh

die(){
    echo "ERROR: $*" 1>&2
    exit 1
}

[ -z "`which tinygo`" ] && die "Unable to find the TinyGo embedded compiler in path."

# This uses TinyGo to output proper WASM/WASI
tinygo build -target=wasi \
    -scheduler=none -gc=conservative \
    -o contract.wasm ./...
