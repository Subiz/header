// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package scheduler

object SchedulerProto extends _root_.scalapb.GeneratedFileObject {
  lazy val dependencies: Seq[_root_.scalapb.GeneratedFileObject] = Seq(
  )
  lazy val messagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq(
    scheduler.Task,
    scheduler.Id
  )
  private lazy val ProtoBytes: Array[Byte] =
      scalapb.Encoding.fromBase64(scala.collection.Seq(
  """CjRiaXRidWNrZXQub3JnL3N1Yml6L2hlYWRlci9zY2hlZHVsZXIvc2NoZWR1bGVyLnByb3RvEglzY2hlZHVsZXIiswEKBFRhc
  2sSDgoCaWQYAiABKAlSAmlkEiMKDWNhbGxiYWNrX3RpbWUYAyABKANSDGNhbGxiYWNrVGltZRIUCgV0b3BpYxgEIAEoCVIFdG9wa
  WMSEgoEZGF0YRgFIAEoDFIEZGF0YRIQCgNrZXkYBiABKAlSA2tleRIWCgZjYWxsZWQYCCABKANSBmNhbGxlZBIQCgNzZWMYCSABK
  ANSA3NlYxIQCgNwYXIYCiABKAlSA3BhciI5CgJJZBIOCgJpZBgBIAEoCVICaWQSIwoNY2FsbGJhY2tfdGltZRgCIAEoA1IMY2Fsb
  GJhY2tUaW1lKjcKBUV2ZW50EhYKElNjaGVkdWxlclJlcXVlc3RlZBAAEhYKElNjaGVkdWxlckNhbmNlbGxlZBABYgZwcm90bzM="""
      ).mkString)
  lazy val scalaDescriptor: _root_.scalapb.descriptors.FileDescriptor = {
    val scalaProto = com.google.protobuf.descriptor.FileDescriptorProto.parseFrom(ProtoBytes)
    _root_.scalapb.descriptors.FileDescriptor.buildFrom(scalaProto, dependencies.map(_.scalaDescriptor))
  }
  lazy val javaDescriptor: com.google.protobuf.Descriptors.FileDescriptor = {
    val javaProto = com.google.protobuf.DescriptorProtos.FileDescriptorProto.parseFrom(ProtoBytes)
    com.google.protobuf.Descriptors.FileDescriptor.buildFrom(javaProto, Array(
    ))
  }
  @deprecated("Use javaDescriptor instead. In a future version this will refer to scalaDescriptor.", "ScalaPB 0.5.47")
  def descriptor: com.google.protobuf.Descriptors.FileDescriptor = javaDescriptor
}