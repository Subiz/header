# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: lang/lang.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='lang/lang.proto',
  package='lang',
  syntax='proto2',
  serialized_pb=_b('\n\x0flang/lang.proto\x12\x04lang*\xc1\x05\n\x01T\x12\r\n\tundefined\x10\x00\x12$\n user_has_already_in_conversation\x10\x01\x12\x17\n\x13\x63onversation_closed\x10\x02\x12\x12\n\x0einvalid_invite\x10\x03\x12\x11\n\rinvalid_agent\x10\x04\x12\x1f\n\x1buser_is_not_in_conversation\x10\x05\x12\x1a\n\x16\x63onversation_not_found\x10\x06\x12\x12\n\x0einternal_error\x10\x1e\x12\x11\n\rinvalid_input\x10\x16\x12\x10\n\x0cinvalid_form\x10\x14\x12\x18\n\x14\x61\x63\x63\x65ss_token_expired\x10\x15\x12\x16\n\x12invalid_credential\x10\r\x12\x16\n\x12\x63redential_not_set\x10\x07\x12\x1f\n\x1bwrong_account_in_credential\x10\x08\x12\x1c\n\x18wrong_user_in_credential\x10\n\x12\x0f\n\x0b\x61\x63\x63\x65ss_deny\x10\t\x12\x1a\n\x16unable_to_send_message\x10\x0b\x12\x12\n\x0etopic_is_empty\x10\x0c\x12\x10\n\x0cinvalid_json\x10\x0f\x12\x14\n\x10invalid_protobuf\x10\x10\x12\x12\n\x0eunable_to_lock\x10(\x12\t\n\x05\x65mpty\x10)\x12\x0e\n\nwrong_type\x10*\x12\x17\n\x13invalid_kafka_topic\x10+\x12\x12\n\x0e\x64\x61tabase_error\x10,\x12\x0b\n\x07timeout\x10-\x12\x13\n\x0fwebsocket_error\x10.\x12\x0f\n\x0bkafka_error\x10/\x12\x11\n\rinvalid_token\x10\x30\x12\x15\n\x11\x61\x63\x63ount_not_found\x10\x31\x12\x13\n\x0f\x61gent_not_found\x10\x32\x12\x11\n\rinvalid_email\x10<')
)

_T = _descriptor.EnumDescriptor(
  name='T',
  full_name='lang.T',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='undefined', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='user_has_already_in_conversation', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_closed', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_invite', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_agent', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='user_is_not_in_conversation', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_not_found', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='internal_error', index=7, number=30,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_input', index=8, number=22,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_form', index=9, number=20,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='access_token_expired', index=10, number=21,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_credential', index=11, number=13,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='credential_not_set', index=12, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='wrong_account_in_credential', index=13, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='wrong_user_in_credential', index=14, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='access_deny', index=15, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='unable_to_send_message', index=16, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='topic_is_empty', index=17, number=12,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_json', index=18, number=15,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_protobuf', index=19, number=16,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='unable_to_lock', index=20, number=40,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='empty', index=21, number=41,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='wrong_type', index=22, number=42,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_kafka_topic', index=23, number=43,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='database_error', index=24, number=44,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='timeout', index=25, number=45,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='websocket_error', index=26, number=46,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='kafka_error', index=27, number=47,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_token', index=28, number=48,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='account_not_found', index=29, number=49,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='agent_not_found', index=30, number=50,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='invalid_email', index=31, number=60,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=26,
  serialized_end=731,
)
_sym_db.RegisterEnumDescriptor(_T)

T = enum_type_wrapper.EnumTypeWrapper(_T)
undefined = 0
user_has_already_in_conversation = 1
conversation_closed = 2
invalid_invite = 3
invalid_agent = 4
user_is_not_in_conversation = 5
conversation_not_found = 6
internal_error = 30
invalid_input = 22
invalid_form = 20
access_token_expired = 21
invalid_credential = 13
credential_not_set = 7
wrong_account_in_credential = 8
wrong_user_in_credential = 10
access_deny = 9
unable_to_send_message = 11
topic_is_empty = 12
invalid_json = 15
invalid_protobuf = 16
unable_to_lock = 40
empty = 41
wrong_type = 42
invalid_kafka_topic = 43
database_error = 44
timeout = 45
websocket_error = 46
kafka_error = 47
invalid_token = 48
account_not_found = 49
agent_not_found = 50
invalid_email = 60


DESCRIPTOR.enum_types_by_name['T'] = _T
_sym_db.RegisterFileDescriptor(DESCRIPTOR)


# @@protoc_insertion_point(module_scope)
