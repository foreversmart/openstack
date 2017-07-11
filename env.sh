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
    if [ "$1" = "travis" ]; then
        parent=$(cd $APPROOT/../; pwd)

        if [ ! -d "$parent/gopkg/src/github.com/kirk-enterprise/openstack-golang-sdk" ]; then
            mkdir -p "$parent/gopkg/src/github.com/kirk-enterprise"

            cp -r "$APPROOT" "$parent/gopkg/src/github.com/kirk-enterprise"
        fi
    else
        if [ ! -d "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk" ]; then
            mkdir -p "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk"

            files=`ls $APPROOT/.`
            for file in $files; do
                if  [ -d "$APPROOT/$file" ] && [ "$file" != "src" ] && [ "$file" != "src/" ] ; then
                    ln -s "$APPROOT/$file" "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk/"
                fi

                if [ -f "$APPROOT/$file" ]; then
                    ln -s "$APPROOT/$file" "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk/"
                fi
            done
        fi
    fi
fi
