// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: internal/controller/grpc/proto/user.proto

package geo

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

type Geocode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat string `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng string `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Geocode) Reset() {
	*x = Geocode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_controller_grpc_proto_geo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geocode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geocode) ProtoMessage() {}

func (x *Geocode) ProtoReflect() protoreflect.Message {
	mi := &file_internal_controller_grpc_proto_geo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geocode.ProtoReflect.Descriptor instead.
func (*Geocode) Descriptor() ([]byte, []int) {
	return file_internal_controller_grpc_proto_geo_proto_rawDescGZIP(), []int{0}
}

func (x *Geocode) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Geocode) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_controller_grpc_proto_geo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_internal_controller_grpc_proto_geo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_internal_controller_grpc_proto_geo_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

var File_internal_controller_grpc_proto_geo_proto protoreflect.FileDescriptor

var file_internal_controller_grpc_proto_geo_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x65, 0x6f, 0x22,
	0x2d, 0x0a, 0x07, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x37,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x32, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65,
	0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0c, 0x2e, 0x67, 0x65, 0x6f, 0x2e,
	0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x0c, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_controller_grpc_proto_geo_proto_rawDescOnce sync.Once
	file_internal_controller_grpc_proto_geo_proto_rawDescData = file_internal_controller_grpc_proto_geo_proto_rawDesc
)

func file_internal_controller_grpc_proto_geo_proto_rawDescGZIP() []byte {
	file_internal_controller_grpc_proto_geo_proto_rawDescOnce.Do(func() {
		file_internal_controller_grpc_proto_geo_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_controller_grpc_proto_geo_proto_rawDescData)
	})
	return file_internal_controller_grpc_proto_geo_proto_rawDescData
}

var file_internal_controller_grpc_proto_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_controller_grpc_proto_geo_proto_goTypes = []any{
	(*Geocode)(nil), // 0: geo.Geocode
	(*Address)(nil), // 1: geo.Address
}
var file_internal_controller_grpc_proto_geo_proto_depIdxs = []int32{
	0, // 0: geo.GeoService.GeocodeToAddress:input_type -> geo.Geocode
	1, // 1: geo.GeoService.GeocodeToAddress:output_type -> geo.Address
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_controller_grpc_proto_geo_proto_init() }
func file_internal_controller_grpc_proto_geo_proto_init() {
	if File_internal_controller_grpc_proto_geo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_controller_grpc_proto_geo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Geocode); i {
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
		file_internal_controller_grpc_proto_geo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_internal_controller_grpc_proto_geo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_controller_grpc_proto_geo_proto_goTypes,
		DependencyIndexes: file_internal_controller_grpc_proto_geo_proto_depIdxs,
		MessageInfos:      file_internal_controller_grpc_proto_geo_proto_msgTypes,
	}.Build()
	File_internal_controller_grpc_proto_geo_proto = out.File
	file_internal_controller_grpc_proto_geo_proto_rawDesc = nil
	file_internal_controller_grpc_proto_geo_proto_goTypes = nil
	file_internal_controller_grpc_proto_geo_proto_depIdxs = nil
}
