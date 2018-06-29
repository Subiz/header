// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

sealed trait AttributeDataState extends _root_.scalapb.GeneratedEnum {
  type EnumType = AttributeDataState
  def islive: _root_.scala.Boolean = false
  def isdeleted: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[AttributeDataState] = user.AttributeDataState
}

object AttributeDataState extends _root_.scalapb.GeneratedEnumCompanion[AttributeDataState] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[AttributeDataState] = this
  @SerialVersionUID(0L)
  case object live extends AttributeDataState {
    val value = 0
    val index = 0
    val name = "live"
    override def islive: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object deleted extends AttributeDataState {
    val value = 1
    val index = 1
    val name = "deleted"
    override def isdeleted: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends AttributeDataState with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(live, deleted)
  def fromValue(value: _root_.scala.Int): AttributeDataState = value match {
    case 0 => live
    case 1 => deleted
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = UserProto.javaDescriptor.getEnumTypes.get(2)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = UserProto.scalaDescriptor.enums(2)
}