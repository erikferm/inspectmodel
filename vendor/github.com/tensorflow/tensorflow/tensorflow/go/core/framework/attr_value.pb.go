// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attr_value.proto

package framework

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Protocol buffer representing the value for an attr used to configure an Op.
// Comment indicates the corresponding attr type.  Only the field matching the
// attr type may be filled.
type AttrValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttrValue_S
	//	*AttrValue_I
	//	*AttrValue_F
	//	*AttrValue_B
	//	*AttrValue_Type
	//	*AttrValue_Shape
	//	*AttrValue_Tensor
	//	*AttrValue_List
	//	*AttrValue_Func
	//	*AttrValue_Placeholder
	Value                isAttrValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttrValue) Reset()         { *m = AttrValue{} }
func (m *AttrValue) String() string { return proto.CompactTextString(m) }
func (*AttrValue) ProtoMessage()    {}
func (*AttrValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a312f2b3a701292, []int{0}
}

func (m *AttrValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrValue.Unmarshal(m, b)
}
func (m *AttrValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrValue.Marshal(b, m, deterministic)
}
func (m *AttrValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrValue.Merge(m, src)
}
func (m *AttrValue) XXX_Size() int {
	return xxx_messageInfo_AttrValue.Size(m)
}
func (m *AttrValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttrValue proto.InternalMessageInfo

type isAttrValue_Value interface {
	isAttrValue_Value()
}

type AttrValue_S struct {
	S []byte `protobuf:"bytes,2,opt,name=s,proto3,oneof"`
}

type AttrValue_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,proto3,oneof"`
}

type AttrValue_F struct {
	F float32 `protobuf:"fixed32,4,opt,name=f,proto3,oneof"`
}

type AttrValue_B struct {
	B bool `protobuf:"varint,5,opt,name=b,proto3,oneof"`
}

type AttrValue_Type struct {
	Type DataType `protobuf:"varint,6,opt,name=type,proto3,enum=tensorflow.DataType,oneof"`
}

type AttrValue_Shape struct {
	Shape *TensorShapeProto `protobuf:"bytes,7,opt,name=shape,proto3,oneof"`
}

type AttrValue_Tensor struct {
	Tensor *TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"`
}

type AttrValue_List struct {
	List *AttrValue_ListValue `protobuf:"bytes,1,opt,name=list,proto3,oneof"`
}

type AttrValue_Func struct {
	Func *NameAttrList `protobuf:"bytes,10,opt,name=func,proto3,oneof"`
}

type AttrValue_Placeholder struct {
	Placeholder string `protobuf:"bytes,9,opt,name=placeholder,proto3,oneof"`
}

func (*AttrValue_S) isAttrValue_Value() {}

func (*AttrValue_I) isAttrValue_Value() {}

func (*AttrValue_F) isAttrValue_Value() {}

func (*AttrValue_B) isAttrValue_Value() {}

func (*AttrValue_Type) isAttrValue_Value() {}

func (*AttrValue_Shape) isAttrValue_Value() {}

func (*AttrValue_Tensor) isAttrValue_Value() {}

func (*AttrValue_List) isAttrValue_Value() {}

func (*AttrValue_Func) isAttrValue_Value() {}

func (*AttrValue_Placeholder) isAttrValue_Value() {}

func (m *AttrValue) GetValue() isAttrValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttrValue) GetS() []byte {
	if x, ok := m.GetValue().(*AttrValue_S); ok {
		return x.S
	}
	return nil
}

func (m *AttrValue) GetI() int64 {
	if x, ok := m.GetValue().(*AttrValue_I); ok {
		return x.I
	}
	return 0
}

func (m *AttrValue) GetF() float32 {
	if x, ok := m.GetValue().(*AttrValue_F); ok {
		return x.F
	}
	return 0
}

func (m *AttrValue) GetB() bool {
	if x, ok := m.GetValue().(*AttrValue_B); ok {
		return x.B
	}
	return false
}

