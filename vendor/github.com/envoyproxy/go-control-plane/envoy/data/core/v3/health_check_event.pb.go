// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/core/v3/health_check_event.proto

package envoy_data_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HealthCheckFailureType int32

const (
	HealthCheckFailureType_ACTIVE  HealthCheckFailureType = 0
	HealthCheckFailureType_PASSIVE HealthCheckFailureType = 1
	HealthCheckFailureType_NETWORK HealthCheckFailureType = 2
)

var HealthCheckFailureType_name = map[int32]string{
	0: "ACTIVE",
	1: "PASSIVE",
	2: "NETWORK",
}

var HealthCheckFailureType_value = map[string]int32{
	"ACTIVE":  0,
	"PASSIVE": 1,
	"NETWORK": 2,
}

func (x HealthCheckFailureType) String() string {
	return proto.EnumName(HealthCheckFailureType_name, int32(x))
}

func (HealthCheckFailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{0}
}

type HealthCheckerType int32

const (
	HealthCheckerType_HTTP  HealthCheckerType = 0
	HealthCheckerType_TCP   HealthCheckerType = 1
	HealthCheckerType_GRPC  HealthCheckerType = 2
	HealthCheckerType_REDIS HealthCheckerType = 3
)

var HealthCheckerType_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "GRPC",
	3: "REDIS",
}

var HealthCheckerType_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"GRPC":  2,
	"REDIS": 3,
}

func (x HealthCheckerType) String() string {
	return proto.EnumName(HealthCheckerType_name, int32(x))
}

func (HealthCheckerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{1}
}

