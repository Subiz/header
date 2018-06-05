// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package template

sealed trait Type extends _root_.scalapb.GeneratedEnum {
  type EnumType = Type
  def isEmail: _root_.scala.Boolean = false
  def isWebPush: _root_.scala.Boolean = false
  def isNotification: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[Type] = template.Type
}

object Type extends _root_.scalapb.GeneratedEnumCompanion[Type] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Type] = this
  @SerialVersionUID(0L)
  case object Email extends Type {
    val value = 0
    val index = 0
    val name = "Email"
    override def isEmail: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WebPush extends Type {
    val value = 1
    val index = 1
    val name = "WebPush"
    override def isWebPush: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object Notification extends Type {
    val value = 2
    val index = 2
    val name = "Notification"
    override def isNotification: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends Type with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(Email, WebPush, Notification)
  def fromValue(value: _root_.scala.Int): Type = value match {
    case 0 => Email
    case 1 => WebPush
    case 2 => Notification
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = TemplateProto.javaDescriptor.getEnumTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = TemplateProto.scalaDescriptor.enums(0)
}