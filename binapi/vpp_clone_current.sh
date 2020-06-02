#!/bin/bash
if [ ! -d $1 ]; then
	git clone "https://gerrit.fd.io/r/vpp" $1
	cd $1
	git reset --hard dcd4aa211
else
	cd $1
	git fetch "https://gerrit.fd.io/r/vpp" && git reset --hard dcd4aa211
fi
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/04/27104/3 && git cherry-pick FETCH_HEAD # GSO checksum fix
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/10/25810/26 && git cherry-pick FETCH_HEAD # GRO (coalesce)
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/85/27085/1 && git cherry-pick FETCH_HEAD # ikev2 cross tunnel fix
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/69/27269/1 && git cherry-pick FETCH_HEAD # tap interrupt
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/62/27162/9 && git cherry-pick FETCH_HEAD # calico_plugin
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/87/27387/1 && git cherry-pick FETCH_HEAD # Fix get route API
