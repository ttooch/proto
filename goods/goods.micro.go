// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: goods.proto

/*
Package go_micro_api_v1_goods is a generated protocol buffer package.

It is generated from these files:
	goods.proto

It has these top-level messages:
	Good
	DetailRequest
	DetailResponse
	ListRequest
	ListResponse
*/
package go_micro_api_v1_goods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Goods service

type GoodsClient interface {
	Detail(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type goodsClient struct {
	c           client.Client
	serviceName string
}

func NewGoodsClient(serviceName string, c client.Client) GoodsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.api.v1.goods"
	}
	return &goodsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *goodsClient) Detail(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.Detail", in)
	out := new(DetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goods service

type GoodsHandler interface {
	Detail(context.Context, *DetailRequest, *DetailResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterGoodsHandler(s server.Server, hdlr GoodsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Goods{hdlr}, opts...))
}

type Goods struct {
	GoodsHandler
}

func (h *Goods) Detail(ctx context.Context, in *DetailRequest, out *DetailResponse) error {
	return h.GoodsHandler.Detail(ctx, in, out)
}

func (h *Goods) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.GoodsHandler.List(ctx, in, out)
}
