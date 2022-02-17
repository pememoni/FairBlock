// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairblock/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSubmitEncrypted struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MessageHash  string `protobuf:"bytes,2,opt,name=messageHash,proto3" json:"messageHash,omitempty"`
	Encryption   string `protobuf:"bytes,3,opt,name=encryption,proto3" json:"encryption,omitempty"`
	TargetHeight string `protobuf:"bytes,4,opt,name=targetHeight,proto3" json:"targetHeight,omitempty"`
	Deposit      string `protobuf:"bytes,5,opt,name=deposit,proto3" json:"deposit,omitempty"`
}

func (m *MsgSubmitEncrypted) Reset()         { *m = MsgSubmitEncrypted{} }
func (m *MsgSubmitEncrypted) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEncrypted) ProtoMessage()    {}
func (*MsgSubmitEncrypted) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0a1902eeba049d, []int{0}
}
func (m *MsgSubmitEncrypted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEncrypted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEncrypted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEncrypted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEncrypted.Merge(m, src)
}
func (m *MsgSubmitEncrypted) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEncrypted) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEncrypted.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEncrypted proto.InternalMessageInfo

func (m *MsgSubmitEncrypted) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitEncrypted) GetMessageHash() string {
	if m != nil {
		return m.MessageHash
	}
	return ""
}

func (m *MsgSubmitEncrypted) GetEncryption() string {
	if m != nil {
		return m.Encryption
	}
	return ""
}

func (m *MsgSubmitEncrypted) GetTargetHeight() string {
	if m != nil {
		return m.TargetHeight
	}
	return ""
}

func (m *MsgSubmitEncrypted) GetDeposit() string {
	if m != nil {
		return m.Deposit
	}
	return ""
}

type MsgSubmitEncryptedResponse struct {
}

func (m *MsgSubmitEncryptedResponse) Reset()         { *m = MsgSubmitEncryptedResponse{} }
func (m *MsgSubmitEncryptedResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEncryptedResponse) ProtoMessage()    {}
func (*MsgSubmitEncryptedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0a1902eeba049d, []int{1}
}
func (m *MsgSubmitEncryptedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEncryptedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEncryptedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEncryptedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEncryptedResponse.Merge(m, src)
}
func (m *MsgSubmitEncryptedResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEncryptedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEncryptedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEncryptedResponse proto.InternalMessageInfo

type MsgCommitDecryption struct {
	Creator                string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PlaintextHash          string `protobuf:"bytes,2,opt,name=plaintextHash,proto3" json:"plaintextHash,omitempty"`
	PlaintextDecryptorHash string `protobuf:"bytes,3,opt,name=plaintextDecryptorHash,proto3" json:"plaintextDecryptorHash,omitempty"`
}

func (m *MsgCommitDecryption) Reset()         { *m = MsgCommitDecryption{} }
func (m *MsgCommitDecryption) String() string { return proto.CompactTextString(m) }
func (*MsgCommitDecryption) ProtoMessage()    {}
func (*MsgCommitDecryption) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0a1902eeba049d, []int{2}
}
func (m *MsgCommitDecryption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitDecryption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitDecryption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitDecryption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitDecryption.Merge(m, src)
}
func (m *MsgCommitDecryption) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitDecryption) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitDecryption.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitDecryption proto.InternalMessageInfo

func (m *MsgCommitDecryption) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCommitDecryption) GetPlaintextHash() string {
	if m != nil {
		return m.PlaintextHash
	}
	return ""
}

func (m *MsgCommitDecryption) GetPlaintextDecryptorHash() string {
	if m != nil {
		return m.PlaintextDecryptorHash
	}
	return ""
}

type MsgCommitDecryptionResponse struct {
}

func (m *MsgCommitDecryptionResponse) Reset()         { *m = MsgCommitDecryptionResponse{} }
func (m *MsgCommitDecryptionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCommitDecryptionResponse) ProtoMessage()    {}
func (*MsgCommitDecryptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0a1902eeba049d, []int{3}
}
func (m *MsgCommitDecryptionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitDecryptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitDecryptionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitDecryptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitDecryptionResponse.Merge(m, src)
}
func (m *MsgCommitDecryptionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitDecryptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitDecryptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitDecryptionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitEncrypted)(nil), "pememoni.fairblock.fairblock.MsgSubmitEncrypted")
	proto.RegisterType((*MsgSubmitEncryptedResponse)(nil), "pememoni.fairblock.fairblock.MsgSubmitEncryptedResponse")
	proto.RegisterType((*MsgCommitDecryption)(nil), "pememoni.fairblock.fairblock.MsgCommitDecryption")
	proto.RegisterType((*MsgCommitDecryptionResponse)(nil), "pememoni.fairblock.fairblock.MsgCommitDecryptionResponse")
}

