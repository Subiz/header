// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package api

sealed trait Method extends _root_.scalapb.GeneratedEnum {
  type EnumType = Method
  def isGet: _root_.scala.Boolean = false
  def isPost: _root_.scala.Boolean = false
  def isPut: _root_.scala.Boolean = false
  def isDelete: _root_.scala.Boolean = false
  def isPatch: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[Method] = api.Method
}

object Method extends _root_.scalapb.GeneratedEnumCompanion[Method] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Method] = this
  @SerialVersionUID(0L)
  case object GET extends Method {
    val value = 0
    val index = 0
    val name = "GET"
    override def isGet: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object POST extends Method {
    val value = 1
    val index = 1
    val name = "POST"
    override def isPost: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object PUT extends Method {
    val value = 2
    val index = 2
    val name = "PUT"
    override def isPut: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object DELETE extends Method {
    val value = 3
    val index = 3
    val name = "DELETE"
    override def isDelete: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object PATCH extends Method {
    val value = 4
    val index = 4
    val name = "PATCH"
    override def isPatch: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends Method with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(GET, POST, PUT, DELETE, PATCH)
  def fromValue(value: _root_.scala.Int): Method = value match {
    case 0 => GET
    case 1 => POST
    case 2 => PUT
    case 3 => DELETE
    case 4 => PATCH
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = ApiProto.javaDescriptor.getEnumTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = ApiProto.scalaDescriptor.enums(0)
}