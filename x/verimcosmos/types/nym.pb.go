// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verimcosmos/nym.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Nym struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Alias   string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Verkey  string `protobuf:"bytes,4,opt,name=verkey,proto3" json:"verkey,omitempty"`
	Did     string `protobuf:"bytes,5,opt,name=did,proto3" json:"did,omitempty"`
	Role    string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *Nym) Reset()         { *m = Nym{} }
func (m *Nym) String() string { return proto.CompactTextString(m) }
func (*Nym) ProtoMessage()    {}
func (*Nym) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ef0d1568db26786, []int{0}
}
func (m *Nym) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nym) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nym.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nym) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nym.Merge(m, src)
}
func (m *Nym) XXX_Size() int {
	return m.Size()
}
func (m *Nym) XXX_DiscardUnknown() {
	xxx_messageInfo_Nym.DiscardUnknown(m)
}

var xxx_messageInfo_Nym proto.InternalMessageInfo

func (m *Nym) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Nym) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nym) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Nym) GetVerkey() string {
	if m != nil {
		return m.Verkey
	}
	return ""
}

func (m *Nym) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Nym) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*Nym)(nil), "verimid.verimcosmos.verimcosmos.Nym")
}

func init() { proto.RegisterFile("verimcosmos/nym.proto", fileDescriptor_0ef0d1568db26786) }

var fileDescriptor_0ef0d1568db26786 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4b, 0x2d, 0xca,
	0xcc, 0x4d, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xcf, 0xab, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x07, 0x0b, 0x67, 0xa6, 0xe8, 0x21, 0x49, 0x23, 0xb3, 0xa5, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x6a, 0xf5, 0x41, 0x2c, 0x88, 0x36, 0xa5, 0x66, 0x46, 0x2e, 0x66, 0xbf, 0xca,
	0x5c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x18, 0x57, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x25, 0x88, 0x29, 0x33, 0x45, 0x48, 0x84, 0x8b, 0x35, 0x31, 0x27, 0x33, 0xb1, 0x58, 0x82, 0x19,
	0xac, 0x0e, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x4b, 0x2d, 0xca, 0x4e, 0xad, 0x94, 0x60, 0x01,
	0x0b, 0x43, 0x79, 0x42, 0x02, 0x5c, 0xcc, 0x29, 0x99, 0x29, 0x12, 0xac, 0x60, 0x41, 0x10, 0x53,
	0x48, 0x88, 0x8b, 0xa5, 0x28, 0x3f, 0x27, 0x55, 0x82, 0x0d, 0x2c, 0x04, 0x66, 0x3b, 0xf9, 0x9d,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x49, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd8, 0x37, 0xba, 0x99, 0x29, 0x50, 0x06, 0x34, 0x08, 0x2a,
	0xf4, 0x91, 0x03, 0xa4, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x39, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x70, 0x00, 0xee, 0x87, 0x2c, 0x01, 0x00, 0x00,
}

func (m *Nym) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nym) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nym) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintNym(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintNym(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Verkey) > 0 {
		i -= len(m.Verkey)
		copy(dAtA[i:], m.Verkey)
		i = encodeVarintNym(dAtA, i, uint64(len(m.Verkey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Alias) > 0 {
		i -= len(m.Alias)
		copy(dAtA[i:], m.Alias)
		i = encodeVarintNym(dAtA, i, uint64(len(m.Alias)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintNym(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNym(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNym(dAtA []byte, offset int, v uint64) int {
	offset -= sovNym(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Nym) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNym(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovNym(uint64(m.Id))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovNym(uint64(l))
	}
	l = len(m.Verkey)
	if l > 0 {
		n += 1 + l + sovNym(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovNym(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovNym(uint64(l))
	}
	return n
}

func sovNym(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNym(x uint64) (n int) {
	return sovNym(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Nym) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNym
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
			return fmt.Errorf("proto: Nym: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nym: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNym
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
				return ErrInvalidLengthNym
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNym
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNym
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNym
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
				return ErrInvalidLengthNym
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNym
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNym
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
				return ErrInvalidLengthNym
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNym
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Verkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNym
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
				return ErrInvalidLengthNym
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNym
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNym
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
				return ErrInvalidLengthNym
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNym
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNym(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNym
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
func skipNym(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNym
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
					return 0, ErrIntOverflowNym
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
					return 0, ErrIntOverflowNym
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
				return 0, ErrInvalidLengthNym
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNym
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNym
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNym        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNym          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNym = fmt.Errorf("proto: unexpected end of group")
)
