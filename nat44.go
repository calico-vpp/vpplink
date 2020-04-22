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

package vpplink

import (
	"fmt"
	"net"

	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/nat"
	"github.com/calico-vpp/vpplink/types"
	"github.com/pkg/errors"
)

func parseIP4Address(address string) nat.IP4Address {
	var ip nat.IP4Address
	copy(ip[:], net.ParseIP(address).To4()[0:4])
	return ip
}

func (v *VppLink) EnableNatForwarding() (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &nat.Nat44ForwardingEnableDisableReply{}
	request := &nat.Nat44ForwardingEnableDisable{
		Enable: true,
	}
	v.log.Debug("Enabling NAT44 forwarding")
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "NAT44 forwarding enable failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("NAT44 forwarding enable failed with retval: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) addDelNat44Address(isAdd bool, address string) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &nat.Nat44AddDelAddressRangeReply{}
	request := &nat.Nat44AddDelAddressRange{
		FirstIPAddress: parseIP4Address(address),
		LastIPAddress:  parseIP4Address(address),
		VrfID:          0,
		IsAdd:          isAdd,
		Flags:          nat.NAT_IS_NONE,
	}
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Nat44 address add failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Nat44 address add failed with retval %d", response.Retval)
	}
	return nil
}

func (v *VppLink) AddNat44InterfaceAddress(swIfIndex uint32, flags types.NatFlags) error {
	return v.addDelNat44InterfaceAddress(true, swIfIndex, flags)
}

func (v *VppLink) DelNat44InterfaceAddress(swIfIndex uint32, flags types.NatFlags) error {
	return v.addDelNat44InterfaceAddress(false, swIfIndex, flags)
}

func (v *VppLink) addDelNat44InterfaceAddress(isAdd bool, swIfIndex uint32, flags types.NatFlags) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &nat.Nat44AddDelInterfaceAddrReply{}
	request := &nat.Nat44AddDelInterfaceAddr{
		IsAdd:     isAdd,
		SwIfIndex: nat.InterfaceIndex(swIfIndex),
		Flags:     types.ToVppNatConfigFlags(flags),
	}
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Nat44 addDel interface address failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Nat44 addDel interface address failed: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) AddNat44Address(address string) error {
	return v.addDelNat44Address(true, address)
}

func (v *VppLink) DelNat44Address(address string) error {
	return v.addDelNat44Address(false, address)
}

func (v *VppLink) addDelNat44Interface(isAdd bool, flags types.NatFlags, swIfIndex uint32) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &nat.Nat44InterfaceAddDelFeatureReply{}
	request := &nat.Nat44InterfaceAddDelFeature{
		IsAdd:     isAdd,
		Flags:     types.ToVppNatConfigFlags(flags),
		SwIfIndex: nat.InterfaceIndex(swIfIndex),
	}
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Nat44 addDel interface failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Nat44 addDel interface failed: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) AddNat44InsideInterface(swIfIndex uint32) error {
	return v.addDelNat44Interface(true, types.NatInside, swIfIndex)
}

func (v *VppLink) AddNat44OutsideInterface(swIfIndex uint32) error {
	return v.addDelNat44Interface(true, types.NatOutside, swIfIndex)
}

func (v *VppLink) DelNat44InsideInterface(swIfIndex uint32) error {
	return v.addDelNat44Interface(false, types.NatInside, swIfIndex)
}

func (v *VppLink) DelNat44OutsideInterface(swIfIndex uint32) error {
	return v.addDelNat44Interface(false, types.NatOutside, swIfIndex)
}

func (v *VppLink) getLBLocals(backends []string, port int32) (locals []nat.Nat44LbAddrPort) {
	for _, ip := range backends {
		v.log.Debugf("Adding local %s:%d", ip, port)
		locals = append(locals, nat.Nat44LbAddrPort{
			Addr:        parseIP4Address(ip),
			Port:        uint16(port),
			Probability: uint8(10),
		})
	}
	return locals
}

