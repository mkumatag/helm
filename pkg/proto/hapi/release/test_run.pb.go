// Code generated by protoc-gen-go.
// source: hapi/release/test_run.proto
// DO NOT EDIT!

package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestRun_Status int32

const (
	TestRun_UNKNOWN TestRun_Status = 0
	TestRun_SUCCESS TestRun_Status = 1
	TestRun_FAILURE TestRun_Status = 2
)

var TestRun_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAILURE",
}
var TestRun_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"SUCCESS": 1,
	"FAILURE": 2,
}

func (x TestRun_Status) String() string {
	return proto.EnumName(TestRun_Status_name, int32(x))
}
func (TestRun_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

type TestRun struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      TestRun_Status             `protobuf:"varint,2,opt,name=status,enum=hapi.release.TestRun_Status" json:"status,omitempty"`
	Info        string                     `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	StartedAt   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	CompletedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=completed_at,json=completedAt" json:"completed_at,omitempty"`
}

func (m *TestRun) Reset()                    { *m = TestRun{} }
func (m *TestRun) String() string            { return proto.CompactTextString(m) }
func (*TestRun) ProtoMessage()               {}
func (*TestRun) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *TestRun) GetStartedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *TestRun) GetCompletedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CompletedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*TestRun)(nil), "hapi.release.TestRun")
	proto.RegisterEnum("hapi.release.TestRun_Status", TestRun_Status_name, TestRun_Status_value)
}

func init() { proto.RegisterFile("hapi/release/test_run.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0xac, 0x09, 0x99, 0x14, 0x09, 0x7b, 0x0a, 0x55, 0x50, 0x7a, 0xf2, 0xb4, 0x0b,
	0xd5, 0x8b, 0x07, 0x0f, 0xb1, 0x54, 0x10, 0x25, 0xc2, 0xa6, 0x41, 0xf0, 0x52, 0xb6, 0xba, 0xad,
	0x81, 0x24, 0x1b, 0x92, 0xc9, 0x1f, 0xf1, 0x17, 0xbb, 0x9b, 0x6c, 0xc5, 0x5b, 0x6f, 0x33, 0xf3,
	0xbe, 0xf7, 0x78, 0x03, 0x17, 0xdf, 0xa2, 0x29, 0x58, 0x2b, 0x4b, 0x29, 0x3a, 0xc9, 0x50, 0x76,
	0xb8, 0x69, 0xfb, 0x9a, 0x36, 0xad, 0x42, 0x45, 0xa6, 0x46, 0xa4, 0x56, 0x9c, 0x5d, 0xed, 0x95,
	0xda, 0x97, 0x92, 0x0d, 0xda, 0xb6, 0xdf, 0x31, 0x2c, 0x2a, 0xcd, 0x8b, 0xaa, 0x19, 0xf1, 0xf9,
	0x8f, 0x0b, 0xfe, 0x5a, 0x5f, 0x78, 0x5f, 0x13, 0x02, 0x93, 0x5a, 0x54, 0x32, 0x76, 0xae, 0x9d,
	0x9b, 0x80, 0x0f, 0x33, 0xb9, 0x03, 0x4f, 0xe3, 0xd8, 0x77, 0xb1, 0xab, 0xaf, 0xe7, 0x8b, 0x4b,
	0xfa, 0x3f, 0x9f, 0x5a, 0x2b, 0xcd, 0x06, 0x86, 0x5b, 0xd6, 0x24, 0x15, 0xf5, 0x4e, 0xc5, 0xa7,
	0x63, 0x92, 0x99, 0xc9, 0x3d, 0x80, 0x56, 0x5b, 0x94, 0x5f, 0x1b, 0x81, 0xf1, 0x44, 0x2b, 0xe1,
	0x62, 0x46, 0xc7, 0x7e, 0xf4, 0xd0, 0x8f, 0xae, 0x0f, 0xfd, 0x78, 0x60, 0xe9, 0x04, 0xc9, 0x03,
	0x4c, 0x3f, 0x55, 0xd5, 0x94, 0xd2, 0x9a, 0xcf, 0x8e, 0x9a, 0xc3, 0x3f, 0x3e, 0xc1, 0x39, 0x03,
	0x6f, 0xec, 0x47, 0x42, 0xf0, 0xf3, 0xf4, 0x25, 0x7d, 0x7b, 0x4f, 0xa3, 0x13, 0xb3, 0x64, 0xf9,
	0x72, 0xb9, 0xca, 0xb2, 0xc8, 0x31, 0xcb, 0x53, 0xf2, 0xfc, 0x9a, 0xf3, 0x55, 0xe4, 0x3e, 0x06,
	0x1f, 0xbe, 0x7d, 0x70, 0xeb, 0x0d, 0xe1, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xb9,
	0xce, 0x57, 0x74, 0x01, 0x00, 0x00,
}
