#!/bin/bash

VPP_COMMIT=4128c7b46

if [ ! -d $1 ]; then
	git clone "https://gerrit.fd.io/r/vpp" $1
	cd $1
	git reset --hard ${VPP_COMMIT}
else
	cd $1
	git fetch "https://gerrit.fd.io/r/vpp" && git reset --hard ${VPP_COMMIT}
fi

git fetch "https://gerrit.fd.io/r/vpp" refs/changes/10/25810/34 && git cherry-pick FETCH_HEAD # GRO (coalesce)
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/04/27104/8 && git cherry-pick FETCH_HEAD # TAP GRO
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/62/27162/13 && git cherry-pick FETCH_HEAD # calico_plugin
