// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package account

@SerialVersionUID(0L)
final case class PasswordChangedEmail(
    ctx: scala.Option[common.Context] = None,
    to: scala.Option[_root_.scala.Predef.String] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    agentId: scala.Option[_root_.scala.Predef.String] = None,
    agentName: scala.Option[_root_.scala.Predef.String] = None,
    lang: scala.Option[_root_.lang.L] = None,
    from: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[PasswordChangedEmail] with scalapb.lenses.Updatable[PasswordChangedEmail] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (to.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, to.get) }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, accountId.get) }
      if (agentId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, agentId.get) }
      if (agentName.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, agentName.get) }
      if (lang.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeEnumSize(7, lang.get.value) }
      if (from.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, from.get) }
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
      ctx.foreach { __v =>
        _output__.writeTag(1, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      to.foreach { __v =>
        _output__.writeString(2, __v)
      };
      accountId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      agentId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      agentName.foreach { __v =>
        _output__.writeString(5, __v)
      };
      lang.foreach { __v =>
        _output__.writeEnum(7, __v.value)
      };
      from.foreach { __v =>
        _output__.writeString(9, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): account.PasswordChangedEmail = {
      var __ctx = this.ctx
      var __to = this.to
      var __accountId = this.accountId
      var __agentId = this.agentId
      var __agentName = this.agentName
      var __lang = this.lang
      var __from = this.from
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __to = Option(_input__.readString())
          case 26 =>
            __accountId = Option(_input__.readString())
          case 34 =>
            __agentId = Option(_input__.readString())
          case 42 =>
            __agentName = Option(_input__.readString())
          case 56 =>
            __lang = Option(_root_.lang.L.fromValue(_input__.readEnum()))
          case 74 =>
            __from = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      account.PasswordChangedEmail(
          ctx = __ctx,
          to = __to,
          accountId = __accountId,
          agentId = __agentId,
          agentName = __agentName,
          lang = __lang,
          from = __from
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: PasswordChangedEmail = copy(ctx = None)
    def withCtx(__v: common.Context): PasswordChangedEmail = copy(ctx = Option(__v))
    def getTo: _root_.scala.Predef.String = to.getOrElse("")
    def clearTo: PasswordChangedEmail = copy(to = None)
    def withTo(__v: _root_.scala.Predef.String): PasswordChangedEmail = copy(to = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: PasswordChangedEmail = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): PasswordChangedEmail = copy(accountId = Option(__v))
    def getAgentId: _root_.scala.Predef.String = agentId.getOrElse("")
    def clearAgentId: PasswordChangedEmail = copy(agentId = None)
    def withAgentId(__v: _root_.scala.Predef.String): PasswordChangedEmail = copy(agentId = Option(__v))
    def getAgentName: _root_.scala.Predef.String = agentName.getOrElse("")
    def clearAgentName: PasswordChangedEmail = copy(agentName = None)
    def withAgentName(__v: _root_.scala.Predef.String): PasswordChangedEmail = copy(agentName = Option(__v))
    def getLang: _root_.lang.L = lang.getOrElse(_root_.lang.L.en)
    def clearLang: PasswordChangedEmail = copy(lang = None)
    def withLang(__v: _root_.lang.L): PasswordChangedEmail = copy(lang = Option(__v))
    def getFrom: _root_.scala.Predef.String = from.getOrElse("")
    def clearFrom: PasswordChangedEmail = copy(from = None)
    def withFrom(__v: _root_.scala.Predef.String): PasswordChangedEmail = copy(from = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => to.orNull
        case 3 => accountId.orNull
        case 4 => agentId.orNull
        case 5 => agentName.orNull
        case 7 => lang.map(_.javaValueDescriptor).orNull
        case 9 => from.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => to.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => agentId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => agentName.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => lang.map(__e => _root_.scalapb.descriptors.PEnum(__e.scalaValueDescriptor)).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => from.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = account.PasswordChangedEmail
}

object PasswordChangedEmail extends scalapb.GeneratedMessageCompanion[account.PasswordChangedEmail] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[account.PasswordChangedEmail] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): account.PasswordChangedEmail = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    account.PasswordChangedEmail(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.com.google.protobuf.Descriptors.EnumValueDescriptor]].map(__e => lang.L.fromValue(__e.getNumber)),
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[account.PasswordChangedEmail] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      account.PasswordChangedEmail(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scalapb.descriptors.EnumValueDescriptor]]).map(__e => lang.L.fromValue(__e.number)),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AccountProto.javaDescriptor.getMessageTypes.get(17)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AccountProto.scalaDescriptor.messages(17)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = {
    (__fieldNumber: @_root_.scala.unchecked) match {
      case 7 => lang.L
    }
  }
  lazy val defaultInstance = account.PasswordChangedEmail(
  )
  implicit class PasswordChangedEmailLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, account.PasswordChangedEmail]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, account.PasswordChangedEmail](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def to: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getTo)((c_, f_) => c_.copy(to = Option(f_)))
    def optionalTo: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.to)((c_, f_) => c_.copy(to = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def agentId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentId)((c_, f_) => c_.copy(agentId = Option(f_)))
    def optionalAgentId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentId)((c_, f_) => c_.copy(agentId = f_))
    def agentName: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentName)((c_, f_) => c_.copy(agentName = Option(f_)))
    def optionalAgentName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentName)((c_, f_) => c_.copy(agentName = f_))
    def lang: _root_.scalapb.lenses.Lens[UpperPB, _root_.lang.L] = field(_.getLang)((c_, f_) => c_.copy(lang = Option(f_)))
    def optionalLang: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.lang.L]] = field(_.lang)((c_, f_) => c_.copy(lang = f_))
    def from: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getFrom)((c_, f_) => c_.copy(from = Option(f_)))
    def optionalFrom: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.from)((c_, f_) => c_.copy(from = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val TO_FIELD_NUMBER = 2
  final val ACCOUNT_ID_FIELD_NUMBER = 3
  final val AGENT_ID_FIELD_NUMBER = 4
  final val AGENT_NAME_FIELD_NUMBER = 5
  final val LANG_FIELD_NUMBER = 7
  final val FROM_FIELD_NUMBER = 9
}
