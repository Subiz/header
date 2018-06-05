// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package api

@SerialVersionUID(0L)
final case class SendMessageState(
    topic: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[SendMessageState] with scalapb.lenses.Updatable[SendMessageState] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (topic.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, topic.get) }
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
      topic.foreach { __v =>
        _output__.writeString(1, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): api.SendMessageState = {
      var __topic = this.topic
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __topic = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      api.SendMessageState(
          topic = __topic
      )
    }
    def getTopic: _root_.scala.Predef.String = topic.getOrElse("")
    def clearTopic: SendMessageState = copy(topic = None)
    def withTopic(__v: _root_.scala.Predef.String): SendMessageState = copy(topic = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => topic.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => topic.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = api.SendMessageState
}

object SendMessageState extends scalapb.GeneratedMessageCompanion[api.SendMessageState] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[api.SendMessageState] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): api.SendMessageState = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    api.SendMessageState(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[api.SendMessageState] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      api.SendMessageState(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ApiProto.javaDescriptor.getMessageTypes.get(2)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ApiProto.scalaDescriptor.messages(2)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = api.SendMessageState(
  )
  implicit class SendMessageStateLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, api.SendMessageState]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, api.SendMessageState](_l) {
    def topic: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getTopic)((c_, f_) => c_.copy(topic = Option(f_)))
    def optionalTopic: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.topic)((c_, f_) => c_.copy(topic = f_))
  }
  final val TOPIC_FIELD_NUMBER = 1
}
