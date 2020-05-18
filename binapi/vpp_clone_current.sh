#!/bin/bash
if [ ! -d $1 ]; then
	git clone https://gerrit.fd.io/r/vpp $1
fi
cd $1

git fetch origin master
git reset --hard 59a78e966c7813e1bd86310bc0edd8e60a506ccf
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/10/25810/23 && git cherry-pick FETCH_HEAD # GRO
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/04/27104/2 && git cherry-pick FETCH_HEAD  # GRO-tap
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/27/27127/1 && git cherry-pick FETCH_HEAD  # SA pinning
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/95/26695/2 && git cherry-pick FETCH_HEAD  # Calico plugin
