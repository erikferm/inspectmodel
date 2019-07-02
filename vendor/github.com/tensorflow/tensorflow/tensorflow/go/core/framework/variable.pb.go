// Code generated by protoc-gen-go. DO NOT EDIT.
// source: variable.proto

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

// Indicates when a distributed variable will be synced.
type VariableSynchronization int32

const (
	// `AUTO`: Indicates that the synchronization will be determined by the
	// current `DistributionStrategy` (eg. With `MirroredStrategy` this would be
	// `ON_WRITE`).
	VariableSynchronization_VARIABLE_SYNCHRONIZATION_AUTO VariableSynchronization = 0
	// `NONE`: Indicates that there will only be one copy of the variable, so
	// there is no need to sync.
	VariableSynchronization_VARIABLE_SYNCHRONIZATION_NONE VariableSynchronization = 1
	// `ON_WRITE`: Indicates that the variable will be updated across devices
	// every time it is written.
	VariableSynchronization_VARIABLE_SYNCHRONIZATION_ON_WRITE VariableSynchronization = 2
	// `ON_READ`: Indicates that the variable will be aggregated across devices
	// when it is read (eg. when checkpointing or when evaluating an op that uses
	// the variable).
	VariableSynchronization_VARIABLE_SYNCHRONIZATION_ON_READ VariableSynchronization = 3
)

var VariableSynchronization_name = map[int32]string{
	0: "VARIABLE_SYNCHRONIZATION_AUTO",
	1: "VARIABLE_SYNCHRONIZATION_NONE",
	2: "VARIABLE_SYNCHRONIZATION_ON_WRITE",
	3: "VARIABLE_SYNCHRONIZATION_ON_READ",
}

var VariableSynchronization_value = map[string]int32{
	"VARIABLE_SYNCHRONIZATION_AUTO":     0,
	"VARIABLE_SYNCHRONIZATION_NONE":     1,
	"VARIABLE_SYNCHRONIZATION_ON_WRITE": 2,
	"VARIABLE_SYNCHRONIZATION_ON_READ":  3,
}

func (x VariableSynchronization) String() string {
	return proto.EnumName(VariableSynchronization_name, int32(x))
}

func (VariableSynchronization) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_444e051a4c91cdda, []int{0}
}

// Indicates how a distributed variable will be aggregated.
type VariableAggregation int32

const (
	// `NONE`: This is the default, giving an error if you use a
	// variable-update operation with multiple replicas.
	VariableAggregation_VARIABLE_AGGREGATION_NONE VariableAggregation = 0
	// `SUM`: Add the updates across replicas.
	VariableAggregation_VARIABLE_AGGREGATION_SUM VariableAggregation = 1
	// `MEAN`: Take the arithmetic mean ("average") of the updates across
	// replicas.
	VariableAggregation_VARIABLE_AGGREGATION_MEAN VariableAggregation = 2
	// `ONLY_FIRST_REPLICA`: This is for when every replica is performing the same
	// update, but we only want to perform the update once. Used, e.g., for the
	// global step counter.
	VariableAggregation_VARIABLE_AGGREGATION_ONLY_FIRST_REPLICA VariableAggregation = 3
)

var VariableAggregation_name = map[int32]string{
	0: "VARIABLE_AGGREGATION_NONE",
	1: "VARIABLE_AGGREGATION_SUM",
	2: "VARIABLE_AGGREGATION_MEAN",
	3: "VARIABLE_AGGREGATION_ONLY_FIRST_REPLICA",
}

var VariableAggregation_value = map[string]int32{
	"VARIABLE_AGGREGATION_NONE":               0,
	"VARIABLE_AGGREGATION_SUM":                1,
	"VARIABLE_AGGREGATION_MEAN":               2,
	"VARIABLE_AGGREGATION_ONLY_FIRST_REPLICA": 3,
}

func (x VariableAggregation) String() string {
	return proto.EnumName(VariableAggregation_name, int32(x))
}

func (VariableAggregation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_444e051a4c91cdda, []int{1}
}

