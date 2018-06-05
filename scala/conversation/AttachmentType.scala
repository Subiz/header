// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

sealed trait AttachmentType extends _root_.scalapb.GeneratedEnum {
  type EnumType = AttachmentType
  def isfile: _root_.scala.Boolean = false
  def isgeneric: _root_.scala.Boolean = false
  def ispreview: _root_.scala.Boolean = false
  def isbutton: _root_.scala.Boolean = false
  def isinput: _root_.scala.Boolean = false
  def isaskinfoform: _root_.scala.Boolean = false
  def isaskinfoformanswer: _root_.scala.Boolean = false
  def isform: _root_.scala.Boolean = false
  def isformsubmit: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[AttachmentType] = conversation.AttachmentType
}

object AttachmentType extends _root_.scalapb.GeneratedEnumCompanion[AttachmentType] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[AttachmentType] = this
  @SerialVersionUID(0L)
  case object file extends AttachmentType {
    val value = 2
    val index = 0
    val name = "file"
    override def isfile: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object generic extends AttachmentType {
    val value = 3
    val index = 1
    val name = "generic"
    override def isgeneric: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object preview extends AttachmentType {
    val value = 4
    val index = 2
    val name = "preview"
    override def ispreview: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object button extends AttachmentType {
    val value = 5
    val index = 3
    val name = "button"
    override def isbutton: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object input extends AttachmentType {
    val value = 6
    val index = 4
    val name = "input"
    override def isinput: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object ask_info_form extends AttachmentType {
    val value = 7
    val index = 5
    val name = "ask_info_form"
    override def isaskinfoform: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object ask_info_form_answer extends AttachmentType {
    val value = 8
    val index = 6
    val name = "ask_info_form_answer"
    override def isaskinfoformanswer: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object form extends AttachmentType {
    val value = 9
    val index = 7
    val name = "form"
    override def isform: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object form_submit extends AttachmentType {
    val value = 10
    val index = 8
    val name = "form_submit"
    override def isformsubmit: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends AttachmentType with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(file, generic, preview, button, input, ask_info_form, ask_info_form_answer, form, form_submit)
  def fromValue(value: _root_.scala.Int): AttachmentType = value match {
    case 2 => file
    case 3 => generic
    case 4 => preview
    case 5 => button
    case 6 => input
    case 7 => ask_info_form
    case 8 => ask_info_form_answer
    case 9 => form
    case 10 => form_submit
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = ConversationProto.javaDescriptor.getEnumTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = ConversationProto.scalaDescriptor.enums(1)
}