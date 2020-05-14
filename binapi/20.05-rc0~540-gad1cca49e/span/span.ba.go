// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /Users/nskrzypc/wrk/vpp/build-root/install-vpp-native/vpp/share/vpp/api/core/span.api.json

/*
Package span is a generated VPP binary API for 'span' module.

It consists of:
	  7 enums
	  1 alias
	  4 messages
	  2 services
*/
package span

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "span"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x9f4dec7c
)

// IfStatusFlags represents VPP binary API enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var IfStatusFlags_name = map[uint32]string{
	1: "IF_STATUS_API_FLAG_ADMIN_UP",
	2: "IF_STATUS_API_FLAG_LINK_UP",
}

var IfStatusFlags_value = map[string]uint32{
	"IF_STATUS_API_FLAG_ADMIN_UP": 1,
	"IF_STATUS_API_FLAG_LINK_UP":  2,
}

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfType represents VPP binary API enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 0
	IF_API_TYPE_SUB      IfType = 1
	IF_API_TYPE_P2P      IfType = 2
	IF_API_TYPE_PIPE     IfType = 3
)

var IfType_name = map[uint32]string{
	0: "IF_API_TYPE_HARDWARE",
	1: "IF_API_TYPE_SUB",
	2: "IF_API_TYPE_P2P",
	3: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 0,
	"IF_API_TYPE_SUB":      1,
	"IF_API_TYPE_P2P":      2,
	"IF_API_TYPE_PIPE":     3,
}

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LinkDuplex represents VPP binary API enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var LinkDuplex_name = map[uint32]string{
	0: "LINK_DUPLEX_API_UNKNOWN",
	1: "LINK_DUPLEX_API_HALF",
	2: "LINK_DUPLEX_API_FULL",
}

var LinkDuplex_value = map[string]uint32{
	"LINK_DUPLEX_API_UNKNOWN": 0,
	"LINK_DUPLEX_API_HALF":    1,
	"LINK_DUPLEX_API_FULL":    2,
}

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 0
	MTU_PROTO_API_IP4  MtuProto = 1
	MTU_PROTO_API_IP6  MtuProto = 2
	MTU_PROTO_API_MPLS MtuProto = 3
)

var MtuProto_name = map[uint32]string{
	0: "MTU_PROTO_API_L3",
	1: "MTU_PROTO_API_IP4",
	2: "MTU_PROTO_API_IP6",
	3: "MTU_PROTO_API_MPLS",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   0,
	"MTU_PROTO_API_IP4":  1,
	"MTU_PROTO_API_IP6":  2,
	"MTU_PROTO_API_MPLS": 3,
}

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RxMode represents VPP binary API enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var RxMode_name = map[uint32]string{
	0: "RX_MODE_API_UNKNOWN",
	1: "RX_MODE_API_POLLING",
	2: "RX_MODE_API_INTERRUPT",
	3: "RX_MODE_API_ADAPTIVE",
	4: "RX_MODE_API_DEFAULT",
}

var RxMode_value = map[string]uint32{
	"RX_MODE_API_UNKNOWN":   0,
	"RX_MODE_API_POLLING":   1,
	"RX_MODE_API_INTERRUPT": 2,
	"RX_MODE_API_ADAPTIVE":  3,
	"RX_MODE_API_DEFAULT":   4,
}

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SpanState represents VPP binary API enum 'span_state'.
type SpanState uint32

const (
	SPAN_STATE_API_DISABLED SpanState = 0
	SPAN_STATE_API_RX       SpanState = 1
	SPAN_STATE_API_TX       SpanState = 2
	SPAN_STATE_API_RX_TX    SpanState = 3
)

var SpanState_name = map[uint32]string{
	0: "SPAN_STATE_API_DISABLED",
	1: "SPAN_STATE_API_RX",
	2: "SPAN_STATE_API_TX",
	3: "SPAN_STATE_API_RX_TX",
}

var SpanState_value = map[string]uint32{
	"SPAN_STATE_API_DISABLED": 0,
	"SPAN_STATE_API_RX":       1,
	"SPAN_STATE_API_TX":       2,
	"SPAN_STATE_API_RX_TX":    3,
}

