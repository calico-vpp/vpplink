// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /Users/aloaugus/src/vpp-api-repo/20.05-rc0~540_g77ea42b~b9261/api/plugins/tls_openssl.api.json

/*
Package tls_openssl is a generated VPP binary API for 'tls_openssl' module.

It consists of:
	  2 messages
	  1 service
*/
package tls_openssl

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "tls_openssl"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x7386fbcd
)

// TLSOpensslSetEngine represents VPP binary API message 'tls_openssl_set_engine'.
type TLSOpensslSetEngine struct {
	AsyncEnable uint32
	Engine      []byte `struc:"[64]byte"`
	Algorithm   []byte `struc:"[64]byte"`
	Ciphers     []byte `struc:"[64]byte"`
}

func (m *TLSOpensslSetEngine) Reset()                        { *m = TLSOpensslSetEngine{} }
func (*TLSOpensslSetEngine) GetMessageName() string          { return "tls_openssl_set_engine" }
func (*TLSOpensslSetEngine) GetCrcString() string            { return "e34d95c1" }
func (*TLSOpensslSetEngine) GetMessageType() api.MessageType { return api.RequestMessage }

// TLSOpensslSetEngineReply represents VPP binary API message 'tls_openssl_set_engine_reply'.
type TLSOpensslSetEngineReply struct {
	Retval int32
}

func (m *TLSOpensslSetEngineReply) Reset()                        { *m = TLSOpensslSetEngineReply{} }
func (*TLSOpensslSetEngineReply) GetMessageName() string          { return "tls_openssl_set_engine_reply" }
func (*TLSOpensslSetEngineReply) GetCrcString() string            { return "e8d4e804" }
func (*TLSOpensslSetEngineReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*TLSOpensslSetEngine)(nil), "tls_openssl.TLSOpensslSetEngine")
	api.RegisterMessage((*TLSOpensslSetEngineReply)(nil), "tls_openssl.TLSOpensslSetEngineReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TLSOpensslSetEngine)(nil),
		(*TLSOpensslSetEngineReply)(nil),
	}
}

// RPCService represents RPC service API for tls_openssl module.
type RPCService interface {
	TLSOpensslSetEngine(ctx context.Context, in *TLSOpensslSetEngine) (*TLSOpensslSetEngineReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) TLSOpensslSetEngine(ctx context.Context, in *TLSOpensslSetEngine) (*TLSOpensslSetEngineReply, error) {
	out := new(TLSOpensslSetEngineReply)
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