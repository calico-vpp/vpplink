#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
SCRIPTDIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

echo "Input VPP's location : "
read VPP_DIR

pushd $VPP_DIR
VPP_VERSION=$(./build-root/scripts/version)
# make json-api-files
popd

find $VPP_DIR/build-root/install-vpp-native/vpp/share/vpp/api/ -name '*.json' \
	-exec binapi-generator --input-file={} --output-dir=$SCRIPTDIR/$VPP_VERSION \;

echo "Update version number with $VPP_VERSION ? [yes/no] "
read RESP

if [[ x$RESP = xyes ]]; then
	find .. -path ./binapi -prune -o -name '*.go' \
		-exec sed -i "s@github.com/calico-vpp/vpplink/binapi/[0-9a-z_~-]*/@github.com/calico-vpp/vpplink/binapi/$VPP_VERSION/@g" {} \;
fi

