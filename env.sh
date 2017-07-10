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
    if [ "$1" == "travis" ]; then
        parent=$(cd $APPROOT/../; pwd)

        if [ ! -d "$parent/gopkg/src/github.com/kirk-enterprise/openstack" ]; then
            mkdir -p "$parent/gopkg/src/github.com/kirk-enterprise"

            cp -r "$APPROOT" "$parent/gopkg/src/github.com/kirk-enterprise"
        fi
    else
        if [ ! -d "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk" ]; then
            mkdir -p "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk"

            FileList=`ls $APPROOT/.`
            for dir in ${FileList};do
                if [ ${dir} != "src" ];then
                    if [ -d "$APPROOT/${dir}" ];then
                        ln -s "$APPROOT/${dir}" "$APPROOT/src/github.com/kirk-enterprise/openstack-golang-sdk/${dir}"
                    fi
                fi
            done
        fi
    fi
fi
