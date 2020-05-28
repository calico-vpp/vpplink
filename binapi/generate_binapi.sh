#!/bin/bash

set -e

SOURCE="${BASH_SOURCE[0]}"
SCRIPTDIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

echo "Input VPP full path : "
read VPP_DIR

if [[ ! -d $VPP_DIR ]]; then
	echo "Couldnt find anything at <$VPP_DIR>"
	exit 1
fi

pushd $VPP_DIR > /dev/null
VPP_REMOTE_NAME=$(echo `git remote`)
if [[ x$(git remote | wc -l) != x1 ]]; then
	echo "Input VPP's remote [ $VPP_REMOTE_NAME ] : "
	echo "ommitting won't update ./vpp_clone_current.sh"
	read VPP_REMOTE_NAME
fi

VPP_VERSION=$(./build-root/scripts/version)
echo "Using commit : $VPP_COMMIT"
make json-api-files
popd > /dev/null

find $VPP_DIR/build-root/install-vpp-native/vpp/share/vpp/api/ -name '*.json' \
	-exec binapi-generator --input-file={} --output-dir=$SCRIPTDIR/$VPP_VERSION \;

echo "Update version number with $VPP_VERSION ? [yes/no] "
read RESP

if [[ x$RESP = xyes ]]; then
	find . -path ./binapi -prune -o -name '*.go' \
		-exec sed -i 's@github.com/calico-vpp/vpplink/binapi/[.~0-9a-z_-]*/'"@github.com/calico-vpp/vpplink/binapi/$VPP_VERSION/@g" {} \;
fi

if [[ x$VPP_REMOTE_NAME = x ]]; then
	exit 0
fi

VPP_COMMIT=$(git rev-parse --short HEAD)
VPP_REMOTE_URL=$(git config --get remote.$VPP_REMOTE_NAME.url)

echo "#!/bin/bash
if [ ! -d \$1 ]; then
	git clone $VPP_REMOTE_URL \$1
fi
cd \$1
git fetch $VPP_REMOTE_URL
git checkout $VPP_COMMIT" > $SCRIPTDIR/vpp_clone_current.sh
chmod +x $SCRIPTDIR/vpp_clone_current.sh

echo "Using remote : $VPP_REMOTE_URL"
echo "Using commit : $VPP_COMMIT"
