// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: v0/cluster.proto

package protocol

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v0_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_v0_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_v0_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type NodeLoadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime         uint64  `protobuf:"varint,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
	CpuLoad1       float64 `protobuf:"fixed64,2,opt,name=cpu_load1,json=cpuLoad1,proto3" json:"cpu_load1,omitempty"`
	CpuLoad5       float64 `protobuf:"fixed64,3,opt,name=cpu_load5,json=cpuLoad5,proto3" json:"cpu_load5,omitempty"`
	CpuLoad15      float64 `protobuf:"fixed64,4,opt,name=cpu_load15,json=cpuLoad15,proto3" json:"cpu_load15,omitempty"`
	MemTotal       uint64  `protobuf:"varint,5,opt,name=mem_total,json=memTotal,proto3" json:"mem_total,omitempty"`
	MemAvailable   uint64  `protobuf:"varint,6,opt,name=mem_available,json=memAvailable,proto3" json:"mem_available,omitempty"`
	MemUsed        uint64  `protobuf:"varint,7,opt,name=mem_used,json=memUsed,proto3" json:"mem_used,omitempty"`
	MemFree        uint64  `protobuf:"varint,8,opt,name=mem_free,json=memFree,proto3" json:"mem_free,omitempty"`
	MemUsedPercent float64 `protobuf:"fixed64,9,opt,name=mem_used_percent,json=memUsedPercent,proto3" json:"mem_used_percent,omitempty"`
}

func (x *NodeLoadInfo) Reset() {
	*x = NodeLoadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v0_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeLoadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeLoadInfo) ProtoMessage() {}

func (x *NodeLoadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v0_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeLoadInfo.ProtoReflect.Descriptor instead.
func (*NodeLoadInfo) Descriptor() ([]byte, []int) {
	return file_v0_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *NodeLoadInfo) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *NodeLoadInfo) GetCpuLoad1() float64 {
	if x != nil {
		return x.CpuLoad1
	}
	return 0
}

func (x *NodeLoadInfo) GetCpuLoad5() float64 {
	if x != nil {
		return x.CpuLoad5
	}
	return 0
}

func (x *NodeLoadInfo) GetCpuLoad15() float64 {
	if x != nil {
		return x.CpuLoad15
	}
	return 0
}

func (x *NodeLoadInfo) GetMemTotal() uint64 {
	if x != nil {
		return x.MemTotal
	}
	return 0
}

func (x *NodeLoadInfo) GetMemAvailable() uint64 {
	if x != nil {
		return x.MemAvailable
	}
	return 0
}

func (x *NodeLoadInfo) GetMemUsed() uint64 {
	if x != nil {
		return x.MemUsed
	}
	return 0
}

func (x *NodeLoadInfo) GetMemFree() uint64 {
	if x != nil {
		return x.MemFree
	}
	return 0
}

func (x *NodeLoadInfo) GetMemUsedPercent() float64 {
	if x != nil {
		return x.MemUsedPercent
	}
	return 0
}

var File_v0_cluster_proto protoreflect.FileDescriptor

var file_v0_cluster_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x30, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x6e, 0x64, 0x78, 0x76, 0x2e, 0x76, 0x30, 0x1a, 0x0d, 0x76,
	0x30, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xa1,
	0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x6c,
	0x6f, 0x61, 0x64, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4c,
	0x6f, 0x61, 0x64, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x6f, 0x61, 0x64,
	0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64,
	0x35, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x35, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x31, 0x35,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x65, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x6d, 0x65, 0x6d, 0x46, 0x72, 0x65, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x5f,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x32, 0xec, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x17, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x78, 0x76, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x6d, 0x69, 0x6e, 0x64,
	0x78, 0x76, 0x2e, 0x76, 0x30, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x78, 0x76,
	0x2e, 0x76, 0x30, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x0f, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x78,
	0x76, 0x2e, 0x76, 0x30, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0a, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x6d, 0x69, 0x6e, 0x64,
	0x78, 0x76, 0x2e, 0x76, 0x30, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x17, 0x2e, 0x6d, 0x69, 0x6e,
	0x64, 0x78, 0x76, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x78, 0x76, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x6d, 0x69, 0x6e, 0x64, 0x78,
	0x76, 0x2e, 0x76, 0x30, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6a, 0x79, 0x2d, 0x64, 0x76, 0x2f, 0x6d, 0x69, 0x6e, 0x64, 0x2d, 0x78, 0x2f, 0x6d, 0x69,
	0x6e, 0x64, 0x78, 0x2d, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v0_cluster_proto_rawDescOnce sync.Once
	file_v0_cluster_proto_rawDescData = file_v0_cluster_proto_rawDesc
)

func file_v0_cluster_proto_rawDescGZIP() []byte {
	file_v0_cluster_proto_rawDescOnce.Do(func() {
		file_v0_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_v0_cluster_proto_rawDescData)
	})
	return file_v0_cluster_proto_rawDescData
}

var file_v0_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v0_cluster_proto_goTypes = []interface{}{
	(*Node)(nil),         // 0: mindxv.v0.Node
	(*NodeLoadInfo)(nil), // 1: mindxv.v0.NodeLoadInfo
	(*EmptyMessage)(nil), // 2: mindxv.v0.EmptyMessage
}
var file_v0_cluster_proto_depIdxs = []int32{
	2, // 0: mindxv.v0.NodesManager.ListNodes:input_type -> mindxv.v0.EmptyMessage
	0, // 1: mindxv.v0.NodesManager.AddNode:input_type -> mindxv.v0.Node
	0, // 2: mindxv.v0.NodesManager.RemoveNode:input_type -> mindxv.v0.Node
	2, // 3: mindxv.v0.NodesManager.LoadInfo:input_type -> mindxv.v0.EmptyMessage
	0, // 4: mindxv.v0.NodesManager.ListNodes:output_type -> mindxv.v0.Node
	0, // 5: mindxv.v0.NodesManager.AddNode:output_type -> mindxv.v0.Node
	2, // 6: mindxv.v0.NodesManager.RemoveNode:output_type -> mindxv.v0.EmptyMessage
	1, // 7: mindxv.v0.NodesManager.LoadInfo:output_type -> mindxv.v0.NodeLoadInfo
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v0_cluster_proto_init() }
func file_v0_cluster_proto_init() {
	if File_v0_cluster_proto != nil {
		return
	}
	file_v0_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v0_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_v0_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeLoadInfo); i {
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
			RawDescriptor: file_v0_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v0_cluster_proto_goTypes,
		DependencyIndexes: file_v0_cluster_proto_depIdxs,
		MessageInfos:      file_v0_cluster_proto_msgTypes,
	}.Build()
	File_v0_cluster_proto = out.File
	file_v0_cluster_proto_rawDesc = nil
	file_v0_cluster_proto_goTypes = nil
	file_v0_cluster_proto_depIdxs = nil
}
