// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ModSecurity struct {
	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Global rule sets for the current http connection manager
	RuleSets             []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModSecurity) Reset()         { *m = ModSecurity{} }
func (m *ModSecurity) String() string { return proto.CompactTextString(m) }
func (*ModSecurity) ProtoMessage()    {}
func (*ModSecurity) Descriptor() ([]byte, []int) {
	return fileDescriptor_259d6c35da0f75a1, []int{0}
}
func (m *ModSecurity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModSecurity.Unmarshal(m, b)
}
func (m *ModSecurity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModSecurity.Marshal(b, m, deterministic)
}
func (m *ModSecurity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModSecurity.Merge(m, src)
}
func (m *ModSecurity) XXX_Size() int {
	return xxx_messageInfo_ModSecurity.Size(m)
}
func (m *ModSecurity) XXX_DiscardUnknown() {
	xxx_messageInfo_ModSecurity.DiscardUnknown(m)
}

var xxx_messageInfo_ModSecurity proto.InternalMessageInfo

func (m *ModSecurity) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ModSecurity) GetRuleSets() []*RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

//
//String options are not recommended unless they are relatively short as they will be sent over the wire quite often.
//
//Any files referenced by this proto should be mounted into the relevant envoy pod prior to use or
//the filter will fail to initialize and the configuration will be rejected
type RuleSet struct {
	// string of rules which are added directly
	RuleStr string `protobuf:"bytes,1,opt,name=rule_str,json=ruleStr,proto3" json:"rule_str,omitempty"`
	// array of files to include
	Files                []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleSet) Reset()         { *m = RuleSet{} }
