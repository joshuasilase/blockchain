// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dao.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type ClusterHistoricalRewards struct {
	CumulativeRewardRatio github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=cumulative_reward_ratio,json=cumulativeRewardRatio,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"cumulative_reward_ratio"`
	ReferenceCount        uint32                                      `protobuf:"varint,2,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	HisReward             github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,3,rep,name=his_reward,json=hisReward,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"his_reward"`
}

func (m *ClusterHistoricalRewards) Reset()         { *m = ClusterHistoricalRewards{} }
func (m *ClusterHistoricalRewards) String() string { return proto.CompactTextString(m) }
func (*ClusterHistoricalRewards) ProtoMessage()    {}
func (*ClusterHistoricalRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a2fcf5e349a9ce, []int{0}
}
func (m *ClusterHistoricalRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterHistoricalRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterHistoricalRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterHistoricalRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHistoricalRewards.Merge(m, src)
}
func (m *ClusterHistoricalRewards) XXX_Size() int {
	return m.Size()
}
func (m *ClusterHistoricalRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHistoricalRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHistoricalRewards proto.InternalMessageInfo

func (m *ClusterHistoricalRewards) GetCumulativeRewardRatio() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.CumulativeRewardRatio
	}
	return nil
}

func (m *ClusterHistoricalRewards) GetReferenceCount() uint32 {
	if m != nil {
		return m.ReferenceCount
	}
	return 0
}

func (m *ClusterHistoricalRewards) GetHisReward() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.HisReward
	}
	return nil
}

type ClusterCurrentRewards struct {
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
	Period  uint64                                      `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
}

func (m *ClusterCurrentRewards) Reset()         { *m = ClusterCurrentRewards{} }
func (m *ClusterCurrentRewards) String() string { return proto.CompactTextString(m) }
func (*ClusterCurrentRewards) ProtoMessage()    {}
func (*ClusterCurrentRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a2fcf5e349a9ce, []int{1}
}
func (m *ClusterCurrentRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterCurrentRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterCurrentRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterCurrentRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterCurrentRewards.Merge(m, src)
}
func (m *ClusterCurrentRewards) XXX_Size() int {
	return m.Size()
}
func (m *ClusterCurrentRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterCurrentRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterCurrentRewards proto.InternalMessageInfo

func (m *ClusterCurrentRewards) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *ClusterCurrentRewards) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type ClusterOutstandingRewards struct {
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
}

func (m *ClusterOutstandingRewards) Reset()         { *m = ClusterOutstandingRewards{} }
func (m *ClusterOutstandingRewards) String() string { return proto.CompactTextString(m) }
func (*ClusterOutstandingRewards) ProtoMessage()    {}
func (*ClusterOutstandingRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a2fcf5e349a9ce, []int{2}
}
func (m *ClusterOutstandingRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterOutstandingRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterOutstandingRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterOutstandingRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterOutstandingRewards.Merge(m, src)
}
func (m *ClusterOutstandingRewards) XXX_Size() int {
	return m.Size()
}
func (m *ClusterOutstandingRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterOutstandingRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterOutstandingRewards proto.InternalMessageInfo

func (m *ClusterOutstandingRewards) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

type BurnStartingInfo struct {
	PreviousPeriod uint64                                 `protobuf:"varint,1,opt,name=previous_period,json=previousPeriod,proto3" json:"previous_period,omitempty"`
	Stake          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=stake,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stake"`
	Height         uint64                                 `protobuf:"varint,3,opt,name=height,proto3" json:"creation_height"`
}

func (m *BurnStartingInfo) Reset()         { *m = BurnStartingInfo{} }
func (m *BurnStartingInfo) String() string { return proto.CompactTextString(m) }
func (*BurnStartingInfo) ProtoMessage()    {}
func (*BurnStartingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a2fcf5e349a9ce, []int{3}
}
func (m *BurnStartingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnStartingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnStartingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnStartingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnStartingInfo.Merge(m, src)
}
func (m *BurnStartingInfo) XXX_Size() int {
	return m.Size()
}
func (m *BurnStartingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnStartingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BurnStartingInfo proto.InternalMessageInfo

func (m *BurnStartingInfo) GetPreviousPeriod() uint64 {
	if m != nil {
		return m.PreviousPeriod
	}
	return 0
}

func (m *BurnStartingInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterHistoricalRewards)(nil), "freemasonry.dao.v1.ClusterHistoricalRewards")
	proto.RegisterType((*ClusterCurrentRewards)(nil), "freemasonry.dao.v1.ClusterCurrentRewards")
	proto.RegisterType((*ClusterOutstandingRewards)(nil), "freemasonry.dao.v1.ClusterOutstandingRewards")
	proto.RegisterType((*BurnStartingInfo)(nil), "freemasonry.dao.v1.BurnStartingInfo")
}

