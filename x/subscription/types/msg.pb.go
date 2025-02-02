// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCancelRequest defines the SDK message for cancelling a subscription
type MsgCancelRequest struct {
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID   uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCancelRequest) Reset()         { *m = MsgCancelRequest{} }
func (m *MsgCancelRequest) String() string { return proto.CompactTextString(m) }
func (*MsgCancelRequest) ProtoMessage()    {}
func (*MsgCancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb96050c6471ce0, []int{0}
}
func (m *MsgCancelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelRequest.Merge(m, src)
}
func (m *MsgCancelRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelRequest proto.InternalMessageInfo

// MsgAllocateRequest defines the SDK message for allocating the bytes of a
// subscription for an address
type MsgAllocateRequest struct {
	From    string                                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID      uint64                                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Address string                                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Bytes   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=bytes,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bytes"`
}

func (m *MsgAllocateRequest) Reset()         { *m = MsgAllocateRequest{} }
func (m *MsgAllocateRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAllocateRequest) ProtoMessage()    {}
func (*MsgAllocateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb96050c6471ce0, []int{1}
}
func (m *MsgAllocateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAllocateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAllocateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAllocateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAllocateRequest.Merge(m, src)
}
func (m *MsgAllocateRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAllocateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAllocateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAllocateRequest proto.InternalMessageInfo

// MsgCancelResponse defines the response of message MsgCancelRequest
type MsgCancelResponse struct {
}

func (m *MsgCancelResponse) Reset()         { *m = MsgCancelResponse{} }
func (m *MsgCancelResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelResponse) ProtoMessage()    {}
func (*MsgCancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb96050c6471ce0, []int{2}
}
func (m *MsgCancelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelResponse.Merge(m, src)
}
func (m *MsgCancelResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelResponse proto.InternalMessageInfo

// MsgAllocateResponse defines the response of message MsgAllocateRequest
type MsgAllocateResponse struct {
}

func (m *MsgAllocateResponse) Reset()         { *m = MsgAllocateResponse{} }
func (m *MsgAllocateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAllocateResponse) ProtoMessage()    {}
func (*MsgAllocateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb96050c6471ce0, []int{3}
}
func (m *MsgAllocateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAllocateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAllocateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAllocateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAllocateResponse.Merge(m, src)
}
func (m *MsgAllocateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAllocateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAllocateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAllocateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCancelRequest)(nil), "sentinel.subscription.v2.MsgCancelRequest")
	proto.RegisterType((*MsgAllocateRequest)(nil), "sentinel.subscription.v2.MsgAllocateRequest")
	proto.RegisterType((*MsgCancelResponse)(nil), "sentinel.subscription.v2.MsgCancelResponse")
	proto.RegisterType((*MsgAllocateResponse)(nil), "sentinel.subscription.v2.MsgAllocateResponse")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v2/msg.proto", fileDescriptor_4cb96050c6471ce0)
}

var fileDescriptor_4cb96050c6471ce0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0xae, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x50, 0x8a, 0x6a, 0x16, 0x70, 0x01, 0x45, 0x1d, 0xdc, 0x2a, 0x03, 0xaa, 0x80,
	0xda, 0xa2, 0x2c, 0x4c, 0x48, 0x94, 0x2e, 0x1d, 0xba, 0x84, 0xad, 0x5b, 0x3e, 0xdc, 0xd4, 0x90,
	0xc4, 0x21, 0xc7, 0x89, 0xe8, 0x5b, 0xf0, 0x08, 0x8c, 0x3c, 0x4a, 0xc7, 0x8e, 0x88, 0x21, 0x82,
	0xf4, 0x45, 0x50, 0x13, 0x8a, 0xd2, 0x4a, 0xf7, 0xaa, 0xba, 0x93, 0x8f, 0x8f, 0x7e, 0xff, 0xf3,
	0xf1, 0xd7, 0xc1, 0x36, 0x88, 0x44, 0xcb, 0x44, 0x44, 0x1c, 0x72, 0x0f, 0xfc, 0x4c, 0xa6, 0x5a,
	0xaa, 0x84, 0x17, 0x53, 0x1e, 0x43, 0xc8, 0xd2, 0x4c, 0x69, 0x45, 0xac, 0x13, 0xc3, 0xda, 0x0c,
	0x2b, 0xa6, 0x83, 0x27, 0xa1, 0x0a, 0x55, 0x0d, 0xf1, 0x63, 0xd4, 0xf0, 0xf6, 0x3b, 0xfc, 0x68,
	0x09, 0xe1, 0x07, 0x37, 0xf1, 0x45, 0xe4, 0x88, 0x2f, 0xb9, 0x00, 0x4d, 0x08, 0xee, 0xac, 0x33,
	0x15, 0x5b, 0x68, 0x84, 0xc6, 0x3d, 0xa7, 0x8e, 0xc9, 0x33, 0x6c, 0xca, 0xc0, 0x32, 0x47, 0x68,
	0xdc, 0x99, 0x75, 0xab, 0x72, 0x68, 0x2e, 0xe6, 0x8e, 0x29, 0x03, 0xfb, 0x3b, 0xc2, 0x64, 0x09,
	0xe1, 0xfb, 0x28, 0x52, 0xbe, 0xab, 0xc5, 0x1d, 0x4a, 0x10, 0x0b, 0x3f, 0x70, 0x83, 0x20, 0x13,
	0x00, 0xd6, 0xbd, 0x1a, 0x3f, 0x7d, 0xc9, 0x1c, 0xdf, 0xf7, 0xb6, 0x5a, 0x80, 0xd5, 0x39, 0xe6,
	0x67, 0x6c, 0x57, 0x0e, 0x8d, 0x5f, 0xe5, 0xf0, 0x79, 0x28, 0xf5, 0x26, 0xf7, 0x98, 0xaf, 0x62,
	0xee, 0x2b, 0x88, 0x15, 0xfc, 0x7b, 0x26, 0x10, 0x7c, 0xe6, 0x7a, 0x9b, 0x0a, 0x60, 0x8b, 0x44,
	0x3b, 0x8d, 0xd8, 0xee, 0xe3, 0xc7, 0xad, 0x15, 0x21, 0x55, 0x09, 0x08, 0xfb, 0x29, 0xee, 0x9f,
	0x8d, 0xdd, 0xa4, 0xa7, 0x25, 0xc2, 0x78, 0x09, 0xe1, 0x47, 0x91, 0x15, 0xd2, 0x17, 0x24, 0xc0,
	0xbd, 0xff, 0x52, 0xf2, 0x82, 0xdd, 0xe4, 0x2d, 0xbb, 0xb4, 0x70, 0xf0, 0xf2, 0x2a, 0xb6, 0x69,
	0x4a, 0x3e, 0xe1, 0x87, 0xad, 0x59, 0xc8, 0xab, 0x5b, 0xb5, 0x17, 0x4e, 0x0f, 0x26, 0x57, 0xd2,
	0x4d, 0xaf, 0xd9, 0x6a, 0xf7, 0x87, 0x1a, 0x3f, 0x2a, 0x6a, 0xec, 0x2a, 0x8a, 0xf6, 0x15, 0x45,
	0xbf, 0x2b, 0x8a, 0xbe, 0x1d, 0xa8, 0xb1, 0x3f, 0x50, 0xe3, 0xe7, 0x81, 0x1a, 0xab, 0xb7, 0x2d,
	0x77, 0x4f, 0xa5, 0x27, 0x6a, 0xbd, 0x96, 0xbe, 0x74, 0x23, 0xbe, 0xc9, 0x3d, 0x5e, 0xbc, 0xe6,
	0x5f, 0xcf, 0x4f, 0xb0, 0xf6, 0xdc, 0xeb, 0xd6, 0x27, 0xf5, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6d, 0xd8, 0x68, 0x02, 0xa8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgCancel(ctx context.Context, in *MsgCancelRequest, opts ...grpc.CallOption) (*MsgCancelResponse, error)
	MsgAllocate(ctx context.Context, in *MsgAllocateRequest, opts ...grpc.CallOption) (*MsgAllocateResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgCancel(ctx context.Context, in *MsgCancelRequest, opts ...grpc.CallOption) (*MsgCancelResponse, error) {
	out := new(MsgCancelResponse)
	err := c.cc.Invoke(ctx, "/sentinel.subscription.v2.MsgService/MsgCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) MsgAllocate(ctx context.Context, in *MsgAllocateRequest, opts ...grpc.CallOption) (*MsgAllocateResponse, error) {
	out := new(MsgAllocateResponse)
	err := c.cc.Invoke(ctx, "/sentinel.subscription.v2.MsgService/MsgAllocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgCancel(context.Context, *MsgCancelRequest) (*MsgCancelResponse, error)
	MsgAllocate(context.Context, *MsgAllocateRequest) (*MsgAllocateResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgCancel(ctx context.Context, req *MsgCancelRequest) (*MsgCancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgCancel not implemented")
}
func (*UnimplementedMsgServiceServer) MsgAllocate(ctx context.Context, req *MsgAllocateRequest) (*MsgAllocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgAllocate not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.subscription.v2.MsgService/MsgCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgCancel(ctx, req.(*MsgCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_MsgAllocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAllocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgAllocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.subscription.v2.MsgService/MsgAllocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgAllocate(ctx, req.(*MsgAllocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.subscription.v2.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgCancel",
			Handler:    _MsgService_MsgCancel_Handler,
		},
		{
			MethodName: "MsgAllocate",
			Handler:    _MsgService_MsgAllocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/subscription/v2/msg.proto",
}

func (m *MsgCancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAllocateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAllocateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAllocateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Bytes.Size()
		i -= size
		if _, err := m.Bytes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAllocateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAllocateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAllocateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	return n
}

func (m *MsgAllocateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = m.Bytes.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgCancelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAllocateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAllocateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAllocateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAllocateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCancelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCancelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAllocateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAllocateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAllocateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
