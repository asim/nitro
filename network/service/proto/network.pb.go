// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network/service/proto/network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "router/service/proto"
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

// Query is passed in a LookupRequest
type Query struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Gateway              string   `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Router               string   `protobuf:"bytes,4,opt,name=router,proto3" json:"router,omitempty"`
	Network              string   `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Query) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Query) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Query) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Query) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type ConnectRequest struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ConnectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{2}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

// PeerRequest requests list of peers
type NodesRequest struct {
	// node topology depth
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{3}
}

func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (m *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(m, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

func (m *NodesRequest) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

// PeerResponse is returned by ListPeers
type NodesResponse struct {
	// return peer topology
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{4}
}

func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (m *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(m, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GraphRequest struct {
	// node topology depth
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphRequest) Reset()         { *m = GraphRequest{} }
func (m *GraphRequest) String() string { return proto.CompactTextString(m) }
func (*GraphRequest) ProtoMessage()    {}
func (*GraphRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{5}
}

func (m *GraphRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphRequest.Unmarshal(m, b)
}
func (m *GraphRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphRequest.Marshal(b, m, deterministic)
}
func (m *GraphRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphRequest.Merge(m, src)
}
func (m *GraphRequest) XXX_Size() int {
	return xxx_messageInfo_GraphRequest.Size(m)
}
func (m *GraphRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GraphRequest proto.InternalMessageInfo

func (m *GraphRequest) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

type GraphResponse struct {
	Root                 *Peer    `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphResponse) Reset()         { *m = GraphResponse{} }
func (m *GraphResponse) String() string { return proto.CompactTextString(m) }
func (*GraphResponse) ProtoMessage()    {}
func (*GraphResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{6}
}

func (m *GraphResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphResponse.Unmarshal(m, b)
}
func (m *GraphResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphResponse.Marshal(b, m, deterministic)
}
func (m *GraphResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphResponse.Merge(m, src)
}
func (m *GraphResponse) XXX_Size() int {
	return xxx_messageInfo_GraphResponse.Size(m)
}
func (m *GraphResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GraphResponse proto.InternalMessageInfo

func (m *GraphResponse) GetRoot() *Peer {
	if m != nil {
		return m.Root
	}
	return nil
}

type RoutesRequest struct {
	// filter based on
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutesRequest) Reset()         { *m = RoutesRequest{} }
func (m *RoutesRequest) String() string { return proto.CompactTextString(m) }
func (*RoutesRequest) ProtoMessage()    {}
func (*RoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{7}
}

func (m *RoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesRequest.Unmarshal(m, b)
}
func (m *RoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesRequest.Marshal(b, m, deterministic)
}
func (m *RoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesRequest.Merge(m, src)
}
func (m *RoutesRequest) XXX_Size() int {
	return xxx_messageInfo_RoutesRequest.Size(m)
}
func (m *RoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesRequest proto.InternalMessageInfo

func (m *RoutesRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type RoutesResponse struct {
	Routes               []*proto1.Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RoutesResponse) Reset()         { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()    {}
func (*RoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{8}
}

func (m *RoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesResponse.Unmarshal(m, b)
}
func (m *RoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesResponse.Marshal(b, m, deterministic)
}
func (m *RoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesResponse.Merge(m, src)
}
func (m *RoutesResponse) XXX_Size() int {
	return xxx_messageInfo_RoutesResponse.Size(m)
}
func (m *RoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesResponse proto.InternalMessageInfo

func (m *RoutesResponse) GetRoutes() []*proto1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesRequest) Reset()         { *m = ServicesRequest{} }
func (m *ServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ServicesRequest) ProtoMessage()    {}
func (*ServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{9}
}

func (m *ServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesRequest.Unmarshal(m, b)
}
func (m *ServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesRequest.Marshal(b, m, deterministic)
}
func (m *ServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesRequest.Merge(m, src)
}
func (m *ServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ServicesRequest.Size(m)
}
func (m *ServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesRequest proto.InternalMessageInfo

type ServicesResponse struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesResponse) Reset()         { *m = ServicesResponse{} }
func (m *ServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ServicesResponse) ProtoMessage()    {}
func (*ServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{10}
}

func (m *ServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesResponse.Unmarshal(m, b)
}
func (m *ServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesResponse.Marshal(b, m, deterministic)
}
func (m *ServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesResponse.Merge(m, src)
}
func (m *ServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ServicesResponse.Size(m)
}
func (m *ServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesResponse proto.InternalMessageInfo

func (m *ServicesResponse) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{11}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{12}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Error tracks network errors
type Error struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{13}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// Status is node status
type Status struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{14}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Node is network node
type Node struct {
	// node id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// the network
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// associated metadata
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// node status
	Status               *Status  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{15}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Connect is sent when the node connects to the network
type Connect struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{16}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{17}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Peer is used to advertise node peers
type Peer struct {
	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// node peers
	Peers                []*Peer  `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{18}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Peer) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Sync is network sync message
type Sync struct {
	// peer origin
	Peer *Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// node routes
	Routes               []*proto1.Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Sync) Reset()         { *m = Sync{} }
func (m *Sync) String() string { return proto.CompactTextString(m) }
func (*Sync) ProtoMessage()    {}
func (*Sync) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aab434177f140e0, []int{19}
}

func (m *Sync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sync.Unmarshal(m, b)
}
func (m *Sync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sync.Marshal(b, m, deterministic)
}
func (m *Sync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sync.Merge(m, src)
}
func (m *Sync) XXX_Size() int {
	return xxx_messageInfo_Sync.Size(m)
}
func (m *Sync) XXX_DiscardUnknown() {
	xxx_messageInfo_Sync.DiscardUnknown(m)
}

var xxx_messageInfo_Sync proto.InternalMessageInfo

func (m *Sync) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *Sync) GetRoutes() []*proto1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "go.micro.network.Query")
	proto.RegisterType((*ConnectRequest)(nil), "go.micro.network.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "go.micro.network.ConnectResponse")
	proto.RegisterType((*NodesRequest)(nil), "go.micro.network.NodesRequest")
	proto.RegisterType((*NodesResponse)(nil), "go.micro.network.NodesResponse")
	proto.RegisterType((*GraphRequest)(nil), "go.micro.network.GraphRequest")
	proto.RegisterType((*GraphResponse)(nil), "go.micro.network.GraphResponse")
	proto.RegisterType((*RoutesRequest)(nil), "go.micro.network.RoutesRequest")
	proto.RegisterType((*RoutesResponse)(nil), "go.micro.network.RoutesResponse")
	proto.RegisterType((*ServicesRequest)(nil), "go.micro.network.ServicesRequest")
	proto.RegisterType((*ServicesResponse)(nil), "go.micro.network.ServicesResponse")
	proto.RegisterType((*StatusRequest)(nil), "go.micro.network.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "go.micro.network.StatusResponse")
	proto.RegisterType((*Error)(nil), "go.micro.network.Error")
	proto.RegisterType((*Status)(nil), "go.micro.network.Status")
	proto.RegisterType((*Node)(nil), "go.micro.network.Node")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.network.Node.MetadataEntry")
	proto.RegisterType((*Connect)(nil), "go.micro.network.Connect")
	proto.RegisterType((*Close)(nil), "go.micro.network.Close")
	proto.RegisterType((*Peer)(nil), "go.micro.network.Peer")
	proto.RegisterType((*Sync)(nil), "go.micro.network.Sync")
}

func init() {
	proto.RegisterFile("network/service/proto/network.proto", fileDescriptor_1aab434177f140e0)
}

var fileDescriptor_1aab434177f140e0 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0xc5, 0xb1, 0x1d, 0x60, 0x3e, 0x1c, 0xf8, 0x56, 0x15, 0xb5, 0x7c, 0x51, 0xc2, 0x96, 0x0b,
	0x54, 0xb5, 0x4e, 0x05, 0xaa, 0x5a, 0x15, 0x15, 0xa1, 0x22, 0x54, 0xa9, 0x12, 0x88, 0x3a, 0x2f,
	0x50, 0x13, 0xaf, 0x20, 0x02, 0xbc, 0x61, 0xbd, 0x06, 0xe5, 0x09, 0xfa, 0xa6, 0x7d, 0x89, 0xde,
	0x54, 0xbb, 0x3b, 0x36, 0x36, 0xb1, 0xd3, 0x70, 0x97, 0x99, 0x3d, 0x67, 0xc6, 0xf3, 0x77, 0x02,
	0xaf, 0x53, 0x26, 0x1f, 0xb8, 0xb8, 0x1e, 0x64, 0x4c, 0xdc, 0x8f, 0x47, 0x6c, 0x30, 0x11, 0x5c,
	0xf2, 0x01, 0x7a, 0x43, 0x6d, 0x91, 0x8d, 0x4b, 0x1e, 0xde, 0x8e, 0x47, 0x82, 0x87, 0xe8, 0x0f,
	0xb6, 0x05, 0xcf, 0x25, 0x13, 0x4f, 0x58, 0xc6, 0x69, 0x48, 0xf4, 0x97, 0x05, 0xee, 0x8f, 0x9c,
	0x89, 0x29, 0xf1, 0x61, 0x19, 0x71, 0xbe, 0xd5, 0xb7, 0x76, 0x57, 0xa3, 0xc2, 0x54, 0x2f, 0x71,
	0x92, 0x08, 0x96, 0x65, 0x7e, 0xc7, 0xbc, 0xa0, 0xa9, 0x5e, 0x2e, 0x63, 0xc9, 0x1e, 0xe2, 0xa9,
	0x6f, 0x9b, 0x17, 0x34, 0xc9, 0x26, 0x74, 0x4d, 0x1e, 0xdf, 0xd1, 0x0f, 0x68, 0x29, 0x06, 0x7e,
	0x9d, 0xef, 0x1a, 0x06, 0x9a, 0xf4, 0x10, 0x7a, 0xc7, 0x3c, 0x4d, 0xd9, 0x48, 0x46, 0xec, 0x2e,
	0x67, 0x99, 0x24, 0x6f, 0xc1, 0x4d, 0x79, 0xc2, 0x32, 0xdf, 0xea, 0xdb, 0xbb, 0xff, 0xed, 0x6d,
	0x86, 0x4f, 0x0b, 0x0c, 0xcf, 0x78, 0xc2, 0x22, 0x03, 0xa2, 0xff, 0xc3, 0x7a, 0xc9, 0xcf, 0x26,
	0x3c, 0xcd, 0x18, 0xdd, 0x81, 0x35, 0x85, 0xc8, 0x8a, 0x80, 0x2f, 0xc0, 0x4d, 0xd8, 0x44, 0x5e,
	0xe9, 0x02, 0xbd, 0xc8, 0x18, 0xf4, 0x0b, 0x78, 0x88, 0x32, 0xb4, 0x67, 0xe6, 0xdd, 0x81, 0xb5,
	0x6f, 0x22, 0x9e, 0x5c, 0xcd, 0x4f, 0x72, 0x00, 0x1e, 0xa2, 0x30, 0xc9, 0x1b, 0x70, 0x04, 0xe7,
	0x52, 0xa3, 0x1a, 0x73, 0x9c, 0x33, 0x26, 0x22, 0x8d, 0xa1, 0x87, 0xe0, 0x45, 0xaa, 0x7d, 0x65,
	0x21, 0xef, 0xc0, 0xbd, 0x53, 0x43, 0x43, 0xf6, 0xcb, 0x59, 0xb6, 0x9e, 0x69, 0x64, 0x50, 0xf4,
	0x08, 0x7a, 0x05, 0x1f, 0xb3, 0x87, 0x38, 0x9e, 0x86, 0x1a, 0x71, 0x3d, 0x34, 0x01, 0xc7, 0xa6,
	0x9b, 0x3b, 0x34, 0xdb, 0x50, 0x7c, 0x03, 0x0d, 0x61, 0xe3, 0xd1, 0x85, 0x61, 0x03, 0x58, 0xc1,
	0xa5, 0x31, 0x81, 0x57, 0xa3, 0xd2, 0xa6, 0xeb, 0xe0, 0x0d, 0x65, 0x2c, 0xf3, 0x32, 0xc0, 0x57,
	0xe8, 0x15, 0x0e, 0xa4, 0xbf, 0x87, 0x6e, 0xa6, 0x3d, 0x58, 0x97, 0x3f, 0x5b, 0x17, 0x32, 0x10,
	0x47, 0x07, 0xe0, 0x9e, 0x08, 0xc1, 0x85, 0xea, 0xfa, 0x88, 0xe7, 0xa9, 0x2c, 0xba, 0xae, 0x0d,
	0xb2, 0x01, 0xf6, 0x6d, 0x76, 0x89, 0x5b, 0xab, 0x7e, 0xd2, 0x8f, 0xd0, 0x35, 0x21, 0x54, 0x0f,
	0x99, 0xa2, 0xb6, 0xf7, 0x50, 0x47, 0x8e, 0x0c, 0x8a, 0xfe, 0xb1, 0xc0, 0x51, 0x63, 0x27, 0x3d,
	0xe8, 0x8c, 0x13, 0x3c, 0x91, 0xce, 0x38, 0x99, 0x7f, 0x1d, 0xc5, 0xae, 0xdb, 0xb5, 0x5d, 0x27,
	0x47, 0xb0, 0x72, 0xcb, 0x64, 0x9c, 0xc4, 0x32, 0xf6, 0x1d, 0x3d, 0x80, 0x9d, 0xe6, 0x25, 0x0b,
	0x4f, 0x11, 0x76, 0x92, 0x4a, 0x31, 0x8d, 0x4a, 0x56, 0xa5, 0x55, 0xee, 0x62, 0xad, 0x0a, 0x0e,
	0xc0, 0xab, 0x05, 0x53, 0xcd, 0xb9, 0x66, 0x53, 0xac, 0x44, 0xfd, 0x54, 0x4d, 0xbc, 0x8f, 0x6f,
	0x72, 0x86, 0x85, 0x18, 0xe3, 0x73, 0xe7, 0x93, 0x45, 0x3f, 0xc0, 0x32, 0x1e, 0x97, 0x5a, 0x5c,
	0xb5, 0xf8, 0xed, 0x8b, 0xab, 0x8f, 0x43, 0x63, 0xe8, 0x3e, 0xb8, 0xc7, 0x37, 0xdc, 0x6c, 0xfb,
	0xc2, 0xa4, 0x9f, 0xe0, 0xa8, 0xdd, 0x7f, 0x0e, 0x47, 0x9d, 0xec, 0x84, 0x31, 0xa1, 0x46, 0x60,
	0xcf, 0x39, 0x27, 0x03, 0xa2, 0x17, 0xe0, 0x0c, 0xa7, 0xe9, 0x48, 0x65, 0x50, 0x8e, 0x7f, 0xdd,
	0xa0, 0xc2, 0x54, 0x2e, 0xa6, 0xb3, 0xc8, 0xc5, 0xec, 0xfd, 0xb6, 0x61, 0xf9, 0x0c, 0xc7, 0x7d,
	0xfe, 0xd8, 0xbd, 0xfe, 0x6c, 0x92, 0xba, 0xea, 0x05, 0xdb, 0x73, 0x10, 0xa8, 0x6b, 0x4b, 0xe4,
	0x3b, 0xb8, 0x5a, 0x4e, 0xc8, 0xab, 0x59, 0x74, 0x55, 0x8d, 0x82, 0xad, 0xd6, 0xf7, 0x6a, 0x2c,
	0xad, 0x7f, 0x4d, 0xb1, 0xaa, 0xf2, 0xd9, 0x14, 0xab, 0x26, 0x9c, 0x74, 0x89, 0x9c, 0x42, 0xd7,
	0x28, 0x0d, 0x69, 0x00, 0xd7, 0x34, 0x2c, 0xe8, 0xb7, 0x03, 0xca, 0x70, 0x43, 0x58, 0x29, 0x34,
	0x86, 0x34, 0xf4, 0xe5, 0x89, 0x24, 0x05, 0x74, 0x1e, 0xa4, 0xfa, 0x8d, 0x28, 0x01, 0x5b, 0xad,
	0x47, 0xd3, 0xfe, 0x8d, 0x75, 0xc9, 0xa2, 0x4b, 0x17, 0x5d, 0xfd, 0x47, 0xba, 0xff, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x13, 0x9f, 0x0c, 0xc2, 0xa4, 0x07, 0x00, 0x00,
}
