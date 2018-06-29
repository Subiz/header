// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package widget

sealed trait Event extends _root_.scalapb.GeneratedEnum {
  type EnumType = Event
  def isWidgetUserSettingReadRequested: _root_.scala.Boolean = false
  def isWidgetUserSettingUpdateRequested: _root_.scala.Boolean = false
  def isWidgetSettingReadRequested: _root_.scala.Boolean = false
  def isWidgetSettingUpdateRequested: _root_.scala.Boolean = false
  def isWidgetSettingCssVersionUpdated: _root_.scala.Boolean = false
  def isWidgetSettingLanguageUpdated: _root_.scala.Boolean = false
  def isWidgetSettingCssVersionUpdateRequested: _root_.scala.Boolean = false
  def isWidgetSettingLanguageUpdateRequested: _root_.scala.Boolean = false
  def isWidgetRequested: _root_.scala.Boolean = false
  def isWidgetSettingUpserted: _root_.scala.Boolean = false
  def isWidgetSynced: _root_.scala.Boolean = false
  def isWidgetV3Synced: _root_.scala.Boolean = false
  def companion: _root_.scalapb.GeneratedEnumCompanion[Event] = widget.Event
}

object Event extends _root_.scalapb.GeneratedEnumCompanion[Event] {
  implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Event] = this
  @SerialVersionUID(0L)
  case object WidgetUserSettingReadRequested extends Event {
    val value = 0
    val index = 0
    val name = "WidgetUserSettingReadRequested"
    override def isWidgetUserSettingReadRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetUserSettingUpdateRequested extends Event {
    val value = 1
    val index = 1
    val name = "WidgetUserSettingUpdateRequested"
    override def isWidgetUserSettingUpdateRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingReadRequested extends Event {
    val value = 2
    val index = 2
    val name = "WidgetSettingReadRequested"
    override def isWidgetSettingReadRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingUpdateRequested extends Event {
    val value = 3
    val index = 3
    val name = "WidgetSettingUpdateRequested"
    override def isWidgetSettingUpdateRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingCssVersionUpdated extends Event {
    val value = 4
    val index = 4
    val name = "WidgetSettingCssVersionUpdated"
    override def isWidgetSettingCssVersionUpdated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingLanguageUpdated extends Event {
    val value = 5
    val index = 5
    val name = "WidgetSettingLanguageUpdated"
    override def isWidgetSettingLanguageUpdated: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingCssVersionUpdateRequested extends Event {
    val value = 6
    val index = 6
    val name = "WidgetSettingCssVersionUpdateRequested"
    override def isWidgetSettingCssVersionUpdateRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingLanguageUpdateRequested extends Event {
    val value = 7
    val index = 7
    val name = "WidgetSettingLanguageUpdateRequested"
    override def isWidgetSettingLanguageUpdateRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetRequested extends Event {
    val value = 8
    val index = 8
    val name = "WidgetRequested"
    override def isWidgetRequested: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSettingUpserted extends Event {
    val value = 10
    val index = 9
    val name = "WidgetSettingUpserted"
    override def isWidgetSettingUpserted: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetSynced extends Event {
    val value = 100
    val index = 10
    val name = "WidgetSynced"
    override def isWidgetSynced: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  case object WidgetV3Synced extends Event {
    val value = 103
    val index = 11
    val name = "WidgetV3Synced"
    override def isWidgetV3Synced: _root_.scala.Boolean = true
  }
  
  @SerialVersionUID(0L)
  final case class Unrecognized(value: _root_.scala.Int) extends Event with _root_.scalapb.UnrecognizedEnum
  
  lazy val values = scala.collection.Seq(WidgetUserSettingReadRequested, WidgetUserSettingUpdateRequested, WidgetSettingReadRequested, WidgetSettingUpdateRequested, WidgetSettingCssVersionUpdated, WidgetSettingLanguageUpdated, WidgetSettingCssVersionUpdateRequested, WidgetSettingLanguageUpdateRequested, WidgetRequested, WidgetSettingUpserted, WidgetSynced, WidgetV3Synced)
  def fromValue(value: _root_.scala.Int): Event = value match {
    case 0 => WidgetUserSettingReadRequested
    case 1 => WidgetUserSettingUpdateRequested
    case 2 => WidgetSettingReadRequested
    case 3 => WidgetSettingUpdateRequested
    case 4 => WidgetSettingCssVersionUpdated
    case 5 => WidgetSettingLanguageUpdated
    case 6 => WidgetSettingCssVersionUpdateRequested
    case 7 => WidgetSettingLanguageUpdateRequested
    case 8 => WidgetRequested
    case 10 => WidgetSettingUpserted
    case 100 => WidgetSynced
    case 103 => WidgetV3Synced
    case __other => Unrecognized(__other)
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = WidgetProto.javaDescriptor.getEnumTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = WidgetProto.scalaDescriptor.enums(0)
}