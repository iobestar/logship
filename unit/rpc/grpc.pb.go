// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package rpc

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type NLogRQ struct {
	UnitId               string   `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NLogRQ) Reset()         { *m = NLogRQ{} }
func (m *NLogRQ) String() string { return proto.CompactTextString(m) }
func (*NLogRQ) ProtoMessage()    {}
func (*NLogRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{1}
}
func (m *NLogRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NLogRQ.Unmarshal(m, b)
}
func (m *NLogRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NLogRQ.Marshal(b, m, deterministic)
}
func (dst *NLogRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NLogRQ.Merge(dst, src)
}
func (m *NLogRQ) XXX_Size() int {
	return xxx_messageInfo_NLogRQ.Size(m)
}
func (m *NLogRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_NLogRQ.DiscardUnknown(m)
}

var xxx_messageInfo_NLogRQ proto.InternalMessageInfo

func (m *NLogRQ) GetUnitId() string {
	if m != nil {
		return m.UnitId
	}
	return ""
}

func (m *NLogRQ) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TLogRQ struct {
	UnitId               string   `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Duration             string   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLogRQ) Reset()         { *m = TLogRQ{} }
func (m *TLogRQ) String() string { return proto.CompactTextString(m) }
func (*TLogRQ) ProtoMessage()    {}
func (*TLogRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{2}
}
func (m *TLogRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLogRQ.Unmarshal(m, b)
}
func (m *TLogRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLogRQ.Marshal(b, m, deterministic)
}
func (dst *TLogRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLogRQ.Merge(dst, src)
}
func (m *TLogRQ) XXX_Size() int {
	return xxx_messageInfo_TLogRQ.Size(m)
}
func (m *TLogRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_TLogRQ.DiscardUnknown(m)
}

var xxx_messageInfo_TLogRQ proto.InternalMessageInfo

func (m *TLogRQ) GetUnitId() string {
	if m != nil {
		return m.UnitId
	}
	return ""
}

func (m *TLogRQ) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *TLogRQ) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

type FLogRQ struct {
	UnitId               string   `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FLogRQ) Reset()         { *m = FLogRQ{} }
func (m *FLogRQ) String() string { return proto.CompactTextString(m) }
func (*FLogRQ) ProtoMessage()    {}
func (*FLogRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{3}
}
func (m *FLogRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FLogRQ.Unmarshal(m, b)
}
func (m *FLogRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FLogRQ.Marshal(b, m, deterministic)
}
func (dst *FLogRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FLogRQ.Merge(dst, src)
}
func (m *FLogRQ) XXX_Size() int {
	return xxx_messageInfo_FLogRQ.Size(m)
}
func (m *FLogRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_FLogRQ.DiscardUnknown(m)
}

var xxx_messageInfo_FLogRQ proto.InternalMessageInfo

func (m *FLogRQ) GetUnitId() string {
	if m != nil {
		return m.UnitId
	}
	return ""
}

type LogRS struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRS) Reset()         { *m = LogRS{} }
func (m *LogRS) String() string { return proto.CompactTextString(m) }
func (*LogRS) ProtoMessage()    {}
func (*LogRS) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{4}
}
func (m *LogRS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRS.Unmarshal(m, b)
}
func (m *LogRS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRS.Marshal(b, m, deterministic)
}
func (dst *LogRS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRS.Merge(dst, src)
}
func (m *LogRS) XXX_Size() int {
	return xxx_messageInfo_LogRS.Size(m)
}
func (m *LogRS) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRS.DiscardUnknown(m)
}

var xxx_messageInfo_LogRS proto.InternalMessageInfo

func (m *LogRS) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type NLineRQ struct {
	UnitId               string   `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NLineRQ) Reset()         { *m = NLineRQ{} }
