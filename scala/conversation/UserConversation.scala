// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

/** used in database
  */
@SerialVersionUID(0L)
final case class UserConversation(
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    state: scala.Option[_root_.scala.Predef.String] = None,
    userId: scala.Option[_root_.scala.Predef.String] = None,
    convoId: scala.Option[_root_.scala.Predef.String] = None,
    lastSeenEvent: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[UserConversation] with scalapb.lenses.Updatable[UserConversation] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, accountId.get) }
      if (state.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, state.get) }
      if (userId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, userId.get) }
      if (convoId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, convoId.get) }
      if (lastSeenEvent.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, lastSeenEvent.get) }
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
      userId.foreach { __v =>
        _output__.writeString(2, __v)
      };
      convoId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      lastSeenEvent.foreach { __v =>
        _output__.writeString(5, __v)
      };
      accountId.foreach { __v =>
        _output__.writeString(6, __v)
      };
      state.foreach { __v =>
        _output__.writeString(11, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.UserConversation = {
      var __accountId = this.accountId
      var __state = this.state
      var __userId = this.userId
      var __convoId = this.convoId
      var __lastSeenEvent = this.lastSeenEvent
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 50 =>
            __accountId = Option(_input__.readString())
          case 90 =>
            __state = Option(_input__.readString())
          case 18 =>
            __userId = Option(_input__.readString())
          case 34 =>
            __convoId = Option(_input__.readString())
          case 42 =>
            __lastSeenEvent = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      conversation.UserConversation(
          accountId = __accountId,
          state = __state,
          userId = __userId,
          convoId = __convoId,
          lastSeenEvent = __lastSeenEvent
      )
    }
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: UserConversation = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): UserConversation = copy(accountId = Option(__v))
    def getState: _root_.scala.Predef.String = state.getOrElse("")
    def clearState: UserConversation = copy(state = None)
    def withState(__v: _root_.scala.Predef.String): UserConversation = copy(state = Option(__v))
    def getUserId: _root_.scala.Predef.String = userId.getOrElse("")
    def clearUserId: UserConversation = copy(userId = None)
    def withUserId(__v: _root_.scala.Predef.String): UserConversation = copy(userId = Option(__v))
    def getConvoId: _root_.scala.Predef.String = convoId.getOrElse("")
    def clearConvoId: UserConversation = copy(convoId = None)
    def withConvoId(__v: _root_.scala.Predef.String): UserConversation = copy(convoId = Option(__v))
    def getLastSeenEvent: _root_.scala.Predef.String = lastSeenEvent.getOrElse("")
    def clearLastSeenEvent: UserConversation = copy(lastSeenEvent = None)
    def withLastSeenEvent(__v: _root_.scala.Predef.String): UserConversation = copy(lastSeenEvent = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 6 => accountId.orNull
        case 11 => state.orNull
        case 2 => userId.orNull
        case 4 => convoId.orNull
        case 5 => lastSeenEvent.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 6 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 11 => state.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => userId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => convoId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => lastSeenEvent.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.UserConversation
}

object UserConversation extends scalapb.GeneratedMessageCompanion[conversation.UserConversation] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.UserConversation] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.UserConversation = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.UserConversation(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.UserConversation] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.UserConversation(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(10)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(10)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.UserConversation(
  )
  implicit class UserConversationLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.UserConversation]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.UserConversation](_l) {
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def state: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getState)((c_, f_) => c_.copy(state = Option(f_)))
    def optionalState: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.state)((c_, f_) => c_.copy(state = f_))
    def userId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getUserId)((c_, f_) => c_.copy(userId = Option(f_)))
    def optionalUserId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.userId)((c_, f_) => c_.copy(userId = f_))
    def convoId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getConvoId)((c_, f_) => c_.copy(convoId = Option(f_)))
    def optionalConvoId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.convoId)((c_, f_) => c_.copy(convoId = f_))
    def lastSeenEvent: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getLastSeenEvent)((c_, f_) => c_.copy(lastSeenEvent = Option(f_)))
    def optionalLastSeenEvent: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.lastSeenEvent)((c_, f_) => c_.copy(lastSeenEvent = f_))
  }
  final val ACCOUNT_ID_FIELD_NUMBER = 6
  final val STATE_FIELD_NUMBER = 11
  final val USER_ID_FIELD_NUMBER = 2
  final val CONVO_ID_FIELD_NUMBER = 4
  final val LAST_SEEN_EVENT_FIELD_NUMBER = 5
}
