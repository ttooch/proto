// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goods.proto

/*
Package go_micro_api_v1_goods is a generated protocol buffer package.

It is generated from these files:
	goods.proto

It has these top-level messages:
	RpcGoods
	DetailRequest
	DetailResponse
	ListRequest
	ListResponse
*/
package go_micro_api_v1_goods

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

type RpcGoods struct {
	GoodsId    int32  `protobuf:"varint,1,opt,name=goods_id" json:"goods_id"`
	GoodsName  string `protobuf:"bytes,2,opt,name=goods_name" json:"goods_name"`
	GoodsBrief string `protobuf:"bytes,3,opt,name=goods_brief" json:"goods_brief"`
	BrandName  string `protobuf:"bytes,4,opt,name=brand_name" json:"brand_name"`
}

func (m *RpcGoods) Reset()                    { *m = RpcGoods{} }
func (m *RpcGoods) String() string            { return proto.CompactTextString(m) }
func (*RpcGoods) ProtoMessage()               {}
func (*RpcGoods) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpcGoods) GetGoodsId() int32 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *RpcGoods) GetGoodsName() string {
	if m != nil {
		return m.GoodsName
	}
	return ""
}

func (m *RpcGoods) GetGoodsBrief() string {
	if m != nil {
		return m.GoodsBrief
	}
	return ""
}

func (m *RpcGoods) GetBrandName() string {
	if m != nil {
		return m.BrandName
	}
	return ""
}

type DetailRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name"`
	BarCode string `protobuf:"bytes,2,opt,name=bar_code" json:"bar_code"`
}

func (m *DetailRequest) Reset()                    { *m = DetailRequest{} }
func (m *DetailRequest) String() string            { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()               {}
func (*DetailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DetailRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DetailRequest) GetBarCode() string {
	if m != nil {
		return m.BarCode
	}
	return ""
}

type DetailResponse struct {
	Status int32     `protobuf:"varint,1,opt,name=status" json:"status"`
	Msg    string    `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data   *RpcGoods `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *DetailResponse) Reset()                    { *m = DetailResponse{} }
func (m *DetailResponse) String() string            { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()               {}
func (*DetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DetailResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DetailResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DetailResponse) GetData() *RpcGoods {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListRequest struct {
	Page string `protobuf:"bytes,1,opt,name=page" json:"page"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListRequest) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

type ListResponse struct {
	Status int32       `protobuf:"varint,1,opt,name=status" json:"status"`
	Msg    string      `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data   []*RpcGoods `protobuf:"bytes,3,rep,name=data" json:"data"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ListResponse) GetData() []*RpcGoods {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RpcGoods)(nil), "go.micro.api.v1.goods.RpcGoods")
	proto.RegisterType((*DetailRequest)(nil), "go.micro.api.v1.goods.DetailRequest")
	proto.RegisterType((*DetailResponse)(nil), "go.micro.api.v1.goods.DetailResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.api.v1.goods.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.api.v1.goods.ListResponse")
}

func init() { proto.RegisterFile("goods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0xfd, 0xf2, 0x75, 0xad, 0x76, 0x5a, 0x45, 0x06, 0x94, 0xa5, 0x88, 0xd6, 0xa8, 0xd0, 0x53,
	0xc0, 0xf6, 0xec, 0xa5, 0x0a, 0x22, 0x88, 0xe0, 0x7a, 0xf0, 0x9c, 0x76, 0xe3, 0x12, 0x70, 0x9b,
	0x35, 0x49, 0x7b, 0xf6, 0x77, 0xf9, 0xeb, 0x24, 0xb3, 0x4d, 0xad, 0x60, 0x55, 0xf0, 0xf6, 0xde,
	0xe4, 0xbd, 0xd9, 0x37, 0x8f, 0x85, 0x76, 0x61, 0x4c, 0xee, 0x44, 0x65, 0x8d, 0x37, 0xb8, 0x57,
	0x18, 0x51, 0xea, 0x89, 0x35, 0x42, 0x56, 0x5a, 0xcc, 0xcf, 0x05, 0x3d, 0xf2, 0x57, 0x06, 0x5b,
	0x59, 0x35, 0xb9, 0x0e, 0x04, 0x53, 0xd8, 0x24, 0x70, 0x93, 0xa7, 0xac, 0xc7, 0xfa, 0x1b, 0x59,
	0xa4, 0x78, 0x00, 0x2d, 0x82, 0x77, 0xb2, 0x54, 0xe9, 0xff, 0x1e, 0xeb, 0xb7, 0xb2, 0x8f, 0x01,
	0x1e, 0x02, 0x10, 0x19, 0x59, 0xad, 0x9e, 0xd2, 0x06, 0x3d, 0xaf, 0x4c, 0x82, 0x7b, 0x64, 0xe5,
	0x34, 0x27, 0x77, 0x52, 0xbb, 0x97, 0x03, 0x7e, 0x01, 0xdb, 0x57, 0xca, 0x4b, 0xfd, 0x9c, 0xa9,
	0x97, 0x99, 0x72, 0x1e, 0x11, 0x92, 0x69, 0x50, 0x32, 0x52, 0x12, 0x0e, 0xd1, 0xc6, 0xd2, 0x5e,
	0x9a, 0x3c, 0x7e, 0x3e, 0x52, 0x6e, 0x60, 0x27, 0xda, 0x5d, 0x65, 0xa6, 0x4e, 0xe1, 0x3e, 0x34,
	0x9d, 0x97, 0x7e, 0xe6, 0x16, 0x57, 0x2c, 0x18, 0xee, 0x42, 0xa3, 0x74, 0xc5, 0xc2, 0x1f, 0x20,
	0x0e, 0x21, 0xc9, 0xa5, 0x97, 0x14, 0xb9, 0x3d, 0x38, 0x12, 0x5f, 0x76, 0x24, 0x62, 0x3f, 0x19,
	0x89, 0xf9, 0x31, 0xb4, 0x6f, 0xb5, 0xf3, 0x2b, 0x69, 0x2b, 0x59, 0x2c, 0xd3, 0x06, 0xcc, 0x4b,
	0xe8, 0xd4, 0x92, 0x3f, 0x24, 0x6a, 0xfc, 0x3a, 0xd1, 0xe0, 0x8d, 0x41, 0x87, 0xf8, 0x83, 0xb2,
	0x73, 0x3d, 0x51, 0xf8, 0x08, 0xcd, 0xba, 0x13, 0x3c, 0x5d, 0xb3, 0xe1, 0x53, 0xe3, 0xdd, 0xb3,
	0x1f, 0x54, 0xf5, 0x19, 0xfc, 0x1f, 0xde, 0x43, 0x12, 0x0e, 0x43, 0xbe, 0xc6, 0xb0, 0x52, 0x4c,
	0xf7, 0xe4, 0x5b, 0x4d, 0x5c, 0x39, 0x6e, 0xd2, 0xff, 0x39, 0x7c, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0x87, 0x9f, 0xd6, 0xae, 0x02, 0x00, 0x00,
}
