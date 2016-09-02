// Code generated by protoc-gen-go.
// source: as.proto
// DO NOT EDIT!

/*
Package as is a generated protocol buffer package.

It is generated from these files:
	as.proto

It has these top-level messages:
	DataRate
	RXInfo
	TXInfo
	JoinRequestRequest
	JoinRequestResponse
	PublishDataUpRequest
	PublishDataUpResponse
	PublishDataDownACKRequest
	PublishDataDownACKResponse
	PublishErrorRequest
	PublishErrorResponse
*/
package as

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}
var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}
func (RXWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorType int32

const (
	ErrorType_Generic         ErrorType = 0
	ErrorType_OTAA            ErrorType = 1
	ErrorType_DATA_UP_FCNT    ErrorType = 2
	ErrorType_DATA_UP_MIC     ErrorType = 3
	ErrorType_DATA_DOWN_QUEUE ErrorType = 4
)

var ErrorType_name = map[int32]string{
	0: "Generic",
	1: "OTAA",
	2: "DATA_UP_FCNT",
	3: "DATA_UP_MIC",
	4: "DATA_DOWN_QUEUE",
}
var ErrorType_value = map[string]int32{
	"Generic":         0,
	"OTAA":            1,
	"DATA_UP_FCNT":    2,
	"DATA_UP_MIC":     3,
	"DATA_DOWN_QUEUE": 4,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}
func (ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DataRate struct {
	Modulation   string `protobuf:"bytes,1,opt,name=modulation" json:"modulation,omitempty"`
	BandWidth    uint32 `protobuf:"varint,2,opt,name=bandWidth" json:"bandWidth,omitempty"`
	SpreadFactor uint32 `protobuf:"varint,3,opt,name=spreadFactor" json:"spreadFactor,omitempty"`
	Bitrate      uint32 `protobuf:"varint,4,opt,name=bitrate" json:"bitrate,omitempty"`
}

func (m *DataRate) Reset()                    { *m = DataRate{} }
func (m *DataRate) String() string            { return proto.CompactTextString(m) }
func (*DataRate) ProtoMessage()               {}
func (*DataRate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RXInfo struct {
	Mac     []byte  `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Time    string  `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	Rssi    int32   `protobuf:"varint,3,opt,name=rssi" json:"rssi,omitempty"`
	LoRaSNR float64 `protobuf:"fixed64,4,opt,name=loRaSNR" json:"loRaSNR,omitempty"`
}

func (m *RXInfo) Reset()                    { *m = RXInfo{} }
func (m *RXInfo) String() string            { return proto.CompactTextString(m) }
func (*RXInfo) ProtoMessage()               {}
func (*RXInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TXInfo struct {
	Frequency int64     `protobuf:"varint,1,opt,name=frequency" json:"frequency,omitempty"`
	DataRate  *DataRate `protobuf:"bytes,2,opt,name=dataRate" json:"dataRate,omitempty"`
	Adr       bool      `protobuf:"varint,3,opt,name=adr" json:"adr,omitempty"`
	CodeRate  string    `protobuf:"bytes,4,opt,name=codeRate" json:"codeRate,omitempty"`
}

func (m *TXInfo) Reset()                    { *m = TXInfo{} }
func (m *TXInfo) String() string            { return proto.CompactTextString(m) }
func (*TXInfo) ProtoMessage()               {}
func (*TXInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TXInfo) GetDataRate() *DataRate {
	if m != nil {
		return m.DataRate
	}
	return nil
}

type JoinRequestRequest struct {
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	DevAddr    []byte `protobuf:"bytes,2,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
	NetID      []byte `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
}

func (m *JoinRequestRequest) Reset()                    { *m = JoinRequestRequest{} }
func (m *JoinRequestRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestRequest) ProtoMessage()               {}
func (*JoinRequestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type JoinRequestResponse struct {
	PhyPayload  []byte   `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	NwkSKey     []byte   `protobuf:"bytes,2,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	RxDelay     uint32   `protobuf:"varint,3,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset uint32   `protobuf:"varint,4,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList      []uint32 `protobuf:"varint,5,rep,packed,name=cFList" json:"cFList,omitempty"`
	RxWindow    RXWindow `protobuf:"varint,6,opt,name=rxWindow,enum=as.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR       uint32   `protobuf:"varint,7,opt,name=rx2DR" json:"rx2DR,omitempty"`
}

func (m *JoinRequestResponse) Reset()                    { *m = JoinRequestResponse{} }
func (m *JoinRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestResponse) ProtoMessage()               {}
func (*JoinRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PublishDataUpRequest struct {
	DevEUI []byte    `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	FCnt   uint32    `protobuf:"varint,2,opt,name=fCnt" json:"fCnt,omitempty"`
	Data   []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	TxInfo *TXInfo   `protobuf:"bytes,4,opt,name=txInfo" json:"txInfo,omitempty"`
	RxInfo []*RXInfo `protobuf:"bytes,5,rep,name=rxInfo" json:"rxInfo,omitempty"`
}

func (m *PublishDataUpRequest) Reset()                    { *m = PublishDataUpRequest{} }
func (m *PublishDataUpRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishDataUpRequest) ProtoMessage()               {}
func (*PublishDataUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PublishDataUpRequest) GetTxInfo() *TXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *PublishDataUpRequest) GetRxInfo() []*RXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type PublishDataUpResponse struct {
}

func (m *PublishDataUpResponse) Reset()                    { *m = PublishDataUpResponse{} }
func (m *PublishDataUpResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishDataUpResponse) ProtoMessage()               {}
func (*PublishDataUpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PublishDataDownACKRequest struct {
	DevEUI []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	FCnt   uint32 `protobuf:"varint,2,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *PublishDataDownACKRequest) Reset()                    { *m = PublishDataDownACKRequest{} }
func (m *PublishDataDownACKRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishDataDownACKRequest) ProtoMessage()               {}
func (*PublishDataDownACKRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type PublishDataDownACKResponse struct {
}

func (m *PublishDataDownACKResponse) Reset()                    { *m = PublishDataDownACKResponse{} }
func (m *PublishDataDownACKResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishDataDownACKResponse) ProtoMessage()               {}
func (*PublishDataDownACKResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type PublishErrorRequest struct {
	DevEUI []byte    `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	Type   ErrorType `protobuf:"varint,2,opt,name=type,enum=as.ErrorType" json:"type,omitempty"`
	Error  string    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *PublishErrorRequest) Reset()                    { *m = PublishErrorRequest{} }
func (m *PublishErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishErrorRequest) ProtoMessage()               {}
func (*PublishErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PublishErrorResponse struct {
}

func (m *PublishErrorResponse) Reset()                    { *m = PublishErrorResponse{} }
func (m *PublishErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishErrorResponse) ProtoMessage()               {}
func (*PublishErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*DataRate)(nil), "as.DataRate")
	proto.RegisterType((*RXInfo)(nil), "as.RXInfo")
	proto.RegisterType((*TXInfo)(nil), "as.TXInfo")
	proto.RegisterType((*JoinRequestRequest)(nil), "as.JoinRequestRequest")
	proto.RegisterType((*JoinRequestResponse)(nil), "as.JoinRequestResponse")
	proto.RegisterType((*PublishDataUpRequest)(nil), "as.PublishDataUpRequest")
	proto.RegisterType((*PublishDataUpResponse)(nil), "as.PublishDataUpResponse")
	proto.RegisterType((*PublishDataDownACKRequest)(nil), "as.PublishDataDownACKRequest")
	proto.RegisterType((*PublishDataDownACKResponse)(nil), "as.PublishDataDownACKResponse")
	proto.RegisterType((*PublishErrorRequest)(nil), "as.PublishErrorRequest")
	proto.RegisterType((*PublishErrorResponse)(nil), "as.PublishErrorResponse")
	proto.RegisterEnum("as.RXWindow", RXWindow_name, RXWindow_value)
	proto.RegisterEnum("as.ErrorType", ErrorType_name, ErrorType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ApplicationServer service

type ApplicationServerClient interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error)
	// PublishDataUp publishes data received from an end-device.
	PublishDataUp(ctx context.Context, in *PublishDataUpRequest, opts ...grpc.CallOption) (*PublishDataUpResponse, error)
	// PublishDataDownACK publishes a data-down ack response.
	PublishDataDownACK(ctx context.Context, in *PublishDataDownACKRequest, opts ...grpc.CallOption) (*PublishDataDownACKResponse, error)
	// PublishError publishes an error message.
	PublishError(ctx context.Context, in *PublishErrorRequest, opts ...grpc.CallOption) (*PublishErrorResponse, error)
}

type applicationServerClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServerClient(cc *grpc.ClientConn) ApplicationServerClient {
	return &applicationServerClient{cc}
}

func (c *applicationServerClient) JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error) {
	out := new(JoinRequestResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/JoinRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) PublishDataUp(ctx context.Context, in *PublishDataUpRequest, opts ...grpc.CallOption) (*PublishDataUpResponse, error) {
	out := new(PublishDataUpResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/PublishDataUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) PublishDataDownACK(ctx context.Context, in *PublishDataDownACKRequest, opts ...grpc.CallOption) (*PublishDataDownACKResponse, error) {
	out := new(PublishDataDownACKResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/PublishDataDownACK", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) PublishError(ctx context.Context, in *PublishErrorRequest, opts ...grpc.CallOption) (*PublishErrorResponse, error) {
	out := new(PublishErrorResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/PublishError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationServer service

type ApplicationServerServer interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(context.Context, *JoinRequestRequest) (*JoinRequestResponse, error)
	// PublishDataUp publishes data received from an end-device.
	PublishDataUp(context.Context, *PublishDataUpRequest) (*PublishDataUpResponse, error)
	// PublishDataDownACK publishes a data-down ack response.
	PublishDataDownACK(context.Context, *PublishDataDownACKRequest) (*PublishDataDownACKResponse, error)
	// PublishError publishes an error message.
	PublishError(context.Context, *PublishErrorRequest) (*PublishErrorResponse, error)
}

func RegisterApplicationServerServer(s *grpc.Server, srv ApplicationServerServer) {
	s.RegisterService(&_ApplicationServer_serviceDesc, srv)
}

func _ApplicationServer_JoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).JoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/JoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).JoinRequest(ctx, req.(*JoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_PublishDataUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishDataUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).PublishDataUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/PublishDataUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).PublishDataUp(ctx, req.(*PublishDataUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_PublishDataDownACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishDataDownACKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).PublishDataDownACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/PublishDataDownACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).PublishDataDownACK(ctx, req.(*PublishDataDownACKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_PublishError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).PublishError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/PublishError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).PublishError(ctx, req.(*PublishErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "as.ApplicationServer",
	HandlerType: (*ApplicationServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinRequest",
			Handler:    _ApplicationServer_JoinRequest_Handler,
		},
		{
			MethodName: "PublishDataUp",
			Handler:    _ApplicationServer_PublishDataUp_Handler,
		},
		{
			MethodName: "PublishDataDownACK",
			Handler:    _ApplicationServer_PublishDataDownACK_Handler,
		},
		{
			MethodName: "PublishError",
			Handler:    _ApplicationServer_PublishError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("as.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0xc6, 0x49, 0xc8, 0xcf, 0x24, 0x01, 0x9f, 0x85, 0x13, 0x4c, 0x94, 0x83, 0x72, 0x7c, 0x15,
	0x71, 0x81, 0x44, 0xfa, 0x02, 0x8d, 0xe2, 0x04, 0xa5, 0xb4, 0x40, 0x87, 0x44, 0xa0, 0xaa, 0x12,
	0xda, 0xc4, 0x1b, 0x61, 0x35, 0xd8, 0xee, 0x7a, 0x81, 0xe4, 0xa6, 0x97, 0x7d, 0x8c, 0xbe, 0x47,
	0x9f, 0xa7, 0x2f, 0x52, 0xed, 0x7a, 0x1d, 0x1c, 0x42, 0xd5, 0xaa, 0x57, 0xcc, 0x7c, 0xdf, 0xee,
	0x7c, 0x3b, 0xf3, 0x8d, 0x09, 0x14, 0x69, 0x74, 0x14, 0xf2, 0x40, 0x04, 0x24, 0x43, 0x23, 0xfb,
	0xab, 0x01, 0x45, 0x87, 0x0a, 0x8a, 0x54, 0x30, 0x72, 0x00, 0x70, 0x17, 0xb8, 0xf7, 0x33, 0x2a,
	0xbc, 0xc0, 0xb7, 0x8c, 0xa6, 0xd1, 0x2a, 0x61, 0x0a, 0x21, 0x0d, 0x28, 0x8d, 0xa9, 0xef, 0x5e,
	0x79, 0xae, 0xb8, 0xb5, 0x32, 0x4d, 0xa3, 0x55, 0xc5, 0x27, 0x80, 0xd8, 0x50, 0x89, 0x42, 0xce,
	0xa8, 0xdb, 0xa7, 0x13, 0x11, 0x70, 0x2b, 0xab, 0x0e, 0xac, 0x60, 0xc4, 0x82, 0xc2, 0xd8, 0x13,
	0x9c, 0x0a, 0x66, 0xe5, 0x14, 0x9d, 0xa4, 0xf6, 0x47, 0xc8, 0xe3, 0xf5, 0xc0, 0x9f, 0x06, 0xc4,
	0x84, 0xec, 0x1d, 0x9d, 0x28, 0xf9, 0x0a, 0xca, 0x90, 0x10, 0xc8, 0x09, 0xef, 0x8e, 0x29, 0xc9,
	0x12, 0xaa, 0x58, 0x62, 0x3c, 0x8a, 0x3c, 0xa5, 0xb2, 0x89, 0x2a, 0x96, 0xd5, 0x67, 0x01, 0xd2,
	0xcb, 0x33, 0x54, 0xd5, 0x0d, 0x4c, 0x52, 0xfb, 0x0b, 0xe4, 0x87, 0x71, 0xf5, 0x06, 0x94, 0xa6,
	0x9c, 0x7d, 0xbe, 0x67, 0xfe, 0x64, 0xa1, 0x34, 0xb2, 0xf8, 0x04, 0x90, 0x16, 0x14, 0x5d, 0x3d,
	0x0d, 0xa5, 0x56, 0x6e, 0x57, 0x8e, 0x68, 0x74, 0x94, 0x4c, 0x08, 0x97, 0xac, 0x7c, 0x25, 0x75,
	0xe3, 0x26, 0x8b, 0x28, 0x43, 0x52, 0x87, 0xe2, 0x24, 0x70, 0x19, 0x26, 0xcd, 0x95, 0x70, 0x99,
	0xdb, 0x2e, 0x90, 0x37, 0x81, 0xe7, 0xa3, 0xd4, 0x89, 0x84, 0xfe, 0x23, 0xe7, 0x1d, 0xde, 0x2e,
	0x2e, 0xe8, 0x62, 0x16, 0x50, 0x57, 0x37, 0x9c, 0x42, 0x64, 0x3f, 0x2e, 0x7b, 0xe8, 0xb8, 0x2e,
	0x57, 0x8f, 0xa9, 0x60, 0x92, 0x92, 0x5d, 0xd8, 0xf4, 0x99, 0x18, 0x38, 0x4a, 0xbf, 0x82, 0x71,
	0x62, 0xff, 0x30, 0x60, 0x67, 0x45, 0x26, 0x0a, 0x03, 0x3f, 0x62, 0x7f, 0xa2, 0xe3, 0x3f, 0x7e,
	0xba, 0x3c, 0x65, 0x8b, 0x44, 0x47, 0xa7, 0x92, 0xe1, 0x73, 0x87, 0xcd, 0xe8, 0x42, 0xdb, 0x99,
	0xa4, 0xa4, 0x09, 0x65, 0x3e, 0x3f, 0x76, 0xf0, 0x7c, 0x3a, 0x8d, 0x98, 0xd0, 0x6e, 0xa6, 0x21,
	0x52, 0x83, 0xfc, 0xa4, 0xff, 0xd6, 0x8b, 0x84, 0xb5, 0xd9, 0xcc, 0xb6, 0xaa, 0xa8, 0x33, 0x39,
	0x63, 0x3e, 0xbf, 0xf2, 0x7c, 0x37, 0x78, 0xb4, 0xf2, 0x4d, 0xa3, 0xb5, 0x15, 0xcf, 0x18, 0xaf,
	0x63, 0x0c, 0x97, 0xac, 0xec, 0x92, 0xcf, 0xdb, 0x0e, 0x5a, 0x05, 0x55, 0x3d, 0x4e, 0xec, 0x6f,
	0x06, 0xec, 0x5e, 0xdc, 0x8f, 0x67, 0x5e, 0x74, 0x2b, 0x7d, 0x19, 0x85, 0xc9, 0x38, 0x6b, 0x90,
	0x77, 0xd9, 0x43, 0x6f, 0x34, 0xd0, 0x2d, 0xea, 0x4c, 0xae, 0xca, 0xb4, 0xeb, 0x0b, 0xbd, 0xb1,
	0x2a, 0x96, 0x98, 0xb4, 0x52, 0xcf, 0x4f, 0xc5, 0xc4, 0x86, 0xbc, 0x98, 0xcb, 0x25, 0x51, 0xdd,
	0x94, 0xdb, 0x20, 0x9f, 0x15, 0xaf, 0x0d, 0x6a, 0x46, 0x9e, 0xe1, 0xf1, 0x19, 0xd9, 0x94, 0x3e,
	0x83, 0xfa, 0x4c, 0xcc, 0xd8, 0x7b, 0xf0, 0xef, 0xb3, 0xf7, 0xc5, 0x3e, 0xd8, 0x27, 0xb0, 0x9f,
	0x22, 0x9c, 0xe0, 0xd1, 0xef, 0x74, 0x4f, 0xff, 0xe2, 0xf5, 0x76, 0x03, 0xea, 0x2f, 0x15, 0xd2,
	0x32, 0x53, 0xd8, 0xd1, 0x6c, 0x8f, 0xf3, 0x80, 0xff, 0x4e, 0xe0, 0x7f, 0xc8, 0x89, 0x45, 0x18,
	0xef, 0xfb, 0x56, 0xbb, 0x2a, 0x1b, 0x52, 0xf7, 0x86, 0x8b, 0x90, 0xa1, 0xa2, 0xa4, 0x11, 0x4c,
	0x42, 0x6a, 0x5c, 0x25, 0x8c, 0x13, 0xbb, 0xb6, 0xf4, 0x41, 0xeb, 0xc4, 0xfa, 0x87, 0x0d, 0x28,
	0x26, 0x66, 0x92, 0x02, 0x64, 0xf1, 0xfa, 0xd8, 0xdc, 0x88, 0x83, 0xb6, 0x69, 0x1c, 0x7e, 0x80,
	0xd2, 0xb2, 0x3c, 0x29, 0x43, 0xe1, 0x84, 0xf9, 0x8c, 0x7b, 0x13, 0x73, 0x83, 0x14, 0x21, 0x77,
	0x3e, 0xec, 0x74, 0x4c, 0x83, 0x98, 0x50, 0x71, 0x3a, 0xc3, 0xce, 0xcd, 0xe8, 0xe2, 0xa6, 0xdf,
	0x3d, 0x1b, 0x9a, 0x19, 0xb2, 0x0d, 0xe5, 0x04, 0x79, 0x37, 0xe8, 0x9a, 0x59, 0xb2, 0x03, 0xdb,
	0x0a, 0x70, 0xce, 0xaf, 0xce, 0x6e, 0xde, 0x8f, 0x7a, 0xa3, 0x9e, 0x99, 0x6b, 0x7f, 0xcf, 0xc0,
	0x3f, 0x9d, 0x30, 0x9c, 0x79, 0x13, 0xf5, 0x0f, 0xeb, 0x92, 0xf1, 0x07, 0xc6, 0xc9, 0x6b, 0x28,
	0xa7, 0xbe, 0x0a, 0x52, 0x93, 0x1d, 0xae, 0x7f, 0x8d, 0xf5, 0xbd, 0x35, 0x5c, 0xcf, 0x73, 0x83,
	0xf4, 0xa1, 0xba, 0xe2, 0x28, 0xb1, 0xe4, 0xd9, 0x97, 0x96, 0xb0, 0xbe, 0xff, 0x02, 0xb3, 0xac,
	0x33, 0x02, 0xb2, 0xee, 0x1b, 0xf9, 0xef, 0xd9, 0x95, 0xd5, 0xc5, 0xa8, 0x1f, 0xfc, 0x8a, 0x5e,
	0x96, 0xed, 0x42, 0x25, 0x6d, 0x04, 0xd9, 0x4b, 0xdd, 0x48, 0xaf, 0x40, 0xdd, 0x5a, 0x27, 0x92,
	0x22, 0xe3, 0xbc, 0xfa, 0x51, 0x78, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0x04, 0x4e, 0x1a, 0x2c,
	0x20, 0x06, 0x00, 0x00,
}