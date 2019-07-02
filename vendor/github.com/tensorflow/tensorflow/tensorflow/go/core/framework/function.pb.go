// Code generated by protoc-gen-go. DO NOT EDIT.
// source: function.proto

package framework

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function             []*FunctionDef `protobuf:"bytes,1,rep,name=function,proto3" json:"function,omitempty"`
	Gradient             []*GradientDef `protobuf:"bytes,2,rep,name=gradient,proto3" json:"gradient,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FunctionDefLibrary) Reset()         { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()    {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{0}
}

func (m *FunctionDefLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDefLibrary.Unmarshal(m, b)
}
func (m *FunctionDefLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDefLibrary.Marshal(b, m, deterministic)
}
func (m *FunctionDefLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDefLibrary.Merge(m, src)
}
func (m *FunctionDefLibrary) XXX_Size() int {
	return xxx_messageInfo_FunctionDefLibrary.Size(m)
}
func (m *FunctionDefLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDefLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDefLibrary proto.InternalMessageInfo

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr    map[string]*AttrValue            `protobuf:"bytes,5,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ArgAttr map[uint32]*FunctionDef_ArgAttrs `protobuf:"bytes,7,rep,name=arg_attr,json=argAttr,proto3" json:"arg_attr,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef,proto3" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret,proto3" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A mapping from control output names from `signature` to node names in
	// `node_def` which should be control outputs of this function.
	ControlRet           map[string]string `protobuf:"bytes,6,rep,name=control_ret,json=controlRet,proto3" json:"control_ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FunctionDef) Reset()         { *m = FunctionDef{} }
func (m *FunctionDef) String() string { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()    {}
func (*FunctionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{1}
}

func (m *FunctionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDef.Unmarshal(m, b)
}
func (m *FunctionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDef.Marshal(b, m, deterministic)
}
func (m *FunctionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDef.Merge(m, src)
}
func (m *FunctionDef) XXX_Size() int {
	return xxx_messageInfo_FunctionDef.Size(m)
}
func (m *FunctionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDef.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDef proto.InternalMessageInfo

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetArgAttr() map[uint32]*FunctionDef_ArgAttrs {
	if m != nil {
		return m.ArgAttr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *FunctionDef) GetControlRet() map[string]string {
	if m != nil {
		return m.ControlRet
	}
	return nil
}

// Attributes for function arguments. These attributes are the same set of
// valid attributes as to _Arg nodes.
type FunctionDef_ArgAttrs struct {
	Attr                 map[string]*AttrValue `protobuf:"bytes,1,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FunctionDef_ArgAttrs) Reset()         { *m = FunctionDef_ArgAttrs{} }
func (m *FunctionDef_ArgAttrs) String() string { return proto.CompactTextString(m) }
func (*FunctionDef_ArgAttrs) ProtoMessage()    {}
func (*FunctionDef_ArgAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{1, 1}
}

func (m *FunctionDef_ArgAttrs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDef_ArgAttrs.Unmarshal(m, b)
}
func (m *FunctionDef_ArgAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDef_ArgAttrs.Marshal(b, m, deterministic)
}
func (m *FunctionDef_ArgAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDef_ArgAttrs.Merge(m, src)
}
func (m *FunctionDef_ArgAttrs) XXX_Size() int {
	return xxx_messageInfo_FunctionDef_ArgAttrs.Size(m)
}
func (m *FunctionDef_ArgAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDef_ArgAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDef_ArgAttrs proto.InternalMessageInfo

func (m *FunctionDef_ArgAttrs) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName         string   `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	GradientFunc         string   `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc,proto3" json:"gradient_func,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GradientDef) Reset()         { *m = GradientDef{} }
func (m *GradientDef) String() string { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()    {}
func (*GradientDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{2}
}

func (m *GradientDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientDef.Unmarshal(m, b)
}
func (m *GradientDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientDef.Marshal(b, m, deterministic)
}
func (m *GradientDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientDef.Merge(m, src)
}
func (m *GradientDef) XXX_Size() int {
	return xxx_messageInfo_GradientDef.Size(m)
}
func (m *GradientDef) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientDef.DiscardUnknown(m)
}

