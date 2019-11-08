// Code generated by protoc-gen-go. DO NOT EDIT.
// source: judge.proto

package protobuf

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

type JudgeRequest struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Tid                  string   `protobuf:"bytes,2,opt,name=tid,proto3" json:"tid,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Language             string   `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Code                 []byte   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JudgeRequest) Reset()         { *m = JudgeRequest{} }
func (m *JudgeRequest) String() string { return proto.CompactTextString(m) }
func (*JudgeRequest) ProtoMessage()    {}
func (*JudgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{0}
}

func (m *JudgeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeRequest.Unmarshal(m, b)
}
func (m *JudgeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeRequest.Marshal(b, m, deterministic)
}
func (m *JudgeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeRequest.Merge(m, src)
}
func (m *JudgeRequest) XXX_Size() int {
	return xxx_messageInfo_JudgeRequest.Size(m)
}
func (m *JudgeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeRequest proto.InternalMessageInfo

func (m *JudgeRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *JudgeRequest) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

func (m *JudgeRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JudgeRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *JudgeRequest) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type JudgeCaseResult struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SpaceUsed            uint32   `protobuf:"varint,2,opt,name=space_used,json=spaceUsed,proto3" json:"space_used,omitempty"`
	TimeUsed             uint32   `protobuf:"varint,3,opt,name=time_used,json=timeUsed,proto3" json:"time_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JudgeCaseResult) Reset()         { *m = JudgeCaseResult{} }
func (m *JudgeCaseResult) String() string { return proto.CompactTextString(m) }
func (*JudgeCaseResult) ProtoMessage()    {}
func (*JudgeCaseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{1}
}

func (m *JudgeCaseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeCaseResult.Unmarshal(m, b)
}
func (m *JudgeCaseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeCaseResult.Marshal(b, m, deterministic)
}
func (m *JudgeCaseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeCaseResult.Merge(m, src)
}
func (m *JudgeCaseResult) XXX_Size() int {
	return xxx_messageInfo_JudgeCaseResult.Size(m)
}
func (m *JudgeCaseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeCaseResult.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeCaseResult proto.InternalMessageInfo

func (m *JudgeCaseResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *JudgeCaseResult) GetSpaceUsed() uint32 {
	if m != nil {
		return m.SpaceUsed
	}
	return 0
}

func (m *JudgeCaseResult) GetTimeUsed() uint32 {
	if m != nil {
		return m.TimeUsed
	}
	return 0
}

type JudgeResponse struct {
	Sid                  string             `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Result               []*JudgeCaseResult `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JudgeResponse) Reset()         { *m = JudgeResponse{} }
func (m *JudgeResponse) String() string { return proto.CompactTextString(m) }
func (*JudgeResponse) ProtoMessage()    {}
func (*JudgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{2}
}

func (m *JudgeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeResponse.Unmarshal(m, b)
}
func (m *JudgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeResponse.Marshal(b, m, deterministic)
}
func (m *JudgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeResponse.Merge(m, src)
}
func (m *JudgeResponse) XXX_Size() int {
	return xxx_messageInfo_JudgeResponse.Size(m)
}
func (m *JudgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeResponse proto.InternalMessageInfo

func (m *JudgeResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *JudgeResponse) GetResult() []*JudgeCaseResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type Case struct {
	Stdin                []byte   `protobuf:"bytes,1,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout               []byte   `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Case) Reset()         { *m = Case{} }
func (m *Case) String() string { return proto.CompactTextString(m) }
func (*Case) ProtoMessage()    {}
func (*Case) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{3}
}

func (m *Case) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Case.Unmarshal(m, b)
}
func (m *Case) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Case.Marshal(b, m, deterministic)
}
func (m *Case) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Case.Merge(m, src)
}
func (m *Case) XXX_Size() int {
	return xxx_messageInfo_Case.Size(m)
}
func (m *Case) XXX_DiscardUnknown() {
	xxx_messageInfo_Case.DiscardUnknown(m)
}

var xxx_messageInfo_Case proto.InternalMessageInfo

func (m *Case) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *Case) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

type TestCaseRequest struct {
	Tid                  string   `protobuf:"bytes,1,opt,name=tid,proto3" json:"tid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestCaseRequest) Reset()         { *m = TestCaseRequest{} }
func (m *TestCaseRequest) String() string { return proto.CompactTextString(m) }
func (*TestCaseRequest) ProtoMessage()    {}
func (*TestCaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{4}
}

func (m *TestCaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestCaseRequest.Unmarshal(m, b)
}
func (m *TestCaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestCaseRequest.Marshal(b, m, deterministic)
}
func (m *TestCaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCaseRequest.Merge(m, src)
}
func (m *TestCaseRequest) XXX_Size() int {
	return xxx_messageInfo_TestCaseRequest.Size(m)
}
func (m *TestCaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestCaseRequest proto.InternalMessageInfo

func (m *TestCaseRequest) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

type TestCaseResponse struct {
	Tid                  string   `protobuf:"bytes,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Case                 []*Case  `protobuf:"bytes,2,rep,name=case,proto3" json:"case,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestCaseResponse) Reset()         { *m = TestCaseResponse{} }
func (m *TestCaseResponse) String() string { return proto.CompactTextString(m) }
func (*TestCaseResponse) ProtoMessage()    {}
func (*TestCaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b4fa635fb2baaa, []int{5}
}

