// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lang/lang.proto

/*
Package lang is a generated protocol buffer package.

It is generated from these files:
	lang/lang.proto

It has these top-level messages:
*/
package lang

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

type T int32

const (
	T_undefined                        T = 0
	T_user_has_already_in_conversation T = 1
	T_conversation_closed              T = 2
	T_invalid_invite                   T = 3
	T_invalid_agent                    T = 4
	T_user_is_not_in_conversation      T = 5
	T_conversation_not_found           T = 6
	T_internal_error                   T = 30
	T_invalid_input                    T = 22
	T_invalid_form                     T = 20
	T_access_token_expired             T = 21
	T_invalid_credential               T = 13
	T_credential_not_set               T = 7
	T_wrong_account_in_credential      T = 8
	T_wrong_user_in_credential         T = 10
	T_access_deny                      T = 9
	T_unable_to_send_message           T = 11
	T_topic_is_empty                   T = 12
	T_invalid_json                     T = 15
	T_invalid_protobuf                 T = 16
	T_unable_to_lock                   T = 40
	T_empty                            T = 41
	T_wrong_type                       T = 42
	T_invalid_kafka_topic              T = 43
	T_database_error                   T = 44
	T_timeout                          T = 45
)

var T_name = map[int32]string{
	0:  "undefined",
	1:  "user_has_already_in_conversation",
	2:  "conversation_closed",
	3:  "invalid_invite",
	4:  "invalid_agent",
	5:  "user_is_not_in_conversation",
	6:  "conversation_not_found",
	30: "internal_error",
	22: "invalid_input",
	20: "invalid_form",
	21: "access_token_expired",
	13: "invalid_credential",
	7:  "credential_not_set",
	8:  "wrong_account_in_credential",
	10: "wrong_user_in_credential",
	9:  "access_deny",
	11: "unable_to_send_message",
	12: "topic_is_empty",
	15: "invalid_json",
	16: "invalid_protobuf",
	40: "unable_to_lock",
	41: "empty",
	42: "wrong_type",
	43: "invalid_kafka_topic",
	44: "database_error",
	45: "timeout",
}
var T_value = map[string]int32{
	"undefined":                        0,
	"user_has_already_in_conversation": 1,
	"conversation_closed":              2,
	"invalid_invite":                   3,
	"invalid_agent":                    4,
	"user_is_not_in_conversation":      5,
	"conversation_not_found":           6,
	"internal_error":                   30,
	"invalid_input":                    22,
	"invalid_form":                     20,
	"access_token_expired":             21,
	"invalid_credential":               13,
	"credential_not_set":               7,
	"wrong_account_in_credential":      8,
	"wrong_user_in_credential":         10,
	"access_deny":                      9,
	"unable_to_send_message":           11,
	"topic_is_empty":                   12,
	"invalid_json":                     15,
	"invalid_protobuf":                 16,
	"unable_to_lock":                   40,
	"empty":                            41,
	"wrong_type":                       42,
	"invalid_kafka_topic":              43,
	"database_error":                   44,
	"timeout":                          45,
}

func (x T) Enum() *T {
	p := new(T)
	*p = x
	return p
}
func (x T) String() string {
	return proto.EnumName(T_name, int32(x))
}
func (x *T) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(T_value, data, "T")
	if err != nil {
		return err
	}
	*x = T(value)
	return nil
}
func (T) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("lang.T", T_name, T_value)
}

