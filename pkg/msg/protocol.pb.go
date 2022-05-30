// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: protocol.proto

package msg

import (
	actor "github.com/asynkron/protoactor-go/actor"
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

type MessageType int32

const (
	MessageType_CREATED     MessageType = 0
	MessageType_SPAWN       MessageType = 1
	MessageType_SPAWNED     MessageType = 2
	MessageType_RUN         MessageType = 3
	MessageType_SUBSCRIBE   MessageType = 4
	MessageType_UNSUBSCRIBE MessageType = 5
	MessageType_LOAD_PLUGIN MessageType = 6
	MessageType_NOTIFY      MessageType = 7
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "CREATED",
		1: "SPAWN",
		2: "SPAWNED",
		3: "RUN",
		4: "SUBSCRIBE",
		5: "UNSUBSCRIBE",
		6: "LOAD_PLUGIN",
		7: "NOTIFY",
	}
	MessageType_value = map[string]int32{
		"CREATED":     0,
		"SPAWN":       1,
		"SPAWNED":     2,
		"RUN":         3,
		"SUBSCRIBE":   4,
		"UNSUBSCRIBE": 5,
		"LOAD_PLUGIN": 6,
		"NOTIFY":      7,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_protocol_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

//
// Protocol message
type Created struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remotes []*RemoteAddress `protobuf:"bytes,1,rep,name=remotes,proto3" json:"remotes,omitempty"`
	Peers   []*actor.PID     `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Created) Reset() {
	*x = Created{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Created) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Created) ProtoMessage() {}

func (x *Created) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Created.ProtoReflect.Descriptor instead.
func (*Created) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Created) GetRemotes() []*RemoteAddress {
	if x != nil {
		return x.Remotes
	}
	return nil
}

func (x *Created) GetPeers() []*actor.PID {
	if x != nil {
		return x.Peers
	}
	return nil
}

type Spawn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host *RemoteAddress `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
}

func (x *Spawn) Reset() {
	*x = Spawn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spawn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spawn) ProtoMessage() {}

func (x *Spawn) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spawn.ProtoReflect.Descriptor instead.
func (*Spawn) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Spawn) GetHost() *RemoteAddress {
	if x != nil {
		return x.Host
	}
	return nil
}

type Spawned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bot *actor.PID `protobuf:"bytes,1,opt,name=Bot,proto3" json:"Bot,omitempty"`
}

func (x *Spawned) Reset() {
	*x = Spawned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spawned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spawned) ProtoMessage() {}

func (x *Spawned) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spawned.ProtoReflect.Descriptor instead.
func (*Spawned) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *Spawned) GetBot() *actor.PID {
	if x != nil {
		return x.Bot
	}
	return nil
}

type LoadPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin       *PluginIdentifier `protobuf:"bytes,1,opt,name=Plugin,proto3" json:"Plugin,omitempty"`
	RunAfterLoad bool              `protobuf:"varint,2,opt,name=RunAfterLoad,proto3" json:"RunAfterLoad,omitempty"`
}

func (x *LoadPlugin) Reset() {
	*x = LoadPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadPlugin) ProtoMessage() {}

func (x *LoadPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadPlugin.ProtoReflect.Descriptor instead.
func (*LoadPlugin) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *LoadPlugin) GetPlugin() *PluginIdentifier {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *LoadPlugin) GetRunAfterLoad() bool {
	if x != nil {
		return x.RunAfterLoad
	}
	return false
}

type UnloadPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin *PluginIdentifier `protobuf:"bytes,1,opt,name=Plugin,proto3" json:"Plugin,omitempty"`
}

func (x *UnloadPlugin) Reset() {
	*x = UnloadPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadPlugin) ProtoMessage() {}

func (x *UnloadPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadPlugin.ProtoReflect.Descriptor instead.
func (*UnloadPlugin) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *UnloadPlugin) GetPlugin() *PluginIdentifier {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{5}
}

type Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriber   *actor.PID    `protobuf:"bytes,1,opt,name=Subscriber,proto3" json:"Subscriber,omitempty"`
	MessageTypes []MessageType `protobuf:"varint,2,rep,packed,name=MessageTypes,proto3,enum=messages.MessageType" json:"MessageTypes,omitempty"`
}

func (x *Subscribe) Reset() {
	*x = Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscribe) ProtoMessage() {}

func (x *Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscribe.ProtoReflect.Descriptor instead.
func (*Subscribe) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *Subscribe) GetSubscriber() *actor.PID {
	if x != nil {
		return x.Subscriber
	}
	return nil
}

func (x *Subscribe) GetMessageTypes() []MessageType {
	if x != nil {
		return x.MessageTypes
	}
	return nil
}

type Unsubscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unsubscriber *actor.PID    `protobuf:"bytes,1,opt,name=Unsubscriber,proto3" json:"Unsubscriber,omitempty"`
	MessageTypes []MessageType `protobuf:"varint,2,rep,packed,name=MessageTypes,proto3,enum=messages.MessageType" json:"MessageTypes,omitempty"`
}

