// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package widget

object WidgetProto extends _root_.scalapb.GeneratedFileObject {
  lazy val dependencies: Seq[_root_.scalapb.GeneratedFileObject] = Seq(
    account.AccountProto,
    common.CommonProto,
    user.UserProto
  )
  lazy val messagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq(
    widget.AllType,
    widget.Theme,
    widget.Sound,
    widget.Setting,
    widget.UserSetting,
    widget.Global
  )
  private lazy val ProtoBytes: Array[Byte] =
      scalapb.Encoding.fromBase64(scala.collection.Seq(
  """Ci5iaXRidWNrZXQub3JnL3N1Yml6L2hlYWRlci93aWRnZXQvd2lkZ2V0LnByb3RvEgZ3aWRnZXQaMGJpdGJ1Y2tldC5vcmcvc
  3ViaXovaGVhZGVyL2FjY291bnQvYWNjb3VudC5wcm90bxouYml0YnVja2V0Lm9yZy9zdWJpei9oZWFkZXIvY29tbW9uL2NvbW1vb
  i5wcm90bxoqYml0YnVja2V0Lm9yZy9zdWJpei9oZWFkZXIvdXNlci91c2VyLnByb3RvIqMBCgdBbGxUeXBlEiMKBXRoZW1lGAIgA
  SgLMg0ud2lkZ2V0LlRoZW1lUgV0aGVtZRIjCgVzb3VuZBgDIAEoCzINLndpZGdldC5Tb3VuZFIFc291bmQSKQoHc2V0dGluZxgEI
  AEoCzIPLndpZGdldC5TZXR0aW5nUgdzZXR0aW5nEiMKAnVzGAUgASgLMhMud2lkZ2V0LlVzZXJTZXR0aW5nUgJ1cyLYAQoFVGhlb
  WUSHQoKYWNjb3VudF9pZBgCIAEoCVIJYWNjb3VudElkEh0KCmN1c3RvbV9jc3MYCCABKAlSCWN1c3RvbUNzcxInCg93aWRnZXRfc
  G9zaXRpb24YAyABKAlSDndpZGdldFBvc2l0aW9uEh8KC3dpbmRvd19tb2RlGAUgASgJUgp3aW5kb3dNb2RlIiUKDkJ1dHRvblBvc
  2l0aW9uEggKBGxlZnQQABIJCgVyaWdodBABIiAKCldpbmRvd01vZGUSCAoEbWluaRAAEggKBGZ1bGwQASLdAQoFU291bmQSHQoKY
  WNjb3VudF9pZBgBIAEoCVIJYWNjb3VudElkEhgKB2VuYWJsZWQYAiABKAhSB2VuYWJsZWQSKQoQbmV3X2NvbnZlcnNhdGlvbhgDI
  AEoCVIPbmV3Q29udmVyc2F0aW9uEh8KC2ZpbGVfY3JlYXRlGAQgASgJUgpmaWxlQ3JlYXRlEh8KC25ld19tZXNzYWdlGAUgASgJU
  gpuZXdNZXNzYWdlEi4KE21lc3NhZ2Vfc2VuZF9mYWlsZWQYBiABKAlSEW1lc3NhZ2VTZW5kRmFpbGVkIvYDCgdTZXR0aW5nEiEKA
  2N0eBgBIAEoCzIPLmNvbW1vbi5Db250ZXh0UgNjdHgSHQoKYWNjb3VudF9pZBgCIAEoCVIJYWNjb3VudElkEiUKDndpZGdldF92Z
  XJzaW9uGAMgASgJUg13aWRnZXRWZXJzaW9uEiMKBXNvdW5kGAQgASgLMg0ud2lkZ2V0LlNvdW5kUgVzb3VuZBIaCghsYW5ndWFnZ
  RgFIAEoCVIIbGFuZ3VhZ2USIwoFdGhlbWUYByABKAsyDS53aWRnZXQuVGhlbWVSBXRoZW1lEhwKCXJlcGx5dGltZRgJIAEoBVIJc
  mVwbHl0aW1lEiYKBmFnZW50cxgKIAMoCzIOLmFjY291bnQuQWdlbnRSBmFnZW50cxIbCglhZ2VudF9pZHMYCyADKAlSCGFnZW50S
  WRzEiEKDGxhbmd1YWdlX3VybBgNIAEoCVILbGFuZ3VhZ2VVcmwSLgoTY3VzdG9tX2xhbmd1YWdlX3VybBgOIAEoCVIRY3VzdG9tT
  GFuZ3VhZ2VVcmwSFwoHY3NzX3VybBgMIAEoCVIGY3NzVXJsEiQKDmN1c3RvbV9jc3NfdXJsGA8gASgJUgxjdXN0b21Dc3NVcmwSJ
  woPY3VzdG9tX2xhbmd1YWdlGBAgASgJUg5jdXN0b21MYW5ndWFnZSKLAwoLVXNlclNldHRpbmcSIQoDY3R4GAEgASgLMg8uY29tb
  W9uLkNvbnRleHRSA2N0eBIqCgdhY2NvdW50GAMgASgLMhAuYWNjb3VudC5BY2NvdW50UgdhY2NvdW50Eh0KCmFjY291bnRfaWQYA
  iABKAlSCWFjY291bnRJZBIeCgR1c2VyGAQgASgLMgoudXNlci5Vc2VyUgR1c2VyEhcKB3VzZXJfaWQYByABKAlSBnVzZXJJZBIjC
  g1zb3VuZF9lbmFibGVkGAUgASgIUgxzb3VuZEVuYWJsZWQSGgoIbGFuZ3VhZ2UYBiABKAlSCGxhbmd1YWdlEicKD3NlbmRfdHJhb
  nNjcmlwdBgIIAEoCFIOc2VuZFRyYW5zY3JpcHQSOAoPYWNjb3VudF9zZXR0aW5nGAkgASgLMg8ud2lkZ2V0LlNldHRpbmdSDmFjY
  291bnRTZXR0aW5nEjEKFGRlc2t0b3Bfbm90aWZpY2F0aW9uGAogASgIUhNkZXNrdG9wTm90aWZpY2F0aW9uImMKBkdsb2JhbBIhC
  gNjdHgYASABKAsyDy5jb21tb24uQ29udGV4dFIDY3R4EhIKBG5hbWUYAiABKAlSBG5hbWUSEgoEZGF0YRgDIAEoCVIEZGF0YRIOC
  gJwaxgEIAEoCVICcGsqhQMKBUV2ZW50EiIKHldpZGdldFVzZXJTZXR0aW5nUmVhZFJlcXVlc3RlZBAAEiQKIFdpZGdldFVzZXJTZ
  XR0aW5nVXBkYXRlUmVxdWVzdGVkEAESHgoaV2lkZ2V0U2V0dGluZ1JlYWRSZXF1ZXN0ZWQQAhIgChxXaWRnZXRTZXR0aW5nVXBkY
  XRlUmVxdWVzdGVkEAMSIgoeV2lkZ2V0U2V0dGluZ0Nzc1ZlcnNpb25VcGRhdGVkEAQSIAocV2lkZ2V0U2V0dGluZ0xhbmd1YWdlV
  XBkYXRlZBAFEioKJldpZGdldFNldHRpbmdDc3NWZXJzaW9uVXBkYXRlUmVxdWVzdGVkEAYSKAokV2lkZ2V0U2V0dGluZ0xhbmd1Y
  WdlVXBkYXRlUmVxdWVzdGVkEAcSEwoPV2lkZ2V0UmVxdWVzdGVkEAgSGQoVV2lkZ2V0U2V0dGluZ1Vwc2VydGVkEAoSEAoMV2lkZ
  2V0U3luY2VkEGQSEgoOV2lkZ2V0VjNTeW5jZWQQZzI0CghNeVNlcnZlchIoCgJEbxIPLndpZGdldC5BbGxUeXBlGg8ud2lkZ2V0L
  kFsbFR5cGUiADLbAQoNV2lkZ2V0U2VydmljZRIlCgRSZWFkEgouY29tbW9uLklkGg8ud2lkZ2V0LlNldHRpbmciABIsCgZVcGRhd
  GUSDy53aWRnZXQuU2V0dGluZxoPLndpZGdldC5TZXR0aW5nIgASNAoPUmVhZFVzZXJTZXR0aW5nEgouY29tbW9uLklkGhMud2lkZ
  2V0LlVzZXJTZXR0aW5nIgASPwoRVXBkYXRlVXNlclNldHRpbmcSEy53aWRnZXQuVXNlclNldHRpbmcaEy53aWRnZXQuVXNlclNld
  HRpbmciAA=="""
      ).mkString)
  lazy val scalaDescriptor: _root_.scalapb.descriptors.FileDescriptor = {
    val scalaProto = com.google.protobuf.descriptor.FileDescriptorProto.parseFrom(ProtoBytes)
    _root_.scalapb.descriptors.FileDescriptor.buildFrom(scalaProto, dependencies.map(_.scalaDescriptor))
  }
  lazy val javaDescriptor: com.google.protobuf.Descriptors.FileDescriptor = {
    val javaProto = com.google.protobuf.DescriptorProtos.FileDescriptorProto.parseFrom(ProtoBytes)
    com.google.protobuf.Descriptors.FileDescriptor.buildFrom(javaProto, Array(
      account.AccountProto.javaDescriptor,
      common.CommonProto.javaDescriptor,
      user.UserProto.javaDescriptor
    ))
  }
  @deprecated("Use javaDescriptor instead. In a future version this will refer to scalaDescriptor.", "ScalaPB 0.5.47")
  def descriptor: com.google.protobuf.Descriptors.FileDescriptor = javaDescriptor
}