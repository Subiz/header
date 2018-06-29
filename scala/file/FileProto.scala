// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package file

object FileProto extends _root_.scalapb.GeneratedFileObject {
  lazy val dependencies: Seq[_root_.scalapb.GeneratedFileObject] = Seq(
    common.CommonProto
  )
  lazy val messagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq(
    file.AllType,
    file.FileHeader,
    file.PresignResult,
    file.FileRequest,
    file.File
  )
  private lazy val ProtoBytes: Array[Byte] =
      scalapb.Encoding.fromBase64(scala.collection.Seq(
  """CipiaXRidWNrZXQub3JnL3N1Yml6L2hlYWRlci9maWxlL2ZpbGUucHJvdG8SBGZpbGUaLmJpdGJ1Y2tldC5vcmcvc3ViaXova
  GVhZGVyL2NvbW1vbi9jb21tb24ucHJvdG8iagoHQWxsVHlwZRIgCgJmaBgCIAEoCzIQLmZpbGUuRmlsZUhlYWRlclICZmgSIwoCc
  HMYAyABKAsyEy5maWxlLlByZXNpZ25SZXN1bHRSAnBzEhgKAWYYBCABKAsyCi5maWxlLkZpbGVSAWYivgEKCkZpbGVIZWFkZXISI
  QoDY3R4GAEgASgLMg8uY29tbW9uLkNvbnRleHRSA2N0eBIdCgphY2NvdW50X2lkGAIgASgJUglhY2NvdW50SWQSEgoEbmFtZRgDI
  AEoCVIEbmFtZRISCgR0eXBlGAQgASgJUgR0eXBlEhIKBHNpemUYBSABKANSBHNpemUSEAoDbWQ1GAYgASgJUgNtZDUSIAoLZGVzY
  3JpcHRpb24YByABKAlSC2Rlc2NyaXB0aW9uIpIBCg1QcmVzaWduUmVzdWx0EiEKA2N0eBgBIAEoCzIPLmNvbW1vbi5Db250ZXh0U
  gNjdHgSHQoKYWNjb3VudF9pZBgDIAEoCVIJYWNjb3VudElkEhAKA3VybBgEIAEoCVIDdXJsEg4KAmlkGAYgASgJUgJpZBIdCgpza
  WduZWRfdXJsGAUgASgJUglzaWduZWRVcmwiXwoLRmlsZVJlcXVlc3QSIQoDY3R4GAEgASgLMg8uY29tbW9uLkNvbnRleHRSA2N0e
  BIdCgphY2NvdW50X2lkGAIgASgJUglhY2NvdW50SWQSDgoCaWQYAyABKAlSAmlkIo4CCgRGaWxlEiEKA2N0eBgBIAEoCzIPLmNvb
  W1vbi5Db250ZXh0UgNjdHgSHQoKYWNjb3VudF9pZBgCIAEoCVIJYWNjb3VudElkEhIKBG5hbWUYAyABKAlSBG5hbWUSEgoEdHlwZ
  RgEIAEoCVIEdHlwZRISCgRzaXplGAUgASgDUgRzaXplEhAKA21kNRgGIAEoCVIDbWQ1EiAKC2Rlc2NyaXB0aW9uGAogASgJUgtkZ
  XNjcmlwdGlvbhIYCgdjcmVhdGVkGAcgASgDUgdjcmVhdGVkEhAKA3VybBgIIAEoCVIDdXJsEhgKB2NyZWF0b3IYCSABKAlSB2NyZ
  WF0b3ISDgoCaWQYCyABKAlSAmlkKm4KBUV2ZW50EhgKFEZpbGVQcmVzaWduUmVxdWVzdGVkEAASDwoLRmlsZUNyZWF0ZWQQAxIVC
  hFGaWxlSW5mb1JlcXVlc3RlZBAEEhAKDEZpbGVVcGxvYWRlZBAGEhEKDUZpbGVSZXF1ZXN0ZWQQCjIwCghNeVNlcnZlchIkCgJEb
  xINLmZpbGUuQWxsVHlwZRoNLmZpbGUuQWxsVHlwZSIAMpMBCgdGaWxlTWdyEjIKB1ByZXNpZ24SEC5maWxlLkZpbGVIZWFkZXIaE
  y5maWxlLlByZXNpZ25SZXN1bHQiABInCgRSZWFkEhEuZmlsZS5GaWxlUmVxdWVzdBoKLmZpbGUuRmlsZSIAEisKCFVwbG9hZGVkE
  hEuZmlsZS5GaWxlUmVxdWVzdBoKLmZpbGUuRmlsZSIAYgZwcm90bzM="""
      ).mkString)
  lazy val scalaDescriptor: _root_.scalapb.descriptors.FileDescriptor = {
    val scalaProto = com.google.protobuf.descriptor.FileDescriptorProto.parseFrom(ProtoBytes)
    _root_.scalapb.descriptors.FileDescriptor.buildFrom(scalaProto, dependencies.map(_.scalaDescriptor))
  }
  lazy val javaDescriptor: com.google.protobuf.Descriptors.FileDescriptor = {
    val javaProto = com.google.protobuf.DescriptorProtos.FileDescriptorProto.parseFrom(ProtoBytes)
    com.google.protobuf.Descriptors.FileDescriptor.buildFrom(javaProto, Array(
      common.CommonProto.javaDescriptor
    ))
  }
  @deprecated("Use javaDescriptor instead. In a future version this will refer to scalaDescriptor.", "ScalaPB 0.5.47")
  def descriptor: com.google.protobuf.Descriptors.FileDescriptor = javaDescriptor
}