func init() { proto.RegisterFile("fairblock/tx.proto", fileDescriptor_ad0a1902eeba049d) }

var fileDescriptor_ad0a1902eeba049d = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3f, 0x4f, 0x3a, 0x41,
	0x10, 0x65, 0xe1, 0xf7, 0xd3, 0x38, 0x6a, 0x34, 0x6b, 0x62, 0x2e, 0x88, 0x17, 0x72, 0xb1, 0xb0,
	0xba, 0xf3, 0x4f, 0x62, 0xb4, 0xc5, 0x3f, 0xc1, 0x82, 0x06, 0x3b, 0xbb, 0xbb, 0x63, 0x3c, 0x36,
	0xb2, 0xb7, 0x9b, 0xdd, 0x25, 0x81, 0xc2, 0xc4, 0xce, 0xca, 0xc4, 0x4f, 0xe2, 0xe7, 0xb0, 0xa4,
	0xb4, 0x34, 0xf0, 0x45, 0x8c, 0x07, 0x87, 0x07, 0x88, 0x86, 0x6e, 0xf7, 0xcd, 0xcc, 0x9b, 0x37,
	0x33, 0x0f, 0xe8, 0x9d, 0xcf, 0x54, 0xd0, 0x12, 0xe1, 0xbd, 0x67, 0x3a, 0xae, 0x54, 0xc2, 0x08,
	0x5a, 0x92, 0xc8, 0x91, 0x8b, 0x98, 0xb9, 0xe3, 0xe0, 0xf7, 0xcb, 0x79, 0x25, 0x40, 0x6b, 0x3a,
	0xba, 0x69, 0x07, 0x9c, 0x99, 0xcb, 0x38, 0x54, 0x5d, 0x69, 0xb0, 0x41, 0x2d, 0x58, 0x0e, 0x15,
	0xfa, 0x46, 0x28, 0x8b, 0x94, 0xc9, 0xfe, 0x4a, 0x3d, 0xfd, 0xd2, 0x32, 0xac, 0x72, 0xd4, 0xda,
	0x8f, 0xb0, 0xea, 0xeb, 0xa6, 0x95, 0x4f, 0xa2, 0x59, 0x88, 0xda, 0x00, 0x38, 0x24, 0x62, 0x22,
	0xb6, 0x0a, 0x49, 0x42, 0x06, 0xa1, 0x0e, 0xac, 0x19, 0x5f, 0x45, 0x68, 0xaa, 0xc8, 0xa2, 0xa6,
	0xb1, 0xfe, 0x25, 0x19, 0x13, 0xd8, 0x57, 0xff, 0x06, 0x4a, 0xa1, 0x99, 0xb1, 0xfe, 0x0f, 0xfb,
	0x8f, 0xbe, 0x4e, 0x09, 0x8a, 0xb3, 0x7a, 0xeb, 0xa8, 0xa5, 0x88, 0x35, 0x3a, 0xcf, 0x04, 0xb6,
	0x6a, 0x3a, 0x3a, 0x17, 0x9c, 0x33, 0x73, 0x81, 0xe3, 0x9e, 0xf3, 0xe7, 0xd9, 0x83, 0x75, 0xd9,
	0xf2, 0x59, 0x6c, 0xb0, 0x63, 0x32, 0x13, 0x4d, 0x82, 0xf4, 0x04, 0xb6, 0xc7, 0xc0, 0x88, 0x56,
	0xa8, 0x24, 0x7d, 0x38, 0xdf, 0x9c, 0xa8, 0xb3, 0x0b, 0x3b, 0x3f, 0xc8, 0x49, 0xe5, 0x1e, 0x3d,
	0xe5, 0xa1, 0x50, 0xd3, 0x11, 0x7d, 0x80, 0x8d, 0xe9, 0x0b, 0x1c, 0xb8, 0xbf, 0xdd, 0xcd, 0x9d,
	0xdd, 0x41, 0xf1, 0x74, 0xd1, 0x8a, 0x54, 0x06, 0x7d, 0x24, 0xb0, 0x39, 0xb3, 0xb2, 0xc3, 0x3f,
	0xe9, 0xa6, 0x4b, 0x8a, 0x67, 0x0b, 0x97, 0xa4, 0x12, 0x2a, 0xd7, 0x6f, 0x7d, 0x9b, 0xf4, 0xfa,
	0x36, 0xf9, 0xe8, 0xdb, 0xe4, 0x65, 0x60, 0xe7, 0x7a, 0x03, 0x3b, 0xf7, 0x3e, 0xb0, 0x73, 0xb7,
	0x5e, 0xc4, 0x4c, 0xb3, 0x1d, 0xb8, 0xa1, 0xe0, 0x5e, 0x4a, 0xef, 0x5d, 0xf9, 0x4c, 0x55, 0x12,
	0x9f, 0x77, 0xbc, 0x8c, 0xe7, 0xbb, 0x12, 0x75, 0xb0, 0x94, 0xf8, 0xfe, 0xf8, 0x33, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0x0e, 0x17, 0x20, 0x0d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SubmitEncrypted(ctx context.Context, in *MsgSubmitEncrypted, opts ...grpc.CallOption) (*MsgSubmitEncryptedResponse, error)
	CommitDecryption(ctx context.Context, in *MsgCommitDecryption, opts ...grpc.CallOption) (*MsgCommitDecryptionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitEncrypted(ctx context.Context, in *MsgSubmitEncrypted, opts ...grpc.CallOption) (*MsgSubmitEncryptedResponse, error) {
	out := new(MsgSubmitEncryptedResponse)
	err := c.cc.Invoke(ctx, "/pememoni.fairblock.fairblock.Msg/SubmitEncrypted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommitDecryption(ctx context.Context, in *MsgCommitDecryption, opts ...grpc.CallOption) (*MsgCommitDecryptionResponse, error) {
	out := new(MsgCommitDecryptionResponse)
	err := c.cc.Invoke(ctx, "/pememoni.fairblock.fairblock.Msg/CommitDecryption", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitEncrypted(context.Context, *MsgSubmitEncrypted) (*MsgSubmitEncryptedResponse, error)
	CommitDecryption(context.Context, *MsgCommitDecryption) (*MsgCommitDecryptionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitEncrypted(ctx context.Context, req *MsgSubmitEncrypted) (*MsgSubmitEncryptedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitEncrypted not implemented")
}
func (*UnimplementedMsgServer) CommitDecryption(ctx context.Context, req *MsgCommitDecryption) (*MsgCommitDecryptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitDecryption not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitEncrypted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitEncrypted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitEncrypted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pememoni.fairblock.fairblock.Msg/SubmitEncrypted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitEncrypted(ctx, req.(*MsgSubmitEncrypted))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommitDecryption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommitDecryption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommitDecryption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pememoni.fairblock.fairblock.Msg/CommitDecryption",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommitDecryption(ctx, req.(*MsgCommitDecryption))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pememoni.fairblock.fairblock.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitEncrypted",
			Handler:    _Msg_SubmitEncrypted_Handler,
		},
		{
			MethodName: "CommitDecryption",
			Handler:    _Msg_CommitDecryption_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fairblock/tx.proto",
}

func (m *MsgSubmitEncrypted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEncrypted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEncrypted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TargetHeight) > 0 {
		i -= len(m.TargetHeight)
		copy(dAtA[i:], m.TargetHeight)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TargetHeight)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Encryption) > 0 {
		i -= len(m.Encryption)
		copy(dAtA[i:], m.Encryption)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Encryption)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MessageHash) > 0 {
		i -= len(m.MessageHash)
		copy(dAtA[i:], m.MessageHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MessageHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitEncryptedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEncryptedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEncryptedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCommitDecryption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitDecryption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitDecryption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlaintextDecryptorHash) > 0 {
		i -= len(m.PlaintextDecryptorHash)
		copy(dAtA[i:], m.PlaintextDecryptorHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlaintextDecryptorHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PlaintextHash) > 0 {
		i -= len(m.PlaintextHash)
		copy(dAtA[i:], m.PlaintextHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlaintextHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCommitDecryptionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitDecryptionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitDecryptionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitEncrypted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MessageHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Encryption)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TargetHeight)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitEncryptedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCommitDecryption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PlaintextHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PlaintextDecryptorHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCommitDecryptionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitEncrypted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitEncrypted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEncrypted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encryption", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Encryption = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitEncryptedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitEncryptedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEncryptedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCommitDecryption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCommitDecryption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitDecryption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlaintextHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlaintextHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlaintextDecryptorHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlaintextDecryptorHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCommitDecryptionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCommitDecryptionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitDecryptionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