type HealthCheckEvent struct {
	HealthCheckerType HealthCheckerType `protobuf:"varint,1,opt,name=health_checker_type,json=healthCheckerType,proto3,enum=envoy.data.core.v3.HealthCheckerType" json:"health_checker_type,omitempty"`
	Host              *v3.Address       `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	ClusterName       string            `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*HealthCheckEvent_EjectUnhealthyEvent
	//	*HealthCheckEvent_AddHealthyEvent
	//	*HealthCheckEvent_HealthCheckFailureEvent
	//	*HealthCheckEvent_DegradedHealthyHost
	//	*HealthCheckEvent_NoLongerDegradedHost
	Event                isHealthCheckEvent_Event `protobuf_oneof:"event"`
	Timestamp            *timestamp.Timestamp     `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HealthCheckEvent) Reset()         { *m = HealthCheckEvent{} }
func (m *HealthCheckEvent) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEvent) ProtoMessage()    {}
func (*HealthCheckEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{0}
}

func (m *HealthCheckEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckEvent.Unmarshal(m, b)
}
func (m *HealthCheckEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckEvent.Marshal(b, m, deterministic)
}
func (m *HealthCheckEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEvent.Merge(m, src)
}
func (m *HealthCheckEvent) XXX_Size() int {
	return xxx_messageInfo_HealthCheckEvent.Size(m)
}
func (m *HealthCheckEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEvent.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEvent proto.InternalMessageInfo

func (m *HealthCheckEvent) GetHealthCheckerType() HealthCheckerType {
	if m != nil {
		return m.HealthCheckerType
	}
	return HealthCheckerType_HTTP
}

func (m *HealthCheckEvent) GetHost() *v3.Address {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *HealthCheckEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type isHealthCheckEvent_Event interface {
	isHealthCheckEvent_Event()
}

type HealthCheckEvent_EjectUnhealthyEvent struct {
	EjectUnhealthyEvent *HealthCheckEjectUnhealthy `protobuf:"bytes,4,opt,name=eject_unhealthy_event,json=ejectUnhealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_AddHealthyEvent struct {
	AddHealthyEvent *HealthCheckAddHealthy `protobuf:"bytes,5,opt,name=add_healthy_event,json=addHealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_HealthCheckFailureEvent struct {
	HealthCheckFailureEvent *HealthCheckFailure `protobuf:"bytes,7,opt,name=health_check_failure_event,json=healthCheckFailureEvent,proto3,oneof"`
}

type HealthCheckEvent_DegradedHealthyHost struct {
	DegradedHealthyHost *DegradedHealthyHost `protobuf:"bytes,8,opt,name=degraded_healthy_host,json=degradedHealthyHost,proto3,oneof"`
}

type HealthCheckEvent_NoLongerDegradedHost struct {
	NoLongerDegradedHost *NoLongerDegradedHost `protobuf:"bytes,9,opt,name=no_longer_degraded_host,json=noLongerDegradedHost,proto3,oneof"`
}

func (*HealthCheckEvent_EjectUnhealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_AddHealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_HealthCheckFailureEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_DegradedHealthyHost) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_NoLongerDegradedHost) isHealthCheckEvent_Event() {}

func (m *HealthCheckEvent) GetEvent() isHealthCheckEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *HealthCheckEvent) GetEjectUnhealthyEvent() *HealthCheckEjectUnhealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_EjectUnhealthyEvent); ok {
		return x.EjectUnhealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetAddHealthyEvent() *HealthCheckAddHealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_AddHealthyEvent); ok {
		return x.AddHealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetHealthCheckFailureEvent() *HealthCheckFailure {
	if x, ok := m.GetEvent().(*HealthCheckEvent_HealthCheckFailureEvent); ok {
		return x.HealthCheckFailureEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetDegradedHealthyHost() *DegradedHealthyHost {
	if x, ok := m.GetEvent().(*HealthCheckEvent_DegradedHealthyHost); ok {
		return x.DegradedHealthyHost
	}
	return nil
}

func (m *HealthCheckEvent) GetNoLongerDegradedHost() *NoLongerDegradedHost {
	if x, ok := m.GetEvent().(*HealthCheckEvent_NoLongerDegradedHost); ok {
		return x.NoLongerDegradedHost
	}
	return nil
}

func (m *HealthCheckEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheckEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheckEvent_EjectUnhealthyEvent)(nil),
		(*HealthCheckEvent_AddHealthyEvent)(nil),
		(*HealthCheckEvent_HealthCheckFailureEvent)(nil),
		(*HealthCheckEvent_DegradedHealthyHost)(nil),
		(*HealthCheckEvent_NoLongerDegradedHost)(nil),
	}
}

type HealthCheckEjectUnhealthy struct {
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v3.HealthCheckFailureType" json:"failure_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckEjectUnhealthy) Reset()         { *m = HealthCheckEjectUnhealthy{} }
func (m *HealthCheckEjectUnhealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEjectUnhealthy) ProtoMessage()    {}
func (*HealthCheckEjectUnhealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{1}
}

func (m *HealthCheckEjectUnhealthy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Unmarshal(m, b)
}
func (m *HealthCheckEjectUnhealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Marshal(b, m, deterministic)
}
func (m *HealthCheckEjectUnhealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEjectUnhealthy.Merge(m, src)
}
func (m *HealthCheckEjectUnhealthy) XXX_Size() int {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Size(m)
}
func (m *HealthCheckEjectUnhealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEjectUnhealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEjectUnhealthy proto.InternalMessageInfo

func (m *HealthCheckEjectUnhealthy) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

type HealthCheckAddHealthy struct {
	FirstCheck           bool     `protobuf:"varint,1,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckAddHealthy) Reset()         { *m = HealthCheckAddHealthy{} }
func (m *HealthCheckAddHealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckAddHealthy) ProtoMessage()    {}
func (*HealthCheckAddHealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{2}
}

func (m *HealthCheckAddHealthy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckAddHealthy.Unmarshal(m, b)
}
func (m *HealthCheckAddHealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckAddHealthy.Marshal(b, m, deterministic)
}
func (m *HealthCheckAddHealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckAddHealthy.Merge(m, src)
}
func (m *HealthCheckAddHealthy) XXX_Size() int {
	return xxx_messageInfo_HealthCheckAddHealthy.Size(m)
}
func (m *HealthCheckAddHealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckAddHealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckAddHealthy proto.InternalMessageInfo

func (m *HealthCheckAddHealthy) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

type HealthCheckFailure struct {
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v3.HealthCheckFailureType" json:"failure_type,omitempty"`
	FirstCheck           bool                   `protobuf:"varint,2,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckFailure) Reset()         { *m = HealthCheckFailure{} }
func (m *HealthCheckFailure) String() string { return proto.CompactTextString(m) }
func (*HealthCheckFailure) ProtoMessage()    {}
func (*HealthCheckFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{3}
}

func (m *HealthCheckFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckFailure.Unmarshal(m, b)
}
func (m *HealthCheckFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckFailure.Marshal(b, m, deterministic)
}
func (m *HealthCheckFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckFailure.Merge(m, src)
}
func (m *HealthCheckFailure) XXX_Size() int {
	return xxx_messageInfo_HealthCheckFailure.Size(m)
}
func (m *HealthCheckFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckFailure.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckFailure proto.InternalMessageInfo

func (m *HealthCheckFailure) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

func (m *HealthCheckFailure) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

type DegradedHealthyHost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DegradedHealthyHost) Reset()         { *m = DegradedHealthyHost{} }
func (m *DegradedHealthyHost) String() string { return proto.CompactTextString(m) }
func (*DegradedHealthyHost) ProtoMessage()    {}
func (*DegradedHealthyHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{4}
}

func (m *DegradedHealthyHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DegradedHealthyHost.Unmarshal(m, b)
}
func (m *DegradedHealthyHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DegradedHealthyHost.Marshal(b, m, deterministic)
}
func (m *DegradedHealthyHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DegradedHealthyHost.Merge(m, src)
}
func (m *DegradedHealthyHost) XXX_Size() int {
	return xxx_messageInfo_DegradedHealthyHost.Size(m)
}
func (m *DegradedHealthyHost) XXX_DiscardUnknown() {
	xxx_messageInfo_DegradedHealthyHost.DiscardUnknown(m)
}

var xxx_messageInfo_DegradedHealthyHost proto.InternalMessageInfo

type NoLongerDegradedHost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoLongerDegradedHost) Reset()         { *m = NoLongerDegradedHost{} }
func (m *NoLongerDegradedHost) String() string { return proto.CompactTextString(m) }
func (*NoLongerDegradedHost) ProtoMessage()    {}
func (*NoLongerDegradedHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{5}
}

func (m *NoLongerDegradedHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoLongerDegradedHost.Unmarshal(m, b)
}
func (m *NoLongerDegradedHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoLongerDegradedHost.Marshal(b, m, deterministic)
}
func (m *NoLongerDegradedHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoLongerDegradedHost.Merge(m, src)
}
func (m *NoLongerDegradedHost) XXX_Size() int {
	return xxx_messageInfo_NoLongerDegradedHost.Size(m)
}
func (m *NoLongerDegradedHost) XXX_DiscardUnknown() {
	xxx_messageInfo_NoLongerDegradedHost.DiscardUnknown(m)
}

var xxx_messageInfo_NoLongerDegradedHost proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.data.core.v3.HealthCheckFailureType", HealthCheckFailureType_name, HealthCheckFailureType_value)
	proto.RegisterEnum("envoy.data.core.v3.HealthCheckerType", HealthCheckerType_name, HealthCheckerType_value)
	proto.RegisterType((*HealthCheckEvent)(nil), "envoy.data.core.v3.HealthCheckEvent")
	proto.RegisterType((*HealthCheckEjectUnhealthy)(nil), "envoy.data.core.v3.HealthCheckEjectUnhealthy")
	proto.RegisterType((*HealthCheckAddHealthy)(nil), "envoy.data.core.v3.HealthCheckAddHealthy")
	proto.RegisterType((*HealthCheckFailure)(nil), "envoy.data.core.v3.HealthCheckFailure")
	proto.RegisterType((*DegradedHealthyHost)(nil), "envoy.data.core.v3.DegradedHealthyHost")
	proto.RegisterType((*NoLongerDegradedHost)(nil), "envoy.data.core.v3.NoLongerDegradedHost")
}

func init() {
	proto.RegisterFile("envoy/data/core/v3/health_check_event.proto", fileDescriptor_b418bf0decf4b3ef)
}

var fileDescriptor_b418bf0decf4b3ef = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xe2, 0x46,
	0x18, 0xc6, 0xe6, 0xfb, 0x25, 0x6a, 0x9d, 0x49, 0x68, 0x28, 0x52, 0x1a, 0x8a, 0xd4, 0x96, 0xd2,
	0xc4, 0x56, 0xa0, 0x87, 0x88, 0x4a, 0x95, 0x30, 0xa1, 0x25, 0x4a, 0x95, 0x22, 0x87, 0x36, 0xa7,
	0xca, 0x9a, 0xe0, 0x01, 0xdc, 0x82, 0x07, 0xd9, 0x03, 0x2a, 0xb7, 0xaa, 0xa7, 0xfe, 0x86, 0xfe,
	0x88, 0x3d, 0xec, 0x71, 0x6f, 0x7b, 0x58, 0x69, 0xaf, 0xfb, 0x6f, 0x56, 0x39, 0xad, 0x66, 0x6c,
	0x20, 0x60, 0x47, 0xe1, 0xb2, 0x37, 0xe6, 0xfd, 0x78, 0x9e, 0xf7, 0x79, 0x3f, 0x30, 0x7c, 0x47,
	0x9c, 0x39, 0x5d, 0x68, 0x16, 0x66, 0x58, 0xeb, 0x53, 0x97, 0x68, 0xf3, 0xba, 0x36, 0x22, 0x78,
	0xcc, 0x46, 0x66, 0x7f, 0x44, 0xfa, 0x7f, 0x99, 0x64, 0x4e, 0x1c, 0xa6, 0x4e, 0x5d, 0xca, 0x28,
	0x42, 0x22, 0x58, 0xe5, 0xc1, 0x2a, 0x0f, 0x56, 0xe7, 0xf5, 0x62, 0xd9, 0x07, 0xe8, 0x53, 0x67,
	0x60, 0x0f, 0x57, 0x10, 0xd8, 0xb2, 0x5c, 0xe2, 0x79, 0x7e, 0x5e, 0xf1, 0x64, 0x48, 0xe9, 0x70,
	0x4c, 0x34, 0xf1, 0xba, 0x9f, 0x0d, 0x34, 0x66, 0x4f, 0x88, 0xc7, 0xf0, 0x64, 0x1a, 0x04, 0x1c,
	0xcf, 0xac, 0x29, 0xd6, 0xb0, 0xe3, 0x50, 0x86, 0x99, 0x4d, 0x1d, 0x4f, 0xf3, 0x18, 0x66, 0xb3,
	0x65, 0xfe, 0x97, 0x21, 0xf7, 0x9c, 0xb8, 0x9e, 0x4d, 0x1d, 0xdb, 0x19, 0x06, 0x21, 0x47, 0x73,
	0x3c, 0xb6, 0x2d, 0xcc, 0x88, 0xb6, 0xfc, 0xe1, 0x3b, 0xca, 0x2f, 0x53, 0xa0, 0x74, 0x84, 0xa0,
	0x16, 0xd7, 0xd3, 0xe6, 0x72, 0x90, 0x09, 0x07, 0x8f, 0x45, 0x12, 0xd7, 0x64, 0x8b, 0x29, 0x29,
	0x48, 0x25, 0xa9, 0xf2, 0x49, 0xed, 0x2b, 0x35, 0x2c, 0x53, 0x7d, 0x04, 0x41, 0xdc, 0xde, 0x62,
	0x4a, 0xf4, 0xcc, 0x83, 0x9e, 0xfc, 0x57, 0x92, 0x15, 0xc9, 0xd8, 0x1f, 0x6d, 0x3b, 0xd1, 0x39,
	0x24, 0x46, 0xd4, 0x63, 0x05, 0xb9, 0x24, 0x55, 0x72, 0xb5, 0xe3, 0x00, 0xd1, 0x6f, 0xd2, 0x0a,
	0xb3, 0xe9, 0x37, 0xc9, 0x10, 0xa1, 0xa8, 0x0a, 0x7b, 0xfd, 0xf1, 0xcc, 0x63, 0xc4, 0x35, 0x1d,
	0x3c, 0x21, 0x85, 0x78, 0x49, 0xaa, 0x64, 0xf5, 0xf4, 0x83, 0x9e, 0x70, 0xe5, 0x92, 0x64, 0xe4,
	0x02, 0xe7, 0x0d, 0x9e, 0x10, 0xd4, 0x87, 0x3c, 0xf9, 0x93, 0xf4, 0x99, 0x39, 0x73, 0x7c, 0xee,
	0x85, 0x3f, 0xa7, 0x42, 0x42, 0xf0, 0x9d, 0x3d, 0xa3, 0xa0, 0xcd, 0x73, 0x7f, 0x5b, 0xa6, 0x76,
	0x62, 0xc6, 0x01, 0xd9, 0xb0, 0xf8, 0x4d, 0xba, 0x83, 0x7d, 0x6c, 0x59, 0xe6, 0x26, 0x41, 0x52,
	0x10, 0x7c, 0xfb, 0x0c, 0x41, 0xd3, 0xb2, 0x3a, 0x2b, 0xf0, 0x4f, 0xf1, 0xea, 0xe5, 0x03, 0x13,
	0x28, 0x6e, 0xac, 0xd8, 0x00, 0xdb, 0xe3, 0x99, 0x4b, 0x02, 0x86, 0xb4, 0x60, 0xf8, 0xfa, 0x19,
	0x86, 0x9f, 0xfc, 0x9c, 0x4e, 0xcc, 0x38, 0x1a, 0x85, 0xac, 0x3e, 0xcd, 0x1f, 0x90, 0xb7, 0xc8,
	0xd0, 0xc5, 0x16, 0x59, 0x8b, 0x10, 0x43, 0xc9, 0x08, 0x86, 0x6f, 0xa2, 0x18, 0x2e, 0x83, 0x84,
	0x65, 0xf5, 0xd4, 0x63, 0xbc, 0x3d, 0x56, 0xd8, 0x8c, 0x30, 0x1c, 0x39, 0xd4, 0x1c, 0x53, 0x67,
	0x48, 0x5c, 0x73, 0x4d, 0xc4, 0x09, 0xb2, 0x82, 0xa0, 0x12, 0x45, 0x70, 0x43, 0x7f, 0x11, 0x19,
	0x2b, 0x22, 0x9f, 0xe1, 0xd0, 0x89, 0xb0, 0xa3, 0x0b, 0xc8, 0xae, 0x2e, 0xa5, 0x90, 0x12, 0xa0,
	0x45, 0xd5, 0xbf, 0x25, 0x75, 0x79, 0x4b, 0x6a, 0x6f, 0x19, 0x61, 0xac, 0x83, 0x1b, 0xda, 0xff,
	0x6f, 0xfe, 0xfb, 0xa2, 0x0a, 0x95, 0x50, 0x05, 0x35, 0x3c, 0x9e, 0x8e, 0xb0, 0xba, 0x7d, 0x11,
	0xfa, 0x1e, 0x24, 0x45, 0xfb, 0x51, 0xfc, 0xbd, 0x2e, 0x95, 0x5f, 0x48, 0xf0, 0xf9, 0x93, 0xfb,
	0x82, 0xee, 0x60, 0x6f, 0x39, 0xb2, 0x47, 0x67, 0x53, 0xdd, 0x6d, 0x62, 0x5b, 0xb7, 0x93, 0x1b,
	0xac, 0xcd, 0x8d, 0x0b, 0x5e, 0x75, 0x1d, 0xce, 0x77, 0xa9, 0x7a, 0xa3, 0xa4, 0xb2, 0x03, 0xf9,
	0xc8, 0xf5, 0x43, 0x27, 0x90, 0x1b, 0xd8, 0xae, 0xc7, 0xfc, 0x55, 0x13, 0xa5, 0x66, 0x0c, 0x10,
	0x26, 0x11, 0xda, 0xf8, 0x9e, 0x73, 0x6a, 0x70, 0xb6, 0x03, 0xe7, 0x1a, 0xb6, 0xfc, 0x5a, 0x02,
	0x14, 0xd6, 0xf6, 0xd1, 0x3a, 0xb3, 0x2d, 0x43, 0x0e, 0xc9, 0x38, 0xe7, 0x32, 0x4e, 0xa1, 0xba,
	0x83, 0x8c, 0x80, 0xae, 0x7c, 0x05, 0x07, 0x11, 0xeb, 0xde, 0xa8, 0x71, 0xa4, 0xb3, 0xe0, 0xc3,
	0x10, 0x81, 0x14, 0x91, 0x53, 0xbe, 0x86, 0xc3, 0xa8, 0xc5, 0x6e, 0xd4, 0x39, 0x96, 0x0a, 0xa7,
	0x4f, 0x61, 0x45, 0x25, 0x55, 0x7f, 0x84, 0xcf, 0xa2, 0x9b, 0x83, 0x00, 0x52, 0xcd, 0x56, 0xef,
	0xea, 0xf7, 0xb6, 0x12, 0x43, 0x39, 0x48, 0x77, 0x9b, 0xb7, 0xb7, 0xfc, 0x21, 0xf1, 0xc7, 0x4d,
	0xbb, 0x77, 0xf7, 0xab, 0x71, 0xad, 0xc8, 0xd5, 0x1f, 0x60, 0x3f, 0xf4, 0x6f, 0x8d, 0x32, 0x90,
	0xe8, 0xf4, 0x7a, 0x5d, 0x25, 0x86, 0xd2, 0x10, 0xef, 0xb5, 0xba, 0x8a, 0xc4, 0x4d, 0x3f, 0x1b,
	0xdd, 0x96, 0x22, 0xa3, 0x2c, 0x24, 0x8d, 0xf6, 0xe5, 0xd5, 0xad, 0x12, 0xd7, 0x5b, 0xaf, 0xfe,
	0x79, 0xfb, 0x2e, 0x25, 0x2b, 0x32, 0x94, 0x6c, 0xea, 0xcf, 0x6d, 0xea, 0xd2, 0xbf, 0x17, 0x11,
	0x23, 0xd4, 0xf3, 0xdb, 0x57, 0xd4, 0xe5, 0x37, 0xd9, 0x95, 0xee, 0x53, 0xe2, 0x38, 0xeb, 0x1f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0x2e, 0x8e, 0xbb, 0x5d, 0x07, 0x00, 0x00,
}