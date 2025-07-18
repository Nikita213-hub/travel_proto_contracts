// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: auth/v1/auth.proto

package authV1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail     string                 `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenTokenReq) Reset() {
	*x = GenTokenReq{}
	mi := &file_auth_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenTokenReq) ProtoMessage() {}

func (x *GenTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenTokenReq.ProtoReflect.Descriptor instead.
func (*GenTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GenTokenReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GenTokenReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type GenTokenRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Exp           int64                  `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenTokenRes) Reset() {
	*x = GenTokenRes{}
	mi := &file_auth_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenTokenRes) ProtoMessage() {}

func (x *GenTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenTokenRes.ProtoReflect.Descriptor instead.
func (*GenTokenRes) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GenTokenRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GenTokenRes) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

type VerifyTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyTokenReq) Reset() {
	*x = VerifyTokenReq{}
	mi := &file_auth_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenReq) ProtoMessage() {}

func (x *VerifyTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenReq.ProtoReflect.Descriptor instead.
func (*VerifyTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail     string                 `protobuf:"bytes,3,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Exp           int64                  `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyTokenRes) Reset() {
	*x = VerifyTokenRes{}
	mi := &file_auth_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRes) ProtoMessage() {}

func (x *VerifyTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRes.ProtoReflect.Descriptor instead.
func (*VerifyTokenRes) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyTokenRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VerifyTokenRes) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *VerifyTokenRes) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_auth_v1_auth_proto protoreflect.FileDescriptor

const file_auth_v1_auth_proto_rawDesc = "" +
	"\n" +
	"\x12auth/v1/auth.proto\x12\aauth.v1\"E\n" +
	"\vGenTokenReq\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"user_email\x18\x02 \x01(\tR\tuserEmail\"5\n" +
	"\vGenTokenRes\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x10\n" +
	"\x03exp\x18\x02 \x01(\x03R\x03exp\"&\n" +
	"\x0eVerifyTokenReq\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"Z\n" +
	"\x0eVerifyTokenRes\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"user_email\x18\x03 \x01(\tR\tuserEmail\x12\x10\n" +
	"\x03exp\x18\x04 \x01(\x03R\x03exp2\x86\x01\n" +
	"\vAuthService\x126\n" +
	"\bGenToken\x12\x14.auth.v1.GenTokenReq\x1a\x14.auth.v1.GenTokenRes\x12?\n" +
	"\vVerifyToken\x12\x17.auth.v1.VerifyTokenReq\x1a\x17.auth.v1.VerifyTokenResBEZCgithub.com/Nikita213-hub/travel_proto_contracts/pkg/proto/v1;authV1b\x06proto3"

var (
	file_auth_v1_auth_proto_rawDescOnce sync.Once
	file_auth_v1_auth_proto_rawDescData []byte
)

func file_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_v1_auth_proto_rawDesc), len(file_auth_v1_auth_proto_rawDesc)))
	})
	return file_auth_v1_auth_proto_rawDescData
}

var file_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_v1_auth_proto_goTypes = []any{
	(*GenTokenReq)(nil),    // 0: auth.v1.GenTokenReq
	(*GenTokenRes)(nil),    // 1: auth.v1.GenTokenRes
	(*VerifyTokenReq)(nil), // 2: auth.v1.VerifyTokenReq
	(*VerifyTokenRes)(nil), // 3: auth.v1.VerifyTokenRes
}
var file_auth_v1_auth_proto_depIdxs = []int32{
	0, // 0: auth.v1.AuthService.GenToken:input_type -> auth.v1.GenTokenReq
	2, // 1: auth.v1.AuthService.VerifyToken:input_type -> auth.v1.VerifyTokenReq
	1, // 2: auth.v1.AuthService.GenToken:output_type -> auth.v1.GenTokenRes
	3, // 3: auth.v1.AuthService.VerifyToken:output_type -> auth.v1.VerifyTokenRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_proto_init() }
func file_auth_v1_auth_proto_init() {
	if File_auth_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_v1_auth_proto_rawDesc), len(file_auth_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_auth_v1_auth_proto = out.File
	file_auth_v1_auth_proto_goTypes = nil
	file_auth_v1_auth_proto_depIdxs = nil
}
