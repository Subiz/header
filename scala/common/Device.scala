// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package common

/** @param userAgent
  *   code
  */
@SerialVersionUID(0L)
final case class Device(
    ip: _root_.scala.Predef.String = "",
    userAgent: _root_.scala.Predef.String = "",
    screenResolution: _root_.scala.Predef.String = "",
    timezone: _root_.scala.Predef.String = "",
    language: _root_.scala.Predef.String = "",
    referrer: _root_.scala.Predef.String = "",
    `type`: _root_.scala.Predef.String = ""
    ) extends scalapb.GeneratedMessage with scalapb.Message[Device] with scalapb.lenses.Updatable[Device] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ip != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, ip) }
      if (userAgent != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, userAgent) }
      if (screenResolution != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, screenResolution) }
      if (timezone != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, timezone) }
      if (language != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, language) }
      if (referrer != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, referrer) }
      if (`type` != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, `type`) }
      __size
    }
    final override def serializedSize: _root_.scala.Int = {
      var read = __serializedSizeCachedValue
      if (read == 0) {
        read = __computeSerializedValue()
        __serializedSizeCachedValue = read
      }
      read
    }
    def writeTo(`_output__`: _root_.com.google.protobuf.CodedOutputStream): _root_.scala.Unit = {
      {
        val __v = ip
        if (__v != "") {
          _output__.writeString(3, __v)
        }
      };
      {
        val __v = userAgent
        if (__v != "") {
          _output__.writeString(4, __v)
        }
      };
      {
        val __v = screenResolution
        if (__v != "") {
          _output__.writeString(5, __v)
        }
      };
      {
        val __v = timezone
        if (__v != "") {
          _output__.writeString(6, __v)
        }
      };
      {
        val __v = language
        if (__v != "") {
          _output__.writeString(7, __v)
        }
      };
      {
        val __v = referrer
        if (__v != "") {
          _output__.writeString(8, __v)
        }
      };
      {
        val __v = `type`
        if (__v != "") {
          _output__.writeString(9, __v)
        }
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): common.Device = {
      var __ip = this.ip
      var __userAgent = this.userAgent
      var __screenResolution = this.screenResolution
      var __timezone = this.timezone
      var __language = this.language
      var __referrer = this.referrer
      var __type = this.`type`
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 26 =>
            __ip = _input__.readString()
          case 34 =>
            __userAgent = _input__.readString()
          case 42 =>
            __screenResolution = _input__.readString()
          case 50 =>
            __timezone = _input__.readString()
          case 58 =>
            __language = _input__.readString()
          case 66 =>
            __referrer = _input__.readString()
          case 74 =>
            __type = _input__.readString()
          case tag => _input__.skipField(tag)
        }
      }
      common.Device(
          ip = __ip,
          userAgent = __userAgent,
          screenResolution = __screenResolution,
          timezone = __timezone,
          language = __language,
          referrer = __referrer,
          `type` = __type
      )
    }
    def withIp(__v: _root_.scala.Predef.String): Device = copy(ip = __v)
    def withUserAgent(__v: _root_.scala.Predef.String): Device = copy(userAgent = __v)
    def withScreenResolution(__v: _root_.scala.Predef.String): Device = copy(screenResolution = __v)
    def withTimezone(__v: _root_.scala.Predef.String): Device = copy(timezone = __v)
    def withLanguage(__v: _root_.scala.Predef.String): Device = copy(language = __v)
    def withReferrer(__v: _root_.scala.Predef.String): Device = copy(referrer = __v)
    def withType(__v: _root_.scala.Predef.String): Device = copy(`type` = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 3 => {
          val __t = ip
          if (__t != "") __t else null
        }
        case 4 => {
          val __t = userAgent
          if (__t != "") __t else null
        }
        case 5 => {
          val __t = screenResolution
          if (__t != "") __t else null
        }
        case 6 => {
          val __t = timezone
          if (__t != "") __t else null
        }
        case 7 => {
          val __t = language
          if (__t != "") __t else null
        }
        case 8 => {
          val __t = referrer
          if (__t != "") __t else null
        }
        case 9 => {
          val __t = `type`
          if (__t != "") __t else null
        }
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 3 => _root_.scalapb.descriptors.PString(ip)
        case 4 => _root_.scalapb.descriptors.PString(userAgent)
        case 5 => _root_.scalapb.descriptors.PString(screenResolution)
        case 6 => _root_.scalapb.descriptors.PString(timezone)
        case 7 => _root_.scalapb.descriptors.PString(language)
        case 8 => _root_.scalapb.descriptors.PString(referrer)
        case 9 => _root_.scalapb.descriptors.PString(`type`)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = common.Device
}

