// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

@SerialVersionUID(0L)
final case class Note(
    message: scala.Option[_root_.scala.Predef.String] = None,
    creator: scala.Option[_root_.scala.Predef.String] = None,
    created: scala.Option[_root_.scala.Long] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Note] with scalapb.lenses.Updatable[Note] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (message.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, message.get) }
      if (creator.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, creator.get) }
      if (created.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(6, created.get) }
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
      message.foreach { __v =>
        _output__.writeString(4, __v)
      };
      creator.foreach { __v =>
        _output__.writeString(5, __v)
      };
      created.foreach { __v =>
        _output__.writeInt64(6, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.Note = {
      var __message = this.message
      var __creator = this.creator
      var __created = this.created
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 34 =>
            __message = Option(_input__.readString())
          case 42 =>
            __creator = Option(_input__.readString())
          case 48 =>
            __created = Option(_input__.readInt64())
          case tag => _input__.skipField(tag)
        }
      }
      payment.Note(
          message = __message,
          creator = __creator,
          created = __created
      )
    }
    def getMessage: _root_.scala.Predef.String = message.getOrElse("")
    def clearMessage: Note = copy(message = None)
    def withMessage(__v: _root_.scala.Predef.String): Note = copy(message = Option(__v))
    def getCreator: _root_.scala.Predef.String = creator.getOrElse("")
    def clearCreator: Note = copy(creator = None)
    def withCreator(__v: _root_.scala.Predef.String): Note = copy(creator = Option(__v))
    def getCreated: _root_.scala.Long = created.getOrElse(0L)
    def clearCreated: Note = copy(created = None)
    def withCreated(__v: _root_.scala.Long): Note = copy(created = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 4 => message.orNull
        case 5 => creator.orNull
        case 6 => created.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 4 => message.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => creator.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => created.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.Note
}

object Note extends scalapb.GeneratedMessageCompanion[payment.Note] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.Note] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.Note = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.Note(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Long]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.Note] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.Note(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Long]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(8)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(8)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.Note(
  )
  implicit class NoteLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.Note]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.Note](_l) {
    def message: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getMessage)((c_, f_) => c_.copy(message = Option(f_)))
    def optionalMessage: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.message)((c_, f_) => c_.copy(message = f_))
    def creator: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getCreator)((c_, f_) => c_.copy(creator = Option(f_)))
    def optionalCreator: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.creator)((c_, f_) => c_.copy(creator = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getCreated)((c_, f_) => c_.copy(created = Option(f_)))
    def optionalCreated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.created)((c_, f_) => c_.copy(created = f_))
  }
  final val MESSAGE_FIELD_NUMBER = 4
  final val CREATOR_FIELD_NUMBER = 5
  final val CREATED_FIELD_NUMBER = 6
}