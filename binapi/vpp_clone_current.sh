#!/bin/bash
if [ ! -d $1 ]; then
	git clone ssh://sknat@gerrit.fd.io:29418/vpp $1
fi
cd $1
git fetch ssh://sknat@gerrit.fd.io:29418/vpp
git checkout 6f2c5a55f
