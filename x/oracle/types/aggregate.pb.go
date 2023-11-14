// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/oracle/aggregate.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Aggregate struct {
	QueryId           string               `protobuf:"bytes,1,opt,name=queryId,proto3" json:"queryId,omitempty"`
	AggregateValue    string               `protobuf:"bytes,2,opt,name=aggregateValue,proto3" json:"aggregateValue,omitempty"`
	AggregateReporter string               `protobuf:"bytes,3,opt,name=aggregateReporter,proto3" json:"aggregateReporter,omitempty"`
	ReporterPower     int64                `protobuf:"varint,4,opt,name=reporterPower,proto3" json:"reporterPower,omitempty"`
	StandardDeviation float64              `protobuf:"fixed64,5,opt,name=standardDeviation,proto3" json:"standardDeviation,omitempty"`
	Reporters         []*AggregateReporter `protobuf:"bytes,6,rep,name=reporters,proto3" json:"reporters,omitempty"`
}

func (m *Aggregate) Reset()         { *m = Aggregate{} }
func (m *Aggregate) String() string { return proto.CompactTextString(m) }
func (*Aggregate) ProtoMessage()    {}
func (*Aggregate) Descriptor() ([]byte, []int) {
	return fileDescriptor_788ad347f505f8a6, []int{0}
}
func (m *Aggregate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Aggregate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Aggregate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Aggregate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aggregate.Merge(m, src)
}
func (m *Aggregate) XXX_Size() int {
	return m.Size()
}
func (m *Aggregate) XXX_DiscardUnknown() {
	xxx_messageInfo_Aggregate.DiscardUnknown(m)
}

var xxx_messageInfo_Aggregate proto.InternalMessageInfo

func (m *Aggregate) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Aggregate) GetAggregateValue() string {
	if m != nil {
		return m.AggregateValue
	}
	return ""
}

func (m *Aggregate) GetAggregateReporter() string {
	if m != nil {
		return m.AggregateReporter
	}
	return ""
}

func (m *Aggregate) GetReporterPower() int64 {
	if m != nil {
		return m.ReporterPower
	}
	return 0
}

func (m *Aggregate) GetStandardDeviation() float64 {
	if m != nil {
		return m.StandardDeviation
	}
	return 0
}

func (m *Aggregate) GetReporters() []*AggregateReporter {
	if m != nil {
		return m.Reporters
	}
	return nil
}

func init() {
	proto.RegisterType((*Aggregate)(nil), "layer.oracle.Aggregate")
}

func init() { proto.RegisterFile("layer/oracle/aggregate.proto", fileDescriptor_788ad347f505f8a6) }

var fileDescriptor_788ad347f505f8a6 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x4f, 0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f,
	0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0xcb, 0xea, 0x41, 0x64, 0xa5,
	0x54, 0xb1, 0xab, 0x8d, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x49, 0x2d, 0x82, 0x68, 0x52, 0xea,
	0x63, 0xe2, 0xe2, 0x74, 0x84, 0x49, 0x0a, 0x49, 0x70, 0xb1, 0x17, 0x96, 0xa6, 0x16, 0x55, 0x7a,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x6a, 0x5c, 0x7c, 0x70, 0x33,
	0xc2, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x0a, 0xd0, 0x44, 0x85, 0x74, 0xb8, 0x04, 0xe1,
	0x22, 0x41, 0x50, 0xab, 0x24, 0x98, 0xc1, 0x4a, 0x31, 0x25, 0x84, 0x54, 0xb8, 0x78, 0x61, 0xee,
	0x09, 0xc8, 0x2f, 0x4f, 0x2d, 0x92, 0x60, 0x51, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x15, 0x04, 0x99,
	0x59, 0x5c, 0x92, 0x98, 0x97, 0x92, 0x58, 0x94, 0xe2, 0x92, 0x5a, 0x96, 0x99, 0x58, 0x92, 0x99,
	0x9f, 0x27, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x18, 0x84, 0x29, 0x21, 0x64, 0xcb, 0xc5, 0x09, 0xd3,
	0x5e, 0x2c, 0xc1, 0xa6, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xaf, 0x87, 0x1c, 0x34, 0x7a, 0x8e, 0xe8,
	0xee, 0x08, 0x42, 0xe8, 0x70, 0x72, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x92, 0xd4, 0x9c,
	0x9c, 0xfc, 0x22, 0xdd, 0xcc, 0x7c, 0x7d, 0x48, 0x30, 0x57, 0xc0, 0x02, 0xba, 0xa4, 0xb2, 0x20,
	0xb5, 0x38, 0x89, 0x0d, 0x1c, 0xb8, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x53, 0x19,
	0xdd, 0xb1, 0x01, 0x00, 0x00,
}

func (m *Aggregate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Aggregate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Aggregate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reporters) > 0 {
		for iNdEx := len(m.Reporters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reporters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAggregate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.StandardDeviation != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.StandardDeviation))))
		i--
		dAtA[i] = 0x29
	}
	if m.ReporterPower != 0 {
		i = encodeVarintAggregate(dAtA, i, uint64(m.ReporterPower))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AggregateReporter) > 0 {
		i -= len(m.AggregateReporter)
		copy(dAtA[i:], m.AggregateReporter)
		i = encodeVarintAggregate(dAtA, i, uint64(len(m.AggregateReporter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AggregateValue) > 0 {
		i -= len(m.AggregateValue)
		copy(dAtA[i:], m.AggregateValue)
		i = encodeVarintAggregate(dAtA, i, uint64(len(m.AggregateValue)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintAggregate(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAggregate(dAtA []byte, offset int, v uint64) int {
	offset -= sovAggregate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Aggregate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovAggregate(uint64(l))
	}
	l = len(m.AggregateValue)
	if l > 0 {
		n += 1 + l + sovAggregate(uint64(l))
	}
	l = len(m.AggregateReporter)
	if l > 0 {
		n += 1 + l + sovAggregate(uint64(l))
	}
	if m.ReporterPower != 0 {
		n += 1 + sovAggregate(uint64(m.ReporterPower))
	}
	if m.StandardDeviation != 0 {
		n += 9
	}
	if len(m.Reporters) > 0 {
		for _, e := range m.Reporters {
			l = e.Size()
			n += 1 + l + sovAggregate(uint64(l))
		}
	}
	return n
}

func sovAggregate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAggregate(x uint64) (n int) {
	return sovAggregate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Aggregate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAggregate
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
			return fmt.Errorf("proto: Aggregate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Aggregate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateReporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateReporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReporterPower", wireType)
			}
			m.ReporterPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReporterPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardDeviation", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.StandardDeviation = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregate
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
				return ErrInvalidLengthAggregate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAggregate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporters = append(m.Reporters, &AggregateReporter{})
			if err := m.Reporters[len(m.Reporters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAggregate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAggregate
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
func skipAggregate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAggregate
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
					return 0, ErrIntOverflowAggregate
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
					return 0, ErrIntOverflowAggregate
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
				return 0, ErrInvalidLengthAggregate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAggregate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAggregate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAggregate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAggregate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAggregate = fmt.Errorf("proto: unexpected end of group")
)