func (m *AttrValue) GetType() DataType {
	if x, ok := m.GetValue().(*AttrValue_Type); ok {
		return x.Type
	}
	return DataType_DT_INVALID
}

func (m *AttrValue) GetShape() *TensorShapeProto {
	if x, ok := m.GetValue().(*AttrValue_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *AttrValue) GetTensor() *TensorProto {
	if x, ok := m.GetValue().(*AttrValue_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (m *AttrValue) GetList() *AttrValue_ListValue {
	if x, ok := m.GetValue().(*AttrValue_List); ok {
		return x.List
	}
	return nil
}

func (m *AttrValue) GetFunc() *NameAttrList {
	if x, ok := m.GetValue().(*AttrValue_Func); ok {
		return x.Func
	}
	return nil
}

func (m *AttrValue) GetPlaceholder() string {
	if x, ok := m.GetValue().(*AttrValue_Placeholder); ok {
		return x.Placeholder
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AttrValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AttrValue_S)(nil),
		(*AttrValue_I)(nil),
		(*AttrValue_F)(nil),
		(*AttrValue_B)(nil),
		(*AttrValue_Type)(nil),
		(*AttrValue_Shape)(nil),
		(*AttrValue_Tensor)(nil),
		(*AttrValue_List)(nil),
		(*AttrValue_Func)(nil),
		(*AttrValue_Placeholder)(nil),
	}
}

// LINT.IfChange
type AttrValue_ListValue struct {
	S                    [][]byte            `protobuf:"bytes,2,rep,name=s,proto3" json:"s,omitempty"`
	I                    []int64             `protobuf:"varint,3,rep,packed,name=i,proto3" json:"i,omitempty"`
	F                    []float32           `protobuf:"fixed32,4,rep,packed,name=f,proto3" json:"f,omitempty"`
	B                    []bool              `protobuf:"varint,5,rep,packed,name=b,proto3" json:"b,omitempty"`
	Type                 []DataType          `protobuf:"varint,6,rep,packed,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	Shape                []*TensorShapeProto `protobuf:"bytes,7,rep,name=shape,proto3" json:"shape,omitempty"`
	Tensor               []*TensorProto      `protobuf:"bytes,8,rep,name=tensor,proto3" json:"tensor,omitempty"`
	Func                 []*NameAttrList     `protobuf:"bytes,9,rep,name=func,proto3" json:"func,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AttrValue_ListValue) Reset()         { *m = AttrValue_ListValue{} }
func (m *AttrValue_ListValue) String() string { return proto.CompactTextString(m) }
func (*AttrValue_ListValue) ProtoMessage()    {}
func (*AttrValue_ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a312f2b3a701292, []int{0, 0}
}

func (m *AttrValue_ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrValue_ListValue.Unmarshal(m, b)
}
func (m *AttrValue_ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrValue_ListValue.Marshal(b, m, deterministic)
}
func (m *AttrValue_ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrValue_ListValue.Merge(m, src)
}
func (m *AttrValue_ListValue) XXX_Size() int {
	return xxx_messageInfo_AttrValue_ListValue.Size(m)
}
func (m *AttrValue_ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrValue_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttrValue_ListValue proto.InternalMessageInfo

func (m *AttrValue_ListValue) GetS() [][]byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *AttrValue_ListValue) GetI() []int64 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *AttrValue_ListValue) GetF() []float32 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *AttrValue_ListValue) GetB() []bool {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *AttrValue_ListValue) GetType() []DataType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AttrValue_ListValue) GetShape() []*TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *AttrValue_ListValue) GetTensor() []*TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *AttrValue_ListValue) GetFunc() []*NameAttrList {
	if m != nil {
		return m.Func
	}
	return nil
}

