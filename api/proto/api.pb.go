// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/api.proto

package go_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df576b66d12087a, []int{0}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// A HTTP request as RPC
// Forward by the api handler
type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df576b66d12087a, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A HTTP response as RPC
// Expected response for the api handler
type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df576b66d12087a, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP event as RPC
// Forwarded by the event handler
type Event struct {
	// e.g login
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// uuid
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// event headers
	Header map[string]*Pair `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the event data
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df576b66d12087a, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.api.Pair")
	proto.RegisterType((*Request)(nil), "go.api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "go.api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Response.HeaderEntry")
	proto.RegisterType((*Event)(nil), "go.api.Event")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.api.Event.HeaderEntry")
}

func init() {
	proto.RegisterFile("api/proto/api.proto", fileDescriptor_2df576b66d12087a)
}

var fileDescriptor_2df576b66d12087a = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0xce, 0xd3, 0x30,
	0x10, 0x54, 0xe2, 0x24, 0x6d, 0xb6, 0x08, 0x21, 0x23, 0x21, 0x53, 0x2a, 0x54, 0xe5, 0x54, 0x21,
	0x91, 0x42, 0xcb, 0x01, 0x71, 0x85, 0xaa, 0x1c, 0x2b, 0xbf, 0x81, 0xab, 0x58, 0x6d, 0x44, 0x13,
	0x9b, 0xd8, 0xa9, 0xd4, 0x87, 0xe3, 0xc0, 0x63, 0xf0, 0x36, 0xc8, 0x1b, 0xf7, 0xe7, 0xab, 0xfa,
	0x5d, 0xbe, 0xaf, 0xb7, 0x89, 0x3d, 0x3b, 0x3b, 0x3b, 0xeb, 0xc0, 0x6b, 0xa1, 0xcb, 0xa9, 0x6e,
	0x94, 0x55, 0x53, 0xa1, 0xcb, 0x1c, 0x11, 0x4d, 0x36, 0x2a, 0x17, 0xba, 0xcc, 0x3e, 0x41, 0xb4,
	0x12, 0x65, 0x43, 0x5f, 0x01, 0xf9, 0x25, 0x0f, 0x2c, 0x18, 0x07, 0x93, 0x94, 0x3b, 0x48, 0xdf,
	0x40, 0xb2, 0x17, 0xbb, 0x56, 0x1a, 0x16, 0x8e, 0xc9, 0x24, 0xe5, 0xfe, 0x2b, 0xfb, 0x4b, 0xa0,
	0xc7, 0xe5, 0xef, 0x56, 0x1a, 0xeb, 0x38, 0x95, 0xb4, 0x5b, 0x55, 0xf8, 0x42, 0xff, 0x45, 0x29,
	0x44, 0x5a, 0xd8, 0x2d, 0x0b, 0xf1, 0x14, 0x31, 0x9d, 0x43, 0xb2, 0x95, 0xa2, 0x90, 0x0d, 0x23,
	0x63, 0x32, 0x19, 0xcc, 0xde, 0xe5, 0x9d, 0x85, 0xdc, 0x8b, 0xe5, 0x3f, 0xf1, 0x76, 0x51, 0xdb,
	0xe6, 0xc0, 0x3d, 0x95, 0x7e, 0x00, 0xb2, 0x91, 0x96, 0x45, 0x58, 0xc1, 0xae, 0x2b, 0x96, 0xd2,
	0x76, 0x74, 0x47, 0xa2, 0x1f, 0x21, 0xd2, 0xca, 0x58, 0x16, 0x23, 0xf9, 0xed, 0x35, 0x79, 0xa5,
	0x8c, 0x67, 0x23, 0xcd, 0x79, 0x5c, 0xab, 0xe2, 0xc0, 0x92, 0xce, 0xa3, 0xc3, 0x2e, 0x85, 0xb6,
	0xd9, 0xb1, 0x5e, 0x97, 0x42, 0xdb, 0xec, 0x86, 0x4b, 0x18, 0x5c, 0xf8, 0xba, 0x11, 0x53, 0x06,
	0x31, 0x06, 0x83, 0xb3, 0x0e, 0x66, 0x2f, 0x8e, 0x6d, 0x5d, 0xaa, 0xbc, 0xbb, 0xfa, 0x16, 0x7e,
	0x0d, 0x86, 0x3f, 0xa0, 0x7f, 0xb4, 0xfb, 0x0c, 0x95, 0x05, 0xa4, 0xa7, 0x39, 0x9e, 0x2e, 0x93,
	0xfd, 0x09, 0xa0, 0xcf, 0xa5, 0xd1, 0xaa, 0x36, 0x92, 0xbe, 0x07, 0x30, 0x56, 0xd8, 0xd6, 0x7c,
	0x57, 0x85, 0x44, 0xb5, 0x98, 0x5f, 0x9c, 0xd0, 0x2f, 0xa7, 0xc5, 0x85, 0x98, 0xec, 0xe8, 0x9c,
	0x6c, 0xa7, 0x70, 0x73, 0x73, 0xc7, 0x78, 0xc9, 0x39, 0xde, 0xbb, 0x85, 0x99, 0xfd, 0x0b, 0x20,
	0x5e, 0xec, 0x65, 0x8d, 0x5b, 0xac, 0x45, 0x25, 0xbd, 0x08, 0x62, 0xfa, 0x12, 0xc2, 0xb2, 0xf0,
	0x6f, 0x2f, 0x2c, 0x0b, 0x3a, 0x82, 0xd4, 0x96, 0x95, 0x34, 0x56, 0x54, 0x1a, 0xfd, 0x10, 0x7e,
	0x3e, 0xa0, 0x9f, 0x4f, 0xe3, 0x45, 0x0f, 0x1f, 0x0e, 0x36, 0x78, 0x6c, 0xb6, 0x42, 0x58, 0xc1,
	0xe2, 0xae, 0xa9, 0xc3, 0x77, 0x9b, 0x6d, 0x9d, 0xe0, 0x0f, 0x3a, 0xff, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xd4, 0x6d, 0x70, 0x51, 0xb7, 0x03, 0x00, 0x00,
}
