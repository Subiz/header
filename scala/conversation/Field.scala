// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

@SerialVersionUID(0L)
final case class Field(
    value: scala.Option[_root_.scala.Predef.String] = None,
    key: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Field] with scalapb.lenses.Updatable[Field] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (value.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, value.get) }
      if (key.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, key.get) }
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
      key.foreach { __v =>
        _output__.writeString(2, __v)
      };
      value.foreach { __v =>
        _output__.writeString(3, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.Field = {
      var __value = this.value
      var __key = this.key
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 26 =>
            __value = Option(_input__.readString())
          case 18 =>
            __key = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      conversation.Field(
          value = __value,
          key = __key
      )
    }
    def getValue: _root_.scala.Predef.String = value.getOrElse("")
    def clearValue: Field = copy(value = None)
    def withValue(__v: _root_.scala.Predef.String): Field = copy(value = Option(__v))
    def getKey: _root_.scala.Predef.String = key.getOrElse("")
    def clearKey: Field = copy(key = None)
    def withKey(__v: _root_.scala.Predef.String): Field = copy(key = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 3 => value.orNull
        case 2 => key.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 3 => value.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => key.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.Field
}

object Field extends scalapb.GeneratedMessageCompanion[conversation.Field] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.Field] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.Field = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.Field(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.Field] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.Field(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(25)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(25)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.Field(
  )
  implicit class FieldLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.Field]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.Field](_l) {
    def value: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getValue)((c_, f_) => c_.copy(value = Option(f_)))
    def optionalValue: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.value)((c_, f_) => c_.copy(value = f_))
    def key: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getKey)((c_, f_) => c_.copy(key = Option(f_)))
    def optionalKey: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.key)((c_, f_) => c_.copy(key = f_))
  }
  final val VALUE_FIELD_NUMBER = 3
  final val KEY_FIELD_NUMBER = 2
}