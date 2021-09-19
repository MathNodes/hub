// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sentinel-official/hub/types"
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

type EventRegister struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty" yaml:"provider"`
}

func (m *EventRegister) Reset()         { *m = EventRegister{} }
func (m *EventRegister) String() string { return proto.CompactTextString(m) }
func (*EventRegister) ProtoMessage()    {}
func (*EventRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe9719c295e6061, []int{0}
}
func (m *EventRegister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegister.Merge(m, src)
}
func (m *EventRegister) XXX_Size() int {
	return m.Size()
}
func (m *EventRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegister.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegister proto.InternalMessageInfo

type EventUpdate struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty" yaml:"provider"`
}

func (m *EventUpdate) Reset()         { *m = EventUpdate{} }
func (m *EventUpdate) String() string { return proto.CompactTextString(m) }
func (*EventUpdate) ProtoMessage()    {}
func (*EventUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe9719c295e6061, []int{1}
}
func (m *EventUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdate.Merge(m, src)
}
func (m *EventUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdate proto.InternalMessageInfo

type EventSetStatus struct {
	Address string       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Status  types.Status `protobuf:"varint,2,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
}

func (m *EventSetStatus) Reset()         { *m = EventSetStatus{} }
func (m *EventSetStatus) String() string { return proto.CompactTextString(m) }
func (*EventSetStatus) ProtoMessage()    {}
func (*EventSetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe9719c295e6061, []int{2}
}
func (m *EventSetStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetStatus.Merge(m, src)
}
func (m *EventSetStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventSetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRegister)(nil), "sentinel.node.v1.EventRegister")
	proto.RegisterType((*EventUpdate)(nil), "sentinel.node.v1.EventUpdate")
	proto.RegisterType((*EventSetStatus)(nil), "sentinel.node.v1.EventSetStatus")
}

func init() { proto.RegisterFile("sentinel/node/v1/events.proto", fileDescriptor_5fe9719c295e6061) }

var fileDescriptor_5fe9719c295e6061 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xe3, 0x7f, 0xe8, 0x0f, 0x46, 0x2d, 0x10, 0x18, 0x4a, 0x25, 0x5c, 0xe4, 0x89, 0x81,
	0xda, 0x2a, 0x6c, 0x8c, 0x15, 0xcc, 0x48, 0xa9, 0x58, 0xd8, 0xd2, 0xe6, 0x6b, 0x6a, 0x29, 0x8d,
	0xa3, 0xd8, 0x8d, 0xe8, 0xce, 0x01, 0x38, 0x06, 0x47, 0xe9, 0xd8, 0x91, 0xa9, 0x82, 0xe4, 0x06,
	0x39, 0x01, 0x8a, 0x9d, 0x64, 0x67, 0x60, 0xb3, 0xf4, 0xbe, 0xdf, 0xf3, 0xc8, 0x9f, 0x8d, 0x2f,
	0x15, 0xc4, 0x5a, 0xc4, 0x10, 0xf1, 0x58, 0x06, 0xc0, 0xb3, 0x31, 0x87, 0x0c, 0x62, 0xad, 0x58,
	0x92, 0x4a, 0x2d, 0xdd, 0x93, 0x26, 0x66, 0x55, 0xcc, 0xb2, 0xf1, 0xe0, 0x3c, 0x94, 0xa1, 0x34,
	0x21, 0xaf, 0x4e, 0xb6, 0x37, 0x20, 0x2d, 0x46, 0x6f, 0x12, 0x50, 0x15, 0x47, 0x69, 0x5f, 0xaf,
	0x6b, 0x0e, 0x8d, 0x71, 0xf7, 0xb1, 0xe2, 0x7a, 0x10, 0x0a, 0xa5, 0x21, 0x75, 0x6f, 0xf0, 0x7f,
	0x3f, 0x08, 0x52, 0x50, 0xaa, 0x8f, 0xae, 0xd0, 0xf5, 0xe1, 0xc4, 0x2d, 0xf7, 0xc3, 0xde, 0xc6,
	0x5f, 0x45, 0xf7, 0xb4, 0x0e, 0xa8, 0xd7, 0x54, 0x5c, 0x8e, 0x0f, 0x92, 0x54, 0x66, 0x22, 0x80,
	0xb4, 0xff, 0xcf, 0xd4, 0xcf, 0xca, 0xfd, 0xf0, 0xd8, 0xd6, 0x9b, 0x84, 0x7a, 0x6d, 0x89, 0x46,
	0xf8, 0xc8, 0xf8, 0x9e, 0x93, 0xc0, 0xd7, 0xf0, 0xd7, 0xb6, 0x37, 0x84, 0x7b, 0x46, 0x37, 0x05,
	0x3d, 0x35, 0xd7, 0xfe, 0xa5, 0xf1, 0x01, 0x77, 0xec, 0xba, 0x8c, 0xaf, 0x77, 0x7b, 0xc1, 0xda,
	0xbd, 0x9b, 0x7d, 0xb2, 0x6c, 0xcc, 0x2c, 0x78, 0x72, 0x5a, 0xee, 0x87, 0x5d, 0xcb, 0xb1, 0x23,
	0xd4, 0xab, 0x67, 0x27, 0x4f, 0xdb, 0x6f, 0xe2, 0x7c, 0xe4, 0xc4, 0xd9, 0xe6, 0x04, 0xed, 0x72,
	0x82, 0xbe, 0x72, 0x82, 0xde, 0x0b, 0xe2, 0xec, 0x0a, 0xe2, 0x7c, 0x16, 0xc4, 0x79, 0x19, 0x85,
	0x42, 0x2f, 0xd7, 0x33, 0x36, 0x97, 0x2b, 0xde, 0x18, 0x46, 0x72, 0xb1, 0x10, 0x73, 0xe1, 0x47,
	0x7c, 0xb9, 0x9e, 0xf1, 0x57, 0xfb, 0x0f, 0x8c, 0x75, 0xd6, 0x31, 0x8f, 0x77, 0xf7, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x11, 0x2d, 0xd0, 0x25, 0x02, 0x00, 0x00,
}

func (m *EventRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRegister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRegister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRegister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
