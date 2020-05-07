#!/bin/bash
if [ ! -d $1 ]; then
	git clone "https://gerrit.fd.io/r/vpp" $1
fi
cd $1
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/95/26695/2
git checkout FETCH_HEAD

# GRO patch
git fetch "ssh://aloys@gerrit.fd.io:29418/vpp" refs/changes/10/25810/19 && git cherry-pick FETCH_HEAD

# NAT44 multi-worker patch
# git fetch "ssh://aloys@gerrit.fd.io:29418/vpp" refs/changes/24/26924/1 && git cherry-pick FETCH_HEAD


