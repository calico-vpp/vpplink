package examples

import (
	_ "github.com/calico-vpp/vpplink/pkg"
	_ "go.fd.io/govpp/cmd/binapi-generator"
)

//go:generate go build -buildmode=plugin -o ./.bin/vpplink_plugin.so github.com/calico-vpp/vpplink/pkg
//go:generate go run go.fd.io/govpp/cmd/binapi-generator --gen ./.bin/vpplink_plugin.so,rpc --input ../../vpp -o ./impl
