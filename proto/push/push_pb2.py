# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: push/push.proto

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
  name='push/push.proto',
  package='push',
  syntax='proto3',
  serialized_pb=_b('\n\x0fpush/push.proto\x12\x04push\"\x07\n\x05\x45mpty\"\x10\n\x02Id\x12\n\n\x02Id\x18\x01 \x01(\t\"_\n\x05Token\x12\r\n\x05Token\x18\x01 \x01(\t\x12\x0e\n\x06UserId\x18\x02 \x01(\t\x12\x11\n\tExpiredIn\x18\x03 \x01(\t\x12$\n\nDeviceType\x18\x04 \x01(\x0e\x32\x10.push.DeviceType\"3\n\x12TokenRemoveRequest\x12\x0e\n\x06UserId\x18\x01 \x01(\t\x12\r\n\x05Token\x18\x02 \x01(\t\"\xf2\x01\n\x15UserPushConfiguration\x12\x11\n\tAccountId\x18\x01 \x01(\t\x12\x0e\n\x06UserId\x18\x02 \x01(\t\x12\r\n\x05\x45mail\x18\x03 \x01(\t\x12 \n\x18\x45mailNotificationEnabled\x18\x04 \x01(\x08\x12&\n\x1e\x44\x65sktopPushNotificationEnabled\x18\x05 \x01(\x08\x12%\n\x1dMobilePushNotificationEnabled\x18\x06 \x01(\x08\x12\x1e\n\x16MobilePushDelayTimming\x18\x07 \x01(\x05\x12\x16\n\x0e\x45mailThreshold\x18\x08 \x01(\x05*%\n\nDeviceType\x12\x0b\n\x07\x44\x45SKTOP\x10\x00\x12\n\n\x06MOBILE\x10\x01\x32\xdb\x01\n\x08NotiDeli\x12+\n\rRegisterToken\x12\x0b.push.Token\x1a\x0b.push.Empty\"\x00\x12\x36\n\x0bRemoveToken\x12\x18.push.TokenRemoveRequest\x1a\x0b.push.Empty\"\x00\x12\x34\n\x06\x43onfig\x12\x1b.push.UserPushConfiguration\x1a\x0b.push.Empty\"\x00\x12\x34\n\tGetConfig\x12\x08.push.Id\x1a\x1b.push.UserPushConfiguration\"\x00\x62\x06proto3')
)

_DEVICETYPE = _descriptor.EnumDescriptor(
  name='DeviceType',
  full_name='push.DeviceType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DESKTOP', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MOBILE', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=447,
  serialized_end=484,
)
_sym_db.RegisterEnumDescriptor(_DEVICETYPE)

DeviceType = enum_type_wrapper.EnumTypeWrapper(_DEVICETYPE)
DESKTOP = 0
MOBILE = 1



_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='push.Empty',
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
  serialized_start=25,
  serialized_end=32,
)


_ID = _descriptor.Descriptor(
  name='Id',
  full_name='push.Id',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='push.Id.Id', index=0,
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
  serialized_start=34,
  serialized_end=50,
)


_TOKEN = _descriptor.Descriptor(
  name='Token',
  full_name='push.Token',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Token', full_name='push.Token.Token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UserId', full_name='push.Token.UserId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ExpiredIn', full_name='push.Token.ExpiredIn', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeviceType', full_name='push.Token.DeviceType', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=52,
  serialized_end=147,
)


_TOKENREMOVEREQUEST = _descriptor.Descriptor(
  name='TokenRemoveRequest',
  full_name='push.TokenRemoveRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UserId', full_name='push.TokenRemoveRequest.UserId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Token', full_name='push.TokenRemoveRequest.Token', index=1,
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
  serialized_start=149,
  serialized_end=200,
)


