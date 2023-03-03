//go:build generate

package examples

import (
	_ "go.fd.io/govpp/cmd/binapi-generator"

	_ "github.com/calico-vpp/vpplink/pkg"
)

//go:generate go build -buildmode=plugin -o ./.bin/vpplink_plugin.so github.com/calico-vpp/vpplink/pkg
//go:generate go run go.fd.io/govpp/cmd/binapi-generator --gen ./.bin/vpplink_plugin.so,rpc --input ${VPP_DIR:-} -o ./impl
