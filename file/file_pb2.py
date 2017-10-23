# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: file/file.proto

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


from bitbucket.org.subiz.header.common import common_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='file/file.proto',
  package='file',
  syntax='proto2',
  serialized_pb=_b('\n\x0f\x66ile/file.proto\x12\x04\x66ile\x1a.bitbucket.org/subiz/header/common/common.proto\"\x8a\x01\n\nFileHeader\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\x12\n\naccount_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x0c\n\x04size\x18\x05 \x01(\x03\x12\x0b\n\x03md5\x18\x06 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x07 \x01(\t\".\n\x02Id\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\n\n\x02id\x18\x03 \x01(\t\"n\n\rPresignResult\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\x12\n\naccount_id\x18\x03 \x01(\t\x12\x0b\n\x03url\x18\x04 \x01(\t\x12\n\n\x02id\x18\x06 \x01(\t\x12\x12\n\nsigned_url\x18\x05 \x01(\t\"\xbf\x01\n\x04\x46ile\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\x12\n\naccount_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x0c\n\x04size\x18\x05 \x01(\x03\x12\x0b\n\x03md5\x18\x06 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\n \x01(\t\x12\x0f\n\x07\x63reated\x18\x07 \x01(\t\x12\x0b\n\x03url\x18\x08 \x01(\t\x12\x0f\n\x07\x63reator\x18\t \x01(\t\x12\n\n\x02id\x18\x0b \x01(\t*F\n\x05\x45vent\x12\x15\n\x11\x46ileSignRequested\x10\x02\x12\x0f\n\x0b\x46ileCreated\x10\x03\x12\x15\n\x11\x46ileInfoRequested\x10\x04')
  ,
  dependencies=[bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2.DESCRIPTOR,])

_EVENT = _descriptor.EnumDescriptor(
  name='Event',
  full_name='file.Event',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FileSignRequested', index=0, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FileCreated', index=1, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FileInfoRequested', index=2, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=568,
  serialized_end=638,
)
_sym_db.RegisterEnumDescriptor(_EVENT)

Event = enum_type_wrapper.EnumTypeWrapper(_EVENT)
FileSignRequested = 2
FileCreated = 3
FileInfoRequested = 4



_FILEHEADER = _descriptor.Descriptor(
  name='FileHeader',
  full_name='file.FileHeader',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='file.FileHeader.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='file.FileHeader.account_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='file.FileHeader.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='file.FileHeader.type', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='size', full_name='file.FileHeader.size', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='md5', full_name='file.FileHeader.md5', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='file.FileHeader.description', index=6,
      number=7, type=9, cpp_type=9, label=1,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=74,
  serialized_end=212,
)


_ID = _descriptor.Descriptor(
  name='Id',
  full_name='file.Id',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='file.Id.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='file.Id.id', index=1,
      number=3, type=9, cpp_type=9, label=1,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=214,
  serialized_end=260,
)


_PRESIGNRESULT = _descriptor.Descriptor(
  name='PresignResult',
  full_name='file.PresignResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='file.PresignResult.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='file.PresignResult.account_id', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='url', full_name='file.PresignResult.url', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='file.PresignResult.id', index=3,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='signed_url', full_name='file.PresignResult.signed_url', index=4,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=262,
  serialized_end=372,
)


_FILE = _descriptor.Descriptor(
  name='File',
  full_name='file.File',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='file.File.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='file.File.account_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='file.File.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='file.File.type', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='size', full_name='file.File.size', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='md5', full_name='file.File.md5', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='file.File.description', index=6,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created', full_name='file.File.created', index=7,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='url', full_name='file.File.url', index=8,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='creator', full_name='file.File.creator', index=9,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='file.File.id', index=10,
      number=11, type=9, cpp_type=9, label=1,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=375,
  serialized_end=566,
)

_FILEHEADER.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_ID.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_PRESIGNRESULT.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_FILE.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
DESCRIPTOR.message_types_by_name['FileHeader'] = _FILEHEADER
DESCRIPTOR.message_types_by_name['Id'] = _ID
DESCRIPTOR.message_types_by_name['PresignResult'] = _PRESIGNRESULT
DESCRIPTOR.message_types_by_name['File'] = _FILE
DESCRIPTOR.enum_types_by_name['Event'] = _EVENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FileHeader = _reflection.GeneratedProtocolMessageType('FileHeader', (_message.Message,), dict(
  DESCRIPTOR = _FILEHEADER,
  __module__ = 'file.file_pb2'
  # @@protoc_insertion_point(class_scope:file.FileHeader)
  ))
_sym_db.RegisterMessage(FileHeader)

Id = _reflection.GeneratedProtocolMessageType('Id', (_message.Message,), dict(
  DESCRIPTOR = _ID,
  __module__ = 'file.file_pb2'
  # @@protoc_insertion_point(class_scope:file.Id)
  ))
_sym_db.RegisterMessage(Id)

PresignResult = _reflection.GeneratedProtocolMessageType('PresignResult', (_message.Message,), dict(
  DESCRIPTOR = _PRESIGNRESULT,
  __module__ = 'file.file_pb2'
  # @@protoc_insertion_point(class_scope:file.PresignResult)
  ))
_sym_db.RegisterMessage(PresignResult)

File = _reflection.GeneratedProtocolMessageType('File', (_message.Message,), dict(
  DESCRIPTOR = _FILE,
  __module__ = 'file.file_pb2'
  # @@protoc_insertion_point(class_scope:file.File)
  ))
_sym_db.RegisterMessage(File)


# @@protoc_insertion_point(module_scope)