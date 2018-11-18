// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pending_references.proto

package domain // import "jaytaylor.com/andromeda/domain"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PendingReferences struct {
	PackagePath          string                        `protobuf:"bytes,1,opt,name=package_path,json=packagePath,proto3" json:"package_path,omitempty"`
	ImportedBy           map[string]*PackageReferences `protobuf:"bytes,2,rep,name=imported_by,json=importedBy" json:"imported_by,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *PendingReferences) Reset()         { *m = PendingReferences{} }
func (m *PendingReferences) String() string { return proto.CompactTextString(m) }
func (*PendingReferences) ProtoMessage()    {}
func (*PendingReferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_pending_references_58a71acddbc22875, []int{0}
}
func (m *PendingReferences) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingReferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingReferences.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PendingReferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingReferences.Merge(dst, src)
}
func (m *PendingReferences) XXX_Size() int {
	return m.Size()
}
func (m *PendingReferences) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingReferences.DiscardUnknown(m)
}

var xxx_messageInfo_PendingReferences proto.InternalMessageInfo

func (m *PendingReferences) GetPackagePath() string {
	if m != nil {
		return m.PackagePath
	}
	return ""
}

func (m *PendingReferences) GetImportedBy() map[string]*PackageReferences {
	if m != nil {
		return m.ImportedBy
	}
	return nil
}

func init() {
	proto.RegisterType((*PendingReferences)(nil), "domain.PendingReferences")
	proto.RegisterMapType((map[string]*PackageReferences)(nil), "domain.PendingReferences.ImportedByEntry")
}
func (m *PendingReferences) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingReferences) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PackagePath) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPendingReferences(dAtA, i, uint64(len(m.PackagePath)))
		i += copy(dAtA[i:], m.PackagePath)
	}
	if len(m.ImportedBy) > 0 {
		for k, _ := range m.ImportedBy {
			dAtA[i] = 0x12
			i++
			v := m.ImportedBy[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovPendingReferences(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovPendingReferences(uint64(len(k))) + msgSize
			i = encodeVarintPendingReferences(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintPendingReferences(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintPendingReferences(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPendingReferences(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PendingReferences) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PackagePath)
	if l > 0 {
		n += 1 + l + sovPendingReferences(uint64(l))
	}
	if len(m.ImportedBy) > 0 {
		for k, v := range m.ImportedBy {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovPendingReferences(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovPendingReferences(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovPendingReferences(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPendingReferences(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPendingReferences(x uint64) (n int) {
	return sovPendingReferences(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PendingReferences) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPendingReferences
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PendingReferences: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingReferences: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackagePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingReferences
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPendingReferences
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackagePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImportedBy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingReferences
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPendingReferences
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ImportedBy == nil {
				m.ImportedBy = make(map[string]*PackageReferences)
			}
			var mapkey string
			var mapvalue *PackageReferences
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPendingReferences
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPendingReferences
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPendingReferences
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPendingReferences
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthPendingReferences
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthPendingReferences
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &PackageReferences{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPendingReferences(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthPendingReferences
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ImportedBy[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPendingReferences(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPendingReferences
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
func skipPendingReferences(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPendingReferences
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
					return 0, ErrIntOverflowPendingReferences
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPendingReferences
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPendingReferences
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPendingReferences
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPendingReferences(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPendingReferences = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPendingReferences   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("pending_references.proto", fileDescriptor_pending_references_58a71acddbc22875)
}

var fileDescriptor_pending_references_58a71acddbc22875 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x25, 0x5b, 0xbe, 0xf2, 0x99, 0x55, 0xd4, 0x3d, 0x6d, 0x7b, 0x28, 0x55, 0x10, 0xea, 0xa1,
	0xbb, 0xd0, 0x5e, 0xa4, 0xde, 0x0a, 0x1e, 0xf4, 0x54, 0xf6, 0x24, 0x5e, 0x4a, 0x76, 0x77, 0x9a,
	0xc6, 0x76, 0x33, 0x21, 0x4d, 0xc5, 0xfc, 0x34, 0xfd, 0x05, 0x1e, 0xfd, 0x09, 0x52, 0xfc, 0x21,
	0xd2, 0x64, 0x59, 0x45, 0x4f, 0x99, 0x37, 0xf3, 0xde, 0xbc, 0x97, 0xa1, 0xb1, 0x02, 0x59, 0x0a,
	0xc9, 0xe7, 0x1a, 0x16, 0xa0, 0x41, 0x16, 0xb0, 0x49, 0x94, 0x46, 0x83, 0x51, 0xbb, 0xc4, 0x8a,
	0x09, 0xd9, 0x1d, 0x72, 0x61, 0x96, 0xdb, 0x3c, 0x29, 0xb0, 0x4a, 0x39, 0x72, 0x4c, 0xdd, 0x38,
	0xdf, 0x2e, 0x1c, 0x72, 0xc0, 0x55, 0x5e, 0xd6, 0x9d, 0xfc, 0xa0, 0x0b, 0xb9, 0xc0, 0x7c, 0x8d,
	0xcf, 0xa8, 0x40, 0x7a, 0x59, 0x31, 0xe4, 0x20, 0x87, 0x1c, 0x75, 0x95, 0xa2, 0x32, 0x02, 0xe5,
	0x26, 0xdd, 0x83, 0x5a, 0x7b, 0xa4, 0x58, 0xb1, 0x62, 0x1c, 0x3c, 0x3c, 0xff, 0x24, 0xf4, 0x74,
	0xe6, 0xe3, 0x65, 0x4d, 0xba, 0xe8, 0x8c, 0x1e, 0xd6, 0xb4, 0xb9, 0x62, 0x66, 0x19, 0x93, 0x3e,
	0x19, 0x1c, 0x64, 0x61, 0xdd, 0x9b, 0x31, 0xb3, 0x8c, 0xee, 0x68, 0x28, 0x2a, 0x85, 0xda, 0x40,
	0x39, 0xcf, 0x6d, 0x1c, 0xf4, 0x5b, 0x83, 0x70, 0x74, 0x99, 0xf8, 0x0f, 0x25, 0x7f, 0x56, 0x26,
	0xb7, 0x35, 0x79, 0x6a, 0x6f, 0xa4, 0xd1, 0x36, 0xa3, 0xa2, 0x69, 0x74, 0xef, 0xe9, 0xf1, 0xaf,
	0x71, 0x74, 0x42, 0x5b, 0x2b, 0xb0, 0xb5, 0xf1, 0xbe, 0x8c, 0x52, 0xfa, 0xef, 0x89, 0xad, 0xb7,
	0x10, 0x07, 0x7d, 0x32, 0x08, 0x47, 0x9d, 0xc6, 0xca, 0x87, 0xfa, 0xb6, 0xca, 0x3c, 0x6f, 0x12,
	0x5c, 0x91, 0x49, 0xfb, 0xf5, 0xa5, 0x13, 0xfc, 0x27, 0xd3, 0xf1, 0xdb, 0xae, 0x47, 0xde, 0x77,
	0x3d, 0xf2, 0xb1, 0xeb, 0x91, 0x87, 0x8b, 0x47, 0x66, 0x0d, 0xb3, 0x6b, 0xd4, 0xee, 0x84, 0x4c,
	0x96, 0x1a, 0x2b, 0x28, 0x59, 0xea, 0x57, 0x5e, 0xfb, 0x27, 0x6f, 0xbb, 0x13, 0x8d, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x16, 0xed, 0xef, 0xca, 0xc0, 0x01, 0x00, 0x00,
}