func (m *NLineRQ) String() string { return proto.CompactTextString(m) }
func (*NLineRQ) ProtoMessage()    {}
func (*NLineRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{5}
}
func (m *NLineRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NLineRQ.Unmarshal(m, b)
}
func (m *NLineRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NLineRQ.Marshal(b, m, deterministic)
}
func (dst *NLineRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NLineRQ.Merge(dst, src)
}
func (m *NLineRQ) XXX_Size() int {
	return xxx_messageInfo_NLineRQ.Size(m)
}
func (m *NLineRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_NLineRQ.DiscardUnknown(m)
}

var xxx_messageInfo_NLineRQ proto.InternalMessageInfo

func (m *NLineRQ) GetUnitId() string {
	if m != nil {
		return m.UnitId
	}
	return ""
}

func (m *NLineRQ) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type FLineRQ struct {
	UnitId               string   `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FLineRQ) Reset()         { *m = FLineRQ{} }
func (m *FLineRQ) String() string { return proto.CompactTextString(m) }
func (*FLineRQ) ProtoMessage()    {}
func (*FLineRQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{6}
}
func (m *FLineRQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FLineRQ.Unmarshal(m, b)
}
func (m *FLineRQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FLineRQ.Marshal(b, m, deterministic)
}
func (dst *FLineRQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FLineRQ.Merge(dst, src)
}
func (m *FLineRQ) XXX_Size() int {
	return xxx_messageInfo_FLineRQ.Size(m)
}
func (m *FLineRQ) XXX_DiscardUnknown() {
	xxx_messageInfo_FLineRQ.DiscardUnknown(m)
}

var xxx_messageInfo_FLineRQ proto.InternalMessageInfo

func (m *FLineRQ) GetUnitId() string {
	if m != nil {
		return m.UnitId
	}
	return ""
}

type NLineRS struct {
	Line                 string   `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NLineRS) Reset()         { *m = NLineRS{} }
func (m *NLineRS) String() string { return proto.CompactTextString(m) }
func (*NLineRS) ProtoMessage()    {}
func (*NLineRS) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{7}
}
func (m *NLineRS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NLineRS.Unmarshal(m, b)
}
func (m *NLineRS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NLineRS.Marshal(b, m, deterministic)
}
func (dst *NLineRS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NLineRS.Merge(dst, src)
}
func (m *NLineRS) XXX_Size() int {
	return xxx_messageInfo_NLineRS.Size(m)
}
func (m *NLineRS) XXX_DiscardUnknown() {
	xxx_messageInfo_NLineRS.DiscardUnknown(m)
}

var xxx_messageInfo_NLineRS proto.InternalMessageInfo