_USERPUSHCONFIGURATION = _descriptor.Descriptor(
  name='UserPushConfiguration',
  full_name='push.UserPushConfiguration',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='AccountId', full_name='push.UserPushConfiguration.AccountId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UserId', full_name='push.UserPushConfiguration.UserId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Email', full_name='push.UserPushConfiguration.Email', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='EmailNotificationEnabled', full_name='push.UserPushConfiguration.EmailNotificationEnabled', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DesktopPushNotificationEnabled', full_name='push.UserPushConfiguration.DesktopPushNotificationEnabled', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='MobilePushNotificationEnabled', full_name='push.UserPushConfiguration.MobilePushNotificationEnabled', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='MobilePushDelayTimming', full_name='push.UserPushConfiguration.MobilePushDelayTimming', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='EmailThreshold', full_name='push.UserPushConfiguration.EmailThreshold', index=7,
      number=8, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=203,
  serialized_end=445,
)

_TOKEN.fields_by_name['DeviceType'].enum_type = _DEVICETYPE
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Id'] = _ID
DESCRIPTOR.message_types_by_name['Token'] = _TOKEN
DESCRIPTOR.message_types_by_name['TokenRemoveRequest'] = _TOKENREMOVEREQUEST
DESCRIPTOR.message_types_by_name['UserPushConfiguration'] = _USERPUSHCONFIGURATION
DESCRIPTOR.enum_types_by_name['DeviceType'] = _DEVICETYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'push.push_pb2'
  # @@protoc_insertion_point(class_scope:push.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Id = _reflection.GeneratedProtocolMessageType('Id', (_message.Message,), dict(
  DESCRIPTOR = _ID,
  __module__ = 'push.push_pb2'
  # @@protoc_insertion_point(class_scope:push.Id)
  ))
_sym_db.RegisterMessage(Id)

Token = _reflection.GeneratedProtocolMessageType('Token', (_message.Message,), dict(
  DESCRIPTOR = _TOKEN,
  __module__ = 'push.push_pb2'
  # @@protoc_insertion_point(class_scope:push.Token)
  ))
_sym_db.RegisterMessage(Token)

TokenRemoveRequest = _reflection.GeneratedProtocolMessageType('TokenRemoveRequest', (_message.Message,), dict(
  DESCRIPTOR = _TOKENREMOVEREQUEST,
  __module__ = 'push.push_pb2'
  # @@protoc_insertion_point(class_scope:push.TokenRemoveRequest)
  ))
_sym_db.RegisterMessage(TokenRemoveRequest)

UserPushConfiguration = _reflection.GeneratedProtocolMessageType('UserPushConfiguration', (_message.Message,), dict(
  DESCRIPTOR = _USERPUSHCONFIGURATION,
  __module__ = 'push.push_pb2'
  # @@protoc_insertion_point(class_scope:push.UserPushConfiguration)
  ))
_sym_db.RegisterMessage(UserPushConfiguration)



_NOTIDELI = _descriptor.ServiceDescriptor(
  name='NotiDeli',
  full_name='push.NotiDeli',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=487,
  serialized_end=706,
  methods=[
  _descriptor.MethodDescriptor(
    name='RegisterToken',
    full_name='push.NotiDeli.RegisterToken',
    index=0,
    containing_service=None,
    input_type=_TOKEN,
    output_type=_EMPTY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveToken',
    full_name='push.NotiDeli.RemoveToken',
    index=1,
    containing_service=None,
    input_type=_TOKENREMOVEREQUEST,
    output_type=_EMPTY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Config',
    full_name='push.NotiDeli.Config',
    index=2,
    containing_service=None,
    input_type=_USERPUSHCONFIGURATION,
    output_type=_EMPTY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetConfig',
    full_name='push.NotiDeli.GetConfig',
    index=3,
    containing_service=None,
    input_type=_ID,
    output_type=_USERPUSHCONFIGURATION,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_NOTIDELI)

DESCRIPTOR.services_by_name['NotiDeli'] = _NOTIDELI

# @@protoc_insertion_point(module_scope)