object Device extends scalapb.GeneratedMessageCompanion[common.Device] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[common.Device] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): common.Device = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    common.Device(
      __fieldsMap.getOrElse(__fields.get(0), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(1), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(2), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(3), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(4), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(5), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(6), "").asInstanceOf[_root_.scala.Predef.String]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[common.Device] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      common.Device(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).map(_.as[_root_.scala.Predef.String]).getOrElse("")
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = CommonProto.javaDescriptor.getMessageTypes.get(6)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = CommonProto.scalaDescriptor.messages(6)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = common.Device(
  )
  sealed trait DeviceType extends _root_.scalapb.GeneratedEnum {
    type EnumType = DeviceType
    def isunknown: _root_.scala.Boolean = false
    def ismobile: _root_.scala.Boolean = false
    def istablet: _root_.scala.Boolean = false
    def isdesktop: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[DeviceType] = common.Device.DeviceType
  }
  
  object DeviceType extends _root_.scalapb.GeneratedEnumCompanion[DeviceType] {
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[DeviceType] = this
    @SerialVersionUID(0L)
    case object unknown extends DeviceType {
      val value = 0
      val index = 0
      val name = "unknown"
      override def isunknown: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object mobile extends DeviceType {
      val value = 1
      val index = 1
      val name = "mobile"
      override def ismobile: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object tablet extends DeviceType {
      val value = 2
      val index = 2
      val name = "tablet"
      override def istablet: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object desktop extends DeviceType {
      val value = 3
      val index = 3
      val name = "desktop"
      override def isdesktop: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(value: _root_.scala.Int) extends DeviceType with _root_.scalapb.UnrecognizedEnum
    
    lazy val values = scala.collection.Seq(unknown, mobile, tablet, desktop)
    def fromValue(value: _root_.scala.Int): DeviceType = value match {
      case 0 => unknown
      case 1 => mobile
      case 2 => tablet
      case 3 => desktop
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = common.Device.javaDescriptor.getEnumTypes.get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = common.Device.scalaDescriptor.enums(0)
  }
  implicit class DeviceLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, common.Device]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, common.Device](_l) {
    def ip: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.ip)((c_, f_) => c_.copy(ip = f_))
    def userAgent: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.userAgent)((c_, f_) => c_.copy(userAgent = f_))
    def screenResolution: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.screenResolution)((c_, f_) => c_.copy(screenResolution = f_))
    def timezone: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.timezone)((c_, f_) => c_.copy(timezone = f_))
    def language: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.language)((c_, f_) => c_.copy(language = f_))
    def referrer: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.referrer)((c_, f_) => c_.copy(referrer = f_))
    def `type`: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.`type`)((c_, f_) => c_.copy(`type` = f_))
  }
  final val IP_FIELD_NUMBER = 3
  final val USER_AGENT_FIELD_NUMBER = 4
  final val SCREEN_RESOLUTION_FIELD_NUMBER = 5
  final val TIMEZONE_FIELD_NUMBER = 6
  final val LANGUAGE_FIELD_NUMBER = 7
  final val REFERRER_FIELD_NUMBER = 8
  final val TYPE_FIELD_NUMBER = 9
}
