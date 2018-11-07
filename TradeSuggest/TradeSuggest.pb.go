// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TradeSuggest.proto

package TradeSuggest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TradeSuggestRequest struct {
	Resolution           int32    `protobuf:"varint,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeSuggestRequest) Reset()         { *m = TradeSuggestRequest{} }
func (m *TradeSuggestRequest) String() string { return proto.CompactTextString(m) }
func (*TradeSuggestRequest) ProtoMessage()    {}
func (*TradeSuggestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_TradeSuggest_2642414f3c54c537, []int{0}
}
func (m *TradeSuggestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeSuggestRequest.Unmarshal(m, b)
}
func (m *TradeSuggestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeSuggestRequest.Marshal(b, m, deterministic)
}
func (dst *TradeSuggestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeSuggestRequest.Merge(dst, src)
}
func (m *TradeSuggestRequest) XXX_Size() int {
	return xxx_messageInfo_TradeSuggestRequest.Size(m)
}
func (m *TradeSuggestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeSuggestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeSuggestRequest proto.InternalMessageInfo

func (m *TradeSuggestRequest) GetResolution() int32 {
	if m != nil {
		return m.Resolution
	}
	return 0
}

type Suggestion struct {
	Stockname            string   `protobuf:"bytes,1,opt,name=stockname,proto3" json:"stockname,omitempty"`
	Rating               int32    `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Direction            int32    `protobuf:"varint,3,opt,name=direction,proto3" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Suggestion) Reset()         { *m = Suggestion{} }
func (m *Suggestion) String() string { return proto.CompactTextString(m) }
func (*Suggestion) ProtoMessage()    {}
func (*Suggestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_TradeSuggest_2642414f3c54c537, []int{1}
}
func (m *Suggestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suggestion.Unmarshal(m, b)
}
func (m *Suggestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suggestion.Marshal(b, m, deterministic)
}
func (dst *Suggestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suggestion.Merge(dst, src)
}
func (m *Suggestion) XXX_Size() int {
	return xxx_messageInfo_Suggestion.Size(m)
}
func (m *Suggestion) XXX_DiscardUnknown() {
	xxx_messageInfo_Suggestion.DiscardUnknown(m)
}

var xxx_messageInfo_Suggestion proto.InternalMessageInfo

func (m *Suggestion) GetStockname() string {
	if m != nil {
		return m.Stockname
	}
	return ""
}

func (m *Suggestion) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Suggestion) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type TradeSuggestResponse struct {
	Suggestions          []*Suggestion `protobuf:"bytes,1,rep,name=Suggestions,proto3" json:"Suggestions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TradeSuggestResponse) Reset()         { *m = TradeSuggestResponse{} }
func (m *TradeSuggestResponse) String() string { return proto.CompactTextString(m) }
func (*TradeSuggestResponse) ProtoMessage()    {}
func (*TradeSuggestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_TradeSuggest_2642414f3c54c537, []int{2}
}
func (m *TradeSuggestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeSuggestResponse.Unmarshal(m, b)
}
func (m *TradeSuggestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeSuggestResponse.Marshal(b, m, deterministic)
}
func (dst *TradeSuggestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeSuggestResponse.Merge(dst, src)
}
func (m *TradeSuggestResponse) XXX_Size() int {
	return xxx_messageInfo_TradeSuggestResponse.Size(m)
}
func (m *TradeSuggestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeSuggestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeSuggestResponse proto.InternalMessageInfo

func (m *TradeSuggestResponse) GetSuggestions() []*Suggestion {
	if m != nil {
		return m.Suggestions
	}
	return nil
}

func init() {
	proto.RegisterType((*TradeSuggestRequest)(nil), "TradeSuggestRequest")
	proto.RegisterType((*Suggestion)(nil), "Suggestion")
	proto.RegisterType((*TradeSuggestResponse)(nil), "TradeSuggestResponse")
}

func init() { proto.RegisterFile("TradeSuggest.proto", fileDescriptor_TradeSuggest_2642414f3c54c537) }

var fileDescriptor_TradeSuggest_2642414f3c54c537 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0a, 0x29, 0x4a, 0x4c,
	0x49, 0x0d, 0x2e, 0x4d, 0x4f, 0x4f, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32,
	0xe5, 0x12, 0x46, 0x16, 0x0d, 0x4a, 0x2d, 0x2c, 0x05, 0x52, 0x42, 0x72, 0x5c, 0x5c, 0x45, 0xa9,
	0xc5, 0xf9, 0x39, 0xa5, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x48,
	0x22, 0x4a, 0x09, 0x5c, 0x5c, 0x50, 0x1d, 0x40, 0x9e, 0x90, 0x0c, 0x17, 0x67, 0x71, 0x49, 0x7e,
	0x72, 0x76, 0x5e, 0x62, 0x6e, 0x2a, 0x58, 0x31, 0x67, 0x10, 0x42, 0x40, 0x48, 0x8c, 0x8b, 0xad,
	0x28, 0xb1, 0x24, 0x33, 0x2f, 0x5d, 0x82, 0x09, 0x6c, 0x0e, 0x94, 0x07, 0xd2, 0x95, 0x92, 0x59,
	0x94, 0x9a, 0x0c, 0xb6, 0x82, 0x19, 0x2c, 0x85, 0x10, 0x50, 0x72, 0xe5, 0x12, 0x41, 0x75, 0x58,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x2e, 0x17, 0x37, 0xc2, 0xe6, 0x62, 0xa0, 0x6d, 0xcc,
	0x1a, 0xdc, 0x46, 0xdc, 0x7a, 0x08, 0xb1, 0x20, 0x64, 0x79, 0xa3, 0x30, 0x54, 0xff, 0x05, 0xa7,
	0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xd9, 0x73, 0xf1, 0xb9, 0xa7, 0x96, 0x20, 0x29, 0x14, 0x12,
	0xd1, 0xc3, 0x12, 0x0e, 0x52, 0xa2, 0x7a, 0xd8, 0x1c, 0xa1, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x3e,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xe5, 0x0e, 0x29, 0x54, 0x01, 0x00, 0x00,
}