func init() { proto.RegisterFile("dao.proto", fileDescriptor_47a2fcf5e349a9ce) }

var fileDescriptor_47a2fcf5e349a9ce = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x31, 0x6f, 0xd3, 0x4e,
	0x18, 0xc6, 0x73, 0xcd, 0xff, 0x1f, 0x94, 0x43, 0xb4, 0xc8, 0x50, 0x70, 0x2b, 0xe4, 0x54, 0x91,
	0x80, 0x4a, 0x55, 0x6d, 0x85, 0xae, 0x48, 0x48, 0x4e, 0x07, 0x98, 0x40, 0x66, 0x63, 0xb1, 0xce,
	0xe7, 0x37, 0xf6, 0x29, 0xc9, 0xbd, 0xd6, 0xdd, 0x39, 0xd0, 0x6f, 0x00, 0x1b, 0x1f, 0x80, 0x4f,
	0xc0, 0xcc, 0xc8, 0xc2, 0xd6, 0xb1, 0x62, 0x42, 0x0c, 0x01, 0x25, 0x1b, 0x9f, 0x02, 0xd9, 0x77,
	0x01, 0x46, 0x96, 0x8a, 0xc9, 0x77, 0xcf, 0xeb, 0xf7, 0x9e, 0xdf, 0xfb, 0xd8, 0x47, 0xfb, 0x39,
	0xc3, 0xb0, 0x52, 0x68, 0xd0, 0xf3, 0x26, 0x0a, 0x60, 0xce, 0x34, 0x4a, 0x75, 0x16, 0x36, 0xf2,
	0x62, 0xb4, 0x1f, 0x70, 0xd4, 0x73, 0xd4, 0x51, 0xc6, 0x34, 0x44, 0x8b, 0x51, 0x06, 0x86, 0x8d,
	0x22, 0x8e, 0x42, 0xda, 0x9e, 0xfd, 0x9b, 0x05, 0x16, 0xd8, 0x2e, 0xa3, 0x66, 0xe5, 0xd4, 0x3d,
	0xdb, 0x95, 0xda, 0x82, 0xdd, 0xd8, 0xd2, 0xf0, 0xd3, 0x16, 0xf5, 0xc7, 0xb3, 0x5a, 0x1b, 0x50,
	0x8f, 0x85, 0x36, 0xa8, 0x04, 0x67, 0xb3, 0x04, 0x5e, 0x32, 0x95, 0x6b, 0xef, 0x0d, 0xa1, 0xb7,
	0x79, 0x3d, 0xaf, 0x67, 0xcc, 0x88, 0x05, 0xa4, 0xaa, 0x95, 0x53, 0xc5, 0x8c, 0x40, 0x9f, 0x1c,
	0x74, 0x0f, 0xaf, 0x3e, 0xb8, 0x13, 0xba, 0xd3, 0x1a, 0xa0, 0xd0, 0x01, 0x85, 0xa7, 0xc0, 0xc7,
	0x28, 0x64, 0x7c, 0x72, 0xbe, 0x1c, 0x74, 0xde, 0x7f, 0x1b, 0x1c, 0x15, 0xc2, 0x94, 0x75, 0x16,
	0x72, 0x9c, 0x3b, 0x77, 0xf7, 0x38, 0xd6, 0xf9, 0x34, 0x32, 0x67, 0x15, 0xe8, 0x4d, 0x8f, 0x4e,
	0x76, 0x7f, 0x3b, 0x5a, 0x8e, 0xa4, 0xf1, 0xf3, 0xee, 0xd3, 0x1d, 0x05, 0x13, 0x50, 0x20, 0x39,
	0xa4, 0x1c, 0x6b, 0x69, 0xfc, 0xad, 0x03, 0x72, 0x78, 0x2d, 0xd9, 0xfe, 0x25, 0x8f, 0x1b, 0xd5,
	0xab, 0x28, 0x2d, 0x85, 0x76, 0xb0, 0x7e, 0xf7, 0xb2, 0x30, 0xfb, 0xa5, 0xd0, 0x96, 0x6f, 0xf8,
	0x8e, 0xd0, 0x5d, 0x97, 0xe1, 0xb8, 0x56, 0x0a, 0xa4, 0xd9, 0x04, 0x38, 0xa5, 0x57, 0x2c, 0x87,
	0xbe, 0xbc, 0xbc, 0x36, 0x0e, 0xde, 0x2d, 0xda, 0xab, 0x40, 0x09, 0xcc, 0xdb, 0x60, 0xfe, 0x4b,
	0xdc, 0x6e, 0xf8, 0x9a, 0xd0, 0x3d, 0x87, 0xf7, 0xb4, 0x36, 0xda, 0x30, 0x99, 0x0b, 0x59, 0xfc,
	0x0b, 0xc4, 0xe1, 0x47, 0x42, 0xaf, 0xc7, 0xb5, 0x92, 0xcf, 0x0d, 0x53, 0x46, 0xc8, 0xe2, 0x89,
	0x9c, 0xb4, 0x5f, 0xb6, 0x52, 0xb0, 0x10, 0x58, 0xeb, 0xd4, 0x0d, 0x40, 0xda, 0x01, 0xb6, 0x37,
	0xf2, 0xb3, 0x56, 0xf5, 0x12, 0xfa, 0xbf, 0x36, 0x6c, 0x0a, 0xed, 0x7c, 0xfd, 0xf8, 0x61, 0x83,
	0xf2, 0x75, 0x39, 0xb8, 0xf7, 0x77, 0x28, 0x9f, 0x3f, 0x1c, 0x53, 0x37, 0xd9, 0x29, 0xf0, 0xc4,
	0x1e, 0xe5, 0x1d, 0xd1, 0x5e, 0x09, 0xa2, 0x28, 0x8d, 0xdf, 0x6d, 0x3c, 0xe3, 0x1b, 0x3f, 0x96,
	0x83, 0x1d, 0xae, 0xa0, 0xf9, 0xe7, 0x64, 0x6a, 0x4b, 0x89, 0x7b, 0x25, 0x7e, 0x74, 0xbe, 0x0a,
	0xc8, 0xc5, 0x2a, 0x20, 0xdf, 0x57, 0x01, 0x79, 0xbb, 0x0e, 0x3a, 0x17, 0xeb, 0xa0, 0xf3, 0x65,
	0x1d, 0x74, 0x5e, 0xdc, 0xfd, 0xf3, 0xae, 0x72, 0x1e, 0x65, 0x33, 0xe4, 0x53, 0x5e, 0x32, 0x21,
	0xa3, 0x57, 0x51, 0xce, 0xd0, 0x62, 0x64, 0xbd, 0xf6, 0xd2, 0x9d, 0xfc, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x5b, 0xfb, 0x56, 0xe6, 0x03, 0x00, 0x00,
}

