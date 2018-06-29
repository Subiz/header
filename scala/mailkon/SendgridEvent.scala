// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package mailkon

@SerialVersionUID(0L)
final case class SendgridEvent(
    eventId: scala.Option[_root_.scala.Predef.String] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    conversationId: scala.Option[_root_.scala.Predef.String] = None,
    messageId: scala.Option[_root_.scala.Predef.String] = None,
    data: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    fullMessageId: scala.Option[_root_.scala.Predef.String] = None,
    subject: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[SendgridEvent] with scalapb.lenses.Updatable[SendgridEvent] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (eventId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, eventId.get) }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, accountId.get) }
      if (conversationId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, conversationId.get) }
      if (messageId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, messageId.get) }
      data.foreach(data => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, data))
      if (fullMessageId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, fullMessageId.get) }
      if (subject.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, subject.get) }
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
      eventId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      accountId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      conversationId.foreach { __v =>
        _output__.writeString(6, __v)
      };
      messageId.foreach { __v =>
        _output__.writeString(7, __v)
      };
      data.foreach { __v =>
        _output__.writeString(8, __v)
      };
      fullMessageId.foreach { __v =>
        _output__.writeString(9, __v)
      };
      subject.foreach { __v =>
        _output__.writeString(10, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): mailkon.SendgridEvent = {
      var __eventId = this.eventId
      var __accountId = this.accountId
      var __conversationId = this.conversationId
      var __messageId = this.messageId
      val __data = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.data)
      var __fullMessageId = this.fullMessageId
      var __subject = this.subject
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 26 =>
            __eventId = Option(_input__.readString())
          case 34 =>
            __accountId = Option(_input__.readString())
          case 50 =>
            __conversationId = Option(_input__.readString())
          case 58 =>
            __messageId = Option(_input__.readString())
          case 66 =>
            __data += _input__.readString()
          case 74 =>
            __fullMessageId = Option(_input__.readString())
          case 82 =>
            __subject = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      mailkon.SendgridEvent(
          eventId = __eventId,
          accountId = __accountId,
          conversationId = __conversationId,
          messageId = __messageId,
          data = __data.result(),
          fullMessageId = __fullMessageId,
          subject = __subject
      )
    }
    def getEventId: _root_.scala.Predef.String = eventId.getOrElse("")
    def clearEventId: SendgridEvent = copy(eventId = None)
    def withEventId(__v: _root_.scala.Predef.String): SendgridEvent = copy(eventId = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: SendgridEvent = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): SendgridEvent = copy(accountId = Option(__v))
    def getConversationId: _root_.scala.Predef.String = conversationId.getOrElse("")
    def clearConversationId: SendgridEvent = copy(conversationId = None)
    def withConversationId(__v: _root_.scala.Predef.String): SendgridEvent = copy(conversationId = Option(__v))
    def getMessageId: _root_.scala.Predef.String = messageId.getOrElse("")
    def clearMessageId: SendgridEvent = copy(messageId = None)
    def withMessageId(__v: _root_.scala.Predef.String): SendgridEvent = copy(messageId = Option(__v))
    def clearData = copy(data = _root_.scala.collection.Seq.empty)
    def addData(__vs: _root_.scala.Predef.String*): SendgridEvent = addAllData(__vs)
    def addAllData(__vs: TraversableOnce[_root_.scala.Predef.String]): SendgridEvent = copy(data = data ++ __vs)
    def withData(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): SendgridEvent = copy(data = __v)
    def getFullMessageId: _root_.scala.Predef.String = fullMessageId.getOrElse("")
    def clearFullMessageId: SendgridEvent = copy(fullMessageId = None)
    def withFullMessageId(__v: _root_.scala.Predef.String): SendgridEvent = copy(fullMessageId = Option(__v))
    def getSubject: _root_.scala.Predef.String = subject.getOrElse("")
    def clearSubject: SendgridEvent = copy(subject = None)
    def withSubject(__v: _root_.scala.Predef.String): SendgridEvent = copy(subject = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 3 => eventId.orNull
        case 4 => accountId.orNull
        case 6 => conversationId.orNull
        case 7 => messageId.orNull
        case 8 => data
        case 9 => fullMessageId.orNull
        case 10 => subject.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 3 => eventId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => conversationId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => messageId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => _root_.scalapb.descriptors.PRepeated(data.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 9 => fullMessageId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => subject.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = mailkon.SendgridEvent
}

object SendgridEvent extends scalapb.GeneratedMessageCompanion[mailkon.SendgridEvent] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[mailkon.SendgridEvent] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): mailkon.SendgridEvent = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    mailkon.SendgridEvent(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(4), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[mailkon.SendgridEvent] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      mailkon.SendgridEvent(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = MailkonProto.javaDescriptor.getMessageTypes.get(4)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = MailkonProto.scalaDescriptor.messages(4)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = mailkon.SendgridEvent(
  )
  implicit class SendgridEventLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, mailkon.SendgridEvent]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, mailkon.SendgridEvent](_l) {
    def eventId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getEventId)((c_, f_) => c_.copy(eventId = Option(f_)))
    def optionalEventId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.eventId)((c_, f_) => c_.copy(eventId = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def conversationId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getConversationId)((c_, f_) => c_.copy(conversationId = Option(f_)))
    def optionalConversationId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.conversationId)((c_, f_) => c_.copy(conversationId = f_))
    def messageId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getMessageId)((c_, f_) => c_.copy(messageId = Option(f_)))
    def optionalMessageId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.messageId)((c_, f_) => c_.copy(messageId = f_))
    def data: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.data)((c_, f_) => c_.copy(data = f_))
    def fullMessageId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getFullMessageId)((c_, f_) => c_.copy(fullMessageId = Option(f_)))
    def optionalFullMessageId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.fullMessageId)((c_, f_) => c_.copy(fullMessageId = f_))
    def subject: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getSubject)((c_, f_) => c_.copy(subject = Option(f_)))
    def optionalSubject: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.subject)((c_, f_) => c_.copy(subject = f_))
  }
  final val EVENT_ID_FIELD_NUMBER = 3
  final val ACCOUNT_ID_FIELD_NUMBER = 4
  final val CONVERSATION_ID_FIELD_NUMBER = 6
  final val MESSAGE_ID_FIELD_NUMBER = 7
  final val DATA_FIELD_NUMBER = 8
  final val FULL_MESSAGE_ID_FIELD_NUMBER = 9
  final val SUBJECT_FIELD_NUMBER = 10
}