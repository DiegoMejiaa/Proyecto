// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/juegos.proto

package juegos

import (
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

type JuegosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JuegosRequest) Reset() {
	*x = JuegosRequest{}
	mi := &file_proto_juegos_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JuegosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegosRequest) ProtoMessage() {}

func (x *JuegosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juegos_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegosRequest.ProtoReflect.Descriptor instead.
func (*JuegosRequest) Descriptor() ([]byte, []int) {
	return file_proto_juegos_proto_rawDescGZIP(), []int{0}
}

func (x *JuegosRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JuegosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Win   string `protobuf:"bytes,2,opt,name=win,proto3" json:"win,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *JuegosResponse) Reset() {
	*x = JuegosResponse{}
	mi := &file_proto_juegos_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JuegosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegosResponse) ProtoMessage() {}

func (x *JuegosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juegos_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegosResponse.ProtoReflect.Descriptor instead.
func (*JuegosResponse) Descriptor() ([]byte, []int) {
	return file_proto_juegos_proto_rawDescGZIP(), []int{1}
}

func (x *JuegosResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JuegosResponse) GetWin() string {
	if x != nil {
		return x.Win
	}
	return ""
}

func (x *JuegosResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type NewJuegosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Win   string `protobuf:"bytes,2,opt,name=win,proto3" json:"win,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *NewJuegosRequest) Reset() {
	*x = NewJuegosRequest{}
	mi := &file_proto_juegos_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewJuegosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJuegosRequest) ProtoMessage() {}

func (x *NewJuegosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juegos_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewJuegosRequest.ProtoReflect.Descriptor instead.
func (*NewJuegosRequest) Descriptor() ([]byte, []int) {
	return file_proto_juegos_proto_rawDescGZIP(), []int{2}
}

func (x *NewJuegosRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewJuegosRequest) GetWin() string {
	if x != nil {
		return x.Win
	}
	return ""
}

func (x *NewJuegosRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type AddJuegosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddJuegosResponse) Reset() {
	*x = AddJuegosResponse{}
	mi := &file_proto_juegos_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddJuegosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddJuegosResponse) ProtoMessage() {}

func (x *AddJuegosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juegos_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddJuegosResponse.ProtoReflect.Descriptor instead.
func (*AddJuegosResponse) Descriptor() ([]byte, []int) {
	return file_proto_juegos_proto_rawDescGZIP(), []int{3}
}

func (x *AddJuegosResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_juegos_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juegos_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_juegos_proto_rawDescGZIP(), []int{4}
}

type JuegosWinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Win string `protobuf:"bytes,1,opt,name=win,proto3" json:"win,omitempty"`
}

func (x *JuegosWinRequest) Reset() {
	*x = JuegosWinRequest{}
	mi := &file_proto_juegos_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JuegosWinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegosWinRequest) ProtoMessage() {}

func (x *JuegosWinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_juegos_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegosWinRequest.ProtoReflect.Descriptor instead.
func (*JuegosWinRequest) Descriptor() ([]byte, []int) {
	return file_proto_juegos_proto_rawDescGZIP(), []int{5}
}

func (x *JuegosWinRequest) GetWin() string {
	if x != nil {
		return x.Win
	}
	return ""
}

var File_proto_juegos_proto protoreflect.FileDescriptor

var file_proto_juegos_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x22, 0x23, 0x0a, 0x0d,
	0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4c, 0x0a, 0x0e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22,
	0x4e, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22,
	0x29, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x57, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x32, 0x95, 0x02, 0x0a, 0x0d, 0x4a, 0x75,
	0x65, 0x67, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x6a,
	0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x4a, 0x75, 0x65,
	0x67, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x6a,
	0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x6a, 0x75,
	0x65, 0x67, 0x6f, 0x73, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4a, 0x75, 0x65, 0x67,
	0x6f, 0x73, 0x12, 0x18, 0x2e, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x4a,
	0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6a,
	0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x42, 0x79, 0x57, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x6a, 0x75,
	0x65, 0x67, 0x6f, 0x73, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x57, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x2e, 0x4a,
	0x75, 0x65, 0x67, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x1a, 0x5a, 0x18, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x5f, 0x67, 0x72, 0x63, 0x70,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_juegos_proto_rawDescOnce sync.Once
	file_proto_juegos_proto_rawDescData = file_proto_juegos_proto_rawDesc
)

func file_proto_juegos_proto_rawDescGZIP() []byte {
	file_proto_juegos_proto_rawDescOnce.Do(func() {
		file_proto_juegos_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_juegos_proto_rawDescData)
	})
	return file_proto_juegos_proto_rawDescData
}

var file_proto_juegos_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_juegos_proto_goTypes = []any{
	(*JuegosRequest)(nil),     // 0: juegos.JuegosRequest
	(*JuegosResponse)(nil),    // 1: juegos.JuegosResponse
	(*NewJuegosRequest)(nil),  // 2: juegos.NewJuegosRequest
	(*AddJuegosResponse)(nil), // 3: juegos.AddJuegosResponse
	(*Empty)(nil),             // 4: juegos.Empty
	(*JuegosWinRequest)(nil),  // 5: juegos.JuegosWinRequest
}
var file_proto_juegos_proto_depIdxs = []int32{
	0, // 0: juegos.JuegosService.GetJuegosInfo:input_type -> juegos.JuegosRequest
	4, // 1: juegos.JuegosService.GetJuegosList:input_type -> juegos.Empty
	2, // 2: juegos.JuegosService.AddJuegos:input_type -> juegos.NewJuegosRequest
	5, // 3: juegos.JuegosService.GetJuegosByWin:input_type -> juegos.JuegosWinRequest
	1, // 4: juegos.JuegosService.GetJuegosInfo:output_type -> juegos.JuegosResponse
	1, // 5: juegos.JuegosService.GetJuegosList:output_type -> juegos.JuegosResponse
	3, // 6: juegos.JuegosService.AddJuegos:output_type -> juegos.AddJuegosResponse
	1, // 7: juegos.JuegosService.GetJuegosByWin:output_type -> juegos.JuegosResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_juegos_proto_init() }
func file_proto_juegos_proto_init() {
	if File_proto_juegos_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_juegos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_juegos_proto_goTypes,
		DependencyIndexes: file_proto_juegos_proto_depIdxs,
		MessageInfos:      file_proto_juegos_proto_msgTypes,
	}.Build()
	File_proto_juegos_proto = out.File
	file_proto_juegos_proto_rawDesc = nil
	file_proto_juegos_proto_goTypes = nil
	file_proto_juegos_proto_depIdxs = nil
}