func (v *VppLink) addDelNat44LBStaticMapping(
	isAdd bool,
	extAddr string,
	proto types.IPProto,
	extPort int32,
	backends []string,
	backendPort int32,
) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	locals := v.getLBLocals(backends, backendPort)
	response := &nat.Nat44AddDelLbStaticMappingReply{}
	request := &nat.Nat44AddDelLbStaticMapping{
		IsAdd:        isAdd,
		Flags:        nat.NAT_IS_SELF_TWICE_NAT | nat.NAT_IS_OUT2IN_ONLY,
		ExternalAddr: parseIP4Address(extAddr),
		ExternalPort: uint16(extPort),
		Protocol:     types.ToVppIPProto(proto),
		Locals:       locals,
	}
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Nat44 add LB static failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Nat44 add LB static failed: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) AddNat44LBStaticMapping(
	externalAddr string,
	serviceProto types.IPProto,
	externalPort int32,
	backendIPs []string,
	backendPort int32,
) error {
	return v.addDelNat44LBStaticMapping(true, externalAddr, serviceProto, externalPort, backendIPs, backendPort)
}

func (v *VppLink) DelNat44LBStaticMapping(
	externalAddr string,
	serviceProto types.IPProto,
	externalPort int32,
) error {
	return v.addDelNat44LBStaticMapping(false, externalAddr, serviceProto, externalPort, []string{}, 0)
}

func (v *VppLink) addDelNat44StaticMapping(
	isAdd bool,
	externalAddr string,
	serviceProto types.IPProto,
	externalPort int32,
	backendIP string,
	backendPort int32,
) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &nat.Nat44AddDelStaticMappingReply{}
	request := &nat.Nat44AddDelStaticMapping{
		IsAdd:             isAdd,
		Flags:             nat.NAT_IS_SELF_TWICE_NAT | nat.NAT_IS_OUT2IN_ONLY,
		LocalIPAddress:    parseIP4Address(backendIP),
		ExternalIPAddress: parseIP4Address(externalAddr),
		Protocol:          types.ToVppIPProto(serviceProto),
		LocalPort:         uint16(backendPort),
		ExternalPort:      uint16(externalPort),
		ExternalSwIfIndex: 0xffffffff,
	}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Nat44 static mapping failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Nat44 add LB static failed: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) AddNat44StaticMapping(
	externalAddr string,
	serviceProto types.IPProto,
	externalPort int32,
	backendIP string,
	backendPort int32,
) error {
	return v.addDelNat44StaticMapping(true, externalAddr, serviceProto, externalPort, backendIP, backendPort)
}

func (v *VppLink) DelNat44StaticMapping(
	externalAddr string,
	serviceProto types.IPProto,
	externalPort int32,
) error {
	return v.addDelNat44StaticMapping(false, externalAddr, serviceProto, externalPort, "0.0.0.0", 0)
}

func (v *VppLink) AddNat44LB(
	serviceIP string,
	serviceProto types.IPProto,
	servicePort int32,
	backendIPs []string,
	backendPort int32,
) error {
	if len(backendIPs) == 0 {
		return fmt.Errorf("No backends provided for NAT44")
	}
	if len(backendIPs) == 1 {
		return v.AddNat44StaticMapping(serviceIP, serviceProto, servicePort, backendIPs[0], backendPort)
	}
	return v.AddNat44LBStaticMapping(serviceIP, serviceProto, servicePort, backendIPs, backendPort)
}

func (v *VppLink) DelNat44LB(
	serviceIP string,
	serviceProto types.IPProto,
	servicePort int32,
	backendCount int,
) error {
	if backendCount == 0 {
		return fmt.Errorf("No backends provided for NAT44")
	}
	if backendCount == 1 {
		return v.DelNat44StaticMapping(serviceIP, serviceProto, servicePort)
	}
	return v.DelNat44LBStaticMapping(serviceIP, serviceProto, servicePort)
}
