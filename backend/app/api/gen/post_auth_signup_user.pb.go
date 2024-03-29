// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: post_auth_signup_user.proto

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

type User struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstNameHuri        string   `protobuf:"bytes,4,opt,name=first_name_huri,json=firstNameHuri,proto3" json:"first_name_huri,omitempty"`
	LastNameHuri         string   `protobuf:"bytes,5,opt,name=last_name_huri,json=lastNameHuri,proto3" json:"last_name_huri,omitempty"`
	ScreenName           string   `protobuf:"bytes,6,opt,name=screen_name,json=screenName,proto3" json:"screen_name,omitempty"`
	Birthday             string   `protobuf:"bytes,7,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_25bf66c039dc580e, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetFirstNameHuri() string {
	if m != nil {
		return m.FirstNameHuri
	}
	return ""
}

func (m *User) GetLastNameHuri() string {
	if m != nil {
		return m.LastNameHuri
	}
	return ""
}

func (m *User) GetScreenName() string {
	if m != nil {
		return m.ScreenName
	}
	return ""
}

func (m *User) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

type PostAuthProvisionalSignupRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm      string   `protobuf:"bytes,3,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostAuthProvisionalSignupRequest) Reset()         { *m = PostAuthProvisionalSignupRequest{} }
func (m *PostAuthProvisionalSignupRequest) String() string { return proto.CompactTextString(m) }
func (*PostAuthProvisionalSignupRequest) ProtoMessage()    {}
func (*PostAuthProvisionalSignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25bf66c039dc580e, []int{1}
}
func (m *PostAuthProvisionalSignupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostAuthProvisionalSignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostAuthProvisionalSignupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostAuthProvisionalSignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostAuthProvisionalSignupRequest.Merge(m, src)
}
func (m *PostAuthProvisionalSignupRequest) XXX_Size() int {
	return m.Size()
}
func (m *PostAuthProvisionalSignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostAuthProvisionalSignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostAuthProvisionalSignupRequest proto.InternalMessageInfo

func (m *PostAuthProvisionalSignupRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PostAuthProvisionalSignupRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PostAuthProvisionalSignupRequest) GetPasswordConfirm() string {
	if m != nil {
		return m.PasswordConfirm
	}
	return ""
}

type PostAuthProvisionalSignupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostAuthProvisionalSignupResponse) Reset()         { *m = PostAuthProvisionalSignupResponse{} }
func (m *PostAuthProvisionalSignupResponse) String() string { return proto.CompactTextString(m) }
func (*PostAuthProvisionalSignupResponse) ProtoMessage()    {}
func (*PostAuthProvisionalSignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25bf66c039dc580e, []int{2}
}
func (m *PostAuthProvisionalSignupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostAuthProvisionalSignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostAuthProvisionalSignupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostAuthProvisionalSignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostAuthProvisionalSignupResponse.Merge(m, src)
}
func (m *PostAuthProvisionalSignupResponse) XXX_Size() int {
	return m.Size()
}
func (m *PostAuthProvisionalSignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostAuthProvisionalSignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostAuthProvisionalSignupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*PostAuthProvisionalSignupRequest)(nil), "api.PostAuthProvisionalSignupRequest")
	proto.RegisterType((*PostAuthProvisionalSignupResponse)(nil), "api.PostAuthProvisionalSignupResponse")
}

func init() { proto.RegisterFile("post_auth_signup_user.proto", fileDescriptor_25bf66c039dc580e) }

