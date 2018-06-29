// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

@SerialVersionUID(0L)
final case class Ack(
    memberId: scala.Option[_root_.scala.Predef.String] = None,
    error: scala.Option[_root_.scala.Predef.String] = None,
    at: scala.Option[_root_.scala.Long] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Ack] with scalapb.lenses.Updatable[Ack] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (memberId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, memberId.get) }
      if (error.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, error.get) }
      if (at.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(4, at.get) }
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
      at.foreach { __v =>
        _output__.writeInt64(4, __v)
      };
      memberId.foreach { __v =>
        _output__.writeString(5, __v)
      };
      error.foreach { __v =>
        _output__.writeString(6, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.Ack = {
      var __memberId = this.memberId
      var __error = this.error
      var __at = this.at
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 42 =>
            __memberId = Option(_input__.readString())
          case 50 =>
            __error = Option(_input__.readString())
          case 32 =>
            __at = Option(_input__.readInt64())
          case tag => _input__.skipField(tag)
        }
      }
      conversation.Ack(
          memberId = __memberId,
          error = __error,
          at = __at
      )
    }
    def getMemberId: _root_.scala.Predef.String = memberId.getOrElse("")
    def clearMemberId: Ack = copy(memberId = None)
    def withMemberId(__v: _root_.scala.Predef.String): Ack = copy(memberId = Option(__v))
    def getError: _root_.scala.Predef.String = error.getOrElse("")
    def clearError: Ack = copy(error = None)
    def withError(__v: _root_.scala.Predef.String): Ack = copy(error = Option(__v))
    def getAt: _root_.scala.Long = at.getOrElse(0L)
    def clearAt: Ack = copy(at = None)
    def withAt(__v: _root_.scala.Long): Ack = copy(at = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 5 => memberId.orNull
        case 6 => error.orNull
        case 4 => at.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 5 => memberId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => error.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => at.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.Ack
}

object Ack extends scalapb.GeneratedMessageCompanion[conversation.Ack] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.Ack] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.Ack = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.Ack(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Long]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.Ack] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.Ack(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Long]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(23)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(23)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.Ack(
  )
  implicit class AckLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.Ack]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.Ack](_l) {
    def memberId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getMemberId)((c_, f_) => c_.copy(memberId = Option(f_)))
    def optionalMemberId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.memberId)((c_, f_) => c_.copy(memberId = f_))
    def error: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getError)((c_, f_) => c_.copy(error = Option(f_)))
    def optionalError: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.error)((c_, f_) => c_.copy(error = f_))
    def at: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getAt)((c_, f_) => c_.copy(at = Option(f_)))
    def optionalAt: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.at)((c_, f_) => c_.copy(at = f_))
  }
  final val MEMBER_ID_FIELD_NUMBER = 5
  final val ERROR_FIELD_NUMBER = 6
  final val AT_FIELD_NUMBER = 4
}