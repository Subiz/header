// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package auth

sealed trait AuthorizationType extends _root_.scalapb.GeneratedEnum {
  type EnumType = AuthorizationType
  def isclientauth: _root_.scala.Boolean = false
  def ischannelauth: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[AuthorizationType] = auth.AuthorizationType
}

object AuthorizationType extends _root_.scalapb.GeneratedEnumCompanion[AuthorizationType] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[AuthorizationType] = this
  @SerialVersionUID(0L)
  case object client_auth extends AuthorizationType {
    val value = 0
    val index = 0
    val name = "client_auth"
    override def isclientauth: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object channel_auth extends AuthorizationType {
    val value = 1
    val index = 1
    val name = "channel_auth"
    override def ischannelauth: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends AuthorizationType with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(client_auth, channel_auth)
  def fromValue(value: _root_.scala.Int): AuthorizationType = value match {
    case 0 => client_auth
    case 1 => channel_auth
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = AuthProto.javaDescriptor.getEnumTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = AuthProto.scalaDescriptor.enums(1)
}