var fileDescriptor_25bf66c039dc580e = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x86, 0xbf, 0x42, 0x81, 0xf6, 0xf0, 0x29, 0xa4, 0x0b, 0x6d, 0x34, 0x34, 0x08, 0xc6, 0xb0,
	0x11, 0x12, 0xbc, 0x02, 0x71, 0xe3, 0xca, 0x90, 0x1a, 0x37, 0x6e, 0x9a, 0x01, 0x06, 0x3a, 0x49,
	0xdb, 0x19, 0xe7, 0x07, 0xe3, 0x9d, 0x18, 0x97, 0x5e, 0x8d, 0x4b, 0x2f, 0xc1, 0xe0, 0x5d, 0xb8,
	0x32, 0x9d, 0xa6, 0xd5, 0xd9, 0xb8, 0x3b, 0x39, 0xcf, 0x93, 0x77, 0x4e, 0xde, 0x0c, 0x1c, 0x33,
	0x2a, 0x64, 0x84, 0x94, 0x8c, 0x23, 0x41, 0x36, 0x99, 0x62, 0x91, 0x12, 0x98, 0x8f, 0x19, 0xa7,
	0x92, 0x7a, 0x75, 0xc4, 0xc8, 0xd1, 0xe1, 0x16, 0x25, 0x64, 0x85, 0x24, 0x9e, 0x94, 0x43, 0x41,
	0x07, 0xaf, 0x35, 0xb0, 0xef, 0x04, 0xe6, 0x5e, 0x0f, 0x1a, 0x38, 0x45, 0x24, 0xf1, 0xad, 0xbe,
	0x35, 0x72, 0x67, 0xad, 0xaf, 0x99, 0xcd, 0x6b, 0x5d, 0x2b, 0x2c, 0xb6, 0xde, 0x19, 0xc0, 0x9a,
	0x70, 0x21, 0xa3, 0x0c, 0xa5, 0xd8, 0xaf, 0x99, 0x8e, 0xab, 0xd1, 0x0d, 0x4a, 0xb1, 0x77, 0x0a,
	0x6e, 0x82, 0x4a, 0xad, 0x6e, 0x6a, 0x4e, 0x4e, 0xb4, 0x35, 0x81, 0xce, 0x4f, 0x5a, 0x14, 0x2b,
	0x4e, 0x7c, 0xdb, 0x74, 0xf7, 0xaa, 0xc8, 0x6b, 0xc5, 0x89, 0x77, 0x0e, 0xfb, 0x55, 0x6c, 0xe1,
	0x37, 0x4c, 0xff, 0x7f, 0x99, 0xad, 0xf5, 0x11, 0xb4, 0xc5, 0x92, 0x63, 0x9c, 0x15, 0x77, 0x34,
	0x4d, 0x17, 0x0a, 0xa6, 0x2f, 0x19, 0x82, 0xb3, 0x20, 0x5c, 0xc6, 0x2b, 0xf4, 0xe4, 0xb7, 0x7e,
	0x6b, 0x10, 0x56, 0x60, 0xf0, 0x62, 0x41, 0x7f, 0x4e, 0x85, 0xbc, 0x54, 0x32, 0x9e, 0x73, 0xba,
	0x25, 0x82, 0xd0, 0x0c, 0x25, 0xb7, 0xba, 0xeb, 0x10, 0x3f, 0x28, 0x2c, 0xa4, 0xd7, 0x03, 0x3b,
	0x6f, 0x5d, 0xf7, 0xd7, 0x9e, 0xba, 0x63, 0xc4, 0xc8, 0x38, 0x6f, 0x36, 0xd4, 0xeb, 0xfc, 0x21,
	0x86, 0x84, 0x78, 0xa4, 0x7c, 0x65, 0xd6, 0xe7, 0x84, 0x15, 0xf0, 0xa6, 0xd0, 0x2d, 0xe7, 0x68,
	0x49, 0xb3, 0x35, 0xe1, 0xa9, 0x59, 0xa2, 0x13, 0x76, 0x4a, 0xe1, 0xaa, 0xe0, 0x83, 0x21, 0x9c,
	0xfc, 0x71, 0x9b, 0x60, 0x34, 0x13, 0x78, 0x76, 0xf0, 0xb6, 0x0b, 0xac, 0xf7, 0x5d, 0x60, 0x7d,
	0xec, 0x02, 0xeb, 0xf9, 0x33, 0xf8, 0x77, 0x6f, 0x4f, 0x36, 0x38, 0x5b, 0x34, 0xf5, 0x2f, 0xb8,
	0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xed, 0x5b, 0xbe, 0x3e, 0x42, 0x02, 0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Birthday) > 0 {
		i -= len(m.Birthday)
		copy(dAtA[i:], m.Birthday)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.Birthday)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ScreenName) > 0 {
		i -= len(m.ScreenName)
		copy(dAtA[i:], m.ScreenName)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.ScreenName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LastNameHuri) > 0 {
		i -= len(m.LastNameHuri)
		copy(dAtA[i:], m.LastNameHuri)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.LastNameHuri)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FirstNameHuri) > 0 {
		i -= len(m.FirstNameHuri)
		copy(dAtA[i:], m.FirstNameHuri)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.FirstNameHuri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostAuthProvisionalSignupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostAuthProvisionalSignupRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostAuthProvisionalSignupRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PasswordConfirm) > 0 {
		i -= len(m.PasswordConfirm)
		copy(dAtA[i:], m.PasswordConfirm)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.PasswordConfirm)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if m.User != nil {
		{
			size, err := m.User.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPostAuthSignupUser(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostAuthProvisionalSignupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostAuthProvisionalSignupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostAuthProvisionalSignupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintPostAuthSignupUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovPostAuthSignupUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.FirstNameHuri)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.LastNameHuri)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.ScreenName)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.Birthday)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostAuthProvisionalSignupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	l = len(m.PasswordConfirm)
	if l > 0 {
		n += 1 + l + sovPostAuthSignupUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostAuthProvisionalSignupResponse) Size() (n int) {
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

func sovPostAuthSignupUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPostAuthSignupUser(x uint64) (n int) {
	return sovPostAuthSignupUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstNameHuri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstNameHuri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastNameHuri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastNameHuri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScreenName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScreenName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Birthday", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Birthday = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupUser
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
func (m *PostAuthProvisionalSignupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupUser
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
			return fmt.Errorf("proto: PostAuthProvisionalSignupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostAuthProvisionalSignupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
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
					return ErrIntOverflowPostAuthSignupUser
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
				return ErrInvalidLengthPostAuthSignupUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostAuthSignupUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PasswordConfirm = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupUser
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
func (m *PostAuthProvisionalSignupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostAuthSignupUser
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
			return fmt.Errorf("proto: PostAuthProvisionalSignupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostAuthProvisionalSignupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPostAuthSignupUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPostAuthSignupUser
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
func skipPostAuthSignupUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPostAuthSignupUser
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
					return 0, ErrIntOverflowPostAuthSignupUser
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
					return 0, ErrIntOverflowPostAuthSignupUser
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
				return 0, ErrInvalidLengthPostAuthSignupUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPostAuthSignupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPostAuthSignupUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPostAuthSignupUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPostAuthSignupUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPostAuthSignupUser = fmt.Errorf("proto: unexpected end of group")
)