var xxx_messageInfo_GradientDef proto.InternalMessageInfo

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterMapType((map[uint32]*FunctionDef_ArgAttrs)(nil), "tensorflow.FunctionDef.ArgAttrEntry")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.FunctionDef.AttrEntry")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.FunctionDef.ControlRetEntry")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.FunctionDef.RetEntry")
	proto.RegisterType((*FunctionDef_ArgAttrs)(nil), "tensorflow.FunctionDef.ArgAttrs")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.FunctionDef.ArgAttrs.AttrEntry")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("function.proto", fileDescriptor_8ac74addf543d91a) }

var fileDescriptor_8ac74addf543d91a = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x6f, 0x8b, 0xd3, 0x30,
	0x1c, 0xc7, 0xe9, 0xba, 0xdb, 0xba, 0x5f, 0x77, 0xe7, 0x19, 0x15, 0xcb, 0x1e, 0xcd, 0x29, 0x3a,
	0x4e, 0x68, 0x61, 0x87, 0x87, 0x08, 0xa7, 0x78, 0xfe, 0x45, 0x64, 0x1e, 0x7d, 0xa0, 0x20, 0x42,
	0xc9, 0xba, 0xb4, 0x96, 0xdb, 0x92, 0x91, 0x66, 0x1e, 0x7b, 0xe2, 0x0b, 0xf1, 0x35, 0xf8, 0x02,
	0x7d, 0x28, 0x49, 0x9b, 0x36, 0xce, 0x2b, 0x43, 0xf0, 0x59, 0x9a, 0x7c, 0x3f, 0xdf, 0xfc, 0xfa,
	0x4d, 0xf2, 0x83, 0x83, 0x64, 0x4d, 0x63, 0x91, 0x31, 0xea, 0xaf, 0x38, 0x13, 0x0c, 0x81, 0x20,
	0x34, 0x67, 0x3c, 0x59, 0xb0, 0xcb, 0xc1, 0x51, 0x3d, 0x0e, 0x62, 0xc6, 0x49, 0x90, 0x70, 0xbc,
	0x24, 0x97, 0x8c, 0x5f, 0x04, 0x58, 0x08, 0x1e, 0x7d, 0xc3, 0x8b, 0x35, 0x29, 0xb8, 0xc1, 0xb8,
	0x59, 0x4b, 0xd9, 0x9c, 0x44, 0x73, 0x92, 0x94, 0xca, 0xfb, 0xcd, 0x4a, 0xb6, 0xaa, 0x75, 0xa3,
	0xef, 0x80, 0x5e, 0x97, 0xb5, 0xbd, 0x24, 0xc9, 0xfb, 0x6c, 0xc6, 0x31, 0xdf, 0xa0, 0x63, 0x70,
	0x74, 0xc5, 0x9e, 0x35, 0xb4, 0xc7, 0xee, 0xe4, 0xb6, 0x5f, 0x1b, 0xfa, 0x06, 0x11, 0x56, 0x42,
	0x09, 0xa5, 0x1c, 0xcf, 0x33, 0x42, 0x85, 0xd7, 0xfa, 0x1b, 0x7a, 0x53, 0xae, 0x29, 0x48, 0x0b,
	0x47, 0x3f, 0x3b, 0xe0, 0x1a, 0x76, 0x28, 0x80, 0x5e, 0x9e, 0xa5, 0x14, 0x8b, 0x35, 0x27, 0x9e,
	0x35, 0xb4, 0xc6, 0xee, 0xe4, 0xba, 0xe9, 0xf2, 0x61, 0x25, 0xf9, 0x5a, 0x83, 0x1e, 0x41, 0x5b,
	0xc6, 0xe4, 0xed, 0xa9, 0x1d, 0xef, 0x34, 0x94, 0xe9, 0x3f, 0x17, 0x82, 0xbf, 0xa2, 0x82, 0x6f,
	0x42, 0x25, 0x47, 0xcf, 0xc0, 0xc1, 0x3c, 0x8d, 0x14, 0xda, 0x55, 0xe8, 0xbd, 0x46, 0x94, 0xa7,
	0x35, 0xdd, 0xc5, 0xc5, 0x17, 0xf2, 0xc1, 0xd1, 0x91, 0x7b, 0xb6, 0x32, 0xb8, 0x61, 0x1a, 0x4c,
	0xd9, 0x9c, 0xc8, 0x4a, 0xbb, 0xb4, 0x18, 0xa0, 0x09, 0xd8, 0x9c, 0x08, 0xaf, 0xad, 0xa4, 0xc3,
	0xa6, 0xbd, 0x42, 0x22, 0x8a, 0x7d, 0xa4, 0x18, 0xbd, 0x05, 0x37, 0x66, 0x54, 0x70, 0xb6, 0x88,
	0x24, 0xdb, 0x51, 0xec, 0x83, 0x26, 0xf6, 0x45, 0x21, 0xad, 0x2c, 0x20, 0xae, 0x26, 0x06, 0x53,
	0xe8, 0x55, 0xff, 0x80, 0x0e, 0xc1, 0xbe, 0x20, 0x1b, 0x95, 0x6e, 0x2f, 0x94, 0x43, 0xf4, 0x10,
	0xf6, 0xd4, 0x35, 0xf3, 0x5a, 0x2a, 0xf1, 0x5b, 0xe6, 0x16, 0x92, 0xfb, 0x28, 0x17, 0xc3, 0x42,
	0xf3, 0xa4, 0xf5, 0xd8, 0x1a, 0xfc, 0xb0, 0xc0, 0x29, 0x73, 0xc9, 0xd1, 0xd3, 0xf2, 0x08, 0x8a,
	0x9b, 0x72, 0xb4, 0x23, 0xc7, 0x7c, 0xfb, 0x2c, 0xfe, 0x7b, 0x71, 0x5f, 0xa0, 0x6f, 0x9e, 0x99,
	0x69, 0xb9, 0x5f, 0x58, 0x9e, 0xfc, 0x69, 0x39, 0xdc, 0x55, 0xb2, 0xe9, 0x7e, 0x02, 0x8e, 0x8e,
	0xf8, 0x8a, 0x62, 0x6f, 0x9a, 0xce, 0x3d, 0x93, 0x3b, 0x85, 0x6b, 0x5b, 0x27, 0xf4, 0x2f, 0xf8,
	0xbb, 0xb6, 0xd3, 0x3a, 0xb4, 0x47, 0x9f, 0xc0, 0x35, 0xde, 0x11, 0xba, 0x0b, 0xfb, 0xfa, 0xf9,
	0x45, 0x14, 0x2f, 0x49, 0x69, 0xd5, 0xd7, 0x93, 0x53, 0xbc, 0x24, 0x52, 0xa4, 0x9f, 0x5b, 0x24,
	0x17, 0x4a, 0xef, 0xbe, 0x9e, 0x94, 0x3f, 0x7c, 0x46, 0xc1, 0x63, 0x3c, 0x35, 0x73, 0xa8, 0x1a,
	0xc6, 0xd9, 0x81, 0x8e, 0xe4, 0x5c, 0xb6, 0x8c, 0xfc, 0xdc, 0xfa, 0x7c, 0x9a, 0x66, 0xe2, 0xeb,
	0x7a, 0xe6, 0xc7, 0x6c, 0x19, 0x18, 0x8d, 0xe6, 0xea, 0x61, 0xca, 0xb6, 0x3a, 0xd0, 0x2f, 0xcb,
	0x9a, 0x75, 0x54, 0xfb, 0x39, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x19, 0x38, 0x41, 0x52, 0x1a,
	0x05, 0x00, 0x00,
}
