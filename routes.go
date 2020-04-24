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

	vppip "github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/ip"
	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/ip_neighbor"
	"github.com/calico-vpp/vpplink/types"
	"github.com/pkg/errors"
)

const (
	AnyInterface = ^uint32(0)
)

func (v *VppLink) GetRoutes(tableID uint32, isIPv6 bool) (routes []types.Route, err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.IPRouteDump{
		Table: vppip.IPTable{
			TableID: tableID,
			IsIP6:   isIPv6,
		},
	}
	response := &vppip.IPRouteDetails{}
	v.log.Debug("Listing VPP routes")
	stream := v.ch.SendMultiRequest(request)
	for {
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			return nil, errors.Wrap(err, "error listing VPP routes")
		}
		if stop {
			return routes, nil
		}
		vppRoute := response.Route
		if vppRoute.NPaths != 1 {
			// Multipath not supported yet
			v.log.Info("Skipping route with %d paths : %+v", vppRoute.NPaths, vppRoute)
			continue
		}
		vppPath := vppRoute.Paths[0]
		route := types.Route{
			Dst:       types.FromVppIpPrefix(vppRoute.Prefix),
			Table:     int(vppRoute.TableID),
			Gw:        types.FromVppIpAddressUnion(vppPath.Nh.Address, vppRoute.Prefix.Address.Af == vppip.ADDRESS_IP6),
			DstTable:  int(vppPath.TableID),
			SwIfIndex: vppPath.SwIfIndex,
		}
		routes = append(routes, route)
	}
}

func (v *VppLink) AddNeighbor(neighbor *types.Neighbor) error {
	return v.addDelNeighbor(neighbor, true)
}

func (v *VppLink) DelNeighbor(neighbor *types.Neighbor) error {
	return v.addDelNeighbor(neighbor, false)
}

func (v *VppLink) addDelNeighbor(neighbor *types.Neighbor, isAdd bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &ip_neighbor.IPNeighborAddDel{
		IsAdd: isAdd,
		Neighbor: ip_neighbor.IPNeighbor{
			SwIfIndex:  ip_neighbor.InterfaceIndex(neighbor.SwIfIndex),
			Flags:      neighbor.GetVppNeighborFlags(),
			MacAddress: ip_neighbor.MacAddress(neighbor.GetVppMacAddress()),
			IPAddress:  neighbor.GetVppNeighborAddress(),
		},
	}
	response := &ip_neighbor.IPNeighborAddDelReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "failed to add/delete (%d) neighbor from VPP", isAdd)
	} else if response.Retval != 0 {
		return fmt.Errorf("failed to add/delete (%d) neighbor from VPP (retval %d)", isAdd, response.Retval)
	}
	v.log.Debugf("added/deleted (%d) neighbor %+v", isAdd, neighbor)
	return nil
}

func (v *VppLink) RouteAdd(route *types.Route) error {
	return v.addDelIPRoute(route, true)
}

func (v *VppLink) RouteDel(route *types.Route) error {
	return v.addDelIPRoute(route, false)
}

func (v *VppLink) addDelIPRoute(route *types.Route, isAdd bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	prefixLen, _ := route.Dst.Mask.Size()

	proto := vppip.FIB_API_PATH_NH_PROTO_IP4
	if IsIP6(route.Dst.IP) {
		proto = vppip.FIB_API_PATH_NH_PROTO_IP6
	}

	paths := []vppip.FibPath{
		{
			SwIfIndex:  uint32(route.SwIfIndex),
			TableID:    uint32(route.DstTable),
			RpfID:      0,
			Weight:     1,
			Preference: 0,
			Type:       vppip.FIB_API_PATH_TYPE_NORMAL,
			Flags:      vppip.FIB_API_PATH_FLAG_NONE,
			Proto:      proto,
		},
	}
	if route.Gw != nil {
		paths[0].Nh.Address = route.GetVppGwAddress().Un
	}

	vppRoute := vppip.IPRoute{
		TableID: uint32(route.DstTable),
		Prefix: vppip.Prefix{
			Len:     uint8(prefixLen),
			Address: route.GetVppDstAddress(),
		},
		Paths: paths,
	}

	request := &vppip.IPRouteAddDel{
		IsAdd: isAdd,
		Route: vppRoute,
	}
	response := &vppip.IPRouteAddDelReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "failed to add/delete (%d) route from VPP", isAdd)
	} else if response.Retval != 0 {
		return fmt.Errorf("failed to add/delete (%d) route from VPP (retval %d)", isAdd, response.Retval)
	}
	v.log.Debugf("added/deleted (%d) route %+v", isAdd, route)
	return nil
}
