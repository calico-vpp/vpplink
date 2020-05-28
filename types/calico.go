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
	"fmt"
	"net"

	"github.com/calico-vpp/vpplink/binapi/20.09-rc0~76-g6ec3f62e7/calico"
)

type CalicoTranslateEntry struct {
	SrcPort    uint16
	Vip        net.IP
	DestPort   uint16
	ID         uint32
	BackendIPs []net.IP
	Proto      IPProto
}

func (n *CalicoTranslateEntry) String() string {
	return fmt.Sprintf("%s %s:%d -> %+v:%d",
		formatProto(n.Proto),
		n.Vip.String(),
		n.SrcPort,
		n.BackendIPs,
		n.DestPort,
	)
}

func ToCalicoProto(proto IPProto) calico.IPProto {
	switch proto {
	case UDP:
		return calico.IP_API_PROTO_TCP
	case TCP:
		return calico.IP_API_PROTO_TCP
	case SCTP:
		return calico.IP_API_PROTO_SCTP
	case ICMP:
		return calico.IP_API_PROTO_ICMP
	case ICMP6:
		return calico.IP_API_PROTO_ICMP6
	default:
		return calico.IP_API_PROTO_RESERVED
	}
}

func ToCalicoEndpoint(addr net.IP, port uint16) calico.CalicoEndpoint {
	a := calico.CalicoEndpoint{
		Port: port,
	}
	if addr.To4() == nil {
		a.Addr.Af = calico.ADDRESS_IP6
		ip := [16]uint8{}
		copy(ip[:], addr)
		a.Addr.Un = calico.AddressUnionIP6(ip)
	} else {
		a.Addr.Af = calico.ADDRESS_IP4
		ip := [4]uint8{}
		copy(ip[:], addr.To4())
		a.Addr.Un = calico.AddressUnionIP4(ip)
	}
	return a
}
