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
	T_undefined                         T = 0
	T_user_has_already_in_conversation  T = 1
	T_conversation_closed               T = 2
	T_invalid_invite                    T = 3
	T_invalid_agent                     T = 4
	T_member_is_not_in_conversation     T = 5
	T_conversation_not_found            T = 6
	T_invalid_left                      T = 14
	T_internal_error                    T = 30
	T_invalid_input                     T = 22
	T_invalid_form                      T = 20
	T_access_token_expired              T = 21
	T_invalid_credential                T = 13
	T_credential_not_set                T = 7
	T_wrong_account_in_credential       T = 8
	T_wrong_password                    T = 18
	T_wrong_user_in_credential          T = 10
	T_invalid_agent_state               T = 19
	T_access_deny                       T = 9
	T_unable_to_send_message            T = 11
	T_topic_is_empty                    T = 12
	T_invalid_json                      T = 15
	T_invalid_protobuf                  T = 16
	T_unable_to_lock                    T = 40
	T_empty                             T = 41
	T_wrong_type                        T = 42
	T_invalid_kafka_topic               T = 43
	T_database_error                    T = 44
	T_timeout                           T = 45
	T_websocket_error                   T = 46
	T_kafka_error                       T = 47
	T_invalid_token                     T = 48
	T_account_not_found                 T = 49
	T_agent_not_found                   T = 50
	T_invalid_email                     T = 60
	T_plan_not_found                    T = 61
	T_agent_group_not_found             T = 62
	T_empty_client_type                 T = 63
	T_empty_url                         T = 64
	T_empty_name                        T = 65
	T_client_not_found                  T = 66
	T_empty_account                     T = 70
	T_invalid_conversation_state        T = 71
	T_invalid_message_id                T = 80
	T_invalid_mask                      T = 81
	T_randomize_error                   T = 82
	T_duplicated_message_received_error T = 83
	T_network_error                     T = 84
	T_rsa_key_not_found                 T = 85
	T_jwt_sign_error                    T = 86
	T_env_config_error                  T = 87
	T_scrypt_error                      T = 90
	T_challenge_not_found               T = 91
	T_wrong_answer                      T = 92
	T_server_listen_error               T = 93
	T_scrypt_file_not_found             T = 94
	T_invalid_topic                     T = 95
	T_file_not_found                    T = 99
	T_user_not_found                    T = 100
	T_empty_md5                         T = 101
	T_base_convert_error                T = 102
	T_s3_error                          T = 103
	T_aws_credential_error              T = 104
	T_aws_send_error                    T = 105
	T_elasticsearch_error               T = 200
	T_invalid_template                  T = 203
	T_sendgrid_send_error               T = 204
	T_whitelist_domain_not_found        T = 205
	T_blacklist_ip_not_found            T = 206
	T_blocked_user_not_found            T = 207
	T_invalid_content_type              T = 210
	T_parse_file_error                  T = 211
	T_invalid_integration_id            T = 220
	T_invalid_integration               T = 221
)