func init() { proto.RegisterFile("lang/lang.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xdb, 0x72, 0x13, 0x3d,
	0x0c, 0xfe, 0xfb, 0x93, 0x52, 0xa2, 0x34, 0x8d, 0x10, 0x21, 0x74, 0x80, 0x01, 0x2e, 0xb8, 0x80,
	0x72, 0x7a, 0x19, 0xee, 0x3d, 0x8a, 0xad, 0x0d, 0x26, 0x1b, 0x79, 0xc7, 0xd6, 0x06, 0xf2, 0x88,
	0xbc, 0x15, 0xe3, 0x6c, 0x77, 0x92, 0x0e, 0x37, 0x1e, 0xfb, 0xb3, 0xf4, 0x1d, 0x24, 0x58, 0xb4,
	0xac, 0x9b, 0x6f, 0xf5, 0xf8, 0xda, 0xe5, 0x64, 0x89, 0x26, 0xf5, 0x7e, 0xf7, 0x67, 0x02, 0x17,
	0xdf, 0x69, 0x0e, 0xd3, 0x5e, 0x83, 0x34, 0x51, 0x25, 0xe0, 0x7f, 0xf4, 0x1e, 0xde, 0xf5, 0x45,
	0xb2, 0xfb, 0xc1, 0xc5, 0x71, 0x9b, 0x85, 0xc3, 0xc1, 0x45, 0x75, 0x3e, 0xe9, 0x5e, 0x72, 0x61,
	0x8b, 0x49, 0xf1, 0x82, 0x5e, 0xc0, 0xb3, 0x73, 0xc4, 0xf9, 0x36, 0x15, 0x09, 0xf8, 0x3f, 0x11,
	0xdc, 0x44, 0xdd, 0x73, 0x1b, 0x83, 0x8b, 0xba, 0x8f, 0x26, 0xf8, 0x88, 0x9e, 0xc2, 0x7c, 0xc4,
	0x78, 0x23, 0x6a, 0x38, 0xa1, 0xb7, 0xf0, 0xea, 0xa8, 0x12, 0x8b, 0xd3, 0x64, 0xff, 0x08, 0x5c,
	0xd2, 0x4b, 0x58, 0x3d, 0x10, 0xa8, 0x55, 0x4d, 0xea, 0x35, 0xe0, 0xe3, 0x41, 0xc3, 0x24, 0x2b,
	0xb7, 0x4e, 0x72, 0x4e, 0x19, 0xdf, 0x9c, 0x6b, 0x44, 0xed, 0x7a, 0xc3, 0x15, 0x21, 0x5c, 0x8f,
	0x50, 0x93, 0xf2, 0x0e, 0x97, 0x74, 0x0b, 0x4b, 0xf6, 0x5e, 0x4a, 0x71, 0x96, 0xb6, 0xa2, 0x4e,
	0x7e, 0x77, 0x31, 0x4b, 0xc0, 0xe7, 0xb4, 0x02, 0x1a, 0x6b, 0x7d, 0x96, 0x20, 0x6a, 0x91, 0x5b,
	0x9c, 0x57, 0xfc, 0xf4, 0x3e, 0x9a, 0x28, 0x62, 0x78, 0x55, 0xfd, 0xff, 0xca, 0x49, 0x37, 0x8e,
	0xbd, 0x4f, 0xbd, 0x0e, 0x09, 0x4e, 0x8d, 0x4f, 0xe8, 0x35, 0xdc, 0x0e, 0x05, 0x43, 0xcc, 0x07,
	0xbf, 0x40, 0x0b, 0x98, 0xdd, 0x1b, 0x09, 0xa2, 0x07, 0x9c, 0xd6, 0xb8, 0xbd, 0xf2, 0xba, 0x15,
	0x67, 0xc9, 0x15, 0xd1, 0xe0, 0x76, 0x52, 0x0a, 0x6f, 0x04, 0x67, 0x35, 0xae, 0xa5, 0x2e, 0xfa,
	0x3a, 0x2c, 0xd9, 0x75, 0x76, 0xc0, 0xeb, 0xf3, 0x6c, 0x3f, 0x4b, 0x52, 0x5c, 0xd0, 0x12, 0x70,
	0x44, 0x8e, 0x3b, 0x5e, 0xf7, 0x0d, 0x62, 0xed, 0x3d, 0xf1, 0xb6, 0xc9, 0x6f, 0xf1, 0x03, 0x4d,
	0xe1, 0x72, 0xa0, 0xf9, 0x48, 0x37, 0x00, 0x83, 0x4b, 0x3b, 0x74, 0x82, 0x77, 0x75, 0xad, 0x23,
	0xc9, 0x96, 0x9b, 0x2d, 0xbb, 0xa3, 0x30, 0x7e, 0xaa, 0x3c, 0x81, 0x8d, 0xd7, 0x5c, 0xe4, 0x7e,
	0xe4, 0x9f, 0x69, 0x06, 0x57, 0x16, 0x77, 0x92, 0x7a, 0xc3, 0x2f, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x8f, 0xd9, 0x8c, 0x63, 0x02, 0x00, 0x00,
}
