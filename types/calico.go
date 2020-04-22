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

	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/calico"
)

func ToCalicoProto(proto IPProto) uint8 {
	switch proto {
	case UDP:
		return uint8(calico.IP_API_PROTO_TCP)
	case TCP:
		return uint8(calico.IP_API_PROTO_TCP)
	case SCTP:
		return uint8(calico.IP_API_PROTO_SCTP)
	case ICMP:
		return uint8(calico.IP_API_PROTO_ICMP)
	case ICMP6:
		return uint8(calico.IP_API_PROTO_ICMP6)
	default:
		return ^uint8(0)
	}
}

func CalicoSwif(swIfIndex uint32) calico.InterfaceIndex {
	return calico.InterfaceIndex(swIfIndex)
}

func ToCalicoVppIpAddress(addr net.IP) calico.Address {
	a := calico.Address{}
	if addr.To4() == nil {
		a.Af = calico.ADDRESS_IP6
		ip := [16]uint8{}
		copy(ip[:], addr)
		a.Un = calico.AddressUnionIP6(ip)
	} else {
		a.Af = calico.ADDRESS_IP4
		ip := [4]uint8{}
		copy(ip[:], addr.To4())
		a.Un = calico.AddressUnionIP4(ip)
	}
	return a
}

func ToCalicoVppIpAddressWithPrefix(addr *net.IPNet) calico.AddressWithPrefix {
	prefixLen, _ := addr.Mask.Size()
	return calico.AddressWithPrefix{
		Address: ToCalicoVppIpAddress(addr.IP),
		Len:     uint8(prefixLen),
	}
}
