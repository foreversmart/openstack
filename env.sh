#!/usr/bin/env bash

if [ "$APPROOT" = "" ]; then
    export APPROOT=$(pwd)
fi

if [ "$APPGOPATH" = "" ]; then
    export APPGOPATH=$(dirname $(pwd))/gopkg
fi

# adjust GOPATH
case ":$GOPATH:" in
    *":$APPGOPATH:"*) :;;
    *) GOPATH=$APPGOPATH:$GOPATH;;
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
    if [ "$1" = "travis" ]; then
        parent=$(cd $APPROOT/../; pwd)

        if [ ! -d "$parent/gopkg/src/github.com/kirk-enterprise/openstack-golang-sdk" ]; then
            mkdir -p "$parent/gopkg/src/github.com/kirk-enterprise"

            cp -r "$APPROOT" "$parent/gopkg/src/github.com/kirk-enterprise"
        fi
    else
        if [ ! -d "$APPGOPATH/src/github.com/kirk-enterprise/openstack-golang-sdk" ]; then
            mkdir -p "$APPGOPATH/src/github.com/kirk-enterprise"

            ln -s "$APPROOT" "$APPGOPATH/src/github.com/kirk-enterprise"
        fi
    fi
fi
