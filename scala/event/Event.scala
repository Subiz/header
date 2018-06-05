// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package event

sealed trait Event extends _root_.scalapb.GeneratedEnum {
  type EnumType = Event
  def isSub: _root_.scala.Boolean = false
  def isRawEventCreated: _root_.scala.Boolean = false
  def isSubscribe: _root_.scala.Boolean = false
  def isSubscribeReply: _root_.scala.Boolean = false
  def isUnsubscribe: _root_.scala.Boolean = false
  def isUnsubscribeReply: _root_.scala.Boolean = false
  def isEventSync: _root_.scala.Boolean = false
  def isEventCreated: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[Event] = event.Event
}

object Event extends _root_.scalapb.GeneratedEnumCompanion[Event] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Event] = this
  @SerialVersionUID(0L)
  case object Sub extends Event {
    val value = 0
    val index = 0
    val name = "Sub"
    override def isSub: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object RawEventCreated extends Event {
    val value = 3
    val index = 1
    val name = "RawEventCreated"
    override def isRawEventCreated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object Subscribe extends Event {
    val value = 4
    val index = 2
    val name = "Subscribe"
    override def isSubscribe: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object SubscribeReply extends Event {
    val value = 6
    val index = 3
    val name = "SubscribeReply"
    override def isSubscribeReply: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object Unsubscribe extends Event {
    val value = 5
    val index = 4
    val name = "Unsubscribe"
    override def isUnsubscribe: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object UnsubscribeReply extends Event {
    val value = 7
    val index = 5
    val name = "UnsubscribeReply"
    override def isUnsubscribeReply: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object EventSync extends Event {
    val value = 8
    val index = 6
    val name = "EventSync"
    override def isEventSync: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object EventCreated extends Event {
    val value = 9
    val index = 7
    val name = "EventCreated"
    override def isEventCreated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends Event with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(Sub, RawEventCreated, Subscribe, SubscribeReply, Unsubscribe, UnsubscribeReply, EventSync, EventCreated)
  def fromValue(value: _root_.scala.Int): Event = value match {
    case 0 => Sub
    case 3 => RawEventCreated
    case 4 => Subscribe
    case 5 => Unsubscribe
    case 6 => SubscribeReply
    case 7 => UnsubscribeReply
    case 8 => EventSync
    case 9 => EventCreated
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = EventProto.javaDescriptor.getEnumTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = EventProto.scalaDescriptor.enums(0)
}