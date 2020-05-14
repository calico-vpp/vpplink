#!/bin/bash
if [ ! -d $1 ]; then
	git clone https://gerrit.fd.io/r/vpp $1
fi
cd $1
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/24/27024/2 && git checkout FETCH_HEAD
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/10/25810/21 && git cherry-pick FETCH_HEAD # GRO
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/06/27006/4 && git cherry-pick FETCH_HEAD # IPIP
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/17/27017/6 && git cherry-pick FETCH_HEAD # NAT