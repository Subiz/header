// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

@SerialVersionUID(0L)
final case class Search(
    ctx: scala.Option[common.Context] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    keyword: scala.Option[_root_.scala.Predef.String] = None,
    limit: scala.Option[_root_.scala.Int] = None,
    beforeId: scala.Option[_root_.scala.Predef.String] = None,
    afterId: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Search] with scalapb.lenses.Updatable[Search] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, accountId.get) }
      if (keyword.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, keyword.get) }
      if (limit.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(5, limit.get) }
      if (beforeId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, beforeId.get) }
      if (afterId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, afterId.get) }
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
      accountId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      keyword.foreach { __v =>
        _output__.writeString(4, __v)
      };
      limit.foreach { __v =>
        _output__.writeInt32(5, __v)
      };
      beforeId.foreach { __v =>
        _output__.writeString(6, __v)
      };
      afterId.foreach { __v =>
        _output__.writeString(7, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.Search = {
      var __ctx = this.ctx
      var __accountId = this.accountId
      var __keyword = this.keyword
      var __limit = this.limit
      var __beforeId = this.beforeId
      var __afterId = this.afterId
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 26 =>
            __accountId = Option(_input__.readString())
          case 34 =>
            __keyword = Option(_input__.readString())
          case 40 =>
            __limit = Option(_input__.readInt32())
          case 50 =>
            __beforeId = Option(_input__.readString())
          case 58 =>
            __afterId = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      conversation.Search(
          ctx = __ctx,
          accountId = __accountId,
          keyword = __keyword,
          limit = __limit,
          beforeId = __beforeId,
          afterId = __afterId
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: Search = copy(ctx = None)
    def withCtx(__v: common.Context): Search = copy(ctx = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: Search = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): Search = copy(accountId = Option(__v))
    def getKeyword: _root_.scala.Predef.String = keyword.getOrElse("")
    def clearKeyword: Search = copy(keyword = None)
    def withKeyword(__v: _root_.scala.Predef.String): Search = copy(keyword = Option(__v))
    def getLimit: _root_.scala.Int = limit.getOrElse(0)
    def clearLimit: Search = copy(limit = None)
    def withLimit(__v: _root_.scala.Int): Search = copy(limit = Option(__v))
    def getBeforeId: _root_.scala.Predef.String = beforeId.getOrElse("")
    def clearBeforeId: Search = copy(beforeId = None)
    def withBeforeId(__v: _root_.scala.Predef.String): Search = copy(beforeId = Option(__v))
    def getAfterId: _root_.scala.Predef.String = afterId.getOrElse("")
    def clearAfterId: Search = copy(afterId = None)
    def withAfterId(__v: _root_.scala.Predef.String): Search = copy(afterId = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 3 => accountId.orNull
        case 4 => keyword.orNull
        case 5 => limit.orNull
        case 6 => beforeId.orNull
        case 7 => afterId.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => keyword.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => limit.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => beforeId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => afterId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.Search
}

object Search extends scalapb.GeneratedMessageCompanion[conversation.Search] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.Search] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.Search = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.Search(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Int]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.Search] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.Search(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Int]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(13)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(13)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.Search(
  )
  implicit class SearchLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.Search]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.Search](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def keyword: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getKeyword)((c_, f_) => c_.copy(keyword = Option(f_)))
    def optionalKeyword: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.keyword)((c_, f_) => c_.copy(keyword = f_))
    def limit: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getLimit)((c_, f_) => c_.copy(limit = Option(f_)))
    def optionalLimit: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.limit)((c_, f_) => c_.copy(limit = f_))
    def beforeId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getBeforeId)((c_, f_) => c_.copy(beforeId = Option(f_)))
    def optionalBeforeId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.beforeId)((c_, f_) => c_.copy(beforeId = f_))
    def afterId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAfterId)((c_, f_) => c_.copy(afterId = Option(f_)))
    def optionalAfterId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.afterId)((c_, f_) => c_.copy(afterId = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ACCOUNT_ID_FIELD_NUMBER = 3
  final val KEYWORD_FIELD_NUMBER = 4
  final val LIMIT_FIELD_NUMBER = 5
  final val BEFORE_ID_FIELD_NUMBER = 6
  final val AFTER_ID_FIELD_NUMBER = 7
}
