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
	"bytes"
	"fmt"
	"net"

	"github.com/calico-vpp/vpplink/binapi/19_08/interfaces"
	vppip "github.com/calico-vpp/vpplink/binapi/19_08/ip"
	"github.com/calico-vpp/vpplink/binapi/19_08/tapv2"
	"github.com/calico-vpp/vpplink/types"
	"github.com/pkg/errors"
)

const (
	INVALID_SW_IF_INDEX = ^uint32(0)
)

func (v *VppLink) CreateTapV2(tap *types.TapV2) (SwIfIndex uint32, err error) {
	v.lock.Lock()
	defer v.lock.Unlock()
	response := &tapv2.TapCreateV2Reply{}
	request := &tapv2.TapCreateV2{
		// TODO check namespace len < 64?
		// TODO set MTU?
		ID:               ^uint32(0),
		HostNamespace:    []byte(tap.ContNS),
		HostNamespaceSet: 1,
		HostIfName:       []byte(tap.ContIfName),
		HostIfNameSet:    1,
		Tag:              []byte(tap.Tag),
		MacAddress:       tap.GetVppMacAddress(),
		HostMacAddr:      tap.GetVppHostMacAddress(),
		HostMacAddrSet:   1,
	}

	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return INVALID_SW_IF_INDEX, errors.Wrap(err, "Tap creation request failed")
	} else if response.Retval != 0 {
		return INVALID_SW_IF_INDEX, fmt.Errorf("Tap creation failed (retval %d). Request: %+v", response.Retval, request)
	}
	return response.SwIfIndex, err

	// Add VPP side fake address
	// TODO: Only if v4 is enabled
	// There is currently a hard limit in VPP to 1024 taps - so this should be safe
	// vppIPAddress = []byte{169, 254, byte(response.SwIfIndex >> 8), byte(response.SwIfIndex)}
	// err = v.AddInterfaceAddress(response.SwIfIndex, vppIPAddress, 32)
	// if err != nil {
	// 	return INVALID_SW_IF_INDEX, vppIPAddress, errors.Wrap(err, "error adding address to new tap")
	// }

	// // Set interface up
	// err = v.InterfaceAdminUp(response.SwIfIndex)
	// if err != nil {
	// 	return INVALID_SW_IF_INDEX, vppIPAddress, errors.Wrap(err, "error setting new tap up")
	// }

	// // Add IPv6 neighbor entry if v6 is enabled
	// if EnableIp6 {
	// 	err = v.EnableInterfaceIP6(response.SwIfIndex)
	// 	if err != nil {
	// 		return INVALID_SW_IF_INDEX, vppIPAddress, errors.Wrap(err, "error enabling IPv6 on new tap")
	// 	}
	// 	// Compute a link local address from mac address, and set it
	// }
	// return response.SwIfIndex, vppIPAddress, err
}

func (v *VppLink) addDelInterfaceAddress(swIfIndex uint32, addr *net.IPNet, isAdd uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	addrLen, _ := addr.Mask.Size()
	request := &interfaces.SwInterfaceAddDelAddress{
		SwIfIndex:     swIfIndex,
		IsAdd:         isAdd,
		IsIPv6:        BoolToU8(IsIP6(addr.IP)),
		AddressLength: uint8(addrLen),
	}
	if IsIP4(addr.IP) {
		request.Address = addr.IP.To4()
	} else {
		request.Address = addr.IP.To16()
	}

	response := &interfaces.SwInterfaceAddDelAddressReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "Adding IP address failed: req %+v reply %+v", request, response)
	}
	return nil
}

func (v *VppLink) DelInterfaceAddress(swIfIndex uint32, addr *net.IPNet) error {
	return v.addDelInterfaceAddress(swIfIndex, addr, 0)
}

func (v *VppLink) AddInterfaceAddress(swIfIndex uint32, addr *net.IPNet) error {
	return v.addDelInterfaceAddress(swIfIndex, addr, 1)
}

func (v *VppLink) enableDisableInterfaceIP6(swIfIndex uint32, enable uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.SwInterfaceIP6EnableDisable{
		SwIfIndex: swIfIndex,
		Enable:    enable,
	}
	response := &vppip.SwInterfaceIP6EnableDisableReply{}
	return v.ch.SendRequest(request).ReceiveReply(response)
}

func (v *VppLink) DisableInterfaceIP6(swIfIndex uint32) error {
	return v.enableDisableInterfaceIP6(swIfIndex, 0)
}

func (v *VppLink) EnableInterfaceIP6(swIfIndex uint32) error {
	return v.enableDisableInterfaceIP6(swIfIndex, 1)
}

