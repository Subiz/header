# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: presence/presence.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='presence/presence.proto',
  package='presence',
  syntax='proto3',
  serialized_pb=_b('\n\x17presence/presence.proto\x12\x08presence\"\'\n\x06UserId\x12\x11\n\tAccountId\x18\x01 \x01(\t\x12\n\n\x02Id\x18\x02 \x01(\t\",\n\x07UserIds\x12!\n\x07UserIds\x18\x01 \x03(\x0b\x32\x10.presence.UserId\"\x07\n\x05\x45mpty\"\x14\n\x04Time\x12\x0c\n\x04Time\x18\x01 \x01(\t\"!\n\x05Times\x12\n\n\x02Id\x18\x01 \x01(\t\x12\x0c\n\x04Time\x18\x02 \x01(\t2\xd8\x01\n\x0bPresenceMgr\x12+\n\x04Ping\x12\x10.presence.UserId\x1a\x0f.presence.Empty\"\x00\x12*\n\x03\x42ye\x12\x10.presence.UserId\x1a\x0f.presence.Empty\"\x00\x12\x35\n\x0fGetLastPingTime\x12\x10.presence.UserId\x1a\x0e.presence.Time\"\x00\x12\x39\n\x11ListLastPingTimes\x12\x11.presence.UserIds\x1a\x0f.presence.Times\"\x00\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_USERID = _descriptor.Descriptor(
  name='UserId',
  full_name='presence.UserId',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='AccountId', full_name='presence.UserId.AccountId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Id', full_name='presence.UserId.Id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=37,
  serialized_end=76,
)


_USERIDS = _descriptor.Descriptor(
  name='UserIds',
  full_name='presence.UserIds',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UserIds', full_name='presence.UserIds.UserIds', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=78,
  serialized_end=122,
)


_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='presence.Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=124,
  serialized_end=131,
)


_TIME = _descriptor.Descriptor(
  name='Time',
  full_name='presence.Time',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Time', full_name='presence.Time.Time', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=133,
  serialized_end=153,
)


_TIMES = _descriptor.Descriptor(
  name='Times',
  full_name='presence.Times',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='presence.Times.Id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Time', full_name='presence.Times.Time', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=155,
  serialized_end=188,
)

_USERIDS.fields_by_name['UserIds'].message_type = _USERID
DESCRIPTOR.message_types_by_name['UserId'] = _USERID
DESCRIPTOR.message_types_by_name['UserIds'] = _USERIDS
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Time'] = _TIME
DESCRIPTOR.message_types_by_name['Times'] = _TIMES

UserId = _reflection.GeneratedProtocolMessageType('UserId', (_message.Message,), dict(
  DESCRIPTOR = _USERID,
  __module__ = 'presence.presence_pb2'
  # @@protoc_insertion_point(class_scope:presence.UserId)
  ))
_sym_db.RegisterMessage(UserId)

UserIds = _reflection.GeneratedProtocolMessageType('UserIds', (_message.Message,), dict(
  DESCRIPTOR = _USERIDS,
  __module__ = 'presence.presence_pb2'
  # @@protoc_insertion_point(class_scope:presence.UserIds)
  ))
_sym_db.RegisterMessage(UserIds)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'presence.presence_pb2'
  # @@protoc_insertion_point(class_scope:presence.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Time = _reflection.GeneratedProtocolMessageType('Time', (_message.Message,), dict(
  DESCRIPTOR = _TIME,
  __module__ = 'presence.presence_pb2'
  # @@protoc_insertion_point(class_scope:presence.Time)
  ))
_sym_db.RegisterMessage(Time)

Times = _reflection.GeneratedProtocolMessageType('Times', (_message.Message,), dict(
  DESCRIPTOR = _TIMES,
  __module__ = 'presence.presence_pb2'
  # @@protoc_insertion_point(class_scope:presence.Times)
  ))
_sym_db.RegisterMessage(Times)


# @@protoc_insertion_point(module_scope)
