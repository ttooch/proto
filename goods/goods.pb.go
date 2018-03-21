// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goods.proto

/*
Package go_micro_srv_goods is a generated protocol buffer package.

It is generated from these files:
	goods.proto

It has these top-level messages:
	Ads
	Request
	Response
*/
package go_micro_srv_goods

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

type Ads struct {
	AdsId      int32  `protobuf:"varint,1,opt,name=ads_id" json:"ads_id"`
	AdsTitle   string `protobuf:"bytes,2,opt,name=ads_title" json:"ads_title"`
	Href       string `protobuf:"bytes,3,opt,name=href" json:"href"`
	ImgPath    string `protobuf:"bytes,4,opt,name=img_path" json:"img_path"`
	ImgBaseUrl string `protobuf:"bytes,5,opt,name=img_base_url" json:"img_base_url"`
	CreatedAt  int32  `protobuf:"varint,6,opt,name=created_at" json:"created_at"`
	UpdatedAt  int32  `protobuf:"varint,7,opt,name=updated_at" json:"updated_at"`
}

func (m *Ads) Reset()                    { *m = Ads{} }
func (m *Ads) String() string            { return proto.CompactTextString(m) }
func (*Ads) ProtoMessage()               {}
func (*Ads) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ads) GetAdsId() int32 {
	if m != nil {
		return m.AdsId
	}
	return 0
}

func (m *Ads) GetAdsTitle() string {
	if m != nil {
		return m.AdsTitle
	}
	return ""
}

func (m *Ads) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *Ads) GetImgPath() string {
	if m != nil {
		return m.ImgPath
	}
	return ""
}

func (m *Ads) GetImgBaseUrl() string {
	if m != nil {
		return m.ImgBaseUrl
	}
	return ""
}

func (m *Ads) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Ads) GetUpdatedAt() int32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type Request struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name"`
	BarCode string `protobuf:"bytes,2,opt,name=bar_code" json:"bar_code"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetBarCode() string {
	if m != nil {
		return m.BarCode
	}
	return ""
}

type Response struct {
	Status int32  `protobuf:"varint,1,opt,name=status" json:"status"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data   []*Ads `protobuf:"bytes,3,rep,name=data" json:"data"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetData() []*Ads {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Ads)(nil), "go.micro.srv.goods.Ads")
	proto.RegisterType((*Request)(nil), "go.micro.srv.goods.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.goods.Response")
}

func init() { proto.RegisterFile("goods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xff, 0xf9, 0xa7, 0x49, 0xdb, 0xe9, 0x45, 0x06, 0xd1, 0xa5, 0x16, 0x29, 0x39, 0x05,
	0x84, 0x3d, 0xd4, 0x83, 0xe7, 0x58, 0x45, 0x73, 0x11, 0x59, 0xec, 0x07, 0xd8, 0xba, 0x6b, 0x0c,
	0x24, 0xdd, 0xb8, 0xb3, 0xf5, 0xfb, 0xf9, 0xcd, 0x24, 0x9b, 0xa4, 0x0a, 0xf6, 0x36, 0xef, 0xfd,
	0x76, 0xe0, 0xbd, 0x59, 0x98, 0x15, 0xc6, 0x28, 0xe2, 0x8d, 0x35, 0xce, 0x20, 0x16, 0x86, 0xd7,
	0xe5, 0xab, 0x35, 0x9c, 0xec, 0x27, 0xf7, 0x24, 0xf9, 0x0a, 0x20, 0xcc, 0x14, 0xe1, 0x29, 0x44,
	0x99, 0xa2, 0x5c, 0xb1, 0x60, 0x19, 0xa4, 0x91, 0xe8, 0x04, 0xce, 0x61, 0x92, 0x29, 0x7a, 0x29,
	0x5d, 0xa5, 0xd9, 0xff, 0x65, 0x90, 0x4e, 0xc5, 0x41, 0x23, 0xc2, 0xe8, 0xd1, 0xea, 0x37, 0x16,
	0x7a, 0xdf, 0xcf, 0xc8, 0x60, 0x9c, 0xd7, 0xc5, 0xb3, 0x74, 0xef, 0x6c, 0xe4, 0xed, 0x41, 0xe2,
	0x25, 0x40, 0x5e, 0x17, 0xb7, 0x92, 0xf4, 0xc6, 0x56, 0x2c, 0xf2, 0xf0, 0x97, 0x83, 0x0b, 0x98,
	0xae, 0xad, 0x96, 0x4e, 0xab, 0xcc, 0xb1, 0xd8, 0x67, 0xf8, 0x31, 0x5a, 0xba, 0x69, 0x54, 0x4f,
	0xc7, 0x1d, 0x3d, 0x18, 0xc9, 0x0d, 0x8c, 0x85, 0xfe, 0xd8, 0x6b, 0x72, 0x6d, 0xa8, 0x9d, 0xac,
	0xb5, 0x6f, 0x31, 0x15, 0x7e, 0x6e, 0x43, 0x6d, 0xa5, 0x5d, 0x1b, 0x35, 0x74, 0x18, 0x64, 0x22,
	0x61, 0x22, 0x34, 0x35, 0x66, 0x47, 0x1a, 0xcf, 0x20, 0x26, 0x27, 0xdd, 0x9e, 0xfa, 0x0b, 0xf4,
	0x0a, 0x4f, 0x20, 0xac, 0xa9, 0xe8, 0x37, 0xdb, 0x11, 0xaf, 0x60, 0xa4, 0xa4, 0x93, 0x2c, 0x5c,
	0x86, 0xe9, 0x6c, 0x75, 0xce, 0xff, 0x5e, 0x95, 0x67, 0x8a, 0x84, 0x7f, 0xb4, 0x7a, 0x82, 0xe8,
	0xa1, 0xb5, 0xf0, 0x1e, 0xe2, 0x3b, 0xed, 0x64, 0x59, 0xe1, 0xc5, 0xb1, 0x8d, 0xbe, 0xc0, 0x7c,
	0x71, 0x1c, 0x76, 0x21, 0x93, 0x7f, 0xdb, 0xd8, 0x7f, 0xe5, 0xf5, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x04, 0x2b, 0xde, 0xd9, 0x01, 0x00, 0x00,
}