func (x SpanState) String() string {
	s, ok := SpanState_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SubIfFlags represents VPP binary API enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var SubIfFlags_name = map[uint32]string{
	1:   "SUB_IF_API_FLAG_NO_TAGS",
	2:   "SUB_IF_API_FLAG_ONE_TAG",
	4:   "SUB_IF_API_FLAG_TWO_TAGS",
	8:   "SUB_IF_API_FLAG_DOT1AD",
	16:  "SUB_IF_API_FLAG_EXACT_MATCH",
	32:  "SUB_IF_API_FLAG_DEFAULT",
	64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
	128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
	254: "SUB_IF_API_FLAG_MASK_VNET",
	256: "SUB_IF_API_FLAG_DOT1AH",
}

var SubIfFlags_value = map[string]uint32{
	"SUB_IF_API_FLAG_NO_TAGS":           1,
	"SUB_IF_API_FLAG_ONE_TAG":           2,
	"SUB_IF_API_FLAG_TWO_TAGS":          4,
	"SUB_IF_API_FLAG_DOT1AD":            8,
	"SUB_IF_API_FLAG_EXACT_MATCH":       16,
	"SUB_IF_API_FLAG_DEFAULT":           32,
	"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
	"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
	"SUB_IF_API_FLAG_MASK_VNET":         254,
	"SUB_IF_API_FLAG_DOT1AH":            256,
}

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// SwInterfaceSpanDetails represents VPP binary API message 'sw_interface_span_details'.
type SwInterfaceSpanDetails struct {
	SwIfIndexFrom InterfaceIndex
	SwIfIndexTo   InterfaceIndex
	State         SpanState
	IsL2          bool
}

func (m *SwInterfaceSpanDetails) Reset()                        { *m = SwInterfaceSpanDetails{} }
func (*SwInterfaceSpanDetails) GetMessageName() string          { return "sw_interface_span_details" }
func (*SwInterfaceSpanDetails) GetCrcString() string            { return "055643fc" }
func (*SwInterfaceSpanDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceSpanDump represents VPP binary API message 'sw_interface_span_dump'.
type SwInterfaceSpanDump struct {
	IsL2 bool
}

func (m *SwInterfaceSpanDump) Reset()                        { *m = SwInterfaceSpanDump{} }
func (*SwInterfaceSpanDump) GetMessageName() string          { return "sw_interface_span_dump" }
func (*SwInterfaceSpanDump) GetCrcString() string            { return "d6cf0c3d" }
func (*SwInterfaceSpanDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSpanEnableDisable represents VPP binary API message 'sw_interface_span_enable_disable'.
type SwInterfaceSpanEnableDisable struct {
	SwIfIndexFrom InterfaceIndex
	SwIfIndexTo   InterfaceIndex
	State         SpanState
	IsL2          bool
}

func (m *SwInterfaceSpanEnableDisable) Reset() { *m = SwInterfaceSpanEnableDisable{} }
func (*SwInterfaceSpanEnableDisable) GetMessageName() string {
	return "sw_interface_span_enable_disable"
}
func (*SwInterfaceSpanEnableDisable) GetCrcString() string            { return "acc8fea1" }
func (*SwInterfaceSpanEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSpanEnableDisableReply represents VPP binary API message 'sw_interface_span_enable_disable_reply'.
type SwInterfaceSpanEnableDisableReply struct {
	Retval int32
}

func (m *SwInterfaceSpanEnableDisableReply) Reset() { *m = SwInterfaceSpanEnableDisableReply{} }
func (*SwInterfaceSpanEnableDisableReply) GetMessageName() string {
	return "sw_interface_span_enable_disable_reply"
}
func (*SwInterfaceSpanEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSpanEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*SwInterfaceSpanDetails)(nil), "span.SwInterfaceSpanDetails")
	api.RegisterMessage((*SwInterfaceSpanDump)(nil), "span.SwInterfaceSpanDump")
	api.RegisterMessage((*SwInterfaceSpanEnableDisable)(nil), "span.SwInterfaceSpanEnableDisable")
	api.RegisterMessage((*SwInterfaceSpanEnableDisableReply)(nil), "span.SwInterfaceSpanEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceSpanDetails)(nil),
		(*SwInterfaceSpanDump)(nil),
		(*SwInterfaceSpanEnableDisable)(nil),
		(*SwInterfaceSpanEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for span module.
type RPCService interface {
	DumpSwInterfaceSpan(ctx context.Context, in *SwInterfaceSpanDump) (RPCService_DumpSwInterfaceSpanClient, error)
	SwInterfaceSpanEnableDisable(ctx context.Context, in *SwInterfaceSpanEnableDisable) (*SwInterfaceSpanEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwInterfaceSpan(ctx context.Context, in *SwInterfaceSpanDump) (RPCService_DumpSwInterfaceSpanClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceSpanClient{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceSpanClient interface {
	Recv() (*SwInterfaceSpanDetails, error)
}

type serviceClient_DumpSwInterfaceSpanClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceSpanClient) Recv() (*SwInterfaceSpanDetails, error) {
	m := new(SwInterfaceSpanDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) SwInterfaceSpanEnableDisable(ctx context.Context, in *SwInterfaceSpanEnableDisable) (*SwInterfaceSpanEnableDisableReply, error) {
	out := new(SwInterfaceSpanEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
