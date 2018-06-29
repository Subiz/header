// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

sealed trait AutomationScope extends _root_.scalapb.GeneratedEnum {
  type EnumType = AutomationScope
  def isconversation: _root_.scala.Boolean = false
  def isuser: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[AutomationScope] = user.AutomationScope
}

object AutomationScope extends _root_.scalapb.GeneratedEnumCompanion[AutomationScope] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[AutomationScope] = this
  @SerialVersionUID(0L)
  case object conversation extends AutomationScope {
    val value = 2
    val index = 0
    val name = "conversation"
    override def isconversation: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object user extends AutomationScope {
    val value = 3
    val index = 1
    val name = "user"
    override def isuser: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends AutomationScope with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(conversation, user)
  def fromValue(value: _root_.scala.Int): AutomationScope = value match {
    case 2 => conversation
    case 3 => user
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = UserProto.javaDescriptor.getEnumTypes.get(4)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = UserProto.scalaDescriptor.enums(4)
}