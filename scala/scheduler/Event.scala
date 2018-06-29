// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package scheduler

sealed trait Event extends _root_.scalapb.GeneratedEnum {
  type EnumType = Event
  def isSchedulerRequested: _root_.scala.Boolean = false
  def isSchedulerCancelled: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[Event] = scheduler.Event
}

object Event extends _root_.scalapb.GeneratedEnumCompanion[Event] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Event] = this
  @SerialVersionUID(0L)
  case object SchedulerRequested extends Event {
    val value = 0
    val index = 0
    val name = "SchedulerRequested"
    override def isSchedulerRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object SchedulerCancelled extends Event {
    val value = 1
    val index = 1
    val name = "SchedulerCancelled"
    override def isSchedulerCancelled: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends Event with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(SchedulerRequested, SchedulerCancelled)
  def fromValue(value: _root_.scala.Int): Event = value match {
    case 0 => SchedulerRequested
    case 1 => SchedulerCancelled
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = SchedulerProto.javaDescriptor.getEnumTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = SchedulerProto.scalaDescriptor.enums(0)
}