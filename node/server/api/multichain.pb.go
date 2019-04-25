// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multichain.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddressMessage struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMessage) Reset()         { *m = AddressMessage{} }
func (m *AddressMessage) String() string { return proto.CompactTextString(m) }
func (*AddressMessage) ProtoMessage()    {}
func (*AddressMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{0}
}

func (m *AddressMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMessage.Unmarshal(m, b)
}
func (m *AddressMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMessage.Marshal(b, m, deterministic)
}
func (m *AddressMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMessage.Merge(m, src)
}
func (m *AddressMessage) XXX_Size() int {
	return xxx_messageInfo_AddressMessage.Size(m)
}
func (m *AddressMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMessage proto.InternalMessageInfo

func (m *AddressMessage) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AddressesRequest struct {
	Addresses            []*AddressMessage `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddressesRequest) Reset()         { *m = AddressesRequest{} }
func (m *AddressesRequest) String() string { return proto.CompactTextString(m) }
func (*AddressesRequest) ProtoMessage()    {}
func (*AddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{1}
}

func (m *AddressesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressesRequest.Unmarshal(m, b)
}
func (m *AddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressesRequest.Marshal(b, m, deterministic)
}
func (m *AddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressesRequest.Merge(m, src)
}
func (m *AddressesRequest) XXX_Size() int {
	return xxx_messageInfo_AddressesRequest.Size(m)
}
func (m *AddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressesRequest proto.InternalMessageInfo

func (m *AddressesRequest) GetAddresses() []*AddressMessage {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type BalancesResponse struct {
	Balances             map[string]float64 `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BalancesResponse) Reset()         { *m = BalancesResponse{} }
func (m *BalancesResponse) String() string { return proto.CompactTextString(m) }
func (*BalancesResponse) ProtoMessage()    {}
func (*BalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd326252c7e9300e, []int{2}
}

func (m *BalancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancesResponse.Unmarshal(m, b)
}
func (m *BalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancesResponse.Marshal(b, m, deterministic)
}
func (m *BalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancesResponse.Merge(m, src)
}
func (m *BalancesResponse) XXX_Size() int {
	return xxx_messageInfo_BalancesResponse.Size(m)
}
func (m *BalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalancesResponse proto.InternalMessageInfo

func (m *BalancesResponse) GetBalances() map[string]float64 {
	if m != nil {
		return m.Balances
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressMessage)(nil), "api.AddressMessage")
	proto.RegisterType((*AddressesRequest)(nil), "api.AddressesRequest")
	proto.RegisterType((*BalancesResponse)(nil), "api.BalancesResponse")
	proto.RegisterMapType((map[string]float64)(nil), "api.BalancesResponse.BalancesEntry")
}

func init() { proto.RegisterFile("multichain.proto", fileDescriptor_dd326252c7e9300e) }

var fileDescriptor_dd326252c7e9300e = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x2d, 0xcd, 0x29,
	0xc9, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c,
	0xc8, 0x54, 0xd2, 0xe2, 0xe2, 0x73, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xf6, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe0, 0x62, 0x4f, 0x84, 0x88, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xb8, 0x4a, 0xae, 0x5c, 0x02, 0x50, 0xb5, 0xa9, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x86, 0x5c, 0x9c, 0x89, 0x30, 0x31, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23,
	0x61, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0x54, 0x53, 0x83, 0x10, 0xaa, 0x94, 0x26, 0x30, 0x72, 0x09,
	0x38, 0x25, 0xe6, 0x24, 0xe6, 0x25, 0x83, 0x8c, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2,
	0xe7, 0xe2, 0x48, 0x82, 0x8a, 0x41, 0x8d, 0x51, 0x06, 0x1b, 0x83, 0xae, 0x10, 0x2e, 0xe0, 0x9a,
	0x57, 0x52, 0x54, 0x19, 0x04, 0xd7, 0x24, 0x65, 0xcd, 0xc5, 0x8b, 0x22, 0x25, 0x24, 0xc0, 0xc5,
	0x9c, 0x9d, 0x5a, 0x09, 0xf5, 0x03, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a,
	0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x18, 0x04, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x1a, 0x79, 0x72,
	0xb1, 0x43, 0x35, 0x0b, 0xd9, 0x71, 0xf1, 0xb8, 0xa5, 0x96, 0x24, 0x67, 0xc0, 0xf8, 0xa2, 0xc8,
	0xbe, 0x81, 0xfb, 0x5b, 0x4a, 0x14, 0xab, 0xeb, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x81, 0x6b, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x02, 0x47, 0xdb, 0xb7, 0x70, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceClient interface {
	FetchBalance(ctx context.Context, in *AddressesRequest, opts ...grpc.CallOption) (*BalancesResponse, error)
}

type balanceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceClient(cc *grpc.ClientConn) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) FetchBalance(ctx context.Context, in *AddressesRequest, opts ...grpc.CallOption) (*BalancesResponse, error) {
	out := new(BalancesResponse)
	err := c.cc.Invoke(ctx, "/api.Balance/FetchBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
type BalanceServer interface {
	FetchBalance(context.Context, *AddressesRequest) (*BalancesResponse, error)
}

// UnimplementedBalanceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServer struct {
}

func (*UnimplementedBalanceServer) FetchBalance(ctx context.Context, req *AddressesRequest) (*BalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBalance not implemented")
}

func RegisterBalanceServer(s *grpc.Server, srv BalanceServer) {
	s.RegisterService(&_Balance_serviceDesc, srv)
}

func _Balance_FetchBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).FetchBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balance/FetchBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).FetchBalance(ctx, req.(*AddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Balance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchBalance",
			Handler:    _Balance_FetchBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multichain.proto",
}