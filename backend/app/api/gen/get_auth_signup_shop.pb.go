// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_auth_signup_shop.proto

package gen

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetAuthSignupShopRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthSignupShopRequest) Reset()         { *m = GetAuthSignupShopRequest{} }
func (m *GetAuthSignupShopRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthSignupShopRequest) ProtoMessage()    {}
func (*GetAuthSignupShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae68f55cc96af8a, []int{0}
}
func (m *GetAuthSignupShopRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthSignupShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthSignupShopRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthSignupShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthSignupShopRequest.Merge(m, src)
}
func (m *GetAuthSignupShopRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthSignupShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthSignupShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthSignupShopRequest proto.InternalMessageInfo

func (m *GetAuthSignupShopRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetAuthSignupShopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthSignupShopResponse) Reset()         { *m = GetAuthSignupShopResponse{} }
func (m *GetAuthSignupShopResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthSignupShopResponse) ProtoMessage()    {}
func (*GetAuthSignupShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae68f55cc96af8a, []int{1}
}
func (m *GetAuthSignupShopResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthSignupShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthSignupShopResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthSignupShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthSignupShopResponse.Merge(m, src)
}
func (m *GetAuthSignupShopResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthSignupShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthSignupShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthSignupShopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetAuthSignupShopRequest)(nil), "api.GetAuthSignupShopRequest")
	proto.RegisterType((*GetAuthSignupShopResponse)(nil), "api.GetAuthSignupShopResponse")
}

func init() { proto.RegisterFile("get_auth_signup_shop.proto", fileDescriptor_aae68f55cc96af8a) }

var fileDescriptor_aae68f55cc96af8a = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x4f, 0x2d, 0x89,
	0x4f, 0x2c, 0x2d, 0xc9, 0x88, 0x2f, 0xce, 0x4c, 0xcf, 0x2b, 0x2d, 0x88, 0x2f, 0xce, 0xc8, 0x2f,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x12, 0x2f, 0x4b, 0xcc,
	0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0xb2, 0x4a, 0x96, 0x5c, 0x12, 0xee, 0xa9,
	0x25, 0x8e, 0xa5, 0x25, 0x19, 0xc1, 0x60, 0x9d, 0xc1, 0x19, 0xf9, 0x05, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0xb2, 0x5c, 0xac, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x4e, 0xec, 0xbf, 0x9c, 0x58, 0x8a, 0x98, 0x04, 0x18, 0x83, 0x20, 0xa2, 0x4a, 0xd2,
	0x5c, 0x92, 0x58, 0xb4, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x3a, 0x89, 0x9d, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0xe8,
	0xa7, 0xa7, 0xe6, 0x25, 0xb1, 0x81, 0xad, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xc1,
	0xa2, 0x8e, 0xb2, 0x00, 0x00, 0x00,
}

func (m *GetAuthSignupShopRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthSignupShopRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthSignupShopRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintGetAuthSignupShop(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAuthSignupShopResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthSignupShopResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthSignupShopResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintGetAuthSignupShop(dAtA []byte, offset int, v uint64) int {
	offset -= sovGetAuthSignupShop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetAuthSignupShopRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovGetAuthSignupShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAuthSignupShopResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGetAuthSignupShop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGetAuthSignupShop(x uint64) (n int) {
	return sovGetAuthSignupShop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetAuthSignupShopRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGetAuthSignupShop
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
			return fmt.Errorf("proto: GetAuthSignupShopRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthSignupShopRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGetAuthSignupShop
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
				return ErrInvalidLengthGetAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGetAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGetAuthSignupShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGetAuthSignupShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetAuthSignupShopResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGetAuthSignupShop
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
			return fmt.Errorf("proto: GetAuthSignupShopResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthSignupShopResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGetAuthSignupShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGetAuthSignupShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGetAuthSignupShop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGetAuthSignupShop
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
					return 0, ErrIntOverflowGetAuthSignupShop
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
					return 0, ErrIntOverflowGetAuthSignupShop
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
				return 0, ErrInvalidLengthGetAuthSignupShop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGetAuthSignupShop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGetAuthSignupShop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGetAuthSignupShop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGetAuthSignupShop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGetAuthSignupShop = fmt.Errorf("proto: unexpected end of group")
)