func (m *TestCaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestCaseResponse.Unmarshal(m, b)
}
func (m *TestCaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestCaseResponse.Marshal(b, m, deterministic)
}
func (m *TestCaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCaseResponse.Merge(m, src)
}
func (m *TestCaseResponse) XXX_Size() int {
	return xxx_messageInfo_TestCaseResponse.Size(m)
}
func (m *TestCaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestCaseResponse proto.InternalMessageInfo

func (m *TestCaseResponse) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

func (m *TestCaseResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TestCaseResponse) GetCase() []*Case {
	if m != nil {
		return m.Case
	}
	return nil
}

func init() {
	proto.RegisterType((*JudgeRequest)(nil), "protobuf.JudgeRequest")
	proto.RegisterType((*JudgeCaseResult)(nil), "protobuf.JudgeCaseResult")
	proto.RegisterType((*JudgeResponse)(nil), "protobuf.JudgeResponse")
	proto.RegisterType((*Case)(nil), "protobuf.Case")
	proto.RegisterType((*TestCaseRequest)(nil), "protobuf.TestCaseRequest")
	proto.RegisterType((*TestCaseResponse)(nil), "protobuf.TestCaseResponse")
}

func init() { proto.RegisterFile("judge.proto", fileDescriptor_a2b4fa635fb2baaa) }

var fileDescriptor_a2b4fa635fb2baaa = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x5f, 0x6b, 0xe2, 0x40,
	0x14, 0xc5, 0xd1, 0x44, 0x37, 0x5e, 0xe3, 0x1f, 0x86, 0x65, 0x37, 0x9b, 0x65, 0x41, 0xb2, 0x2f,
	0x3e, 0x09, 0xb5, 0xfd, 0x02, 0xa5, 0x05, 0xc1, 0xa7, 0x32, 0xd8, 0xe7, 0x12, 0x9d, 0x5b, 0x49,
	0xb1, 0x89, 0xf5, 0xde, 0xe9, 0x53, 0x3f, 0x7c, 0xc9, 0xcd, 0xc4, 0xb4, 0x56, 0xfa, 0x94, 0x39,
	0xe7, 0x0c, 0x37, 0xbf, 0x73, 0x07, 0xfa, 0x4f, 0xd6, 0x6c, 0x71, 0xb6, 0x3f, 0x14, 0x5c, 0xa8,
	0x40, 0x3e, 0x6b, 0xfb, 0x98, 0xbc, 0x41, 0xb8, 0x2c, 0x03, 0x8d, 0x2f, 0x16, 0x89, 0xd5, 0x18,
	0x3c, 0xca, 0x4c, 0xd4, 0x9a, 0xb4, 0xa6, 0x3d, 0x5d, 0x1e, 0x4b, 0x87, 0x33, 0x13, 0xb5, 0x2b,
	0x87, 0x33, 0xa3, 0x22, 0xf8, 0xf1, 0x8a, 0x07, 0xca, 0x8a, 0x3c, 0xf2, 0xc4, 0xad, 0xa5, 0x8a,
	0x21, 0xd8, 0xa5, 0xf9, 0xd6, 0xa6, 0x5b, 0x8c, 0x7c, 0x89, 0x8e, 0x5a, 0x29, 0xf0, 0x37, 0x85,
	0xc1, 0xa8, 0x33, 0x69, 0x4d, 0x43, 0x2d, 0xe7, 0x04, 0x61, 0x24, 0x7f, 0xbf, 0x49, 0x09, 0x35,
	0x92, 0xdd, 0xb1, 0xfa, 0x05, 0x5d, 0xe2, 0x94, 0x2d, 0x39, 0x06, 0xa7, 0xd4, 0x3f, 0x00, 0xda,
	0xa7, 0x1b, 0x7c, 0xb0, 0x84, 0x15, 0xcd, 0x40, 0xf7, 0xc4, 0xb9, 0x27, 0x34, 0xea, 0x2f, 0xf4,
	0x38, 0x7b, 0x76, 0xa9, 0x27, 0x69, 0x50, 0x1a, 0x65, 0x98, 0xac, 0x60, 0xe0, 0x4a, 0xd2, 0xbe,
	0xc8, 0x09, 0xcf, 0xb4, 0xbc, 0x80, 0xee, 0x41, 0x00, 0xa2, 0xf6, 0xc4, 0x9b, 0xf6, 0xe7, 0x7f,
	0x66, 0xf5, 0x8a, 0x66, 0x27, 0x84, 0xda, 0x5d, 0x4c, 0xae, 0xc0, 0x2f, 0x5d, 0xf5, 0x13, 0x3a,
	0xc4, 0x26, 0xcb, 0x65, 0x5c, 0xa8, 0x2b, 0x51, 0xf5, 0x30, 0x85, 0x65, 0x61, 0x0d, 0xb5, 0x53,
	0xc9, 0x7f, 0x18, 0xad, 0x90, 0xb8, 0x9a, 0x77, 0xdc, 0x39, 0x37, 0x34, 0x9c, 0x99, 0x64, 0x0d,
	0xe3, 0xe6, 0x52, 0xc3, 0xfc, 0xf9, 0xd6, 0x37, 0xef, 0x90, 0x80, 0xbf, 0x49, 0x09, 0x5d, 0x97,
	0x61, 0xd3, 0x45, 0x26, 0x4a, 0x36, 0x5f, 0x42, 0x47, 0x9a, 0xa9, 0x6b, 0x18, 0x2e, 0x90, 0xeb,
	0x05, 0xc9, 0x1b, 0x9c, 0x94, 0x77, 0xa0, 0xf1, 0xef, 0x2f, 0x7e, 0xc5, 0x36, 0xbf, 0x83, 0xa0,
	0xe6, 0x55, 0xb7, 0xd0, 0x5f, 0x20, 0x1f, 0xe5, 0x87, 0x45, 0x9e, 0xf4, 0x8e, 0xe3, 0x73, 0x51,
	0x35, 0x71, 0xdd, 0x95, 0xe8, 0xf2, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x81, 0xc1, 0xf4, 0xb7,
	0x02, 0x00, 0x00,
}
