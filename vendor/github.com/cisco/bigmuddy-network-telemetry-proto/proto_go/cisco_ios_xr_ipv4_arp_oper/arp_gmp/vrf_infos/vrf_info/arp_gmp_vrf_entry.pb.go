// Code generated by protoc-gen-go.
// source: arp_gmp_vrf_entry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_arp_oper_arp_gmp_vrf_infos_vrf_info is a generated protocol buffer package.

It is generated from these files:
	arp_gmp_vrf_entry.proto

It has these top-level messages:
	ArpGmpVrfEntry_KEYS
	ArpGmpVrfEntry
*/
package cisco_ios_xr_ipv4_arp_oper_arp_gmp_vrf_infos_vrf_info

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

// ARP GMP VRF entry
type ArpGmpVrfEntry_KEYS struct {
	VrfName string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *ArpGmpVrfEntry_KEYS) Reset()                    { *m = ArpGmpVrfEntry_KEYS{} }
func (m *ArpGmpVrfEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*ArpGmpVrfEntry_KEYS) ProtoMessage()               {}
func (*ArpGmpVrfEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArpGmpVrfEntry_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type ArpGmpVrfEntry struct {
	// VRF Name
	VrfName string `protobuf:"bytes,50,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// VRF ID
	VrfIdNumber uint32 `protobuf:"varint,51,opt,name=vrf_id_number,json=vrfIdNumber" json:"vrf_id_number,omitempty"`
	// IPv4 unicast table ID
	TableId uint32 `protobuf:"varint,52,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	// RSI registration handle
	RsiHandle uint32 `protobuf:"varint,53,opt,name=rsi_handle,json=rsiHandle" json:"rsi_handle,omitempty"`
	// RSI registration handle (top 32-bits)
	RsiHandleHigh uint32 `protobuf:"varint,54,opt,name=rsi_handle_high,json=rsiHandleHigh" json:"rsi_handle_high,omitempty"`
}

func (m *ArpGmpVrfEntry) Reset()                    { *m = ArpGmpVrfEntry{} }
func (m *ArpGmpVrfEntry) String() string            { return proto.CompactTextString(m) }
func (*ArpGmpVrfEntry) ProtoMessage()               {}
func (*ArpGmpVrfEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArpGmpVrfEntry) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *ArpGmpVrfEntry) GetVrfIdNumber() uint32 {
	if m != nil {
		return m.VrfIdNumber
	}
	return 0
}

func (m *ArpGmpVrfEntry) GetTableId() uint32 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *ArpGmpVrfEntry) GetRsiHandle() uint32 {
	if m != nil {
		return m.RsiHandle
	}
	return 0
}

func (m *ArpGmpVrfEntry) GetRsiHandleHigh() uint32 {
	if m != nil {
		return m.RsiHandleHigh
	}
	return 0
}

func init() {
	proto.RegisterType((*ArpGmpVrfEntry_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp_gmp.vrf_infos.vrf_info.arp_gmp_vrf_entry_KEYS")
	proto.RegisterType((*ArpGmpVrfEntry)(nil), "cisco_ios_xr_ipv4_arp_oper.arp_gmp.vrf_infos.vrf_info.arp_gmp_vrf_entry")
}

func init() { proto.RegisterFile("arp_gmp_vrf_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0xf1, 0xd2, 0x34, 0x2a, 0xa6, 0x54, 0x43, 0xab, 0x0e, 0x85, 0xe0, 0xa1, 0x64, 0xf2,
	0x50, 0x27, 0x7d, 0x83, 0x42, 0x42, 0x21, 0x43, 0x3a, 0x75, 0x3a, 0xe4, 0x48, 0xb6, 0x0f, 0x62,
	0x49, 0x9c, 0x5c, 0xd3, 0xbe, 0x57, 0x1f, 0xb0, 0xe8, 0x28, 0x09, 0xc1, 0x9b, 0xf4, 0xff, 0xdf,
	0x77, 0x07, 0x27, 0x1e, 0x34, 0x05, 0x68, 0xfb, 0x00, 0x23, 0x35, 0x60, 0xdd, 0x40, 0x3f, 0x65,
	0x20, 0x3f, 0x78, 0xb9, 0x3e, 0x60, 0x3c, 0x78, 0x40, 0x1f, 0xe1, 0x9b, 0x00, 0xc3, 0xb8, 0x82,
	0x84, 0xfa, 0x60, 0xa9, 0xfc, 0x77, 0xca, 0xe4, 0xa0, 0x6b, 0x7c, 0x3c, 0xbd, 0x8a, 0x4a, 0xdc,
	0x4f, 0x26, 0xc2, 0xfb, 0xdb, 0xe7, 0x87, 0x7c, 0x14, 0xd7, 0x29, 0x71, 0xba, 0xb7, 0x2a, 0x5b,
	0x64, 0xcb, 0xf9, 0x7e, 0x36, 0x52, 0xb3, 0xd3, 0xbd, 0x2d, 0x7e, 0x33, 0x71, 0x37, 0xb1, 0x2e,
	0x84, 0x97, 0x0b, 0x41, 0x16, 0x22, 0xe7, 0x8d, 0x06, 0xdc, 0x57, 0x5f, 0x5b, 0x52, 0xd5, 0x22,
	0x5b, 0xe6, 0xfb, 0x9b, 0x91, 0x9a, 0xad, 0xd9, 0x71, 0x94, 0xf4, 0x41, 0xd7, 0x47, 0x0b, 0x68,
	0xd4, 0x8a, 0xeb, 0x19, 0xff, 0xb7, 0x46, 0x3e, 0x09, 0x41, 0x11, 0xa1, 0xd3, 0xce, 0x1c, 0xad,
	0x5a, 0x73, 0x39, 0xa7, 0x88, 0x1b, 0x0e, 0xe4, 0xb3, 0xb8, 0x3d, 0xd7, 0xd0, 0x61, 0xdb, 0xa9,
	0x57, 0x66, 0xf2, 0x13, 0xb3, 0xc1, 0xb6, 0xab, 0xaf, 0xf8, 0x52, 0xd5, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9c, 0xb7, 0x22, 0x25, 0x44, 0x01, 0x00, 0x00,
}