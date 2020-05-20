#!/bin/bash
if [ ! -d $1 ]; then
	git clone "https://gerrit.fd.io/r/vpp" $1
fi
cd $1
git reset --hard dcd4aa211
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/04/27104/3 && git cherry-pick FETCH_HEAD
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/10/25810/26 && git cherry-pick FETCH_HEAD
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/85/27085/1 && git cherry-pick FETCH_HEAD
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/69/27269/1 && git cherry-pick FETCH_HEAD
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/62/27162/7 && git cherry-pick FETCH_HEAD
