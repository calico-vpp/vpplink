module github.com/calico-vpp/vpplink

go 1.14

require (
	git.fd.io/govpp.git v0.3.4
	github.com/lunixbochs/struc v0.0.0-20190916212049-a5c72983bc42
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894
)

replace github.com/lunixbochs/struc => ../struc
