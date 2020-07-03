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

	vppnat "github.com/calico-vpp/vpplink/binapi/20.09-rc0~187-gf9d9cd97b/nat"
)

type NatFlags uint8

const (
	NatNone         NatFlags = NatFlags(vppnat.NAT_IS_NONE)
	NatTwice        NatFlags = NatFlags(vppnat.NAT_IS_TWICE_NAT)
	NatSelfTwice    NatFlags = NatFlags(vppnat.NAT_IS_SELF_TWICE_NAT)
	NatOut2In       NatFlags = NatFlags(vppnat.NAT_IS_OUT2IN_ONLY)
	NatAddrOnly     NatFlags = NatFlags(vppnat.NAT_IS_ADDR_ONLY)
	NatOutside      NatFlags = NatFlags(vppnat.NAT_IS_OUTSIDE)
	NatInside       NatFlags = NatFlags(vppnat.NAT_IS_INSIDE)
	NatStatic       NatFlags = NatFlags(vppnat.NAT_IS_STATIC)
	NatExtHostValid NatFlags = NatFlags(vppnat.NAT_IS_EXT_HOST_VALID)
)

func ToVppNatConfigFlags(flags NatFlags) vppnat.NatConfigFlags {
	return vppnat.NatConfigFlags(flags)
}

type Nat44Entry struct {
	ServiceIP   net.IP
	ServicePort int32
	Protocol    IPProto
	BackendIPs  []net.IP
	BackendPort int32
}

func (n *Nat44Entry) String() string {
	return fmt.Sprintf("%s %s:%d -> %+v:%d",
		formatProto(n.Protocol),
		n.ServiceIP.String(),
		n.ServicePort,
		n.BackendIPs,
		n.BackendPort,
	)
}
