// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package lang

sealed trait L extends _root_.scalapb.GeneratedEnum {
  type EnumType = L
  def isen: _root_.scala.Boolean = false
  def isvn: _root_.scala.Boolean = false
  def iscn: _root_.scala.Boolean = false
  def isfr: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[L] = lang.L
}

object L extends _root_.scalapb.GeneratedEnumCompanion[L] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[L] = this
  @SerialVersionUID(0L)
  case object en extends L {
    val value = 0
    val index = 0
    val name = "en"
    override def isen: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object vn extends L {
    val value = 1
    val index = 1
    val name = "vn"
    override def isvn: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object cn extends L {
    val value = 3
    val index = 2
    val name = "cn"
    override def iscn: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object fr extends L {
    val value = 5
    val index = 3
    val name = "fr"
    override def isfr: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends L with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(en, vn, cn, fr)
  def fromValue(value: _root_.scala.Int): L = value match {
    case 0 => en
    case 1 => vn
    case 3 => cn
    case 5 => fr
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = LangProto.javaDescriptor.getEnumTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = LangProto.scalaDescriptor.enums(1)
}