// A list of attr names and their values. The whole list is attached
// with a string name.  E.g., MatMul[T=float].
type NameAttrList struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Attr                 map[string]*AttrValue `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NameAttrList) Reset()         { *m = NameAttrList{} }
func (m *NameAttrList) String() string { return proto.CompactTextString(m) }
func (*NameAttrList) ProtoMessage()    {}
func (*NameAttrList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a312f2b3a701292, []int{1}
}

func (m *NameAttrList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameAttrList.Unmarshal(m, b)
}
func (m *NameAttrList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameAttrList.Marshal(b, m, deterministic)
}
func (m *NameAttrList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameAttrList.Merge(m, src)
}
func (m *NameAttrList) XXX_Size() int {
	return xxx_messageInfo_NameAttrList.Size(m)
}
func (m *NameAttrList) XXX_DiscardUnknown() {
	xxx_messageInfo_NameAttrList.DiscardUnknown(m)
}

var xxx_messageInfo_NameAttrList proto.InternalMessageInfo

func (m *NameAttrList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameAttrList) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*AttrValue)(nil), "tensorflow.AttrValue")
	proto.RegisterType((*AttrValue_ListValue)(nil), "tensorflow.AttrValue.ListValue")
	proto.RegisterType((*NameAttrList)(nil), "tensorflow.NameAttrList")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.NameAttrList.AttrEntry")
}

func init() { proto.RegisterFile("attr_value.proto", fileDescriptor_2a312f2b3a701292) }

var fileDescriptor_2a312f2b3a701292 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0x99, 0x74, 0xb7, 0x99, 0x96, 0xb5, 0x0c, 0x8a, 0x43, 0x10, 0x0c, 0x05, 0x65,
	0xd0, 0x92, 0x60, 0xfc, 0x83, 0x08, 0x5e, 0x58, 0x14, 0x7a, 0x21, 0xcb, 0x12, 0x17, 0x2f, 0xbc,
	0x59, 0x92, 0x3a, 0x69, 0xc3, 0xa6, 0x9d, 0x30, 0x99, 0xba, 0xf4, 0x09, 0xbc, 0xf5, 0x39, 0x7c,
	0x42, 0x2f, 0xe5, 0xcc, 0x64, 0xd3, 0xa0, 0xdb, 0xdd, 0xbb, 0x73, 0xce, 0x7c, 0xdf, 0xe4, 0xe4,
	0x37, 0xe7, 0x90, 0x71, 0xaa, 0xb5, 0xba, 0xf8, 0x91, 0x96, 0x5b, 0x11, 0x56, 0x4a, 0x6a, 0x49,
	0x89, 0x16, 0x9b, 0x5a, 0xaa, 0xbc, 0x94, 0x57, 0xfe, 0xd3, 0x7d, 0x1c, 0x2d, 0xa4, 0x12, 0x51,
	0xae, 0xd2, 0xb5, 0xb8, 0x92, 0xea, 0x32, 0xb2, 0x27, 0xd6, 0xe3, 0x4f, 0xef, 0xd2, 0x5d, 0xd4,
	0xab, 0xb4, 0x6a, 0xbe, 0xe0, 0x3f, 0xb9, 0x45, 0xbd, 0xab, 0x44, 0x6d, 0x65, 0x93, 0x9f, 0x7d,
	0xe2, 0x7d, 0xd0, 0x5a, 0x7d, 0x85, 0xe6, 0xe8, 0x09, 0x41, 0x35, 0x73, 0x02, 0xc4, 0x47, 0xf3,
	0x5e, 0x82, 0x6a, 0xc8, 0x0b, 0x86, 0x03, 0xc4, 0x31, 0xe4, 0x05, 0xe4, 0x39, 0x73, 0x03, 0xc4,
	0x1d, 0xc8, 0x73, 0xc8, 0x33, 0xd6, 0x0f, 0x10, 0x1f, 0x40, 0x9e, 0xd1, 0x67, 0xc4, 0x85, 0xcb,
	0xd9, 0x51, 0x80, 0xf8, 0x49, 0x7c, 0x3f, 0xdc, 0xf7, 0x10, 0x7e, 0x4c, 0x75, 0x7a, 0xbe, 0xab,
	0xc4, 0xbc, 0x97, 0x18, 0x0d, 0x7d, 0x45, 0xfa, 0xa6, 0x5f, 0x76, 0x1c, 0x20, 0x3e, 0x8c, 0x1f,
	0x75, 0xc5, 0xe7, 0x26, 0xfc, 0x02, 0xc7, 0x67, 0xd0, 0xe6, 0xbc, 0x97, 0x58, 0x31, 0x7d, 0x41,
	0x8e, 0xac, 0x8e, 0x0d, 0x8c, 0xed, 0xe1, 0xff, 0xb6, 0x6b, 0x47, 0x23, 0xa4, 0xaf, 0x89, 0x5b,
	0x16, 0xb5, 0x66, 0xc8, 0x18, 0x1e, 0x77, 0x0d, 0xed, 0x9f, 0x87, 0x9f, 0x8b, 0x5a, 0x9b, 0x08,
	0xfa, 0x03, 0x39, 0x0d, 0x89, 0x9b, 0x6f, 0x37, 0x0b, 0x46, 0x8c, 0x8d, 0x75, 0x6d, 0xa7, 0xe9,
	0x5a, 0x80, 0x15, 0x4c, 0xa0, 0x07, 0x1d, 0x9d, 0x90, 0x61, 0x55, 0xa6, 0x0b, 0xb1, 0x92, 0xe5,
	0x77, 0xa1, 0x98, 0x17, 0x20, 0xee, 0xcd, 0x7b, 0x49, 0xb7, 0xe8, 0xff, 0x72, 0x88, 0xd7, 0x7e,
	0x89, 0x8e, 0x2c, 0x6d, 0xcc, 0x47, 0xc0, 0x7a, 0x6c, 0x59, 0x63, 0x8e, 0x67, 0xce, 0x18, 0x01,
	0xed, 0xb1, 0xa5, 0x8d, 0xb9, 0x63, 0x2b, 0x39, 0x54, 0x80, 0x37, 0xe6, 0x03, 0x5b, 0xc9, 0xe8,
	0xb4, 0x25, 0x8e, 0x0f, 0x11, 0x37, 0x52, 0xcb, 0x3c, 0xde, 0x33, 0xc7, 0x77, 0x31, 0xbf, 0x26,
	0x1e, 0x75, 0x88, 0xe3, 0x5b, 0x88, 0xb7, 0xbc, 0xa7, 0x0d, 0x38, 0xcf, 0xc8, 0x0f, 0x82, 0xb3,
	0xd8, 0x66, 0xc7, 0xa4, 0x6f, 0x16, 0x63, 0xf2, 0x1b, 0x91, 0x51, 0xf7, 0x9c, 0x52, 0xe2, 0x6e,
	0xd2, 0xb5, 0x30, 0xef, 0xe6, 0x25, 0x26, 0xa6, 0x6f, 0x88, 0x0b, 0xbb, 0x64, 0xa8, 0x0d, 0xe3,
	0xc9, 0xa1, 0xbb, 0xcd, 0xc3, 0x7e, 0xda, 0x68, 0xb5, 0x4b, 0x8c, 0xde, 0x3f, 0xb5, 0x53, 0x6e,
	0x4a, 0x74, 0x4c, 0xf0, 0xa5, 0xd8, 0x35, 0xf7, 0x42, 0x48, 0x9f, 0x37, 0x4d, 0x98, 0xd9, 0x1f,
	0xc6, 0x0f, 0x6e, 0x9c, 0x91, 0xc4, 0x6a, 0xde, 0x39, 0x6f, 0xd1, 0x4c, 0x12, 0x26, 0xd5, 0xb2,
	0x2b, 0x6b, 0xd7, 0x6b, 0x76, 0xaf, 0x75, 0x18, 0x2e, 0xf5, 0x19, 0xfa, 0xf6, 0x7e, 0x59, 0xe8,
	0xd5, 0x36, 0x0b, 0x17, 0x72, 0x1d, 0x75, 0xf6, 0xf2, 0xe6, 0x70, 0x29, 0xff, 0x59, 0xd8, 0x3f,
	0x08, 0x65, 0x47, 0x66, 0x5d, 0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x30, 0x8e, 0x3c,
	0x4b, 0x04, 0x00, 0x00,
}