var T_name = map[int32]string{
	0:   "undefined",
	1:   "user_has_already_in_conversation",
	2:   "conversation_closed",
	3:   "invalid_invite",
	4:   "invalid_agent",
	5:   "member_is_not_in_conversation",
	6:   "conversation_not_found",
	14:  "invalid_left",
	30:  "internal_error",
	22:  "invalid_input",
	20:  "invalid_form",
	21:  "access_token_expired",
	13:  "invalid_credential",
	7:   "credential_not_set",
	8:   "wrong_account_in_credential",
	18:  "wrong_password",
	10:  "wrong_user_in_credential",
	19:  "invalid_agent_state",
	9:   "access_deny",
	11:  "unable_to_send_message",
	12:  "topic_is_empty",
	15:  "invalid_json",
	16:  "invalid_protobuf",
	40:  "unable_to_lock",
	41:  "empty",
	42:  "wrong_type",
	43:  "invalid_kafka_topic",
	44:  "database_error",
	45:  "timeout",
	46:  "websocket_error",
	47:  "kafka_error",
	48:  "invalid_token",
	49:  "account_not_found",
	50:  "agent_not_found",
	60:  "invalid_email",
	61:  "plan_not_found",
	62:  "agent_group_not_found",
	63:  "empty_client_type",
	64:  "empty_url",
	65:  "empty_name",
	66:  "client_not_found",
	70:  "empty_account",
	71:  "invalid_conversation_state",
	80:  "invalid_message_id",
	81:  "invalid_mask",
	82:  "randomize_error",
	83:  "duplicated_message_received_error",
	84:  "network_error",
	85:  "rsa_key_not_found",
	86:  "jwt_sign_error",
	87:  "env_config_error",
	90:  "scrypt_error",
	91:  "challenge_not_found",
	92:  "wrong_answer",
	93:  "server_listen_error",
	94:  "scrypt_file_not_found",
	95:  "invalid_topic",
	99:  "file_not_found",
	100: "user_not_found",
	101: "empty_md5",
	102: "base_convert_error",
	103: "s3_error",
	104: "aws_credential_error",
	105: "aws_send_error",
	200: "elasticsearch_error",
	203: "invalid_template",
	204: "sendgrid_send_error",
	205: "whitelist_domain_not_found",
	206: "blacklist_ip_not_found",
	207: "blocked_user_not_found",
	210: "invalid_content_type",
	211: "parse_file_error",
	220: "invalid_integration_id",
	221: "invalid_integration",
}
var T_value = map[string]int32{
	"undefined":                         0,
	"user_has_already_in_conversation":  1,
	"conversation_closed":               2,
	"invalid_invite":                    3,
	"invalid_agent":                     4,
	"member_is_not_in_conversation":     5,
	"conversation_not_found":            6,
	"invalid_left":                      14,
	"internal_error":                    30,
	"invalid_input":                     22,
	"invalid_form":                      20,
	"access_token_expired":              21,
	"invalid_credential":                13,
	"credential_not_set":                7,
	"wrong_account_in_credential":       8,
	"wrong_password":                    18,
	"wrong_user_in_credential":          10,
	"invalid_agent_state":               19,
	"access_deny":                       9,
	"unable_to_send_message":            11,
	"topic_is_empty":                    12,
	"invalid_json":                      15,
	"invalid_protobuf":                  16,
	"unable_to_lock":                    40,
	"empty":                             41,
	"wrong_type":                        42,
	"invalid_kafka_topic":               43,
	"database_error":                    44,
	"timeout":                           45,
	"websocket_error":                   46,
	"kafka_error":                       47,
	"invalid_token":                     48,
	"account_not_found":                 49,
	"agent_not_found":                   50,
	"invalid_email":                     60,
	"plan_not_found":                    61,
	"agent_group_not_found":             62,
	"empty_client_type":                 63,
	"empty_url":                         64,
	"empty_name":                        65,
	"client_not_found":                  66,
	"empty_account":                     70,
	"invalid_conversation_state":        71,
	"invalid_message_id":                80,
	"invalid_mask":                      81,
	"randomize_error":                   82,
	"duplicated_message_received_error": 83,
	"network_error":                     84,
	"rsa_key_not_found":                 85,
	"jwt_sign_error":                    86,
	"env_config_error":                  87,
	"scrypt_error":                      90,
	"challenge_not_found":               91,
	"wrong_answer":                      92,
	"server_listen_error":               93,
	"scrypt_file_not_found":             94,
	"invalid_topic":                     95,
	"file_not_found":                    99,
	"user_not_found":                    100,
	"empty_md5":                         101,
	"base_convert_error":                102,
	"s3_error":                          103,
	"aws_credential_error":              104,
	"aws_send_error":                    105,
	"elasticsearch_error":               200,
	"invalid_template":                  203,
	"sendgrid_send_error":               204,
	"whitelist_domain_not_found":        205,
	"blacklist_ip_not_found":            206,
	"blocked_user_not_found":            207,
	"invalid_content_type":              210,
	"parse_file_error":                  211,
	"invalid_integration_id":            220,
	"invalid_integration":               221,
}

func (x T) String() string {
	return proto.EnumName(T_name, int32(x))
}
func (T) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type L int32

const (
	L_en L = 0
	L_vn L = 1
	L_cn L = 3
	L_fr L = 5
)

var L_name = map[int32]string{
	0: "en",
	1: "vn",
	3: "cn",
	5: "fr",
}
var L_value = map[string]int32{
	"en": 0,
	"vn": 1,
	"cn": 3,
	"fr": 5,
}

