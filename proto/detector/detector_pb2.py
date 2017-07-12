# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: detector/detector.proto

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
  name='detector/detector.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x17\x64\x65tector/detector.proto\"\xcf\x01\n\tUserAgent\x12\n\n\x02Id\x18\x01 \x01(\x05\x12\x1a\n\x12\x42rowserVersionFull\x18\x02 \x01(\t\x12\x14\n\x0cLayoutEngine\x18\x03 \x01(\t\x12\x1b\n\x13LayoutEngineVersion\x18\x04 \x01(\t\x12\x1c\n\x14HardwareArchitecture\x18\x05 \x01(\t\x12\n\n\x02Os\x18\x06 \x01(\t\x12\x16\n\x0e\x42rowserVersion\x18\x07 \x01(\t\x12\x12\n\nDeviceType\x18\x08 \x01(\t\x12\x11\n\tUserAgent\x18\t \x01(\t\"\x18\n\x06String\x12\x0e\n\x06String\x18\x01 \x01(\t\"\x14\n\x06Zipped\x12\n\n\x02Id\x18\x01 \x01(\x05\"Y\n\x08Location\x12\n\n\x02Id\x18\x01 \x01(\x05\x12\x0f\n\x07\x43ountry\x18\x02 \x01(\t\x12\x0c\n\x04\x43ity\x18\x03 \x01(\t\x12\x10\n\x08TimeZone\x18\x04 \x01(\t\x12\x10\n\x08Language\x18\x05 \x01(\t2\xf1\x01\n\x08\x44\x65tector\x12(\n\x0f\x44\x65tectUserAgent\x12\x07.String\x1a\n.UserAgent\"\x00\x12\"\n\x0cZipUserAgent\x12\x07.String\x1a\x07.Zipped\"\x00\x12\'\n\x0eUnzipUserAgent\x12\x07.Zipped\x1a\n.UserAgent\"\x00\x12&\n\x0e\x44\x65tectLocation\x12\x07.String\x1a\t.Location\"\x00\x12!\n\x0bZipLanguage\x12\x07.String\x1a\x07.Zipped\"\x00\x12#\n\rUnzipLanguage\x12\x07.Zipped\x1a\x07.String\"\x00\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_USERAGENT = _descriptor.Descriptor(
  name='UserAgent',
  full_name='UserAgent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='UserAgent.Id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='BrowserVersionFull', full_name='UserAgent.BrowserVersionFull', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='LayoutEngine', full_name='UserAgent.LayoutEngine', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='LayoutEngineVersion', full_name='UserAgent.LayoutEngineVersion', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='HardwareArchitecture', full_name='UserAgent.HardwareArchitecture', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Os', full_name='UserAgent.Os', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='BrowserVersion', full_name='UserAgent.BrowserVersion', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeviceType', full_name='UserAgent.DeviceType', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UserAgent', full_name='UserAgent.UserAgent', index=8,
      number=9, type=9, cpp_type=9, label=1,
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
  serialized_start=28,
  serialized_end=235,
)


_STRING = _descriptor.Descriptor(
  name='String',
  full_name='String',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='String', full_name='String.String', index=0,
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
  serialized_start=237,
  serialized_end=261,
)


_ZIPPED = _descriptor.Descriptor(
  name='Zipped',
  full_name='Zipped',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='Zipped.Id', index=0,
      number=1, type=5, cpp_type=1, label=1,
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
  serialized_start=263,
  serialized_end=283,
)


_LOCATION = _descriptor.Descriptor(
  name='Location',
  full_name='Location',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='Location.Id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Country', full_name='Location.Country', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='City', full_name='Location.City', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='TimeZone', full_name='Location.TimeZone', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Language', full_name='Location.Language', index=4,
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
  serialized_start=285,
  serialized_end=374,
)

DESCRIPTOR.message_types_by_name['UserAgent'] = _USERAGENT
DESCRIPTOR.message_types_by_name['String'] = _STRING
DESCRIPTOR.message_types_by_name['Zipped'] = _ZIPPED
DESCRIPTOR.message_types_by_name['Location'] = _LOCATION

UserAgent = _reflection.GeneratedProtocolMessageType('UserAgent', (_message.Message,), dict(
  DESCRIPTOR = _USERAGENT,
  __module__ = 'detector.detector_pb2'
  # @@protoc_insertion_point(class_scope:UserAgent)
  ))
_sym_db.RegisterMessage(UserAgent)

String = _reflection.GeneratedProtocolMessageType('String', (_message.Message,), dict(
  DESCRIPTOR = _STRING,
  __module__ = 'detector.detector_pb2'
  # @@protoc_insertion_point(class_scope:String)
  ))
_sym_db.RegisterMessage(String)

Zipped = _reflection.GeneratedProtocolMessageType('Zipped', (_message.Message,), dict(
  DESCRIPTOR = _ZIPPED,
  __module__ = 'detector.detector_pb2'
  # @@protoc_insertion_point(class_scope:Zipped)
  ))
_sym_db.RegisterMessage(Zipped)

Location = _reflection.GeneratedProtocolMessageType('Location', (_message.Message,), dict(
  DESCRIPTOR = _LOCATION,
  __module__ = 'detector.detector_pb2'
  # @@protoc_insertion_point(class_scope:Location)
  ))
_sym_db.RegisterMessage(Location)


# @@protoc_insertion_point(module_scope)