func (m *NLineRS) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *NLineRS) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UnitRS struct {
	Unit                 string   `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnitRS) Reset()         { *m = UnitRS{} }
func (m *UnitRS) String() string { return proto.CompactTextString(m) }
func (*UnitRS) ProtoMessage()    {}
func (*UnitRS) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_aab5eddc8f6e1e32, []int{8}
}
func (m *UnitRS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnitRS.Unmarshal(m, b)
}
func (m *UnitRS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnitRS.Marshal(b, m, deterministic)
}
func (dst *UnitRS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnitRS.Merge(dst, src)
}
func (m *UnitRS) XXX_Size() int {
	return xxx_messageInfo_UnitRS.Size(m)
}
func (m *UnitRS) XXX_DiscardUnknown() {
	xxx_messageInfo_UnitRS.DiscardUnknown(m)
}

var xxx_messageInfo_UnitRS proto.InternalMessageInfo

func (m *UnitRS) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "rpc.Empty")
	proto.RegisterType((*NLogRQ)(nil), "rpc.NLogRQ")
	proto.RegisterType((*TLogRQ)(nil), "rpc.TLogRQ")
	proto.RegisterType((*FLogRQ)(nil), "rpc.FLogRQ")
	proto.RegisterType((*LogRS)(nil), "rpc.LogRS")
	proto.RegisterType((*NLineRQ)(nil), "rpc.NLineRQ")
	proto.RegisterType((*FLineRQ)(nil), "rpc.FLineRQ")
	proto.RegisterType((*NLineRS)(nil), "rpc.NLineRS")
	proto.RegisterType((*UnitRS)(nil), "rpc.UnitRS")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogUnitServiceClient is the client API for LogUnitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogUnitServiceClient interface {
	GetUnits(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LogUnitService_GetUnitsClient, error)
	GetNLogs(ctx context.Context, in *NLogRQ, opts ...grpc.CallOption) (LogUnitService_GetNLogsClient, error)
	GetTLogs(ctx context.Context, in *TLogRQ, opts ...grpc.CallOption) (LogUnitService_GetTLogsClient, error)
	GetNLines(ctx context.Context, in *NLineRQ, opts ...grpc.CallOption) (LogUnitService_GetNLinesClient, error)
	GetFLines(ctx context.Context, in *FLineRQ, opts ...grpc.CallOption) (LogUnitService_GetFLinesClient, error)
}

type logUnitServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogUnitServiceClient(cc *grpc.ClientConn) LogUnitServiceClient {
	return &logUnitServiceClient{cc}
}

func (c *logUnitServiceClient) GetUnits(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LogUnitService_GetUnitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogUnitService_serviceDesc.Streams[0], "/rpc.LogUnitService/GetUnits", opts...)
	if err != nil {
		return nil, err
	}
	x := &logUnitServiceGetUnitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogUnitService_GetUnitsClient interface {
	Recv() (*UnitRS, error)
	grpc.ClientStream
}

type logUnitServiceGetUnitsClient struct {
	grpc.ClientStream
}

func (x *logUnitServiceGetUnitsClient) Recv() (*UnitRS, error) {
	m := new(UnitRS)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logUnitServiceClient) GetNLogs(ctx context.Context, in *NLogRQ, opts ...grpc.CallOption) (LogUnitService_GetNLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogUnitService_serviceDesc.Streams[1], "/rpc.LogUnitService/GetNLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logUnitServiceGetNLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogUnitService_GetNLogsClient interface {
	Recv() (*LogRS, error)
	grpc.ClientStream
}

type logUnitServiceGetNLogsClient struct {
	grpc.ClientStream
}

func (x *logUnitServiceGetNLogsClient) Recv() (*LogRS, error) {
	m := new(LogRS)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logUnitServiceClient) GetTLogs(ctx context.Context, in *TLogRQ, opts ...grpc.CallOption) (LogUnitService_GetTLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogUnitService_serviceDesc.Streams[2], "/rpc.LogUnitService/GetTLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logUnitServiceGetTLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogUnitService_GetTLogsClient interface {
	Recv() (*LogRS, error)
	grpc.ClientStream
}

type logUnitServiceGetTLogsClient struct {
	grpc.ClientStream
}

func (x *logUnitServiceGetTLogsClient) Recv() (*LogRS, error) {
	m := new(LogRS)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logUnitServiceClient) GetNLines(ctx context.Context, in *NLineRQ, opts ...grpc.CallOption) (LogUnitService_GetNLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogUnitService_serviceDesc.Streams[3], "/rpc.LogUnitService/GetNLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &logUnitServiceGetNLinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogUnitService_GetNLinesClient interface {
	Recv() (*NLineRS, error)
	grpc.ClientStream
}

type logUnitServiceGetNLinesClient struct {
	grpc.ClientStream
}

func (x *logUnitServiceGetNLinesClient) Recv() (*NLineRS, error) {
	m := new(NLineRS)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logUnitServiceClient) GetFLines(ctx context.Context, in *FLineRQ, opts ...grpc.CallOption) (LogUnitService_GetFLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogUnitService_serviceDesc.Streams[4], "/rpc.LogUnitService/GetFLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &logUnitServiceGetFLinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogUnitService_GetFLinesClient interface {
	Recv() (*NLineRS, error)
	grpc.ClientStream
}

type logUnitServiceGetFLinesClient struct {
	grpc.ClientStream
}

func (x *logUnitServiceGetFLinesClient) Recv() (*NLineRS, error) {
	m := new(NLineRS)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogUnitServiceServer is the server API for LogUnitService service.
type LogUnitServiceServer interface {
	GetUnits(*Empty, LogUnitService_GetUnitsServer) error
	GetNLogs(*NLogRQ, LogUnitService_GetNLogsServer) error
	GetTLogs(*TLogRQ, LogUnitService_GetTLogsServer) error
	GetNLines(*NLineRQ, LogUnitService_GetNLinesServer) error
	GetFLines(*FLineRQ, LogUnitService_GetFLinesServer) error
}

func RegisterLogUnitServiceServer(s *grpc.Server, srv LogUnitServiceServer) {
	s.RegisterService(&_LogUnitService_serviceDesc, srv)
}

func _LogUnitService_GetUnits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogUnitServiceServer).GetUnits(m, &logUnitServiceGetUnitsServer{stream})
}

type LogUnitService_GetUnitsServer interface {
	Send(*UnitRS) error
	grpc.ServerStream
}

type logUnitServiceGetUnitsServer struct {
	grpc.ServerStream
}

func (x *logUnitServiceGetUnitsServer) Send(m *UnitRS) error {
	return x.ServerStream.SendMsg(m)
}

func _LogUnitService_GetNLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NLogRQ)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogUnitServiceServer).GetNLogs(m, &logUnitServiceGetNLogsServer{stream})
}

type LogUnitService_GetNLogsServer interface {
	Send(*LogRS) error
	grpc.ServerStream
}

type logUnitServiceGetNLogsServer struct {
	grpc.ServerStream
}

func (x *logUnitServiceGetNLogsServer) Send(m *LogRS) error {
	return x.ServerStream.SendMsg(m)
}

func _LogUnitService_GetTLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TLogRQ)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogUnitServiceServer).GetTLogs(m, &logUnitServiceGetTLogsServer{stream})
}

type LogUnitService_GetTLogsServer interface {
	Send(*LogRS) error
	grpc.ServerStream
}

type logUnitServiceGetTLogsServer struct {
	grpc.ServerStream
}

func (x *logUnitServiceGetTLogsServer) Send(m *LogRS) error {
	return x.ServerStream.SendMsg(m)
}

func _LogUnitService_GetNLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NLineRQ)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogUnitServiceServer).GetNLines(m, &logUnitServiceGetNLinesServer{stream})
}

type LogUnitService_GetNLinesServer interface {
	Send(*NLineRS) error
	grpc.ServerStream
}

type logUnitServiceGetNLinesServer struct {
	grpc.ServerStream
}

func (x *logUnitServiceGetNLinesServer) Send(m *NLineRS) error {
	return x.ServerStream.SendMsg(m)
}

func _LogUnitService_GetFLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FLineRQ)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogUnitServiceServer).GetFLines(m, &logUnitServiceGetFLinesServer{stream})
}

type LogUnitService_GetFLinesServer interface {
	Send(*NLineRS) error
	grpc.ServerStream
}

type logUnitServiceGetFLinesServer struct {
	grpc.ServerStream
}

func (x *logUnitServiceGetFLinesServer) Send(m *NLineRS) error {
	return x.ServerStream.SendMsg(m)
}

var _LogUnitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LogUnitService",
	HandlerType: (*LogUnitServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUnits",
			Handler:       _LogUnitService_GetUnits_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetNLogs",
			Handler:       _LogUnitService_GetNLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTLogs",
			Handler:       _LogUnitService_GetTLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetNLines",
			Handler:       _LogUnitService_GetNLines_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFLines",
			Handler:       _LogUnitService_GetFLines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_grpc_aab5eddc8f6e1e32) }

var fileDescriptor_grpc_aab5eddc8f6e1e32 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xc9, 0x7f, 0xff, 0x26, 0xdb, 0xab, 0xec, 0x10, 0x44, 0x4b, 0xf5, 0xb0, 0x05, 0x84,
	0x79, 0x19, 0xe2, 0x0e, 0x7a, 0x1e, 0x58, 0x11, 0xca, 0xd0, 0x6e, 0x3b, 0xcb, 0xec, 0xb2, 0x1a,
	0x98, 0x49, 0x68, 0x33, 0x61, 0x9f, 0xd5, 0x2f, 0x23, 0x49, 0xda, 0xd9, 0xc3, 0x74, 0x78, 0x7b,
	0x9f, 0xbe, 0xbf, 0xe7, 0x81, 0xbe, 0x4f, 0x00, 0xf2, 0x42, 0x67, 0x43, 0x5d, 0x28, 0xa3, 0x68,
	0xab, 0xd0, 0x19, 0x23, 0x10, 0xdc, 0xbf, 0x6b, 0xb3, 0x65, 0xb7, 0x80, 0x27, 0x89, 0xca, 0xd3,
	0x67, 0x7a, 0x06, 0x64, 0x23, 0x85, 0x79, 0x11, 0xcb, 0x10, 0xf5, 0xd0, 0xa0, 0x93, 0x62, 0x2b,
	0x1f, 0x97, 0xf4, 0x04, 0x82, 0x4c, 0x6d, 0xa4, 0x09, 0xff, 0xf5, 0xd0, 0x20, 0x48, 0xbd, 0x60,
	0x73, 0xc0, 0xb3, 0x03, 0xc6, 0x53, 0xc0, 0x6a, 0xb5, 0x2a, 0xb9, 0x77, 0xb6, 0xd2, 0x4a, 0xd1,
	0x08, 0xda, 0xcb, 0x4d, 0xb1, 0x30, 0x42, 0xc9, 0xb0, 0xe5, 0x1c, 0x3b, 0xcd, 0xfa, 0x80, 0xe3,
	0xdf, 0x63, 0x59, 0x1f, 0x02, 0x4b, 0x4c, 0x69, 0x08, 0x44, 0x2f, 0xb6, 0x6b, 0xb5, 0xa8, 0x89,
	0x5a, 0xb2, 0x3b, 0x20, 0x93, 0x44, 0x48, 0xfe, 0xf7, 0xdf, 0x62, 0x40, 0xe2, 0x03, 0x4e, 0x36,
	0xaa, 0xd3, 0xa7, 0x94, 0xc2, 0xff, 0xb5, 0x90, 0xbc, 0x02, 0xdc, 0xfc, 0x43, 0xf0, 0x05, 0xe0,
	0xb9, 0x14, 0xc6, 0x7b, 0x6c, 0x50, 0xed, 0xb1, 0xf3, 0xcd, 0x27, 0x82, 0x6e, 0xa2, 0x72, 0x4b,
	0x4c, 0x79, 0xf1, 0x21, 0x32, 0x4e, 0x2f, 0xa1, 0xfd, 0xc0, 0x8d, 0xfd, 0x52, 0x52, 0x18, 0xda,
	0xfe, 0x5c, 0x63, 0xd1, 0x91, 0x9b, 0x7d, 0xd6, 0x35, 0xaa, 0x30, 0xdb, 0x61, 0x49, 0xfd, 0xca,
	0xf7, 0x19, 0x79, 0x8f, 0xbb, 0xd4, 0x0e, 0x9b, 0x35, 0xb0, 0xd9, 0x3e, 0xec, 0x0a, 0x3a, 0x2e,
	0x4d, 0x48, 0x5e, 0xd2, 0xe3, 0x2a, 0xce, 0x9d, 0x23, 0x6a, 0xaa, 0x6f, 0x34, 0x6e, 0xa2, 0xf1,
	0x7e, 0x74, 0x7c, 0x0e, 0x5d, 0xa1, 0x86, 0x6b, 0x95, 0x97, 0x6f, 0x42, 0xdb, 0xdd, 0x98, 0x24,
	0x5e, 0x3c, 0xa1, 0x57, 0xec, 0x9e, 0xe5, 0xe8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x38, 0x4e, 0x20,
	0x8a, 0xa4, 0x02, 0x00, 0x00,
}
