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

	vppip "github.com/calico-vpp/vpplink/binapi/20_05-rc0-540-g77ea42b/ip"
)

type IPProto uint32

const (
	INVALID IPProto = 0
	UDP     IPProto = 1
	SCTP    IPProto = 2
	TCP     IPProto = 3
)

type IfAddress struct {
	IPNet     net.IPNet
	SwIfIndex uint32
}

func ToVppIPProto(proto IPProto) uint8 {
	switch proto {
	case UDP:
		return uint8(vppip.IP_API_PROTO_TCP)
	case TCP:
		return uint8(vppip.IP_API_PROTO_TCP)
	case SCTP:
		return uint8(vppip.IP_API_PROTO_SCTP)
	default:
		return ^uint8(0)
	}
}
