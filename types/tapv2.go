// Copyright (C) 2019 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"net"

	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/tapv2"
)

type TapV2 struct {
	HostNamespace  string
	HostIfName     string
	Tag            string
	MacAddress     net.HardwareAddr
	HostMacAddress net.HardwareAddr
	RxQueues       int
}

func (t *TapV2) GetVppHostMacAddress() tapv2.MacAddress {
	return tapv2.MacAddress(ToVppMacAddress(t.HostMacAddress))
}

func (t *TapV2) GetVppMacAddress() tapv2.MacAddress {
	return tapv2.MacAddress(ToVppMacAddress(t.MacAddress))
}
