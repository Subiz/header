// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package account

@SerialVersionUID(0L)
final case class ConfirmEmail(
    ctx: scala.Option[common.Context] = None,
    to: scala.Option[_root_.scala.Predef.String] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    ownerId: scala.Option[_root_.scala.Predef.String] = None,
    token: scala.Option[_root_.scala.Predef.String] = None,
    expiredIn: scala.Option[_root_.scala.Long] = None,
    accountName: scala.Option[_root_.scala.Predef.String] = None,
    lang: scala.Option[_root_.lang.L] = None,
    ownerName: scala.Option[_root_.scala.Predef.String] = None,
    from: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[ConfirmEmail] with scalapb.lenses.Updatable[ConfirmEmail] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (to.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, to.get) }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, accountId.get) }
      if (ownerId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, ownerId.get) }
      if (token.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, token.get) }
      if (expiredIn.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(6, expiredIn.get) }
      if (accountName.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, accountName.get) }
      if (lang.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeEnumSize(8, lang.get.value) }
      if (ownerName.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, ownerName.get) }
      if (from.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, from.get) }
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
      ownerId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      token.foreach { __v =>
        _output__.writeString(5, __v)
      };
      expiredIn.foreach { __v =>
        _output__.writeInt64(6, __v)
      };
      accountName.foreach { __v =>
        _output__.writeString(7, __v)
      };
      lang.foreach { __v =>
        _output__.writeEnum(8, __v.value)
      };
      ownerName.foreach { __v =>
        _output__.writeString(9, __v)
      };
      from.foreach { __v =>
        _output__.writeString(10, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): account.ConfirmEmail = {
      var __ctx = this.ctx
      var __to = this.to
      var __accountId = this.accountId
      var __ownerId = this.ownerId
      var __token = this.token
      var __expiredIn = this.expiredIn
      var __accountName = this.accountName
      var __lang = this.lang
      var __ownerName = this.ownerName
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
            __ownerId = Option(_input__.readString())
          case 42 =>
            __token = Option(_input__.readString())
          case 48 =>
            __expiredIn = Option(_input__.readInt64())
          case 58 =>
            __accountName = Option(_input__.readString())
          case 64 =>
            __lang = Option(_root_.lang.L.fromValue(_input__.readEnum()))
          case 74 =>
            __ownerName = Option(_input__.readString())
          case 82 =>
            __from = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      account.ConfirmEmail(
          ctx = __ctx,
          to = __to,
          accountId = __accountId,
          ownerId = __ownerId,
          token = __token,
          expiredIn = __expiredIn,
          accountName = __accountName,
          lang = __lang,
          ownerName = __ownerName,
          from = __from
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: ConfirmEmail = copy(ctx = None)
    def withCtx(__v: common.Context): ConfirmEmail = copy(ctx = Option(__v))
    def getTo: _root_.scala.Predef.String = to.getOrElse("")
    def clearTo: ConfirmEmail = copy(to = None)
    def withTo(__v: _root_.scala.Predef.String): ConfirmEmail = copy(to = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: ConfirmEmail = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): ConfirmEmail = copy(accountId = Option(__v))
    def getOwnerId: _root_.scala.Predef.String = ownerId.getOrElse("")
    def clearOwnerId: ConfirmEmail = copy(ownerId = None)
    def withOwnerId(__v: _root_.scala.Predef.String): ConfirmEmail = copy(ownerId = Option(__v))
    def getToken: _root_.scala.Predef.String = token.getOrElse("")
    def clearToken: ConfirmEmail = copy(token = None)
    def withToken(__v: _root_.scala.Predef.String): ConfirmEmail = copy(token = Option(__v))
    def getExpiredIn: _root_.scala.Long = expiredIn.getOrElse(0L)
    def clearExpiredIn: ConfirmEmail = copy(expiredIn = None)
    def withExpiredIn(__v: _root_.scala.Long): ConfirmEmail = copy(expiredIn = Option(__v))
    def getAccountName: _root_.scala.Predef.String = accountName.getOrElse("")
    def clearAccountName: ConfirmEmail = copy(accountName = None)
    def withAccountName(__v: _root_.scala.Predef.String): ConfirmEmail = copy(accountName = Option(__v))
    def getLang: _root_.lang.L = lang.getOrElse(_root_.lang.L.en)
    def clearLang: ConfirmEmail = copy(lang = None)
    def withLang(__v: _root_.lang.L): ConfirmEmail = copy(lang = Option(__v))
    def getOwnerName: _root_.scala.Predef.String = ownerName.getOrElse("")
    def clearOwnerName: ConfirmEmail = copy(ownerName = None)
    def withOwnerName(__v: _root_.scala.Predef.String): ConfirmEmail = copy(ownerName = Option(__v))
    def getFrom: _root_.scala.Predef.String = from.getOrElse("")
    def clearFrom: ConfirmEmail = copy(from = None)
    def withFrom(__v: _root_.scala.Predef.String): ConfirmEmail = copy(from = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => to.orNull
        case 3 => accountId.orNull
        case 4 => ownerId.orNull
        case 5 => token.orNull
        case 6 => expiredIn.orNull
        case 7 => accountName.orNull
        case 8 => lang.map(_.javaValueDescriptor).orNull
        case 9 => ownerName.orNull
        case 10 => from.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => to.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => ownerId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => token.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => expiredIn.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => accountName.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => lang.map(__e => _root_.scalapb.descriptors.PEnum(__e.scalaValueDescriptor)).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => ownerName.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => from.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = account.ConfirmEmail
}

object ConfirmEmail extends scalapb.GeneratedMessageCompanion[account.ConfirmEmail] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[account.ConfirmEmail] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): account.ConfirmEmail = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    account.ConfirmEmail(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.com.google.protobuf.Descriptors.EnumValueDescriptor]].map(__e => lang.L.fromValue(__e.getNumber)),
      __fieldsMap.get(__fields.get(8)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(9)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[account.ConfirmEmail] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      account.ConfirmEmail(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scalapb.descriptors.EnumValueDescriptor]]).map(__e => lang.L.fromValue(__e.number)),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AccountProto.javaDescriptor.getMessageTypes.get(14)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AccountProto.scalaDescriptor.messages(14)
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
      case 8 => lang.L
    }
  }
  lazy val defaultInstance = account.ConfirmEmail(
  )
  implicit class ConfirmEmailLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, account.ConfirmEmail]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, account.ConfirmEmail](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def to: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getTo)((c_, f_) => c_.copy(to = Option(f_)))
    def optionalTo: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.to)((c_, f_) => c_.copy(to = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def ownerId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getOwnerId)((c_, f_) => c_.copy(ownerId = Option(f_)))
    def optionalOwnerId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.ownerId)((c_, f_) => c_.copy(ownerId = f_))
    def token: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getToken)((c_, f_) => c_.copy(token = Option(f_)))
    def optionalToken: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.token)((c_, f_) => c_.copy(token = f_))
    def expiredIn: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getExpiredIn)((c_, f_) => c_.copy(expiredIn = Option(f_)))
    def optionalExpiredIn: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.expiredIn)((c_, f_) => c_.copy(expiredIn = f_))
    def accountName: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountName)((c_, f_) => c_.copy(accountName = Option(f_)))
    def optionalAccountName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountName)((c_, f_) => c_.copy(accountName = f_))
    def lang: _root_.scalapb.lenses.Lens[UpperPB, _root_.lang.L] = field(_.getLang)((c_, f_) => c_.copy(lang = Option(f_)))
    def optionalLang: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.lang.L]] = field(_.lang)((c_, f_) => c_.copy(lang = f_))
    def ownerName: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getOwnerName)((c_, f_) => c_.copy(ownerName = Option(f_)))
    def optionalOwnerName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.ownerName)((c_, f_) => c_.copy(ownerName = f_))
    def from: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getFrom)((c_, f_) => c_.copy(from = Option(f_)))
    def optionalFrom: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.from)((c_, f_) => c_.copy(from = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val TO_FIELD_NUMBER = 2
  final val ACCOUNT_ID_FIELD_NUMBER = 3
  final val OWNER_ID_FIELD_NUMBER = 4
  final val TOKEN_FIELD_NUMBER = 5
  final val EXPIRED_IN_FIELD_NUMBER = 6
  final val ACCOUNT_NAME_FIELD_NUMBER = 7
  final val LANG_FIELD_NUMBER = 8
  final val OWNER_NAME_FIELD_NUMBER = 9
  final val FROM_FIELD_NUMBER = 10
}