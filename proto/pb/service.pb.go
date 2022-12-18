// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Direction int32

const (
	Direction_L Direction = 0
	Direction_R Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "L",
		1: "R",
	}
	Direction_value = map[string]int32{
		"L": 0,
		"R": 1,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_service_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_proto_service_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

type Orientation int32

const (
	Orientation_N Orientation = 0
	Orientation_E Orientation = 90
	Orientation_S Orientation = 180
	Orientation_W Orientation = 270
)

// Enum value maps for Orientation.
var (
	Orientation_name = map[int32]string{
		0:   "N",
		90:  "E",
		180: "S",
		270: "W",
	}
	Orientation_value = map[string]int32{
		"N": 0,
		"E": 90,
		"S": 180,
		"W": 270,
	}
)

func (x Orientation) Enum() *Orientation {
	p := new(Orientation)
	*p = x
	return p
}

func (x Orientation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Orientation) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_service_proto_enumTypes[1].Descriptor()
}

func (Orientation) Type() protoreflect.EnumType {
	return &file_proto_service_proto_enumTypes[1]
}

func (x Orientation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Orientation.Descriptor instead.
func (Orientation) EnumDescriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

type SetPlateauReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upper int32 `protobuf:"varint,1,opt,name=Upper,proto3" json:"Upper,omitempty"`
	Right int32 `protobuf:"varint,2,opt,name=Right,proto3" json:"Right,omitempty"`
}

func (x *SetPlateauReq) Reset() {
	*x = SetPlateauReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlateauReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlateauReq) ProtoMessage() {}

func (x *SetPlateauReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlateauReq.ProtoReflect.Descriptor instead.
func (*SetPlateauReq) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetPlateauReq) GetUpper() int32 {
	if x != nil {
		return x.Upper
	}
	return 0
}

func (x *SetPlateauReq) GetRight() int32 {
	if x != nil {
		return x.Right
	}
	return 0
}

type RotateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoverID   string    `protobuf:"bytes,1,opt,name=RoverID,proto3" json:"RoverID,omitempty"`
	Direction Direction `protobuf:"varint,2,opt,name=Direction,proto3,enum=pb.Direction" json:"Direction,omitempty"`
}

func (x *RotateReq) Reset() {
	*x = RotateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateReq) ProtoMessage() {}

func (x *RotateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateReq.ProtoReflect.Descriptor instead.
func (*RotateReq) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *RotateReq) GetRoverID() string {
	if x != nil {
		return x.RoverID
	}
	return ""
}

func (x *RotateReq) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_L
}

type GoForwardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoverID       string `protobuf:"bytes,1,opt,name=RoverID,proto3" json:"RoverID,omitempty"`
	NumberOfSteps int32  `protobuf:"varint,2,opt,name=NumberOfSteps,proto3" json:"NumberOfSteps,omitempty"`
}

func (x *GoForwardReq) Reset() {
	*x = GoForwardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoForwardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoForwardReq) ProtoMessage() {}

func (x *GoForwardReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoForwardReq.ProtoReflect.Descriptor instead.
func (*GoForwardReq) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *GoForwardReq) GetRoverID() string {
	if x != nil {
		return x.RoverID
	}
	return ""
}

func (x *GoForwardReq) GetNumberOfSteps() int32 {
	if x != nil {
		return x.NumberOfSteps
	}
	return 0
}

type RoverPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X                    int32       `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32       `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Orientation          Orientation `protobuf:"varint,3,opt,name=Orientation,proto3,enum=pb.Orientation" json:"Orientation,omitempty"`
	SquadIncomingMessage string      `protobuf:"bytes,4,opt,name=SquadIncomingMessage,proto3" json:"SquadIncomingMessage,omitempty"`
}

func (x *RoverPosition) Reset() {
	*x = RoverPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoverPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoverPosition) ProtoMessage() {}

func (x *RoverPosition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoverPosition.ProtoReflect.Descriptor instead.
func (*RoverPosition) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *RoverPosition) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *RoverPosition) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *RoverPosition) GetOrientation() Orientation {
	if x != nil {
		return x.Orientation
	}
	return Orientation_N
}

func (x *RoverPosition) GetSquadIncomingMessage() string {
	if x != nil {
		return x.SquadIncomingMessage
	}
	return ""
}