func (x L) String() string {
	return proto.EnumName(L_name, int32(x))
}
func (L) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterEnum("lang.T", T_name, T_value)
	proto.RegisterEnum("lang.L", L_name, L_value)
}

func init() { proto.RegisterFile("lang/lang.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 886 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0xdb, 0x72, 0x1c, 0x35,
	0x10, 0xcd, 0x3a, 0xb1, 0x13, 0xcb, 0xb7, 0x8e, 0x7c, 0xc1, 0x71, 0x20, 0x21, 0x05, 0x54, 0x81,
	0x81, 0x04, 0x48, 0xf1, 0xc6, 0xfd, 0x01, 0x5e, 0x78, 0xe0, 0x12, 0xa0, 0x8a, 0x9b, 0x4a, 0x2b,
	0xf5, 0xcc, 0x2a, 0xab, 0x91, 0xa6, 0x24, 0xcd, 0x2e, 0xcb, 0x17, 0xf0, 0x69, 0x54, 0x71, 0xab,
	0x82, 0x57, 0xf8, 0x17, 0xaa, 0x47, 0xb3, 0x19, 0xd9, 0xf0, 0xe2, 0xf1, 0xb4, 0x5a, 0xa7, 0x4f,
	0x9f, 0x3e, 0x3d, 0xcb, 0x0e, 0xac, 0x74, 0xf5, 0x03, 0xfa, 0x73, 0xbf, 0x0d, 0x3e, 0x79, 0x7e,
	0x8d, 0xfe, 0x3f, 0xff, 0x69, 0x8f, 0x4d, 0x1e, 0xf1, 0x3d, 0xb6, 0xdd, 0x39, 0x8d, 0x95, 0x71,
	0xa8, 0xe1, 0x0a, 0x7f, 0x9e, 0x3d, 0xdb, 0x45, 0x0c, 0x62, 0x26, 0xa3, 0x90, 0x36, 0xa0, 0xd4,
	0x2b, 0x61, 0x9c, 0x50, 0xde, 0x2d, 0x30, 0x44, 0x99, 0x8c, 0x77, 0x30, 0xe1, 0x4f, 0xb1, 0xc3,
	0x32, 0x22, 0x94, 0xf5, 0x11, 0x35, 0x6c, 0x70, 0xce, 0xf6, 0x8d, 0x5b, 0x48, 0x6b, 0xb4, 0x30,
	0x6e, 0x61, 0x12, 0xc2, 0x55, 0x7e, 0x93, 0xed, 0xad, 0x63, 0xb2, 0x46, 0x97, 0xe0, 0x1a, 0xbf,
	0xc7, 0x9e, 0x69, 0xb0, 0x99, 0x62, 0x10, 0x26, 0x0a, 0xe7, 0xd3, 0x7f, 0x4a, 0x6c, 0xf2, 0x33,
	0x76, 0x72, 0xa1, 0x04, 0x65, 0x55, 0xbe, 0x73, 0x1a, 0xb6, 0x38, 0xb0, 0xdd, 0x35, 0xa2, 0xc5,
	0x2a, 0xc1, 0x7e, 0xae, 0x9b, 0x30, 0x38, 0x69, 0x05, 0x86, 0xe0, 0x03, 0xdc, 0x29, 0xeb, 0x1a,
	0xd7, 0x76, 0x09, 0x4e, 0xca, 0x8b, 0x95, 0x0f, 0x0d, 0x1c, 0xf1, 0x53, 0x76, 0x24, 0x95, 0xc2,
	0x18, 0x45, 0xf2, 0x73, 0x74, 0x02, 0x7f, 0x68, 0x4d, 0x40, 0x0d, 0xc7, 0xfc, 0x84, 0xf1, 0x75,
	0xae, 0x0a, 0xa8, 0xd1, 0x25, 0x23, 0x2d, 0xec, 0x51, 0x7c, 0x7c, 0xef, 0x69, 0x45, 0x4c, 0x70,
	0x9d, 0xdf, 0x65, 0xb7, 0x97, 0xc1, 0xbb, 0x5a, 0x48, 0xa5, 0x7c, 0xe7, 0x72, 0x4f, 0xe3, 0xc5,
	0x1b, 0xc4, 0x31, 0x27, 0xb4, 0x32, 0xc6, 0xa5, 0x0f, 0x1a, 0x38, 0x7f, 0x9a, 0x9d, 0xe6, 0x58,
	0x2f, 0xfa, 0xc5, 0x1b, 0x8c, 0x64, 0xbe, 0xa0, 0x9c, 0x88, 0x49, 0x26, 0x84, 0x43, 0x7e, 0xc0,
	0x76, 0x06, 0xd6, 0x1a, 0xdd, 0x0a, 0xb6, 0x49, 0xad, 0xce, 0xc9, 0xa9, 0x45, 0x91, 0xbc, 0x88,
	0xe8, 0xb4, 0x68, 0x30, 0x46, 0x59, 0x23, 0xec, 0x50, 0xdd, 0xe4, 0x5b, 0xa3, 0x48, 0x6b, 0x6c,
	0xda, 0xb4, 0x82, 0xdd, 0x52, 0x88, 0xc7, 0xd1, 0x3b, 0x38, 0xe0, 0x47, 0x0c, 0xd6, 0x91, 0xde,
	0x24, 0xd3, 0xae, 0x02, 0xa0, 0xbb, 0x23, 0xae, 0xf5, 0x6a, 0x0e, 0x2f, 0xf2, 0x6d, 0xb6, 0x99,
	0x61, 0x5e, 0xe2, 0xfb, 0x8c, 0x65, 0xfa, 0x69, 0xd5, 0x22, 0x9c, 0x97, 0x84, 0xe7, 0xb2, 0x9a,
	0x4b, 0xd1, 0x17, 0x86, 0x97, 0x09, 0x47, 0xcb, 0x24, 0xa7, 0x32, 0xe2, 0x30, 0x9f, 0x57, 0xf8,
	0x0e, 0xbb, 0x9e, 0x4c, 0x83, 0xbe, 0x4b, 0xf0, 0x2a, 0x3f, 0x64, 0x07, 0x4b, 0x9c, 0x46, 0xaf,
	0xe6, 0x98, 0x86, 0x8c, 0xfb, 0xd4, 0x66, 0x86, 0xc9, 0x81, 0x07, 0xe5, 0x48, 0xfb, 0x71, 0xc1,
	0x6b, 0xfc, 0x98, 0xdd, 0x5c, 0x0b, 0x3e, 0x5a, 0xe4, 0x75, 0xc2, 0xcb, 0x92, 0x8d, 0xc1, 0x37,
	0xca, 0xeb, 0xd8, 0x48, 0x63, 0xe1, 0x2d, 0x22, 0xd6, 0x5a, 0x59, 0xda, 0xeb, 0x6d, 0x7e, 0x8b,
	0x1d, 0xe7, 0xbb, 0x75, 0xf0, 0x5d, 0x5b, 0x1c, 0xbd, 0x43, 0xd5, 0xfa, 0xde, 0x85, 0xb2, 0x86,
	0x32, 0xfa, 0xbe, 0xdf, 0xa5, 0x25, 0xca, 0xe1, 0x2e, 0x58, 0x78, 0x8f, 0x64, 0xc9, 0xaf, 0x4e,
	0x36, 0x08, 0xef, 0x93, 0xb6, 0x43, 0xfe, 0x88, 0xf5, 0x01, 0xb1, 0xc9, 0x59, 0x03, 0x7f, 0xf8,
	0x90, 0xdf, 0x61, 0x67, 0x4f, 0x3c, 0x57, 0x9a, 0x3f, 0xcf, 0xfd, 0xa3, 0xd2, 0x93, 0xc3, 0x7c,
	0x85, 0xd1, 0xf0, 0x49, 0x39, 0xce, 0x46, 0xc6, 0x39, 0x7c, 0x4a, 0xfd, 0x07, 0xe9, 0xb4, 0x6f,
	0xcc, 0x8f, 0x6b, 0xc5, 0x3f, 0xe3, 0x2f, 0xb0, 0x7b, 0xba, 0x6b, 0xad, 0x51, 0x32, 0xe1, 0x88,
	0x10, 0x50, 0xa1, 0x59, 0xa0, 0x1e, 0xd2, 0x3e, 0x27, 0x62, 0x0e, 0xd3, 0xd2, 0x87, 0xf9, 0x10,
	0x7a, 0x44, 0x7d, 0x87, 0x28, 0xc5, 0x1c, 0x57, 0x45, 0x0b, 0x5f, 0x90, 0x7a, 0x8f, 0x97, 0x49,
	0x44, 0x53, 0xbb, 0x21, 0xf5, 0x4b, 0x6a, 0x16, 0xdd, 0x82, 0xf8, 0x57, 0xa6, 0x1e, 0xa2, 0x5f,
	0x11, 0xc3, 0xa8, 0xc2, 0xaa, 0x5d, 0x0f, 0xf7, 0xeb, 0xfe, 0x1b, 0x32, 0x93, 0xd6, 0xa2, 0xab,
	0xb1, 0x00, 0xfd, 0x86, 0x52, 0x87, 0x45, 0x72, 0x71, 0x89, 0x01, 0xbe, 0xa5, 0xd4, 0x88, 0x61,
	0x81, 0x41, 0x58, 0x13, 0x13, 0xae, 0x6b, 0x7d, 0x47, 0x93, 0x1a, 0x50, 0x2b, 0x63, 0x4b, 0x94,
	0xef, 0x2f, 0x5a, 0x85, 0x4c, 0x28, 0x88, 0xed, 0xa5, 0x34, 0xd5, 0x1b, 0x9c, 0x56, 0x6f, 0x8c,
	0xe9, 0x71, 0x9a, 0x8d, 0x7e, 0x13, 0x90, 0x44, 0xef, 0x7d, 0x9b, 0x27, 0xb2, 0x6e, 0xa0, 0xe2,
	0xbb, 0xec, 0x46, 0x7c, 0x38, 0xbc, 0xd5, 0xfd, 0x87, 0x64, 0x19, 0x8b, 0xfd, 0x1d, 0x4e, 0x66,
	0x54, 0x82, 0x4e, 0xfa, 0xad, 0xcc, 0x31, 0xc3, 0x4f, 0xd9, 0x21, 0x5a, 0x19, 0x93, 0x51, 0x11,
	0x65, 0x50, 0xb3, 0xe1, 0xe0, 0xe7, 0x09, 0x3f, 0x1e, 0xf7, 0x30, 0x61, 0xd3, 0x5a, 0x1a, 0xfc,
	0x2f, 0x13, 0xba, 0x40, 0x00, 0x75, 0x30, 0xba, 0x44, 0xfa, 0x75, 0xc2, 0xef, 0xb2, 0xb3, 0xe5,
	0xcc, 0x24, 0x24, 0x69, 0x84, 0xf6, 0x8d, 0x34, 0xa5, 0x9b, 0x7f, 0x9b, 0xf0, 0xdb, 0xec, 0x64,
	0x6a, 0xa5, 0x9a, 0xf7, 0x09, 0xa6, 0xf4, 0xf3, 0xef, 0xc3, 0x21, 0x6d, 0x9d, 0x16, 0x97, 0x74,
	0xf8, 0x63, 0xc2, 0x6f, 0xb1, 0xa3, 0xc2, 0x8e, 0xe9, 0x89, 0xe1, 0xff, 0xec, 0x69, 0xb6, 0x32,
	0x44, 0xcc, 0xc2, 0x67, 0x32, 0x7f, 0xf5, 0x70, 0xe3, 0x37, 0x37, 0x61, 0x1d, 0xb2, 0x7f, 0x8d,
	0x86, 0xbf, 0xfb, 0x1e, 0xfe, 0xe7, 0x10, 0xfe, 0x99, 0x9c, 0x3f, 0xc7, 0x26, 0x1f, 0xf3, 0x2d,
	0xb6, 0x81, 0x0e, 0xae, 0xd0, 0x73, 0x41, 0x3f, 0x32, 0x5b, 0x6c, 0x43, 0x39, 0xb8, 0x4a, 0xcf,
	0x2a, 0xc0, 0xe6, 0x74, 0xab, 0xff, 0x2e, 0x3d, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x25,
	0xf0, 0xd0, 0xcf, 0x06, 0x00, 0x00,
}
