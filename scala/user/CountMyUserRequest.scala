// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

@SerialVersionUID(0L)
final case class CountMyUserRequest(
    agentId: scala.Option[_root_.scala.Predef.String] = None,
    from: scala.Option[_root_.scala.Int] = None,
    to: scala.Option[_root_.scala.Int] = None,
    range: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[CountMyUserRequest] with scalapb.lenses.Updatable[CountMyUserRequest] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (agentId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, agentId.get) }
      if (from.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(5, from.get) }
      if (to.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(6, to.get) }
      if (range.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, range.get) }
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
      agentId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      from.foreach { __v =>
        _output__.writeInt32(5, __v)
      };
      to.foreach { __v =>
        _output__.writeInt32(6, __v)
      };
      range.foreach { __v =>
        _output__.writeString(8, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): user.CountMyUserRequest = {
      var __agentId = this.agentId
      var __from = this.from
      var __to = this.to
      var __range = this.range
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 34 =>
            __agentId = Option(_input__.readString())
          case 40 =>
            __from = Option(_input__.readInt32())
          case 48 =>
            __to = Option(_input__.readInt32())
          case 66 =>
            __range = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      user.CountMyUserRequest(
          agentId = __agentId,
          from = __from,
          to = __to,
          range = __range
      )
    }
    def getAgentId: _root_.scala.Predef.String = agentId.getOrElse("")
    def clearAgentId: CountMyUserRequest = copy(agentId = None)
    def withAgentId(__v: _root_.scala.Predef.String): CountMyUserRequest = copy(agentId = Option(__v))
    def getFrom: _root_.scala.Int = from.getOrElse(0)
    def clearFrom: CountMyUserRequest = copy(from = None)
    def withFrom(__v: _root_.scala.Int): CountMyUserRequest = copy(from = Option(__v))
    def getTo: _root_.scala.Int = to.getOrElse(0)
    def clearTo: CountMyUserRequest = copy(to = None)
    def withTo(__v: _root_.scala.Int): CountMyUserRequest = copy(to = Option(__v))
    def getRange: _root_.scala.Predef.String = range.getOrElse("")
    def clearRange: CountMyUserRequest = copy(range = None)
    def withRange(__v: _root_.scala.Predef.String): CountMyUserRequest = copy(range = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 4 => agentId.orNull
        case 5 => from.orNull
        case 6 => to.orNull
        case 8 => range.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 4 => agentId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => from.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => to.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => range.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = user.CountMyUserRequest
}

object CountMyUserRequest extends scalapb.GeneratedMessageCompanion[user.CountMyUserRequest] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[user.CountMyUserRequest] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): user.CountMyUserRequest = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    user.CountMyUserRequest(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Int]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Int]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[user.CountMyUserRequest] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      user.CountMyUserRequest(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Int]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Int]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = UserProto.javaDescriptor.getMessageTypes.get(44)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = UserProto.scalaDescriptor.messages(44)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = user.CountMyUserRequest(
  )
  sealed trait Range extends _root_.scalapb.GeneratedEnum {
    type EnumType = Range
    def ishour: _root_.scala.Boolean = false
    def isday: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[Range] = user.CountMyUserRequest.Range
  }
  
  object Range extends _root_.scalapb.GeneratedEnumCompanion[Range] {
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[Range] = this
    @SerialVersionUID(0L)
    case object hour extends Range {
      val value = 0
      val index = 0
      val name = "hour"
      override def ishour: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object day extends Range {
      val value = 1
      val index = 1
      val name = "day"
      override def isday: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(value: _root_.scala.Int) extends Range with _root_.scalapb.UnrecognizedEnum
    
    lazy val values = scala.collection.Seq(hour, day)
    def fromValue(value: _root_.scala.Int): Range = value match {
      case 0 => hour
      case 1 => day
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = user.CountMyUserRequest.javaDescriptor.getEnumTypes.get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = user.CountMyUserRequest.scalaDescriptor.enums(0)
  }
  implicit class CountMyUserRequestLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, user.CountMyUserRequest]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, user.CountMyUserRequest](_l) {
    def agentId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentId)((c_, f_) => c_.copy(agentId = Option(f_)))
    def optionalAgentId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentId)((c_, f_) => c_.copy(agentId = f_))
    def from: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getFrom)((c_, f_) => c_.copy(from = Option(f_)))
    def optionalFrom: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.from)((c_, f_) => c_.copy(from = f_))
    def to: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getTo)((c_, f_) => c_.copy(to = Option(f_)))
    def optionalTo: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.to)((c_, f_) => c_.copy(to = f_))
    def range: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getRange)((c_, f_) => c_.copy(range = Option(f_)))
    def optionalRange: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.range)((c_, f_) => c_.copy(range = f_))
  }
  final val AGENT_ID_FIELD_NUMBER = 4
  final val FROM_FIELD_NUMBER = 5
  final val TO_FIELD_NUMBER = 6
  final val RANGE_FIELD_NUMBER = 8
}
