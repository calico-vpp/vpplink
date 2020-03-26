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

	vppip "github.com/calico-vpp/vpplink/binapi/19_08/ip"
)

type Route struct {
	Dst       *net.IPNet
	Gw        net.IP
	Table     int
	DstTable  int
	SwIfIndex uint32
}

func (r *Route) GetVppGwAddress() vppip.Address {
	return ToVppIpAddress(r.Gw)
}

func (r *Route) GetVppDstAddress() vppip.Address {
	return ToVppIpAddress(r.Dst.IP)
}