type Rover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoverID  string         `protobuf:"bytes,1,opt,name=RoverID,proto3" json:"RoverID,omitempty"`
	Position *RoverPosition `protobuf:"bytes,2,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (x *Rover) Reset() {
	*x = Rover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rover) ProtoMessage() {}

func (x *Rover) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rover.ProtoReflect.Descriptor instead.
func (*Rover) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *Rover) GetRoverID() string {
	if x != nil {
		return x.RoverID
	}
	return ""
}

func (x *Rover) GetPosition() *RoverPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

type LandDownReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoverList []*Rover `protobuf:"bytes,1,rep,name=RoverList,proto3" json:"RoverList,omitempty"`
}

func (x *LandDownReq) Reset() {
	*x = LandDownReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandDownReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandDownReq) ProtoMessage() {}

func (x *LandDownReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandDownReq.ProtoReflect.Descriptor instead.
func (*LandDownReq) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *LandDownReq) GetRoverList() []*Rover {
	if x != nil {
		return x.RoverList
	}
	return nil
}

type RoverSquad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rovers []*Rover `protobuf:"bytes,1,rep,name=Rovers,proto3" json:"Rovers,omitempty"`
}

func (x *RoverSquad) Reset() {
	*x = RoverSquad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoverSquad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoverSquad) ProtoMessage() {}

func (x *RoverSquad) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoverSquad.ProtoReflect.Descriptor instead.
func (*RoverSquad) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *RoverSquad) GetRovers() []*Rover {
	if x != nil {
		return x.Rovers
	}
	return nil
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x74, 0x65, 0x61, 0x75, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x70, 0x70, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x52, 0x0a, 0x09, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x09, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x0c, 0x47, 0x6f, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x74, 0x65,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x53, 0x74, 0x65, 0x70, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x76, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x59, 0x12, 0x31, 0x0a, 0x0b, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x4f, 0x72, 0x69,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x53, 0x71, 0x75, 0x61,
	0x64, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x53, 0x71, 0x75, 0x61, 0x64, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x05,
	0x52, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x2d, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36,
	0x0a, 0x0b, 0x4c, 0x61, 0x6e, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a,
	0x09, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x09, 0x52, 0x6f, 0x76,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x0a, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x53,
	0x71, 0x75, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x52,
	0x06, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x2a, 0x19, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x4c, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x52,
	0x10, 0x01, 0x2a, 0x2b, 0x0a, 0x0b, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x05, 0x0a, 0x01, 0x4e, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x45, 0x10, 0x5a, 0x12,
	0x06, 0x0a, 0x01, 0x53, 0x10, 0xb4, 0x01, 0x12, 0x06, 0x0a, 0x01, 0x57, 0x10, 0x8e, 0x02, 0x32,
	0x8b, 0x02, 0x0a, 0x18, 0x4d, 0x61, 0x72, 0x73, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a,
	0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x61, 0x75, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x61, 0x75, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x64, 0x44, 0x6f, 0x77,
	0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x09, 0x47, 0x6f, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x6f, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x77,
	0x53, 0x71, 0x75, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x71, 0x75, 0x61, 0x64, 0x42, 0x05, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_service_proto_goTypes = []interface{}{
	(Direction)(0),        // 0: pb.Direction
	(Orientation)(0),      // 1: pb.Orientation
	(*SetPlateauReq)(nil), // 2: pb.SetPlateauReq
	(*RotateReq)(nil),     // 3: pb.RotateReq
	(*GoForwardReq)(nil),  // 4: pb.GoForwardReq
	(*RoverPosition)(nil), // 5: pb.RoverPosition
	(*Rover)(nil),         // 6: pb.Rover
	(*LandDownReq)(nil),   // 7: pb.LandDownReq
	(*RoverSquad)(nil),    // 8: pb.RoverSquad
	(*empty.Empty)(nil),   // 9: google.protobuf.Empty
}
var file_proto_service_proto_depIdxs = []int32{
	0,  // 0: pb.RotateReq.Direction:type_name -> pb.Direction
	1,  // 1: pb.RoverPosition.Orientation:type_name -> pb.Orientation
	5,  // 2: pb.Rover.Position:type_name -> pb.RoverPosition
	6,  // 3: pb.LandDownReq.RoverList:type_name -> pb.Rover
	6,  // 4: pb.RoverSquad.Rovers:type_name -> pb.Rover
	2,  // 5: pb.MarsRoverNavigateService.SetPlateau:input_type -> pb.SetPlateauReq
	7,  // 6: pb.MarsRoverNavigateService.LandDown:input_type -> pb.LandDownReq
	3,  // 7: pb.MarsRoverNavigateService.Rotate:input_type -> pb.RotateReq
	4,  // 8: pb.MarsRoverNavigateService.GoForward:input_type -> pb.GoForwardReq
	9,  // 9: pb.MarsRoverNavigateService.ShowSquad:input_type -> google.protobuf.Empty
	9,  // 10: pb.MarsRoverNavigateService.SetPlateau:output_type -> google.protobuf.Empty
	9,  // 11: pb.MarsRoverNavigateService.LandDown:output_type -> google.protobuf.Empty
	6,  // 12: pb.MarsRoverNavigateService.Rotate:output_type -> pb.Rover
	6,  // 13: pb.MarsRoverNavigateService.GoForward:output_type -> pb.Rover
	8,  // 14: pb.MarsRoverNavigateService.ShowSquad:output_type -> pb.RoverSquad
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlateauReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoForwardReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoverPosition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rover); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandDownReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoverSquad); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		EnumInfos:         file_proto_service_proto_enumTypes,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
