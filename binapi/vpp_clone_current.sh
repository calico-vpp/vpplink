#!/bin/bash
if [ ! -d $1 ]; then
	git clone "https://gerrit.fd.io/r/vpp" $1
fi
cd $1
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/95/26695/1
git checkout FETCH_HEAD
