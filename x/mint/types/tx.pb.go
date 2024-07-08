// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/mint/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type MsgInit struct {
	// authority is the address that is allowed calling this msg.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgInit) Reset()         { *m = MsgInit{} }
func (m *MsgInit) String() string { return proto.CompactTextString(m) }
func (*MsgInit) ProtoMessage()    {}
func (*MsgInit) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbaaf682b6a9a290, []int{0}
}
func (m *MsgInit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInit.Merge(m, src)
}
func (m *MsgInit) XXX_Size() int {
	return m.Size()
}
func (m *MsgInit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInit proto.InternalMessageInfo

func (m *MsgInit) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type MsgMsgInitResponse struct {
}

func (m *MsgMsgInitResponse) Reset()         { *m = MsgMsgInitResponse{} }
func (m *MsgMsgInitResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMsgInitResponse) ProtoMessage()    {}
func (*MsgMsgInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbaaf682b6a9a290, []int{1}
}
func (m *MsgMsgInitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMsgInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMsgInitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMsgInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMsgInitResponse.Merge(m, src)
}
func (m *MsgMsgInitResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMsgInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMsgInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMsgInitResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgInit)(nil), "layer.mint.MsgInit")
	proto.RegisterType((*MsgMsgInitResponse)(nil), "layer.mint.MsgMsgInitResponse")
}

func init() { proto.RegisterFile("layer/mint/tx.proto", fileDescriptor_bbaaf682b6a9a290) }

var fileDescriptor_bbaaf682b6a9a290 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0x69, 0xfc, 0x15, 0xba, 0x79, 0x5c, 0x22, 0xde, 0xd0, 0x18, 0x16, 0x0c, 0x09, 0xd7,
	0xa0, 0x89, 0x83, 0x4e, 0xb0, 0x31, 0xb0, 0xe0, 0xe6, 0x62, 0xee, 0xa0, 0x29, 0x35, 0x5c, 0x1f,
	0xe9, 0x2b, 0x04, 0x36, 0xe3, 0xe8, 0xe4, 0x9f, 0xc2, 0xe0, 0x1f, 0xe1, 0x48, 0x9c, 0x1c, 0x0d,
	0x0c, 0xfc, 0x1b, 0xe6, 0xae, 0x4d, 0x2e, 0xb2, 0xb4, 0xfd, 0x7e, 0x3f, 0x2f, 0xaf, 0xef, 0x7d,
	0x69, 0x6d, 0x9a, 0xac, 0x84, 0xe1, 0x99, 0xd2, 0x96, 0xdb, 0x65, 0x3c, 0x33, 0x60, 0x21, 0xa0,
	0x85, 0x19, 0xe7, 0x66, 0x74, 0x9e, 0x64, 0x4a, 0x03, 0x2f, 0x4e, 0x87, 0x23, 0x36, 0x02, 0xcc,
	0x00, 0x79, 0x9a, 0xa0, 0xe0, 0x8b, 0x4e, 0x2a, 0x6c, 0xd2, 0xe1, 0x23, 0x50, 0xda, 0xf3, 0x0b,
	0xcf, 0x33, 0x94, 0x7c, 0xd1, 0xc9, 0x2f, 0x0f, 0x2e, 0x1d, 0x78, 0x2e, 0x14, 0x77, 0xc2, 0xa3,
	0x50, 0x82, 0x04, 0xe7, 0xe7, 0x2f, 0xe7, 0x36, 0x5e, 0xe8, 0xd9, 0x00, 0x65, 0x5f, 0x2b, 0x1b,
	0xdc, 0xd1, 0x6a, 0x32, 0xb7, 0x13, 0x30, 0xca, 0xae, 0xea, 0xe4, 0x8a, 0x5c, 0x57, 0x7b, 0xf5,
	0xef, 0xcf, 0x76, 0xe8, 0xbb, 0x74, 0xc7, 0x63, 0x23, 0x10, 0x1f, 0xad, 0x51, 0x5a, 0x0e, 0xcb,
	0xd2, 0xfb, 0xe6, 0xdb, 0x7e, 0xdd, 0x2a, 0xf5, 0xfb, 0x7e, 0xdd, 0x0a, 0xdd, 0xce, 0x4b, 0xb7,
	0xb5, 0xff, 0xa0, 0x11, 0xd2, 0x60, 0x80, 0xd2, 0xab, 0xa1, 0xc0, 0x19, 0x68, 0x14, 0x37, 0x7d,
	0x7a, 0x34, 0x40, 0x19, 0x3c, 0xd0, 0xe3, 0x62, 0x8a, 0x5a, 0x5c, 0x46, 0x13, 0xfb, 0xda, 0x88,
	0x1d, 0x98, 0x07, 0x3d, 0xa2, 0x93, 0xd7, 0xfd, 0xba, 0x45, 0x7a, 0xdd, 0xaf, 0x2d, 0x23, 0x9b,
	0x2d, 0x23, 0xbf, 0x5b, 0x46, 0x3e, 0x76, 0xac, 0xb2, 0xd9, 0xb1, 0xca, 0xcf, 0x8e, 0x55, 0x9e,
	0x9a, 0x52, 0xd9, 0xc9, 0x3c, 0x8d, 0x47, 0x90, 0x71, 0x2b, 0xa6, 0x53, 0x30, 0x6d, 0x05, 0xfc,
	0xdf, 0x94, 0x76, 0x35, 0x13, 0x98, 0x9e, 0x16, 0xb1, 0xdc, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xa8, 0x7a, 0x3d, 0x2e, 0xb6, 0x01, 0x00, 0x00,
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
	Init(ctx context.Context, in *MsgInit, opts ...grpc.CallOption) (*MsgMsgInitResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Init(ctx context.Context, in *MsgInit, opts ...grpc.CallOption) (*MsgMsgInitResponse, error) {
	out := new(MsgMsgInitResponse)
	err := c.cc.Invoke(ctx, "/layer.mint.Msg/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Init(context.Context, *MsgInit) (*MsgMsgInitResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Init(ctx context.Context, req *MsgInit) (*MsgMsgInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.mint.Msg/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Init(ctx, req.(*MsgInit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "layer.mint.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Msg_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layer/mint/tx.proto",
}

func (m *MsgInit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMsgInitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMsgInitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMsgInitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgInit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgMsgInitResponse) Size() (n int) {
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
func (m *MsgInit) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgInit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
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
func (m *MsgMsgInitResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMsgInitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMsgInitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
