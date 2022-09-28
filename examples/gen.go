package consumer

import (
	_ "go.fd.io/govpp/binapi"
)

//go:generate go build -buildmode=plugin -o ./.bin/vpplink_plugin.so github.com/calico-vpp/vpplink/pkg
//go:generate go run go.fd.io/govpp/cmd/binapi-generator --plugin ./.bin/vpplink_plugin.so --vpp $VPP_DIR -o ./impl