// Protocol buffer representing a Variable.
type VariableDef struct {
	// Name of the variable tensor.
	VariableName string `protobuf:"bytes,1,opt,name=variable_name,json=variableName,proto3" json:"variable_name,omitempty"`
	// Name of the tensor holding the variable's initial value.
	InitialValueName string `protobuf:"bytes,6,opt,name=initial_value_name,json=initialValueName,proto3" json:"initial_value_name,omitempty"`
	// Name of the initializer op.
	InitializerName string `protobuf:"bytes,2,opt,name=initializer_name,json=initializerName,proto3" json:"initializer_name,omitempty"`
	// Name of the snapshot tensor.
	SnapshotName string `protobuf:"bytes,3,opt,name=snapshot_name,json=snapshotName,proto3" json:"snapshot_name,omitempty"`
	// Support for saving variables as slices of a larger variable.
	SaveSliceInfoDef *SaveSliceInfoDef `protobuf:"bytes,4,opt,name=save_slice_info_def,json=saveSliceInfoDef,proto3" json:"save_slice_info_def,omitempty"`
	// Whether to represent this as a ResourceVariable.
	IsResource bool `protobuf:"varint,5,opt,name=is_resource,json=isResource,proto3" json:"is_resource,omitempty"`
	// Whether this variable should be trained.
	Trainable bool `protobuf:"varint,7,opt,name=trainable,proto3" json:"trainable,omitempty"`
	// Indicates when a distributed variable will be synced.
	Synchronization VariableSynchronization `protobuf:"varint,8,opt,name=synchronization,proto3,enum=tensorflow.VariableSynchronization" json:"synchronization,omitempty"`
	// Indicates how a distributed variable will be aggregated.
	Aggregation          VariableAggregation `protobuf:"varint,9,opt,name=aggregation,proto3,enum=tensorflow.VariableAggregation" json:"aggregation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VariableDef) Reset()         { *m = VariableDef{} }
func (m *VariableDef) String() string { return proto.CompactTextString(m) }
func (*VariableDef) ProtoMessage()    {}
func (*VariableDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_444e051a4c91cdda, []int{0}
}

func (m *VariableDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariableDef.Unmarshal(m, b)
}
func (m *VariableDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariableDef.Marshal(b, m, deterministic)
}
func (m *VariableDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariableDef.Merge(m, src)
}
func (m *VariableDef) XXX_Size() int {
	return xxx_messageInfo_VariableDef.Size(m)
}
func (m *VariableDef) XXX_DiscardUnknown() {
	xxx_messageInfo_VariableDef.DiscardUnknown(m)
}

var xxx_messageInfo_VariableDef proto.InternalMessageInfo

func (m *VariableDef) GetVariableName() string {
	if m != nil {
		return m.VariableName
	}
	return ""
}

func (m *VariableDef) GetInitialValueName() string {
	if m != nil {
		return m.InitialValueName
	}
	return ""
}

func (m *VariableDef) GetInitializerName() string {
	if m != nil {
		return m.InitializerName
	}
	return ""
}

func (m *VariableDef) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *VariableDef) GetSaveSliceInfoDef() *SaveSliceInfoDef {
	if m != nil {
		return m.SaveSliceInfoDef
	}
	return nil
}

func (m *VariableDef) GetIsResource() bool {
	if m != nil {
		return m.IsResource
	}
	return false
}

func (m *VariableDef) GetTrainable() bool {
	if m != nil {
		return m.Trainable
	}
	return false
}

func (m *VariableDef) GetSynchronization() VariableSynchronization {
	if m != nil {
		return m.Synchronization
	}
	return VariableSynchronization_VARIABLE_SYNCHRONIZATION_AUTO
}

func (m *VariableDef) GetAggregation() VariableAggregation {
	if m != nil {
		return m.Aggregation
	}
	return VariableAggregation_VARIABLE_AGGREGATION_NONE
}

type SaveSliceInfoDef struct {
	// Name of the full variable of which this is a slice.
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// Shape of the full variable.
	FullShape []int64 `protobuf:"varint,2,rep,packed,name=full_shape,json=fullShape,proto3" json:"full_shape,omitempty"`
	// Offset of this variable into the full variable.
	VarOffset []int64 `protobuf:"varint,3,rep,packed,name=var_offset,json=varOffset,proto3" json:"var_offset,omitempty"`
	// Shape of this variable.
	VarShape             []int64  `protobuf:"varint,4,rep,packed,name=var_shape,json=varShape,proto3" json:"var_shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveSliceInfoDef) Reset()         { *m = SaveSliceInfoDef{} }
func (m *SaveSliceInfoDef) String() string { return proto.CompactTextString(m) }
func (*SaveSliceInfoDef) ProtoMessage()    {}
func (*SaveSliceInfoDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_444e051a4c91cdda, []int{1}
}

func (m *SaveSliceInfoDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveSliceInfoDef.Unmarshal(m, b)
}
func (m *SaveSliceInfoDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveSliceInfoDef.Marshal(b, m, deterministic)
}
func (m *SaveSliceInfoDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveSliceInfoDef.Merge(m, src)
}
func (m *SaveSliceInfoDef) XXX_Size() int {
	return xxx_messageInfo_SaveSliceInfoDef.Size(m)
}
func (m *SaveSliceInfoDef) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveSliceInfoDef.DiscardUnknown(m)
}