func (m *RuleSet) String() string { return proto.CompactTextString(m) }
func (*RuleSet) ProtoMessage()    {}
func (*RuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_259d6c35da0f75a1, []int{1}
}
func (m *RuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleSet.Unmarshal(m, b)
}
func (m *RuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleSet.Marshal(b, m, deterministic)
}
func (m *RuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleSet.Merge(m, src)
}
func (m *RuleSet) XXX_Size() int {
	return xxx_messageInfo_RuleSet.Size(m)
}
func (m *RuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_RuleSet proto.InternalMessageInfo

func (m *RuleSet) GetRuleStr() string {
	if m != nil {
		return m.RuleStr
	}
	return ""
}

func (m *RuleSet) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

type ModSecurityPerRoute struct {
	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Overwite the global rules on this route
	RuleSets             []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModSecurityPerRoute) Reset()         { *m = ModSecurityPerRoute{} }
func (m *ModSecurityPerRoute) String() string { return proto.CompactTextString(m) }
func (*ModSecurityPerRoute) ProtoMessage()    {}
func (*ModSecurityPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_259d6c35da0f75a1, []int{2}
}
func (m *ModSecurityPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModSecurityPerRoute.Unmarshal(m, b)
}
func (m *ModSecurityPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModSecurityPerRoute.Marshal(b, m, deterministic)
}
func (m *ModSecurityPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModSecurityPerRoute.Merge(m, src)
}
func (m *ModSecurityPerRoute) XXX_Size() int {
	return xxx_messageInfo_ModSecurityPerRoute.Size(m)
}
func (m *ModSecurityPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ModSecurityPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ModSecurityPerRoute proto.InternalMessageInfo

func (m *ModSecurityPerRoute) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ModSecurityPerRoute) GetRuleSets() []*RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

func init() {
	proto.RegisterType((*ModSecurity)(nil), "envoy.config.filter.http.modsecurity.v2.ModSecurity")
	proto.RegisterType((*RuleSet)(nil), "envoy.config.filter.http.modsecurity.v2.RuleSet")
	proto.RegisterType((*ModSecurityPerRoute)(nil), "envoy.config.filter.http.modsecurity.v2.ModSecurityPerRoute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/waf/waf.proto", fileDescriptor_259d6c35da0f75a1)
}

var fileDescriptor_259d6c35da0f75a1 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0x54, 0x88, 0xa0, 0x89, 0xbb, 0x85, 0x0e, 0xa1, 0x03, 0x8a, 0xb2, 0x90, 0x05, 0x1b, 0x95,
	0x8d, 0x09, 0x31, 0x74, 0xab, 0x84, 0xdc, 0x8d, 0x05, 0xe5, 0xe3, 0xc5, 0x35, 0xb8, 0x79, 0x91,
	0xfd, 0x52, 0xda, 0x8d, 0x9f, 0xc3, 0xef, 0xe2, 0x97, 0xa0, 0x26, 0x15, 0x62, 0x41, 0xea, 0xd4,
	0xc1, 0xd2, 0xbb, 0xd3, 0x9d, 0xef, 0xa4, 0x63, 0x73, 0xa5, 0x69, 0xd5, 0x15, 0xbc, 0xc4, 0xb5,
	0x70, 0x68, 0xf0, 0x56, 0xa3, 0x50, 0x06, 0x51, 0xb4, 0x16, 0xdf, 0xa0, 0x24, 0x37, 0xa0, 0xbc,
	0xd5, 0x02, 0xb6, 0x04, 0xb6, 0xc9, 0x8d, 0x80, 0x66, 0x83, 0x3b, 0xf1, 0x91, 0xd7, 0xfb, 0xc7,
	0x5b, 0x8b, 0x84, 0xd1, 0x4d, 0x4f, 0xf2, 0x12, 0x9b, 0x5a, 0x2b, 0x5e, 0x6b, 0x43, 0x60, 0xf9,
	0x8a, 0xa8, 0xe5, 0x6b, 0xac, 0x1c, 0x94, 0x9d, 0xd5, 0xb4, 0xe3, 0x9b, 0xd9, 0x74, 0xa2, 0x50,
	0x61, 0xef, 0x11, 0xfb, 0x6b, 0xb0, 0xa7, 0x5b, 0x36, 0x5e, 0x60, 0xb5, 0x3c, 0xe8, 0xa2, 0x29,
	0x0b, 0x2a, 0xed, 0xf2, 0xc2, 0x40, 0x15, 0x7b, 0x89, 0x97, 0x05, 0xf2, 0x17, 0x47, 0x0b, 0x16,
	0xda, 0xce, 0xc0, 0xab, 0x03, 0x72, 0xf1, 0x59, 0xe2, 0x67, 0xe3, 0xd9, 0x1d, 0x3f, 0x32, 0x9d,
	0xcb, 0xce, 0xc0, 0x12, 0x48, 0x06, 0x76, 0x38, 0x5c, 0xfa, 0xc0, 0x46, 0x07, 0x32, 0xba, 0x62,
	0xc1, 0xf0, 0x33, 0xd9, 0x3e, 0x35, 0x94, 0xa3, 0x5e, 0x46, 0x36, 0x9a, 0xb0, 0xf3, 0x5a, 0x1b,
	0x70, 0xb1, 0x9f, 0xf8, 0x59, 0x28, 0x07, 0x90, 0x7e, 0x7a, 0xec, 0xf2, 0x4f, 0xed, 0x67, 0xb0,
	0x12, 0x3b, 0x82, 0x13, 0xd6, 0x7f, 0x9a, 0x7f, 0x7d, 0x5f, 0x7b, 0x2f, 0x8f, 0xc7, 0xad, 0xd8,
	0xbe, 0xab, 0x7f, 0x96, 0x2c, 0x2e, 0xfa, 0x1d, 0xee, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0x93, 0xac, 0x8e, 0x10, 0x02, 0x00, 0x00,
}

func (this *ModSecurity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModSecurity)
	if !ok {
		that2, ok := that.(ModSecurity)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RuleSet)
	if !ok {
		that2, ok := that.(RuleSet)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RuleStr != that1.RuleStr {
		return false
	}
	if len(this.Files) != len(that1.Files) {
		return false
	}
	for i := range this.Files {
		if this.Files[i] != that1.Files[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ModSecurityPerRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModSecurityPerRoute)
	if !ok {
		that2, ok := that.(ModSecurityPerRoute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}