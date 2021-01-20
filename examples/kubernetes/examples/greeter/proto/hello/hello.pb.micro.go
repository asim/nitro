// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/hello/hello.proto

package go_micro_srv_greeter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayService interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type sayService struct {
	c    client.Client
	name string
}

func NewSayService(name string, c client.Client) SayService {
	return &sayService{
		c:    c,
		name: name,
	}
}

func (c *sayService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Say.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler, opts ...server.HandlerOption) error {
	type say interface {
		Hello(ctx context.Context, in *Request, out *Response) error
	}
	type Say struct {
		say
	}
	h := &sayHandler{hdlr}
	return s.Handle(s.NewHandler(&Say{h}, opts...))
}

type sayHandler struct {
	SayHandler
}

func (h *sayHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}
