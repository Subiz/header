// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package account

@SerialVersionUID(0L)
final case class Invitation(
    ctx: scala.Option[common.Context] = None,
    id: scala.Option[_root_.scala.Predef.String] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    fromId: scala.Option[_root_.scala.Predef.String] = None,
    email: scala.Option[_root_.scala.Predef.String] = None,
    agentId: scala.Option[_root_.scala.Predef.String] = None,
    sent: scala.Option[_root_.scala.Long] = None,
    replied: scala.Option[_root_.scala.Long] = None,
    agentFullname: scala.Option[_root_.scala.Predef.String] = None,
    agentJobTitle: scala.Option[_root_.scala.Predef.String] = None,
    token: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Invitation] with scalapb.lenses.Updatable[Invitation] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (id.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, id.get) }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, accountId.get) }
      if (fromId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, fromId.get) }
      if (email.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, email.get) }
      if (agentId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, agentId.get) }
      if (sent.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(7, sent.get) }
      if (replied.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(8, replied.get) }
      if (agentFullname.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, agentFullname.get) }
      if (agentJobTitle.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, agentJobTitle.get) }
      if (token.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, token.get) }
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
      id.foreach { __v =>
        _output__.writeString(2, __v)
      };
      accountId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      fromId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      email.foreach { __v =>
        _output__.writeString(5, __v)
      };
      agentId.foreach { __v =>
        _output__.writeString(6, __v)
      };
      sent.foreach { __v =>
        _output__.writeInt64(7, __v)
      };
      replied.foreach { __v =>
        _output__.writeInt64(8, __v)
      };
      agentFullname.foreach { __v =>
        _output__.writeString(9, __v)
      };
      agentJobTitle.foreach { __v =>
        _output__.writeString(10, __v)
      };
      token.foreach { __v =>
        _output__.writeString(11, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): account.Invitation = {
      var __ctx = this.ctx
      var __id = this.id
      var __accountId = this.accountId
      var __fromId = this.fromId
      var __email = this.email
      var __agentId = this.agentId
      var __sent = this.sent
      var __replied = this.replied
      var __agentFullname = this.agentFullname
      var __agentJobTitle = this.agentJobTitle
      var __token = this.token
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __id = Option(_input__.readString())
          case 26 =>
            __accountId = Option(_input__.readString())
          case 34 =>
            __fromId = Option(_input__.readString())
          case 42 =>
            __email = Option(_input__.readString())
          case 50 =>
            __agentId = Option(_input__.readString())
          case 56 =>
            __sent = Option(_input__.readInt64())
          case 64 =>
            __replied = Option(_input__.readInt64())
          case 74 =>
            __agentFullname = Option(_input__.readString())
          case 82 =>
            __agentJobTitle = Option(_input__.readString())
          case 90 =>
            __token = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      account.Invitation(
          ctx = __ctx,
          id = __id,
          accountId = __accountId,
          fromId = __fromId,
          email = __email,
          agentId = __agentId,
          sent = __sent,
          replied = __replied,
          agentFullname = __agentFullname,
          agentJobTitle = __agentJobTitle,
          token = __token
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: Invitation = copy(ctx = None)
    def withCtx(__v: common.Context): Invitation = copy(ctx = Option(__v))
    def getId: _root_.scala.Predef.String = id.getOrElse("")
    def clearId: Invitation = copy(id = None)
    def withId(__v: _root_.scala.Predef.String): Invitation = copy(id = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: Invitation = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): Invitation = copy(accountId = Option(__v))
    def getFromId: _root_.scala.Predef.String = fromId.getOrElse("")
    def clearFromId: Invitation = copy(fromId = None)
    def withFromId(__v: _root_.scala.Predef.String): Invitation = copy(fromId = Option(__v))
    def getEmail: _root_.scala.Predef.String = email.getOrElse("")
    def clearEmail: Invitation = copy(email = None)
    def withEmail(__v: _root_.scala.Predef.String): Invitation = copy(email = Option(__v))
    def getAgentId: _root_.scala.Predef.String = agentId.getOrElse("")
    def clearAgentId: Invitation = copy(agentId = None)
    def withAgentId(__v: _root_.scala.Predef.String): Invitation = copy(agentId = Option(__v))
    def getSent: _root_.scala.Long = sent.getOrElse(0L)
    def clearSent: Invitation = copy(sent = None)
    def withSent(__v: _root_.scala.Long): Invitation = copy(sent = Option(__v))
    def getReplied: _root_.scala.Long = replied.getOrElse(0L)
    def clearReplied: Invitation = copy(replied = None)
    def withReplied(__v: _root_.scala.Long): Invitation = copy(replied = Option(__v))
    def getAgentFullname: _root_.scala.Predef.String = agentFullname.getOrElse("")
    def clearAgentFullname: Invitation = copy(agentFullname = None)
    def withAgentFullname(__v: _root_.scala.Predef.String): Invitation = copy(agentFullname = Option(__v))
    def getAgentJobTitle: _root_.scala.Predef.String = agentJobTitle.getOrElse("")
    def clearAgentJobTitle: Invitation = copy(agentJobTitle = None)
    def withAgentJobTitle(__v: _root_.scala.Predef.String): Invitation = copy(agentJobTitle = Option(__v))
    def getToken: _root_.scala.Predef.String = token.getOrElse("")
    def clearToken: Invitation = copy(token = None)
    def withToken(__v: _root_.scala.Predef.String): Invitation = copy(token = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => id.orNull
        case 3 => accountId.orNull
        case 4 => fromId.orNull
        case 5 => email.orNull
        case 6 => agentId.orNull
        case 7 => sent.orNull
        case 8 => replied.orNull
        case 9 => agentFullname.orNull
        case 10 => agentJobTitle.orNull
        case 11 => token.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => id.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => fromId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => email.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => agentId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => sent.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => replied.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => agentFullname.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => agentJobTitle.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 11 => token.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = account.Invitation
}

object Invitation extends scalapb.GeneratedMessageCompanion[account.Invitation] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[account.Invitation] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): account.Invitation = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    account.Invitation(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(8)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(9)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(10)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[account.Invitation] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      account.Invitation(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AccountProto.javaDescriptor.getMessageTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AccountProto.scalaDescriptor.messages(1)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = account.Invitation(
  )
  implicit class InvitationLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, account.Invitation]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, account.Invitation](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getId)((c_, f_) => c_.copy(id = Option(f_)))
    def optionalId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.id)((c_, f_) => c_.copy(id = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def fromId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getFromId)((c_, f_) => c_.copy(fromId = Option(f_)))
    def optionalFromId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.fromId)((c_, f_) => c_.copy(fromId = f_))
    def email: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getEmail)((c_, f_) => c_.copy(email = Option(f_)))
    def optionalEmail: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.email)((c_, f_) => c_.copy(email = f_))
    def agentId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentId)((c_, f_) => c_.copy(agentId = Option(f_)))
    def optionalAgentId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentId)((c_, f_) => c_.copy(agentId = f_))
    def sent: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getSent)((c_, f_) => c_.copy(sent = Option(f_)))
    def optionalSent: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.sent)((c_, f_) => c_.copy(sent = f_))
    def replied: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getReplied)((c_, f_) => c_.copy(replied = Option(f_)))
    def optionalReplied: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.replied)((c_, f_) => c_.copy(replied = f_))
    def agentFullname: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentFullname)((c_, f_) => c_.copy(agentFullname = Option(f_)))
    def optionalAgentFullname: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentFullname)((c_, f_) => c_.copy(agentFullname = f_))
    def agentJobTitle: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentJobTitle)((c_, f_) => c_.copy(agentJobTitle = Option(f_)))
    def optionalAgentJobTitle: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentJobTitle)((c_, f_) => c_.copy(agentJobTitle = f_))
    def token: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getToken)((c_, f_) => c_.copy(token = Option(f_)))
    def optionalToken: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.token)((c_, f_) => c_.copy(token = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ID_FIELD_NUMBER = 2
  final val ACCOUNT_ID_FIELD_NUMBER = 3
  final val FROM_ID_FIELD_NUMBER = 4
  final val EMAIL_FIELD_NUMBER = 5
  final val AGENT_ID_FIELD_NUMBER = 6
  final val SENT_FIELD_NUMBER = 7
  final val REPLIED_FIELD_NUMBER = 8
  final val AGENT_FULLNAME_FIELD_NUMBER = 9
  final val AGENT_JOB_TITLE_FIELD_NUMBER = 10
  final val TOKEN_FIELD_NUMBER = 11
}