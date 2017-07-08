#!/usr/bin/env bash

if [ "$APPROOT" = "" ]; then
    export APPROOT=$(pwd)
fi

# adjust GOPATH
case ":$GOPATH:" in
    *":$APPROOT:"*) :;;
    *) GOPATH=$APPROOT:$GOPATH;;
esac
export GOPATH


# adjust PATH
if [ -n "$ZSH_VERSION" ]; then
    readopts="rA"
else
    readopts="ra"
fi
while IFS=':' read -$readopts ARR; do
    for i in "${ARR[@]}"; do
        case ":$PATH:" in
            *":$i/bin:"*) :;;
            *) PATH=$i/bin:$PATH
        esac
    done
done <<< "$GOPATH"
export PATH


# mock development && test envs
if [ -f "$APPROOT/openstack.go" ]; then
    if [ ! -d "$APPROOT/src/github.com/kirk-enterprise/openstack" ]; then
        mkdir -p "$APPROOT/src/github.com/kirk-enterprise"

        if [ "$1" == "travis" ]; then
            cp -r "$APPROOT" "$APPROOT/src/github.com/kirk-enterprise"
        else
            ln -s "$APPROOT" "$APPROOT/src/github.com/kirk-enterprise"
        fi
    fi
fi
