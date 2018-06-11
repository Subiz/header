// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package auth

object AuthProto extends _root_.scalapb.GeneratedFileObject {
  lazy val dependencies: Seq[_root_.scalapb.GeneratedFileObject] = Seq(
  )
  lazy val messagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq(
    auth.Credential,
    auth.Scope,
    auth.Method,
    auth.UserAuth,
    auth.PasswordCredential,
    auth.SuperPasswordCredential,
    auth.AuthCookie,
    auth.OauthAccessToken,
    auth.AccessToken,
    auth.CookieExpire,
    auth.User
  )
  private lazy val ProtoBytes: Array[Byte] =
      scalapb.Encoding.fromBase64(scala.collection.Seq(
  """CipiaXRidWNrZXQub3JnL3N1Yml6L2hlYWRlci9hdXRoL2F1dGgucHJvdG8SBGF1dGgi4AIKCkNyZWRlbnRpYWwSHQoKYWNjb
  3VudF9pZBgBIAEoCVIJYWNjb3VudElkEhYKBmlzc3VlchgDIAEoCVIGaXNzdWVyEh4KBHR5cGUYBCABKA4yCi5hdXRoLlR5cGVSB
  HR5cGUSJAoGbWV0aG9kGAUgASgLMgwuYXV0aC5NZXRob2RSBm1ldGhvZBIbCgljbGllbnRfaWQYByABKAlSCGNsaWVudElkEisKC
  2NsaWVudF90eXBlGAggASgOMgouYXV0aC5UeXBlUgpjbGllbnRUeXBlEioKEWNsaWVudF9hY2NvdW50X2lkGAogASgJUg9jbGllb
  nRBY2NvdW50SWQSFgoGc2NvcGVzGAkgAygJUgZzY29wZXMSHQoKYXZhdGFyX3VybBgPIAEoCVIJYXZhdGFyVXJsEhIKBG5hbWUYE
  CABKAlSBG5hbWUSFAoFZW1haWwYESABKAlSBWVtYWlsIroBCgVTY29wZRIOCgJpZBgBIAEoCVICaWQSEgoEbmFtZRgCIAEoCVIEb
  mFtZRIZCghsb2dvX3VybBgDIAEoCVIHbG9nb1VybBIUCgV0aXRsZRgEIAEoCVIFdGl0bGUSIAoLZGVzY3JpcHRpb24YBSABKAlSC
  2Rlc2NyaXB0aW9uEiQKBm1ldGhvZBgGIAEoCzIMLmF1dGguTWV0aG9kUgZtZXRob2QSFAoFZXZlbnQYByADKAlSBWV2ZW50IrUWC
  gZNZXRob2QSEgoEcGluZxgBIAEoCFIEcGluZxIlCg51cGRhdGVfdHJpZ2dlchgyIAEoCFINdXBkYXRlVHJpZ2dlchIlCg5kZWxld
  GVfdHJpZ2dlchgzIAEoCFINZGVsZXRlVHJpZ2dlchIlCg5jcmVhdGVfdHJpZ2dlchg0IAEoCFINY3JlYXRlVHJpZ2dlchIhCgxyZ
  WFkX3RyaWdnZXIYNSABKAhSC3JlYWRUcmlnZ2VyEisKEXJlYWRfc2VnbWVudGF0aW9uGDwgASgIUhByZWFkU2VnbWVudGF0aW9uE
  i8KE3VwZGF0ZV9zZWdtZW50YXRpb24YPSABKAhSEnVwZGF0ZVNlZ21lbnRhdGlvbhIvChNkZWxldGVfc2VnbWVudGF0aW9uGD4gA
  SgIUhJkZWxldGVTZWdtZW50YXRpb24SLwoTY3JlYXRlX3NlZ21lbnRhdGlvbhg/IAEoCFISY3JlYXRlU2VnbWVudGF0aW9uEiEKD
  Gludml0ZV9hZ2VudBhGIAEoCFILaW52aXRlQWdlbnQSIQoMdXBkYXRlX2FnZW50GEggASgIUgt1cGRhdGVBZ2VudBIjCg11cGRhd
  GVfYWdlbnRzGEkgASgIUgx1cGRhdGVBZ2VudHMSHQoKcmVhZF9hZ2VudBhMIAEoCFIJcmVhZEFnZW50Eh8KC3JlYWRfYWdlbnRzG
  E0gASgIUgpyZWFkQWdlbnRzEiUKDnJlc2V0X3Bhc3N3b3JkGE4gASgIUg1yZXNldFBhc3N3b3JkEjgKGHVwZGF0ZV9hZ2VudHNfc
  GVybWlzc2lvbhhRIAEoCFIWdXBkYXRlQWdlbnRzUGVybWlzc2lvbhIyChVyZWFkX2FnZW50X3Blcm1pc3Npb24YVSABKAhSE3JlY
  WRBZ2VudFBlcm1pc3Npb24SLgoTdXBkYXRlX2FnZW50c19zdGF0ZRhWIAEoCFIRdXBkYXRlQWdlbnRzU3RhdGUSIQoMcmVhZF9hY
  2NvdW50GHggASgIUgtyZWFkQWNjb3VudBIsChJjcmVhdGVfYWdlbnRfZ3JvdXAYfCABKAhSEGNyZWF0ZUFnZW50R3JvdXASLAoSZ
  GVsZXRlX2FnZW50X2dyb3VwGH0gASgIUhBkZWxldGVBZ2VudEdyb3VwEigKEHJlYWRfYWdlbnRfZ3JvdXAYfiABKAhSDnJlYWRBZ
  2VudEdyb3VwEiwKEnVwZGF0ZV9hZ2VudF9ncm91cBh/IAEoCFIQdXBkYXRlQWdlbnRHcm91cBIfCgt1cGRhdGVfcGxhbhh7IAEoC
  FIKdXBkYXRlUGxhbhI6Chl1cGRhdGVfYWNjb3VudF9pbmZvbWF0aW9uGHogASgIUhd1cGRhdGVBY2NvdW50SW5mb21hdGlvbhIgC
  gtyZWFkX2NsaWVudBiXASABKAhSCnJlYWRDbGllbnQSJAoNdXBkYXRlX2NsaWVudBiZASABKAhSDHVwZGF0ZUNsaWVudBIkCg1kZ
  WxldGVfY2xpZW50GJ0BIAEoCFIMZGVsZXRlQ2xpZW50EiQKDWNyZWF0ZV9jbGllbnQYngEgASgIUgxjcmVhdGVDbGllbnQSHAoJc
  mVhZF9ydWxlGLQBIAEoCFIIcmVhZFJ1bGUSIAoLY3JlYXRlX3J1bGUYtQEgASgIUgpjcmVhdGVSdWxlEiAKC2RlbGV0ZV9ydWxlG
  LYBIAEoCFIKZGVsZXRlUnVsZRIgCgt1cGRhdGVfcnVsZRi3ASABKAhSCnVwZGF0ZVJ1bGUSLgoSc3RhcnRfY29udmVyc2F0aW9uG
  L4BIAEoCFIRc3RhcnRDb252ZXJzYXRpb24SLAoRcmVhZF9jb252ZXJzYXRpb24YvwEgASgIUhByZWFkQ29udmVyc2F0aW9uEjIKF
  GV4cG9ydF9jb252ZXJzYXRpb25zGMABIAEoCFITZXhwb3J0Q29udmVyc2F0aW9ucxJBChxyZWFkX3RlYW1tYXRlc19jb252ZXJzY
  XRpb25zGMEBIAEoCFIacmVhZFRlYW1tYXRlc0NvbnZlcnNhdGlvbnMSIgoMc2VuZF9tZXNzYWdlGMIBIAEoCFILc2VuZE1lc3NhZ
  2USMAoTaW50ZWdyYXRlX2Nvbm5lY3RvchjIASABKAhSEmludGVncmF0ZUNvbm5lY3RvchInCg9yZWFkX3VzZXJfZW1haWwYzQEgA
  SgIUg1yZWFkVXNlckVtYWlsEjIKFXJlYWRfdXNlcl9mYWNlYm9va19pZBjOASABKAhSEnJlYWRVc2VyRmFjZWJvb2tJZBIpChByZ
  WFkX3VzZXJfcGhvbmVzGM8BIAEoCFIOcmVhZFVzZXJQaG9uZXMSOAoYcmVhZF91c2VyX3dpZGdldF9zZXR0aW5nGNABIAEoCFIVc
  mVhZFVzZXJXaWRnZXRTZXR0aW5nEhoKCHJlYWRfdGFnGNEBIAEoCFIHcmVhZFRhZxIeCgp1cGRhdGVfdGFnGNIBIAEoCFIJdXBkY
  XRlVGFnEh4KCmRlbGV0ZV90YWcY0wEgASgIUglkZWxldGVUYWcSMwoVdXBkYXRlX3dpZGdldF9zZXR0aW5nGKsCIAEoCFITdXBkY
  XRlV2lkZ2V0U2V0dGluZxI3ChdjcmVhdGVfd2hpdGVsaXN0X2RvbWFpbhisAiABKAhSFWNyZWF0ZVdoaXRlbGlzdERvbWFpbhIvC
  hNjcmVhdGVfd2hpdGVsaXN0X2lwGK0CIAEoCFIRY3JlYXRlV2hpdGVsaXN0SXASMwoVY3JlYXRlX3doaXRlbGlzdF91c2VyGK4CI
  AEoCFITY3JlYXRlV2hpdGVsaXN0VXNlchI3ChdkZWxldGVfd2hpdGVsaXN0X2RvbWFpbhivAiABKAhSFWRlbGV0ZVdoaXRlbGlzd
  ERvbWFpbhIvChNkZWxldGVfd2hpdGVsaXN0X2lwGLACIAEoCFIRZGVsZXRlV2hpdGVsaXN0SXASMwoVZGVsZXRlX3doaXRlbGlzd
  F91c2VyGLECIAEoCFITZGVsZXRlV2hpdGVsaXN0VXNlchIrChFyZWFkX3doaXRlbGlzdF9pcBiyAiABKAhSD3JlYWRXaGl0ZWxpc
  3RJcBIvChNyZWFkX3doaXRlbGlzdF91c2VyGLMCIAEoCFIRcmVhZFdoaXRlbGlzdFVzZXISMwoVcmVhZF93aGl0ZWxpc3RfZG9tY
  WluGLQCIAEoCFITcmVhZFdoaXRlbGlzdERvbWFpbhIqChBwdXJjaGFzZV9zZXJ2aWNlGLYCIAEoCFIPcHVyY2hhc2VTZXJ2aWNlE
  jMKFXVwZGF0ZV9wYXltZW50X21ldGhvZBi3AiABKAhSE3VwZGF0ZVBheW1lbnRNZXRob2QSHgoKYWRkX2NyZWRpdBi4AiABKAhSC
  WFkZENyZWRpdBIxChR1cGRhdGVfYmlsbGluZ19jeWNsZRi5AiABKAhSEnVwZGF0ZUJpbGxpbmdDeWNsZRIkCg1yZWFkX2ludm9pY
  2VzGLoCIAEoCFIMcmVhZEludm9pY2VzEi4KEnJlYWRfc3Vic2NyaXB0aW9ucxi7AiABKAhSEXJlYWRTdWJzY3JpcHRpb25zEiYKD
  nJlYWRfYXR0cmlidXRlGLwCIAEoCFINcmVhZEF0dHJpYnV0ZRIqChBjcmVhdGVfYXR0cmlidXRlGL0CIAEoCFIPY3JlYXRlQXR0c
  mlidXRlEioKEHVwZGF0ZV9hdHRyaWJ1dGUYvgIgASgIUg91cGRhdGVBdHRyaWJ1dGUSKgoQZGVsZXRlX2F0dHJpYnV0ZRi/AiABK
  AhSD2RlbGV0ZUF0dHJpYnV0ZSJJCghVc2VyQXV0aBIXCgd1c2VyX2lkGAEgASgJUgZ1c2VySWQSJAoGbWV0aG9kGAIgASgLMgwuY
  XV0aC5NZXRob2RSBm1ldGhvZCJMChJQYXNzd29yZENyZWRlbnRpYWwSGgoIdXNlcm5hbWUYASABKAlSCHVzZXJuYW1lEhoKCHBhc
  3N3b3JkGAIgASgJUghwYXNzd29yZCL2AQoXU3VwZXJQYXNzd29yZENyZWRlbnRpYWwSJQoOc3ViaXpfdXNlcm5hbWUYASABKAlSD
  XN1Yml6VXNlcm5hbWUSJQoOc3ViaXpfcGFzc3dvcmQYAiABKAlSDXN1Yml6UGFzc3dvcmQSHwoLc3ViaXpfdG9rZW4YAyABKAlSC
  nN1Yml6VG9rZW4SGwoJaXNzdWVyX2lkGAUgASgJUghpc3N1ZXJJZBIdCgphY2NvdW50X2lkGAkgASgJUglhY2NvdW50SWQSGAoHZ
  XhwaXJlZBgLIAEoBVIHZXhwaXJlZBIWCgZzY29wZXMYDCADKAlSBnNjb3BlcyKKAQoKQXV0aENvb2tpZRIXCgd1c2VyX2lkGAEgA
  SgJUgZ1c2VySWQSHQoKYWNjb3VudF9pZBgCIAEoCVIJYWNjb3VudElkEhgKB2V4cGlyZWQYBCABKAVSB2V4cGlyZWQSFgoGaXNzd
  WVkGAUgASgFUgZpc3N1ZWQSEgoEdHlwZRgDIAEoCVIEdHlwZSKkAgoQT2F1dGhBY2Nlc3NUb2tlbhIhCgxhY2Nlc3NfdG9rZW4YA
  SABKAlSC2FjY2Vzc1Rva2VuEh0KCnRva2VuX3R5cGUYAyABKAlSCXRva2VuVHlwZRIdCgpleHBpcmVzX2luGAQgASgDUglleHBpc
  mVzSW4SIwoNcmVmcmVzaF90b2tlbhgFIAEoCVIMcmVmcmVzaFRva2VuEhQKBXNjb3BlGAYgASgJUgVzY29wZRIUCgVzdGF0ZRgHI
  AEoCVIFc3RhdGUSFAoFZXJyb3IYCCABKAlSBWVycm9yEisKEWVycm9yX2Rlc2NyaXB0aW9uGAkgASgJUhBlcnJvckRlc2NyaXB0a
  W9uEhsKCWVycm9yX3VyaRgKIAEoCVIIZXJyb3JVcmkihgIKC0FjY2Vzc1Rva2VuEhsKCWlzc3Vlcl9pZBgBIAEoCVIIaXNzdWVyS
  WQSHwoLaXNzdWVyX3R5cGUYByABKAlSCmlzc3VlclR5cGUSGwoJY2xpZW50X2lkGAIgASgJUghjbGllbnRJZBIfCgtjbGllbnRfd
  HlwZRgDIAEoCVIKY2xpZW50VHlwZRIdCgphY2NvdW50X2lkGAQgASgJUglhY2NvdW50SWQSKgoRY2xpZW50X2FjY291bnRfaWQYB
  SABKAlSD2NsaWVudEFjY291bnRJZBIYCgdleHBpcmVkGAggASgFUgdleHBpcmVkEhYKBnNjb3BlcxgGIAMoCVIGc2NvcGVzIq8BC
  gxDb29raWVFeHBpcmUSFwoHdXNlcl9pZBgBIAEoCVIGdXNlcklkEh0KCmFjY291bnRfaWQYAiABKAlSCWFjY291bnRJZBIjCg1le
  HBpcmVkX3Rva2VuGAMgASgJUgxleHBpcmVkVG9rZW4SHwoLYmVmb3JlX3RpbWUYBCABKANSCmJlZm9yZVRpbWUSIQoMZXhjZXB0X
  3Rva2VuGAUgASgJUgtleGNlcHRUb2tlbiKcAgoEVXNlchIOCgJpZBgCIAEoCVICaWQSHQoKYWNjb3VudF9pZBgDIAEoCVIJYWNjb
  3VudElkEhQKBWVtYWlsGAUgASgJUgVlbWFpbBItChJlbmNyeXB0ZWRfcGFzc3dvcmQYESABKAlSEWVuY3J5cHRlZFBhc3N3b3JkE
  hsKCWlzX2FjdGl2ZRgSIAEoCFIIaXNBY3RpdmUSGgoIdXBzZXJ0ZWQYEyABKANSCHVwc2VydGVkEhkKCHYzX3N0YXRlGBQgASgFU
  gd2M1N0YXRlSgQIARACSgQIBBAFSgQIBhAHSgQIBxAISgQICBAJSgQICRAKSgQIChALSgQICxAMSgQIDBANSgQIDRAOSgQIDhAPS
  gQIDxAQSgQIEBARKm0KBFR5cGUSCwoHdW5rbm93bhAAEggKBHVzZXIQARIJCgVhZ2VudBACEgkKBXN1Yml6EAMSBwoDYXBwEAUSC
  AoEcmVzdBAIEg0KCWNvbm5lY3RvchAGEgcKA2JvdBAHEg0KCWRhc2hib2FyZBAKKjYKEUF1dGhvcml6YXRpb25UeXBlEg8KC2Nsa
  WVudF9hdXRoEAASEAoMY2hhbm5lbF9hdXRoEAEqJwoFRXZlbnQSCAoEQVVUSBAAEhQKEEF1dGhFeHBpcmVDb29raWUQBGIGcHJvd
  G8z"""
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