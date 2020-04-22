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

	"github.com/calico-vpp/vpplink/binapi/19_08/calico"
	"github.com/calico-vpp/vpplink/types"
	"github.com/pkg/errors"
)

func (v *VppLink) calicoAddDelIfNat4(swIfIndex uint32, isAdd bool) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &calico.CalicoAddDelIntfNat4Reply{}
	request := &calico.CalicoAddDelIntfNat4{
		SwIfIndex: swIfIndex,
		IsAdd:     isAdd,
	}
	v.log.Debug("Add/del interface nat6")
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Add/del interface nat6 failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Add/del interface nat6 failed with retval: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) CalicoAddInterfaceNat4(swIfIndex uint32) (err error) {
	return v.calicoAddDelIfNat4(swIfIndex, true /* isAdd */)
}

func (v *VppLink) CalicoDelInterfaceNat4(swIfIndex uint32) (err error) {
	return v.calicoAddDelIfNat4(swIfIndex, false /* isAdd */)
}

// --------------

func (v *VppLink) calicoAddDelIfNat6(swIfIndex uint32, isAdd bool) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &calico.CalicoAddDelIntfNat6Reply{}
	request := &calico.CalicoAddDelIntfNat6{
		SwIfIndex: swIfIndex,
		IsAdd:     isAdd,
	}
	v.log.Debug("Add/del interface nat6")
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Add/del interface nat6 failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Add/del interface nat6 failed with retval: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) CalicoAddInterfaceNat6(swIfIndex uint32) (err error) {
	return v.calicoAddDelIfNat6(swIfIndex, true /* isAdd */)
}

func (v *VppLink) CalicoDelInterfaceNat6(swIfIndex uint32) (err error) {
	return v.calicoAddDelIfNat6(swIfIndex, false /* isAdd */)
}

// --------------

func (v *VppLink) calicoAddDelAs(asAddress net.IP, vipAddress *net.IPNet, port int32, isAdd bool) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &calico.CalicoAddDelAsReply{}
	request := &calico.CalicoAddDelAs{
		Pfx:       types.ToCalicoVppIpAddressWithPrefix(vipAddress),
		Protocol:  uint8(calico.IP_API_PROTO_TCP),
		Port:      uint16(port),
		AsAddress: types.ToCalicoVppIpAddress(asAddress),
		IsDel:     !isAdd,
		IsFlush:   false,
	}
	v.log.Debug("Add/del As")
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Add/del As failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Add/del As failed with retval: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) CalicoAddAs(asAddress net.IP, vipAddress *net.IPNet, port int32) (err error) {
	return v.calicoAddDelAs(asAddress, vipAddress, port, true /* isAdd */)
}

func (v *VppLink) CalicoDelAs(asAddress net.IP, vipAddress *net.IPNet, port int32) (err error) {
	return v.calicoAddDelAs(asAddress, vipAddress, port, false /* isAdd */)
}

// --------------

func (v *VppLink) calicoAddDelVip(vipAddress *net.IPNet, port int32, targetPort int32, encapIsv6 bool, isAdd bool) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	encap := calico.CALICO_API_ENCAP_TYPE_NAT4
	if encapIsv6 {
		encap = calico.CALICO_API_ENCAP_TYPE_NAT6
	}

	response := &calico.CalicoAddDelVipReply{}
	request := &calico.CalicoAddDelVip{
		Pfx:                 types.ToCalicoVppIpAddressWithPrefix(vipAddress),
		Protocol:            uint8(calico.IP_API_PROTO_TCP),
		Port:                uint16(port),
		Encap:               encap,
		TargetPort:          uint16(targetPort),
		NewFlowsTableLength: 1024,
		IsDel:               !isAdd,
	}
	v.log.Debug("Adding Vip")
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Add/del vip failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Add/del vip failed with retval: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) CalicoAddVip(vipAddress *net.IPNet, port int32, targetPort int32, encapIsv6 bool) (err error) {
	return v.calicoAddDelVip(vipAddress, port, targetPort, encapIsv6, true /* isAdd */)
}

func (v *VppLink) CalicoDelVip(vipAddress *net.IPNet, port int32, targetPort int32, encapIsv6 bool) (err error) {
	return v.calicoAddDelVip(vipAddress, port, targetPort, encapIsv6, false /* isAdd */)
}

// --------------

func (v *VppLink) calicoAddDelSnatEntry(addr net.IP, prefix *net.IPNet, tableID uint32, isAdd bool) (err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	response := &calico.CalicoAddDelSnatEntryReply{}
	request := &calico.CalicoAddDelSnatEntry{
		Pfx:     types.ToCalicoVppIpAddressWithPrefix(prefix),
		Addr:    types.ToCalicoVppIpAddress(addr),
		TableID: tableID,
		IsAdd:   isAdd,
	}
	v.log.Debug("Add/Del SnatEntry")
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "Add/del SnatEntry failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("Add/del SnatEntry failed with retval: %d", response.Retval)
	}
	return nil
}

func (v *VppLink) CalicoAddSnatEntry(addr net.IP, prefix *net.IPNet, tableID uint32) (err error) {
	return v.calicoAddDelSnatEntry(addr, prefix, tableID, true /* isAdd */)
}

func (v *VppLink) CalicoDelSnatEntry(addr net.IP, prefix *net.IPNet, tableID uint32) (err error) {
	return v.calicoAddDelSnatEntry(addr, prefix, tableID, false /* isAdd */)
}
