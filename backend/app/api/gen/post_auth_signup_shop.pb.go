// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: post_auth_signup_shop.proto

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

type Owner struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NameKana             string   `protobuf:"bytes,2,opt,name=name_kana,json=nameKana,proto3" json:"name_kana,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_308dcc44252e1e0e, []int{0}
}
func (m *Owner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return m.Size()
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Owner) GetNameKana() string {
	if m != nil {
		return m.NameKana
	}
	return ""
}

func (m *Owner) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Shop struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NameKana             string   `protobuf:"bytes,2,opt,name=name_kana,json=nameKana,proto3" json:"name_kana,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Owner                *Owner   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shop) Reset()         { *m = Shop{} }
func (m *Shop) String() string { return proto.CompactTextString(m) }
func (*Shop) ProtoMessage()    {}
func (*Shop) Descriptor() ([]byte, []int) {
	return fileDescriptor_308dcc44252e1e0e, []int{1}
}
func (m *Shop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Shop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Shop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Shop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shop.Merge(m, src)
}
func (m *Shop) XXX_Size() int {
	return m.Size()
}
func (m *Shop) XXX_DiscardUnknown() {
	xxx_messageInfo_Shop.DiscardUnknown(m)
}

var xxx_messageInfo_Shop proto.InternalMessageInfo

func (m *Shop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Shop) GetNameKana() string {
	if m != nil {
		return m.NameKana
	}
	return ""
}

func (m *Shop) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Shop) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Shop) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