func (m *ClusterHistoricalRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterHistoricalRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterHistoricalRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HisReward) > 0 {
		for iNdEx := len(m.HisReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HisReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ReferenceCount != 0 {
		i = encodeVarintDao(dAtA, i, uint64(m.ReferenceCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CumulativeRewardRatio) > 0 {
		for iNdEx := len(m.CumulativeRewardRatio) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CumulativeRewardRatio[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterCurrentRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterCurrentRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterCurrentRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Period != 0 {
		i = encodeVarintDao(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterOutstandingRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterOutstandingRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterOutstandingRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BurnStartingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnStartingInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnStartingInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintDao(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Stake.Size()
		i -= size
		if _, err := m.Stake.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDao(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PreviousPeriod != 0 {
		i = encodeVarintDao(dAtA, i, uint64(m.PreviousPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDao(dAtA []byte, offset int, v uint64) int {
	offset -= sovDao(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterHistoricalRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CumulativeRewardRatio) > 0 {
		for _, e := range m.CumulativeRewardRatio {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	if m.ReferenceCount != 0 {
		n += 1 + sovDao(uint64(m.ReferenceCount))
	}
	if len(m.HisReward) > 0 {
		for _, e := range m.HisReward {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func (m *ClusterCurrentRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	if m.Period != 0 {
		n += 1 + sovDao(uint64(m.Period))
	}
	return n
}

func (m *ClusterOutstandingRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func (m *BurnStartingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PreviousPeriod != 0 {
		n += 1 + sovDao(uint64(m.PreviousPeriod))
	}
	l = m.Stake.Size()
	n += 1 + l + sovDao(uint64(l))
	if m.Height != 0 {
		n += 1 + sovDao(uint64(m.Height))
	}
	return n
}

func sovDao(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDao(x uint64) (n int) {
	return sovDao(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClusterHistoricalRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: ClusterHistoricalRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterHistoricalRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeRewardRatio", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CumulativeRewardRatio = append(m.CumulativeRewardRatio, types.DecCoin{})
			if err := m.CumulativeRewardRatio[len(m.CumulativeRewardRatio)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceCount", wireType)
			}
			m.ReferenceCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReferenceCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HisReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HisReward = append(m.HisReward, types.DecCoin{})
			if err := m.HisReward[len(m.HisReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *ClusterCurrentRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: ClusterCurrentRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterCurrentRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *ClusterOutstandingRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: ClusterOutstandingRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterOutstandingRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *BurnStartingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: BurnStartingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnStartingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousPeriod", wireType)
			}
			m.PreviousPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreviousPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func skipDao(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDao
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
					return 0, ErrIntOverflowDao
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
					return 0, ErrIntOverflowDao
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
				return 0, ErrInvalidLengthDao
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDao
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDao
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDao        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDao          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDao = fmt.Errorf("proto: unexpected end of group")
)