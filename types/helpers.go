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

func ToVppIpAddress(addr net.IP) vppip.Address {
	a := vppip.Address{}
	if addr.To4() == nil {
		a.Af = vppip.ADDRESS_IP6
		ip := [16]uint8{}
		copy(ip[:], addr)
		a.Un = vppip.AddressUnionIP6(ip)
	} else {
		a.Af = vppip.ADDRESS_IP4
		ip := [4]uint8{}
		copy(ip[:], addr.To4())
		a.Un = vppip.AddressUnionIP4(ip)
	}
	return a
}

func ToVppMacAddress(hardwareAddr net.HardwareAddr) vppip.MacAddress {
	hwAddr := [6]uint8{}
	copy(hwAddr[:], hardwareAddr)
	return vppip.MacAddress(hwAddr)
}

func TobytesMacAddress(hardwareAddr net.HardwareAddr) []byte {
	hwAddr := [6]uint8{}
	copy(hwAddr[:], hardwareAddr)
	return hwAddr[:]
}
