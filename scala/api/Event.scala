// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package api

sealed trait Event extends _root_.scalapb.GeneratedEnum {
  type EnumType = Event
  def isApiRequested: _root_.scala.Boolean = false
  def isConversationStart: _root_.scala.Boolean = false
  def isApiReply: _root_.scala.Boolean = false
  def isApiWhiteListSynced: _root_.scala.Boolean = false
  def isDomainWhiteListCreated: _root_.scala.Boolean = false
  def isDomainWhiteListDeleted: _root_.scala.Boolean = false
  def isUserWhiteListCreated: _root_.scala.Boolean = false
  def isUserWhiteListDeleted: _root_.scala.Boolean = false
  def isIpWhiteListCreated: _root_.scala.Boolean = false
  def isIpWhiteListDeleted: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[Event] = api.Event
}

object Event extends _root_.scalapb.GeneratedEnumCompanion[Event] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Event] = this
  @SerialVersionUID(0L)
  case object ApiRequested extends Event {
    val value = 0
    val index = 0
    val name = "ApiRequested"
    override def isApiRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object ConversationStart extends Event {
    val value = 2
    val index = 1
    val name = "ConversationStart"
    override def isConversationStart: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object ApiReply extends Event {
    val value = 4
    val index = 2
    val name = "ApiReply"
    override def isApiReply: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object ApiWhiteListSynced extends Event {
    val value = 5
    val index = 3
    val name = "ApiWhiteListSynced"
    override def isApiWhiteListSynced: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object DomainWhiteListCreated extends Event {
    val value = 6
    val index = 4
    val name = "DomainWhiteListCreated"
    override def isDomainWhiteListCreated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object DomainWhiteListDeleted extends Event {
    val value = 7
    val index = 5
    val name = "DomainWhiteListDeleted"
    override def isDomainWhiteListDeleted: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object UserWhiteListCreated extends Event {
    val value = 8
    val index = 6
    val name = "UserWhiteListCreated"
    override def isUserWhiteListCreated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object UserWhiteListDeleted extends Event {
    val value = 9
    val index = 7
    val name = "UserWhiteListDeleted"
    override def isUserWhiteListDeleted: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object IpWhiteListCreated extends Event {
    val value = 10
    val index = 8
    val name = "IpWhiteListCreated"
    override def isIpWhiteListCreated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object IpWhiteListDeleted extends Event {
    val value = 11
    val index = 9
    val name = "IpWhiteListDeleted"
    override def isIpWhiteListDeleted: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends Event with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(ApiRequested, ConversationStart, ApiReply, ApiWhiteListSynced, DomainWhiteListCreated, DomainWhiteListDeleted, UserWhiteListCreated, UserWhiteListDeleted, IpWhiteListCreated, IpWhiteListDeleted)
  def fromValue(value: _root_.scala.Int): Event = value match {
    case 0 => ApiRequested
    case 2 => ConversationStart
    case 4 => ApiReply
    case 5 => ApiWhiteListSynced
    case 6 => DomainWhiteListCreated
    case 7 => DomainWhiteListDeleted
    case 8 => UserWhiteListCreated
    case 9 => UserWhiteListDeleted
    case 10 => IpWhiteListCreated
    case 11 => IpWhiteListDeleted
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = ApiProto.javaDescriptor.getEnumTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = ApiProto.scalaDescriptor.enums(1)
}