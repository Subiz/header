# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gateway/gateway.proto

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
  name='gateway/gateway.proto',
  package='gateway',
  syntax='proto3',
  serialized_pb=_b('\n\x15gateway/gateway.proto\x12\x07gateway\"\x07\n\x05\x45mpty\"C\n\x0cUsersPackage\x12\x0f\n\x07UserIds\x18\x02 \x03(\t\x12\x0f\n\x07Payload\x18\x03 \x01(\t\x12\x11\n\tAccountId\x18\x04 \x01(\t\"G\n\x0e\x43hannelPackage\x12\x11\n\tChannelId\x18\x02 \x01(\t\x12\x0f\n\x07Payload\x18\x03 \x01(\t\x12\x11\n\tAccountId\x18\x04 \x01(\t\"S\n\x0bSubsPackage\x12\x0f\n\x07UserIds\x18\x02 \x03(\t\x12\x0f\n\x07\x43onnIds\x18\x03 \x03(\t\x12\x0f\n\x07Payload\x18\x04 \x01(\t\x12\x11\n\tAccountId\x18\x05 \x01(\t*:\n\nClientType\x12\x0c\n\x08\x46orAgent\x10\x00\x12\x0e\n\nForVisitor\x10\x01\x12\x0e\n\nForAccount\x10\x02\x32\xb3\x01\n\x07Gateway\x12\x36\n\x0bSendToUsers\x12\x15.gateway.UsersPackage\x1a\x0e.gateway.Empty\"\x00\x12:\n\rSendToChannel\x12\x17.gateway.ChannelPackage\x1a\x0e.gateway.Empty\"\x00\x12\x34\n\nSendToSubs\x12\x14.gateway.SubsPackage\x1a\x0e.gateway.Empty\"\x00\x62\x06proto3')
)

_CLIENTTYPE = _descriptor.EnumDescriptor(
  name='ClientType',
  full_name='gateway.ClientType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ForAgent', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ForVisitor', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ForAccount', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=270,
  serialized_end=328,
)
_sym_db.RegisterEnumDescriptor(_CLIENTTYPE)

ClientType = enum_type_wrapper.EnumTypeWrapper(_CLIENTTYPE)
ForAgent = 0
ForVisitor = 1
ForAccount = 2



_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='gateway.Empty',
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
  serialized_start=34,
  serialized_end=41,
)


_USERSPACKAGE = _descriptor.Descriptor(
  name='UsersPackage',
  full_name='gateway.UsersPackage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UserIds', full_name='gateway.UsersPackage.UserIds', index=0,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Payload', full_name='gateway.UsersPackage.Payload', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='AccountId', full_name='gateway.UsersPackage.AccountId', index=2,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=43,
  serialized_end=110,
)


_CHANNELPACKAGE = _descriptor.Descriptor(
  name='ChannelPackage',
  full_name='gateway.ChannelPackage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ChannelId', full_name='gateway.ChannelPackage.ChannelId', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Payload', full_name='gateway.ChannelPackage.Payload', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='AccountId', full_name='gateway.ChannelPackage.AccountId', index=2,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=112,
  serialized_end=183,
)


_SUBSPACKAGE = _descriptor.Descriptor(
  name='SubsPackage',
  full_name='gateway.SubsPackage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UserIds', full_name='gateway.SubsPackage.UserIds', index=0,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ConnIds', full_name='gateway.SubsPackage.ConnIds', index=1,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Payload', full_name='gateway.SubsPackage.Payload', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='AccountId', full_name='gateway.SubsPackage.AccountId', index=3,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=185,
  serialized_end=268,
)

DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['UsersPackage'] = _USERSPACKAGE
DESCRIPTOR.message_types_by_name['ChannelPackage'] = _CHANNELPACKAGE
DESCRIPTOR.message_types_by_name['SubsPackage'] = _SUBSPACKAGE
DESCRIPTOR.enum_types_by_name['ClientType'] = _CLIENTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'gateway.gateway_pb2'
  # @@protoc_insertion_point(class_scope:gateway.Empty)
  ))
_sym_db.RegisterMessage(Empty)

UsersPackage = _reflection.GeneratedProtocolMessageType('UsersPackage', (_message.Message,), dict(
  DESCRIPTOR = _USERSPACKAGE,
  __module__ = 'gateway.gateway_pb2'
  # @@protoc_insertion_point(class_scope:gateway.UsersPackage)
  ))
_sym_db.RegisterMessage(UsersPackage)

ChannelPackage = _reflection.GeneratedProtocolMessageType('ChannelPackage', (_message.Message,), dict(
  DESCRIPTOR = _CHANNELPACKAGE,
  __module__ = 'gateway.gateway_pb2'
  # @@protoc_insertion_point(class_scope:gateway.ChannelPackage)
  ))
_sym_db.RegisterMessage(ChannelPackage)

SubsPackage = _reflection.GeneratedProtocolMessageType('SubsPackage', (_message.Message,), dict(
  DESCRIPTOR = _SUBSPACKAGE,
  __module__ = 'gateway.gateway_pb2'
  # @@protoc_insertion_point(class_scope:gateway.SubsPackage)
  ))
_sym_db.RegisterMessage(SubsPackage)



_GATEWAY = _descriptor.ServiceDescriptor(
  name='Gateway',
  full_name='gateway.Gateway',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=331,
  serialized_end=510,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendToUsers',
    full_name='gateway.Gateway.SendToUsers',
    index=0,
    containing_service=None,
    input_type=_USERSPACKAGE,
    output_type=_EMPTY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SendToChannel',
    full_name='gateway.Gateway.SendToChannel',
    index=1,
    containing_service=None,
    input_type=_CHANNELPACKAGE,
    output_type=_EMPTY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SendToSubs',
    full_name='gateway.Gateway.SendToSubs',
    index=2,
    containing_service=None,
    input_type=_SUBSPACKAGE,
    output_type=_EMPTY,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_GATEWAY)

DESCRIPTOR.services_by_name['Gateway'] = _GATEWAY

# @@protoc_insertion_point(module_scope)