type PostAuthProvisionalShopSignupRequest struct {
	Shop                 *Shop    `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm      string   `protobuf:"bytes,3,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
	OpeningHourd         uint64   `protobuf:"varint,4,opt,name=opening_hourd,json=openingHourd,proto3" json:"opening_hourd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostAuthProvisionalShopSignupRequest) Reset()         { *m = PostAuthProvisionalShopSignupRequest{} }
func (m *PostAuthProvisionalShopSignupRequest) String() string { return proto.CompactTextString(m) }
func (*PostAuthProvisionalShopSignupRequest) ProtoMessage()    {}
func (*PostAuthProvisionalShopSignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_308dcc44252e1e0e, []int{2}
}
func (m *PostAuthProvisionalShopSignupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostAuthProvisionalShopSignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostAuthProvisionalShopSignupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostAuthProvisionalShopSignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostAuthProvisionalShopSignupRequest.Merge(m, src)
}
func (m *PostAuthProvisionalShopSignupRequest) XXX_Size() int {
	return m.Size()
}
func (m *PostAuthProvisionalShopSignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostAuthProvisionalShopSignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostAuthProvisionalShopSignupRequest proto.InternalMessageInfo

func (m *PostAuthProvisionalShopSignupRequest) GetShop() *Shop {
	if m != nil {
		return m.Shop
	}
	return nil
}

func (m *PostAuthProvisionalShopSignupRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PostAuthProvisionalShopSignupRequest) GetPasswordConfirm() string {
	if m != nil {
		return m.PasswordConfirm
	}
	return ""
}

func (m *PostAuthProvisionalShopSignupRequest) GetOpeningHourd() uint64 {
	if m != nil {
		return m.OpeningHourd
	}
	return 0
}

type PostAuthProvisionalShopSignupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostAuthProvisionalShopSignupResponse) Reset()         { *m = PostAuthProvisionalShopSignupResponse{} }
func (m *PostAuthProvisionalShopSignupResponse) String() string { return proto.CompactTextString(m) }
func (*PostAuthProvisionalShopSignupResponse) ProtoMessage()    {}
func (*PostAuthProvisionalShopSignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_308dcc44252e1e0e, []int{3}
}
func (m *PostAuthProvisionalShopSignupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostAuthProvisionalShopSignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostAuthProvisionalShopSignupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostAuthProvisionalShopSignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostAuthProvisionalShopSignupResponse.Merge(m, src)
}
func (m *PostAuthProvisionalShopSignupResponse) XXX_Size() int {
	return m.Size()
}
func (m *PostAuthProvisionalShopSignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostAuthProvisionalShopSignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostAuthProvisionalShopSignupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Owner)(nil), "api.Owner")
	proto.RegisterType((*Shop)(nil), "api.Shop")
	proto.RegisterType((*PostAuthProvisionalShopSignupRequest)(nil), "api.PostAuthProvisionalShopSignupRequest")
	proto.RegisterType((*PostAuthProvisionalShopSignupResponse)(nil), "api.PostAuthProvisionalShopSignupResponse")
}

func init() { proto.RegisterFile("post_auth_signup_shop.proto", fileDescriptor_308dcc44252e1e0e) }

var fileDescriptor_308dcc44252e1e0e = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x19, 0x68, 0xef, 0x2d, 0xc3, 0xbd, 0x91, 0xcc, 0x42, 0x1b, 0x89, 0x0d, 0xa9, 0x18,
	0x88, 0x31, 0x90, 0xd4, 0x27, 0xb0, 0x6e, 0x4c, 0x5c, 0x48, 0xca, 0xce, 0x4d, 0x1d, 0x65, 0xa4,
	0x13, 0x61, 0xce, 0xd8, 0x69, 0xe1, 0x1d, 0x7c, 0x02, 0xdf, 0xc3, 0x97, 0x70, 0xe9, 0xca, 0xb5,
	0xc1, 0xb7, 0x60, 0x65, 0x66, 0x2a, 0x6a, 0x13, 0x13, 0x37, 0xae, 0x7a, 0x72, 0xfe, 0xaf, 0xe7,
	0x9c, 0xff, 0xcf, 0xe0, 0x96, 0x04, 0x95, 0xc5, 0x34, 0xcf, 0x92, 0x58, 0xf1, 0x89, 0xc8, 0x65,
	0xac, 0x12, 0x90, 0x7d, 0x99, 0x42, 0x06, 0xa4, 0x46, 0x25, 0xdf, 0xde, 0x9a, 0xd3, 0x29, 0x1f,
	0xd3, 0x8c, 0x0d, 0xd6, 0x45, 0xa1, 0xfa, 0x1c, 0xdb, 0x67, 0x0b, 0xc1, 0x52, 0xd2, 0xc2, 0x96,
	0xa0, 0x33, 0xe6, 0xa2, 0x36, 0xea, 0xd5, 0xc3, 0xbf, 0xab, 0xd0, 0x4a, 0xab, 0x4d, 0x14, 0x99,
	0x26, 0xe9, 0xe0, 0xba, 0xfe, 0xc6, 0x37, 0x54, 0x50, 0xb7, 0x5a, 0x26, 0x1c, 0xad, 0x9c, 0x52,
	0x41, 0xc9, 0x0e, 0xb6, 0x65, 0x02, 0x82, 0xb9, 0xb5, 0x32, 0x51, 0x74, 0xfd, 0x07, 0x84, 0xad,
	0x51, 0x02, 0xf2, 0x97, 0x56, 0xb1, 0x19, 0xe5, 0xd3, 0xd2, 0xaa, 0x0b, 0x14, 0x15, 0xdd, 0xcf,
	0x4b, 0xac, 0xef, 0x2e, 0x21, 0xfb, 0xd8, 0x06, 0x6d, 0xda, 0xb5, 0xdb, 0xa8, 0xd7, 0x08, 0x70,
	0x9f, 0x4a, 0xde, 0x37, 0x31, 0x84, 0xce, 0x2a, 0xb4, 0xef, 0x90, 0x61, 0x0d, 0xe2, 0x3f, 0x23,
	0xdc, 0x19, 0x82, 0xca, 0x8e, 0xf2, 0x2c, 0x19, 0xa6, 0x30, 0xe7, 0x8a, 0x83, 0xa0, 0x53, 0x6d,
	0x64, 0x64, 0xb2, 0x8e, 0xd8, 0x6d, 0xce, 0x54, 0x46, 0xba, 0xd8, 0xd2, 0xa9, 0x1b, 0x57, 0x8d,
	0xa0, 0x6e, 0x66, 0x6a, 0xea, 0xcb, 0x48, 0x03, 0x90, 0x5d, 0xec, 0x48, 0xaa, 0xd4, 0x02, 0xd2,
	0x71, 0xd9, 0xa0, 0x13, 0x7d, 0x08, 0x24, 0xc0, 0xcd, 0x75, 0x1d, 0x5f, 0x81, 0xb8, 0xe6, 0xe9,
	0xac, 0x1c, 0xab, 0x13, 0x6d, 0xac, 0x81, 0xe3, 0x42, 0x27, 0x07, 0xf8, 0x3f, 0x48, 0x26, 0xb8,
	0x98, 0xc4, 0x09, 0xe4, 0xe9, 0xd8, 0xb8, 0xb7, 0xcc, 0x0f, 0x41, 0xb5, 0x5d, 0x89, 0xfe, 0xbd,
	0xab, 0x27, 0x5a, 0xf4, 0xbb, 0x78, 0xef, 0x07, 0x5f, 0x4a, 0x82, 0x50, 0x2c, 0xdc, 0x7c, 0x5c,
	0x7a, 0xe8, 0x69, 0xe9, 0xa1, 0x97, 0xa5, 0x87, 0xee, 0x5f, 0xbd, 0xca, 0xb9, 0x35, 0x98, 0x30,
	0x71, 0xf9, 0xc7, 0xbc, 0xa0, 0xc3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xbe, 0x43, 0xde,
	0x7e, 0x02, 0x00, 0x00,
}

func (m *Owner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Owner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Owner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Phone) > 0 {
		i -= len(m.Phone)
		copy(dAtA[i:], m.Phone)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.Phone)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NameKana) > 0 {
		i -= len(m.NameKana)
		copy(dAtA[i:], m.NameKana)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.NameKana)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Shop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Shop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Shop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Owner != nil {
		{
			size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Phone) > 0 {
		i -= len(m.Phone)
		copy(dAtA[i:], m.Phone)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.Phone)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NameKana) > 0 {
		i -= len(m.NameKana)
		copy(dAtA[i:], m.NameKana)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.NameKana)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostAuthProvisionalShopSignupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostAuthProvisionalShopSignupRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostAuthProvisionalShopSignupRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OpeningHourd != 0 {
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(m.OpeningHourd))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PasswordConfirm) > 0 {
		i -= len(m.PasswordConfirm)
		copy(dAtA[i:], m.PasswordConfirm)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.PasswordConfirm)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if m.Shop != nil {
		{
			size, err := m.Shop.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPostAuthSignupShop(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostAuthProvisionalShopSignupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostAuthProvisionalShopSignupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostAuthProvisionalShopSignupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintPostAuthSignupShop(dAtA []byte, offset int, v uint64) int {
	offset -= sovPostAuthSignupShop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Owner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.NameKana)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Shop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.NameKana)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	if m.Owner != nil {
		l = m.Owner.Size()
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostAuthProvisionalShopSignupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Shop != nil {
		l = m.Shop.Size()
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	l = len(m.PasswordConfirm)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupShop(uint64(l))
	}
	if m.OpeningHourd != 0 {
		n += 1 + sovPostAuthSignupShop(uint64(m.OpeningHourd))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostAuthProvisionalShopSignupResponse) Size() (n int) {
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

func sovPostAuthSignupShop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPostAuthSignupShop(x uint64) (n int) {
	return sovPostAuthSignupShop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Owner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupShop
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
			return fmt.Errorf("proto: Owner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Owner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NameKana", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NameKana = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupShop
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
func (m *Shop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupShop
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
			return fmt.Errorf("proto: Shop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Shop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NameKana", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NameKana = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Owner == nil {
				m.Owner = &Owner{}
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupShop
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
func (m *PostAuthProvisionalShopSignupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupShop
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
			return fmt.Errorf("proto: PostAuthProvisionalShopSignupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostAuthProvisionalShopSignupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shop == nil {
				m.Shop = &Shop{}
			}
			if err := m.Shop.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PasswordConfirm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
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
				return ErrInvalidLengthPostAuthSignupShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PasswordConfirm = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpeningHourd", wireType)
			}
			m.OpeningHourd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpeningHourd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupShop
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
func (m *PostAuthProvisionalShopSignupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupShop
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
			return fmt.Errorf("proto: PostAuthProvisionalShopSignupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostAuthProvisionalShopSignupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupShop
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
func skipPostAuthSignupShop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPostAuthSignupShop
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
					return 0, ErrIntOverflowPostAuthSignupShop
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
					return 0, ErrIntOverflowPostAuthSignupShop
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
				return 0, ErrInvalidLengthPostAuthSignupShop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPostAuthSignupShop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPostAuthSignupShop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPostAuthSignupShop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPostAuthSignupShop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPostAuthSignupShop = fmt.Errorf("proto: unexpected end of group")
)