func (x *Unsubscribe) Reset() {
	*x = Unsubscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unsubscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unsubscribe) ProtoMessage() {}

func (x *Unsubscribe) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unsubscribe.ProtoReflect.Descriptor instead.
func (*Unsubscribe) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *Unsubscribe) GetUnsubscriber() *actor.PID {
	if x != nil {
		return x.Unsubscriber
	}
	return nil
}

func (x *Unsubscribe) GetMessageTypes() []MessageType {
	if x != nil {
		return x.MessageTypes
	}
	return nil
}

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      *actor.PID  `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	MessageType MessageType `protobuf:"varint,2,opt,name=MessageType,proto3,enum=messages.MessageType" json:"MessageType,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{8}
}

func (x *Notify) GetSource() *actor.PID {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Notify) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_CREATED
}

//
// Models
type RemoteAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *RemoteAddress) Reset() {
	*x = RemoteAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteAddress) ProtoMessage() {}

func (x *RemoteAddress) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteAddress.ProtoReflect.Descriptor instead.
func (*RemoteAddress) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{9}
}

func (x *RemoteAddress) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *RemoteAddress) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PluginIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *PluginIdentifier) Reset() {
	*x = PluginIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginIdentifier) ProtoMessage() {}

func (x *PluginIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginIdentifier.ProtoReflect.Descriptor instead.
func (*PluginIdentifier) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{10}
}

func (x *PluginIdentifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginIdentifier) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x11, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a,
	0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a,
	0x05, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x48,
	0x6f, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x03, 0x42, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x03, 0x42, 0x6f, 0x74, 0x22, 0x64, 0x0a, 0x0a,
	0x4c, 0x6f, 0x61, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x52, 0x75, 0x6e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4c, 0x6f,
	0x61, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x06,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x05, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x22, 0x72, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x0a, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x78, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x2e, 0x0a, 0x0c, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x49, 0x44, 0x52, 0x0c, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x12, 0x39, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x06, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49,
	0x44, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x3f, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x10, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x78, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x41, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x50, 0x41, 0x57, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x55, 0x4e,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x50, 0x4c, 0x55, 0x47, 0x49,
	0x4e, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x07, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x75,
	0x73, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2d, 0x62, 0x6f, 0x74, 0x73, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protocol_proto_goTypes = []interface{}{
	(MessageType)(0),         // 0: messages.MessageType
	(*Created)(nil),          // 1: messages.Created
	(*Spawn)(nil),            // 2: messages.Spawn
	(*Spawned)(nil),          // 3: messages.Spawned
	(*LoadPlugin)(nil),       // 4: messages.LoadPlugin
	(*UnloadPlugin)(nil),     // 5: messages.UnloadPlugin
	(*Run)(nil),              // 6: messages.Run
	(*Subscribe)(nil),        // 7: messages.Subscribe
	(*Unsubscribe)(nil),      // 8: messages.Unsubscribe
	(*Notify)(nil),           // 9: messages.Notify
	(*RemoteAddress)(nil),    // 10: messages.RemoteAddress
	(*PluginIdentifier)(nil), // 11: messages.PluginIdentifier
	(*actor.PID)(nil),        // 12: actor.PID
}
var file_protocol_proto_depIdxs = []int32{
	10, // 0: messages.Created.remotes:type_name -> messages.RemoteAddress
	12, // 1: messages.Created.peers:type_name -> actor.PID
	10, // 2: messages.Spawn.Host:type_name -> messages.RemoteAddress
	12, // 3: messages.Spawned.Bot:type_name -> actor.PID
	11, // 4: messages.LoadPlugin.Plugin:type_name -> messages.PluginIdentifier
	11, // 5: messages.UnloadPlugin.Plugin:type_name -> messages.PluginIdentifier
	12, // 6: messages.Subscribe.Subscriber:type_name -> actor.PID
	0,  // 7: messages.Subscribe.MessageTypes:type_name -> messages.MessageType
	12, // 8: messages.Unsubscribe.Unsubscriber:type_name -> actor.PID
	0,  // 9: messages.Unsubscribe.MessageTypes:type_name -> messages.MessageType
	12, // 10: messages.Notify.Source:type_name -> actor.PID
	0,  // 11: messages.Notify.MessageType:type_name -> messages.MessageType
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Created); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spawn); i {
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
		file_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spawned); i {
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
		file_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadPlugin); i {
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
		file_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadPlugin); i {
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
		file_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
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
		file_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscribe); i {
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
		file_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unsubscribe); i {
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
		file_protocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
		file_protocol_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteAddress); i {
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
		file_protocol_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginIdentifier); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		EnumInfos:         file_protocol_proto_enumTypes,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}
