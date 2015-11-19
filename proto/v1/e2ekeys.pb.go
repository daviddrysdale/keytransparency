// Code generated by protoc-gen-go.
// source: proto/v1/e2ekeys.proto
// DO NOT EDIT!

/*
Package google_security_e2ekeys_v1 is a generated protocol buffer package.

It is generated from these files:
	proto/v1/e2ekeys.proto

It has these top-level messages:
	HkpLookupRequest
	HttpResponse
*/
package google_security_e2ekeys_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_security_e2ekeys_v2 "github.com/google/e2e-key-server/proto/v2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HkpLookupRequest contains query parameters for retrieving PGP keys.
type HkpLookupRequest struct {
	// Op specifies the operation to be performed on the keyserver.
	// - "get" returns the pgp key specified in the search parameter.
	// - "index" returns 501 (not implemented).
	// - "vindex" returns 501 (not implemented).
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Search specifies the email address or key id being queried.
	Search string `protobuf:"bytes,2,opt,name=search" json:"search,omitempty"`
	// Options specifies what output format to use.
	// - "mr" machine readable will set the content type to "application/pgp-keys"
	// - other options will be ignored.
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	// Exact specifies an exact match on search. Always on. If specified in the
	// URL, its value will be ignored.
	Exact string `protobuf:"bytes,4,opt,name=exact" json:"exact,omitempty"`
	// fingerprint is ignored.
	Fingerprint string `protobuf:"bytes,5,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *HkpLookupRequest) Reset()                    { *m = HkpLookupRequest{} }
func (m *HkpLookupRequest) String() string            { return proto.CompactTextString(m) }
func (*HkpLookupRequest) ProtoMessage()               {}
func (*HkpLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// HttpBody represents an http body.
type HttpResponse struct {
	// Header content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type" json:"content_type,omitempty"`
	// The http body itself.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HttpResponse) Reset()                    { *m = HttpResponse{} }
func (m *HttpResponse) String() string            { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()               {}
func (*HttpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HkpLookupRequest)(nil), "google.security.e2ekeys.v1.HkpLookupRequest")
	proto.RegisterType((*HttpResponse)(nil), "google.security.e2ekeys.v1.HttpResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for E2EKeyProxy service

type E2EKeyProxyClient interface {
	// GetEntry returns a user's current profile.
	GetEntry(ctx context.Context, in *google_security_e2ekeys_v2.GetEntryRequest, opts ...grpc.CallOption) (*google_security_e2ekeys_v2.Profile, error)
	HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error)
}

type e2EKeyProxyClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyProxyClient(cc *grpc.ClientConn) E2EKeyProxyClient {
	return &e2EKeyProxyClient{cc}
}

func (c *e2EKeyProxyClient) GetEntry(ctx context.Context, in *google_security_e2ekeys_v2.GetEntryRequest, opts ...grpc.CallOption) (*google_security_e2ekeys_v2.Profile, error) {
	out := new(google_security_e2ekeys_v2.Profile)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyProxyClient) HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	out := new(HttpResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/HkpLookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyProxy service

type E2EKeyProxyServer interface {
	// GetEntry returns a user's current profile.
	GetEntry(context.Context, *google_security_e2ekeys_v2.GetEntryRequest) (*google_security_e2ekeys_v2.Profile, error)
	HkpLookup(context.Context, *HkpLookupRequest) (*HttpResponse, error)
}

func RegisterE2EKeyProxyServer(s *grpc.Server, srv E2EKeyProxyServer) {
	s.RegisterService(&_E2EKeyProxy_serviceDesc, srv)
}

func _E2EKeyProxy_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_security_e2ekeys_v2.GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).GetEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyProxy_HkpLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HkpLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).HkpLookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _E2EKeyProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.security.e2ekeys.v1.E2EKeyProxy",
	HandlerType: (*E2EKeyProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _E2EKeyProxy_GetEntry_Handler,
		},
		{
			MethodName: "HkpLookup",
			Handler:    _E2EKeyProxy_HkpLookup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4e, 0x32, 0x41,
	0x10, 0xc4, 0x03, 0x1f, 0xf0, 0x49, 0xb3, 0xfe, 0xc9, 0x68, 0xcc, 0x66, 0x4f, 0x06, 0x2f, 0x24,
	0xca, 0x6c, 0x58, 0x0f, 0x3e, 0x01, 0x91, 0x44, 0x0f, 0xc6, 0xb3, 0x89, 0x81, 0xb1, 0x59, 0x36,
	0xe0, 0xf4, 0x38, 0xd3, 0x4b, 0x98, 0xb7, 0xf4, 0x91, 0x5c, 0x46, 0x30, 0x6a, 0x02, 0xc7, 0xae,
	0xae, 0xaa, 0x4e, 0xff, 0xe0, 0xdc, 0x58, 0x62, 0x4a, 0x97, 0x83, 0x14, 0x33, 0x9c, 0xa3, 0x77,
	0x32, 0x08, 0x22, 0xc9, 0x89, 0xf2, 0x05, 0x4a, 0x87, 0xaa, 0xb4, 0x05, 0x7b, 0xb9, 0x5d, 0x2f,
	0x07, 0xc9, 0x6d, 0x5e, 0xf0, 0xac, 0x9c, 0x48, 0x45, 0x6f, 0xe9, 0x97, 0x6d, 0x1d, 0xee, 0x57,
	0xeb, 0xbe, 0x43, 0xbb, 0x44, 0x9b, 0x6e, 0x4a, 0xb3, 0xdf, 0xa5, 0x5d, 0x05, 0x27, 0xa3, 0xb9,
	0x79, 0x20, 0x9a, 0x97, 0xe6, 0x09, 0xdf, 0x4b, 0x74, 0x2c, 0x00, 0xea, 0x64, 0xe2, 0xda, 0x45,
	0xad, 0xd7, 0x16, 0x47, 0xd0, 0x72, 0x38, 0xb6, 0x6a, 0x16, 0xd7, 0xc3, 0x7c, 0x0c, 0xff, 0xc9,
	0x70, 0x41, 0xda, 0xc5, 0xff, 0x82, 0x70, 0x08, 0x4d, 0x5c, 0x8d, 0x15, 0xc7, 0x8d, 0x30, 0x9e,
	0x42, 0x67, 0x5a, 0xe8, 0x1c, 0xad, 0xb1, 0x85, 0xe6, 0xb8, 0xb9, 0x16, 0xbb, 0x19, 0x44, 0x23,
	0xe6, 0xaa, 0xdf, 0x99, 0x2a, 0x88, 0xe2, 0x0c, 0x22, 0x45, 0x9a, 0x51, 0xf3, 0x0b, 0x7b, 0x83,
	0x9b, 0x53, 0x11, 0x34, 0x26, 0xf4, 0xea, 0xc3, 0xa1, 0x28, 0xfb, 0xa8, 0x41, 0x67, 0x98, 0x0d,
	0xef, 0xd1, 0x3f, 0x5a, 0x5a, 0x79, 0xf1, 0x0c, 0x07, 0x77, 0xc8, 0x43, 0xcd, 0xd6, 0x8b, 0x2b,
	0xb9, 0x13, 0x45, 0x26, 0xb7, 0xae, 0xcd, 0x37, 0xc9, 0xe5, 0x3e, 0x73, 0xd5, 0x3d, 0x2d, 0x16,
	0x28, 0x14, 0xb4, 0xbf, 0x31, 0x88, 0xeb, 0xdd, 0x89, 0x81, 0xfc, 0x4b, 0x2b, 0xe9, 0xed, 0x75,
	0xff, 0x78, 0x7b, 0xd2, 0x0a, 0xc8, 0x6f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x95, 0xe2,
	0xc3, 0xe1, 0x01, 0x00, 0x00,
}