var xxx_messageInfo_SaveSliceInfoDef proto.InternalMessageInfo

func (m *SaveSliceInfoDef) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *SaveSliceInfoDef) GetFullShape() []int64 {
	if m != nil {
		return m.FullShape
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarOffset() []int64 {
	if m != nil {
		return m.VarOffset
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarShape() []int64 {
	if m != nil {
		return m.VarShape
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.VariableSynchronization", VariableSynchronization_name, VariableSynchronization_value)
	proto.RegisterEnum("tensorflow.VariableAggregation", VariableAggregation_name, VariableAggregation_value)
	proto.RegisterType((*VariableDef)(nil), "tensorflow.VariableDef")
	proto.RegisterType((*SaveSliceInfoDef)(nil), "tensorflow.SaveSliceInfoDef")
}

func init() { proto.RegisterFile("variable.proto", fileDescriptor_444e051a4c91cdda) }

var fileDescriptor_444e051a4c91cdda = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xd1, 0x4e, 0xdb, 0x30,
	0x14, 0x86, 0x31, 0x61, 0xac, 0x3d, 0x1d, 0x10, 0x99, 0x8b, 0x65, 0x1a, 0x88, 0x02, 0x9b, 0xd6,
	0xb1, 0xa9, 0x48, 0xec, 0x7a, 0x17, 0x01, 0x32, 0x16, 0x0d, 0x12, 0xe4, 0x14, 0x26, 0xb8, 0xb1,
	0x4c, 0xe7, 0xb4, 0xd6, 0xd2, 0x18, 0xd9, 0x69, 0xd0, 0x78, 0x84, 0x3d, 0xc4, 0x9e, 0x60, 0xcf,
	0xb2, 0xe7, 0xd9, 0xe5, 0x14, 0xb7, 0xa1, 0x59, 0xd5, 0x72, 0x17, 0xfd, 0xff, 0xf7, 0x9f, 0x9c,
	0x73, 0x9c, 0x18, 0x56, 0x73, 0xa6, 0x04, 0xbb, 0x49, 0x78, 0xfb, 0x56, 0xc9, 0x4c, 0x62, 0xc8,
	0x78, 0xaa, 0xa5, 0x8a, 0x13, 0x79, 0xb7, 0xf3, 0xc7, 0x82, 0xc6, 0xe5, 0xd8, 0x3e, 0xe6, 0x31,
	0xde, 0x85, 0x95, 0x92, 0xa6, 0x29, 0x1b, 0x70, 0x07, 0x35, 0x51, 0xab, 0x4e, 0x9e, 0x95, 0x62,
	0xc0, 0x06, 0x1c, 0xbf, 0x07, 0x2c, 0x52, 0x91, 0x09, 0x96, 0xd0, 0x9c, 0x25, 0xc3, 0x31, 0xb9,
	0x6c, 0x48, 0x7b, 0xec, 0x5c, 0x16, 0x86, 0xa1, 0xdf, 0x42, 0xa9, 0x89, 0x7b, 0xae, 0x46, 0xec,
	0xa2, 0x61, 0xd7, 0x2a, 0xba, 0x41, 0x77, 0x61, 0x45, 0xa7, 0xec, 0x56, 0xf7, 0x65, 0x36, 0xe2,
	0xac, 0xd1, 0xdb, 0x4b, 0xd1, 0x40, 0x5f, 0x60, 0x5d, 0xb3, 0x9c, 0x53, 0x9d, 0x88, 0x2e, 0xa7,
	0x22, 0x8d, 0x25, 0xfd, 0xc6, 0x63, 0x67, 0xa9, 0x89, 0x5a, 0x8d, 0x83, 0x8d, 0xf6, 0x64, 0xb8,
	0x76, 0xc4, 0x72, 0x1e, 0x15, 0x94, 0x9f, 0xc6, 0xf2, 0x98, 0xc7, 0xc4, 0xd6, 0x53, 0x0a, 0xde,
	0x82, 0x86, 0xd0, 0x54, 0x71, 0x2d, 0x87, 0xaa, 0xcb, 0x9d, 0x27, 0x4d, 0xd4, 0xaa, 0x11, 0x10,
	0x9a, 0x8c, 0x15, 0xbc, 0x01, 0xf5, 0x4c, 0x31, 0x91, 0x16, 0xc3, 0x3b, 0x4f, 0x8d, 0x3d, 0x11,
	0xf0, 0x19, 0xac, 0xe9, 0x1f, 0x69, 0xb7, 0xaf, 0x64, 0x2a, 0xee, 0x59, 0x26, 0x64, 0xea, 0xd4,
	0x9a, 0xa8, 0xb5, 0x7a, 0xb0, 0x5b, 0xed, 0xa3, 0x5c, 0x70, 0xf4, 0x3f, 0x4a, 0xa6, 0xb3, 0xd8,
	0x85, 0x06, 0xeb, 0xf5, 0x14, 0xef, 0x8d, 0x4a, 0xd5, 0x4d, 0xa9, 0xad, 0x59, 0xa5, 0xdc, 0x09,
	0x46, 0xaa, 0x99, 0x9d, 0x9f, 0x08, 0xec, 0xe9, 0xb9, 0xf1, 0x4b, 0xa8, 0xc7, 0xc3, 0x24, 0xa9,
	0x9e, 0x68, 0xad, 0x10, 0xcc, 0x3e, 0x37, 0x01, 0x8c, 0xa9, 0xfb, 0xec, 0xb6, 0x38, 0x19, 0xab,
	0x65, 0x11, 0x83, 0x47, 0x85, 0x50, 0xd8, 0x39, 0x53, 0x54, 0xc6, 0xb1, 0xe6, 0x99, 0x63, 0x8d,
	0xec, 0x9c, 0xa9, 0xd0, 0x08, 0x45, 0xe9, 0xc2, 0x1e, 0x85, 0x97, 0x8c, 0x5b, 0xcb, 0x99, 0x32,
	0xd9, 0xbd, 0xdf, 0x08, 0x9e, 0xcf, 0x19, 0x1e, 0x6f, 0xc3, 0xe6, 0xa5, 0x4b, 0x7c, 0xf7, 0xf0,
	0xd4, 0xa3, 0xd1, 0x55, 0x70, 0xf4, 0x99, 0x84, 0x81, 0x7f, 0xed, 0x76, 0xfc, 0x30, 0xa0, 0xee,
	0x45, 0x27, 0xb4, 0x17, 0x1e, 0x45, 0x82, 0x30, 0xf0, 0x6c, 0x84, 0x5f, 0xc3, 0xf6, 0x5c, 0x24,
	0x0c, 0xe8, 0x57, 0xe2, 0x77, 0x3c, 0x7b, 0x11, 0xbf, 0x82, 0xe6, 0x63, 0x18, 0xf1, 0xdc, 0x63,
	0xdb, 0xda, 0xfb, 0x85, 0x60, 0x7d, 0xc6, 0x82, 0xf1, 0x26, 0xbc, 0x78, 0x48, 0xbb, 0x27, 0x27,
	0xc4, 0x3b, 0xa9, 0xf4, 0xb0, 0x80, 0x37, 0xc0, 0x99, 0x69, 0x47, 0x17, 0x67, 0x36, 0x9a, 0x1b,
	0x3e, 0xf3, 0xdc, 0xc0, 0x5e, 0xc4, 0xef, 0xe0, 0xcd, 0x4c, 0x3b, 0x0c, 0x4e, 0xaf, 0xe8, 0x27,
	0x9f, 0x44, 0x1d, 0x4a, 0xbc, 0xf3, 0x53, 0xff, 0xc8, 0xb5, 0xad, 0xc3, 0x14, 0x1c, 0xa9, 0x7a,
	0xd5, 0xef, 0x21, 0x56, 0x6c, 0xc0, 0xef, 0xa4, 0xfa, 0x7e, 0xb8, 0x5a, 0x76, 0x7e, 0x5e, 0xfc,
	0xe4, 0xfa, 0x1c, 0x5d, 0x7f, 0xec, 0x89, 0xac, 0x3f, 0xbc, 0x69, 0x77, 0xe5, 0x60, 0x7f, 0x12,
	0x99, 0xf3, 0xd8, 0x93, 0xfb, 0x5d, 0xa9, 0xf8, 0xfe, 0x43, 0xc1, 0xbf, 0x08, 0xdd, 0x2c, 0x9b,
	0x0b, 0xe3, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x95, 0x68, 0x20, 0x42, 0x04, 0x00,
	0x00,
}
