// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: goan/tx.proto

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

type MsgOddiyTx struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Sender      string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver    string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount      string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee         string `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee,omitempty"`
	FeeReceiver string `protobuf:"bytes,6,opt,name=feeReceiver,proto3" json:"feeReceiver,omitempty"`
	TxType      string `protobuf:"bytes,7,opt,name=txType,proto3" json:"txType,omitempty"`
	ServiceName string `protobuf:"bytes,8,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (m *MsgOddiyTx) Reset()         { *m = MsgOddiyTx{} }
func (m *MsgOddiyTx) String() string { return proto.CompactTextString(m) }
func (*MsgOddiyTx) ProtoMessage()    {}
func (*MsgOddiyTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_60954180b4b874ef, []int{0}
}
func (m *MsgOddiyTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOddiyTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOddiyTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOddiyTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOddiyTx.Merge(m, src)
}
func (m *MsgOddiyTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgOddiyTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOddiyTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOddiyTx proto.InternalMessageInfo

func (m *MsgOddiyTx) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgOddiyTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgOddiyTx) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgOddiyTx) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MsgOddiyTx) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *MsgOddiyTx) GetFeeReceiver() string {
	if m != nil {
		return m.FeeReceiver
	}
	return ""
}

func (m *MsgOddiyTx) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *MsgOddiyTx) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type MsgOddiyTxResponse struct {
}

func (m *MsgOddiyTxResponse) Reset()         { *m = MsgOddiyTxResponse{} }
func (m *MsgOddiyTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgOddiyTxResponse) ProtoMessage()    {}
func (*MsgOddiyTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60954180b4b874ef, []int{1}
}
func (m *MsgOddiyTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOddiyTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOddiyTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOddiyTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOddiyTxResponse.Merge(m, src)
}
func (m *MsgOddiyTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgOddiyTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOddiyTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOddiyTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgOddiyTx)(nil), "goan.goan.MsgOddiyTx")
	proto.RegisterType((*MsgOddiyTxResponse)(nil), "goan.goan.MsgOddiyTxResponse")
}

func init() { proto.RegisterFile("goan/tx.proto", fileDescriptor_60954180b4b874ef) }

var fileDescriptor_60954180b4b874ef = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x1b, 0xab, 0xed, 0xf6, 0x17, 0x41, 0x83, 0x4a, 0x18, 0x18, 0xc6, 0x4e, 0x82, 0xd0,
	0x81, 0x7e, 0x00, 0xc1, 0x83, 0xb7, 0x29, 0x94, 0x9d, 0xbc, 0xd5, 0xf6, 0xbf, 0xd2, 0xc3, 0x9a,
	0x92, 0xc4, 0xd1, 0x7e, 0x0b, 0x3f, 0x96, 0xc7, 0x1d, 0x3d, 0x8e, 0xf6, 0x8b, 0x48, 0x12, 0xbb,
	0xf5, 0xb0, 0x4b, 0xc8, 0xef, 0x3d, 0xde, 0x83, 0xff, 0x83, 0x8b, 0x5c, 0x24, 0xe5, 0x5c, 0xd7,
	0x51, 0x25, 0x85, 0x16, 0x74, 0x6c, 0x30, 0x32, 0xcf, 0x6c, 0x47, 0x00, 0x16, 0x2a, 0x7f, 0xcf,
	0xb2, 0xa2, 0x59, 0xd6, 0x94, 0x41, 0x98, 0x4a, 0x4c, 0xb4, 0x90, 0x8c, 0x4c, 0xc9, 0xfd, 0x38,
	0xee, 0x91, 0xde, 0x42, 0xa0, 0xb0, 0xcc, 0x50, 0xb2, 0x13, 0x6b, 0xfc, 0x13, 0x9d, 0xc0, 0x48,
	0x62, 0x8a, 0xc5, 0x06, 0x25, 0xf3, 0xad, 0xb3, 0x67, 0x93, 0x49, 0xd6, 0xe2, 0xab, 0xd4, 0xec,
	0xd4, 0x65, 0x1c, 0xd1, 0x4b, 0xf0, 0x57, 0x88, 0xec, 0xcc, 0x8a, 0xe6, 0x4b, 0xa7, 0x70, 0xbe,
	0x42, 0x8c, 0xfb, 0xa2, 0xc0, 0x3a, 0x43, 0xc9, 0x74, 0xe9, 0x7a, 0xd9, 0x54, 0xc8, 0x42, 0xd7,
	0xe5, 0xc8, 0x24, 0x15, 0xca, 0x4d, 0x91, 0xe2, 0x5b, 0xb2, 0x46, 0x36, 0x72, 0xc9, 0x81, 0x34,
	0xbb, 0x06, 0x7a, 0xb8, 0x30, 0x46, 0x55, 0x89, 0x52, 0xe1, 0xe3, 0x2b, 0xf8, 0x0b, 0x95, 0xd3,
	0x67, 0x08, 0xfb, 0xdb, 0x6f, 0xa2, 0xfd, 0x2c, 0xd1, 0x21, 0x30, 0xb9, 0x3b, 0x2a, 0xf7, 0x3d,
	0x2f, 0x0f, 0x3f, 0x2d, 0x27, 0xdb, 0x96, 0x93, 0x5d, 0xcb, 0xc9, 0x77, 0xc7, 0xbd, 0x6d, 0xc7,
	0xbd, 0xdf, 0x8e, 0x7b, 0x1f, 0x57, 0x76, 0xf4, 0x7a, 0xee, 0xb6, 0x6f, 0x2a, 0x54, 0x9f, 0x81,
	0xdd, 0xff, 0xe9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x53, 0x44, 0xbc, 0x6b, 0x90, 0x01, 0x00, 0x00,
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
	OddiyTx(ctx context.Context, in *MsgOddiyTx, opts ...grpc.CallOption) (*MsgOddiyTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) OddiyTx(ctx context.Context, in *MsgOddiyTx, opts ...grpc.CallOption) (*MsgOddiyTxResponse, error) {
	out := new(MsgOddiyTxResponse)
	err := c.cc.Invoke(ctx, "/goan.goan.Msg/OddiyTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	OddiyTx(context.Context, *MsgOddiyTx) (*MsgOddiyTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) OddiyTx(ctx context.Context, req *MsgOddiyTx) (*MsgOddiyTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OddiyTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_OddiyTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOddiyTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OddiyTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goan.goan.Msg/OddiyTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).OddiyTx(ctx, req.(*MsgOddiyTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goan.goan.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OddiyTx",
			Handler:    _Msg_OddiyTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goan/tx.proto",
}

func (m *MsgOddiyTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOddiyTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOddiyTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TxType) > 0 {
		i -= len(m.TxType)
		copy(dAtA[i:], m.TxType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TxType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FeeReceiver) > 0 {
		i -= len(m.FeeReceiver)
		copy(dAtA[i:], m.FeeReceiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FeeReceiver)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
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

func (m *MsgOddiyTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOddiyTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOddiyTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgOddiyTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FeeReceiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TxType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgOddiyTxResponse) Size() (n int) {
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
func (m *MsgOddiyTx) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgOddiyTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOddiyTx: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
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
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeReceiver", wireType)
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
			m.FeeReceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
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
			m.TxType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
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
			m.ServiceName = string(dAtA[iNdEx:postIndex])
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
func (m *MsgOddiyTxResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgOddiyTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOddiyTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