func (v *VppLink) SearchInterfaceWithTag(tag string) (err error, swIfIndex uint32) {
	v.lock.Lock()
	defer v.lock.Unlock()

	swIfIndex = INVALID_SW_IF_INDEX
	request := &interfaces.SwInterfaceDump{}
	stream := v.ch.SendMultiRequest(request)
	for {
		response := &interfaces.SwInterfaceDetails{}
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP interfaces: %v", err)
			return err, INVALID_SW_IF_INDEX
		}
		if stop {
			break
		}
		intfTag := string(bytes.Trim([]byte(response.Tag), "\x00"))
		v.log.Debugf("found interface %d, tag: %s (len %d)", response.SwIfIndex, intfTag, len(intfTag))
		if intfTag == tag {
			swIfIndex = response.SwIfIndex
		}
	}
	if swIfIndex == INVALID_SW_IF_INDEX {
		v.log.Errorf("Interface with tag %s not found", tag)
		return errors.New("Interface not found"), INVALID_SW_IF_INDEX
	}
	return nil, swIfIndex
}

func (v *VppLink) SearchInterfaceWithName(name string) (err error, swIfIndex uint32) {
	v.lock.Lock()
	defer v.lock.Unlock()

	swIfIndex = INVALID_SW_IF_INDEX
	request := &interfaces.SwInterfaceDump{
		SwIfIndex: interfaces.InterfaceIndex(INVALID_SW_IF_INDEX),
		// TODO: filter by name with NameFilter
	}
	reqCtx := v.ch.SendMultiRequest(request)
	for {
		response := &interfaces.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("SwInterfaceDump failed: %v", err)
			return err, INVALID_SW_IF_INDEX
		}
		if stop {
			break
		}
		interfaceName := string(bytes.Trim([]byte(response.InterfaceName), "\x00"))
		v.log.Debugf("Found interface: %s", interfaceName)
		if interfaceName == name {
			swIfIndex = response.SwIfIndex
		}

	}
	if swIfIndex == INVALID_SW_IF_INDEX {
		v.log.Errorf("Interface %s not found", name)
		return errors.New("Interface not found"), INVALID_SW_IF_INDEX
	}
	return nil, swIfIndex
}

func (v *VppLink) interfaceAdminUpDown(swIfIndex uint32, updown uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	// Set interface down
	request := &interfaces.SwInterfaceSetFlags{
		SwIfIndex:   swIfIndex,
		AdminUpDown: updown,
	}
	response := &interfaces.SwInterfaceSetFlagsReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "setting interface down failed")
	}
	return nil
}

func (v *VppLink) InterfaceAdminDown(swIfIndex uint32) error {
	return v.interfaceAdminUpDown(swIfIndex, 0)
}

func (v *VppLink) InterfaceAdminUp(swIfIndex uint32) error {
	return v.interfaceAdminUpDown(swIfIndex, 1)
}

func (v *VppLink) GetInterfaceNeighbors(swIfIndex uint32, isIPv6 uint8) (err error, neighbors []vppip.IPNeighbor) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.IPNeighborDump{
		SwIfIndex: swIfIndex,
		IsIPv6:    isIPv6,
	}
	response := &vppip.IPNeighborDetails{}
	stream := v.ch.SendMultiRequest(request)
	for {
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP neighbors: %v", err)
			return err, nil
		}
		if stop {
			return nil, neighbors
		}
		neighbors = append(neighbors, response.Neighbor)
	}
}

func (v *VppLink) DelTap(swIfIndex uint32) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &tapv2.TapDeleteV2{
		SwIfIndex: swIfIndex,
	}
	response := &tapv2.TapDeleteV2Reply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "failed to delete tap from VPP")
	}
	return nil
}

func (v *VppLink) interfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32, isAdd uint8) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	// Set interface down
	request := &interfaces.SwInterfaceSetUnnumbered{
		SwIfIndex:           swIfIndex,
		UnnumberedSwIfIndex: unnumberedSwIfIndex,
		IsAdd:               isAdd,
	}
	response := &interfaces.SwInterfaceSetUnnumberedReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "setting interface unnumbered failed %d -> %d", unnumberedSwIfIndex, swIfIndex)
	}
	return nil
}

func (v *VppLink) InterfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error {
	return v.interfaceSetUnnumbered(unnumberedSwIfIndex, swIfIndex, 1)
}

func (v *VppLink) InterfaceUnsetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error {
	return v.interfaceSetUnnumbered(unnumberedSwIfIndex, swIfIndex, 0)
}

func (v *VppLink) PuntRedirect(sourceSwIfIndex, destSwIfIndex uint32, nh net.IP) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	request := &vppip.IPPuntRedirect{
		Punt: vppip.PuntRedirect{
			RxSwIfIndex: sourceSwIfIndex,
			TxSwIfIndex: destSwIfIndex,
			Nh:          types.ToVppIpAddress(nh),
		},
		IsAdd: 1,
	}
	response := &vppip.IPPuntRedirectReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil || response.Retval != 0 {
		return fmt.Errorf("cannot set punt in VPP: %v %d", err, response.Retval)
	}
	return nil
}
